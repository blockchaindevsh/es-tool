package cmd

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	mrand "math/rand"
	"os"
	"time"

	"github.com/ethereum-optimism/optimism/op-batcher/compressor"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum-optimism/optimism/op-service/testutils"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
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
		opBatchRatioCmd,
		opCompressCmd,
		opEstimateGasCmd,
		opDecodeBlobCmd,
		opBlobIndexCmd,
		opTxSizeCmd,
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

var opBatchRatioCmd = cli.Command{
	Name:  "batch_ratio",
	Usage: "compute batch ratio",
	Flags: []cli.Flag{
		flag.SpanFlag,
		flag.TPSFlag,
		flag.TxDataSizeFlag,
	},
	Action: opBatchRatio,
}

var opCompressCmd = cli.Command{
	Name:  "compress",
	Usage: "do comprest test",
	Flags: []cli.Flag{
		flag.TxDataSizeFlag,
	},
	Action: opCompress,
}

var opEstimateGasCmd = cli.Command{
	Name:  "estimate_gas",
	Usage: "do gas estimation",
	Flags: []cli.Flag{
		flag.TPSFlag,
		flag.ESInboxFlag,
		flag.SpanFlag,
		flag.DailyProposeTimesFlag,
		flag.BlobBaseFeeFlag,
		flag.TxDataSizeFlag,
		flag.TxBlobsFlag,
	},
	Action: opEstimateGas,
}

var opDecodeBlobCmd = cli.Command{
	Name:  "decode_blob",
	Usage: "decode batcher blob",
	Flags: []cli.Flag{
		flag.BlobFileFlag,
	},
	Action: opDecodeBlob,
}

var opBlobIndexCmd = cli.Command{
	Name:  "blob_index",
	Usage: "compute blob index by tx hash",
	Flags: []cli.Flag{
		flag.TxFlag,
		flag.NetworkFlag,
	},
	Action: opBlobIndex,
}

var opTxSizeCmd = cli.Command{
	Name:   "tx_size",
	Usage:  "show random tx size",
	Flags:  []cli.Flag{flag.TxDataSizeFlag},
	Action: opTxSize,
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
		blobHashes   = []eth.IndexedBlobHash{{Hash: tx.BlobHashes()[0]}}
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
		blobHashes   = []eth.IndexedBlobHash{{Hash: blobHash}}
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
		panic(fmt.Errorf("cannot compute KZG commitment of blob in tx candidate: %w", err))
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
		return fmt.Errorf("value overflow")
	}
	return nil
}

var (
	two          = big.NewInt(2)
	minBlobTxFee = big.NewInt(params.GWei)
)

func opBatchRatio(ctx *cli.Context) (err error) {

	chainID := big.NewInt(50)
	rng := mrand.New(mrand.NewSource(0x5432177))

	var spanBatchBuilder *derive.SpanBatchBuilder
	spanType := derive.SingularBatchType
	if ctx.Bool(flag.SpanFlag.Name) {
		spanBatchBuilder = derive.NewSpanBatchBuilder(uint64(time.Now().Unix()), chainID)
		spanType = derive.SpanBatchType
	}

	// should keep consistent with https://github.com/ethereum-optimism/optimism/blob/c87a469d7d679e8a4efbace56c3646b925bcc009/op-batcher/compressor/cli.go#L71
	cliConfig := compressor.CLIConfig{Kind: compressor.ShadowKind, TargetL1TxSizeBytes: 100_000, TargetNumFrames: 1, ApproxComprRatio: 0.4}
	c, err := compressor.NewShadowCompressor(cliConfig.Config())
	if err != nil {
		return
	}
	co, err := derive.NewChannelOut(uint(spanType), c, spanBatchBuilder)
	if err != nil {
		return
	}

	singularBatches := RandomValidConsecutiveSingularBatches(rng, chainID, 200000, ctx.Int(flag.TPSFlag.Name)*2, ctx.Int(flag.TxDataSizeFlag.Name))
	fmt.Println("#singularBatches", len(singularBatches))
	full := len(singularBatches) - 1
	for i, singularBatch := range singularBatches {
		if i > 1800 {
			// assume we submit batches every 1800 L2 blocks, which is 1 hour
			full = i
			// force close the channel for submission
			err = co.Close()
			if err != nil {
				return
			}
			break
		}
		_, err = co.AddSingularBatch(singularBatch, uint64(i))
		if err == nil {
			continue
		}

		if errors.Is(err, derive.ErrTooManyRLPBytes) || errors.Is(err, derive.CompressorFullErr) {
			fmt.Println(err)
			full = i
			break
		}
		return
	}

	fmt.Println("added batches", full+1)

	uncompressedBytes := 0
	for i := 0; i < full; i++ {
		var buf bytes.Buffer
		if err = rlp.Encode(&buf, derive.NewBatchData(singularBatches[i])); err != nil {
			return
		}
		uncompressedBytes += buf.Len()
	}
	fmt.Println("compressed bytes", co.ReadyBytes() /* "\ninput bytes", co.InputBytes(), */, "\nuncompressed bytes", uncompressedBytes, "\nratio", float32(co.ReadyBytes())/float32(uncompressedBytes))

	return
}

