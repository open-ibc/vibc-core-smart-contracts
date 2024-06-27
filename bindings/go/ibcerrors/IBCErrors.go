// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcerrors

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

// IBCErrorsMetaData contains all meta data concerning the IBCErrors contract.
var IBCErrorsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"ackPacketCommitmentAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"channelIdNotFound\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"channelNotOwnedByPortAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"channelNotOwnedBySender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"clientAlreadyCreated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"clientNotCreated\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"consensusStateVerificationFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidChannelType\",\"inputs\":[{\"name\":\"channelType\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"invalidCharacter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidConnection\",\"inputs\":[{\"name\":\"connection\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"invalidConnectionHops\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidCounterParty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidCounterPartyPortId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidHexStringLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidPacket\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidPacketSequence\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidPortPrefix\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidRelayerAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"lightClientNotFound\",\"inputs\":[{\"name\":\"connection\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"notEnoughGas\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"packetCommitmentNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"packetNotTimedOut\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"packetReceiptAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"receiverNotIntendedPacketDestination\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"receiverNotOriginPacketSender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"unexpectedPacketSequence\",\"inputs\":[]}]",
}

// IBCErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IBCErrorsMetaData.ABI instead.
var IBCErrorsABI = IBCErrorsMetaData.ABI

// IBCErrors is an auto generated Go binding around an Ethereum contract.
type IBCErrors struct {
	IBCErrorsCaller     // Read-only binding to the contract
	IBCErrorsTransactor // Write-only binding to the contract
	IBCErrorsFilterer   // Log filterer for contract events
}

// IBCErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBCErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBCErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBCErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBCErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBCErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBCErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBCErrorsSession struct {
	Contract     *IBCErrors        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBCErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBCErrorsCallerSession struct {
	Contract *IBCErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IBCErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBCErrorsTransactorSession struct {
	Contract     *IBCErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IBCErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBCErrorsRaw struct {
	Contract *IBCErrors // Generic contract binding to access the raw methods on
}

// IBCErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBCErrorsCallerRaw struct {
	Contract *IBCErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IBCErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBCErrorsTransactorRaw struct {
	Contract *IBCErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBCErrors creates a new instance of IBCErrors, bound to a specific deployed contract.
func NewIBCErrors(address common.Address, backend bind.ContractBackend) (*IBCErrors, error) {
	contract, err := bindIBCErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBCErrors{IBCErrorsCaller: IBCErrorsCaller{contract: contract}, IBCErrorsTransactor: IBCErrorsTransactor{contract: contract}, IBCErrorsFilterer: IBCErrorsFilterer{contract: contract}}, nil
}

// NewIBCErrorsCaller creates a new read-only instance of IBCErrors, bound to a specific deployed contract.
func NewIBCErrorsCaller(address common.Address, caller bind.ContractCaller) (*IBCErrorsCaller, error) {
	contract, err := bindIBCErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBCErrorsCaller{contract: contract}, nil
}

// NewIBCErrorsTransactor creates a new write-only instance of IBCErrors, bound to a specific deployed contract.
func NewIBCErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IBCErrorsTransactor, error) {
	contract, err := bindIBCErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBCErrorsTransactor{contract: contract}, nil
}

// NewIBCErrorsFilterer creates a new log filterer instance of IBCErrors, bound to a specific deployed contract.
func NewIBCErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IBCErrorsFilterer, error) {
	contract, err := bindIBCErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBCErrorsFilterer{contract: contract}, nil
}

// bindIBCErrors binds a generic wrapper to an already deployed contract.
func bindIBCErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBCErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBCErrors *IBCErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBCErrors.Contract.IBCErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBCErrors *IBCErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBCErrors.Contract.IBCErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBCErrors *IBCErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBCErrors.Contract.IBCErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBCErrors *IBCErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBCErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBCErrors *IBCErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBCErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBCErrors *IBCErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBCErrors.Contract.contract.Transact(opts, method, params...)
}
