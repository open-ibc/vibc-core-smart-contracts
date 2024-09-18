// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optimisticproofverifier

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

// OptimisticProofVerifierMetaData contains all meta data concerning the OptimisticProofVerifier contract.
var OptimisticProofVerifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_l2OutputOracleAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l2OutputOracleAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofs\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyStateUpdate\",\"inputs\":[{\"name\":\"l1header\",\"type\":\"tuple\",\"internalType\":\"structL1Header\",\"components\":[{\"name\":\"header\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"number\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structOpL2StateProof\",\"components\":[{\"name\":\"accountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"outputRootProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"l2OutputProposalKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"peptideAppHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"trustedL1BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"trustedL1BlockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"InvalidAppHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIbcStateProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1BlockHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1BlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPacketProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedL1BlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedL1StateRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MethodNotImplemented\",\"inputs\":[]}]",
}

// OptimisticProofVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use OptimisticProofVerifierMetaData.ABI instead.
var OptimisticProofVerifierABI = OptimisticProofVerifierMetaData.ABI

// OptimisticProofVerifier is an auto generated Go binding around an Ethereum contract.
type OptimisticProofVerifier struct {
	OptimisticProofVerifierCaller     // Read-only binding to the contract
	OptimisticProofVerifierTransactor // Write-only binding to the contract
	OptimisticProofVerifierFilterer   // Log filterer for contract events
}

// OptimisticProofVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptimisticProofVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimisticProofVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptimisticProofVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimisticProofVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptimisticProofVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptimisticProofVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptimisticProofVerifierSession struct {
	Contract     *OptimisticProofVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OptimisticProofVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptimisticProofVerifierCallerSession struct {
	Contract *OptimisticProofVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// OptimisticProofVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptimisticProofVerifierTransactorSession struct {
	Contract     *OptimisticProofVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// OptimisticProofVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptimisticProofVerifierRaw struct {
	Contract *OptimisticProofVerifier // Generic contract binding to access the raw methods on
}

// OptimisticProofVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptimisticProofVerifierCallerRaw struct {
	Contract *OptimisticProofVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// OptimisticProofVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptimisticProofVerifierTransactorRaw struct {
	Contract *OptimisticProofVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptimisticProofVerifier creates a new instance of OptimisticProofVerifier, bound to a specific deployed contract.
func NewOptimisticProofVerifier(address common.Address, backend bind.ContractBackend) (*OptimisticProofVerifier, error) {
	contract, err := bindOptimisticProofVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptimisticProofVerifier{OptimisticProofVerifierCaller: OptimisticProofVerifierCaller{contract: contract}, OptimisticProofVerifierTransactor: OptimisticProofVerifierTransactor{contract: contract}, OptimisticProofVerifierFilterer: OptimisticProofVerifierFilterer{contract: contract}}, nil
}

// NewOptimisticProofVerifierCaller creates a new read-only instance of OptimisticProofVerifier, bound to a specific deployed contract.
func NewOptimisticProofVerifierCaller(address common.Address, caller bind.ContractCaller) (*OptimisticProofVerifierCaller, error) {
	contract, err := bindOptimisticProofVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptimisticProofVerifierCaller{contract: contract}, nil
}

// NewOptimisticProofVerifierTransactor creates a new write-only instance of OptimisticProofVerifier, bound to a specific deployed contract.
func NewOptimisticProofVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*OptimisticProofVerifierTransactor, error) {
	contract, err := bindOptimisticProofVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptimisticProofVerifierTransactor{contract: contract}, nil
}

// NewOptimisticProofVerifierFilterer creates a new log filterer instance of OptimisticProofVerifier, bound to a specific deployed contract.
func NewOptimisticProofVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*OptimisticProofVerifierFilterer, error) {
	contract, err := bindOptimisticProofVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptimisticProofVerifierFilterer{contract: contract}, nil
}

// bindOptimisticProofVerifier binds a generic wrapper to an already deployed contract.
func bindOptimisticProofVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptimisticProofVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimisticProofVerifier *OptimisticProofVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimisticProofVerifier.Contract.OptimisticProofVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimisticProofVerifier *OptimisticProofVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimisticProofVerifier.Contract.OptimisticProofVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimisticProofVerifier *OptimisticProofVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimisticProofVerifier.Contract.OptimisticProofVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptimisticProofVerifier *OptimisticProofVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptimisticProofVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptimisticProofVerifier *OptimisticProofVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptimisticProofVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptimisticProofVerifier *OptimisticProofVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptimisticProofVerifier.Contract.contract.Transact(opts, method, params...)
}

