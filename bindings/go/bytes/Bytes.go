// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bytes

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

// BytesMetaData contains all meta data concerning the Bytes contract.
var BytesMetaData = &bind.MetaData{
	ABI: "[]",
}

// BytesABI is the input ABI used to generate the binding from.
// Deprecated: Use BytesMetaData.ABI instead.
var BytesABI = BytesMetaData.ABI

// Bytes is an auto generated Go binding around an Ethereum contract.
type Bytes struct {
	BytesCaller     // Read-only binding to the contract
	BytesTransactor // Write-only binding to the contract
	BytesFilterer   // Log filterer for contract events
}

// BytesCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesSession struct {
	Contract     *Bytes            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesCallerSession struct {
	Contract *BytesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BytesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesTransactorSession struct {
	Contract     *BytesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesRaw struct {
	Contract *Bytes // Generic contract binding to access the raw methods on
}

// BytesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesCallerRaw struct {
	Contract *BytesCaller // Generic read-only contract binding to access the raw methods on
}

// BytesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesTransactorRaw struct {
	Contract *BytesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytes creates a new instance of Bytes, bound to a specific deployed contract.
func NewBytes(address common.Address, backend bind.ContractBackend) (*Bytes, error) {
	contract, err := bindBytes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bytes{BytesCaller: BytesCaller{contract: contract}, BytesTransactor: BytesTransactor{contract: contract}, BytesFilterer: BytesFilterer{contract: contract}}, nil
}

// NewBytesCaller creates a new read-only instance of Bytes, bound to a specific deployed contract.
func NewBytesCaller(address common.Address, caller bind.ContractCaller) (*BytesCaller, error) {
	contract, err := bindBytes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesCaller{contract: contract}, nil
}

// NewBytesTransactor creates a new write-only instance of Bytes, bound to a specific deployed contract.
func NewBytesTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesTransactor, error) {
	contract, err := bindBytes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesTransactor{contract: contract}, nil
}

// NewBytesFilterer creates a new log filterer instance of Bytes, bound to a specific deployed contract.
func NewBytesFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesFilterer, error) {
	contract, err := bindBytes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesFilterer{contract: contract}, nil
}

// bindBytes binds a generic wrapper to an already deployed contract.
func bindBytes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BytesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bytes *BytesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bytes.Contract.BytesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bytes *BytesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bytes.Contract.BytesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bytes *BytesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bytes.Contract.BytesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bytes *BytesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bytes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bytes *BytesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bytes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bytes *BytesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bytes.Contract.contract.Transact(opts, method, params...)
}
