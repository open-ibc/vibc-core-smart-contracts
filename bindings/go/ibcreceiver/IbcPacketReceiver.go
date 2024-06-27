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

// IbcPacketReceiverMetaData contains all meta data concerning the IbcPacketReceiver contract.
var IbcPacketReceiverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IbcPacketReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcPacketReceiverMetaData.ABI instead.
var IbcPacketReceiverABI = IbcPacketReceiverMetaData.ABI

// IbcPacketReceiver is an auto generated Go binding around an Ethereum contract.
type IbcPacketReceiver struct {
	IbcPacketReceiverCaller     // Read-only binding to the contract
	IbcPacketReceiverTransactor // Write-only binding to the contract
	IbcPacketReceiverFilterer   // Log filterer for contract events
}

// IbcPacketReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcPacketReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcPacketReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcPacketReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcPacketReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcPacketReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcPacketReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcPacketReceiverSession struct {
	Contract     *IbcPacketReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IbcPacketReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcPacketReceiverCallerSession struct {
	Contract *IbcPacketReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IbcPacketReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcPacketReceiverTransactorSession struct {
	Contract     *IbcPacketReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IbcPacketReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcPacketReceiverRaw struct {
	Contract *IbcPacketReceiver // Generic contract binding to access the raw methods on
}

// IbcPacketReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcPacketReceiverCallerRaw struct {
	Contract *IbcPacketReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IbcPacketReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcPacketReceiverTransactorRaw struct {
	Contract *IbcPacketReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcPacketReceiver creates a new instance of IbcPacketReceiver, bound to a specific deployed contract.
func NewIbcPacketReceiver(address common.Address, backend bind.ContractBackend) (*IbcPacketReceiver, error) {
	contract, err := bindIbcPacketReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcPacketReceiver{IbcPacketReceiverCaller: IbcPacketReceiverCaller{contract: contract}, IbcPacketReceiverTransactor: IbcPacketReceiverTransactor{contract: contract}, IbcPacketReceiverFilterer: IbcPacketReceiverFilterer{contract: contract}}, nil
}

// NewIbcPacketReceiverCaller creates a new read-only instance of IbcPacketReceiver, bound to a specific deployed contract.
func NewIbcPacketReceiverCaller(address common.Address, caller bind.ContractCaller) (*IbcPacketReceiverCaller, error) {
	contract, err := bindIbcPacketReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcPacketReceiverCaller{contract: contract}, nil
}

// NewIbcPacketReceiverTransactor creates a new write-only instance of IbcPacketReceiver, bound to a specific deployed contract.
func NewIbcPacketReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcPacketReceiverTransactor, error) {
	contract, err := bindIbcPacketReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcPacketReceiverTransactor{contract: contract}, nil
}

// NewIbcPacketReceiverFilterer creates a new log filterer instance of IbcPacketReceiver, bound to a specific deployed contract.
func NewIbcPacketReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcPacketReceiverFilterer, error) {
	contract, err := bindIbcPacketReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcPacketReceiverFilterer{contract: contract}, nil
}

// bindIbcPacketReceiver binds a generic wrapper to an already deployed contract.
func bindIbcPacketReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcPacketReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcPacketReceiver *IbcPacketReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcPacketReceiver.Contract.IbcPacketReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcPacketReceiver *IbcPacketReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcPacketReceiver.Contract.IbcPacketReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcPacketReceiver *IbcPacketReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcPacketReceiver.Contract.IbcPacketReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcPacketReceiver *IbcPacketReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcPacketReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcPacketReceiver *IbcPacketReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcPacketReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcPacketReceiver *IbcPacketReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcPacketReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IbcPacketReceiver *IbcPacketReceiverTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcPacketReceiver.contract.Transact(opts, "onAcknowledgementPacket", packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IbcPacketReceiver *IbcPacketReceiverSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcPacketReceiver.Contract.OnAcknowledgementPacket(&_IbcPacketReceiver.TransactOpts, packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IbcPacketReceiver *IbcPacketReceiverTransactorSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcPacketReceiver.Contract.OnAcknowledgementPacket(&_IbcPacketReceiver.TransactOpts, packet, ack)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_IbcPacketReceiver *IbcPacketReceiverTransactor) OnRecvPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _IbcPacketReceiver.contract.Transact(opts, "onRecvPacket", packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_IbcPacketReceiver *IbcPacketReceiverSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IbcPacketReceiver.Contract.OnRecvPacket(&_IbcPacketReceiver.TransactOpts, packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_IbcPacketReceiver *IbcPacketReceiverTransactorSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IbcPacketReceiver.Contract.OnRecvPacket(&_IbcPacketReceiver.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IbcPacketReceiver *IbcPacketReceiverTransactor) OnTimeoutPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _IbcPacketReceiver.contract.Transact(opts, "onTimeoutPacket", packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IbcPacketReceiver *IbcPacketReceiverSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IbcPacketReceiver.Contract.OnTimeoutPacket(&_IbcPacketReceiver.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IbcPacketReceiver *IbcPacketReceiverTransactorSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IbcPacketReceiver.Contract.OnTimeoutPacket(&_IbcPacketReceiver.TransactOpts, packet)
}
