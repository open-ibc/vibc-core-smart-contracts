// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package protochannel

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

// ProtoChannelMetaData contains all meta data concerning the ProtoChannel contract.
var ProtoChannelMetaData = &bind.MetaData{
	ABI: "[]",
}

// ProtoChannelABI is the input ABI used to generate the binding from.
// Deprecated: Use ProtoChannelMetaData.ABI instead.
var ProtoChannelABI = ProtoChannelMetaData.ABI

// ProtoChannel is an auto generated Go binding around an Ethereum contract.
type ProtoChannel struct {
	ProtoChannelCaller     // Read-only binding to the contract
	ProtoChannelTransactor // Write-only binding to the contract
	ProtoChannelFilterer   // Log filterer for contract events
}

// ProtoChannelCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtoChannelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoChannelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtoChannelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoChannelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtoChannelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoChannelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtoChannelSession struct {
	Contract     *ProtoChannel     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtoChannelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtoChannelCallerSession struct {
	Contract *ProtoChannelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ProtoChannelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtoChannelTransactorSession struct {
	Contract     *ProtoChannelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ProtoChannelRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtoChannelRaw struct {
	Contract *ProtoChannel // Generic contract binding to access the raw methods on
}

// ProtoChannelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtoChannelCallerRaw struct {
	Contract *ProtoChannelCaller // Generic read-only contract binding to access the raw methods on
}

// ProtoChannelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtoChannelTransactorRaw struct {
	Contract *ProtoChannelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtoChannel creates a new instance of ProtoChannel, bound to a specific deployed contract.
func NewProtoChannel(address common.Address, backend bind.ContractBackend) (*ProtoChannel, error) {
	contract, err := bindProtoChannel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtoChannel{ProtoChannelCaller: ProtoChannelCaller{contract: contract}, ProtoChannelTransactor: ProtoChannelTransactor{contract: contract}, ProtoChannelFilterer: ProtoChannelFilterer{contract: contract}}, nil
}

// NewProtoChannelCaller creates a new read-only instance of ProtoChannel, bound to a specific deployed contract.
func NewProtoChannelCaller(address common.Address, caller bind.ContractCaller) (*ProtoChannelCaller, error) {
	contract, err := bindProtoChannel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtoChannelCaller{contract: contract}, nil
}

// NewProtoChannelTransactor creates a new write-only instance of ProtoChannel, bound to a specific deployed contract.
func NewProtoChannelTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtoChannelTransactor, error) {
	contract, err := bindProtoChannel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtoChannelTransactor{contract: contract}, nil
}

// NewProtoChannelFilterer creates a new log filterer instance of ProtoChannel, bound to a specific deployed contract.
func NewProtoChannelFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtoChannelFilterer, error) {
	contract, err := bindProtoChannel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtoChannelFilterer{contract: contract}, nil
}

// bindProtoChannel binds a generic wrapper to an already deployed contract.
func bindProtoChannel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProtoChannelMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtoChannel *ProtoChannelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtoChannel.Contract.ProtoChannelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtoChannel *ProtoChannelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtoChannel.Contract.ProtoChannelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtoChannel *ProtoChannelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtoChannel.Contract.ProtoChannelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtoChannel *ProtoChannelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtoChannel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtoChannel *ProtoChannelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtoChannel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtoChannel *ProtoChannelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtoChannel.Contract.contract.Transact(opts, method, params...)
}
