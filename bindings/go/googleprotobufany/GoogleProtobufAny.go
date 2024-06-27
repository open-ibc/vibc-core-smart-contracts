// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package googleprotobufany

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

// GoogleProtobufAnyMetaData contains all meta data concerning the GoogleProtobufAny contract.
var GoogleProtobufAnyMetaData = &bind.MetaData{
	ABI: "[]",
}

// GoogleProtobufAnyABI is the input ABI used to generate the binding from.
// Deprecated: Use GoogleProtobufAnyMetaData.ABI instead.
var GoogleProtobufAnyABI = GoogleProtobufAnyMetaData.ABI

// GoogleProtobufAny is an auto generated Go binding around an Ethereum contract.
type GoogleProtobufAny struct {
	GoogleProtobufAnyCaller     // Read-only binding to the contract
	GoogleProtobufAnyTransactor // Write-only binding to the contract
	GoogleProtobufAnyFilterer   // Log filterer for contract events
}

// GoogleProtobufAnyCaller is an auto generated read-only Go binding around an Ethereum contract.
type GoogleProtobufAnyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoogleProtobufAnyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GoogleProtobufAnyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoogleProtobufAnyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GoogleProtobufAnyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GoogleProtobufAnySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GoogleProtobufAnySession struct {
	Contract     *GoogleProtobufAny // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GoogleProtobufAnyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GoogleProtobufAnyCallerSession struct {
	Contract *GoogleProtobufAnyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GoogleProtobufAnyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GoogleProtobufAnyTransactorSession struct {
	Contract     *GoogleProtobufAnyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GoogleProtobufAnyRaw is an auto generated low-level Go binding around an Ethereum contract.
type GoogleProtobufAnyRaw struct {
	Contract *GoogleProtobufAny // Generic contract binding to access the raw methods on
}

// GoogleProtobufAnyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GoogleProtobufAnyCallerRaw struct {
	Contract *GoogleProtobufAnyCaller // Generic read-only contract binding to access the raw methods on
}

// GoogleProtobufAnyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GoogleProtobufAnyTransactorRaw struct {
	Contract *GoogleProtobufAnyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGoogleProtobufAny creates a new instance of GoogleProtobufAny, bound to a specific deployed contract.
func NewGoogleProtobufAny(address common.Address, backend bind.ContractBackend) (*GoogleProtobufAny, error) {
	contract, err := bindGoogleProtobufAny(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GoogleProtobufAny{GoogleProtobufAnyCaller: GoogleProtobufAnyCaller{contract: contract}, GoogleProtobufAnyTransactor: GoogleProtobufAnyTransactor{contract: contract}, GoogleProtobufAnyFilterer: GoogleProtobufAnyFilterer{contract: contract}}, nil
}

// NewGoogleProtobufAnyCaller creates a new read-only instance of GoogleProtobufAny, bound to a specific deployed contract.
func NewGoogleProtobufAnyCaller(address common.Address, caller bind.ContractCaller) (*GoogleProtobufAnyCaller, error) {
	contract, err := bindGoogleProtobufAny(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GoogleProtobufAnyCaller{contract: contract}, nil
}

// NewGoogleProtobufAnyTransactor creates a new write-only instance of GoogleProtobufAny, bound to a specific deployed contract.
func NewGoogleProtobufAnyTransactor(address common.Address, transactor bind.ContractTransactor) (*GoogleProtobufAnyTransactor, error) {
	contract, err := bindGoogleProtobufAny(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GoogleProtobufAnyTransactor{contract: contract}, nil
}

// NewGoogleProtobufAnyFilterer creates a new log filterer instance of GoogleProtobufAny, bound to a specific deployed contract.
func NewGoogleProtobufAnyFilterer(address common.Address, filterer bind.ContractFilterer) (*GoogleProtobufAnyFilterer, error) {
	contract, err := bindGoogleProtobufAny(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GoogleProtobufAnyFilterer{contract: contract}, nil
}

// bindGoogleProtobufAny binds a generic wrapper to an already deployed contract.
func bindGoogleProtobufAny(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GoogleProtobufAnyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoogleProtobufAny *GoogleProtobufAnyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoogleProtobufAny.Contract.GoogleProtobufAnyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoogleProtobufAny *GoogleProtobufAnyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoogleProtobufAny.Contract.GoogleProtobufAnyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoogleProtobufAny *GoogleProtobufAnyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoogleProtobufAny.Contract.GoogleProtobufAnyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GoogleProtobufAny *GoogleProtobufAnyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GoogleProtobufAny.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GoogleProtobufAny *GoogleProtobufAnyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GoogleProtobufAny.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GoogleProtobufAny *GoogleProtobufAnyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GoogleProtobufAny.Contract.contract.Transact(opts, method, params...)
}