// L2OutputOracleAddress is a free data retrieval call binding the contract method 0x59c1b56b.
//
// Solidity: function l2OutputOracleAddress() view returns(address)
func (_OptimisticProofVerifier *OptimisticProofVerifierCaller) L2OutputOracleAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OptimisticProofVerifier.contract.Call(opts, &out, "l2OutputOracleAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2OutputOracleAddress is a free data retrieval call binding the contract method 0x59c1b56b.
//
// Solidity: function l2OutputOracleAddress() view returns(address)
func (_OptimisticProofVerifier *OptimisticProofVerifierSession) L2OutputOracleAddress() (common.Address, error) {
	return _OptimisticProofVerifier.Contract.L2OutputOracleAddress(&_OptimisticProofVerifier.CallOpts)
}

// L2OutputOracleAddress is a free data retrieval call binding the contract method 0x59c1b56b.
//
// Solidity: function l2OutputOracleAddress() view returns(address)
func (_OptimisticProofVerifier *OptimisticProofVerifierCallerSession) L2OutputOracleAddress() (common.Address, error) {
	return _OptimisticProofVerifier.Contract.L2OutputOracleAddress(&_OptimisticProofVerifier.CallOpts)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proofs) pure returns()
func (_OptimisticProofVerifier *OptimisticProofVerifierCaller) VerifyMembership(opts *bind.CallOpts, appHash [32]byte, key []byte, value []byte, proofs Ics23Proof) error {
	var out []interface{}
	err := _OptimisticProofVerifier.contract.Call(opts, &out, "verifyMembership", appHash, key, value, proofs)

	if err != nil {
		return err
	}

	return err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proofs) pure returns()
func (_OptimisticProofVerifier *OptimisticProofVerifierSession) VerifyMembership(appHash [32]byte, key []byte, value []byte, proofs Ics23Proof) error {
	return _OptimisticProofVerifier.Contract.VerifyMembership(&_OptimisticProofVerifier.CallOpts, appHash, key, value, proofs)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proofs) pure returns()
func (_OptimisticProofVerifier *OptimisticProofVerifierCallerSession) VerifyMembership(appHash [32]byte, key []byte, value []byte, proofs Ics23Proof) error {
	return _OptimisticProofVerifier.Contract.VerifyMembership(&_OptimisticProofVerifier.CallOpts, appHash, key, value, proofs)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_OptimisticProofVerifier *OptimisticProofVerifierCaller) VerifyNonMembership(opts *bind.CallOpts, arg0 [32]byte, arg1 []byte, arg2 Ics23Proof) error {
	var out []interface{}
	err := _OptimisticProofVerifier.contract.Call(opts, &out, "verifyNonMembership", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_OptimisticProofVerifier *OptimisticProofVerifierSession) VerifyNonMembership(arg0 [32]byte, arg1 []byte, arg2 Ics23Proof) error {
	return _OptimisticProofVerifier.Contract.VerifyNonMembership(&_OptimisticProofVerifier.CallOpts, arg0, arg1, arg2)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_OptimisticProofVerifier *OptimisticProofVerifierCallerSession) VerifyNonMembership(arg0 [32]byte, arg1 []byte, arg2 Ics23Proof) error {
	return _OptimisticProofVerifier.Contract.VerifyNonMembership(&_OptimisticProofVerifier.CallOpts, arg0, arg1, arg2)
}

// VerifyStateUpdate is a free data retrieval call binding the contract method 0x0a1bb8b5.
//
// Solidity: function verifyStateUpdate((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, bytes32 peptideAppHash, bytes32 trustedL1BlockHash, uint64 trustedL1BlockNumber) view returns()
func (_OptimisticProofVerifier *OptimisticProofVerifierCaller) VerifyStateUpdate(opts *bind.CallOpts, l1header L1Header, proof OpL2StateProof, peptideAppHash [32]byte, trustedL1BlockHash [32]byte, trustedL1BlockNumber uint64) error {
	var out []interface{}
	err := _OptimisticProofVerifier.contract.Call(opts, &out, "verifyStateUpdate", l1header, proof, peptideAppHash, trustedL1BlockHash, trustedL1BlockNumber)

	if err != nil {
		return err
	}

	return err

}

// VerifyStateUpdate is a free data retrieval call binding the contract method 0x0a1bb8b5.
//
// Solidity: function verifyStateUpdate((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, bytes32 peptideAppHash, bytes32 trustedL1BlockHash, uint64 trustedL1BlockNumber) view returns()
func (_OptimisticProofVerifier *OptimisticProofVerifierSession) VerifyStateUpdate(l1header L1Header, proof OpL2StateProof, peptideAppHash [32]byte, trustedL1BlockHash [32]byte, trustedL1BlockNumber uint64) error {
	return _OptimisticProofVerifier.Contract.VerifyStateUpdate(&_OptimisticProofVerifier.CallOpts, l1header, proof, peptideAppHash, trustedL1BlockHash, trustedL1BlockNumber)
}

// VerifyStateUpdate is a free data retrieval call binding the contract method 0x0a1bb8b5.
//
// Solidity: function verifyStateUpdate((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, bytes32 peptideAppHash, bytes32 trustedL1BlockHash, uint64 trustedL1BlockNumber) view returns()
func (_OptimisticProofVerifier *OptimisticProofVerifierCallerSession) VerifyStateUpdate(l1header L1Header, proof OpL2StateProof, peptideAppHash [32]byte, trustedL1BlockHash [32]byte, trustedL1BlockNumber uint64) error {
	return _OptimisticProofVerifier.Contract.VerifyStateUpdate(&_OptimisticProofVerifier.CallOpts, l1header, proof, peptideAppHash, trustedL1BlockHash, trustedL1BlockNumber)
}
