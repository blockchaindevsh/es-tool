package cmd

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	"github.com/urfave/cli"
	"github.com/zhiqiangxu/es-tool/cmd/flag"
	"github.com/zhiqiangxu/es-tool/internal/abi"
)

// OPCmd ...
var OPCmd = cli.Command{
	Name:  "op",
	Usage: "op actions",
	Subcommands: []cli.Command{
		opDeployBatchInboxCmd,
		opPutGetCmd,
		opGetCmd,
	},
}

var opDeployBatchInboxCmd = cli.Command{
	Name:  "deploy_batchinbox",
	Usage: "deploy batchinbox contract",
	Flags: []cli.Flag{
		flag.NetworkFlag,
		flag.PKFlag,
	},
	Action: opDeployBatchInbox,
}

var opPutGetCmd = cli.Command{
	Name:  "putget",
	Usage: "put a blob and fetch it",
	Flags: []cli.Flag{
		flag.NetworkFlag,
		flag.PKFlag,
		flag.ContractFlag,
	},
	Action: opPutGet,
}

var opGetCmd = cli.Command{
	Name:  "get",
	Usage: "get a blob",
	Flags: []cli.Flag{
		flag.ContractFlag,
		flag.BlobFlag,
	},
	Action: opGet,
}

func opDeployBatchInbox(ctx *cli.Context) (err error) {
	client, err := ethclient.Dial(ctx.String(flag.NetworkFlag.Name))
	if err != nil {
		return
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return
	}

	privateKey, err := crypto.HexToECDSA(ctx.String(flag.PKFlag.Name))
	if err != nil {
		return
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = defaultGasLimit
	auth.GasPrice = gasPrice

	esAddr := common.HexToAddress("0x804C520d3c084C805E37A35E90057Ac32831F96f")
	inboxAddr, tx, _, err := abi.DeployBatchInbox(auth, client, esAddr)
	if err != nil {
		return
	}

	fmt.Println("inbox addr", inboxAddr, "deploy tx", tx.Hash())

	return
}

const (
	defaultGasLimit = 5000000
)

func opPutGet(ctx *cli.Context) (err error) {
	inboxAddr := common.HexToAddress(ctx.String(flag.ContractFlag.Name))

	// put to sepolia

	client, err := ethclient.Dial(ctx.String(flag.NetworkFlag.Name))
	if err != nil {
		return
	}
	esAddr := common.HexToAddress("0x804C520d3c084C805E37A35E90057Ac32831F96f")
	kv, err := abi.NewDecentralizedKVCaller(esAddr, client)
	if err != nil {
		return
	}
	pay, err := kv.UpfrontPayment(nil)
	if err != nil {
		return
	}

	randomBytes := make([]byte, 20)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return
	}
	var b eth.Blob
	if err = b.FromData(randomBytes); err != nil {
		return
	}
	candidate := &txmgr.TxCandidate{
		To:       &inboxAddr,
		Blobs:    []*eth.Blob{&b},
		GasLimit: defaultGasLimit,
		Value:    pay,
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return
	}

	privateKey, err := crypto.HexToECDSA(ctx.String(flag.PKFlag.Name))
	if err != nil {
		return
	}
	tx, err := craftTx(privateKey, client, candidate, chainID)
	if err != nil {
		return
	}
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		return
	}
	fmt.Println("tx", tx.Hash(), "blob hashes", tx.BlobHashes())

	// get from es-geth
	esGeth := "http://65.108.236.27:9540"
	esClient, err := sources.NewESClient(esGeth, inboxAddr, nil)
	if err != nil {
		return
	}

	var (
		blobHashes   = []eth.IndexedBlobHash{eth.IndexedBlobHash{Hash: tx.BlobHashes()[0]}}
		blobSidecars []*eth.BlobSidecar
	)

	start := time.Now()

	for {
		blobSidecars, err = esClient.GetBlobs(blobHashes)
		if err != nil {
			fmt.Println("GetBlobs err, sleep:", err)
			time.Sleep(time.Second)
			continue
		}
		if len(blobSidecars) != 1 {
			panic(fmt.Sprintf("wrong #blobSidecars:%d", len(blobSidecars)))
		}
		if blobSidecars[0].Blob != b {
			panic("blob not match")
		}

		fmt.Println("found blob hash from es", tx.BlobHashes()[0], "elapsed", time.Since(start))
		break
	}

	return
}

func opGet(ctx *cli.Context) (err error) {
	inboxAddr := common.HexToAddress(ctx.String(flag.ContractFlag.Name))
	blobHash := common.HexToHash(ctx.String(flag.BlobFlag.Name))

	// get from es-geth
	esGeth := "http://65.108.236.27:9540"
	esClient, err := sources.NewESClient(esGeth, inboxAddr, nil)
	if err != nil {
		return
	}
	var (
		blobHashes   = []eth.IndexedBlobHash{eth.IndexedBlobHash{Hash: blobHash}}
		blobSidecars []*eth.BlobSidecar
	)

	for {
		blobSidecars, err = esClient.GetBlobs(blobHashes)
		if err != nil {
			fmt.Println("GetBlobs err, sleep:", err)
			time.Sleep(time.Second)
			continue
		}
		if len(blobSidecars) != 1 {
			panic(fmt.Sprintf("wrong #blobSidecars:%d", len(blobSidecars)))
		}
		if calcBlobHash(blobSidecars[0].Blob) != blobHash {
			panic("invalid blob hash")
		}
		fmt.Println("blob", hex.EncodeToString(blobSidecars[0].Blob[:]))
		break
	}

	return
}

