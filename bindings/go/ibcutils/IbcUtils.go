// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibcutils

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

// UniversalPacket is an auto generated low-level Go binding around an user-defined struct.
type UniversalPacket struct {
	SrcPortAddr  [32]byte
	MwBitmap     *big.Int
	DestPortAddr [32]byte
	AppData      []byte
}

// IbcUtilsMetaData contains all meta data concerning the IbcUtils contract.
var IbcUtilsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"fromUniversalPacketBytes\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"universalPacketData\",\"type\":\"tuple\",\"internalType\":\"structUniversalPacket\",\"components\":[{\"name\":\"srcPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"mwBitmap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"hexStrToAddress\",\"inputs\":[{\"name\":\"hexStr\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidCharacter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidHexStringLength\",\"inputs\":[]}]",
}

// IbcUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use IbcUtilsMetaData.ABI instead.
var IbcUtilsABI = IbcUtilsMetaData.ABI

// IbcUtils is an auto generated Go binding around an Ethereum contract.
type IbcUtils struct {
	IbcUtilsCaller     // Read-only binding to the contract
	IbcUtilsTransactor // Write-only binding to the contract
	IbcUtilsFilterer   // Log filterer for contract events
}

// IbcUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IbcUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IbcUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IbcUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IbcUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IbcUtilsSession struct {
	Contract     *IbcUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IbcUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IbcUtilsCallerSession struct {
	Contract *IbcUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IbcUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IbcUtilsTransactorSession struct {
	Contract     *IbcUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IbcUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IbcUtilsRaw struct {
	Contract *IbcUtils // Generic contract binding to access the raw methods on
}

// IbcUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IbcUtilsCallerRaw struct {
	Contract *IbcUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// IbcUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IbcUtilsTransactorRaw struct {
	Contract *IbcUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIbcUtils creates a new instance of IbcUtils, bound to a specific deployed contract.
func NewIbcUtils(address common.Address, backend bind.ContractBackend) (*IbcUtils, error) {
	contract, err := bindIbcUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IbcUtils{IbcUtilsCaller: IbcUtilsCaller{contract: contract}, IbcUtilsTransactor: IbcUtilsTransactor{contract: contract}, IbcUtilsFilterer: IbcUtilsFilterer{contract: contract}}, nil
}

// NewIbcUtilsCaller creates a new read-only instance of IbcUtils, bound to a specific deployed contract.
func NewIbcUtilsCaller(address common.Address, caller bind.ContractCaller) (*IbcUtilsCaller, error) {
	contract, err := bindIbcUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IbcUtilsCaller{contract: contract}, nil
}

// NewIbcUtilsTransactor creates a new write-only instance of IbcUtils, bound to a specific deployed contract.
func NewIbcUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*IbcUtilsTransactor, error) {
	contract, err := bindIbcUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IbcUtilsTransactor{contract: contract}, nil
}

// NewIbcUtilsFilterer creates a new log filterer instance of IbcUtils, bound to a specific deployed contract.
func NewIbcUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*IbcUtilsFilterer, error) {
	contract, err := bindIbcUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IbcUtilsFilterer{contract: contract}, nil
}

// bindIbcUtils binds a generic wrapper to an already deployed contract.
func bindIbcUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IbcUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcUtils *IbcUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcUtils.Contract.IbcUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcUtils *IbcUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcUtils.Contract.IbcUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcUtils *IbcUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcUtils.Contract.IbcUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IbcUtils *IbcUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IbcUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IbcUtils *IbcUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IbcUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IbcUtils *IbcUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IbcUtils.Contract.contract.Transact(opts, method, params...)
}

// FromUniversalPacketBytes is a free data retrieval call binding the contract method 0xd5c39a9d.
//
// Solidity: function fromUniversalPacketBytes(bytes data) pure returns((bytes32,uint256,bytes32,bytes) universalPacketData)
func (_IbcUtils *IbcUtilsCaller) FromUniversalPacketBytes(opts *bind.CallOpts, data []byte) (UniversalPacket, error) {
	var out []interface{}
	err := _IbcUtils.contract.Call(opts, &out, "fromUniversalPacketBytes", data)

	if err != nil {
		return *new(UniversalPacket), err
	}

	out0 := *abi.ConvertType(out[0], new(UniversalPacket)).(*UniversalPacket)

	return out0, err

}

// FromUniversalPacketBytes is a free data retrieval call binding the contract method 0xd5c39a9d.
//
// Solidity: function fromUniversalPacketBytes(bytes data) pure returns((bytes32,uint256,bytes32,bytes) universalPacketData)
func (_IbcUtils *IbcUtilsSession) FromUniversalPacketBytes(data []byte) (UniversalPacket, error) {
	return _IbcUtils.Contract.FromUniversalPacketBytes(&_IbcUtils.CallOpts, data)
}

// FromUniversalPacketBytes is a free data retrieval call binding the contract method 0xd5c39a9d.
//
// Solidity: function fromUniversalPacketBytes(bytes data) pure returns((bytes32,uint256,bytes32,bytes) universalPacketData)
func (_IbcUtils *IbcUtilsCallerSession) FromUniversalPacketBytes(data []byte) (UniversalPacket, error) {
	return _IbcUtils.Contract.FromUniversalPacketBytes(&_IbcUtils.CallOpts, data)
}

// HexStrToAddress is a free data retrieval call binding the contract method 0xa1ef9a98.
//
// Solidity: function hexStrToAddress(string hexStr) pure returns(address addr)
func (_IbcUtils *IbcUtilsCaller) HexStrToAddress(opts *bind.CallOpts, hexStr string) (common.Address, error) {
	var out []interface{}
	err := _IbcUtils.contract.Call(opts, &out, "hexStrToAddress", hexStr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HexStrToAddress is a free data retrieval call binding the contract method 0xa1ef9a98.
//
// Solidity: function hexStrToAddress(string hexStr) pure returns(address addr)
func (_IbcUtils *IbcUtilsSession) HexStrToAddress(hexStr string) (common.Address, error) {
	return _IbcUtils.Contract.HexStrToAddress(&_IbcUtils.CallOpts, hexStr)
}

// HexStrToAddress is a free data retrieval call binding the contract method 0xa1ef9a98.
//
// Solidity: function hexStrToAddress(string hexStr) pure returns(address addr)
func (_IbcUtils *IbcUtilsCallerSession) HexStrToAddress(hexStr string) (common.Address, error) {
	return _IbcUtils.Contract.HexStrToAddress(&_IbcUtils.CallOpts, hexStr)
}
