// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dummylightclient

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

// DummyLightClientMetaData contains all meta data concerning the DummyLightClient contract.
var DummyLightClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"LIGHT_CLIENT_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumLightClientType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"consensusStates\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLightClientType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumLightClientType\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getState\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"InvalidDummyMembershipProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDummyNonMembershipProof\",\"inputs\":[]}]",
}

// DummyLightClientABI is the input ABI used to generate the binding from.
// Deprecated: Use DummyLightClientMetaData.ABI instead.
var DummyLightClientABI = DummyLightClientMetaData.ABI

// DummyLightClient is an auto generated Go binding around an Ethereum contract.
type DummyLightClient struct {
	DummyLightClientCaller     // Read-only binding to the contract
	DummyLightClientTransactor // Write-only binding to the contract
	DummyLightClientFilterer   // Log filterer for contract events
}

// DummyLightClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type DummyLightClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyLightClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DummyLightClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyLightClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DummyLightClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyLightClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DummyLightClientSession struct {
	Contract     *DummyLightClient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DummyLightClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DummyLightClientCallerSession struct {
	Contract *DummyLightClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DummyLightClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DummyLightClientTransactorSession struct {
	Contract     *DummyLightClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DummyLightClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type DummyLightClientRaw struct {
	Contract *DummyLightClient // Generic contract binding to access the raw methods on
}

// DummyLightClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DummyLightClientCallerRaw struct {
	Contract *DummyLightClientCaller // Generic read-only contract binding to access the raw methods on
}

// DummyLightClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DummyLightClientTransactorRaw struct {
	Contract *DummyLightClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDummyLightClient creates a new instance of DummyLightClient, bound to a specific deployed contract.
func NewDummyLightClient(address common.Address, backend bind.ContractBackend) (*DummyLightClient, error) {
	contract, err := bindDummyLightClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DummyLightClient{DummyLightClientCaller: DummyLightClientCaller{contract: contract}, DummyLightClientTransactor: DummyLightClientTransactor{contract: contract}, DummyLightClientFilterer: DummyLightClientFilterer{contract: contract}}, nil
}

// NewDummyLightClientCaller creates a new read-only instance of DummyLightClient, bound to a specific deployed contract.
func NewDummyLightClientCaller(address common.Address, caller bind.ContractCaller) (*DummyLightClientCaller, error) {
	contract, err := bindDummyLightClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DummyLightClientCaller{contract: contract}, nil
}

// NewDummyLightClientTransactor creates a new write-only instance of DummyLightClient, bound to a specific deployed contract.
func NewDummyLightClientTransactor(address common.Address, transactor bind.ContractTransactor) (*DummyLightClientTransactor, error) {
	contract, err := bindDummyLightClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DummyLightClientTransactor{contract: contract}, nil
}

// NewDummyLightClientFilterer creates a new log filterer instance of DummyLightClient, bound to a specific deployed contract.
func NewDummyLightClientFilterer(address common.Address, filterer bind.ContractFilterer) (*DummyLightClientFilterer, error) {
	contract, err := bindDummyLightClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DummyLightClientFilterer{contract: contract}, nil
}

// bindDummyLightClient binds a generic wrapper to an already deployed contract.
func bindDummyLightClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DummyLightClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyLightClient *DummyLightClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DummyLightClient.Contract.DummyLightClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyLightClient *DummyLightClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyLightClient.Contract.DummyLightClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyLightClient *DummyLightClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyLightClient.Contract.DummyLightClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyLightClient *DummyLightClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DummyLightClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyLightClient *DummyLightClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyLightClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyLightClient *DummyLightClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyLightClient.Contract.contract.Transact(opts, method, params...)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_DummyLightClient *DummyLightClientCaller) LIGHTCLIENTTYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DummyLightClient.contract.Call(opts, &out, "LIGHT_CLIENT_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_DummyLightClient *DummyLightClientSession) LIGHTCLIENTTYPE() (uint8, error) {
	return _DummyLightClient.Contract.LIGHTCLIENTTYPE(&_DummyLightClient.CallOpts)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_DummyLightClient *DummyLightClientCallerSession) LIGHTCLIENTTYPE() (uint8, error) {
	return _DummyLightClient.Contract.LIGHTCLIENTTYPE(&_DummyLightClient.CallOpts)
}

