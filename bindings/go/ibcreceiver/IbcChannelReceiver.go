// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcreceiver

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

// IbcChannelReceiverMetaData contains all meta data concerning the IbcChannelReceiver contract.
var IbcChannelReceiverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"onChanCloseConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseInit\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartychannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenInit\",\"inputs\":[{\"name\":\"order\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortIdentifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenTry\",\"inputs\":[{\"name\":\"order\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortIdentifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartychannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"}]",
}

// IbcChannelReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcChannelReceiverMetaData.ABI instead.
var IbcChannelReceiverABI = IbcChannelReceiverMetaData.ABI

// IbcChannelReceiver is an auto generated Go binding around an Ethereum contract.
type IbcChannelReceiver struct {
	IbcChannelReceiverCaller     // Read-only binding to the contract
	IbcChannelReceiverTransactor // Write-only binding to the contract
	IbcChannelReceiverFilterer   // Log filterer for contract events
}

// IbcChannelReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcChannelReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcChannelReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcChannelReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcChannelReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcChannelReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcChannelReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcChannelReceiverSession struct {
	Contract     *IbcChannelReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IbcChannelReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcChannelReceiverCallerSession struct {
	Contract *IbcChannelReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IbcChannelReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcChannelReceiverTransactorSession struct {
	Contract     *IbcChannelReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IbcChannelReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcChannelReceiverRaw struct {
	Contract *IbcChannelReceiver // Generic contract binding to access the raw methods on
}

// IbcChannelReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcChannelReceiverCallerRaw struct {
	Contract *IbcChannelReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IbcChannelReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcChannelReceiverTransactorRaw struct {
	Contract *IbcChannelReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcChannelReceiver creates a new instance of IbcChannelReceiver, bound to a specific deployed contract.
func NewIbcChannelReceiver(address common.Address, backend bind.ContractBackend) (*IbcChannelReceiver, error) {
	contract, err := bindIbcChannelReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcChannelReceiver{IbcChannelReceiverCaller: IbcChannelReceiverCaller{contract: contract}, IbcChannelReceiverTransactor: IbcChannelReceiverTransactor{contract: contract}, IbcChannelReceiverFilterer: IbcChannelReceiverFilterer{contract: contract}}, nil
}

// NewIbcChannelReceiverCaller creates a new read-only instance of IbcChannelReceiver, bound to a specific deployed contract.
func NewIbcChannelReceiverCaller(address common.Address, caller bind.ContractCaller) (*IbcChannelReceiverCaller, error) {
	contract, err := bindIbcChannelReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcChannelReceiverCaller{contract: contract}, nil
}

// NewIbcChannelReceiverTransactor creates a new write-only instance of IbcChannelReceiver, bound to a specific deployed contract.
func NewIbcChannelReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcChannelReceiverTransactor, error) {
	contract, err := bindIbcChannelReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcChannelReceiverTransactor{contract: contract}, nil
}

// NewIbcChannelReceiverFilterer creates a new log filterer instance of IbcChannelReceiver, bound to a specific deployed contract.
func NewIbcChannelReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcChannelReceiverFilterer, error) {
	contract, err := bindIbcChannelReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcChannelReceiverFilterer{contract: contract}, nil
}

// bindIbcChannelReceiver binds a generic wrapper to an already deployed contract.
func bindIbcChannelReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcChannelReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcChannelReceiver *IbcChannelReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcChannelReceiver.Contract.IbcChannelReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcChannelReceiver *IbcChannelReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.IbcChannelReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcChannelReceiver *IbcChannelReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.IbcChannelReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcChannelReceiver *IbcChannelReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcChannelReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcChannelReceiver *IbcChannelReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcChannelReceiver *IbcChannelReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcChannelReceiver *IbcChannelReceiverTransactor) OnChanCloseConfirm(opts *bind.TransactOpts, channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcChannelReceiver.contract.Transact(opts, "onChanCloseConfirm", channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcChannelReceiver *IbcChannelReceiverSession) OnChanCloseConfirm(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.OnChanCloseConfirm(&_IbcChannelReceiver.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcChannelReceiver *IbcChannelReceiverTransactorSession) OnChanCloseConfirm(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.OnChanCloseConfirm(&_IbcChannelReceiver.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcChannelReceiver *IbcChannelReceiverTransactor) OnChanCloseInit(opts *bind.TransactOpts, channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcChannelReceiver.contract.Transact(opts, "onChanCloseInit", channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcChannelReceiver *IbcChannelReceiverSession) OnChanCloseInit(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.OnChanCloseInit(&_IbcChannelReceiver.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcChannelReceiver *IbcChannelReceiverTransactorSession) OnChanCloseInit(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.OnChanCloseInit(&_IbcChannelReceiver.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string counterpartyVersion) returns()
func (_IbcChannelReceiver *IbcChannelReceiverTransactor) OnChanOpenAck(opts *bind.TransactOpts, channelId [32]byte, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcChannelReceiver.contract.Transact(opts, "onChanOpenAck", channelId, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string counterpartyVersion) returns()
func (_IbcChannelReceiver *IbcChannelReceiverSession) OnChanOpenAck(channelId [32]byte, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.OnChanOpenAck(&_IbcChannelReceiver.TransactOpts, channelId, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string counterpartyVersion) returns()
func (_IbcChannelReceiver *IbcChannelReceiverTransactorSession) OnChanOpenAck(channelId [32]byte, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.OnChanOpenAck(&_IbcChannelReceiver.TransactOpts, channelId, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IbcChannelReceiver *IbcChannelReceiverTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _IbcChannelReceiver.contract.Transact(opts, "onChanOpenConfirm", channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IbcChannelReceiver *IbcChannelReceiverSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.OnChanOpenConfirm(&_IbcChannelReceiver.TransactOpts, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IbcChannelReceiver *IbcChannelReceiverTransactorSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.OnChanOpenConfirm(&_IbcChannelReceiver.TransactOpts, channelId)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 order, string[] connectionHops, string counterpartyPortIdentifier, string version) returns(string selectedVersion)
func (_IbcChannelReceiver *IbcChannelReceiverTransactor) OnChanOpenInit(opts *bind.TransactOpts, order uint8, connectionHops []string, counterpartyPortIdentifier string, version string) (*types.Transaction, error) {
	return _IbcChannelReceiver.contract.Transact(opts, "onChanOpenInit", order, connectionHops, counterpartyPortIdentifier, version)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 order, string[] connectionHops, string counterpartyPortIdentifier, string version) returns(string selectedVersion)
func (_IbcChannelReceiver *IbcChannelReceiverSession) OnChanOpenInit(order uint8, connectionHops []string, counterpartyPortIdentifier string, version string) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.OnChanOpenInit(&_IbcChannelReceiver.TransactOpts, order, connectionHops, counterpartyPortIdentifier, version)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 order, string[] connectionHops, string counterpartyPortIdentifier, string version) returns(string selectedVersion)
func (_IbcChannelReceiver *IbcChannelReceiverTransactorSession) OnChanOpenInit(order uint8, connectionHops []string, counterpartyPortIdentifier string, version string) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.OnChanOpenInit(&_IbcChannelReceiver.TransactOpts, order, connectionHops, counterpartyPortIdentifier, version)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 order, string[] connectionHops, bytes32 channelId, string counterpartyPortIdentifier, bytes32 counterpartychannelId, string counterpartyVersion) returns(string selectedVersion)
func (_IbcChannelReceiver *IbcChannelReceiverTransactor) OnChanOpenTry(opts *bind.TransactOpts, order uint8, connectionHops []string, channelId [32]byte, counterpartyPortIdentifier string, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcChannelReceiver.contract.Transact(opts, "onChanOpenTry", order, connectionHops, channelId, counterpartyPortIdentifier, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 order, string[] connectionHops, bytes32 channelId, string counterpartyPortIdentifier, bytes32 counterpartychannelId, string counterpartyVersion) returns(string selectedVersion)
func (_IbcChannelReceiver *IbcChannelReceiverSession) OnChanOpenTry(order uint8, connectionHops []string, channelId [32]byte, counterpartyPortIdentifier string, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.OnChanOpenTry(&_IbcChannelReceiver.TransactOpts, order, connectionHops, channelId, counterpartyPortIdentifier, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 order, string[] connectionHops, bytes32 channelId, string counterpartyPortIdentifier, bytes32 counterpartychannelId, string counterpartyVersion) returns(string selectedVersion)
func (_IbcChannelReceiver *IbcChannelReceiverTransactorSession) OnChanOpenTry(order uint8, connectionHops []string, channelId [32]byte, counterpartyPortIdentifier string, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcChannelReceiver.Contract.OnChanOpenTry(&_IbcChannelReceiver.TransactOpts, order, connectionHops, channelId, counterpartyPortIdentifier, counterpartychannelId, counterpartyVersion)
}
