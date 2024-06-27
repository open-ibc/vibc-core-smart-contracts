// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package universalchannelhandler

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

// AckPacket is an auto generated low-level Go binding around an user-defined struct.
type AckPacket struct {
	Success bool
	Data    []byte
}

// Height is an auto generated low-level Go binding around an user-defined struct.
type Height struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IbcEndpoint is an auto generated low-level Go binding around an user-defined struct.
type IbcEndpoint struct {
	PortId    string
	ChannelId [32]byte
}

// IbcPacket is an auto generated low-level Go binding around an user-defined struct.
type IbcPacket struct {
	Src              IbcEndpoint
	Dest             IbcEndpoint
	Sequence         uint64
	Data             []byte
	TimeoutHeight    Height
	TimeoutTimestamp uint64
}

// UniversalchannelhandlerMetaData contains all meta data concerning the Universalchannelhandler contract.
var UniversalchannelhandlerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MW_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"closeChannel\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dispatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_dispatcher\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseInit\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onChanOpenConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenInit\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onChanOpenTry\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"openChannel\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortIdentifier\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacketWithFee\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setDispatcher\",\"inputs\":[{\"name\":\"_dispatcher\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UCHPacketSent\",\"inputs\":[{\"name\":\"source\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ChannelNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MwBitmpaCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"notIbcDispatcher\",\"inputs\":[]}]",
}

// UniversalchannelhandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use UniversalchannelhandlerMetaData.ABI instead.
var UniversalchannelhandlerABI = UniversalchannelhandlerMetaData.ABI

// Universalchannelhandler is an auto generated Go binding around an Ethereum contract.
type Universalchannelhandler struct {
	UniversalchannelhandlerCaller     // Read-only binding to the contract
	UniversalchannelhandlerTransactor // Write-only binding to the contract
	UniversalchannelhandlerFilterer   // Log filterer for contract events
}

// UniversalchannelhandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniversalchannelhandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalchannelhandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniversalchannelhandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalchannelhandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniversalchannelhandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalchannelhandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniversalchannelhandlerSession struct {
	Contract     *Universalchannelhandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UniversalchannelhandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniversalchannelhandlerCallerSession struct {
	Contract *UniversalchannelhandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// UniversalchannelhandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniversalchannelhandlerTransactorSession struct {
	Contract     *UniversalchannelhandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// UniversalchannelhandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniversalchannelhandlerRaw struct {
	Contract *Universalchannelhandler // Generic contract binding to access the raw methods on
}

// UniversalchannelhandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniversalchannelhandlerCallerRaw struct {
	Contract *UniversalchannelhandlerCaller // Generic read-only contract binding to access the raw methods on
}

// UniversalchannelhandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniversalchannelhandlerTransactorRaw struct {
	Contract *UniversalchannelhandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversalchannelhandler creates a new instance of Universalchannelhandler, bound to a specific deployed contract.
func NewUniversalchannelhandler(address common.Address, backend bind.ContractBackend) (*Universalchannelhandler, error) {
	contract, err := bindUniversalchannelhandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Universalchannelhandler{UniversalchannelhandlerCaller: UniversalchannelhandlerCaller{contract: contract}, UniversalchannelhandlerTransactor: UniversalchannelhandlerTransactor{contract: contract}, UniversalchannelhandlerFilterer: UniversalchannelhandlerFilterer{contract: contract}}, nil
}

// NewUniversalchannelhandlerCaller creates a new read-only instance of Universalchannelhandler, bound to a specific deployed contract.
func NewUniversalchannelhandlerCaller(address common.Address, caller bind.ContractCaller) (*UniversalchannelhandlerCaller, error) {
	contract, err := bindUniversalchannelhandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalchannelhandlerCaller{contract: contract}, nil
}

// NewUniversalchannelhandlerTransactor creates a new write-only instance of Universalchannelhandler, bound to a specific deployed contract.
func NewUniversalchannelhandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversalchannelhandlerTransactor, error) {
	contract, err := bindUniversalchannelhandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalchannelhandlerTransactor{contract: contract}, nil
}

