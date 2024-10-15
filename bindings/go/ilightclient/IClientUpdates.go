// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ilightclient

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

// IClientUpdatesMetaData contains all meta data concerning the IClientUpdates contract.
var IClientUpdatesMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"LIGHT_CLIENT_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumLightClientType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IClientUpdatesABI is the input ABI used to generate the binding from.
// Deprecated: Use IClientUpdatesMetaData.ABI instead.
var IClientUpdatesABI = IClientUpdatesMetaData.ABI

// IClientUpdates is an auto generated Go binding around an Ethereum contract.
type IClientUpdates struct {
	IClientUpdatesCaller     // Read-only binding to the contract
	IClientUpdatesTransactor // Write-only binding to the contract
	IClientUpdatesFilterer   // Log filterer for contract events
}

// IClientUpdatesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IClientUpdatesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IClientUpdatesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IClientUpdatesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IClientUpdatesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IClientUpdatesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IClientUpdatesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IClientUpdatesSession struct {
	Contract     *IClientUpdates   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IClientUpdatesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IClientUpdatesCallerSession struct {
	Contract *IClientUpdatesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IClientUpdatesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IClientUpdatesTransactorSession struct {
	Contract     *IClientUpdatesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IClientUpdatesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IClientUpdatesRaw struct {
	Contract *IClientUpdates // Generic contract binding to access the raw methods on
}

// IClientUpdatesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IClientUpdatesCallerRaw struct {
	Contract *IClientUpdatesCaller // Generic read-only contract binding to access the raw methods on
}

// IClientUpdatesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IClientUpdatesTransactorRaw struct {
	Contract *IClientUpdatesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIClientUpdates creates a new instance of IClientUpdates, bound to a specific deployed contract.
func NewIClientUpdates(address common.Address, backend bind.ContractBackend) (*IClientUpdates, error) {
	contract, err := bindIClientUpdates(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IClientUpdates{IClientUpdatesCaller: IClientUpdatesCaller{contract: contract}, IClientUpdatesTransactor: IClientUpdatesTransactor{contract: contract}, IClientUpdatesFilterer: IClientUpdatesFilterer{contract: contract}}, nil
}

// NewIClientUpdatesCaller creates a new read-only instance of IClientUpdates, bound to a specific deployed contract.
func NewIClientUpdatesCaller(address common.Address, caller bind.ContractCaller) (*IClientUpdatesCaller, error) {
	contract, err := bindIClientUpdates(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IClientUpdatesCaller{contract: contract}, nil
}

// NewIClientUpdatesTransactor creates a new write-only instance of IClientUpdates, bound to a specific deployed contract.
func NewIClientUpdatesTransactor(address common.Address, transactor bind.ContractTransactor) (*IClientUpdatesTransactor, error) {
	contract, err := bindIClientUpdates(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IClientUpdatesTransactor{contract: contract}, nil
}

// NewIClientUpdatesFilterer creates a new log filterer instance of IClientUpdates, bound to a specific deployed contract.
func NewIClientUpdatesFilterer(address common.Address, filterer bind.ContractFilterer) (*IClientUpdatesFilterer, error) {
	contract, err := bindIClientUpdates(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IClientUpdatesFilterer{contract: contract}, nil
}

// bindIClientUpdates binds a generic wrapper to an already deployed contract.
func bindIClientUpdates(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IClientUpdatesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IClientUpdates *IClientUpdatesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IClientUpdates.Contract.IClientUpdatesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IClientUpdates *IClientUpdatesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IClientUpdates.Contract.IClientUpdatesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IClientUpdates *IClientUpdatesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IClientUpdates.Contract.IClientUpdatesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IClientUpdates *IClientUpdatesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IClientUpdates.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IClientUpdates *IClientUpdatesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IClientUpdates.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IClientUpdates *IClientUpdatesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IClientUpdates.Contract.contract.Transact(opts, method, params...)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_IClientUpdates *IClientUpdatesCaller) LIGHTCLIENTTYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IClientUpdates.contract.Call(opts, &out, "LIGHT_CLIENT_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_IClientUpdates *IClientUpdatesSession) LIGHTCLIENTTYPE() (uint8, error) {
	return _IClientUpdates.Contract.LIGHTCLIENTTYPE(&_IClientUpdates.CallOpts)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_IClientUpdates *IClientUpdatesCallerSession) LIGHTCLIENTTYPE() (uint8, error) {
	return _IClientUpdates.Contract.LIGHTCLIENTTYPE(&_IClientUpdates.CallOpts)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes proof, uint256 height, uint256 appHash) returns()
func (_IClientUpdates *IClientUpdatesTransactor) UpdateClient(opts *bind.TransactOpts, proof []byte, height *big.Int, appHash *big.Int) (*types.Transaction, error) {
	return _IClientUpdates.contract.Transact(opts, "updateClient", proof, height, appHash)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes proof, uint256 height, uint256 appHash) returns()
func (_IClientUpdates *IClientUpdatesSession) UpdateClient(proof []byte, height *big.Int, appHash *big.Int) (*types.Transaction, error) {
	return _IClientUpdates.Contract.UpdateClient(&_IClientUpdates.TransactOpts, proof, height, appHash)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes proof, uint256 height, uint256 appHash) returns()
func (_IClientUpdates *IClientUpdatesTransactorSession) UpdateClient(proof []byte, height *big.Int, appHash *big.Int) (*types.Transaction, error) {
	return _IClientUpdates.Contract.UpdateClient(&_IClientUpdates.TransactOpts, proof, height, appHash)
}
