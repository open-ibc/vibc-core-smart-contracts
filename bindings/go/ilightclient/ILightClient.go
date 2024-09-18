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

// ILightClientMetaData contains all meta data concerning the ILightClient contract.
var ILightClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getFraudProofEndtime\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getState\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fraudProofEndTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ended\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"l1header\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fraudProofEndTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ended\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expectedValue\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// ILightClientABI is the input ABI used to generate the binding from.
// Deprecated: Use ILightClientMetaData.ABI instead.
var ILightClientABI = ILightClientMetaData.ABI

// ILightClient is an auto generated Go binding around an Ethereum contract.
type ILightClient struct {
	ILightClientCaller     // Read-only binding to the contract
	ILightClientTransactor // Write-only binding to the contract
	ILightClientFilterer   // Log filterer for contract events
}

// ILightClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILightClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILightClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILightClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILightClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILightClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILightClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILightClientSession struct {
	Contract     *ILightClient     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILightClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILightClientCallerSession struct {
	Contract *ILightClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ILightClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILightClientTransactorSession struct {
	Contract     *ILightClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ILightClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILightClientRaw struct {
	Contract *ILightClient // Generic contract binding to access the raw methods on
}

// ILightClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILightClientCallerRaw struct {
	Contract *ILightClientCaller // Generic read-only contract binding to access the raw methods on
}

// ILightClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILightClientTransactorRaw struct {
	Contract *ILightClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILightClient creates a new instance of ILightClient, bound to a specific deployed contract.
func NewILightClient(address common.Address, backend bind.ContractBackend) (*ILightClient, error) {
	contract, err := bindILightClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILightClient{ILightClientCaller: ILightClientCaller{contract: contract}, ILightClientTransactor: ILightClientTransactor{contract: contract}, ILightClientFilterer: ILightClientFilterer{contract: contract}}, nil
}

// NewILightClientCaller creates a new read-only instance of ILightClient, bound to a specific deployed contract.
func NewILightClientCaller(address common.Address, caller bind.ContractCaller) (*ILightClientCaller, error) {
	contract, err := bindILightClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILightClientCaller{contract: contract}, nil
}

// NewILightClientTransactor creates a new write-only instance of ILightClient, bound to a specific deployed contract.
func NewILightClientTransactor(address common.Address, transactor bind.ContractTransactor) (*ILightClientTransactor, error) {
	contract, err := bindILightClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILightClientTransactor{contract: contract}, nil
}

// NewILightClientFilterer creates a new log filterer instance of ILightClient, bound to a specific deployed contract.
func NewILightClientFilterer(address common.Address, filterer bind.ContractFilterer) (*ILightClientFilterer, error) {
	contract, err := bindILightClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILightClientFilterer{contract: contract}, nil
}

// bindILightClient binds a generic wrapper to an already deployed contract.
func bindILightClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ILightClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILightClient *ILightClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILightClient.Contract.ILightClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILightClient *ILightClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILightClient.Contract.ILightClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILightClient *ILightClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILightClient.Contract.ILightClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILightClient *ILightClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILightClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILightClient *ILightClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILightClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILightClient *ILightClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILightClient.Contract.contract.Transact(opts, method, params...)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_ILightClient *ILightClientCaller) GetState(opts *bind.CallOpts, height *big.Int) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	var out []interface{}
	err := _ILightClient.contract.Call(opts, &out, "getState", height)

	outstruct := new(struct {
		AppHash           *big.Int
		FraudProofEndTime *big.Int
		Ended             bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AppHash = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FraudProofEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Ended = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_ILightClient *ILightClientSession) GetState(height *big.Int) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	return _ILightClient.Contract.GetState(&_ILightClient.CallOpts, height)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_ILightClient *ILightClientCallerSession) GetState(height *big.Int) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	return _ILightClient.Contract.GetState(&_ILightClient.CallOpts, height)
}

// GetFraudProofEndtime is a paid mutator transaction binding the contract method 0xd56ff842.
//
// Solidity: function getFraudProofEndtime(uint256 height) returns(uint256 endTime)
func (_ILightClient *ILightClientTransactor) GetFraudProofEndtime(opts *bind.TransactOpts, height *big.Int) (*types.Transaction, error) {
	return _ILightClient.contract.Transact(opts, "getFraudProofEndtime", height)
}