// NewUniversalchannelhandlerFilterer creates a new log filterer instance of Universalchannelhandler, bound to a specific deployed contract.
func NewUniversalchannelhandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversalchannelhandlerFilterer, error) {
	contract, err := bindUniversalchannelhandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversalchannelhandlerFilterer{contract: contract}, nil
}

// bindUniversalchannelhandler binds a generic wrapper to an already deployed contract.
func bindUniversalchannelhandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniversalchannelhandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Universalchannelhandler *UniversalchannelhandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Universalchannelhandler.Contract.UniversalchannelhandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Universalchannelhandler *UniversalchannelhandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.UniversalchannelhandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Universalchannelhandler *UniversalchannelhandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.UniversalchannelhandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Universalchannelhandler *UniversalchannelhandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Universalchannelhandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Universalchannelhandler *UniversalchannelhandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Universalchannelhandler *UniversalchannelhandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.contract.Transact(opts, method, params...)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256)
func (_Universalchannelhandler *UniversalchannelhandlerCaller) MWID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Universalchannelhandler.contract.Call(opts, &out, "MW_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256)
func (_Universalchannelhandler *UniversalchannelhandlerSession) MWID() (*big.Int, error) {
	return _Universalchannelhandler.Contract.MWID(&_Universalchannelhandler.CallOpts)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256)
func (_Universalchannelhandler *UniversalchannelhandlerCallerSession) MWID() (*big.Int, error) {
	return _Universalchannelhandler.Contract.MWID(&_Universalchannelhandler.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Universalchannelhandler *UniversalchannelhandlerCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Universalchannelhandler.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Universalchannelhandler *UniversalchannelhandlerSession) VERSION() (string, error) {
	return _Universalchannelhandler.Contract.VERSION(&_Universalchannelhandler.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Universalchannelhandler *UniversalchannelhandlerCallerSession) VERSION() (string, error) {
	return _Universalchannelhandler.Contract.VERSION(&_Universalchannelhandler.CallOpts)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_Universalchannelhandler *UniversalchannelhandlerCaller) Dispatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Universalchannelhandler.contract.Call(opts, &out, "dispatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_Universalchannelhandler *UniversalchannelhandlerSession) Dispatcher() (common.Address, error) {
	return _Universalchannelhandler.Contract.Dispatcher(&_Universalchannelhandler.CallOpts)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_Universalchannelhandler *UniversalchannelhandlerCallerSession) Dispatcher() (common.Address, error) {
	return _Universalchannelhandler.Contract.Dispatcher(&_Universalchannelhandler.CallOpts)
}

// OnChanOpenAck is a free data retrieval call binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) view returns()
func (_Universalchannelhandler *UniversalchannelhandlerCaller) OnChanOpenAck(opts *bind.CallOpts, channelId [32]byte, arg1 [32]byte, counterpartyVersion string) error {
	var out []interface{}
	err := _Universalchannelhandler.contract.Call(opts, &out, "onChanOpenAck", channelId, arg1, counterpartyVersion)

	if err != nil {
		return err
	}

	return err

}

// OnChanOpenAck is a free data retrieval call binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) view returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) OnChanOpenAck(channelId [32]byte, arg1 [32]byte, counterpartyVersion string) error {
	return _Universalchannelhandler.Contract.OnChanOpenAck(&_Universalchannelhandler.CallOpts, channelId, arg1, counterpartyVersion)
}

// OnChanOpenAck is a free data retrieval call binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) view returns()
func (_Universalchannelhandler *UniversalchannelhandlerCallerSession) OnChanOpenAck(channelId [32]byte, arg1 [32]byte, counterpartyVersion string) error {
	return _Universalchannelhandler.Contract.OnChanOpenAck(&_Universalchannelhandler.CallOpts, channelId, arg1, counterpartyVersion)
}

// OnChanOpenInit is a free data retrieval call binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 , string[] , string , string version) view returns(string selectedVersion)
func (_Universalchannelhandler *UniversalchannelhandlerCaller) OnChanOpenInit(opts *bind.CallOpts, arg0 uint8, arg1 []string, arg2 string, version string) (string, error) {
	var out []interface{}
	err := _Universalchannelhandler.contract.Call(opts, &out, "onChanOpenInit", arg0, arg1, arg2, version)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OnChanOpenInit is a free data retrieval call binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 , string[] , string , string version) view returns(string selectedVersion)
func (_Universalchannelhandler *UniversalchannelhandlerSession) OnChanOpenInit(arg0 uint8, arg1 []string, arg2 string, version string) (string, error) {
	return _Universalchannelhandler.Contract.OnChanOpenInit(&_Universalchannelhandler.CallOpts, arg0, arg1, arg2, version)
}

// OnChanOpenInit is a free data retrieval call binding the contract method 0x7a9ccc4b.
//
// Solidity: function onChanOpenInit(uint8 , string[] , string , string version) view returns(string selectedVersion)
func (_Universalchannelhandler *UniversalchannelhandlerCallerSession) OnChanOpenInit(arg0 uint8, arg1 []string, arg2 string, version string) (string, error) {
	return _Universalchannelhandler.Contract.OnChanOpenInit(&_Universalchannelhandler.CallOpts, arg0, arg1, arg2, version)
}

// OnChanOpenTry is a free data retrieval call binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 , string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) view returns(string selectedVersion)
func (_Universalchannelhandler *UniversalchannelhandlerCaller) OnChanOpenTry(opts *bind.CallOpts, arg0 uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (string, error) {
	var out []interface{}
	err := _Universalchannelhandler.contract.Call(opts, &out, "onChanOpenTry", arg0, arg1, channelId, arg3, arg4, counterpartyVersion)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OnChanOpenTry is a free data retrieval call binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 , string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) view returns(string selectedVersion)
func (_Universalchannelhandler *UniversalchannelhandlerSession) OnChanOpenTry(arg0 uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (string, error) {
	return _Universalchannelhandler.Contract.OnChanOpenTry(&_Universalchannelhandler.CallOpts, arg0, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// OnChanOpenTry is a free data retrieval call binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 , string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) view returns(string selectedVersion)
func (_Universalchannelhandler *UniversalchannelhandlerCallerSession) OnChanOpenTry(arg0 uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (string, error) {
	return _Universalchannelhandler.Contract.OnChanOpenTry(&_Universalchannelhandler.CallOpts, arg0, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Universalchannelhandler *UniversalchannelhandlerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Universalchannelhandler.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Universalchannelhandler *UniversalchannelhandlerSession) Owner() (common.Address, error) {
	return _Universalchannelhandler.Contract.Owner(&_Universalchannelhandler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Universalchannelhandler *UniversalchannelhandlerCallerSession) Owner() (common.Address, error) {
	return _Universalchannelhandler.Contract.Owner(&_Universalchannelhandler.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Universalchannelhandler *UniversalchannelhandlerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Universalchannelhandler.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Universalchannelhandler *UniversalchannelhandlerSession) ProxiableUUID() ([32]byte, error) {
	return _Universalchannelhandler.Contract.ProxiableUUID(&_Universalchannelhandler.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Universalchannelhandler *UniversalchannelhandlerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Universalchannelhandler.Contract.ProxiableUUID(&_Universalchannelhandler.CallOpts)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x4c2ee09d.
//
// Solidity: function closeChannel(bytes32 channelId) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) CloseChannel(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "closeChannel", channelId)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x4c2ee09d.
//
// Solidity: function closeChannel(bytes32 channelId) returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) CloseChannel(channelId [32]byte) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.CloseChannel(&_Universalchannelhandler.TransactOpts, channelId)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x4c2ee09d.
//
// Solidity: function closeChannel(bytes32 channelId) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) CloseChannel(channelId [32]byte) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.CloseChannel(&_Universalchannelhandler.TransactOpts, channelId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dispatcher) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) Initialize(opts *bind.TransactOpts, _dispatcher common.Address) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "initialize", _dispatcher)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dispatcher) returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) Initialize(_dispatcher common.Address) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.Initialize(&_Universalchannelhandler.TransactOpts, _dispatcher)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dispatcher) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) Initialize(_dispatcher common.Address) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.Initialize(&_Universalchannelhandler.TransactOpts, _dispatcher)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "onAcknowledgementPacket", packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OnAcknowledgementPacket(&_Universalchannelhandler.TransactOpts, packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OnAcknowledgementPacket(&_Universalchannelhandler.TransactOpts, packet, ack)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string , bytes32 ) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) OnChanCloseConfirm(opts *bind.TransactOpts, channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "onChanCloseConfirm", channelId, arg1, arg2)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string , bytes32 ) returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) OnChanCloseConfirm(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OnChanCloseConfirm(&_Universalchannelhandler.TransactOpts, channelId, arg1, arg2)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string , bytes32 ) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) OnChanCloseConfirm(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OnChanCloseConfirm(&_Universalchannelhandler.TransactOpts, channelId, arg1, arg2)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string , bytes32 ) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) OnChanCloseInit(opts *bind.TransactOpts, channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "onChanCloseInit", channelId, arg1, arg2)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string , bytes32 ) returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) OnChanCloseInit(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OnChanCloseInit(&_Universalchannelhandler.TransactOpts, channelId, arg1, arg2)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string , bytes32 ) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) OnChanCloseInit(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OnChanCloseInit(&_Universalchannelhandler.TransactOpts, channelId, arg1, arg2)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "onChanOpenConfirm", channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OnChanOpenConfirm(&_Universalchannelhandler.TransactOpts, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OnChanOpenConfirm(&_Universalchannelhandler.TransactOpts, channelId)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) OnRecvPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "onRecvPacket", packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_Universalchannelhandler *UniversalchannelhandlerSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OnRecvPacket(&_Universalchannelhandler.TransactOpts, packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket)
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OnRecvPacket(&_Universalchannelhandler.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) OnTimeoutPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "onTimeoutPacket", packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OnTimeoutPacket(&_Universalchannelhandler.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OnTimeoutPacket(&_Universalchannelhandler.TransactOpts, packet)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xace02de7.
//
// Solidity: function openChannel(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortIdentifier) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) OpenChannel(opts *bind.TransactOpts, version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortIdentifier string) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "openChannel", version, ordering, feeEnabled, connectionHops, counterpartyPortIdentifier)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xace02de7.
//
// Solidity: function openChannel(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortIdentifier) returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) OpenChannel(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortIdentifier string) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OpenChannel(&_Universalchannelhandler.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortIdentifier)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xace02de7.
//
// Solidity: function openChannel(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortIdentifier) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) OpenChannel(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortIdentifier string) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.OpenChannel(&_Universalchannelhandler.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortIdentifier)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.RenounceOwnership(&_Universalchannelhandler.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.RenounceOwnership(&_Universalchannelhandler.TransactOpts)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) SendUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "sendUniversalPacket", channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_Universalchannelhandler *UniversalchannelhandlerSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.SendUniversalPacket(&_Universalchannelhandler.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.SendUniversalPacket(&_Universalchannelhandler.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) SendUniversalPacketWithFee(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "sendUniversalPacketWithFee", channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_Universalchannelhandler *UniversalchannelhandlerSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.SendUniversalPacketWithFee(&_Universalchannelhandler.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.SendUniversalPacketWithFee(&_Universalchannelhandler.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address _dispatcher) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) SetDispatcher(opts *bind.TransactOpts, _dispatcher common.Address) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "setDispatcher", _dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address _dispatcher) returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) SetDispatcher(_dispatcher common.Address) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.SetDispatcher(&_Universalchannelhandler.TransactOpts, _dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address _dispatcher) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) SetDispatcher(_dispatcher common.Address) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.SetDispatcher(&_Universalchannelhandler.TransactOpts, _dispatcher)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.TransferOwnership(&_Universalchannelhandler.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.TransferOwnership(&_Universalchannelhandler.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.UpgradeTo(&_Universalchannelhandler.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.UpgradeTo(&_Universalchannelhandler.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.UpgradeToAndCall(&_Universalchannelhandler.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.UpgradeToAndCall(&_Universalchannelhandler.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universalchannelhandler.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Universalchannelhandler *UniversalchannelhandlerSession) Receive() (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.Receive(&_Universalchannelhandler.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Universalchannelhandler *UniversalchannelhandlerTransactorSession) Receive() (*types.Transaction, error) {
	return _Universalchannelhandler.Contract.Receive(&_Universalchannelhandler.TransactOpts)
}

// UniversalchannelhandlerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Universalchannelhandler contract.
type UniversalchannelhandlerAdminChangedIterator struct {
	Event *UniversalchannelhandlerAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniversalchannelhandlerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalchannelhandlerAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniversalchannelhandlerAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniversalchannelhandlerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalchannelhandlerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalchannelhandlerAdminChanged represents a AdminChanged event raised by the Universalchannelhandler contract.
type UniversalchannelhandlerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*UniversalchannelhandlerAdminChangedIterator, error) {

	logs, sub, err := _Universalchannelhandler.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &UniversalchannelhandlerAdminChangedIterator{contract: _Universalchannelhandler.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *UniversalchannelhandlerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Universalchannelhandler.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalchannelhandlerAdminChanged)
				if err := _Universalchannelhandler.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) ParseAdminChanged(log types.Log) (*UniversalchannelhandlerAdminChanged, error) {
	event := new(UniversalchannelhandlerAdminChanged)
	if err := _Universalchannelhandler.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalchannelhandlerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Universalchannelhandler contract.
type UniversalchannelhandlerBeaconUpgradedIterator struct {
	Event *UniversalchannelhandlerBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniversalchannelhandlerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalchannelhandlerBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniversalchannelhandlerBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniversalchannelhandlerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalchannelhandlerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalchannelhandlerBeaconUpgraded represents a BeaconUpgraded event raised by the Universalchannelhandler contract.
type UniversalchannelhandlerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*UniversalchannelhandlerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Universalchannelhandler.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &UniversalchannelhandlerBeaconUpgradedIterator{contract: _Universalchannelhandler.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *UniversalchannelhandlerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Universalchannelhandler.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalchannelhandlerBeaconUpgraded)
				if err := _Universalchannelhandler.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) ParseBeaconUpgraded(log types.Log) (*UniversalchannelhandlerBeaconUpgraded, error) {
	event := new(UniversalchannelhandlerBeaconUpgraded)
	if err := _Universalchannelhandler.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalchannelhandlerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Universalchannelhandler contract.
type UniversalchannelhandlerInitializedIterator struct {
	Event *UniversalchannelhandlerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniversalchannelhandlerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalchannelhandlerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniversalchannelhandlerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniversalchannelhandlerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalchannelhandlerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalchannelhandlerInitialized represents a Initialized event raised by the Universalchannelhandler contract.
type UniversalchannelhandlerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) FilterInitialized(opts *bind.FilterOpts) (*UniversalchannelhandlerInitializedIterator, error) {

	logs, sub, err := _Universalchannelhandler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UniversalchannelhandlerInitializedIterator{contract: _Universalchannelhandler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UniversalchannelhandlerInitialized) (event.Subscription, error) {

	logs, sub, err := _Universalchannelhandler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalchannelhandlerInitialized)
				if err := _Universalchannelhandler.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) ParseInitialized(log types.Log) (*UniversalchannelhandlerInitialized, error) {
	event := new(UniversalchannelhandlerInitialized)
	if err := _Universalchannelhandler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalchannelhandlerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Universalchannelhandler contract.
type UniversalchannelhandlerOwnershipTransferredIterator struct {
	Event *UniversalchannelhandlerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniversalchannelhandlerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalchannelhandlerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniversalchannelhandlerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniversalchannelhandlerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalchannelhandlerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalchannelhandlerOwnershipTransferred represents a OwnershipTransferred event raised by the Universalchannelhandler contract.
type UniversalchannelhandlerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniversalchannelhandlerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Universalchannelhandler.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalchannelhandlerOwnershipTransferredIterator{contract: _Universalchannelhandler.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UniversalchannelhandlerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Universalchannelhandler.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalchannelhandlerOwnershipTransferred)
				if err := _Universalchannelhandler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) ParseOwnershipTransferred(log types.Log) (*UniversalchannelhandlerOwnershipTransferred, error) {
	event := new(UniversalchannelhandlerOwnershipTransferred)
	if err := _Universalchannelhandler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalchannelhandlerUCHPacketSentIterator is returned from FilterUCHPacketSent and is used to iterate over the raw logs and unpacked data for UCHPacketSent events raised by the Universalchannelhandler contract.
type UniversalchannelhandlerUCHPacketSentIterator struct {
	Event *UniversalchannelhandlerUCHPacketSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniversalchannelhandlerUCHPacketSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalchannelhandlerUCHPacketSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniversalchannelhandlerUCHPacketSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniversalchannelhandlerUCHPacketSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalchannelhandlerUCHPacketSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalchannelhandlerUCHPacketSent represents a UCHPacketSent event raised by the Universalchannelhandler contract.
type UniversalchannelhandlerUCHPacketSent struct {
	Source      common.Address
	Destination [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUCHPacketSent is a free log retrieval operation binding the contract event 0x9831d8c66285bfd33de069ced58ad437d6bf08f63446bf06c3713e40b4b7e873.
//
// Solidity: event UCHPacketSent(address source, bytes32 destination)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) FilterUCHPacketSent(opts *bind.FilterOpts) (*UniversalchannelhandlerUCHPacketSentIterator, error) {

	logs, sub, err := _Universalchannelhandler.contract.FilterLogs(opts, "UCHPacketSent")
	if err != nil {
		return nil, err
	}
	return &UniversalchannelhandlerUCHPacketSentIterator{contract: _Universalchannelhandler.contract, event: "UCHPacketSent", logs: logs, sub: sub}, nil
}

// WatchUCHPacketSent is a free log subscription operation binding the contract event 0x9831d8c66285bfd33de069ced58ad437d6bf08f63446bf06c3713e40b4b7e873.
//
// Solidity: event UCHPacketSent(address source, bytes32 destination)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) WatchUCHPacketSent(opts *bind.WatchOpts, sink chan<- *UniversalchannelhandlerUCHPacketSent) (event.Subscription, error) {

	logs, sub, err := _Universalchannelhandler.contract.WatchLogs(opts, "UCHPacketSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalchannelhandlerUCHPacketSent)
				if err := _Universalchannelhandler.contract.UnpackLog(event, "UCHPacketSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUCHPacketSent is a log parse operation binding the contract event 0x9831d8c66285bfd33de069ced58ad437d6bf08f63446bf06c3713e40b4b7e873.
//
// Solidity: event UCHPacketSent(address source, bytes32 destination)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) ParseUCHPacketSent(log types.Log) (*UniversalchannelhandlerUCHPacketSent, error) {
	event := new(UniversalchannelhandlerUCHPacketSent)
	if err := _Universalchannelhandler.contract.UnpackLog(event, "UCHPacketSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalchannelhandlerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Universalchannelhandler contract.
type UniversalchannelhandlerUpgradedIterator struct {
	Event *UniversalchannelhandlerUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UniversalchannelhandlerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalchannelhandlerUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UniversalchannelhandlerUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UniversalchannelhandlerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalchannelhandlerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalchannelhandlerUpgraded represents a Upgraded event raised by the Universalchannelhandler contract.
type UniversalchannelhandlerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UniversalchannelhandlerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Universalchannelhandler.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UniversalchannelhandlerUpgradedIterator{contract: _Universalchannelhandler.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UniversalchannelhandlerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Universalchannelhandler.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalchannelhandlerUpgraded)
				if err := _Universalchannelhandler.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Universalchannelhandler *UniversalchannelhandlerFilterer) ParseUpgraded(log types.Log) (*UniversalchannelhandlerUpgraded, error) {
	event := new(UniversalchannelhandlerUpgraded)
	if err := _Universalchannelhandler.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