// ConsensusStates is a free data retrieval call binding the contract method 0x1b738a22.
//
// Solidity: function consensusStates(uint256 ) view returns(uint256)
func (_DummyLightClient *DummyLightClientCaller) ConsensusStates(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DummyLightClient.contract.Call(opts, &out, "consensusStates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConsensusStates is a free data retrieval call binding the contract method 0x1b738a22.
//
// Solidity: function consensusStates(uint256 ) view returns(uint256)
func (_DummyLightClient *DummyLightClientSession) ConsensusStates(arg0 *big.Int) (*big.Int, error) {
	return _DummyLightClient.Contract.ConsensusStates(&_DummyLightClient.CallOpts, arg0)
}

// ConsensusStates is a free data retrieval call binding the contract method 0x1b738a22.
//
// Solidity: function consensusStates(uint256 ) view returns(uint256)
func (_DummyLightClient *DummyLightClientCallerSession) ConsensusStates(arg0 *big.Int) (*big.Int, error) {
	return _DummyLightClient.Contract.ConsensusStates(&_DummyLightClient.CallOpts, arg0)
}

// GetLightClientType is a free data retrieval call binding the contract method 0x36028d2a.
//
// Solidity: function getLightClientType() pure returns(uint8)
func (_DummyLightClient *DummyLightClientCaller) GetLightClientType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DummyLightClient.contract.Call(opts, &out, "getLightClientType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetLightClientType is a free data retrieval call binding the contract method 0x36028d2a.
//
// Solidity: function getLightClientType() pure returns(uint8)
func (_DummyLightClient *DummyLightClientSession) GetLightClientType() (uint8, error) {
	return _DummyLightClient.Contract.GetLightClientType(&_DummyLightClient.CallOpts)
}

// GetLightClientType is a free data retrieval call binding the contract method 0x36028d2a.
//
// Solidity: function getLightClientType() pure returns(uint8)
func (_DummyLightClient *DummyLightClientCallerSession) GetLightClientType() (uint8, error) {
	return _DummyLightClient.Contract.GetLightClientType(&_DummyLightClient.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 ) pure returns(uint256 appHash)
func (_DummyLightClient *DummyLightClientCaller) GetState(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DummyLightClient.contract.Call(opts, &out, "getState", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 ) pure returns(uint256 appHash)
func (_DummyLightClient *DummyLightClientSession) GetState(arg0 *big.Int) (*big.Int, error) {
	return _DummyLightClient.Contract.GetState(&_DummyLightClient.CallOpts, arg0)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 ) pure returns(uint256 appHash)
func (_DummyLightClient *DummyLightClientCallerSession) GetState(arg0 *big.Int) (*big.Int, error) {
	return _DummyLightClient.Contract.GetState(&_DummyLightClient.CallOpts, arg0)
}

// UpdateClient is a free data retrieval call binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes , uint256 , uint256 ) pure returns()
func (_DummyLightClient *DummyLightClientCaller) UpdateClient(opts *bind.CallOpts, arg0 []byte, arg1 *big.Int, arg2 *big.Int) error {
	var out []interface{}
	err := _DummyLightClient.contract.Call(opts, &out, "updateClient", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// UpdateClient is a free data retrieval call binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes , uint256 , uint256 ) pure returns()
func (_DummyLightClient *DummyLightClientSession) UpdateClient(arg0 []byte, arg1 *big.Int, arg2 *big.Int) error {
	return _DummyLightClient.Contract.UpdateClient(&_DummyLightClient.CallOpts, arg0, arg1, arg2)
}

// UpdateClient is a free data retrieval call binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes , uint256 , uint256 ) pure returns()
func (_DummyLightClient *DummyLightClientCallerSession) UpdateClient(arg0 []byte, arg1 *big.Int, arg2 *big.Int) error {
	return _DummyLightClient.Contract.UpdateClient(&_DummyLightClient.CallOpts, arg0, arg1, arg2)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes , bytes ) pure returns()
func (_DummyLightClient *DummyLightClientCaller) VerifyMembership(opts *bind.CallOpts, proof Ics23Proof, arg1 []byte, arg2 []byte) error {
	var out []interface{}
	err := _DummyLightClient.contract.Call(opts, &out, "verifyMembership", proof, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes , bytes ) pure returns()
func (_DummyLightClient *DummyLightClientSession) VerifyMembership(proof Ics23Proof, arg1 []byte, arg2 []byte) error {
	return _DummyLightClient.Contract.VerifyMembership(&_DummyLightClient.CallOpts, proof, arg1, arg2)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes , bytes ) pure returns()
func (_DummyLightClient *DummyLightClientCallerSession) VerifyMembership(proof Ics23Proof, arg1 []byte, arg2 []byte) error {
	return _DummyLightClient.Contract.VerifyMembership(&_DummyLightClient.CallOpts, proof, arg1, arg2)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes ) pure returns()
func (_DummyLightClient *DummyLightClientCaller) VerifyNonMembership(opts *bind.CallOpts, proof Ics23Proof, arg1 []byte) error {
	var out []interface{}
	err := _DummyLightClient.contract.Call(opts, &out, "verifyNonMembership", proof, arg1)

	if err != nil {
		return err
	}

	return err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes ) pure returns()
func (_DummyLightClient *DummyLightClientSession) VerifyNonMembership(proof Ics23Proof, arg1 []byte) error {
	return _DummyLightClient.Contract.VerifyNonMembership(&_DummyLightClient.CallOpts, proof, arg1)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes ) pure returns()
func (_DummyLightClient *DummyLightClientCallerSession) VerifyNonMembership(proof Ics23Proof, arg1 []byte) error {
	return _DummyLightClient.Contract.VerifyNonMembership(&_DummyLightClient.CallOpts, proof, arg1)
}
