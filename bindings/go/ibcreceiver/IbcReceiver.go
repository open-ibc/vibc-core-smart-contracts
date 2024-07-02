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

// AckPacket is an auto generated low-level Go binding around an user-defined struct.
type AckPacket struct {
	Success bool
	Data    []byte
}

// Height is an auto generated low-level Go binding around an user-defined struct.
type Height struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IbcEndpoint is an auto generated low-level Go binding around an user-defined struct.
type IbcEndpoint struct {
	PortId    string
	ChannelId [32]byte
}

// IbcPacket is an auto generated low-level Go binding around an user-defined struct.
type IbcPacket struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}

// IbcReceiverMetaData contains all meta data concerning the IbcReceiver contract.
var IbcReceiverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseInit\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartychannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenInit\",\"inputs\":[{\"name\":\"order\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortIdentifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenTry\",\"inputs\":[{\"name\":\"order\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortIdentifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartychannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IbcReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcReceiverMetaData.ABI instead.
var IbcReceiverABI = IbcReceiverMetaData.ABI

// IbcReceiver is an auto generated Go binding around an Ethereum contract.
type IbcReceiver struct {
	IbcReceiverCaller     // Read-only binding to the contract
	IbcReceiverTransactor // Write-only binding to the contract
	IbcReceiverFilterer   // Log filterer for contract events
}

// IbcReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcReceiverSession struct {
	Contract     *IbcReceiver      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcReceiverCallerSession struct {
	Contract *IbcReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IbcReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcReceiverTransactorSession struct {
	Contract     *IbcReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IbcReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcReceiverRaw struct {
	Contract *IbcReceiver // Generic contract binding to access the raw methods on
}

// IbcReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcReceiverCallerRaw struct {
	Contract *IbcReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IbcReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcReceiverTransactorRaw struct {
	Contract *IbcReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcReceiver creates a new instance of IbcReceiver, bound to a specific deployed contract.
func NewIbcReceiver(address common.Address, backend bind.ContractBackend) (*IbcReceiver, error) {
	contract, err := bindIbcReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcReceiver{IbcReceiverCaller: IbcReceiverCaller{contract: contract}, IbcReceiverTransactor: IbcReceiverTransactor{contract: contract}, IbcReceiverFilterer: IbcReceiverFilterer{contract: contract}}, nil
}

// NewIbcReceiverCaller creates a new read-only instance of IbcReceiver, bound to a specific deployed contract.
func NewIbcReceiverCaller(address common.Address, caller bind.ContractCaller) (*IbcReceiverCaller, error) {
	contract, err := bindIbcReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverCaller{contract: contract}, nil
}

// NewIbcReceiverTransactor creates a new write-only instance of IbcReceiver, bound to a specific deployed contract.
func NewIbcReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcReceiverTransactor, error) {
	contract, err := bindIbcReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverTransactor{contract: contract}, nil
}

// NewIbcReceiverFilterer creates a new log filterer instance of IbcReceiver, bound to a specific deployed contract.
func NewIbcReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcReceiverFilterer, error) {
	contract, err := bindIbcReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcReceiverFilterer{contract: contract}, nil
}

// bindIbcReceiver binds a generic wrapper to an already deployed contract.
func bindIbcReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcReceiver *IbcReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcReceiver.Contract.IbcReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcReceiver *IbcReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcReceiver.Contract.IbcReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcReceiver *IbcReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcReceiver.Contract.IbcReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcReceiver *IbcReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcReceiver *IbcReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcReceiver *IbcReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IbcReceiver *IbcReceiverTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcReceiver.contract.Transact(opts, "onAcknowledgementPacket", packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IbcReceiver *IbcReceiverSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnAcknowledgementPacket(&_IbcReceiver.TransactOpts, packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IbcReceiver *IbcReceiverTransactorSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnAcknowledgementPacket(&_IbcReceiver.TransactOpts, packet, ack)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcReceiver *IbcReceiverTransactor) OnChanCloseConfirm(opts *bind.TransactOpts, channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcReceiver.contract.Transact(opts, "onChanCloseConfirm", channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcReceiver *IbcReceiverSession) OnChanCloseConfirm(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnChanCloseConfirm(&_IbcReceiver.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcReceiver *IbcReceiverTransactorSession) OnChanCloseConfirm(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnChanCloseConfirm(&_IbcReceiver.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcReceiver *IbcReceiverTransactor) OnChanCloseInit(opts *bind.TransactOpts, channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcReceiver.contract.Transact(opts, "onChanCloseInit", channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcReceiver *IbcReceiverSession) OnChanCloseInit(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnChanCloseInit(&_IbcReceiver.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcReceiver *IbcReceiverTransactorSession) OnChanCloseInit(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnChanCloseInit(&_IbcReceiver.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string counterpartyVersion) returns()
func (_IbcReceiver *IbcReceiverTransactor) OnChanOpenAck(opts *bind.TransactOpts, channelId [32]byte, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcReceiver.contract.Transact(opts, "onChanOpenAck", channelId, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string counterpartyVersion) returns()
func (_IbcReceiver *IbcReceiverSession) OnChanOpenAck(channelId [32]byte, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnChanOpenAck(&_IbcReceiver.TransactOpts, channelId, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string counterpartyVersion) returns()
func (_IbcReceiver *IbcReceiverTransactorSession) OnChanOpenAck(channelId [32]byte, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnChanOpenAck(&_IbcReceiver.TransactOpts, channelId, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IbcReceiver *IbcReceiverTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _IbcReceiver.contract.Transact(opts, "onChanOpenConfirm", channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IbcReceiver *IbcReceiverSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnChanOpenConfirm(&_IbcReceiver.TransactOpts, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IbcReceiver *IbcReceiverTransactorSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnChanOpenConfirm(&_IbcReceiver.TransactOpts, channelId)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 order, string[] connectionHops, string counterpartyPortIdentifier, string version) returns(string selectedVersion)
func (_IbcReceiver *IbcReceiverTransactor) OnChanOpenInit(opts *bind.TransactOpts, order uint8, connectionHops []string, counterpartyPortIdentifier string, version string) (*types.Transaction, error) {
	return _IbcReceiver.contract.Transact(opts, "onChanOpenInit", order, connectionHops, counterpartyPortIdentifier, version)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 order, string[] connectionHops, string counterpartyPortIdentifier, string version) returns(string selectedVersion)
func (_IbcReceiver *IbcReceiverSession) OnChanOpenInit(order uint8, connectionHops []string, counterpartyPortIdentifier string, version string) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnChanOpenInit(&_IbcReceiver.TransactOpts, order, connectionHops, counterpartyPortIdentifier, version)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 order, string[] connectionHops, string counterpartyPortIdentifier, string version) returns(string selectedVersion)
func (_IbcReceiver *IbcReceiverTransactorSession) OnChanOpenInit(order uint8, connectionHops []string, counterpartyPortIdentifier string, version string) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnChanOpenInit(&_IbcReceiver.TransactOpts, order, connectionHops, counterpartyPortIdentifier, version)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 order, string[] connectionHops, bytes32 channelId, string counterpartyPortIdentifier, bytes32 counterpartychannelId, string counterpartyVersion) returns(string selectedVersion)
func (_IbcReceiver *IbcReceiverTransactor) OnChanOpenTry(opts *bind.TransactOpts, order uint8, connectionHops []string, channelId [32]byte, counterpartyPortIdentifier string, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcReceiver.contract.Transact(opts, "onChanOpenTry", order, connectionHops, channelId, counterpartyPortIdentifier, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 order, string[] connectionHops, bytes32 channelId, string counterpartyPortIdentifier, bytes32 counterpartychannelId, string counterpartyVersion) returns(string selectedVersion)
func (_IbcReceiver *IbcReceiverSession) OnChanOpenTry(order uint8, connectionHops []string, channelId [32]byte, counterpartyPortIdentifier string, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnChanOpenTry(&_IbcReceiver.TransactOpts, order, connectionHops, channelId, counterpartyPortIdentifier, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 order, string[] connectionHops, bytes32 channelId, string counterpartyPortIdentifier, bytes32 counterpartychannelId, string counterpartyVersion) returns(string selectedVersion)
func (_IbcReceiver *IbcReceiverTransactorSession) OnChanOpenTry(order uint8, connectionHops []string, channelId [32]byte, counterpartyPortIdentifier string, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnChanOpenTry(&_IbcReceiver.TransactOpts, order, connectionHops, channelId, counterpartyPortIdentifier, counterpartychannelId, counterpartyVersion)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_IbcReceiver *IbcReceiverTransactor) OnRecvPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _IbcReceiver.contract.Transact(opts, "onRecvPacket", packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_IbcReceiver *IbcReceiverSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnRecvPacket(&_IbcReceiver.TransactOpts, packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_IbcReceiver *IbcReceiverTransactorSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnRecvPacket(&_IbcReceiver.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IbcReceiver *IbcReceiverTransactor) OnTimeoutPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _IbcReceiver.contract.Transact(opts, "onTimeoutPacket", packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IbcReceiver *IbcReceiverSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnTimeoutPacket(&_IbcReceiver.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IbcReceiver *IbcReceiverTransactorSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IbcReceiver.Contract.OnTimeoutPacket(&_IbcReceiver.TransactOpts, packet)
}
