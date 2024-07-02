// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcdispatcher

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

// IbcDispatcherMetaData contains all meta data concerning the IbcDispatcher contract.
var IbcDispatcherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"channelCloseConfirm\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelCloseInit\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenInit\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeVault\",\"inputs\":[],\"outputs\":[{\"name\":\"feeVault\",\"type\":\"address\",\"internalType\":\"contractIFeeVault\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"portPrefix\",\"inputs\":[],\"outputs\":[{\"name\":\"portPrefix\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"}]",
}

// IbcDispatcherABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcDispatcherMetaData.ABI instead.
var IbcDispatcherABI = IbcDispatcherMetaData.ABI

// IbcDispatcher is an auto generated Go binding around an Ethereum contract.
type IbcDispatcher struct {
	IbcDispatcherCaller     // Read-only binding to the contract
	IbcDispatcherTransactor // Write-only binding to the contract
	IbcDispatcherFilterer   // Log filterer for contract events
}

// IbcDispatcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcDispatcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcDispatcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcDispatcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcDispatcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcDispatcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcDispatcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcDispatcherSession struct {
	Contract     *IbcDispatcher    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcDispatcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcDispatcherCallerSession struct {
	Contract *IbcDispatcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IbcDispatcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcDispatcherTransactorSession struct {
	Contract     *IbcDispatcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IbcDispatcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcDispatcherRaw struct {
	Contract *IbcDispatcher // Generic contract binding to access the raw methods on
}

// IbcDispatcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcDispatcherCallerRaw struct {
	Contract *IbcDispatcherCaller // Generic read-only contract binding to access the raw methods on
}

// IbcDispatcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcDispatcherTransactorRaw struct {
	Contract *IbcDispatcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcDispatcher creates a new instance of IbcDispatcher, bound to a specific deployed contract.
func NewIbcDispatcher(address common.Address, backend bind.ContractBackend) (*IbcDispatcher, error) {
	contract, err := bindIbcDispatcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcDispatcher{IbcDispatcherCaller: IbcDispatcherCaller{contract: contract}, IbcDispatcherTransactor: IbcDispatcherTransactor{contract: contract}, IbcDispatcherFilterer: IbcDispatcherFilterer{contract: contract}}, nil
}

// NewIbcDispatcherCaller creates a new read-only instance of IbcDispatcher, bound to a specific deployed contract.
func NewIbcDispatcherCaller(address common.Address, caller bind.ContractCaller) (*IbcDispatcherCaller, error) {
	contract, err := bindIbcDispatcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcDispatcherCaller{contract: contract}, nil
}

// NewIbcDispatcherTransactor creates a new write-only instance of IbcDispatcher, bound to a specific deployed contract.
func NewIbcDispatcherTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcDispatcherTransactor, error) {
	contract, err := bindIbcDispatcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcDispatcherTransactor{contract: contract}, nil
}

// NewIbcDispatcherFilterer creates a new log filterer instance of IbcDispatcher, bound to a specific deployed contract.
func NewIbcDispatcherFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcDispatcherFilterer, error) {
	contract, err := bindIbcDispatcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcDispatcherFilterer{contract: contract}, nil
}

// bindIbcDispatcher binds a generic wrapper to an already deployed contract.
func bindIbcDispatcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcDispatcherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcDispatcher *IbcDispatcherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcDispatcher.Contract.IbcDispatcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcDispatcher *IbcDispatcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcDispatcher.Contract.IbcDispatcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcDispatcher *IbcDispatcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcDispatcher.Contract.IbcDispatcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcDispatcher *IbcDispatcherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcDispatcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcDispatcher *IbcDispatcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcDispatcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcDispatcher *IbcDispatcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcDispatcher.Contract.contract.Transact(opts, method, params...)
}