func RandomValidConsecutiveSingularBatches(rng *mrand.Rand, chainID *big.Int, blockCount, txCount, txDataSize int) []*derive.SingularBatch {
	l2BlockTime := uint64(2)

	var singularBatches []*derive.SingularBatch
	for i := 0; i < blockCount; i++ {
		singularBatch := RandomSingularBatch(rng, txCount, chainID, txDataSize)
		singularBatches = append(singularBatches, singularBatch)
	}
	l1BlockNum := rng.Uint64()
	// make sure oldest timestamp is large enough
	singularBatches[0].Timestamp += 256
	for i := 0; i < blockCount; i++ {
		originChangedBit := rng.Intn(2)
		if originChangedBit == 1 {
			l1BlockNum++
			singularBatches[i].EpochHash = testutils.RandomHash(rng)
		} else if i > 0 {
			singularBatches[i].EpochHash = singularBatches[i-1].EpochHash
		}
		singularBatches[i].EpochNum = rollup.Epoch(l1BlockNum)
		if i > 0 {
			singularBatches[i].Timestamp = singularBatches[i-1].Timestamp + l2BlockTime
		}
	}
	return singularBatches
}

func randomTxNonAccesslist(rng *mrand.Rand, baseFee *big.Int, signer types.Signer, size *int) *types.Transaction {
	var tx *types.Transaction
	for {
		if size == nil {
			tx = testutils.RandomTx(rng, baseFee, signer)
		} else {
			tx = testutils.RandomTxWithSize(rng, baseFee, signer, *size)
		}
		if tx.Type() != types.AccessListTxType {
			return tx
		}
	}
}

func RandomSingularBatch(rng *mrand.Rand, txCount int, chainID *big.Int, txDataSize int) *derive.SingularBatch {
	signer := types.NewLondonSigner(chainID)
	baseFee := big.NewInt(rng.Int63n(300_000_000_000))
	txsEncoded := make([]hexutil.Bytes, 0, txCount)
	// force each tx to have equal chainID
	for i := 0; i < txCount; i++ {
		tx := randomTxNonAccesslist(rng, baseFee, signer, &txDataSize)
		txEncoded, err := tx.MarshalBinary()
		if err != nil {
			panic("tx Marshal binary" + err.Error())
		}
		txsEncoded = append(txsEncoded, hexutil.Bytes(txEncoded))
	}
	return &derive.SingularBatch{
		ParentHash:   testutils.RandomHash(rng),
		EpochNum:     rollup.Epoch(1 + rng.Int63n(100_000_000)),
		EpochHash:    testutils.RandomHash(rng),
		Timestamp:    uint64(rng.Int63n(2_000_000_000)),
		Transactions: txsEncoded,
	}
}

func opCompress(ctx *cli.Context) (err error) {

	chainID := big.NewInt(50)
	rng := mrand.New(mrand.NewSource(time.Now().Unix()))

	// should keep consistent with https://github.com/ethereum-optimism/optimism/blob/c87a469d7d679e8a4efbace56c3646b925bcc009/op-batcher/compressor/cli.go#L71
	cliConfig := compressor.CLIConfig{Kind: compressor.ShadowKind, TargetL1TxSizeBytes: 100_000, TargetNumFrames: 1, ApproxComprRatio: 0.4}
	c, err := compressor.NewShadowCompressor(cliConfig.Config())
	if err != nil {
		return
	}

	singularBatch := RandomSingularBatch(rng, 10, chainID, ctx.Int(flag.TxDataSizeFlag.Name))
	var buf bytes.Buffer
	if err = rlp.Encode(&buf, derive.NewBatchData(singularBatch)); err != nil {
		return
	}

	_, err = c.Write(buf.Bytes())
	if err != nil {
		return
	}

	err = c.Flush()
	if err != nil {
		return
	}

	fmt.Println("uncompressed", buf.Len(), "compressed", c.Len())
	return
}

