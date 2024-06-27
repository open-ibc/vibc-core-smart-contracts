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

// IbcMiddlwareProviderMetaData contains all meta data concerning the IbcMiddlwareProvider contract.
var IbcMiddlwareProviderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MW_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"MWID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacketWithFee\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"}]",
}

// IbcMiddlwareProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcMiddlwareProviderMetaData.ABI instead.
var IbcMiddlwareProviderABI = IbcMiddlwareProviderMetaData.ABI

// IbcMiddlwareProvider is an auto generated Go binding around an Ethereum contract.
type IbcMiddlwareProvider struct {
	IbcMiddlwareProviderCaller     // Read-only binding to the contract
	IbcMiddlwareProviderTransactor // Write-only binding to the contract
	IbcMiddlwareProviderFilterer   // Log filterer for contract events
}

// IbcMiddlwareProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcMiddlwareProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMiddlwareProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcMiddlwareProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMiddlwareProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcMiddlwareProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcMiddlwareProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcMiddlwareProviderSession struct {
	Contract     *IbcMiddlwareProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IbcMiddlwareProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcMiddlwareProviderCallerSession struct {
	Contract *IbcMiddlwareProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IbcMiddlwareProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcMiddlwareProviderTransactorSession struct {
	Contract     *IbcMiddlwareProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IbcMiddlwareProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcMiddlwareProviderRaw struct {
	Contract *IbcMiddlwareProvider // Generic contract binding to access the raw methods on
}

// IbcMiddlwareProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcMiddlwareProviderCallerRaw struct {
	Contract *IbcMiddlwareProviderCaller // Generic read-only contract binding to access the raw methods on
}

// IbcMiddlwareProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcMiddlwareProviderTransactorRaw struct {
	Contract *IbcMiddlwareProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcMiddlwareProvider creates a new instance of IbcMiddlwareProvider, bound to a specific deployed contract.
func NewIbcMiddlwareProvider(address common.Address, backend bind.ContractBackend) (*IbcMiddlwareProvider, error) {
	contract, err := bindIbcMiddlwareProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcMiddlwareProvider{IbcMiddlwareProviderCaller: IbcMiddlwareProviderCaller{contract: contract}, IbcMiddlwareProviderTransactor: IbcMiddlwareProviderTransactor{contract: contract}, IbcMiddlwareProviderFilterer: IbcMiddlwareProviderFilterer{contract: contract}}, nil
}

// NewIbcMiddlwareProviderCaller creates a new read-only instance of IbcMiddlwareProvider, bound to a specific deployed contract.
func NewIbcMiddlwareProviderCaller(address common.Address, caller bind.ContractCaller) (*IbcMiddlwareProviderCaller, error) {
	contract, err := bindIbcMiddlwareProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcMiddlwareProviderCaller{contract: contract}, nil
}

// NewIbcMiddlwareProviderTransactor creates a new write-only instance of IbcMiddlwareProvider, bound to a specific deployed contract.
func NewIbcMiddlwareProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcMiddlwareProviderTransactor, error) {
	contract, err := bindIbcMiddlwareProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcMiddlwareProviderTransactor{contract: contract}, nil
}

// NewIbcMiddlwareProviderFilterer creates a new log filterer instance of IbcMiddlwareProvider, bound to a specific deployed contract.
func NewIbcMiddlwareProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcMiddlwareProviderFilterer, error) {
	contract, err := bindIbcMiddlwareProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcMiddlwareProviderFilterer{contract: contract}, nil
}

// bindIbcMiddlwareProvider binds a generic wrapper to an already deployed contract.
func bindIbcMiddlwareProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcMiddlwareProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcMiddlwareProvider *IbcMiddlwareProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcMiddlwareProvider.Contract.IbcMiddlwareProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcMiddlwareProvider *IbcMiddlwareProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcMiddlwareProvider.Contract.IbcMiddlwareProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcMiddlwareProvider *IbcMiddlwareProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcMiddlwareProvider.Contract.IbcMiddlwareProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcMiddlwareProvider *IbcMiddlwareProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcMiddlwareProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcMiddlwareProvider *IbcMiddlwareProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcMiddlwareProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcMiddlwareProvider *IbcMiddlwareProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcMiddlwareProvider.Contract.contract.Transact(opts, method, params...)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256 MWID)
func (_IbcMiddlwareProvider *IbcMiddlwareProviderCaller) MWID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IbcMiddlwareProvider.contract.Call(opts, &out, "MW_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256 MWID)
func (_IbcMiddlwareProvider *IbcMiddlwareProviderSession) MWID() (*big.Int, error) {
	return _IbcMiddlwareProvider.Contract.MWID(&_IbcMiddlwareProvider.CallOpts)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256 MWID)
func (_IbcMiddlwareProvider *IbcMiddlwareProviderCallerSession) MWID() (*big.Int, error) {
	return _IbcMiddlwareProvider.Contract.MWID(&_IbcMiddlwareProvider.CallOpts)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcMiddlwareProvider *IbcMiddlwareProviderTransactor) SendUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcMiddlwareProvider.contract.Transact(opts, "sendUniversalPacket", channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcMiddlwareProvider *IbcMiddlwareProviderSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcMiddlwareProvider.Contract.SendUniversalPacket(&_IbcMiddlwareProvider.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcMiddlwareProvider *IbcMiddlwareProviderTransactorSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcMiddlwareProvider.Contract.SendUniversalPacket(&_IbcMiddlwareProvider.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IbcMiddlwareProvider *IbcMiddlwareProviderTransactor) SendUniversalPacketWithFee(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IbcMiddlwareProvider.contract.Transact(opts, "sendUniversalPacketWithFee", channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IbcMiddlwareProvider *IbcMiddlwareProviderSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IbcMiddlwareProvider.Contract.SendUniversalPacketWithFee(&_IbcMiddlwareProvider.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_IbcMiddlwareProvider *IbcMiddlwareProviderTransactorSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _IbcMiddlwareProvider.Contract.SendUniversalPacketWithFee(&_IbcMiddlwareProvider.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}
