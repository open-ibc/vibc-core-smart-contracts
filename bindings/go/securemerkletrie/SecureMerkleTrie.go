// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package securemerkletrie

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

// SecureMerkleTrieMetaData contains all meta data concerning the SecureMerkleTrie contract.
var SecureMerkleTrieMetaData = &bind.MetaData{
	ABI: "[]",
}

// SecureMerkleTrieABI is the input ABI used to generate the binding from.
// Deprecated: Use SecureMerkleTrieMetaData.ABI instead.
var SecureMerkleTrieABI = SecureMerkleTrieMetaData.ABI

// SecureMerkleTrie is an auto generated Go binding around an Ethereum contract.
type SecureMerkleTrie struct {
	SecureMerkleTrieCaller     // Read-only binding to the contract
	SecureMerkleTrieTransactor // Write-only binding to the contract
	SecureMerkleTrieFilterer   // Log filterer for contract events
}

// SecureMerkleTrieCaller is an auto generated read-only Go binding around an Ethereum contract.
type SecureMerkleTrieCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecureMerkleTrieTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SecureMerkleTrieTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecureMerkleTrieFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SecureMerkleTrieFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SecureMerkleTrieSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SecureMerkleTrieSession struct {
	Contract     *SecureMerkleTrie // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SecureMerkleTrieCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SecureMerkleTrieCallerSession struct {
	Contract *SecureMerkleTrieCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SecureMerkleTrieTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SecureMerkleTrieTransactorSession struct {
	Contract     *SecureMerkleTrieTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SecureMerkleTrieRaw is an auto generated low-level Go binding around an Ethereum contract.
type SecureMerkleTrieRaw struct {
	Contract *SecureMerkleTrie // Generic contract binding to access the raw methods on
}

// SecureMerkleTrieCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SecureMerkleTrieCallerRaw struct {
	Contract *SecureMerkleTrieCaller // Generic read-only contract binding to access the raw methods on
}

// SecureMerkleTrieTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SecureMerkleTrieTransactorRaw struct {
	Contract *SecureMerkleTrieTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSecureMerkleTrie creates a new instance of SecureMerkleTrie, bound to a specific deployed contract.
func NewSecureMerkleTrie(address common.Address, backend bind.ContractBackend) (*SecureMerkleTrie, error) {
	contract, err := bindSecureMerkleTrie(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SecureMerkleTrie{SecureMerkleTrieCaller: SecureMerkleTrieCaller{contract: contract}, SecureMerkleTrieTransactor: SecureMerkleTrieTransactor{contract: contract}, SecureMerkleTrieFilterer: SecureMerkleTrieFilterer{contract: contract}}, nil
}

// NewSecureMerkleTrieCaller creates a new read-only instance of SecureMerkleTrie, bound to a specific deployed contract.
func NewSecureMerkleTrieCaller(address common.Address, caller bind.ContractCaller) (*SecureMerkleTrieCaller, error) {
	contract, err := bindSecureMerkleTrie(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SecureMerkleTrieCaller{contract: contract}, nil
}

// NewSecureMerkleTrieTransactor creates a new write-only instance of SecureMerkleTrie, bound to a specific deployed contract.
func NewSecureMerkleTrieTransactor(address common.Address, transactor bind.ContractTransactor) (*SecureMerkleTrieTransactor, error) {
	contract, err := bindSecureMerkleTrie(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SecureMerkleTrieTransactor{contract: contract}, nil
}

// NewSecureMerkleTrieFilterer creates a new log filterer instance of SecureMerkleTrie, bound to a specific deployed contract.
func NewSecureMerkleTrieFilterer(address common.Address, filterer bind.ContractFilterer) (*SecureMerkleTrieFilterer, error) {
	contract, err := bindSecureMerkleTrie(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SecureMerkleTrieFilterer{contract: contract}, nil
}

// bindSecureMerkleTrie binds a generic wrapper to an already deployed contract.
func bindSecureMerkleTrie(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SecureMerkleTrieMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecureMerkleTrie *SecureMerkleTrieRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SecureMerkleTrie.Contract.SecureMerkleTrieCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecureMerkleTrie *SecureMerkleTrieRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecureMerkleTrie.Contract.SecureMerkleTrieTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecureMerkleTrie *SecureMerkleTrieRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecureMerkleTrie.Contract.SecureMerkleTrieTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SecureMerkleTrie *SecureMerkleTrieCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SecureMerkleTrie.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SecureMerkleTrie *SecureMerkleTrieTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SecureMerkleTrie.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SecureMerkleTrie *SecureMerkleTrieTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SecureMerkleTrie.Contract.contract.Transact(opts, method, params...)
}
