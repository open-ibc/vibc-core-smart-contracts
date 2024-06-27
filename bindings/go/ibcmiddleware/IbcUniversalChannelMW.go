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

// IbcUniversalChannelMWMetaData contains all meta data concerning the IbcUniversalChannelMW contract.
var IbcUniversalChannelMWMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MW_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"MWID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseInit\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartychannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenInit\",\"inputs\":[{\"name\":\"order\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortIdentifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenTry\",\"inputs\":[{\"name\":\"order\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyPortIdentifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartychannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacketWithFee\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"error\",\"name\":\"MwBitmpaCannotBeZero\",\"inputs\":[]}]",
}

// IbcUniversalChannelMWABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcUniversalChannelMWMetaData.ABI instead.
var IbcUniversalChannelMWABI = IbcUniversalChannelMWMetaData.ABI

// IbcUniversalChannelMW is an auto generated Go binding around an Ethereum contract.
type IbcUniversalChannelMW struct {
	IbcUniversalChannelMWCaller     // Read-only binding to the contract
	IbcUniversalChannelMWTransactor // Write-only binding to the contract
	IbcUniversalChannelMWFilterer   // Log filterer for contract events
}

// IbcUniversalChannelMWCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcUniversalChannelMWCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUniversalChannelMWTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcUniversalChannelMWTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUniversalChannelMWFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcUniversalChannelMWFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUniversalChannelMWSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcUniversalChannelMWSession struct {
	Contract     *IbcUniversalChannelMW // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IbcUniversalChannelMWCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcUniversalChannelMWCallerSession struct {
	Contract *IbcUniversalChannelMWCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IbcUniversalChannelMWTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcUniversalChannelMWTransactorSession struct {
	Contract     *IbcUniversalChannelMWTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IbcUniversalChannelMWRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcUniversalChannelMWRaw struct {
	Contract *IbcUniversalChannelMW // Generic contract binding to access the raw methods on
}

// IbcUniversalChannelMWCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcUniversalChannelMWCallerRaw struct {
	Contract *IbcUniversalChannelMWCaller // Generic read-only contract binding to access the raw methods on
}

// IbcUniversalChannelMWTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcUniversalChannelMWTransactorRaw struct {
	Contract *IbcUniversalChannelMWTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcUniversalChannelMW creates a new instance of IbcUniversalChannelMW, bound to a specific deployed contract.
