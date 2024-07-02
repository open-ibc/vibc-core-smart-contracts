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

// L1Header is an auto generated low-level Go binding around an user-defined struct.
type L1Header struct {
	Header    [][]byte
	StateRoot [32]byte
	Number    uint64
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

// OpL2StateProof is an auto generated low-level Go binding around an user-defined struct.
type OpL2StateProof struct {
	AccountProof        [][]byte
	OutputRootProof     [][]byte
	L2OutputProposalKey [32]byte
	L2BlockHash         [32]byte
}

// IProofVerifierMetaData contains all meta data concerning the IProofVerifier contract.
var IProofVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyStateUpdate\",\"inputs\":[{\"name\":\"l1header\",\"type\":\"tuple\",\"internalType\":\"structL1Header\",\"components\":[{\"name\":\"header\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"number\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structOpL2StateProof\",\"components\":[{\"name\":\"accountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"outputRootProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"l2OutputProposalKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"trustedL1BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"trustedL1BlockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidAppHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIbcStateProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1BlockHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1BlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPacketProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedL1BlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedL1StateRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MethodNotImplemented\",\"inputs\":[]}]",
}

// IProofVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IProofVerifierMetaData.ABI instead.
var IProofVerifierABI = IProofVerifierMetaData.ABI

// IProofVerifier is an auto generated Go binding around an Ethereum contract.
type IProofVerifier struct {
	IProofVerifierCaller     // Read-only binding to the contract
	IProofVerifierTransactor // Write-only binding to the contract
	IProofVerifierFilterer   // Log filterer for contract events
}

// IProofVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IProofVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProofVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IProofVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProofVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IProofVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IProofVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IProofVerifierSession struct {
	Contract     *IProofVerifier   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IProofVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IProofVerifierCallerSession struct {
	Contract *IProofVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IProofVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IProofVerifierTransactorSession struct {
	Contract     *IProofVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IProofVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IProofVerifierRaw struct {
	Contract *IProofVerifier // Generic contract binding to access the raw methods on
}

// IProofVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IProofVerifierCallerRaw struct {
	Contract *IProofVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IProofVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IProofVerifierTransactorRaw struct {
	Contract *IProofVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIProofVerifier creates a new instance of IProofVerifier, bound to a specific deployed contract.
func NewIProofVerifier(address common.Address, backend bind.ContractBackend) (*IProofVerifier, error) {
	contract, err := bindIProofVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IProofVerifier{IProofVerifierCaller: IProofVerifierCaller{contract: contract}, IProofVerifierTransactor: IProofVerifierTransactor{contract: contract}, IProofVerifierFilterer: IProofVerifierFilterer{contract: contract}}, nil
}

// NewIProofVerifierCaller creates a new read-only instance of IProofVerifier, bound to a specific deployed contract.
func NewIProofVerifierCaller(address common.Address, caller bind.ContractCaller) (*IProofVerifierCaller, error) {
	contract, err := bindIProofVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IProofVerifierCaller{contract: contract}, nil
}

// NewIProofVerifierTransactor creates a new write-only instance of IProofVerifier, bound to a specific deployed contract.
func NewIProofVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IProofVerifierTransactor, error) {
	contract, err := bindIProofVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IProofVerifierTransactor{contract: contract}, nil
}

// NewIProofVerifierFilterer creates a new log filterer instance of IProofVerifier, bound to a specific deployed contract.
func NewIProofVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IProofVerifierFilterer, error) {
	contract, err := bindIProofVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IProofVerifierFilterer{contract: contract}, nil
}

// bindIProofVerifier binds a generic wrapper to an already deployed contract.
func bindIProofVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IProofVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProofVerifier *IProofVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProofVerifier.Contract.IProofVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProofVerifier *IProofVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProofVerifier.Contract.IProofVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProofVerifier *IProofVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProofVerifier.Contract.IProofVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IProofVerifier *IProofVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IProofVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IProofVerifier *IProofVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IProofVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IProofVerifier *IProofVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IProofVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) pure returns()
func (_IProofVerifier *IProofVerifierCaller) VerifyMembership(opts *bind.CallOpts, appHash [32]byte, key []byte, value []byte, proof Ics23Proof) error {
	var out []interface{}
	err := _IProofVerifier.contract.Call(opts, &out, "verifyMembership", appHash, key, value, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) pure returns()
func (_IProofVerifier *IProofVerifierSession) VerifyMembership(appHash [32]byte, key []byte, value []byte, proof Ics23Proof) error {
	return _IProofVerifier.Contract.VerifyMembership(&_IProofVerifier.CallOpts, appHash, key, value, proof)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) pure returns()
func (_IProofVerifier *IProofVerifierCallerSession) VerifyMembership(appHash [32]byte, key []byte, value []byte, proof Ics23Proof) error {
	return _IProofVerifier.Contract.VerifyMembership(&_IProofVerifier.CallOpts, appHash, key, value, proof)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 appHash, bytes key, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) pure returns()
func (_IProofVerifier *IProofVerifierCaller) VerifyNonMembership(opts *bind.CallOpts, appHash [32]byte, key []byte, proof Ics23Proof) error {
	var out []interface{}
	err := _IProofVerifier.contract.Call(opts, &out, "verifyNonMembership", appHash, key, proof)

	if err != nil {
		return err
	}

	return err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 appHash, bytes key, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) pure returns()
func (_IProofVerifier *IProofVerifierSession) VerifyNonMembership(appHash [32]byte, key []byte, proof Ics23Proof) error {
	return _IProofVerifier.Contract.VerifyNonMembership(&_IProofVerifier.CallOpts, appHash, key, proof)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 appHash, bytes key, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) pure returns()
func (_IProofVerifier *IProofVerifierCallerSession) VerifyNonMembership(appHash [32]byte, key []byte, proof Ics23Proof) error {
	return _IProofVerifier.Contract.VerifyNonMembership(&_IProofVerifier.CallOpts, appHash, key, proof)
}

// VerifyStateUpdate is a free data retrieval call binding the contract method 0x0a1bb8b5.
//
// Solidity: function verifyStateUpdate((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, bytes32 appHash, bytes32 trustedL1BlockHash, uint64 trustedL1BlockNumber) view returns()
func (_IProofVerifier *IProofVerifierCaller) VerifyStateUpdate(opts *bind.CallOpts, l1header L1Header, proof OpL2StateProof, appHash [32]byte, trustedL1BlockHash [32]byte, trustedL1BlockNumber uint64) error {
	var out []interface{}
	err := _IProofVerifier.contract.Call(opts, &out, "verifyStateUpdate", l1header, proof, appHash, trustedL1BlockHash, trustedL1BlockNumber)

	if err != nil {
		return err
	}

	return err

}

// VerifyStateUpdate is a free data retrieval call binding the contract method 0x0a1bb8b5.
//
// Solidity: function verifyStateUpdate((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, bytes32 appHash, bytes32 trustedL1BlockHash, uint64 trustedL1BlockNumber) view returns()
func (_IProofVerifier *IProofVerifierSession) VerifyStateUpdate(l1header L1Header, proof OpL2StateProof, appHash [32]byte, trustedL1BlockHash [32]byte, trustedL1BlockNumber uint64) error {
	return _IProofVerifier.Contract.VerifyStateUpdate(&_IProofVerifier.CallOpts, l1header, proof, appHash, trustedL1BlockHash, trustedL1BlockNumber)
}

// VerifyStateUpdate is a free data retrieval call binding the contract method 0x0a1bb8b5.
//
// Solidity: function verifyStateUpdate((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, bytes32 appHash, bytes32 trustedL1BlockHash, uint64 trustedL1BlockNumber) view returns()
func (_IProofVerifier *IProofVerifierCallerSession) VerifyStateUpdate(l1header L1Header, proof OpL2StateProof, appHash [32]byte, trustedL1BlockHash [32]byte, trustedL1BlockNumber uint64) error {
	return _IProofVerifier.Contract.VerifyStateUpdate(&_IProofVerifier.CallOpts, l1header, proof, appHash, trustedL1BlockHash, trustedL1BlockNumber)
}