// GetFraudProofEndtime is a paid mutator transaction binding the contract method 0xd56ff842.
//
// Solidity: function getFraudProofEndtime(uint256 height) returns(uint256 endTime)
func (_ILightClient *ILightClientSession) GetFraudProofEndtime(height *big.Int) (*types.Transaction, error) {
	return _ILightClient.Contract.GetFraudProofEndtime(&_ILightClient.TransactOpts, height)
}

// GetFraudProofEndtime is a paid mutator transaction binding the contract method 0xd56ff842.
//
// Solidity: function getFraudProofEndtime(uint256 height) returns(uint256 endTime)
func (_ILightClient *ILightClientTransactorSession) GetFraudProofEndtime(height *big.Int) (*types.Transaction, error) {
	return _ILightClient.Contract.GetFraudProofEndtime(&_ILightClient.TransactOpts, height)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xd4af812d.
//
// Solidity: function updateClient(bytes l1header, bytes proof, uint256 height, uint256 appHash) returns(uint256 fraudProofEndTime, bool ended)
func (_ILightClient *ILightClientTransactor) UpdateClient(opts *bind.TransactOpts, l1header []byte, proof []byte, height *big.Int, appHash *big.Int) (*types.Transaction, error) {
	return _ILightClient.contract.Transact(opts, "updateClient", l1header, proof, height, appHash)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xd4af812d.
//
// Solidity: function updateClient(bytes l1header, bytes proof, uint256 height, uint256 appHash) returns(uint256 fraudProofEndTime, bool ended)
func (_ILightClient *ILightClientSession) UpdateClient(l1header []byte, proof []byte, height *big.Int, appHash *big.Int) (*types.Transaction, error) {
	return _ILightClient.Contract.UpdateClient(&_ILightClient.TransactOpts, l1header, proof, height, appHash)
}

// UpdateClient is a paid mutator transaction binding the contract method 0xd4af812d.
//
// Solidity: function updateClient(bytes l1header, bytes proof, uint256 height, uint256 appHash) returns(uint256 fraudProofEndTime, bool ended)
func (_ILightClient *ILightClientTransactorSession) UpdateClient(l1header []byte, proof []byte, height *big.Int, appHash *big.Int) (*types.Transaction, error) {
	return _ILightClient.Contract.UpdateClient(&_ILightClient.TransactOpts, l1header, proof, height, appHash)
}

// VerifyMembership is a paid mutator transaction binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key, bytes expectedValue) returns()
func (_ILightClient *ILightClientTransactor) VerifyMembership(opts *bind.TransactOpts, proof Ics23Proof, key []byte, expectedValue []byte) (*types.Transaction, error) {
	return _ILightClient.contract.Transact(opts, "verifyMembership", proof, key, expectedValue)
}

// VerifyMembership is a paid mutator transaction binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key, bytes expectedValue) returns()
func (_ILightClient *ILightClientSession) VerifyMembership(proof Ics23Proof, key []byte, expectedValue []byte) (*types.Transaction, error) {
	return _ILightClient.Contract.VerifyMembership(&_ILightClient.TransactOpts, proof, key, expectedValue)
}

// VerifyMembership is a paid mutator transaction binding the contract method 0xcb535ab5.
//
// Solidity: function verifyMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key, bytes expectedValue) returns()
func (_ILightClient *ILightClientTransactorSession) VerifyMembership(proof Ics23Proof, key []byte, expectedValue []byte) (*types.Transaction, error) {
	return _ILightClient.Contract.VerifyMembership(&_ILightClient.TransactOpts, proof, key, expectedValue)
}

// VerifyNonMembership is a paid mutator transaction binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key) returns()
func (_ILightClient *ILightClientTransactor) VerifyNonMembership(opts *bind.TransactOpts, proof Ics23Proof, key []byte) (*types.Transaction, error) {
	return _ILightClient.contract.Transact(opts, "verifyNonMembership", proof, key)
}

// VerifyNonMembership is a paid mutator transaction binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key) returns()
func (_ILightClient *ILightClientSession) VerifyNonMembership(proof Ics23Proof, key []byte) (*types.Transaction, error) {
	return _ILightClient.Contract.VerifyNonMembership(&_ILightClient.TransactOpts, proof, key)
}

// VerifyNonMembership is a paid mutator transaction binding the contract method 0xfdaab4e5.
//
// Solidity: function verifyNonMembership((((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof, bytes key) returns()
func (_ILightClient *ILightClientTransactorSession) VerifyNonMembership(proof Ics23Proof, key []byte) (*types.Transaction, error) {
	return _ILightClient.Contract.VerifyNonMembership(&_ILightClient.TransactOpts, proof, key)
}
