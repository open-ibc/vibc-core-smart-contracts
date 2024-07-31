// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iuniversalchannelhandler

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

// IUniversalChannelHandlerMetaData contains all meta data concerning the IUniversalChannelHandler contract.
var IUniversalChannelHandlerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MW_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"MWID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"closeChannel\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"connectedChannels\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"channel\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dispatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"dispatcher\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseInit\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartychannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenTry\",\"inputs\":[{\"name\":\"order\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortIdentifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartychannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"openChannel\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortIdentifier\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerMwStack\",\"inputs\":[{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mwAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacketWithFee\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setDispatcher\",\"inputs\":[{\"name\":\"dispatcher\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"MwBitmpaCannotBeZero\",\"inputs\":[]}]",
}

// IUniversalChannelHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use IUniversalChannelHandlerMetaData.ABI instead.
var IUniversalChannelHandlerABI = IUniversalChannelHandlerMetaData.ABI

// IUniversalChannelHandler is an auto generated Go binding around an Ethereum contract.
type IUniversalChannelHandler struct {
	IUniversalChannelHandlerCaller     // Read-only binding to the contract
	IUniversalChannelHandlerTransactor // Write-only binding to the contract
	IUniversalChannelHandlerFilterer   // Log filterer for contract events
}

// IUniversalChannelHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IUniversalChannelHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniversalChannelHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IUniversalChannelHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniversalChannelHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IUniversalChannelHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IUniversalChannelHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IUniversalChannelHandlerSession struct {
	Contract     *IUniversalChannelHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IUniversalChannelHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IUniversalChannelHandlerCallerSession struct {
	Contract *IUniversalChannelHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IUniversalChannelHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IUniversalChannelHandlerTransactorSession struct {
	Contract     *IUniversalChannelHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IUniversalChannelHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IUniversalChannelHandlerRaw struct {
	Contract *IUniversalChannelHandler // Generic contract binding to access the raw methods on
}

// IUniversalChannelHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IUniversalChannelHandlerCallerRaw struct {
	Contract *IUniversalChannelHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IUniversalChannelHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IUniversalChannelHandlerTransactorRaw struct {
	Contract *IUniversalChannelHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIUniversalChannelHandler creates a new instance of IUniversalChannelHandler, bound to a specific deployed contract.
func NewIUniversalChannelHandler(address common.Address, backend bind.ContractBackend) (*IUniversalChannelHandler, error) {
	contract, err := bindIUniversalChannelHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IUniversalChannelHandler{IUniversalChannelHandlerCaller: IUniversalChannelHandlerCaller{contract: contract}, IUniversalChannelHandlerTransactor: IUniversalChannelHandlerTransactor{contract: contract}, IUniversalChannelHandlerFilterer: IUniversalChannelHandlerFilterer{contract: contract}}, nil
}

// NewIUniversalChannelHandlerCaller creates a new read-only instance of IUniversalChannelHandler, bound to a specific deployed contract.
func NewIUniversalChannelHandlerCaller(address common.Address, caller bind.ContractCaller) (*IUniversalChannelHandlerCaller, error) {
	contract, err := bindIUniversalChannelHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IUniversalChannelHandlerCaller{contract: contract}, nil
}

// NewIUniversalChannelHandlerTransactor creates a new write-only instance of IUniversalChannelHandler, bound to a specific deployed contract.
func NewIUniversalChannelHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IUniversalChannelHandlerTransactor, error) {
	contract, err := bindIUniversalChannelHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IUniversalChannelHandlerTransactor{contract: contract}, nil
}

// NewIUniversalChannelHandlerFilterer creates a new log filterer instance of IUniversalChannelHandler, bound to a specific deployed contract.
func NewIUniversalChannelHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IUniversalChannelHandlerFilterer, error) {
	contract, err := bindIUniversalChannelHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IUniversalChannelHandlerFilterer{contract: contract}, nil
}

// bindIUniversalChannelHandler binds a generic wrapper to an already deployed contract.
func bindIUniversalChannelHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IUniversalChannelHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniversalChannelHandler *IUniversalChannelHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniversalChannelHandler.Contract.IUniversalChannelHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniversalChannelHandler *IUniversalChannelHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.IUniversalChannelHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniversalChannelHandler *IUniversalChannelHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.IUniversalChannelHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IUniversalChannelHandler *IUniversalChannelHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IUniversalChannelHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.contract.Transact(opts, method, params...)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256 MWID)
func (_IUniversalChannelHandler *IUniversalChannelHandlerCaller) MWID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IUniversalChannelHandler.contract.Call(opts, &out, "MW_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256 MWID)
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) MWID() (*big.Int, error) {
	return _IUniversalChannelHandler.Contract.MWID(&_IUniversalChannelHandler.CallOpts)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256 MWID)