func opEstimateGas(ctx *cli.Context) (err error) {
	tps := ctx.Int(flag.TPSFlag.Name)
	txCount := 2 * tps

	chainID := big.NewInt(50)
	rng := mrand.New(mrand.NewSource(0x5432177))
	txDataSize := ctx.Int(flag.TxDataSizeFlag.Name)
	blobsPerTx := ctx.Int(flag.TxBlobsFlag.Name)

	singularBatch := RandomSingularBatch(rng, txCount, chainID, txDataSize)

	var buf bytes.Buffer
	if err = rlp.Encode(&buf, derive.NewBatchData(singularBatch)); err != nil {
		return
	}
	singularBatchSize := buf.Len()
	dailyBytes := float64(singularBatchSize * 24 * 1800)

	// we ignore the compressing ratio since it seems not so effective
	if tps == 0 {
		if ctx.Bool(flag.SpanFlag.Name) {
			dailyBytes *= 0.002
		} else {
			dailyBytes *= 0.63
		}
	}

	// assume frame size is MaxBlobDataSize
	dailyBlobTx := int64(dailyBytes)/int64(blobsPerTx*eth.MaxBlobDataSize) + 1

	callDataGas := int64(21000)
	esInbox := ctx.Bool(flag.ESInboxFlag.Name)
	if esInbox {
		callDataGas = 117_258                      // FYI https://sepolia.etherscan.io/tx/0xed09f77fbd3cb87874d3ea06ec7bb84e784095ac2cbdb44a484f6ee5532d732d
		callDataGas += int64(blobsPerTx-1) * 50000 // 每多一个～50000左右gas成本
	}
	fmt.Println("######Batcher#######")
	if esInbox {
		fmt.Printf(
			`
singularBatchSize:	%d
batcher daily tx:	%d ( ~ singularBatchSize * 24 * 1800 / tx_blobs / MaxBlobDataSize + 1 )
batcher per tx gas:	(%d + 50000*tx_blobs)*base_fee + %d*blob_base_fee*tx_blobs
batcher daily gas:	(%d + %d*tx_blobs)*base_fee + %d*blob_base_fee*tx_blobs
`,
			singularBatchSize,
			dailyBlobTx,
			117_258-50000, params.BlobTxBlobGasPerBlob,
			dailyBlobTx*(117_258-50000), 50000*dailyBlobTx, dailyBlobTx*params.BlobTxBlobGasPerBlob)
	} else {
		fmt.Printf(
			`
singularBatchSize:	%d
batcher daily tx:	%d ( ~ singularBatchSize * 24 * 1800 / tx_blobs / MaxBlobDataSize + 1 )
batcher per tx gas:	21000*base_fee + %d*blob_base_fee*tx_blobs
batcher daily gas:	%d*base_fee + %d*blob_base_fee*tx_blobs
`,
			singularBatchSize,
			dailyBlobTx,
			params.BlobTxBlobGasPerBlob,
			dailyBlobTx*21000, dailyBlobTx*params.BlobTxBlobGasPerBlob)
	}

	drawLine := func() {
		fmt.Println("------------------------------------------------")
	}
	drawSharp := func() {
		fmt.Println("\n###########################")
	}
	drawLine()
	blobBaseFee := ctx.Int64(flag.BlobBaseFeeFlag.Name)
	fmt.Printf(`basefee		daily gas/eth (blob_base_fee = %d)
`, blobBaseFee)
	baseFees := []uint{20, 30, 40}
	batcherDailyCost := func(baseFee uint) *big.Int {
		dailyCost := new(big.Int).Add(
			new(big.Int).Mul(big.NewInt(int64(dailyBlobTx*callDataGas)), big.NewInt(int64(baseFee))),
			new(big.Int).Mul(big.NewInt(int64(dailyBlobTx*params.BlobTxBlobGasPerBlob)*int64(blobsPerTx)), big.NewInt(blobBaseFee)),
		)
		return dailyCost
	}
	for _, baseFee := range baseFees {
		dailyCost := batcherDailyCost(baseFee)
		dailyCostFloat, _ := dailyCost.Float64()
		fmt.Printf("%dGwei\t\t%.10f\n", baseFee, dailyCostFloat/1e9)
	}
	drawSharp()
	fmt.Println("\n######Proposer#######")

	proposeGas := int64(87789)
	dailyPropose := ctx.Int64(flag.DailyProposeTimesFlag.Name)
	fmt.Printf(`
proposer daily tx:	%d
proposer per tx gas:	%d*base_fee
proposer daily gas:	%d*base_fee
`, dailyPropose, proposeGas, proposeGas*dailyPropose)
	drawLine()
	fmt.Printf("basefee\t\tdaily gas/eth\n")
	proposerDailyCost := func(baseFee uint) *big.Int {
		dailyCost := new(big.Int).Mul(big.NewInt(proposeGas*dailyPropose), big.NewInt(int64(baseFee)))
		return dailyCost
	}
	for _, baseFee := range baseFees {
		dailyCost := proposerDailyCost(baseFee)
		dailyCostFloat, _ := dailyCost.Float64()
		fmt.Printf("%dGwei\t\t%.10f\n", baseFee, dailyCostFloat/1e9)
	}

	drawSharp()
	fmt.Println("\n######Average l2 tx cost#######")
	dailyTxCount := tps * 24 * 3600
	fmt.Printf(`
dailyTxCount:	%d ( tps * 24 * 3600)
average cost:	(batcherDailyCost + proposerDailyCost)/dailyTxCount
`, dailyTxCount)
	fmt.Printf("\nbasefee\t\tcost/eth\n")
	for _, baseFee := range baseFees {
		batcherDailyCost := batcherDailyCost(baseFee)
		batcherDailyCostFloat, _ := batcherDailyCost.Float64()
		proposerDailyCost := proposerDailyCost(baseFee)
		proposerDailyCostFloat, _ := proposerDailyCost.Float64()
		fmt.Printf("%dGwei\t\t%.10f\n", baseFee, (batcherDailyCostFloat+proposerDailyCostFloat)/1e9/float64(dailyTxCount))
	}
	drawSharp()
	return
}

