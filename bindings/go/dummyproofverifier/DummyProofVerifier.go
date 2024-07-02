// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dummyproofverifier

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

// DummyProofVerifierMetaData contains all meta data concerning the DummyProofVerifier contract.
var DummyProofVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyStateUpdate\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structL1Header\",\"components\":[{\"name\":\"header\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"number\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOpL2StateProof\",\"components\":[{\"name\":\"accountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"outputRootProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"l2OutputProposalKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"InvalidAppHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIbcStateProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1BlockHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1BlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPacketProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedL1BlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedL1StateRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MethodNotImplemented\",\"inputs\":[]}]",
}

// DummyProofVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use DummyProofVerifierMetaData.ABI instead.
var DummyProofVerifierABI = DummyProofVerifierMetaData.ABI

// DummyProofVerifier is an auto generated Go binding around an Ethereum contract.
type DummyProofVerifier struct {
	DummyProofVerifierCaller     // Read-only binding to the contract
	DummyProofVerifierTransactor // Write-only binding to the contract
	DummyProofVerifierFilterer   // Log filterer for contract events
}

// DummyProofVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type DummyProofVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyProofVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DummyProofVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyProofVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DummyProofVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyProofVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DummyProofVerifierSession struct {
	Contract     *DummyProofVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DummyProofVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DummyProofVerifierCallerSession struct {
	Contract *DummyProofVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DummyProofVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DummyProofVerifierTransactorSession struct {
	Contract     *DummyProofVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DummyProofVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type DummyProofVerifierRaw struct {
	Contract *DummyProofVerifier // Generic contract binding to access the raw methods on
}

// DummyProofVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DummyProofVerifierCallerRaw struct {
	Contract *DummyProofVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// DummyProofVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DummyProofVerifierTransactorRaw struct {
	Contract *DummyProofVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDummyProofVerifier creates a new instance of DummyProofVerifier, bound to a specific deployed contract.
func NewDummyProofVerifier(address common.Address, backend bind.ContractBackend) (*DummyProofVerifier, error) {
	contract, err := bindDummyProofVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DummyProofVerifier{DummyProofVerifierCaller: DummyProofVerifierCaller{contract: contract}, DummyProofVerifierTransactor: DummyProofVerifierTransactor{contract: contract}, DummyProofVerifierFilterer: DummyProofVerifierFilterer{contract: contract}}, nil
}

// NewDummyProofVerifierCaller creates a new read-only instance of DummyProofVerifier, bound to a specific deployed contract.
func NewDummyProofVerifierCaller(address common.Address, caller bind.ContractCaller) (*DummyProofVerifierCaller, error) {
	contract, err := bindDummyProofVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DummyProofVerifierCaller{contract: contract}, nil
}

// NewDummyProofVerifierTransactor creates a new write-only instance of DummyProofVerifier, bound to a specific deployed contract.
func NewDummyProofVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*DummyProofVerifierTransactor, error) {
	contract, err := bindDummyProofVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DummyProofVerifierTransactor{contract: contract}, nil
}

// NewDummyProofVerifierFilterer creates a new log filterer instance of DummyProofVerifier, bound to a specific deployed contract.
func NewDummyProofVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*DummyProofVerifierFilterer, error) {
	contract, err := bindDummyProofVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DummyProofVerifierFilterer{contract: contract}, nil
}

