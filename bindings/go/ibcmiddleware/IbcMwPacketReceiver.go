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

// IbcMwPacketReceiverMetaData contains all meta data concerning the IbcMwPacketReceiver contract.
var IbcMwPacketReceiverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"onRecvMWAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"mwIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mwAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvMWPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"mwIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mwAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvMWTimeout\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"mwIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mwAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IbcMwPacketReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcMwPacketReceiverMetaData.ABI instead.
var IbcMwPacketReceiverABI = IbcMwPacketReceiverMetaData.ABI

// IbcMwPacketReceiver is an auto generated Go binding around an Ethereum contract.
type IbcMwPacketReceiver struct {
	IbcMwPacketReceiverCaller     // Read-only binding to the contract
	IbcMwPacketReceiverTransactor // Write-only binding to the contract
	IbcMwPacketReceiverFilterer   // Log filterer for contract events
}

// IbcMwPacketReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcMwPacketReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMwPacketReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcMwPacketReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMwPacketReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcMwPacketReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMwPacketReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcMwPacketReceiverSession struct {
	Contract     *IbcMwPacketReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IbcMwPacketReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcMwPacketReceiverCallerSession struct {
	Contract *IbcMwPacketReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IbcMwPacketReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcMwPacketReceiverTransactorSession struct {
	Contract     *IbcMwPacketReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IbcMwPacketReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcMwPacketReceiverRaw struct {
	Contract *IbcMwPacketReceiver // Generic contract binding to access the raw methods on
}

// IbcMwPacketReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcMwPacketReceiverCallerRaw struct {
	Contract *IbcMwPacketReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IbcMwPacketReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcMwPacketReceiverTransactorRaw struct {
	Contract *IbcMwPacketReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcMwPacketReceiver creates a new instance of IbcMwPacketReceiver, bound to a specific deployed contract.
func NewIbcMwPacketReceiver(address common.Address, backend bind.ContractBackend) (*IbcMwPacketReceiver, error) {
	contract, err := bindIbcMwPacketReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcMwPacketReceiver{IbcMwPacketReceiverCaller: IbcMwPacketReceiverCaller{contract: contract}, IbcMwPacketReceiverTransactor: IbcMwPacketReceiverTransactor{contract: contract}, IbcMwPacketReceiverFilterer: IbcMwPacketReceiverFilterer{contract: contract}}, nil
}

// NewIbcMwPacketReceiverCaller creates a new read-only instance of IbcMwPacketReceiver, bound to a specific deployed contract.
func NewIbcMwPacketReceiverCaller(address common.Address, caller bind.ContractCaller) (*IbcMwPacketReceiverCaller, error) {
	contract, err := bindIbcMwPacketReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcMwPacketReceiverCaller{contract: contract}, nil
}

// NewIbcMwPacketReceiverTransactor creates a new write-only instance of IbcMwPacketReceiver, bound to a specific deployed contract.
func NewIbcMwPacketReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcMwPacketReceiverTransactor, error) {
	contract, err := bindIbcMwPacketReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcMwPacketReceiverTransactor{contract: contract}, nil
}

// NewIbcMwPacketReceiverFilterer creates a new log filterer instance of IbcMwPacketReceiver, bound to a specific deployed contract.
func NewIbcMwPacketReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcMwPacketReceiverFilterer, error) {
	contract, err := bindIbcMwPacketReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcMwPacketReceiverFilterer{contract: contract}, nil
}

// bindIbcMwPacketReceiver binds a generic wrapper to an already deployed contract.
func bindIbcMwPacketReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcMwPacketReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcMwPacketReceiver *IbcMwPacketReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcMwPacketReceiver.Contract.IbcMwPacketReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcMwPacketReceiver *IbcMwPacketReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.Contract.IbcMwPacketReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcMwPacketReceiver *IbcMwPacketReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.Contract.IbcMwPacketReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcMwPacketReceiver *IbcMwPacketReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcMwPacketReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcMwPacketReceiver *IbcMwPacketReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcMwPacketReceiver *IbcMwPacketReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.Contract.contract.Transact(opts, method, params...)
}

// OnRecvMWAck is a paid mutator transaction binding the contract method 0xa968dc60.
//
// Solidity: function onRecvMWAck(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs, (bool,bytes) ack) returns()
func (_IbcMwPacketReceiver *IbcMwPacketReceiverTransactor) OnRecvMWAck(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address, ack AckPacket) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.contract.Transact(opts, "onRecvMWAck", channelId, ucPacket, mwIndex, mwAddrs, ack)
}

// OnRecvMWAck is a paid mutator transaction binding the contract method 0xa968dc60.
//
// Solidity: function onRecvMWAck(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs, (bool,bytes) ack) returns()
func (_IbcMwPacketReceiver *IbcMwPacketReceiverSession) OnRecvMWAck(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address, ack AckPacket) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.Contract.OnRecvMWAck(&_IbcMwPacketReceiver.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs, ack)
}

// OnRecvMWAck is a paid mutator transaction binding the contract method 0xa968dc60.
//
// Solidity: function onRecvMWAck(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs, (bool,bytes) ack) returns()
func (_IbcMwPacketReceiver *IbcMwPacketReceiverTransactorSession) OnRecvMWAck(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address, ack AckPacket) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.Contract.OnRecvMWAck(&_IbcMwPacketReceiver.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs, ack)
}

// OnRecvMWPacket is a paid mutator transaction binding the contract method 0xce2ce646.
//
// Solidity: function onRecvMWPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns((bool,bytes) ackPacket)
func (_IbcMwPacketReceiver *IbcMwPacketReceiverTransactor) OnRecvMWPacket(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.contract.Transact(opts, "onRecvMWPacket", channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWPacket is a paid mutator transaction binding the contract method 0xce2ce646.
//
// Solidity: function onRecvMWPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns((bool,bytes) ackPacket)
func (_IbcMwPacketReceiver *IbcMwPacketReceiverSession) OnRecvMWPacket(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.Contract.OnRecvMWPacket(&_IbcMwPacketReceiver.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWPacket is a paid mutator transaction binding the contract method 0xce2ce646.
//
// Solidity: function onRecvMWPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns((bool,bytes) ackPacket)
func (_IbcMwPacketReceiver *IbcMwPacketReceiverTransactorSession) OnRecvMWPacket(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.Contract.OnRecvMWPacket(&_IbcMwPacketReceiver.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWTimeout is a paid mutator transaction binding the contract method 0xc2cfcd73.
//
// Solidity: function onRecvMWTimeout(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns()
func (_IbcMwPacketReceiver *IbcMwPacketReceiverTransactor) OnRecvMWTimeout(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.contract.Transact(opts, "onRecvMWTimeout", channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWTimeout is a paid mutator transaction binding the contract method 0xc2cfcd73.
//
// Solidity: function onRecvMWTimeout(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns()
func (_IbcMwPacketReceiver *IbcMwPacketReceiverSession) OnRecvMWTimeout(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.Contract.OnRecvMWTimeout(&_IbcMwPacketReceiver.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWTimeout is a paid mutator transaction binding the contract method 0xc2cfcd73.
//
// Solidity: function onRecvMWTimeout(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns()
func (_IbcMwPacketReceiver *IbcMwPacketReceiverTransactorSession) OnRecvMWTimeout(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IbcMwPacketReceiver.Contract.OnRecvMWTimeout(&_IbcMwPacketReceiver.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs)
}