func opDecodeBlob(ctx *cli.Context) (err error) {

	blobBytes, err := os.ReadFile(ctx.String(flag.BlobFileFlag.Name))
	if err != nil {
		return
	}
	var blob eth.Blob
	err = blob.UnmarshalText(blobBytes)
	if err != nil {
		return
	}

	data, err := blob.ToData()
	if err != nil {
		return
	}

	fmt.Println("blob data size", len(data))

	return
}

func opBlobIndex(ctx *cli.Context) (err error) {
	client, err := ethclient.Dial(ctx.String(flag.NetworkFlag.Name))
	if err != nil {
		return
	}

	txHash := common.HexToHash(ctx.String(flag.TxFlag.Name))
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		return
	}
	if len(tx.BlobHashes()) == 0 {
		err = fmt.Errorf("tx has no blob")
		return
	}

	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return
	}

	block, err := client.BlockByHash(context.Background(), receipt.BlockHash)
	if err != nil {
		return
	}

	startBlobIndex := 0
	for i, transaction := range block.Transactions() {
		if i == int(receipt.TransactionIndex) {
			break
		}
		startBlobIndex += len(transaction.BlobHashes())
	}

	if len(tx.BlobHashes()) == 1 {
		fmt.Println(startBlobIndex)
	} else {
		fmt.Printf("%d-%d\n", startBlobIndex, startBlobIndex+len(tx.BlobHashes())-1)
	}

	return
}

func opTxSize(ctx *cli.Context) (err error) {
	chainID := big.NewInt(50)
	rng := mrand.New(mrand.NewSource(0x5432177))
	signer := types.NewLondonSigner(chainID)
	baseFee := big.NewInt(rng.Int63n(300_000_000_000))
	txDataSize := ctx.Int(flag.TxDataSizeFlag.Name)

	tx := randomTxNonAccesslist(rng, baseFee, signer, &txDataSize)
	txEncoded, err := tx.MarshalBinary()
	if err != nil {
		return
	}

	fmt.Println("encoded", len(txEncoded), "data", len(tx.Data()), "type", tx.Type())
	return
}
