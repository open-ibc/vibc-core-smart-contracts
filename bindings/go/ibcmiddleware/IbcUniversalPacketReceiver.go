// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcmiddleware

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

// UniversalPacket is an auto generated low-level Go binding around an user-defined struct.
type UniversalPacket struct {
	SrcPortAddr  [32]byte
	MwBitmap     *big.Int
	DestPortAddr [32]byte
	AppData      []byte
}

// IbcUniversalPacketReceiverMetaData contains all meta data concerning the IbcUniversalPacketReceiver contract.
var IbcUniversalPacketReceiverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"onRecvUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onUniversalAcknowledgement\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"invalidChannelId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"unauthorizedChannel\",\"inputs\":[]}]",
}

// IbcUniversalPacketReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcUniversalPacketReceiverMetaData.ABI instead.
var IbcUniversalPacketReceiverABI = IbcUniversalPacketReceiverMetaData.ABI

// IbcUniversalPacketReceiver is an auto generated Go binding around an Ethereum contract.
type IbcUniversalPacketReceiver struct {
	IbcUniversalPacketReceiverCaller     // Read-only binding to the contract
	IbcUniversalPacketReceiverTransactor // Write-only binding to the contract
	IbcUniversalPacketReceiverFilterer   // Log filterer for contract events
}

// IbcUniversalPacketReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcUniversalPacketReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUniversalPacketReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcUniversalPacketReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUniversalPacketReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcUniversalPacketReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUniversalPacketReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcUniversalPacketReceiverSession struct {
	Contract     *IbcUniversalPacketReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IbcUniversalPacketReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcUniversalPacketReceiverCallerSession struct {
	Contract *IbcUniversalPacketReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// IbcUniversalPacketReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcUniversalPacketReceiverTransactorSession struct {
	Contract     *IbcUniversalPacketReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// IbcUniversalPacketReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcUniversalPacketReceiverRaw struct {
	Contract *IbcUniversalPacketReceiver // Generic contract binding to access the raw methods on
}

// IbcUniversalPacketReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcUniversalPacketReceiverCallerRaw struct {
	Contract *IbcUniversalPacketReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IbcUniversalPacketReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcUniversalPacketReceiverTransactorRaw struct {
	Contract *IbcUniversalPacketReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcUniversalPacketReceiver creates a new instance of IbcUniversalPacketReceiver, bound to a specific deployed contract.
func NewIbcUniversalPacketReceiver(address common.Address, backend bind.ContractBackend) (*IbcUniversalPacketReceiver, error) {
	contract, err := bindIbcUniversalPacketReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketReceiver{IbcUniversalPacketReceiverCaller: IbcUniversalPacketReceiverCaller{contract: contract}, IbcUniversalPacketReceiverTransactor: IbcUniversalPacketReceiverTransactor{contract: contract}, IbcUniversalPacketReceiverFilterer: IbcUniversalPacketReceiverFilterer{contract: contract}}, nil
}

// NewIbcUniversalPacketReceiverCaller creates a new read-only instance of IbcUniversalPacketReceiver, bound to a specific deployed contract.
func NewIbcUniversalPacketReceiverCaller(address common.Address, caller bind.ContractCaller) (*IbcUniversalPacketReceiverCaller, error) {
	contract, err := bindIbcUniversalPacketReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketReceiverCaller{contract: contract}, nil
}

// NewIbcUniversalPacketReceiverTransactor creates a new write-only instance of IbcUniversalPacketReceiver, bound to a specific deployed contract.
func NewIbcUniversalPacketReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcUniversalPacketReceiverTransactor, error) {
	contract, err := bindIbcUniversalPacketReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketReceiverTransactor{contract: contract}, nil
}

// NewIbcUniversalPacketReceiverFilterer creates a new log filterer instance of IbcUniversalPacketReceiver, bound to a specific deployed contract.
func NewIbcUniversalPacketReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcUniversalPacketReceiverFilterer, error) {
	contract, err := bindIbcUniversalPacketReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketReceiverFilterer{contract: contract}, nil
}

// bindIbcUniversalPacketReceiver binds a generic wrapper to an already deployed contract.
func bindIbcUniversalPacketReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcUniversalPacketReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcUniversalPacketReceiver.Contract.IbcUniversalPacketReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.Contract.IbcUniversalPacketReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.Contract.IbcUniversalPacketReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcUniversalPacketReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket) returns((bool,bytes) ackPacket)
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverTransactor) OnRecvUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.contract.Transact(opts, "onRecvUniversalPacket", channelId, ucPacket)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket) returns((bool,bytes) ackPacket)
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverSession) OnRecvUniversalPacket(channelId [32]byte, ucPacket UniversalPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.Contract.OnRecvUniversalPacket(&_IbcUniversalPacketReceiver.TransactOpts, channelId, ucPacket)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket) returns((bool,bytes) ackPacket)
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverTransactorSession) OnRecvUniversalPacket(channelId [32]byte, ucPacket UniversalPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.Contract.OnRecvUniversalPacket(&_IbcUniversalPacketReceiver.TransactOpts, channelId, ucPacket)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverTransactor) OnTimeoutUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.contract.Transact(opts, "onTimeoutUniversalPacket", channelId, packet)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverSession) OnTimeoutUniversalPacket(channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.Contract.OnTimeoutUniversalPacket(&_IbcUniversalPacketReceiver.TransactOpts, channelId, packet)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverTransactorSession) OnTimeoutUniversalPacket(channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.Contract.OnTimeoutUniversalPacket(&_IbcUniversalPacketReceiver.TransactOpts, channelId, packet)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverTransactor) OnUniversalAcknowledgement(opts *bind.TransactOpts, channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.contract.Transact(opts, "onUniversalAcknowledgement", channelId, packet, ack)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverSession) OnUniversalAcknowledgement(channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.Contract.OnUniversalAcknowledgement(&_IbcUniversalPacketReceiver.TransactOpts, channelId, packet, ack)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_IbcUniversalPacketReceiver *IbcUniversalPacketReceiverTransactorSession) OnUniversalAcknowledgement(channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcUniversalPacketReceiver.Contract.OnUniversalAcknowledgement(&_IbcUniversalPacketReceiver.TransactOpts, channelId, packet, ack)
}
