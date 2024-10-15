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

// INonMembershipVerifierMetaData contains all meta data concerning the INonMembershipVerifier contract.
var INonMembershipVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// INonMembershipVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use INonMembershipVerifierMetaData.ABI instead.
var INonMembershipVerifierABI = INonMembershipVerifierMetaData.ABI

// INonMembershipVerifier is an auto generated Go binding around an Ethereum contract.
type INonMembershipVerifier struct {
	INonMembershipVerifierCaller     // Read-only binding to the contract
	INonMembershipVerifierTransactor // Write-only binding to the contract
	INonMembershipVerifierFilterer   // Log filterer for contract events
}

// INonMembershipVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type INonMembershipVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INonMembershipVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INonMembershipVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INonMembershipVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INonMembershipVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INonMembershipVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INonMembershipVerifierSession struct {
	Contract     *INonMembershipVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// INonMembershipVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INonMembershipVerifierCallerSession struct {
	Contract *INonMembershipVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// INonMembershipVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INonMembershipVerifierTransactorSession struct {
	Contract     *INonMembershipVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// INonMembershipVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type INonMembershipVerifierRaw struct {
	Contract *INonMembershipVerifier // Generic contract binding to access the raw methods on
}

// INonMembershipVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INonMembershipVerifierCallerRaw struct {
	Contract *INonMembershipVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// INonMembershipVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INonMembershipVerifierTransactorRaw struct {
	Contract *INonMembershipVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINonMembershipVerifier creates a new instance of INonMembershipVerifier, bound to a specific deployed contract.
func NewINonMembershipVerifier(address common.Address, backend bind.ContractBackend) (*INonMembershipVerifier, error) {
	contract, err := bindINonMembershipVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INonMembershipVerifier{INonMembershipVerifierCaller: INonMembershipVerifierCaller{contract: contract}, INonMembershipVerifierTransactor: INonMembershipVerifierTransactor{contract: contract}, INonMembershipVerifierFilterer: INonMembershipVerifierFilterer{contract: contract}}, nil
}

// NewINonMembershipVerifierCaller creates a new read-only instance of INonMembershipVerifier, bound to a specific deployed contract.
func NewINonMembershipVerifierCaller(address common.Address, caller bind.ContractCaller) (*INonMembershipVerifierCaller, error) {
	contract, err := bindINonMembershipVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INonMembershipVerifierCaller{contract: contract}, nil
}

// NewINonMembershipVerifierTransactor creates a new write-only instance of INonMembershipVerifier, bound to a specific deployed contract.
func NewINonMembershipVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*INonMembershipVerifierTransactor, error) {
	contract, err := bindINonMembershipVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INonMembershipVerifierTransactor{contract: contract}, nil
}

// NewINonMembershipVerifierFilterer creates a new log filterer instance of INonMembershipVerifier, bound to a specific deployed contract.
func NewINonMembershipVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*INonMembershipVerifierFilterer, error) {
	contract, err := bindINonMembershipVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INonMembershipVerifierFilterer{contract: contract}, nil
}

// bindINonMembershipVerifier binds a generic wrapper to an already deployed contract.
func bindINonMembershipVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INonMembershipVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INonMembershipVerifier *INonMembershipVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INonMembershipVerifier.Contract.INonMembershipVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INonMembershipVerifier *INonMembershipVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INonMembershipVerifier.Contract.INonMembershipVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INonMembershipVerifier *INonMembershipVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INonMembershipVerifier.Contract.INonMembershipVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INonMembershipVerifier *INonMembershipVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INonMembershipVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INonMembershipVerifier *INonMembershipVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INonMembershipVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INonMembershipVerifier *INonMembershipVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INonMembershipVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyNonMembership is a paid mutator transaction binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key) returns()
func (_INonMembershipVerifier *INonMembershipVerifierTransactor) VerifyNonMembership(opts *bind.TransactOpts, proof Ics23Proof, key []byte) (*types.Transaction, error) {
	return _INonMembershipVerifier.contract.Transact(opts, "verifyNonMembership", proof, key)
}

// VerifyNonMembership is a paid mutator transaction binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key) returns()
func (_INonMembershipVerifier *INonMembershipVerifierSession) VerifyNonMembership(proof Ics23Proof, key []byte) (*types.Transaction, error) {
	return _INonMembershipVerifier.Contract.VerifyNonMembership(&_INonMembershipVerifier.TransactOpts, proof, key)
}

// VerifyNonMembership is a paid mutator transaction binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key) returns()
func (_INonMembershipVerifier *INonMembershipVerifierTransactorSession) VerifyNonMembership(proof Ics23Proof, key []byte) (*types.Transaction, error) {
	return _INonMembershipVerifier.Contract.VerifyNonMembership(&_INonMembershipVerifier.TransactOpts, proof, key)
}
