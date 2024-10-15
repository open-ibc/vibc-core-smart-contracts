// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ilightclient

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

// Ics23Proof is an auto generated low-level Go binding around an user-defined struct.
type Ics23Proof struct {
	Proof  []OpIcs23Proof
	Height *big.Int
}

// OpIcs23Proof is an auto generated low-level Go binding around an user-defined struct.
type OpIcs23Proof struct {
	Path   []OpIcs23ProofPath
	Key    []byte
	Value  []byte
	Prefix []byte
}

// OpIcs23ProofPath is an auto generated low-level Go binding around an user-defined struct.
type OpIcs23ProofPath struct {
	Prefix []byte
	Suffix []byte
}

// IMembershipVerifierMetaData contains all meta data concerning the IMembershipVerifier contract.
var IMembershipVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expectedValue\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// IMembershipVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IMembershipVerifierMetaData.ABI instead.
var IMembershipVerifierABI = IMembershipVerifierMetaData.ABI

// IMembershipVerifier is an auto generated Go binding around an Ethereum contract.
type IMembershipVerifier struct {
	IMembershipVerifierCaller     // Read-only binding to the contract
	IMembershipVerifierTransactor // Write-only binding to the contract
	IMembershipVerifierFilterer   // Log filterer for contract events
}

// IMembershipVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMembershipVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMembershipVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMembershipVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMembershipVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMembershipVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMembershipVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMembershipVerifierSession struct {
	Contract     *IMembershipVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMembershipVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMembershipVerifierCallerSession struct {
	Contract *IMembershipVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMembershipVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMembershipVerifierTransactorSession struct {
	Contract     *IMembershipVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMembershipVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMembershipVerifierRaw struct {
	Contract *IMembershipVerifier // Generic contract binding to access the raw methods on
}

// IMembershipVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMembershipVerifierCallerRaw struct {
	Contract *IMembershipVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IMembershipVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMembershipVerifierTransactorRaw struct {
	Contract *IMembershipVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMembershipVerifier creates a new instance of IMembershipVerifier, bound to a specific deployed contract.
func NewIMembershipVerifier(address common.Address, backend bind.ContractBackend) (*IMembershipVerifier, error) {
	contract, err := bindIMembershipVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMembershipVerifier{IMembershipVerifierCaller: IMembershipVerifierCaller{contract: contract}, IMembershipVerifierTransactor: IMembershipVerifierTransactor{contract: contract}, IMembershipVerifierFilterer: IMembershipVerifierFilterer{contract: contract}}, nil
}

// NewIMembershipVerifierCaller creates a new read-only instance of IMembershipVerifier, bound to a specific deployed contract.
func NewIMembershipVerifierCaller(address common.Address, caller bind.ContractCaller) (*IMembershipVerifierCaller, error) {
	contract, err := bindIMembershipVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMembershipVerifierCaller{contract: contract}, nil
}

// NewIMembershipVerifierTransactor creates a new write-only instance of IMembershipVerifier, bound to a specific deployed contract.
func NewIMembershipVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IMembershipVerifierTransactor, error) {
	contract, err := bindIMembershipVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMembershipVerifierTransactor{contract: contract}, nil
}

// NewIMembershipVerifierFilterer creates a new log filterer instance of IMembershipVerifier, bound to a specific deployed contract.
func NewIMembershipVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IMembershipVerifierFilterer, error) {
	contract, err := bindIMembershipVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMembershipVerifierFilterer{contract: contract}, nil
}

// bindIMembershipVerifier binds a generic wrapper to an already deployed contract.
func bindIMembershipVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMembershipVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMembershipVerifier *IMembershipVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMembershipVerifier.Contract.IMembershipVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMembershipVerifier *IMembershipVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMembershipVerifier.Contract.IMembershipVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMembershipVerifier *IMembershipVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMembershipVerifier.Contract.IMembershipVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMembershipVerifier *IMembershipVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMembershipVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMembershipVerifier *IMembershipVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMembershipVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMembershipVerifier *IMembershipVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMembershipVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyMembership is a paid mutator transaction binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key, bytes expectedValue) returns()
func (_IMembershipVerifier *IMembershipVerifierTransactor) VerifyMembership(opts *bind.TransactOpts, proof Ics23Proof, key []byte, expectedValue []byte) (*types.Transaction, error) {
	return _IMembershipVerifier.contract.Transact(opts, "verifyMembership", proof, key, expectedValue)
}

// VerifyMembership is a paid mutator transaction binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key, bytes expectedValue) returns()
func (_IMembershipVerifier *IMembershipVerifierSession) VerifyMembership(proof Ics23Proof, key []byte, expectedValue []byte) (*types.Transaction, error) {
	return _IMembershipVerifier.Contract.VerifyMembership(&_IMembershipVerifier.TransactOpts, proof, key, expectedValue)
}

// VerifyMembership is a paid mutator transaction binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key, bytes expectedValue) returns()
func (_IMembershipVerifier *IMembershipVerifierTransactorSession) VerifyMembership(proof Ics23Proof, key []byte, expectedValue []byte) (*types.Transaction, error) {
	return _IMembershipVerifier.Contract.VerifyMembership(&_IMembershipVerifier.TransactOpts, proof, key, expectedValue)
}