// bindDummyProofVerifier binds a generic wrapper to an already deployed contract.
func bindDummyProofVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DummyProofVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyProofVerifier *DummyProofVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DummyProofVerifier.Contract.DummyProofVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyProofVerifier *DummyProofVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyProofVerifier.Contract.DummyProofVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyProofVerifier *DummyProofVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyProofVerifier.Contract.DummyProofVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyProofVerifier *DummyProofVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DummyProofVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyProofVerifier *DummyProofVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyProofVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyProofVerifier *DummyProofVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyProofVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 , bytes , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_DummyProofVerifier *DummyProofVerifierCaller) VerifyMembership(opts *bind.CallOpts, arg0 [32]byte, arg1 []byte, arg2 []byte, arg3 Ics23Proof) error {
	var out []interface{}
	err := _DummyProofVerifier.contract.Call(opts, &out, "verifyMembership", arg0, arg1, arg2, arg3)

	if err != nil {
		return err
	}

	return err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 , bytes , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_DummyProofVerifier *DummyProofVerifierSession) VerifyMembership(arg0 [32]byte, arg1 []byte, arg2 []byte, arg3 Ics23Proof) error {
	return _DummyProofVerifier.Contract.VerifyMembership(&_DummyProofVerifier.CallOpts, arg0, arg1, arg2, arg3)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 , bytes , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_DummyProofVerifier *DummyProofVerifierCallerSession) VerifyMembership(arg0 [32]byte, arg1 []byte, arg2 []byte, arg3 Ics23Proof) error {
	return _DummyProofVerifier.Contract.VerifyMembership(&_DummyProofVerifier.CallOpts, arg0, arg1, arg2, arg3)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_DummyProofVerifier *DummyProofVerifierCaller) VerifyNonMembership(opts *bind.CallOpts, arg0 [32]byte, arg1 []byte, arg2 Ics23Proof) error {
	var out []interface{}
	err := _DummyProofVerifier.contract.Call(opts, &out, "verifyNonMembership", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_DummyProofVerifier *DummyProofVerifierSession) VerifyNonMembership(arg0 [32]byte, arg1 []byte, arg2 Ics23Proof) error {
	return _DummyProofVerifier.Contract.VerifyNonMembership(&_DummyProofVerifier.CallOpts, arg0, arg1, arg2)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_DummyProofVerifier *DummyProofVerifierCallerSession) VerifyNonMembership(arg0 [32]byte, arg1 []byte, arg2 Ics23Proof) error {
	return _DummyProofVerifier.Contract.VerifyNonMembership(&_DummyProofVerifier.CallOpts, arg0, arg1, arg2)
}

// VerifyStateUpdate is a free data retrieval call binding the contract method 0x0a1bb8b5.
//
// Solidity: function verifyStateUpdate((bytes[],bytes32,uint64) , (bytes[],bytes[],bytes32,bytes32) , bytes32 , bytes32 , uint64 ) pure returns()
func (_DummyProofVerifier *DummyProofVerifierCaller) VerifyStateUpdate(opts *bind.CallOpts, arg0 L1Header, arg1 OpL2StateProof, arg2 [32]byte, arg3 [32]byte, arg4 uint64) error {
	var out []interface{}
	err := _DummyProofVerifier.contract.Call(opts, &out, "verifyStateUpdate", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return err
	}

	return err

}

// VerifyStateUpdate is a free data retrieval call binding the contract method 0x0a1bb8b5.
//
// Solidity: function verifyStateUpdate((bytes[],bytes32,uint64) , (bytes[],bytes[],bytes32,bytes32) , bytes32 , bytes32 , uint64 ) pure returns()
func (_DummyProofVerifier *DummyProofVerifierSession) VerifyStateUpdate(arg0 L1Header, arg1 OpL2StateProof, arg2 [32]byte, arg3 [32]byte, arg4 uint64) error {
	return _DummyProofVerifier.Contract.VerifyStateUpdate(&_DummyProofVerifier.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// VerifyStateUpdate is a free data retrieval call binding the contract method 0x0a1bb8b5.
//
// Solidity: function verifyStateUpdate((bytes[],bytes32,uint64) , (bytes[],bytes[],bytes32,bytes32) , bytes32 , bytes32 , uint64 ) pure returns()
func (_DummyProofVerifier *DummyProofVerifierCallerSession) VerifyStateUpdate(arg0 L1Header, arg1 OpL2StateProof, arg2 [32]byte, arg3 [32]byte, arg4 uint64) error {
	return _DummyProofVerifier.Contract.VerifyStateUpdate(&_DummyProofVerifier.CallOpts, arg0, arg1, arg2, arg3, arg4)
}