func (_IUniversalChannelHandler *IUniversalChannelHandlerCallerSession) MWID() (*big.Int, error) {
	return _IUniversalChannelHandler.Contract.MWID(&_IUniversalChannelHandler.CallOpts)
}

// ConnectedChannels is a free data retrieval call binding the contract method 0xbb3f9f8d.
//
// Solidity: function connectedChannels(uint256 ) view returns(bytes32 channel)
func (_IUniversalChannelHandler *IUniversalChannelHandlerCaller) ConnectedChannels(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IUniversalChannelHandler.contract.Call(opts, &out, "connectedChannels", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConnectedChannels is a free data retrieval call binding the contract method 0xbb3f9f8d.
//
// Solidity: function connectedChannels(uint256 ) view returns(bytes32 channel)
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) ConnectedChannels(arg0 *big.Int) ([32]byte, error) {
	return _IUniversalChannelHandler.Contract.ConnectedChannels(&_IUniversalChannelHandler.CallOpts, arg0)
}

// ConnectedChannels is a free data retrieval call binding the contract method 0xbb3f9f8d.
//
// Solidity: function connectedChannels(uint256 ) view returns(bytes32 channel)
func (_IUniversalChannelHandler *IUniversalChannelHandlerCallerSession) ConnectedChannels(arg0 *big.Int) ([32]byte, error) {
	return _IUniversalChannelHandler.Contract.ConnectedChannels(&_IUniversalChannelHandler.CallOpts, arg0)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x4c2ee09d.
//
// Solidity: function closeChannel(bytes32 channelId) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) CloseChannel(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "closeChannel", channelId)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x4c2ee09d.
//
// Solidity: function closeChannel(bytes32 channelId) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) CloseChannel(channelId [32]byte) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.CloseChannel(&_IUniversalChannelHandler.TransactOpts, channelId)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x4c2ee09d.
//
// Solidity: function closeChannel(bytes32 channelId) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) CloseChannel(channelId [32]byte) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.CloseChannel(&_IUniversalChannelHandler.TransactOpts, channelId)
}

// Dispatcher is a paid mutator transaction binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() returns(address dispatcher)
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) Dispatcher(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "dispatcher")
}

// Dispatcher is a paid mutator transaction binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() returns(address dispatcher)
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) Dispatcher() (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.Dispatcher(&_IUniversalChannelHandler.TransactOpts)
}

