// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package icrossl2prover

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

// ICrossL2ProverMetaData contains all meta data concerning the ICrossL2Prover contract.
var ICrossL2ProverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"LIGHT_CLIENT_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumLightClientType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getState\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateEvent\",\"inputs\":[{\"name\":\"receiptIndex\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiptRLPEncodedBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"logBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateReceipt\",\"inputs\":[{\"name\":\"receiptIndex\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiptRLPEncodedBytes\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"}]",
}

// ICrossL2ProverABI is the input ABI used to generate the binding from.
// Deprecated: Use ICrossL2ProverMetaData.ABI instead.
var ICrossL2ProverABI = ICrossL2ProverMetaData.ABI

// ICrossL2Prover is an auto generated Go binding around an Ethereum contract.
type ICrossL2Prover struct {
	ICrossL2ProverCaller     // Read-only binding to the contract
	ICrossL2ProverTransactor // Write-only binding to the contract
	ICrossL2ProverFilterer   // Log filterer for contract events
}

// ICrossL2ProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ICrossL2ProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossL2ProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ICrossL2ProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossL2ProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICrossL2ProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICrossL2ProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICrossL2ProverSession struct {
	Contract     *ICrossL2Prover   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICrossL2ProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICrossL2ProverCallerSession struct {
	Contract *ICrossL2ProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ICrossL2ProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICrossL2ProverTransactorSession struct {
	Contract     *ICrossL2ProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ICrossL2ProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ICrossL2ProverRaw struct {
	Contract *ICrossL2Prover // Generic contract binding to access the raw methods on
}

// ICrossL2ProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICrossL2ProverCallerRaw struct {
	Contract *ICrossL2ProverCaller // Generic read-only contract binding to access the raw methods on
}

// ICrossL2ProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICrossL2ProverTransactorRaw struct {
	Contract *ICrossL2ProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewICrossL2Prover creates a new instance of ICrossL2Prover, bound to a specific deployed contract.
func NewICrossL2Prover(address common.Address, backend bind.ContractBackend) (*ICrossL2Prover, error) {
	contract, err := bindICrossL2Prover(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICrossL2Prover{ICrossL2ProverCaller: ICrossL2ProverCaller{contract: contract}, ICrossL2ProverTransactor: ICrossL2ProverTransactor{contract: contract}, ICrossL2ProverFilterer: ICrossL2ProverFilterer{contract: contract}}, nil
}

// NewICrossL2ProverCaller creates a new read-only instance of ICrossL2Prover, bound to a specific deployed contract.
func NewICrossL2ProverCaller(address common.Address, caller bind.ContractCaller) (*ICrossL2ProverCaller, error) {
	contract, err := bindICrossL2Prover(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICrossL2ProverCaller{contract: contract}, nil
}

// NewICrossL2ProverTransactor creates a new write-only instance of ICrossL2Prover, bound to a specific deployed contract.
func NewICrossL2ProverTransactor(address common.Address, transactor bind.ContractTransactor) (*ICrossL2ProverTransactor, error) {
	contract, err := bindICrossL2Prover(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICrossL2ProverTransactor{contract: contract}, nil
}

// NewICrossL2ProverFilterer creates a new log filterer instance of ICrossL2Prover, bound to a specific deployed contract.
func NewICrossL2ProverFilterer(address common.Address, filterer bind.ContractFilterer) (*ICrossL2ProverFilterer, error) {
	contract, err := bindICrossL2Prover(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICrossL2ProverFilterer{contract: contract}, nil
}

// bindICrossL2Prover binds a generic wrapper to an already deployed contract.
func bindICrossL2Prover(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICrossL2ProverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrossL2Prover *ICrossL2ProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrossL2Prover.Contract.ICrossL2ProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrossL2Prover *ICrossL2ProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossL2Prover.Contract.ICrossL2ProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrossL2Prover *ICrossL2ProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrossL2Prover.Contract.ICrossL2ProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICrossL2Prover *ICrossL2ProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICrossL2Prover.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICrossL2Prover *ICrossL2ProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICrossL2Prover.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICrossL2Prover *ICrossL2ProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICrossL2Prover.Contract.contract.Transact(opts, method, params...)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_ICrossL2Prover *ICrossL2ProverCaller) LIGHTCLIENTTYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ICrossL2Prover.contract.Call(opts, &out, "LIGHT_CLIENT_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_ICrossL2Prover *ICrossL2ProverSession) LIGHTCLIENTTYPE() (uint8, error) {
	return _ICrossL2Prover.Contract.LIGHTCLIENTTYPE(&_ICrossL2Prover.CallOpts)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_ICrossL2Prover *ICrossL2ProverCallerSession) LIGHTCLIENTTYPE() (uint8, error) {
	return _ICrossL2Prover.Contract.LIGHTCLIENTTYPE(&_ICrossL2Prover.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256)
func (_ICrossL2Prover *ICrossL2ProverCaller) GetState(opts *bind.CallOpts, height *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ICrossL2Prover.contract.Call(opts, &out, "getState", height)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256)
func (_ICrossL2Prover *ICrossL2ProverSession) GetState(height *big.Int) (*big.Int, error) {
	return _ICrossL2Prover.Contract.GetState(&_ICrossL2Prover.CallOpts, height)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256)
func (_ICrossL2Prover *ICrossL2ProverCallerSession) GetState(height *big.Int) (*big.Int, error) {
	return _ICrossL2Prover.Contract.GetState(&_ICrossL2Prover.CallOpts, height)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes proof, uint256 height, uint256 appHash) returns()
func (_ICrossL2Prover *ICrossL2ProverTransactor) UpdateClient(opts *bind.TransactOpts, proof []byte, height *big.Int, appHash *big.Int) (*types.Transaction, error) {
	return _ICrossL2Prover.contract.Transact(opts, "updateClient", proof, height, appHash)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes proof, uint256 height, uint256 appHash) returns()
func (_ICrossL2Prover *ICrossL2ProverSession) UpdateClient(proof []byte, height *big.Int, appHash *big.Int) (*types.Transaction, error) {
	return _ICrossL2Prover.Contract.UpdateClient(&_ICrossL2Prover.TransactOpts, proof, height, appHash)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes proof, uint256 height, uint256 appHash) returns()
func (_ICrossL2Prover *ICrossL2ProverTransactorSession) UpdateClient(proof []byte, height *big.Int, appHash *big.Int) (*types.Transaction, error) {
	return _ICrossL2Prover.Contract.UpdateClient(&_ICrossL2Prover.TransactOpts, proof, height, appHash)
}

// ValidateEvent is a paid mutator transaction binding the contract method 0x81ce4d34.
//
// Solidity: function validateEvent(bytes receiptIndex, bytes receiptRLPEncodedBytes, uint256 logIndex, bytes logBytes, bytes proof) returns(bool)
func (_ICrossL2Prover *ICrossL2ProverTransactor) ValidateEvent(opts *bind.TransactOpts, receiptIndex []byte, receiptRLPEncodedBytes []byte, logIndex *big.Int, logBytes []byte, proof []byte) (*types.Transaction, error) {
	return _ICrossL2Prover.contract.Transact(opts, "validateEvent", receiptIndex, receiptRLPEncodedBytes, logIndex, logBytes, proof)
}

// ValidateEvent is a paid mutator transaction binding the contract method 0x81ce4d34.
//
// Solidity: function validateEvent(bytes receiptIndex, bytes receiptRLPEncodedBytes, uint256 logIndex, bytes logBytes, bytes proof) returns(bool)
func (_ICrossL2Prover *ICrossL2ProverSession) ValidateEvent(receiptIndex []byte, receiptRLPEncodedBytes []byte, logIndex *big.Int, logBytes []byte, proof []byte) (*types.Transaction, error) {
	return _ICrossL2Prover.Contract.ValidateEvent(&_ICrossL2Prover.TransactOpts, receiptIndex, receiptRLPEncodedBytes, logIndex, logBytes, proof)
}

// ValidateEvent is a paid mutator transaction binding the contract method 0x81ce4d34.
//
// Solidity: function validateEvent(bytes receiptIndex, bytes receiptRLPEncodedBytes, uint256 logIndex, bytes logBytes, bytes proof) returns(bool)
func (_ICrossL2Prover *ICrossL2ProverTransactorSession) ValidateEvent(receiptIndex []byte, receiptRLPEncodedBytes []byte, logIndex *big.Int, logBytes []byte, proof []byte) (*types.Transaction, error) {
	return _ICrossL2Prover.Contract.ValidateEvent(&_ICrossL2Prover.TransactOpts, receiptIndex, receiptRLPEncodedBytes, logIndex, logBytes, proof)
}

// ValidateReceipt is a paid mutator transaction binding the contract method 0x85532872.
//
// Solidity: function validateReceipt(bytes receiptIndex, bytes receiptRLPEncodedBytes, bytes proof) returns(bool)
func (_ICrossL2Prover *ICrossL2ProverTransactor) ValidateReceipt(opts *bind.TransactOpts, receiptIndex []byte, receiptRLPEncodedBytes []byte, proof []byte) (*types.Transaction, error) {
	return _ICrossL2Prover.contract.Transact(opts, "validateReceipt", receiptIndex, receiptRLPEncodedBytes, proof)
}

// ValidateReceipt is a paid mutator transaction binding the contract method 0x85532872.
//
// Solidity: function validateReceipt(bytes receiptIndex, bytes receiptRLPEncodedBytes, bytes proof) returns(bool)
func (_ICrossL2Prover *ICrossL2ProverSession) ValidateReceipt(receiptIndex []byte, receiptRLPEncodedBytes []byte, proof []byte) (*types.Transaction, error) {
	return _ICrossL2Prover.Contract.ValidateReceipt(&_ICrossL2Prover.TransactOpts, receiptIndex, receiptRLPEncodedBytes, proof)
}

// ValidateReceipt is a paid mutator transaction binding the contract method 0x85532872.
//
// Solidity: function validateReceipt(bytes receiptIndex, bytes receiptRLPEncodedBytes, bytes proof) returns(bool)
func (_ICrossL2Prover *ICrossL2ProverTransactorSession) ValidateReceipt(receiptIndex []byte, receiptRLPEncodedBytes []byte, proof []byte) (*types.Transaction, error) {
	return _ICrossL2Prover.Contract.ValidateReceipt(&_ICrossL2Prover.TransactOpts, receiptIndex, receiptRLPEncodedBytes, proof)
}
