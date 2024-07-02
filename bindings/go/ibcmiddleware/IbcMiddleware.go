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

// IbcMiddlewareMetaData contains all meta data concerning the IbcMiddleware contract.
var IbcMiddlewareMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MW_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"MWID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onRecvMWAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"mwIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mwAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvMWPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"mwIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mwAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvMWTimeout\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"mwIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mwAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"ucPacket\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onUniversalAcknowledgement\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacketWithFee\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"error\",\"name\":\"invalidChannelId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"unauthorizedChannel\",\"inputs\":[]}]",
}

// IbcMiddlewareABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcMiddlewareMetaData.ABI instead.
var IbcMiddlewareABI = IbcMiddlewareMetaData.ABI

// IbcMiddleware is an auto generated Go binding around an Ethereum contract.
type IbcMiddleware struct {
	IbcMiddlewareCaller     // Read-only binding to the contract
	IbcMiddlewareTransactor // Write-only binding to the contract
	IbcMiddlewareFilterer   // Log filterer for contract events
}

// IbcMiddlewareCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcMiddlewareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMiddlewareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcMiddlewareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMiddlewareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcMiddlewareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMiddlewareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcMiddlewareSession struct {
	Contract     *IbcMiddleware    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcMiddlewareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcMiddlewareCallerSession struct {
	Contract *IbcMiddlewareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IbcMiddlewareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcMiddlewareTransactorSession struct {
	Contract     *IbcMiddlewareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IbcMiddlewareRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcMiddlewareRaw struct {
	Contract *IbcMiddleware // Generic contract binding to access the raw methods on
}

// IbcMiddlewareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcMiddlewareCallerRaw struct {
	Contract *IbcMiddlewareCaller // Generic read-only contract binding to access the raw methods on
}

// IbcMiddlewareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcMiddlewareTransactorRaw struct {
	Contract *IbcMiddlewareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcMiddleware creates a new instance of IbcMiddleware, bound to a specific deployed contract.
func NewIbcMiddleware(address common.Address, backend bind.ContractBackend) (*IbcMiddleware, error) {
	contract, err := bindIbcMiddleware(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcMiddleware{IbcMiddlewareCaller: IbcMiddlewareCaller{contract: contract}, IbcMiddlewareTransactor: IbcMiddlewareTransactor{contract: contract}, IbcMiddlewareFilterer: IbcMiddlewareFilterer{contract: contract}}, nil
}

// NewIbcMiddlewareCaller creates a new read-only instance of IbcMiddleware, bound to a specific deployed contract.
func NewIbcMiddlewareCaller(address common.Address, caller bind.ContractCaller) (*IbcMiddlewareCaller, error) {
	contract, err := bindIbcMiddleware(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcMiddlewareCaller{contract: contract}, nil
}

// NewIbcMiddlewareTransactor creates a new write-only instance of IbcMiddleware, bound to a specific deployed contract.
func NewIbcMiddlewareTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcMiddlewareTransactor, error) {
	contract, err := bindIbcMiddleware(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcMiddlewareTransactor{contract: contract}, nil
}

// NewIbcMiddlewareFilterer creates a new log filterer instance of IbcMiddleware, bound to a specific deployed contract.
func NewIbcMiddlewareFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcMiddlewareFilterer, error) {
	contract, err := bindIbcMiddleware(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcMiddlewareFilterer{contract: contract}, nil
}

// bindIbcMiddleware binds a generic wrapper to an already deployed contract.
func bindIbcMiddleware(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcMiddlewareMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcMiddleware *IbcMiddlewareRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcMiddleware.Contract.IbcMiddlewareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcMiddleware *IbcMiddlewareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.IbcMiddlewareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcMiddleware *IbcMiddlewareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.IbcMiddlewareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcMiddleware *IbcMiddlewareCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcMiddleware.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcMiddleware *IbcMiddlewareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcMiddleware *IbcMiddlewareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.contract.Transact(opts, method, params...)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256 MWID)
func (_IbcMiddleware *IbcMiddlewareCaller) MWID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IbcMiddleware.contract.Call(opts, &out, "MW_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256 MWID)
func (_IbcMiddleware *IbcMiddlewareSession) MWID() (*big.Int, error) {
	return _IbcMiddleware.Contract.MWID(&_IbcMiddleware.CallOpts)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256 MWID)
func (_IbcMiddleware *IbcMiddlewareCallerSession) MWID() (*big.Int, error) {
	return _IbcMiddleware.Contract.MWID(&_IbcMiddleware.CallOpts)
}

// OnRecvMWAck is a paid mutator transaction binding the contract method 0xa968dc60.
//
// Solidity: function onRecvMWAck(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs, (bool,bytes) ack) returns()
func (_IbcMiddleware *IbcMiddlewareTransactor) OnRecvMWAck(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address, ack AckPacket) (*types.Transaction, error) {
	return _IbcMiddleware.contract.Transact(opts, "onRecvMWAck", channelId, ucPacket, mwIndex, mwAddrs, ack)
}

// OnRecvMWAck is a paid mutator transaction binding the contract method 0xa968dc60.
//
// Solidity: function onRecvMWAck(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs, (bool,bytes) ack) returns()
func (_IbcMiddleware *IbcMiddlewareSession) OnRecvMWAck(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address, ack AckPacket) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.OnRecvMWAck(&_IbcMiddleware.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs, ack)
}

// OnRecvMWAck is a paid mutator transaction binding the contract method 0xa968dc60.
//
// Solidity: function onRecvMWAck(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs, (bool,bytes) ack) returns()
func (_IbcMiddleware *IbcMiddlewareTransactorSession) OnRecvMWAck(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address, ack AckPacket) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.OnRecvMWAck(&_IbcMiddleware.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs, ack)
}

// OnRecvMWPacket is a paid mutator transaction binding the contract method 0xce2ce646.
//
// Solidity: function onRecvMWPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns((bool,bytes) ackPacket)
func (_IbcMiddleware *IbcMiddlewareTransactor) OnRecvMWPacket(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IbcMiddleware.contract.Transact(opts, "onRecvMWPacket", channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWPacket is a paid mutator transaction binding the contract method 0xce2ce646.
//
// Solidity: function onRecvMWPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns((bool,bytes) ackPacket)
func (_IbcMiddleware *IbcMiddlewareSession) OnRecvMWPacket(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.OnRecvMWPacket(&_IbcMiddleware.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWPacket is a paid mutator transaction binding the contract method 0xce2ce646.
//
// Solidity: function onRecvMWPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns((bool,bytes) ackPacket)
func (_IbcMiddleware *IbcMiddlewareTransactorSession) OnRecvMWPacket(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.OnRecvMWPacket(&_IbcMiddleware.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWTimeout is a paid mutator transaction binding the contract method 0xc2cfcd73.
//
// Solidity: function onRecvMWTimeout(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns()
func (_IbcMiddleware *IbcMiddlewareTransactor) OnRecvMWTimeout(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IbcMiddleware.contract.Transact(opts, "onRecvMWTimeout", channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWTimeout is a paid mutator transaction binding the contract method 0xc2cfcd73.
//
// Solidity: function onRecvMWTimeout(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns()
func (_IbcMiddleware *IbcMiddlewareSession) OnRecvMWTimeout(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.OnRecvMWTimeout(&_IbcMiddleware.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvMWTimeout is a paid mutator transaction binding the contract method 0xc2cfcd73.
//
// Solidity: function onRecvMWTimeout(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket, uint256 mwIndex, address[] mwAddrs) returns()
func (_IbcMiddleware *IbcMiddlewareTransactorSession) OnRecvMWTimeout(channelId [32]byte, ucPacket UniversalPacket, mwIndex *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.OnRecvMWTimeout(&_IbcMiddleware.TransactOpts, channelId, ucPacket, mwIndex, mwAddrs)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket) returns((bool,bytes) ackPacket)
func (_IbcMiddleware *IbcMiddlewareTransactor) OnRecvUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, ucPacket UniversalPacket) (*types.Transaction, error) {
	return _IbcMiddleware.contract.Transact(opts, "onRecvUniversalPacket", channelId, ucPacket)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket) returns((bool,bytes) ackPacket)
func (_IbcMiddleware *IbcMiddlewareSession) OnRecvUniversalPacket(channelId [32]byte, ucPacket UniversalPacket) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.OnRecvUniversalPacket(&_IbcMiddleware.TransactOpts, channelId, ucPacket)
}

// OnRecvUniversalPacket is a paid mutator transaction binding the contract method 0x5b761585.
//
// Solidity: function onRecvUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) ucPacket) returns((bool,bytes) ackPacket)
func (_IbcMiddleware *IbcMiddlewareTransactorSession) OnRecvUniversalPacket(channelId [32]byte, ucPacket UniversalPacket) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.OnRecvUniversalPacket(&_IbcMiddleware.TransactOpts, channelId, ucPacket)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_IbcMiddleware *IbcMiddlewareTransactor) OnTimeoutUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _IbcMiddleware.contract.Transact(opts, "onTimeoutUniversalPacket", channelId, packet)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_IbcMiddleware *IbcMiddlewareSession) OnTimeoutUniversalPacket(channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.OnTimeoutUniversalPacket(&_IbcMiddleware.TransactOpts, channelId, packet)
}

// OnTimeoutUniversalPacket is a paid mutator transaction binding the contract method 0x400d9f5d.
//
// Solidity: function onTimeoutUniversalPacket(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet) returns()
func (_IbcMiddleware *IbcMiddlewareTransactorSession) OnTimeoutUniversalPacket(channelId [32]byte, packet UniversalPacket) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.OnTimeoutUniversalPacket(&_IbcMiddleware.TransactOpts, channelId, packet)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_IbcMiddleware *IbcMiddlewareTransactor) OnUniversalAcknowledgement(opts *bind.TransactOpts, channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcMiddleware.contract.Transact(opts, "onUniversalAcknowledgement", channelId, packet, ack)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_IbcMiddleware *IbcMiddlewareSession) OnUniversalAcknowledgement(channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.OnUniversalAcknowledgement(&_IbcMiddleware.TransactOpts, channelId, packet, ack)
}

// OnUniversalAcknowledgement is a paid mutator transaction binding the contract method 0x588152ca.
//
// Solidity: function onUniversalAcknowledgement(bytes32 channelId, (bytes32,uint256,bytes32,bytes) packet, (bool,bytes) ack) returns()
func (_IbcMiddleware *IbcMiddlewareTransactorSession) OnUniversalAcknowledgement(channelId [32]byte, packet UniversalPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.OnUniversalAcknowledgement(&_IbcMiddleware.TransactOpts, channelId, packet, ack)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcMiddleware *IbcMiddlewareTransactor) SendUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcMiddleware.contract.Transact(opts, "sendUniversalPacket", channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcMiddleware *IbcMiddlewareSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.SendUniversalPacket(&_IbcMiddleware.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcMiddleware *IbcMiddlewareTransactorSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.SendUniversalPacket(&_IbcMiddleware.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IbcMiddleware *IbcMiddlewareTransactor) SendUniversalPacketWithFee(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IbcMiddleware.contract.Transact(opts, "sendUniversalPacketWithFee", channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IbcMiddleware *IbcMiddlewareSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.SendUniversalPacketWithFee(&_IbcMiddleware.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IbcMiddleware *IbcMiddlewareTransactorSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IbcMiddleware.Contract.SendUniversalPacketWithFee(&_IbcMiddleware.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}
