// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibeacon

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

// IBeaconMetaData contains all meta data concerning the IBeacon contract.
var IBeaconMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
}

// IBeaconABI is the input ABI used to generate the binding from.
// Deprecated: Use IBeaconMetaData.ABI instead.
var IBeaconABI = IBeaconMetaData.ABI

// IBeacon is an auto generated Go binding around an Ethereum contract.
type IBeacon struct {
	IBeaconCaller     // Read-only binding to the contract
	IBeaconTransactor // Write-only binding to the contract
	IBeaconFilterer   // Log filterer for contract events
}

// IBeaconCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBeaconCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBeaconTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBeaconFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBeaconSession struct {
	Contract     *IBeacon          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBeaconCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBeaconCallerSession struct {
	Contract *IBeaconCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IBeaconTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBeaconTransactorSession struct {
	Contract     *IBeaconTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IBeaconRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBeaconRaw struct {
	Contract *IBeacon // Generic contract binding to access the raw methods on
}

// IBeaconCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBeaconCallerRaw struct {
	Contract *IBeaconCaller // Generic read-only contract binding to access the raw methods on
}

// IBeaconTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBeaconTransactorRaw struct {
	Contract *IBeaconTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBeacon creates a new instance of IBeacon, bound to a specific deployed contract.
func NewIBeacon(address common.Address, backend bind.ContractBackend) (*IBeacon, error) {
	contract, err := bindIBeacon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBeacon{IBeaconCaller: IBeaconCaller{contract: contract}, IBeaconTransactor: IBeaconTransactor{contract: contract}, IBeaconFilterer: IBeaconFilterer{contract: contract}}, nil
}

// NewIBeaconCaller creates a new read-only instance of IBeacon, bound to a specific deployed contract.
func NewIBeaconCaller(address common.Address, caller bind.ContractCaller) (*IBeaconCaller, error) {
	contract, err := bindIBeacon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconCaller{contract: contract}, nil
}

// NewIBeaconTransactor creates a new write-only instance of IBeacon, bound to a specific deployed contract.
func NewIBeaconTransactor(address common.Address, transactor bind.ContractTransactor) (*IBeaconTransactor, error) {
	contract, err := bindIBeacon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconTransactor{contract: contract}, nil
}

// NewIBeaconFilterer creates a new log filterer instance of IBeacon, bound to a specific deployed contract.
func NewIBeaconFilterer(address common.Address, filterer bind.ContractFilterer) (*IBeaconFilterer, error) {
	contract, err := bindIBeacon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBeaconFilterer{contract: contract}, nil
}

// bindIBeacon binds a generic wrapper to an already deployed contract.
func bindIBeacon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBeaconMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeacon *IBeaconRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBeacon.Contract.IBeaconCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeacon *IBeaconRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeacon.Contract.IBeaconTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeacon *IBeaconRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeacon.Contract.IBeaconTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeacon *IBeaconCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBeacon.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeacon *IBeaconTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeacon.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeacon *IBeaconTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeacon.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeacon *IBeaconCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBeacon.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeacon *IBeaconSession) Implementation() (common.Address, error) {
	return _IBeacon.Contract.Implementation(&_IBeacon.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeacon *IBeaconCallerSession) Implementation() (common.Address, error) {
	return _IBeacon.Contract.Implementation(&_IBeacon.CallOpts)
}