// PortPrefix is a free data retrieval call binding the contract method 0x7774a6d3.
//
// Solidity: function portPrefix() view returns(string portPrefix)
func (_IbcDispatcher *IbcDispatcherCaller) PortPrefix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IbcDispatcher.contract.Call(opts, &out, "portPrefix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PortPrefix is a free data retrieval call binding the contract method 0x7774a6d3.
//
// Solidity: function portPrefix() view returns(string portPrefix)
func (_IbcDispatcher *IbcDispatcherSession) PortPrefix() (string, error) {
	return _IbcDispatcher.Contract.PortPrefix(&_IbcDispatcher.CallOpts)
}

// PortPrefix is a free data retrieval call binding the contract method 0x7774a6d3.
//
// Solidity: function portPrefix() view returns(string portPrefix)
func (_IbcDispatcher *IbcDispatcherCallerSession) PortPrefix() (string, error) {
	return _IbcDispatcher.Contract.PortPrefix(&_IbcDispatcher.CallOpts)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0xf90b8e96.
//
// Solidity: function channelCloseConfirm(address portAddress, bytes32 channelId, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IbcDispatcher *IbcDispatcherTransactor) ChannelCloseConfirm(opts *bind.TransactOpts, portAddress common.Address, channelId [32]byte, proof Ics23Proof) (*types.Transaction, error) {
	return _IbcDispatcher.contract.Transact(opts, "channelCloseConfirm", portAddress, channelId, proof)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0xf90b8e96.
//
// Solidity: function channelCloseConfirm(address portAddress, bytes32 channelId, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IbcDispatcher *IbcDispatcherSession) ChannelCloseConfirm(portAddress common.Address, channelId [32]byte, proof Ics23Proof) (*types.Transaction, error) {
	return _IbcDispatcher.Contract.ChannelCloseConfirm(&_IbcDispatcher.TransactOpts, portAddress, channelId, proof)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0xf90b8e96.
//
// Solidity: function channelCloseConfirm(address portAddress, bytes32 channelId, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_IbcDispatcher *IbcDispatcherTransactorSession) ChannelCloseConfirm(portAddress common.Address, channelId [32]byte, proof Ics23Proof) (*types.Transaction, error) {
	return _IbcDispatcher.Contract.ChannelCloseConfirm(&_IbcDispatcher.TransactOpts, portAddress, channelId, proof)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0x81bc079b.
//
// Solidity: function channelCloseInit(bytes32 channelId) returns()
func (_IbcDispatcher *IbcDispatcherTransactor) ChannelCloseInit(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _IbcDispatcher.contract.Transact(opts, "channelCloseInit", channelId)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0x81bc079b.
//
// Solidity: function channelCloseInit(bytes32 channelId) returns()
func (_IbcDispatcher *IbcDispatcherSession) ChannelCloseInit(channelId [32]byte) (*types.Transaction, error) {
	return _IbcDispatcher.Contract.ChannelCloseInit(&_IbcDispatcher.TransactOpts, channelId)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0x81bc079b.
//
// Solidity: function channelCloseInit(bytes32 channelId) returns()
func (_IbcDispatcher *IbcDispatcherTransactorSession) ChannelCloseInit(channelId [32]byte) (*types.Transaction, error) {
	return _IbcDispatcher.Contract.ChannelCloseInit(&_IbcDispatcher.TransactOpts, channelId)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0x418925b7.
//
// Solidity: function channelOpenInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_IbcDispatcher *IbcDispatcherTransactor) ChannelOpenInit(opts *bind.TransactOpts, version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _IbcDispatcher.contract.Transact(opts, "channelOpenInit", version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0x418925b7.
//
// Solidity: function channelOpenInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_IbcDispatcher *IbcDispatcherSession) ChannelOpenInit(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _IbcDispatcher.Contract.ChannelOpenInit(&_IbcDispatcher.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0x418925b7.
//
// Solidity: function channelOpenInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_IbcDispatcher *IbcDispatcherTransactorSession) ChannelOpenInit(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _IbcDispatcher.Contract.ChannelOpenInit(&_IbcDispatcher.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// FeeVault is a paid mutator transaction binding the contract method 0x478222c2.
//
// Solidity: function feeVault() returns(address feeVault)
func (_IbcDispatcher *IbcDispatcherTransactor) FeeVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcDispatcher.contract.Transact(opts, "feeVault")
}

// FeeVault is a paid mutator transaction binding the contract method 0x478222c2.
//
// Solidity: function feeVault() returns(address feeVault)
func (_IbcDispatcher *IbcDispatcherSession) FeeVault() (*types.Transaction, error) {
	return _IbcDispatcher.Contract.FeeVault(&_IbcDispatcher.TransactOpts)
}

// FeeVault is a paid mutator transaction binding the contract method 0x478222c2.
//
// Solidity: function feeVault() returns(address feeVault)
func (_IbcDispatcher *IbcDispatcherTransactorSession) FeeVault() (*types.Transaction, error) {
	return _IbcDispatcher.Contract.FeeVault(&_IbcDispatcher.TransactOpts)
}

// SendPacket is a paid mutator transaction binding the contract method 0xc3e1155c.
//
// Solidity: function sendPacket(bytes32 channelId, bytes payload, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcDispatcher *IbcDispatcherTransactor) SendPacket(opts *bind.TransactOpts, channelId [32]byte, payload []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcDispatcher.contract.Transact(opts, "sendPacket", channelId, payload, timeoutTimestamp)
}

// SendPacket is a paid mutator transaction binding the contract method 0xc3e1155c.
//
// Solidity: function sendPacket(bytes32 channelId, bytes payload, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcDispatcher *IbcDispatcherSession) SendPacket(channelId [32]byte, payload []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcDispatcher.Contract.SendPacket(&_IbcDispatcher.TransactOpts, channelId, payload, timeoutTimestamp)
}

// SendPacket is a paid mutator transaction binding the contract method 0xc3e1155c.
//
// Solidity: function sendPacket(bytes32 channelId, bytes payload, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_IbcDispatcher *IbcDispatcherTransactorSession) SendPacket(channelId [32]byte, payload []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _IbcDispatcher.Contract.SendPacket(&_IbcDispatcher.TransactOpts, channelId, payload, timeoutTimestamp)
}
