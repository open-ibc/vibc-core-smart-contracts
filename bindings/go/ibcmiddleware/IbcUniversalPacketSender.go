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

// IbcUniversalPacketSenderMetaData contains all meta data concerning the IbcUniversalPacketSender contract.
var IbcUniversalPacketSenderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"sendUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacketWithFee\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"}]",
}

// IbcUniversalPacketSenderABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcUniversalPacketSenderMetaData.ABI instead.
var IbcUniversalPacketSenderABI = IbcUniversalPacketSenderMetaData.ABI

// IbcUniversalPacketSender is an auto generated Go binding around an Ethereum contract.
type IbcUniversalPacketSender struct {
	IbcUniversalPacketSenderCaller     // Read-only binding to the contract
	IbcUniversalPacketSenderTransactor // Write-only binding to the contract
	IbcUniversalPacketSenderFilterer   // Log filterer for contract events
}

// IbcUniversalPacketSenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcUniversalPacketSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUniversalPacketSenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcUniversalPacketSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUniversalPacketSenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcUniversalPacketSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUniversalPacketSenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcUniversalPacketSenderSession struct {
	Contract     *IbcUniversalPacketSender // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IbcUniversalPacketSenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcUniversalPacketSenderCallerSession struct {
	Contract *IbcUniversalPacketSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IbcUniversalPacketSenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcUniversalPacketSenderTransactorSession struct {
	Contract     *IbcUniversalPacketSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IbcUniversalPacketSenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcUniversalPacketSenderRaw struct {
	Contract *IbcUniversalPacketSender // Generic contract binding to access the raw methods on
}

// IbcUniversalPacketSenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcUniversalPacketSenderCallerRaw struct {
	Contract *IbcUniversalPacketSenderCaller // Generic read-only contract binding to access the raw methods on
}

// IbcUniversalPacketSenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcUniversalPacketSenderTransactorRaw struct {
	Contract *IbcUniversalPacketSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcUniversalPacketSender creates a new instance of IbcUniversalPacketSender, bound to a specific deployed contract.
func NewIbcUniversalPacketSender(address common.Address, backend bind.ContractBackend) (*IbcUniversalPacketSender, error) {
	contract, err := bindIbcUniversalPacketSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketSender{IbcUniversalPacketSenderCaller: IbcUniversalPacketSenderCaller{contract: contract}, IbcUniversalPacketSenderTransactor: IbcUniversalPacketSenderTransactor{contract: contract}, IbcUniversalPacketSenderFilterer: IbcUniversalPacketSenderFilterer{contract: contract}}, nil
}

// NewIbcUniversalPacketSenderCaller creates a new read-only instance of IbcUniversalPacketSender, bound to a specific deployed contract.
func NewIbcUniversalPacketSenderCaller(address common.Address, caller bind.ContractCaller) (*IbcUniversalPacketSenderCaller, error) {
	contract, err := bindIbcUniversalPacketSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketSenderCaller{contract: contract}, nil
}

// NewIbcUniversalPacketSenderTransactor creates a new write-only instance of IbcUniversalPacketSender, bound to a specific deployed contract.
func NewIbcUniversalPacketSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcUniversalPacketSenderTransactor, error) {
	contract, err := bindIbcUniversalPacketSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketSenderTransactor{contract: contract}, nil
}

// NewIbcUniversalPacketSenderFilterer creates a new log filterer instance of IbcUniversalPacketSender, bound to a specific deployed contract.
func NewIbcUniversalPacketSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcUniversalPacketSenderFilterer, error) {
	contract, err := bindIbcUniversalPacketSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcUniversalPacketSenderFilterer{contract: contract}, nil
}

// bindIbcUniversalPacketSender binds a generic wrapper to an already deployed contract.
func bindIbcUniversalPacketSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcUniversalPacketSenderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcUniversalPacketSender *IbcUniversalPacketSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcUniversalPacketSender.Contract.IbcUniversalPacketSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcUniversalPacketSender *IbcUniversalPacketSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcUniversalPacketSender.Contract.IbcUniversalPacketSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcUniversalPacketSender *IbcUniversalPacketSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcUniversalPacketSender.Contract.IbcUniversalPacketSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcUniversalPacketSender *IbcUniversalPacketSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcUniversalPacketSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcUniversalPacketSender *IbcUniversalPacketSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcUniversalPacketSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcUniversalPacketSender *IbcUniversalPacketSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcUniversalPacketSender.Contract.contract.Transact(opts, method, params...)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcUniversalPacketSender *IbcUniversalPacketSenderTransactor) SendUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcUniversalPacketSender.contract.Transact(opts, "sendUniversalPacket", channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcUniversalPacketSender *IbcUniversalPacketSenderSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcUniversalPacketSender.Contract.SendUniversalPacket(&_IbcUniversalPacketSender.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcUniversalPacketSender *IbcUniversalPacketSenderTransactorSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcUniversalPacketSender.Contract.SendUniversalPacket(&_IbcUniversalPacketSender.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IbcUniversalPacketSender *IbcUniversalPacketSenderTransactor) SendUniversalPacketWithFee(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IbcUniversalPacketSender.contract.Transact(opts, "sendUniversalPacketWithFee", channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IbcUniversalPacketSender *IbcUniversalPacketSenderSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IbcUniversalPacketSender.Contract.SendUniversalPacketWithFee(&_IbcUniversalPacketSender.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IbcUniversalPacketSender *IbcUniversalPacketSenderTransactorSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IbcUniversalPacketSender.Contract.SendUniversalPacketWithFee(&_IbcUniversalPacketSender.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}
