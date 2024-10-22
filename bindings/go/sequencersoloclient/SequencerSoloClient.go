// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sequencersoloclient

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

// SequencerSoloClientMetaData contains all meta data concerning the SequencerSoloClient contract.
var SequencerSoloClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"verifier_\",\"type\":\"address\",\"internalType\":\"contractISignatureVerifier\"},{\"name\":\"_l1BlockProvider\",\"type\":\"address\",\"internalType\":\"contractL1Block\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"LIGHT_CLIENT_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumLightClientType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"consensusStates\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getState\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1BlockProvider\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractL1Block\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"peptideHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"peptideAppHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISignatureVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expectedValue\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"AppHashHasNotPassedFraudProofWindow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotUpdateClientWithDifferentAppHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1Origin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoConsensusStateAtHeight\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NonMembershipProofsNotYetImplemented\",\"inputs\":[]}]",
}

// SequencerSoloClientABI is the input ABI used to generate the binding from.
// Deprecated: Use SequencerSoloClientMetaData.ABI instead.
var SequencerSoloClientABI = SequencerSoloClientMetaData.ABI

// SequencerSoloClient is an auto generated Go binding around an Ethereum contract.
type SequencerSoloClient struct {
	SequencerSoloClientCaller     // Read-only binding to the contract
	SequencerSoloClientTransactor // Write-only binding to the contract
	SequencerSoloClientFilterer   // Log filterer for contract events
}

// SequencerSoloClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type SequencerSoloClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSoloClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SequencerSoloClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSoloClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SequencerSoloClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SequencerSoloClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SequencerSoloClientSession struct {
	Contract     *SequencerSoloClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SequencerSoloClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SequencerSoloClientCallerSession struct {
	Contract *SequencerSoloClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SequencerSoloClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SequencerSoloClientTransactorSession struct {
	Contract     *SequencerSoloClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SequencerSoloClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type SequencerSoloClientRaw struct {
	Contract *SequencerSoloClient // Generic contract binding to access the raw methods on
}

// SequencerSoloClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SequencerSoloClientCallerRaw struct {
	Contract *SequencerSoloClientCaller // Generic read-only contract binding to access the raw methods on
}

// SequencerSoloClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SequencerSoloClientTransactorRaw struct {
	Contract *SequencerSoloClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSequencerSoloClient creates a new instance of SequencerSoloClient, bound to a specific deployed contract.
func NewSequencerSoloClient(address common.Address, backend bind.ContractBackend) (*SequencerSoloClient, error) {
	contract, err := bindSequencerSoloClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SequencerSoloClient{SequencerSoloClientCaller: SequencerSoloClientCaller{contract: contract}, SequencerSoloClientTransactor: SequencerSoloClientTransactor{contract: contract}, SequencerSoloClientFilterer: SequencerSoloClientFilterer{contract: contract}}, nil
}

// NewSequencerSoloClientCaller creates a new read-only instance of SequencerSoloClient, bound to a specific deployed contract.
func NewSequencerSoloClientCaller(address common.Address, caller bind.ContractCaller) (*SequencerSoloClientCaller, error) {
	contract, err := bindSequencerSoloClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerSoloClientCaller{contract: contract}, nil
}

// NewSequencerSoloClientTransactor creates a new write-only instance of SequencerSoloClient, bound to a specific deployed contract.
func NewSequencerSoloClientTransactor(address common.Address, transactor bind.ContractTransactor) (*SequencerSoloClientTransactor, error) {
	contract, err := bindSequencerSoloClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SequencerSoloClientTransactor{contract: contract}, nil
}

// NewSequencerSoloClientFilterer creates a new log filterer instance of SequencerSoloClient, bound to a specific deployed contract.
func NewSequencerSoloClientFilterer(address common.Address, filterer bind.ContractFilterer) (*SequencerSoloClientFilterer, error) {
	contract, err := bindSequencerSoloClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SequencerSoloClientFilterer{contract: contract}, nil
}

// bindSequencerSoloClient binds a generic wrapper to an already deployed contract.
func bindSequencerSoloClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SequencerSoloClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerSoloClient *SequencerSoloClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerSoloClient.Contract.SequencerSoloClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerSoloClient *SequencerSoloClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerSoloClient.Contract.SequencerSoloClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerSoloClient *SequencerSoloClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerSoloClient.Contract.SequencerSoloClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SequencerSoloClient *SequencerSoloClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SequencerSoloClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SequencerSoloClient *SequencerSoloClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SequencerSoloClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SequencerSoloClient *SequencerSoloClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SequencerSoloClient.Contract.contract.Transact(opts, method, params...)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_SequencerSoloClient *SequencerSoloClientCaller) LIGHTCLIENTTYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SequencerSoloClient.contract.Call(opts, &out, "LIGHT_CLIENT_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_SequencerSoloClient *SequencerSoloClientSession) LIGHTCLIENTTYPE() (uint8, error) {
	return _SequencerSoloClient.Contract.LIGHTCLIENTTYPE(&_SequencerSoloClient.CallOpts)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_SequencerSoloClient *SequencerSoloClientCallerSession) LIGHTCLIENTTYPE() (uint8, error) {
	return _SequencerSoloClient.Contract.LIGHTCLIENTTYPE(&_SequencerSoloClient.CallOpts)
}

