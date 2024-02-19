// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DecentralizedKVMetaData contains all meta data concerning the DecentralizedKV contract.
var DecentralizedKVMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"kvIdx\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lastKvIdx\",\"type\":\"uint256\"}],\"name\":\"Remove\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxKvSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_storageCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dcfFactor\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"__init_KV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dcfFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"exist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"enumDecentralizedKV.DecodeType\",\"name\":\"decodeType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"off\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kvIndices\",\"type\":\"uint256[]\"}],\"name\":\"getKvMetas\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes24\",\"name\":\"\",\"type\":\"bytes24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastKvIdx\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxKvSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"removeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upfrontPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DecentralizedKVABI is the input ABI used to generate the binding from.
// Deprecated: Use DecentralizedKVMetaData.ABI instead.
var DecentralizedKVABI = DecentralizedKVMetaData.ABI

// DecentralizedKV is an auto generated Go binding around an Ethereum contract.
type DecentralizedKV struct {
	DecentralizedKVCaller     // Read-only binding to the contract
	DecentralizedKVTransactor // Write-only binding to the contract
	DecentralizedKVFilterer   // Log filterer for contract events
}

// DecentralizedKVCaller is an auto generated read-only Go binding around an Ethereum contract.
type DecentralizedKVCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecentralizedKVTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DecentralizedKVTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecentralizedKVFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DecentralizedKVFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DecentralizedKVSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DecentralizedKVSession struct {
	Contract     *DecentralizedKV  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DecentralizedKVCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DecentralizedKVCallerSession struct {
	Contract *DecentralizedKVCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DecentralizedKVTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DecentralizedKVTransactorSession struct {
	Contract     *DecentralizedKVTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DecentralizedKVRaw is an auto generated low-level Go binding around an Ethereum contract.
type DecentralizedKVRaw struct {
	Contract *DecentralizedKV // Generic contract binding to access the raw methods on
}

// DecentralizedKVCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DecentralizedKVCallerRaw struct {
	Contract *DecentralizedKVCaller // Generic read-only contract binding to access the raw methods on
}

// DecentralizedKVTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DecentralizedKVTransactorRaw struct {
	Contract *DecentralizedKVTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDecentralizedKV creates a new instance of DecentralizedKV, bound to a specific deployed contract.
func NewDecentralizedKV(address common.Address, backend bind.ContractBackend) (*DecentralizedKV, error) {
	contract, err := bindDecentralizedKV(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DecentralizedKV{DecentralizedKVCaller: DecentralizedKVCaller{contract: contract}, DecentralizedKVTransactor: DecentralizedKVTransactor{contract: contract}, DecentralizedKVFilterer: DecentralizedKVFilterer{contract: contract}}, nil
}

// NewDecentralizedKVCaller creates a new read-only instance of DecentralizedKV, bound to a specific deployed contract.
func NewDecentralizedKVCaller(address common.Address, caller bind.ContractCaller) (*DecentralizedKVCaller, error) {
	contract, err := bindDecentralizedKV(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DecentralizedKVCaller{contract: contract}, nil
}

// NewDecentralizedKVTransactor creates a new write-only instance of DecentralizedKV, bound to a specific deployed contract.
func NewDecentralizedKVTransactor(address common.Address, transactor bind.ContractTransactor) (*DecentralizedKVTransactor, error) {
	contract, err := bindDecentralizedKV(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DecentralizedKVTransactor{contract: contract}, nil
}

// NewDecentralizedKVFilterer creates a new log filterer instance of DecentralizedKV, bound to a specific deployed contract.
func NewDecentralizedKVFilterer(address common.Address, filterer bind.ContractFilterer) (*DecentralizedKVFilterer, error) {
	contract, err := bindDecentralizedKV(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DecentralizedKVFilterer{contract: contract}, nil
}

// bindDecentralizedKV binds a generic wrapper to an already deployed contract.
func bindDecentralizedKV(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DecentralizedKVMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DecentralizedKV *DecentralizedKVRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DecentralizedKV.Contract.DecentralizedKVCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DecentralizedKV *DecentralizedKVRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecentralizedKV.Contract.DecentralizedKVTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DecentralizedKV *DecentralizedKVRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DecentralizedKV.Contract.DecentralizedKVTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DecentralizedKV *DecentralizedKVCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DecentralizedKV.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DecentralizedKV *DecentralizedKVTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecentralizedKV.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DecentralizedKV *DecentralizedKVTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DecentralizedKV.Contract.contract.Transact(opts, method, params...)
}

// DcfFactor is a free data retrieval call binding the contract method 0xa4a8435e.
//
// Solidity: function dcfFactor() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVCaller) DcfFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DecentralizedKV.contract.Call(opts, &out, "dcfFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DcfFactor is a free data retrieval call binding the contract method 0xa4a8435e.
//
// Solidity: function dcfFactor() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVSession) DcfFactor() (*big.Int, error) {
	return _DecentralizedKV.Contract.DcfFactor(&_DecentralizedKV.CallOpts)
}

// DcfFactor is a free data retrieval call binding the contract method 0xa4a8435e.
//
// Solidity: function dcfFactor() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVCallerSession) DcfFactor() (*big.Int, error) {
	return _DecentralizedKV.Contract.DcfFactor(&_DecentralizedKV.CallOpts)
}

// Exist is a free data retrieval call binding the contract method 0x73e8b3d4.
//
// Solidity: function exist(bytes32 key) view returns(bool)
func (_DecentralizedKV *DecentralizedKVCaller) Exist(opts *bind.CallOpts, key [32]byte) (bool, error) {
	var out []interface{}
	err := _DecentralizedKV.contract.Call(opts, &out, "exist", key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exist is a free data retrieval call binding the contract method 0x73e8b3d4.
//
// Solidity: function exist(bytes32 key) view returns(bool)
func (_DecentralizedKV *DecentralizedKVSession) Exist(key [32]byte) (bool, error) {
	return _DecentralizedKV.Contract.Exist(&_DecentralizedKV.CallOpts, key)
}

// Exist is a free data retrieval call binding the contract method 0x73e8b3d4.
//
// Solidity: function exist(bytes32 key) view returns(bool)
func (_DecentralizedKV *DecentralizedKVCallerSession) Exist(key [32]byte) (bool, error) {
	return _DecentralizedKV.Contract.Exist(&_DecentralizedKV.CallOpts, key)
}

// Get is a free data retrieval call binding the contract method 0xbea94b8b.
//
// Solidity: function get(bytes32 key, uint8 decodeType, uint256 off, uint256 len) view returns(bytes)
func (_DecentralizedKV *DecentralizedKVCaller) Get(opts *bind.CallOpts, key [32]byte, decodeType uint8, off *big.Int, len *big.Int) ([]byte, error) {
	var out []interface{}
	err := _DecentralizedKV.contract.Call(opts, &out, "get", key, decodeType, off, len)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xbea94b8b.
//
// Solidity: function get(bytes32 key, uint8 decodeType, uint256 off, uint256 len) view returns(bytes)
func (_DecentralizedKV *DecentralizedKVSession) Get(key [32]byte, decodeType uint8, off *big.Int, len *big.Int) ([]byte, error) {
	return _DecentralizedKV.Contract.Get(&_DecentralizedKV.CallOpts, key, decodeType, off, len)
}

// Get is a free data retrieval call binding the contract method 0xbea94b8b.
//
// Solidity: function get(bytes32 key, uint8 decodeType, uint256 off, uint256 len) view returns(bytes)
func (_DecentralizedKV *DecentralizedKVCallerSession) Get(key [32]byte, decodeType uint8, off *big.Int, len *big.Int) ([]byte, error) {
	return _DecentralizedKV.Contract.Get(&_DecentralizedKV.CallOpts, key, decodeType, off, len)
}

// GetKvMetas is a free data retrieval call binding the contract method 0xfedee7ea.
//
// Solidity: function getKvMetas(uint256[] kvIndices) view returns(bytes32[])
func (_DecentralizedKV *DecentralizedKVCaller) GetKvMetas(opts *bind.CallOpts, kvIndices []*big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _DecentralizedKV.contract.Call(opts, &out, "getKvMetas", kvIndices)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetKvMetas is a free data retrieval call binding the contract method 0xfedee7ea.
//
// Solidity: function getKvMetas(uint256[] kvIndices) view returns(bytes32[])
func (_DecentralizedKV *DecentralizedKVSession) GetKvMetas(kvIndices []*big.Int) ([][32]byte, error) {
	return _DecentralizedKV.Contract.GetKvMetas(&_DecentralizedKV.CallOpts, kvIndices)
}

// GetKvMetas is a free data retrieval call binding the contract method 0xfedee7ea.
//
// Solidity: function getKvMetas(uint256[] kvIndices) view returns(bytes32[])
func (_DecentralizedKV *DecentralizedKVCallerSession) GetKvMetas(kvIndices []*big.Int) ([][32]byte, error) {
	return _DecentralizedKV.Contract.GetKvMetas(&_DecentralizedKV.CallOpts, kvIndices)
}

// Hash is a free data retrieval call binding the contract method 0xd8389dc5.
//
// Solidity: function hash(bytes32 key) view returns(bytes24)
func (_DecentralizedKV *DecentralizedKVCaller) Hash(opts *bind.CallOpts, key [32]byte) ([24]byte, error) {
	var out []interface{}
	err := _DecentralizedKV.contract.Call(opts, &out, "hash", key)

	if err != nil {
		return *new([24]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0xd8389dc5.
//
// Solidity: function hash(bytes32 key) view returns(bytes24)
func (_DecentralizedKV *DecentralizedKVSession) Hash(key [32]byte) ([24]byte, error) {
	return _DecentralizedKV.Contract.Hash(&_DecentralizedKV.CallOpts, key)
}

// Hash is a free data retrieval call binding the contract method 0xd8389dc5.
//
// Solidity: function hash(bytes32 key) view returns(bytes24)
func (_DecentralizedKV *DecentralizedKVCallerSession) Hash(key [32]byte) ([24]byte, error) {
	return _DecentralizedKV.Contract.Hash(&_DecentralizedKV.CallOpts, key)
}

// LastKvIdx is a free data retrieval call binding the contract method 0x429dd7ad.
//
// Solidity: function lastKvIdx() view returns(uint40)
func (_DecentralizedKV *DecentralizedKVCaller) LastKvIdx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DecentralizedKV.contract.Call(opts, &out, "lastKvIdx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastKvIdx is a free data retrieval call binding the contract method 0x429dd7ad.
//
// Solidity: function lastKvIdx() view returns(uint40)
func (_DecentralizedKV *DecentralizedKVSession) LastKvIdx() (*big.Int, error) {
	return _DecentralizedKV.Contract.LastKvIdx(&_DecentralizedKV.CallOpts)
}

// LastKvIdx is a free data retrieval call binding the contract method 0x429dd7ad.
//
// Solidity: function lastKvIdx() view returns(uint40)
func (_DecentralizedKV *DecentralizedKVCallerSession) LastKvIdx() (*big.Int, error) {
	return _DecentralizedKV.Contract.LastKvIdx(&_DecentralizedKV.CallOpts)
}

// MaxKvSize is a free data retrieval call binding the contract method 0xa097365f.
//
// Solidity: function maxKvSize() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVCaller) MaxKvSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DecentralizedKV.contract.Call(opts, &out, "maxKvSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxKvSize is a free data retrieval call binding the contract method 0xa097365f.
//
// Solidity: function maxKvSize() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVSession) MaxKvSize() (*big.Int, error) {
	return _DecentralizedKV.Contract.MaxKvSize(&_DecentralizedKV.CallOpts)
}

// MaxKvSize is a free data retrieval call binding the contract method 0xa097365f.
//
// Solidity: function maxKvSize() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVCallerSession) MaxKvSize() (*big.Int, error) {
	return _DecentralizedKV.Contract.MaxKvSize(&_DecentralizedKV.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DecentralizedKV *DecentralizedKVCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DecentralizedKV.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DecentralizedKV *DecentralizedKVSession) Owner() (common.Address, error) {
	return _DecentralizedKV.Contract.Owner(&_DecentralizedKV.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DecentralizedKV *DecentralizedKVCallerSession) Owner() (common.Address, error) {
	return _DecentralizedKV.Contract.Owner(&_DecentralizedKV.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0xafd5644d.
//
// Solidity: function size(bytes32 key) view returns(uint256)
func (_DecentralizedKV *DecentralizedKVCaller) Size(opts *bind.CallOpts, key [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DecentralizedKV.contract.Call(opts, &out, "size", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0xafd5644d.
//
// Solidity: function size(bytes32 key) view returns(uint256)
func (_DecentralizedKV *DecentralizedKVSession) Size(key [32]byte) (*big.Int, error) {
	return _DecentralizedKV.Contract.Size(&_DecentralizedKV.CallOpts, key)
}

// Size is a free data retrieval call binding the contract method 0xafd5644d.
//
// Solidity: function size(bytes32 key) view returns(uint256)
func (_DecentralizedKV *DecentralizedKVCallerSession) Size(key [32]byte) (*big.Int, error) {
	return _DecentralizedKV.Contract.Size(&_DecentralizedKV.CallOpts, key)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DecentralizedKV.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVSession) StartTime() (*big.Int, error) {
	return _DecentralizedKV.Contract.StartTime(&_DecentralizedKV.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVCallerSession) StartTime() (*big.Int, error) {
	return _DecentralizedKV.Contract.StartTime(&_DecentralizedKV.CallOpts)
}

// StorageCost is a free data retrieval call binding the contract method 0x3cb2fecc.
//
// Solidity: function storageCost() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVCaller) StorageCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DecentralizedKV.contract.Call(opts, &out, "storageCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StorageCost is a free data retrieval call binding the contract method 0x3cb2fecc.
//
// Solidity: function storageCost() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVSession) StorageCost() (*big.Int, error) {
	return _DecentralizedKV.Contract.StorageCost(&_DecentralizedKV.CallOpts)
}

// StorageCost is a free data retrieval call binding the contract method 0x3cb2fecc.
//
// Solidity: function storageCost() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVCallerSession) StorageCost() (*big.Int, error) {
	return _DecentralizedKV.Contract.StorageCost(&_DecentralizedKV.CallOpts)
}

// UpfrontPayment is a free data retrieval call binding the contract method 0x1ccbc6da.
//
// Solidity: function upfrontPayment() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVCaller) UpfrontPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DecentralizedKV.contract.Call(opts, &out, "upfrontPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpfrontPayment is a free data retrieval call binding the contract method 0x1ccbc6da.
//
// Solidity: function upfrontPayment() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVSession) UpfrontPayment() (*big.Int, error) {
	return _DecentralizedKV.Contract.UpfrontPayment(&_DecentralizedKV.CallOpts)
}

// UpfrontPayment is a free data retrieval call binding the contract method 0x1ccbc6da.
//
// Solidity: function upfrontPayment() view returns(uint256)
func (_DecentralizedKV *DecentralizedKVCallerSession) UpfrontPayment() (*big.Int, error) {
	return _DecentralizedKV.Contract.UpfrontPayment(&_DecentralizedKV.CallOpts)
}

// InitKV is a paid mutator transaction binding the contract method 0x7904f7d9.
//
// Solidity: function __init_KV(uint256 _maxKvSize, uint256 _startTime, uint256 _storageCost, uint256 _dcfFactor, address _owner) returns()
func (_DecentralizedKV *DecentralizedKVTransactor) InitKV(opts *bind.TransactOpts, _maxKvSize *big.Int, _startTime *big.Int, _storageCost *big.Int, _dcfFactor *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _DecentralizedKV.contract.Transact(opts, "__init_KV", _maxKvSize, _startTime, _storageCost, _dcfFactor, _owner)
}

// InitKV is a paid mutator transaction binding the contract method 0x7904f7d9.
//
// Solidity: function __init_KV(uint256 _maxKvSize, uint256 _startTime, uint256 _storageCost, uint256 _dcfFactor, address _owner) returns()
func (_DecentralizedKV *DecentralizedKVSession) InitKV(_maxKvSize *big.Int, _startTime *big.Int, _storageCost *big.Int, _dcfFactor *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _DecentralizedKV.Contract.InitKV(&_DecentralizedKV.TransactOpts, _maxKvSize, _startTime, _storageCost, _dcfFactor, _owner)
}

// InitKV is a paid mutator transaction binding the contract method 0x7904f7d9.
//
// Solidity: function __init_KV(uint256 _maxKvSize, uint256 _startTime, uint256 _storageCost, uint256 _dcfFactor, address _owner) returns()
func (_DecentralizedKV *DecentralizedKVTransactorSession) InitKV(_maxKvSize *big.Int, _startTime *big.Int, _storageCost *big.Int, _dcfFactor *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _DecentralizedKV.Contract.InitKV(&_DecentralizedKV.TransactOpts, _maxKvSize, _startTime, _storageCost, _dcfFactor, _owner)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 key) returns()
func (_DecentralizedKV *DecentralizedKVTransactor) Remove(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _DecentralizedKV.contract.Transact(opts, "remove", key)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 key) returns()
func (_DecentralizedKV *DecentralizedKVSession) Remove(key [32]byte) (*types.Transaction, error) {
	return _DecentralizedKV.Contract.Remove(&_DecentralizedKV.TransactOpts, key)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 key) returns()
func (_DecentralizedKV *DecentralizedKVTransactorSession) Remove(key [32]byte) (*types.Transaction, error) {
	return _DecentralizedKV.Contract.Remove(&_DecentralizedKV.TransactOpts, key)
}

// RemoveTo is a paid mutator transaction binding the contract method 0x49bdd6f5.
//
// Solidity: function removeTo(bytes32 key, address to) returns()
func (_DecentralizedKV *DecentralizedKVTransactor) RemoveTo(opts *bind.TransactOpts, key [32]byte, to common.Address) (*types.Transaction, error) {
	return _DecentralizedKV.contract.Transact(opts, "removeTo", key, to)
}

// RemoveTo is a paid mutator transaction binding the contract method 0x49bdd6f5.
//
// Solidity: function removeTo(bytes32 key, address to) returns()
func (_DecentralizedKV *DecentralizedKVSession) RemoveTo(key [32]byte, to common.Address) (*types.Transaction, error) {
	return _DecentralizedKV.Contract.RemoveTo(&_DecentralizedKV.TransactOpts, key, to)
}

// RemoveTo is a paid mutator transaction binding the contract method 0x49bdd6f5.
//
// Solidity: function removeTo(bytes32 key, address to) returns()
func (_DecentralizedKV *DecentralizedKVTransactorSession) RemoveTo(key [32]byte, to common.Address) (*types.Transaction, error) {
	return _DecentralizedKV.Contract.RemoveTo(&_DecentralizedKV.TransactOpts, key, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DecentralizedKV *DecentralizedKVTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DecentralizedKV.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DecentralizedKV *DecentralizedKVSession) RenounceOwnership() (*types.Transaction, error) {
	return _DecentralizedKV.Contract.RenounceOwnership(&_DecentralizedKV.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DecentralizedKV *DecentralizedKVTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DecentralizedKV.Contract.RenounceOwnership(&_DecentralizedKV.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DecentralizedKV *DecentralizedKVTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DecentralizedKV.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DecentralizedKV *DecentralizedKVSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DecentralizedKV.Contract.TransferOwnership(&_DecentralizedKV.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DecentralizedKV *DecentralizedKVTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DecentralizedKV.Contract.TransferOwnership(&_DecentralizedKV.TransactOpts, newOwner)
}

// DecentralizedKVInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DecentralizedKV contract.
type DecentralizedKVInitializedIterator struct {
	Event *DecentralizedKVInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DecentralizedKVInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecentralizedKVInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DecentralizedKVInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DecentralizedKVInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecentralizedKVInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecentralizedKVInitialized represents a Initialized event raised by the DecentralizedKV contract.
type DecentralizedKVInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DecentralizedKV *DecentralizedKVFilterer) FilterInitialized(opts *bind.FilterOpts) (*DecentralizedKVInitializedIterator, error) {

	logs, sub, err := _DecentralizedKV.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DecentralizedKVInitializedIterator{contract: _DecentralizedKV.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DecentralizedKV *DecentralizedKVFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DecentralizedKVInitialized) (event.Subscription, error) {

	logs, sub, err := _DecentralizedKV.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecentralizedKVInitialized)
				if err := _DecentralizedKV.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DecentralizedKV *DecentralizedKVFilterer) ParseInitialized(log types.Log) (*DecentralizedKVInitialized, error) {
	event := new(DecentralizedKVInitialized)
	if err := _DecentralizedKV.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecentralizedKVOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DecentralizedKV contract.
type DecentralizedKVOwnershipTransferredIterator struct {
	Event *DecentralizedKVOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DecentralizedKVOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecentralizedKVOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DecentralizedKVOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DecentralizedKVOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecentralizedKVOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecentralizedKVOwnershipTransferred represents a OwnershipTransferred event raised by the DecentralizedKV contract.
type DecentralizedKVOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DecentralizedKV *DecentralizedKVFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DecentralizedKVOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DecentralizedKV.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DecentralizedKVOwnershipTransferredIterator{contract: _DecentralizedKV.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DecentralizedKV *DecentralizedKVFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DecentralizedKVOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DecentralizedKV.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecentralizedKVOwnershipTransferred)
				if err := _DecentralizedKV.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DecentralizedKV *DecentralizedKVFilterer) ParseOwnershipTransferred(log types.Log) (*DecentralizedKVOwnershipTransferred, error) {
	event := new(DecentralizedKVOwnershipTransferred)
	if err := _DecentralizedKV.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DecentralizedKVRemoveIterator is returned from FilterRemove and is used to iterate over the raw logs and unpacked data for Remove events raised by the DecentralizedKV contract.
type DecentralizedKVRemoveIterator struct {
	Event *DecentralizedKVRemove // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DecentralizedKVRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DecentralizedKVRemove)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DecentralizedKVRemove)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DecentralizedKVRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DecentralizedKVRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DecentralizedKVRemove represents a Remove event raised by the DecentralizedKV contract.
type DecentralizedKVRemove struct {
	KvIdx     *big.Int
	LastKvIdx *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemove is a free log retrieval operation binding the contract event 0x3fe0d56b27a732f12ef8ae490c8de6bcda1ded03608739ba22b2e63e91407814.
//
// Solidity: event Remove(uint256 indexed kvIdx, uint256 indexed lastKvIdx)
func (_DecentralizedKV *DecentralizedKVFilterer) FilterRemove(opts *bind.FilterOpts, kvIdx []*big.Int, lastKvIdx []*big.Int) (*DecentralizedKVRemoveIterator, error) {

	var kvIdxRule []interface{}
	for _, kvIdxItem := range kvIdx {
		kvIdxRule = append(kvIdxRule, kvIdxItem)
	}
	var lastKvIdxRule []interface{}
	for _, lastKvIdxItem := range lastKvIdx {
		lastKvIdxRule = append(lastKvIdxRule, lastKvIdxItem)
	}

	logs, sub, err := _DecentralizedKV.contract.FilterLogs(opts, "Remove", kvIdxRule, lastKvIdxRule)
	if err != nil {
		return nil, err
	}
	return &DecentralizedKVRemoveIterator{contract: _DecentralizedKV.contract, event: "Remove", logs: logs, sub: sub}, nil
}

// WatchRemove is a free log subscription operation binding the contract event 0x3fe0d56b27a732f12ef8ae490c8de6bcda1ded03608739ba22b2e63e91407814.
//
// Solidity: event Remove(uint256 indexed kvIdx, uint256 indexed lastKvIdx)
func (_DecentralizedKV *DecentralizedKVFilterer) WatchRemove(opts *bind.WatchOpts, sink chan<- *DecentralizedKVRemove, kvIdx []*big.Int, lastKvIdx []*big.Int) (event.Subscription, error) {

	var kvIdxRule []interface{}
	for _, kvIdxItem := range kvIdx {
		kvIdxRule = append(kvIdxRule, kvIdxItem)
	}
	var lastKvIdxRule []interface{}
	for _, lastKvIdxItem := range lastKvIdx {
		lastKvIdxRule = append(lastKvIdxRule, lastKvIdxItem)
	}

	logs, sub, err := _DecentralizedKV.contract.WatchLogs(opts, "Remove", kvIdxRule, lastKvIdxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DecentralizedKVRemove)
				if err := _DecentralizedKV.contract.UnpackLog(event, "Remove", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemove is a log parse operation binding the contract event 0x3fe0d56b27a732f12ef8ae490c8de6bcda1ded03608739ba22b2e63e91407814.
//
// Solidity: event Remove(uint256 indexed kvIdx, uint256 indexed lastKvIdx)
func (_DecentralizedKV *DecentralizedKVFilterer) ParseRemove(log types.Log) (*DecentralizedKVRemove, error) {
	event := new(DecentralizedKVRemove)
	if err := _DecentralizedKV.contract.UnpackLog(event, "Remove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

