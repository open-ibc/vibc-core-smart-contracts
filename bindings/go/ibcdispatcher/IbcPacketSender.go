// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcdispatcher

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

// IbcPacketSenderMetaData contains all meta data concerning the IbcPacketSender contract.
var IbcPacketSenderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"sendPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"}]",
}

// IbcPacketSenderABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcPacketSenderMetaData.ABI instead.
var IbcPacketSenderABI = IbcPacketSenderMetaData.ABI

// IbcPacketSender is an auto generated Go binding around an Ethereum contract.
type IbcPacketSender struct {
	IbcPacketSenderCaller     // Read-only binding to the contract
	IbcPacketSenderTransactor // Write-only binding to the contract
	IbcPacketSenderFilterer   // Log filterer for contract events
}

// IbcPacketSenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcPacketSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcPacketSenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcPacketSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcPacketSenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcPacketSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcPacketSenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcPacketSenderSession struct {
	Contract     *IbcPacketSender  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcPacketSenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcPacketSenderCallerSession struct {
	Contract *IbcPacketSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IbcPacketSenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcPacketSenderTransactorSession struct {
	Contract     *IbcPacketSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IbcPacketSenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcPacketSenderRaw struct {
	Contract *IbcPacketSender // Generic contract binding to access the raw methods on
}

// IbcPacketSenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcPacketSenderCallerRaw struct {
	Contract *IbcPacketSenderCaller // Generic read-only contract binding to access the raw methods on
}

// IbcPacketSenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcPacketSenderTransactorRaw struct {
	Contract *IbcPacketSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcPacketSender creates a new instance of IbcPacketSender, bound to a specific deployed contract.
func NewIbcPacketSender(address common.Address, backend bind.ContractBackend) (*IbcPacketSender, error) {
	contract, err := bindIbcPacketSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcPacketSender{IbcPacketSenderCaller: IbcPacketSenderCaller{contract: contract}, IbcPacketSenderTransactor: IbcPacketSenderTransactor{contract: contract}, IbcPacketSenderFilterer: IbcPacketSenderFilterer{contract: contract}}, nil
}

// NewIbcPacketSenderCaller creates a new read-only instance of IbcPacketSender, bound to a specific deployed contract.
func NewIbcPacketSenderCaller(address common.Address, caller bind.ContractCaller) (*IbcPacketSenderCaller, error) {
	contract, err := bindIbcPacketSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcPacketSenderCaller{contract: contract}, nil
}

// NewIbcPacketSenderTransactor creates a new write-only instance of IbcPacketSender, bound to a specific deployed contract.
func NewIbcPacketSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcPacketSenderTransactor, error) {
	contract, err := bindIbcPacketSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcPacketSenderTransactor{contract: contract}, nil
}

// NewIbcPacketSenderFilterer creates a new log filterer instance of IbcPacketSender, bound to a specific deployed contract.
func NewIbcPacketSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcPacketSenderFilterer, error) {
	contract, err := bindIbcPacketSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcPacketSenderFilterer{contract: contract}, nil
}

// bindIbcPacketSender binds a generic wrapper to an already deployed contract.
func bindIbcPacketSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcPacketSenderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcPacketSender *IbcPacketSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcPacketSender.Contract.IbcPacketSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcPacketSender *IbcPacketSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcPacketSender.Contract.IbcPacketSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcPacketSender *IbcPacketSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcPacketSender.Contract.IbcPacketSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcPacketSender *IbcPacketSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcPacketSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcPacketSender *IbcPacketSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcPacketSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcPacketSender *IbcPacketSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcPacketSender.Contract.contract.Transact(opts, method, params...)
}

// SendPacket is a paid mutator transaction binding the contract method 0xc3e1155c.
//
// Solidity: function sendPacket(bytes32 channelId, bytes payload, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcPacketSender *IbcPacketSenderTransactor) SendPacket(opts *bind.TransactOpts, channelId [32]byte, payload []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcPacketSender.contract.Transact(opts, "sendPacket", channelId, payload, timeoutTimestamp)
}

// SendPacket is a paid mutator transaction binding the contract method 0xc3e1155c.
//
// Solidity: function sendPacket(bytes32 channelId, bytes payload, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcPacketSender *IbcPacketSenderSession) SendPacket(channelId [32]byte, payload []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcPacketSender.Contract.SendPacket(&_IbcPacketSender.TransactOpts, channelId, payload, timeoutTimestamp)
}

// SendPacket is a paid mutator transaction binding the contract method 0xc3e1155c.
//
// Solidity: function sendPacket(bytes32 channelId, bytes payload, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcPacketSender *IbcPacketSenderTransactorSession) SendPacket(channelId [32]byte, payload []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcPacketSender.Contract.SendPacket(&_IbcPacketSender.TransactOpts, channelId, payload, timeoutTimestamp)
}
