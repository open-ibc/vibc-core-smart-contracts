// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iproofverifier

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

// IAppStateVerifierMetaData contains all meta data concerning the IAppStateVerifier contract.
var IAppStateVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"InvalidAppHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIbcStateProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1BlockHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1BlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPacketProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedL1BlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedL1StateRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MethodNotImplemented\",\"inputs\":[]}]",
}

// IAppStateVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IAppStateVerifierMetaData.ABI instead.
var IAppStateVerifierABI = IAppStateVerifierMetaData.ABI

// IAppStateVerifier is an auto generated Go binding around an Ethereum contract.
type IAppStateVerifier struct {
	IAppStateVerifierCaller     // Read-only binding to the contract
	IAppStateVerifierTransactor // Write-only binding to the contract
	IAppStateVerifierFilterer   // Log filterer for contract events
}

// IAppStateVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAppStateVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppStateVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAppStateVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppStateVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAppStateVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAppStateVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAppStateVerifierSession struct {
	Contract     *IAppStateVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IAppStateVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAppStateVerifierCallerSession struct {
	Contract *IAppStateVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IAppStateVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAppStateVerifierTransactorSession struct {
	Contract     *IAppStateVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IAppStateVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAppStateVerifierRaw struct {
	Contract *IAppStateVerifier // Generic contract binding to access the raw methods on
}

// IAppStateVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAppStateVerifierCallerRaw struct {
	Contract *IAppStateVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IAppStateVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAppStateVerifierTransactorRaw struct {
	Contract *IAppStateVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAppStateVerifier creates a new instance of IAppStateVerifier, bound to a specific deployed contract.
func NewIAppStateVerifier(address common.Address, backend bind.ContractBackend) (*IAppStateVerifier, error) {
	contract, err := bindIAppStateVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAppStateVerifier{IAppStateVerifierCaller: IAppStateVerifierCaller{contract: contract}, IAppStateVerifierTransactor: IAppStateVerifierTransactor{contract: contract}, IAppStateVerifierFilterer: IAppStateVerifierFilterer{contract: contract}}, nil
}

// NewIAppStateVerifierCaller creates a new read-only instance of IAppStateVerifier, bound to a specific deployed contract.
func NewIAppStateVerifierCaller(address common.Address, caller bind.ContractCaller) (*IAppStateVerifierCaller, error) {
	contract, err := bindIAppStateVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAppStateVerifierCaller{contract: contract}, nil
}

// NewIAppStateVerifierTransactor creates a new write-only instance of IAppStateVerifier, bound to a specific deployed contract.
func NewIAppStateVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IAppStateVerifierTransactor, error) {
	contract, err := bindIAppStateVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAppStateVerifierTransactor{contract: contract}, nil
}

// NewIAppStateVerifierFilterer creates a new log filterer instance of IAppStateVerifier, bound to a specific deployed contract.
func NewIAppStateVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IAppStateVerifierFilterer, error) {
	contract, err := bindIAppStateVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAppStateVerifierFilterer{contract: contract}, nil
}

// bindIAppStateVerifier binds a generic wrapper to an already deployed contract.
func bindIAppStateVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAppStateVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAppStateVerifier *IAppStateVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAppStateVerifier.Contract.IAppStateVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAppStateVerifier *IAppStateVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAppStateVerifier.Contract.IAppStateVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAppStateVerifier *IAppStateVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAppStateVerifier.Contract.IAppStateVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAppStateVerifier *IAppStateVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAppStateVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAppStateVerifier *IAppStateVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAppStateVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAppStateVerifier *IAppStateVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAppStateVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) pure returns()
func (_IAppStateVerifier *IAppStateVerifierCaller) VerifyMembership(opts *bind.CallOpts, appHash [32]byte, key []byte, value []byte, proof Ics23Proof) error {
	var out []interface{}
	err := _IAppStateVerifier.contract.Call(opts, &out, "verifyMembership", appHash, key, value, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) pure returns()
func (_IAppStateVerifier *IAppStateVerifierSession) VerifyMembership(appHash [32]byte, key []byte, value []byte, proof Ics23Proof) error {
	return _IAppStateVerifier.Contract.VerifyMembership(&_IAppStateVerifier.CallOpts, appHash, key, value, proof)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) pure returns()
func (_IAppStateVerifier *IAppStateVerifierCallerSession) VerifyMembership(appHash [32]byte, key []byte, value []byte, proof Ics23Proof) error {
	return _IAppStateVerifier.Contract.VerifyMembership(&_IAppStateVerifier.CallOpts, appHash, key, value, proof)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 appHash, bytes key, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) pure returns()
func (_IAppStateVerifier *IAppStateVerifierCaller) VerifyNonMembership(opts *bind.CallOpts, appHash [32]byte, key []byte, proof Ics23Proof) error {
	var out []interface{}
	err := _IAppStateVerifier.contract.Call(opts, &out, "verifyNonMembership", appHash, key, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 appHash, bytes key, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) pure returns()
func (_IAppStateVerifier *IAppStateVerifierSession) VerifyNonMembership(appHash [32]byte, key []byte, proof Ics23Proof) error {
	return _IAppStateVerifier.Contract.VerifyNonMembership(&_IAppStateVerifier.CallOpts, appHash, key, proof)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 appHash, bytes key, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) pure returns()
func (_IAppStateVerifier *IAppStateVerifierCallerSession) VerifyNonMembership(appHash [32]byte, key []byte, proof Ics23Proof) error {
	return _IAppStateVerifier.Contract.VerifyNonMembership(&_IAppStateVerifier.CallOpts, appHash, key, proof)
}