// ConsensusStates is a free data retrieval call binding the contract method 0x1b738a22.
//
// Solidity: function consensusStates(uint256 ) view returns(uint256)
func (_SequencerSoloClient *SequencerSoloClientCaller) ConsensusStates(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SequencerSoloClient.contract.Call(opts, &out, "consensusStates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConsensusStates is a free data retrieval call binding the contract method 0x1b738a22.
//
// Solidity: function consensusStates(uint256 ) view returns(uint256)
func (_SequencerSoloClient *SequencerSoloClientSession) ConsensusStates(arg0 *big.Int) (*big.Int, error) {
	return _SequencerSoloClient.Contract.ConsensusStates(&_SequencerSoloClient.CallOpts, arg0)
}

// ConsensusStates is a free data retrieval call binding the contract method 0x1b738a22.
//
// Solidity: function consensusStates(uint256 ) view returns(uint256)
func (_SequencerSoloClient *SequencerSoloClientCallerSession) ConsensusStates(arg0 *big.Int) (*big.Int, error) {
	return _SequencerSoloClient.Contract.ConsensusStates(&_SequencerSoloClient.CallOpts, arg0)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256 appHash)
func (_SequencerSoloClient *SequencerSoloClientCaller) GetState(opts *bind.CallOpts, height *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SequencerSoloClient.contract.Call(opts, &out, "getState", height)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256 appHash)
func (_SequencerSoloClient *SequencerSoloClientSession) GetState(height *big.Int) (*big.Int, error) {
	return _SequencerSoloClient.Contract.GetState(&_SequencerSoloClient.CallOpts, height)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256 appHash)
func (_SequencerSoloClient *SequencerSoloClientCallerSession) GetState(height *big.Int) (*big.Int, error) {
	return _SequencerSoloClient.Contract.GetState(&_SequencerSoloClient.CallOpts, height)
}

// L1BlockProvider is a free data retrieval call binding the contract method 0xeb772058.
//
// Solidity: function l1BlockProvider() view returns(address)
func (_SequencerSoloClient *SequencerSoloClientCaller) L1BlockProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerSoloClient.contract.Call(opts, &out, "l1BlockProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1BlockProvider is a free data retrieval call binding the contract method 0xeb772058.
//
// Solidity: function l1BlockProvider() view returns(address)
func (_SequencerSoloClient *SequencerSoloClientSession) L1BlockProvider() (common.Address, error) {
	return _SequencerSoloClient.Contract.L1BlockProvider(&_SequencerSoloClient.CallOpts)
}

// L1BlockProvider is a free data retrieval call binding the contract method 0xeb772058.
//
// Solidity: function l1BlockProvider() view returns(address)
func (_SequencerSoloClient *SequencerSoloClientCallerSession) L1BlockProvider() (common.Address, error) {
	return _SequencerSoloClient.Contract.L1BlockProvider(&_SequencerSoloClient.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_SequencerSoloClient *SequencerSoloClientCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SequencerSoloClient.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_SequencerSoloClient *SequencerSoloClientSession) Verifier() (common.Address, error) {
	return _SequencerSoloClient.Contract.Verifier(&_SequencerSoloClient.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_SequencerSoloClient *SequencerSoloClientCallerSession) Verifier() (common.Address, error) {
	return _SequencerSoloClient.Contract.Verifier(&_SequencerSoloClient.CallOpts)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key, bytes expectedValue) view returns()
func (_SequencerSoloClient *SequencerSoloClientCaller) VerifyMembership(opts *bind.CallOpts, proof Ics23Proof, key []byte, expectedValue []byte) error {
	var out []interface{}
	err := _SequencerSoloClient.contract.Call(opts, &out, "verifyMembership", proof, key, expectedValue)

	if err != nil {
		return err
	}

	return err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key, bytes expectedValue) view returns()
func (_SequencerSoloClient *SequencerSoloClientSession) VerifyMembership(proof Ics23Proof, key []byte, expectedValue []byte) error {
	return _SequencerSoloClient.Contract.VerifyMembership(&_SequencerSoloClient.CallOpts, proof, key, expectedValue)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key, bytes expectedValue) view returns()
func (_SequencerSoloClient *SequencerSoloClientCallerSession) VerifyMembership(proof Ics23Proof, key []byte, expectedValue []byte) error {
	return _SequencerSoloClient.Contract.VerifyMembership(&_SequencerSoloClient.CallOpts, proof, key, expectedValue)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) , bytes ) pure returns()
func (_SequencerSoloClient *SequencerSoloClientCaller) VerifyNonMembership(opts *bind.CallOpts, arg0 Ics23Proof, arg1 []byte) error {
	var out []interface{}
	err := _SequencerSoloClient.contract.Call(opts, &out, "verifyNonMembership", arg0, arg1)

	if err != nil {
		return err
	}

	return err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) , bytes ) pure returns()
func (_SequencerSoloClient *SequencerSoloClientSession) VerifyNonMembership(arg0 Ics23Proof, arg1 []byte) error {
	return _SequencerSoloClient.Contract.VerifyNonMembership(&_SequencerSoloClient.CallOpts, arg0, arg1)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) , bytes ) pure returns()
func (_SequencerSoloClient *SequencerSoloClientCallerSession) VerifyNonMembership(arg0 Ics23Proof, arg1 []byte) error {
	return _SequencerSoloClient.Contract.VerifyNonMembership(&_SequencerSoloClient.CallOpts, arg0, arg1)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes proof, uint256 peptideHeight, uint256 peptideAppHash) returns()
func (_SequencerSoloClient *SequencerSoloClientTransactor) UpdateClient(opts *bind.TransactOpts, proof []byte, peptideHeight *big.Int, peptideAppHash *big.Int) (*types.Transaction, error) {
	return _SequencerSoloClient.contract.Transact(opts, "updateClient", proof, peptideHeight, peptideAppHash)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes proof, uint256 peptideHeight, uint256 peptideAppHash) returns()
func (_SequencerSoloClient *SequencerSoloClientSession) UpdateClient(proof []byte, peptideHeight *big.Int, peptideAppHash *big.Int) (*types.Transaction, error) {
	return _SequencerSoloClient.Contract.UpdateClient(&_SequencerSoloClient.TransactOpts, proof, peptideHeight, peptideAppHash)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes proof, uint256 peptideHeight, uint256 peptideAppHash) returns()
func (_SequencerSoloClient *SequencerSoloClientTransactorSession) UpdateClient(proof []byte, peptideHeight *big.Int, peptideAppHash *big.Int) (*types.Transaction, error) {
	return _SequencerSoloClient.Contract.UpdateClient(&_SequencerSoloClient.TransactOpts, proof, peptideHeight, peptideAppHash)
}
