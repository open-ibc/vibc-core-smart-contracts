// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sequencersignatureverifier

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

// SequencerSignatureVerifierMetaData contains all meta data concerning the SequencerSignatureVerifier contract.
var SequencerSignatureVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"sequencer_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SEQUENCER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofs\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyStateUpdate\",\"inputs\":[{\"name\":\"l2BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l1BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidAppHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIbcStateProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1BlockHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1BlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPacketProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedL1BlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedL1StateRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSequencerSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MethodNotImplemented\",\"inputs\":[]}]",
}

// SequencerSignatureVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerSignatureVerifierMetaData.ABI instead.
var SequencerSignatureVerifierABI = SequencerSignatureVerifierMetaData.ABI

// SequencerSignatureVerifier is an auto generated Go binding around an Ethereum contract.
type SequencerSignatureVerifier struct {
	SequencerSignatureVerifierCaller     // Read-only binding to the contract
	SequencerSignatureVerifierTransactor // Write-only binding to the contract
	SequencerSignatureVerifierFilterer   // Log filterer for contract events
}

// SequencerSignatureVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerSignatureVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSignatureVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerSignatureVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSignatureVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerSignatureVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSignatureVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerSignatureVerifierSession struct {
	Contract     *SequencerSignatureVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SequencerSignatureVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerSignatureVerifierCallerSession struct {
	Contract *SequencerSignatureVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// SequencerSignatureVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerSignatureVerifierTransactorSession struct {
	Contract     *SequencerSignatureVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// SequencerSignatureVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerSignatureVerifierRaw struct {
	Contract *SequencerSignatureVerifier // Generic contract binding to access the raw methods on
}

// SequencerSignatureVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerSignatureVerifierCallerRaw struct {
	Contract *SequencerSignatureVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerSignatureVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerSignatureVerifierTransactorRaw struct {
	Contract *SequencerSignatureVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerSignatureVerifier creates a new instance of SequencerSignatureVerifier, bound to a specific deployed contract.
func NewSequencerSignatureVerifier(address common.Address, backend bind.ContractBackend) (*SequencerSignatureVerifier, error) {
	contract, err := bindSequencerSignatureVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerSignatureVerifier{SequencerSignatureVerifierCaller: SequencerSignatureVerifierCaller{contract: contract}, SequencerSignatureVerifierTransactor: SequencerSignatureVerifierTransactor{contract: contract}, SequencerSignatureVerifierFilterer: SequencerSignatureVerifierFilterer{contract: contract}}, nil
}

// NewSequencerSignatureVerifierCaller creates a new read-only instance of SequencerSignatureVerifier, bound to a specific deployed contract.
func NewSequencerSignatureVerifierCaller(address common.Address, caller bind.ContractCaller) (*SequencerSignatureVerifierCaller, error) {
	contract, err := bindSequencerSignatureVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerSignatureVerifierCaller{contract: contract}, nil
}

// NewSequencerSignatureVerifierTransactor creates a new write-only instance of SequencerSignatureVerifier, bound to a specific deployed contract.
func NewSequencerSignatureVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerSignatureVerifierTransactor, error) {
	contract, err := bindSequencerSignatureVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerSignatureVerifierTransactor{contract: contract}, nil
}

// NewSequencerSignatureVerifierFilterer creates a new log filterer instance of SequencerSignatureVerifier, bound to a specific deployed contract.
func NewSequencerSignatureVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerSignatureVerifierFilterer, error) {
	contract, err := bindSequencerSignatureVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerSignatureVerifierFilterer{contract: contract}, nil
}