func calcBlobHash(blob eth.Blob) common.Hash {
	rawBlob := *blob.KZGBlob()
	commitment, err := kzg4844.BlobToCommitment(rawBlob)
	if err != nil {
		panic(fmt.Errorf("cannot compute KZG commitment of blob %d in tx candidate: %w", err))
	}

	return eth.KZGToVersionedHash(commitment)
}

func craftTx(privateKey *ecdsa.PrivateKey, client *ethclient.Client, candidate *txmgr.TxCandidate, chainID *big.Int) (*types.Transaction, error) {

	gasTipCap, baseFee, blobBaseFee, err := suggestGasPriceCaps(client)
	if err != nil {
		return nil, err
	}
	gasFeeCap := calcGasFeeCap(baseFee, gasTipCap)

	var sidecar *types.BlobTxSidecar
	var blobHashes []common.Hash
	if len(candidate.Blobs) > 0 {
		if candidate.To == nil {
			return nil, errors.New("blob txs cannot deploy contracts")
		}
		if sidecar, blobHashes, err = txmgr.MakeSidecar(candidate.Blobs); err != nil {
			return nil, fmt.Errorf("failed to make sidecar: %w", err)
		}
	}

	var txMessage types.TxData
	if sidecar != nil {
		if blobBaseFee == nil {
			return nil, fmt.Errorf("expected non-nil blobBaseFee")
		}
		blobFeeCap := calcBlobFeeCap(blobBaseFee)
		message := &types.BlobTx{
			To:         *candidate.To,
			Data:       candidate.TxData,
			Gas:        candidate.GasLimit,
			BlobHashes: blobHashes,
			Sidecar:    sidecar,
		}
		if err := finishBlobTx(message, chainID, gasTipCap, gasFeeCap, blobFeeCap, candidate.Value); err != nil {
			return nil, fmt.Errorf("failed to create blob transaction: %w", err)
		}
		txMessage = message
	} else {
		panic("unsupported")
	}
	return signWithNextNonce(privateKey, client, txMessage, chainID) // signer sets the nonce field of the tx

}

func signWithNextNonce(privateKey *ecdsa.PrivateKey, client *ethclient.Client, txMessage types.TxData, chainID *big.Int) (*types.Transaction, error) {

	nonce, err := client.NonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey), nil)
	if err != nil {
		return nil, err
	}

	switch x := txMessage.(type) {
	case *types.BlobTx:
		x.Nonce = nonce
	default:
		return nil, fmt.Errorf("unrecognized tx type: %T", x)
	}

	signer := types.LatestSignerForChainID(chainID)
	tx := types.NewTx(txMessage)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(signer, signature)
}

func suggestGasPriceCaps(client *ethclient.Client) (*big.Int, *big.Int, *big.Int, error) {

	tip, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, nil, nil, err
	} else if tip == nil {
		return nil, nil, nil, errors.New("the suggested tip was nil")
	}
	head, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, nil, nil, err
	} else if head.BaseFee == nil {
		return nil, nil, nil, errors.New("txmgr does not support pre-london blocks that do not have a base fee")
	}

	baseFee := head.BaseFee

	var blobFee *big.Int
	if head.ExcessBlobGas != nil {
		blobFee = eip4844.CalcBlobFee(*head.ExcessBlobGas)
	}
	return tip, baseFee, blobFee, nil
}

func calcGasFeeCap(baseFee, gasTipCap *big.Int) *big.Int {
	return new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(baseFee, two),
	)
}

func calcBlobFeeCap(blobBaseFee *big.Int) *big.Int {
	cap := new(big.Int).Mul(blobBaseFee, two)
	if cap.Cmp(minBlobTxFee) < 0 {
		cap.Set(minBlobTxFee)
	}
	return cap
}

func finishBlobTx(message *types.BlobTx, chainID, tip, fee, blobFee, value *big.Int) error {
	var o bool
	if message.ChainID, o = uint256.FromBig(chainID); o {
		return fmt.Errorf("ChainID overflow")
	}
	if message.GasTipCap, o = uint256.FromBig(tip); o {
		return fmt.Errorf("GasTipCap overflow")
	}
	if message.GasFeeCap, o = uint256.FromBig(fee); o {
		return fmt.Errorf("GasFeeCap overflow")
	}
	if message.BlobFeeCap, o = uint256.FromBig(blobFee); o {
		return fmt.Errorf("BlobFeeCap overflow")
	}
	if message.Value, o = uint256.FromBig(value); o {
		return fmt.Errorf("Value overflow")
	}
	return nil
}

var (
	two          = big.NewInt(2)
	minBlobTxFee = big.NewInt(params.GWei)
)
