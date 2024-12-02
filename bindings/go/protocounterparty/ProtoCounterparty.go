// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package protocounterparty

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

// ProtoCounterpartyMetaData contains all meta data concerning the ProtoCounterparty contract.
var ProtoCounterpartyMetaData = &bind.MetaData{
	ABI: "[]",
}

// ProtoCounterpartyABI is the input ABI used to generate the binding from.
// Deprecated: Use ProtoCounterpartyMetaData.ABI instead.
var ProtoCounterpartyABI = ProtoCounterpartyMetaData.ABI

// ProtoCounterparty is an auto generated Go binding around an Ethereum contract.
type ProtoCounterparty struct {
	ProtoCounterpartyCaller     // Read-only binding to the contract
	ProtoCounterpartyTransactor // Write-only binding to the contract
	ProtoCounterpartyFilterer   // Log filterer for contract events
}

// ProtoCounterpartyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtoCounterpartyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoCounterpartyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtoCounterpartyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoCounterpartyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtoCounterpartyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtoCounterpartySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtoCounterpartySession struct {
	Contract     *ProtoCounterparty // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProtoCounterpartyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtoCounterpartyCallerSession struct {
	Contract *ProtoCounterpartyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ProtoCounterpartyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtoCounterpartyTransactorSession struct {
	Contract     *ProtoCounterpartyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ProtoCounterpartyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtoCounterpartyRaw struct {
	Contract *ProtoCounterparty // Generic contract binding to access the raw methods on
}

// ProtoCounterpartyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtoCounterpartyCallerRaw struct {
	Contract *ProtoCounterpartyCaller // Generic read-only contract binding to access the raw methods on
}

// ProtoCounterpartyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtoCounterpartyTransactorRaw struct {
	Contract *ProtoCounterpartyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtoCounterparty creates a new instance of ProtoCounterparty, bound to a specific deployed contract.
func NewProtoCounterparty(address common.Address, backend bind.ContractBackend) (*ProtoCounterparty, error) {
	contract, err := bindProtoCounterparty(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtoCounterparty{ProtoCounterpartyCaller: ProtoCounterpartyCaller{contract: contract}, ProtoCounterpartyTransactor: ProtoCounterpartyTransactor{contract: contract}, ProtoCounterpartyFilterer: ProtoCounterpartyFilterer{contract: contract}}, nil
}

// NewProtoCounterpartyCaller creates a new read-only instance of ProtoCounterparty, bound to a specific deployed contract.
func NewProtoCounterpartyCaller(address common.Address, caller bind.ContractCaller) (*ProtoCounterpartyCaller, error) {
	contract, err := bindProtoCounterparty(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtoCounterpartyCaller{contract: contract}, nil
}

// NewProtoCounterpartyTransactor creates a new write-only instance of ProtoCounterparty, bound to a specific deployed contract.
func NewProtoCounterpartyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtoCounterpartyTransactor, error) {
	contract, err := bindProtoCounterparty(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtoCounterpartyTransactor{contract: contract}, nil
}

// NewProtoCounterpartyFilterer creates a new log filterer instance of ProtoCounterparty, bound to a specific deployed contract.
func NewProtoCounterpartyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtoCounterpartyFilterer, error) {
	contract, err := bindProtoCounterparty(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtoCounterpartyFilterer{contract: contract}, nil
}

// bindProtoCounterparty binds a generic wrapper to an already deployed contract.
func bindProtoCounterparty(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProtoCounterpartyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtoCounterparty *ProtoCounterpartyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtoCounterparty.Contract.ProtoCounterpartyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtoCounterparty *ProtoCounterpartyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtoCounterparty.Contract.ProtoCounterpartyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtoCounterparty *ProtoCounterpartyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtoCounterparty.Contract.ProtoCounterpartyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtoCounterparty *ProtoCounterpartyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtoCounterparty.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtoCounterparty *ProtoCounterpartyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtoCounterparty.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtoCounterparty *ProtoCounterpartyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtoCounterparty.Contract.contract.Transact(opts, method, params...)
}