// bindSequencerSignatureVerifier binds a generic wrapper to an already deployed contract.
func bindSequencerSignatureVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SequencerSignatureVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerSignatureVerifier *SequencerSignatureVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerSignatureVerifier.Contract.SequencerSignatureVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerSignatureVerifier *SequencerSignatureVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerSignatureVerifier.Contract.SequencerSignatureVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerSignatureVerifier *SequencerSignatureVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerSignatureVerifier.Contract.SequencerSignatureVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerSignatureVerifier *SequencerSignatureVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerSignatureVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerSignatureVerifier *SequencerSignatureVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerSignatureVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerSignatureVerifier *SequencerSignatureVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerSignatureVerifier.Contract.contract.Transact(opts, method, params...)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(bytes32)
func (_SequencerSignatureVerifier *SequencerSignatureVerifierCaller) CHAINID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SequencerSignatureVerifier.contract.Call(opts, &out, "CHAIN_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(bytes32)
func (_SequencerSignatureVerifier *SequencerSignatureVerifierSession) CHAINID() ([32]byte, error) {
	return _SequencerSignatureVerifier.Contract.CHAINID(&_SequencerSignatureVerifier.CallOpts)
}

// CHAINID is a free data retrieval call binding the contract method 0x85e1f4d0.
//
// Solidity: function CHAIN_ID() view returns(bytes32)
func (_SequencerSignatureVerifier *SequencerSignatureVerifierCallerSession) CHAINID() ([32]byte, error) {
	return _SequencerSignatureVerifier.Contract.CHAINID(&_SequencerSignatureVerifier.CallOpts)
}

// SEQUENCER is a free data retrieval call binding the contract method 0x75fd4ca9.
//
// Solidity: function SEQUENCER() view returns(address)
func (_SequencerSignatureVerifier *SequencerSignatureVerifierCaller) SEQUENCER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerSignatureVerifier.contract.Call(opts, &out, "SEQUENCER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SEQUENCER is a free data retrieval call binding the contract method 0x75fd4ca9.
//
// Solidity: function SEQUENCER() view returns(address)
func (_SequencerSignatureVerifier *SequencerSignatureVerifierSession) SEQUENCER() (common.Address, error) {
	return _SequencerSignatureVerifier.Contract.SEQUENCER(&_SequencerSignatureVerifier.CallOpts)
}

// SEQUENCER is a free data retrieval call binding the contract method 0x75fd4ca9.
//
// Solidity: function SEQUENCER() view returns(address)
func (_SequencerSignatureVerifier *SequencerSignatureVerifierCallerSession) SEQUENCER() (common.Address, error) {
	return _SequencerSignatureVerifier.Contract.SEQUENCER(&_SequencerSignatureVerifier.CallOpts)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proofs) pure returns()
func (_SequencerSignatureVerifier *SequencerSignatureVerifierCaller) VerifyMembership(opts *bind.CallOpts, appHash [32]byte, key []byte, value []byte, proofs Ics23Proof) error {
	var out []interface{}
	err := _SequencerSignatureVerifier.contract.Call(opts, &out, "verifyMembership", appHash, key, value, proofs)

	if err != nil {
		return err
	}

	return err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proofs) pure returns()
func (_SequencerSignatureVerifier *SequencerSignatureVerifierSession) VerifyMembership(appHash [32]byte, key []byte, value []byte, proofs Ics23Proof) error {
	return _SequencerSignatureVerifier.Contract.VerifyMembership(&_SequencerSignatureVerifier.CallOpts, appHash, key, value, proofs)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proofs) pure returns()
func (_SequencerSignatureVerifier *SequencerSignatureVerifierCallerSession) VerifyMembership(appHash [32]byte, key []byte, value []byte, proofs Ics23Proof) error {
	return _SequencerSignatureVerifier.Contract.VerifyMembership(&_SequencerSignatureVerifier.CallOpts, appHash, key, value, proofs)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_SequencerSignatureVerifier *SequencerSignatureVerifierCaller) VerifyNonMembership(opts *bind.CallOpts, arg0 [32]byte, arg1 []byte, arg2 Ics23Proof) error {
	var out []interface{}
	err := _SequencerSignatureVerifier.contract.Call(opts, &out, "verifyNonMembership", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_SequencerSignatureVerifier *SequencerSignatureVerifierSession) VerifyNonMembership(arg0 [32]byte, arg1 []byte, arg2 Ics23Proof) error {
	return _SequencerSignatureVerifier.Contract.VerifyNonMembership(&_SequencerSignatureVerifier.CallOpts, arg0, arg1, arg2)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_SequencerSignatureVerifier *SequencerSignatureVerifierCallerSession) VerifyNonMembership(arg0 [32]byte, arg1 []byte, arg2 Ics23Proof) error {
	return _SequencerSignatureVerifier.Contract.VerifyNonMembership(&_SequencerSignatureVerifier.CallOpts, arg0, arg1, arg2)
}

// VerifyStateUpdate is a free data retrieval call binding the contract method 0xe8d90039.
//
// Solidity: function verifyStateUpdate(uint256 l2BlockNumber, bytes32 appHash, bytes32 l1BlockHash, bytes signature) view returns()
func (_SequencerSignatureVerifier *SequencerSignatureVerifierCaller) VerifyStateUpdate(opts *bind.CallOpts, l2BlockNumber *big.Int, appHash [32]byte, l1BlockHash [32]byte, signature []byte) error {
	var out []interface{}
	err := _SequencerSignatureVerifier.contract.Call(opts, &out, "verifyStateUpdate", l2BlockNumber, appHash, l1BlockHash, signature)

	if err != nil {
		return err
	}

	return err

}

// VerifyStateUpdate is a free data retrieval call binding the contract method 0xe8d90039.
//
// Solidity: function verifyStateUpdate(uint256 l2BlockNumber, bytes32 appHash, bytes32 l1BlockHash, bytes signature) view returns()
func (_SequencerSignatureVerifier *SequencerSignatureVerifierSession) VerifyStateUpdate(l2BlockNumber *big.Int, appHash [32]byte, l1BlockHash [32]byte, signature []byte) error {
	return _SequencerSignatureVerifier.Contract.VerifyStateUpdate(&_SequencerSignatureVerifier.CallOpts, l2BlockNumber, appHash, l1BlockHash, signature)
}

// VerifyStateUpdate is a free data retrieval call binding the contract method 0xe8d90039.
//
// Solidity: function verifyStateUpdate(uint256 l2BlockNumber, bytes32 appHash, bytes32 l1BlockHash, bytes signature) view returns()
func (_SequencerSignatureVerifier *SequencerSignatureVerifierCallerSession) VerifyStateUpdate(l2BlockNumber *big.Int, appHash [32]byte, l1BlockHash [32]byte, signature []byte) error {
	return _SequencerSignatureVerifier.Contract.VerifyStateUpdate(&_SequencerSignatureVerifier.CallOpts, l2BlockNumber, appHash, l1BlockHash, signature)
}