// Dispatcher is a paid mutator transaction binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() returns(address dispatcher)
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) Dispatcher() (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.Dispatcher(&_IUniversalChannelHandler.TransactOpts)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "onAcknowledgementPacket", packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnAcknowledgementPacket(&_IUniversalChannelHandler.TransactOpts, packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnAcknowledgementPacket(&_IUniversalChannelHandler.TransactOpts, packet, ack)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) OnChanCloseConfirm(opts *bind.TransactOpts, channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "onChanCloseConfirm", channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) OnChanCloseConfirm(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnChanCloseConfirm(&_IUniversalChannelHandler.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) OnChanCloseConfirm(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnChanCloseConfirm(&_IUniversalChannelHandler.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) OnChanCloseInit(opts *bind.TransactOpts, channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "onChanCloseInit", channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) OnChanCloseInit(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnChanCloseInit(&_IUniversalChannelHandler.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) OnChanCloseInit(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnChanCloseInit(&_IUniversalChannelHandler.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string counterpartyVersion) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) OnChanOpenAck(opts *bind.TransactOpts, channelId [32]byte, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "onChanOpenAck", channelId, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string counterpartyVersion) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) OnChanOpenAck(channelId [32]byte, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnChanOpenAck(&_IUniversalChannelHandler.TransactOpts, channelId, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string counterpartyVersion) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) OnChanOpenAck(channelId [32]byte, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnChanOpenAck(&_IUniversalChannelHandler.TransactOpts, channelId, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "onChanOpenConfirm", channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnChanOpenConfirm(&_IUniversalChannelHandler.TransactOpts, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnChanOpenConfirm(&_IUniversalChannelHandler.TransactOpts, channelId)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 order, string[] connectionHops, bytes32 channelId, string counterpartyPortIdentifier, bytes32 counterpartychannelId, string counterpartyVersion) returns(string selectedVersion)
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) OnChanOpenTry(opts *bind.TransactOpts, order uint8, connectionHops []string, channelId [32]byte, counterpartyPortIdentifier string, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "onChanOpenTry", order, connectionHops, channelId, counterpartyPortIdentifier, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 order, string[] connectionHops, bytes32 channelId, string counterpartyPortIdentifier, bytes32 counterpartychannelId, string counterpartyVersion) returns(string selectedVersion)
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) OnChanOpenTry(order uint8, connectionHops []string, channelId [32]byte, counterpartyPortIdentifier string, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnChanOpenTry(&_IUniversalChannelHandler.TransactOpts, order, connectionHops, channelId, counterpartyPortIdentifier, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 order, string[] connectionHops, bytes32 channelId, string counterpartyPortIdentifier, bytes32 counterpartychannelId, string counterpartyVersion) returns(string selectedVersion)
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) OnChanOpenTry(order uint8, connectionHops []string, channelId [32]byte, counterpartyPortIdentifier string, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnChanOpenTry(&_IUniversalChannelHandler.TransactOpts, order, connectionHops, channelId, counterpartyPortIdentifier, counterpartychannelId, counterpartyVersion)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) OnRecvPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "onRecvPacket", packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnRecvPacket(&_IUniversalChannelHandler.TransactOpts, packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnRecvPacket(&_IUniversalChannelHandler.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) OnTimeoutPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "onTimeoutPacket", packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnTimeoutPacket(&_IUniversalChannelHandler.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OnTimeoutPacket(&_IUniversalChannelHandler.TransactOpts, packet)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xace02de7.
//
// Solidity: function openChannel(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortIdentifier) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) OpenChannel(opts *bind.TransactOpts, version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortIdentifier string) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "openChannel", version, ordering, feeEnabled, connectionHops, counterpartyPortIdentifier)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xace02de7.
//
// Solidity: function openChannel(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortIdentifier) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) OpenChannel(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortIdentifier string) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OpenChannel(&_IUniversalChannelHandler.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortIdentifier)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xace02de7.
//
// Solidity: function openChannel(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortIdentifier) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) OpenChannel(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortIdentifier string) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.OpenChannel(&_IUniversalChannelHandler.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortIdentifier)
}

// RegisterMwStack is a paid mutator transaction binding the contract method 0x1b532db1.
//
// Solidity: function registerMwStack(uint256 mwBitmap, address[] mwAddrs) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) RegisterMwStack(opts *bind.TransactOpts, mwBitmap *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "registerMwStack", mwBitmap, mwAddrs)
}

// RegisterMwStack is a paid mutator transaction binding the contract method 0x1b532db1.
//
// Solidity: function registerMwStack(uint256 mwBitmap, address[] mwAddrs) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) RegisterMwStack(mwBitmap *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.RegisterMwStack(&_IUniversalChannelHandler.TransactOpts, mwBitmap, mwAddrs)
}

// RegisterMwStack is a paid mutator transaction binding the contract method 0x1b532db1.
//
// Solidity: function registerMwStack(uint256 mwBitmap, address[] mwAddrs) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) RegisterMwStack(mwBitmap *big.Int, mwAddrs []common.Address) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.RegisterMwStack(&_IUniversalChannelHandler.TransactOpts, mwBitmap, mwAddrs)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) SendUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "sendUniversalPacket", channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.SendUniversalPacket(&_IUniversalChannelHandler.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.SendUniversalPacket(&_IUniversalChannelHandler.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) SendUniversalPacketWithFee(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "sendUniversalPacketWithFee", channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.SendUniversalPacketWithFee(&_IUniversalChannelHandler.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.SendUniversalPacketWithFee(&_IUniversalChannelHandler.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address dispatcher) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactor) SetDispatcher(opts *bind.TransactOpts, dispatcher common.Address) (*types.Transaction, error) {
	return _IUniversalChannelHandler.contract.Transact(opts, "setDispatcher", dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address dispatcher) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerSession) SetDispatcher(dispatcher common.Address) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.SetDispatcher(&_IUniversalChannelHandler.TransactOpts, dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address dispatcher) returns()
func (_IUniversalChannelHandler *IUniversalChannelHandlerTransactorSession) SetDispatcher(dispatcher common.Address) (*types.Transaction, error) {
	return _IUniversalChannelHandler.Contract.SetDispatcher(&_IUniversalChannelHandler.TransactOpts, dispatcher)
}
