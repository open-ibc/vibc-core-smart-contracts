// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package protobufruntime

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

// ProtoBufRuntimeMetaData contains all meta data concerning the ProtoBufRuntime contract.
var ProtoBufRuntimeMetaData = &bind.MetaData{
	ABI: "[]",
}

// ProtoBufRuntimeABI is the input ABI used to generate the binding from.
// Deprecated: Use ProtoBufRuntimeMetaData.ABI instead.
var ProtoBufRuntimeABI = ProtoBufRuntimeMetaData.ABI

// ProtoBufRuntime is an auto generated Go binding around an Ethereum contract.
type ProtoBufRuntime struct {
	ProtoBufRuntimeCaller     // Read-only binding to the contract
	ProtoBufRuntimeTransactor // Write-only binding to the contract
	ProtoBufRuntimeFilterer   // Log filterer for contract events
}

// ProtoBufRuntimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtoBufRuntimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoBufRuntimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtoBufRuntimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoBufRuntimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtoBufRuntimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoBufRuntimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtoBufRuntimeSession struct {
	Contract     *ProtoBufRuntime  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtoBufRuntimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtoBufRuntimeCallerSession struct {
	Contract *ProtoBufRuntimeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProtoBufRuntimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtoBufRuntimeTransactorSession struct {
	Contract     *ProtoBufRuntimeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProtoBufRuntimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtoBufRuntimeRaw struct {
	Contract *ProtoBufRuntime // Generic contract binding to access the raw methods on
}

// ProtoBufRuntimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtoBufRuntimeCallerRaw struct {
	Contract *ProtoBufRuntimeCaller // Generic read-only contract binding to access the raw methods on
}

// ProtoBufRuntimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtoBufRuntimeTransactorRaw struct {
	Contract *ProtoBufRuntimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtoBufRuntime creates a new instance of ProtoBufRuntime, bound to a specific deployed contract.
func NewProtoBufRuntime(address common.Address, backend bind.ContractBackend) (*ProtoBufRuntime, error) {
	contract, err := bindProtoBufRuntime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtoBufRuntime{ProtoBufRuntimeCaller: ProtoBufRuntimeCaller{contract: contract}, ProtoBufRuntimeTransactor: ProtoBufRuntimeTransactor{contract: contract}, ProtoBufRuntimeFilterer: ProtoBufRuntimeFilterer{contract: contract}}, nil
}

// NewProtoBufRuntimeCaller creates a new read-only instance of ProtoBufRuntime, bound to a specific deployed contract.
func NewProtoBufRuntimeCaller(address common.Address, caller bind.ContractCaller) (*ProtoBufRuntimeCaller, error) {
	contract, err := bindProtoBufRuntime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtoBufRuntimeCaller{contract: contract}, nil
}

// NewProtoBufRuntimeTransactor creates a new write-only instance of ProtoBufRuntime, bound to a specific deployed contract.
func NewProtoBufRuntimeTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtoBufRuntimeTransactor, error) {
	contract, err := bindProtoBufRuntime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtoBufRuntimeTransactor{contract: contract}, nil
}

// NewProtoBufRuntimeFilterer creates a new log filterer instance of ProtoBufRuntime, bound to a specific deployed contract.
func NewProtoBufRuntimeFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtoBufRuntimeFilterer, error) {
	contract, err := bindProtoBufRuntime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtoBufRuntimeFilterer{contract: contract}, nil
}

// bindProtoBufRuntime binds a generic wrapper to an already deployed contract.
func bindProtoBufRuntime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProtoBufRuntimeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtoBufRuntime *ProtoBufRuntimeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtoBufRuntime.Contract.ProtoBufRuntimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtoBufRuntime *ProtoBufRuntimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtoBufRuntime.Contract.ProtoBufRuntimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtoBufRuntime *ProtoBufRuntimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtoBufRuntime.Contract.ProtoBufRuntimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtoBufRuntime *ProtoBufRuntimeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtoBufRuntime.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtoBufRuntime *ProtoBufRuntimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtoBufRuntime.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtoBufRuntime *ProtoBufRuntimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtoBufRuntime.Contract.contract.Transact(opts, method, params...)
}