func NewIbcUniversalChannelMW(address common.Address, backend bind.ContractBackend) (*IbcUniversalChannelMW, error) {
	contract, err := bindIbcUniversalChannelMW(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalChannelMW{IbcUniversalChannelMWCaller: IbcUniversalChannelMWCaller{contract: contract}, IbcUniversalChannelMWTransactor: IbcUniversalChannelMWTransactor{contract: contract}, IbcUniversalChannelMWFilterer: IbcUniversalChannelMWFilterer{contract: contract}}, nil
}

// NewIbcUniversalChannelMWCaller creates a new read-only instance of IbcUniversalChannelMW, bound to a specific deployed contract.
func NewIbcUniversalChannelMWCaller(address common.Address, caller bind.ContractCaller) (*IbcUniversalChannelMWCaller, error) {
	contract, err := bindIbcUniversalChannelMW(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalChannelMWCaller{contract: contract}, nil
}

// NewIbcUniversalChannelMWTransactor creates a new write-only instance of IbcUniversalChannelMW, bound to a specific deployed contract.
func NewIbcUniversalChannelMWTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcUniversalChannelMWTransactor, error) {
	contract, err := bindIbcUniversalChannelMW(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalChannelMWTransactor{contract: contract}, nil
}

// NewIbcUniversalChannelMWFilterer creates a new log filterer instance of IbcUniversalChannelMW, bound to a specific deployed contract.
func NewIbcUniversalChannelMWFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcUniversalChannelMWFilterer, error) {
	contract, err := bindIbcUniversalChannelMW(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalChannelMWFilterer{contract: contract}, nil
}

// bindIbcUniversalChannelMW binds a generic wrapper to an already deployed contract.
func bindIbcUniversalChannelMW(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcUniversalChannelMWMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcUniversalChannelMW *IbcUniversalChannelMWRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcUniversalChannelMW.Contract.IbcUniversalChannelMWCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcUniversalChannelMW *IbcUniversalChannelMWRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.IbcUniversalChannelMWTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcUniversalChannelMW *IbcUniversalChannelMWRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.IbcUniversalChannelMWTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcUniversalChannelMW *IbcUniversalChannelMWCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcUniversalChannelMW.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.contract.Transact(opts, method, params...)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256 MWID)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWCaller) MWID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IbcUniversalChannelMW.contract.Call(opts, &out, "MW_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256 MWID)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWSession) MWID() (*big.Int, error) {
	return _IbcUniversalChannelMW.Contract.MWID(&_IbcUniversalChannelMW.CallOpts)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256 MWID)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWCallerSession) MWID() (*big.Int, error) {
	return _IbcUniversalChannelMW.Contract.MWID(&_IbcUniversalChannelMW.CallOpts)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.contract.Transact(opts, "onAcknowledgementPacket", packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnAcknowledgementPacket(&_IbcUniversalChannelMW.TransactOpts, packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnAcknowledgementPacket(&_IbcUniversalChannelMW.TransactOpts, packet, ack)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactor) OnChanCloseConfirm(opts *bind.TransactOpts, channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.contract.Transact(opts, "onChanCloseConfirm", channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWSession) OnChanCloseConfirm(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnChanCloseConfirm(&_IbcUniversalChannelMW.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorSession) OnChanCloseConfirm(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnChanCloseConfirm(&_IbcUniversalChannelMW.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactor) OnChanCloseInit(opts *bind.TransactOpts, channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.contract.Transact(opts, "onChanCloseInit", channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWSession) OnChanCloseInit(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnChanCloseInit(&_IbcUniversalChannelMW.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string counterpartyPortId, bytes32 counterpartyChannelId) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorSession) OnChanCloseInit(channelId [32]byte, counterpartyPortId string, counterpartyChannelId [32]byte) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnChanCloseInit(&_IbcUniversalChannelMW.TransactOpts, channelId, counterpartyPortId, counterpartyChannelId)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string counterpartyVersion) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactor) OnChanOpenAck(opts *bind.TransactOpts, channelId [32]byte, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.contract.Transact(opts, "onChanOpenAck", channelId, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string counterpartyVersion) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWSession) OnChanOpenAck(channelId [32]byte, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnChanOpenAck(&_IbcUniversalChannelMW.TransactOpts, channelId, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenAck is a paid mutator transaction binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string counterpartyVersion) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorSession) OnChanOpenAck(channelId [32]byte, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnChanOpenAck(&_IbcUniversalChannelMW.TransactOpts, channelId, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.contract.Transact(opts, "onChanOpenConfirm", channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnChanOpenConfirm(&_IbcUniversalChannelMW.TransactOpts, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnChanOpenConfirm(&_IbcUniversalChannelMW.TransactOpts, channelId)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 order, string[] connectionHops, string counterpartyPortIdentifier, string version) returns(string selectedVersion)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactor) OnChanOpenInit(opts *bind.TransactOpts, order uint8, connectionHops []string, counterpartyPortIdentifier string, version string) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.contract.Transact(opts, "onChanOpenInit", order, connectionHops, counterpartyPortIdentifier, version)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 order, string[] connectionHops, string counterpartyPortIdentifier, string version) returns(string selectedVersion)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWSession) OnChanOpenInit(order uint8, connectionHops []string, counterpartyPortIdentifier string, version string) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnChanOpenInit(&_IbcUniversalChannelMW.TransactOpts, order, connectionHops, counterpartyPortIdentifier, version)
}

// OnChanOpenInit is a paid mutator transaction binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 order, string[] connectionHops, string counterpartyPortIdentifier, string version) returns(string selectedVersion)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorSession) OnChanOpenInit(order uint8, connectionHops []string, counterpartyPortIdentifier string, version string) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnChanOpenInit(&_IbcUniversalChannelMW.TransactOpts, order, connectionHops, counterpartyPortIdentifier, version)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 order, string[] connectionHops, bytes32 channelId, string counterpartyPortIdentifier, bytes32 counterpartychannelId, string counterpartyVersion) returns(string selectedVersion)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactor) OnChanOpenTry(opts *bind.TransactOpts, order uint8, connectionHops []string, channelId [32]byte, counterpartyPortIdentifier string, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.contract.Transact(opts, "onChanOpenTry", order, connectionHops, channelId, counterpartyPortIdentifier, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 order, string[] connectionHops, bytes32 channelId, string counterpartyPortIdentifier, bytes32 counterpartychannelId, string counterpartyVersion) returns(string selectedVersion)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWSession) OnChanOpenTry(order uint8, connectionHops []string, channelId [32]byte, counterpartyPortIdentifier string, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnChanOpenTry(&_IbcUniversalChannelMW.TransactOpts, order, connectionHops, channelId, counterpartyPortIdentifier, counterpartychannelId, counterpartyVersion)
}

// OnChanOpenTry is a paid mutator transaction binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 order, string[] connectionHops, bytes32 channelId, string counterpartyPortIdentifier, bytes32 counterpartychannelId, string counterpartyVersion) returns(string selectedVersion)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorSession) OnChanOpenTry(order uint8, connectionHops []string, channelId [32]byte, counterpartyPortIdentifier string, counterpartychannelId [32]byte, counterpartyVersion string) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnChanOpenTry(&_IbcUniversalChannelMW.TransactOpts, order, connectionHops, channelId, counterpartyPortIdentifier, counterpartychannelId, counterpartyVersion)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactor) OnRecvPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.contract.Transact(opts, "onRecvPacket", packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnRecvPacket(&_IbcUniversalChannelMW.TransactOpts, packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnRecvPacket(&_IbcUniversalChannelMW.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactor) OnTimeoutPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.contract.Transact(opts, "onTimeoutPacket", packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnTimeoutPacket(&_IbcUniversalChannelMW.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.OnTimeoutPacket(&_IbcUniversalChannelMW.TransactOpts, packet)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactor) SendUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.contract.Transact(opts, "sendUniversalPacket", channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.SendUniversalPacket(&_IbcUniversalChannelMW.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.SendUniversalPacket(&_IbcUniversalChannelMW.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactor) SendUniversalPacketWithFee(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.contract.Transact(opts, "sendUniversalPacketWithFee", channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.SendUniversalPacketWithFee(&_IbcUniversalChannelMW.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IbcUniversalChannelMW *IbcUniversalChannelMWTransactorSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IbcUniversalChannelMW.Contract.SendUniversalPacketWithFee(&_IbcUniversalChannelMW.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}
