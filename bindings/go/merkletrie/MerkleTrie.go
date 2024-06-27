// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package merkletrie

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

// MerkleTrieMetaData contains all meta data concerning the MerkleTrie contract.
var MerkleTrieMetaData = &bind.MetaData{
	ABI: "[]",
}

// MerkleTrieABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleTrieMetaData.ABI instead.
var MerkleTrieABI = MerkleTrieMetaData.ABI

// MerkleTrie is an auto generated Go binding around an Ethereum contract.
type MerkleTrie struct {
	MerkleTrieCaller     // Read-only binding to the contract
	MerkleTrieTransactor // Write-only binding to the contract
	MerkleTrieFilterer   // Log filterer for contract events
}

// MerkleTrieCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleTrieCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTrieTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTrieTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTrieFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleTrieFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTrieSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleTrieSession struct {
	Contract     *MerkleTrie       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleTrieCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleTrieCallerSession struct {
	Contract *MerkleTrieCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleTrieTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTrieTransactorSession struct {
	Contract     *MerkleTrieTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleTrieRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleTrieRaw struct {
	Contract *MerkleTrie // Generic contract binding to access the raw methods on
}

// MerkleTrieCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleTrieCallerRaw struct {
	Contract *MerkleTrieCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTrieTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTrieTransactorRaw struct {
	Contract *MerkleTrieTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleTrie creates a new instance of MerkleTrie, bound to a specific deployed contract.
func NewMerkleTrie(address common.Address, backend bind.ContractBackend) (*MerkleTrie, error) {
	contract, err := bindMerkleTrie(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleTrie{MerkleTrieCaller: MerkleTrieCaller{contract: contract}, MerkleTrieTransactor: MerkleTrieTransactor{contract: contract}, MerkleTrieFilterer: MerkleTrieFilterer{contract: contract}}, nil
}

// NewMerkleTrieCaller creates a new read-only instance of MerkleTrie, bound to a specific deployed contract.
func NewMerkleTrieCaller(address common.Address, caller bind.ContractCaller) (*MerkleTrieCaller, error) {
	contract, err := bindMerkleTrie(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTrieCaller{contract: contract}, nil
}

// NewMerkleTrieTransactor creates a new write-only instance of MerkleTrie, bound to a specific deployed contract.
func NewMerkleTrieTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTrieTransactor, error) {
	contract, err := bindMerkleTrie(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTrieTransactor{contract: contract}, nil
}

// NewMerkleTrieFilterer creates a new log filterer instance of MerkleTrie, bound to a specific deployed contract.
func NewMerkleTrieFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleTrieFilterer, error) {
	contract, err := bindMerkleTrie(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleTrieFilterer{contract: contract}, nil
}

// bindMerkleTrie binds a generic wrapper to an already deployed contract.
func bindMerkleTrie(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerkleTrieMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTrie *MerkleTrieRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTrie.Contract.MerkleTrieCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTrie *MerkleTrieRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTrie.Contract.MerkleTrieTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTrie *MerkleTrieRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTrie.Contract.MerkleTrieTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTrie *MerkleTrieCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTrie.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTrie *MerkleTrieTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTrie.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTrie *MerkleTrieTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTrie.Contract.contract.Transact(opts, method, params...)
}
