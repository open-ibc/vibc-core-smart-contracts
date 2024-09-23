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

// UniversalChannelHandlerMetaData contains all meta data concerning the UniversalChannelHandler contract.
var UniversalChannelHandlerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MW_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"closeChannel\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dispatcher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_dispatcher\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onAcknowledgementPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"ack\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanCloseInit\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenAck\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onChanOpenConfirm\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onChanOpenTry\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"counterpartyVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"selectedVersion\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onRecvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[{\"name\":\"ackPacket\",\"type\":\"tuple\",\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"skipAck\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onTimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"openChannel\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortIdentifier\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendUniversalPacketWithFee\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destPortAddr\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"appData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimits\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"gasPrices\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setDispatcher\",\"inputs\":[{\"name\":\"_dispatcher\",\"type\":\"address\",\"internalType\":\"contractIbcDispatcher\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UCHPacketSent\",\"inputs\":[{\"name\":\"source\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ChannelNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MwBitmpaCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnsupportedVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"notIbcDispatcher\",\"inputs\":[]}]",
}

// UniversalChannelHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use UniversalChannelHandlerMetaData.ABI instead.
var UniversalChannelHandlerABI = UniversalChannelHandlerMetaData.ABI

// UniversalChannelHandler is an auto generated Go binding around an Ethereum contract.
type UniversalChannelHandler struct {
	UniversalChannelHandlerCaller     // Read-only binding to the contract
	UniversalChannelHandlerTransactor // Write-only binding to the contract
	UniversalChannelHandlerFilterer   // Log filterer for contract events
}

// UniversalChannelHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniversalChannelHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalChannelHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniversalChannelHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalChannelHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniversalChannelHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalChannelHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniversalChannelHandlerSession struct {
	Contract     *UniversalChannelHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UniversalChannelHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniversalChannelHandlerCallerSession struct {
	Contract *UniversalChannelHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// UniversalChannelHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniversalChannelHandlerTransactorSession struct {
	Contract     *UniversalChannelHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// UniversalChannelHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniversalChannelHandlerRaw struct {
	Contract *UniversalChannelHandler // Generic contract binding to access the raw methods on
}

// UniversalChannelHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniversalChannelHandlerCallerRaw struct {
	Contract *UniversalChannelHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// UniversalChannelHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniversalChannelHandlerTransactorRaw struct {
	Contract *UniversalChannelHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversalChannelHandler creates a new instance of UniversalChannelHandler, bound to a specific deployed contract.
func NewUniversalChannelHandler(address common.Address, backend bind.ContractBackend) (*UniversalChannelHandler, error) {
	contract, err := bindUniversalChannelHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniversalChannelHandler{UniversalChannelHandlerCaller: UniversalChannelHandlerCaller{contract: contract}, UniversalChannelHandlerTransactor: UniversalChannelHandlerTransactor{contract: contract}, UniversalChannelHandlerFilterer: UniversalChannelHandlerFilterer{contract: contract}}, nil
}

// NewUniversalChannelHandlerCaller creates a new read-only instance of UniversalChannelHandler, bound to a specific deployed contract.
func NewUniversalChannelHandlerCaller(address common.Address, caller bind.ContractCaller) (*UniversalChannelHandlerCaller, error) {
	contract, err := bindUniversalChannelHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalChannelHandlerCaller{contract: contract}, nil
}

// NewUniversalChannelHandlerTransactor creates a new write-only instance of UniversalChannelHandler, bound to a specific deployed contract.
func NewUniversalChannelHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversalChannelHandlerTransactor, error) {
	contract, err := bindUniversalChannelHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalChannelHandlerTransactor{contract: contract}, nil
}

// NewUniversalChannelHandlerFilterer creates a new log filterer instance of UniversalChannelHandler, bound to a specific deployed contract.
func NewUniversalChannelHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversalChannelHandlerFilterer, error) {
	contract, err := bindUniversalChannelHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversalChannelHandlerFilterer{contract: contract}, nil
}

// bindUniversalChannelHandler binds a generic wrapper to an already deployed contract.
func bindUniversalChannelHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniversalChannelHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalChannelHandler *UniversalChannelHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalChannelHandler.Contract.UniversalChannelHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalChannelHandler *UniversalChannelHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.UniversalChannelHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalChannelHandler *UniversalChannelHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.UniversalChannelHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniversalChannelHandler *UniversalChannelHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniversalChannelHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.contract.Transact(opts, method, params...)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256)
func (_UniversalChannelHandler *UniversalChannelHandlerCaller) MWID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniversalChannelHandler.contract.Call(opts, &out, "MW_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256)
func (_UniversalChannelHandler *UniversalChannelHandlerSession) MWID() (*big.Int, error) {
	return _UniversalChannelHandler.Contract.MWID(&_UniversalChannelHandler.CallOpts)
}

// MWID is a free data retrieval call binding the contract method 0xc1cb44e5.
//
// Solidity: function MW_ID() view returns(uint256)
func (_UniversalChannelHandler *UniversalChannelHandlerCallerSession) MWID() (*big.Int, error) {
	return _UniversalChannelHandler.Contract.MWID(&_UniversalChannelHandler.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UniversalChannelHandler *UniversalChannelHandlerCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _UniversalChannelHandler.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UniversalChannelHandler *UniversalChannelHandlerSession) VERSION() (string, error) {
	return _UniversalChannelHandler.Contract.VERSION(&_UniversalChannelHandler.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_UniversalChannelHandler *UniversalChannelHandlerCallerSession) VERSION() (string, error) {
	return _UniversalChannelHandler.Contract.VERSION(&_UniversalChannelHandler.CallOpts)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_UniversalChannelHandler *UniversalChannelHandlerCaller) Dispatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalChannelHandler.contract.Call(opts, &out, "dispatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_UniversalChannelHandler *UniversalChannelHandlerSession) Dispatcher() (common.Address, error) {
	return _UniversalChannelHandler.Contract.Dispatcher(&_UniversalChannelHandler.CallOpts)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_UniversalChannelHandler *UniversalChannelHandlerCallerSession) Dispatcher() (common.Address, error) {
	return _UniversalChannelHandler.Contract.Dispatcher(&_UniversalChannelHandler.CallOpts)
}

// OnChanOpenAck is a free data retrieval call binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) view returns()
func (_UniversalChannelHandler *UniversalChannelHandlerCaller) OnChanOpenAck(opts *bind.CallOpts, channelId [32]byte, arg1 [32]byte, counterpartyVersion string) error {
	var out []interface{}
	err := _UniversalChannelHandler.contract.Call(opts, &out, "onChanOpenAck", channelId, arg1, counterpartyVersion)

	if err != nil {
		return err
	}

	return err

}

// OnChanOpenAck is a free data retrieval call binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) view returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) OnChanOpenAck(channelId [32]byte, arg1 [32]byte, counterpartyVersion string) error {
	return _UniversalChannelHandler.Contract.OnChanOpenAck(&_UniversalChannelHandler.CallOpts, channelId, arg1, counterpartyVersion)
}

// OnChanOpenAck is a free data retrieval call binding the contract method 0xe847e280.
//
// Solidity: function onChanOpenAck(bytes32 channelId, bytes32 , string counterpartyVersion) view returns()
func (_UniversalChannelHandler *UniversalChannelHandlerCallerSession) OnChanOpenAck(channelId [32]byte, arg1 [32]byte, counterpartyVersion string) error {
	return _UniversalChannelHandler.Contract.OnChanOpenAck(&_UniversalChannelHandler.CallOpts, channelId, arg1, counterpartyVersion)
}

// OnChanOpenTry is a free data retrieval call binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 , string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) view returns(string selectedVersion)
func (_UniversalChannelHandler *UniversalChannelHandlerCaller) OnChanOpenTry(opts *bind.CallOpts, arg0 uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (string, error) {
	var out []interface{}
	err := _UniversalChannelHandler.contract.Call(opts, &out, "onChanOpenTry", arg0, arg1, channelId, arg3, arg4, counterpartyVersion)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OnChanOpenTry is a free data retrieval call binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 , string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) view returns(string selectedVersion)
func (_UniversalChannelHandler *UniversalChannelHandlerSession) OnChanOpenTry(arg0 uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (string, error) {
	return _UniversalChannelHandler.Contract.OnChanOpenTry(&_UniversalChannelHandler.CallOpts, arg0, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// OnChanOpenTry is a free data retrieval call binding the contract method 0x4bdb5597.
//
// Solidity: function onChanOpenTry(uint8 , string[] , bytes32 channelId, string , bytes32 , string counterpartyVersion) view returns(string selectedVersion)
func (_UniversalChannelHandler *UniversalChannelHandlerCallerSession) OnChanOpenTry(arg0 uint8, arg1 []string, channelId [32]byte, arg3 string, arg4 [32]byte, counterpartyVersion string) (string, error) {
	return _UniversalChannelHandler.Contract.OnChanOpenTry(&_UniversalChannelHandler.CallOpts, arg0, arg1, channelId, arg3, arg4, counterpartyVersion)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniversalChannelHandler *UniversalChannelHandlerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalChannelHandler.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniversalChannelHandler *UniversalChannelHandlerSession) Owner() (common.Address, error) {
	return _UniversalChannelHandler.Contract.Owner(&_UniversalChannelHandler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniversalChannelHandler *UniversalChannelHandlerCallerSession) Owner() (common.Address, error) {
	return _UniversalChannelHandler.Contract.Owner(&_UniversalChannelHandler.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniversalChannelHandler *UniversalChannelHandlerCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniversalChannelHandler.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniversalChannelHandler *UniversalChannelHandlerSession) PendingOwner() (common.Address, error) {
	return _UniversalChannelHandler.Contract.PendingOwner(&_UniversalChannelHandler.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UniversalChannelHandler *UniversalChannelHandlerCallerSession) PendingOwner() (common.Address, error) {
	return _UniversalChannelHandler.Contract.PendingOwner(&_UniversalChannelHandler.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UniversalChannelHandler *UniversalChannelHandlerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniversalChannelHandler.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UniversalChannelHandler *UniversalChannelHandlerSession) ProxiableUUID() ([32]byte, error) {
	return _UniversalChannelHandler.Contract.ProxiableUUID(&_UniversalChannelHandler.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UniversalChannelHandler *UniversalChannelHandlerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _UniversalChannelHandler.Contract.ProxiableUUID(&_UniversalChannelHandler.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) AcceptOwnership() (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.AcceptOwnership(&_UniversalChannelHandler.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.AcceptOwnership(&_UniversalChannelHandler.TransactOpts)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x4c2ee09d.
//
// Solidity: function closeChannel(bytes32 channelId) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) CloseChannel(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "closeChannel", channelId)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x4c2ee09d.
//
// Solidity: function closeChannel(bytes32 channelId) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) CloseChannel(channelId [32]byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.CloseChannel(&_UniversalChannelHandler.TransactOpts, channelId)
}

// CloseChannel is a paid mutator transaction binding the contract method 0x4c2ee09d.
//
// Solidity: function closeChannel(bytes32 channelId) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) CloseChannel(channelId [32]byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.CloseChannel(&_UniversalChannelHandler.TransactOpts, channelId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dispatcher) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) Initialize(opts *bind.TransactOpts, _dispatcher common.Address) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "initialize", _dispatcher)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dispatcher) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) Initialize(_dispatcher common.Address) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.Initialize(&_UniversalChannelHandler.TransactOpts, _dispatcher)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _dispatcher) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) Initialize(_dispatcher common.Address) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.Initialize(&_UniversalChannelHandler.TransactOpts, _dispatcher)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) OnAcknowledgementPacket(opts *bind.TransactOpts, packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "onAcknowledgementPacket", packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OnAcknowledgementPacket(&_UniversalChannelHandler.TransactOpts, packet, ack)
}

// OnAcknowledgementPacket is a paid mutator transaction binding the contract method 0x7e1d42b5.
//
// Solidity: function onAcknowledgementPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (bool,bytes) ack) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) OnAcknowledgementPacket(packet IbcPacket, ack AckPacket) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OnAcknowledgementPacket(&_UniversalChannelHandler.TransactOpts, packet, ack)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string , bytes32 ) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) OnChanCloseConfirm(opts *bind.TransactOpts, channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "onChanCloseConfirm", channelId, arg1, arg2)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string , bytes32 ) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) OnChanCloseConfirm(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OnChanCloseConfirm(&_UniversalChannelHandler.TransactOpts, channelId, arg1, arg2)
}

// OnChanCloseConfirm is a paid mutator transaction binding the contract method 0x3f9fdbe4.
//
// Solidity: function onChanCloseConfirm(bytes32 channelId, string , bytes32 ) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) OnChanCloseConfirm(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OnChanCloseConfirm(&_UniversalChannelHandler.TransactOpts, channelId, arg1, arg2)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string , bytes32 ) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) OnChanCloseInit(opts *bind.TransactOpts, channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "onChanCloseInit", channelId, arg1, arg2)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string , bytes32 ) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) OnChanCloseInit(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OnChanCloseInit(&_UniversalChannelHandler.TransactOpts, channelId, arg1, arg2)
}

// OnChanCloseInit is a paid mutator transaction binding the contract method 0x1eb7dd5e.
//
// Solidity: function onChanCloseInit(bytes32 channelId, string , bytes32 ) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) OnChanCloseInit(channelId [32]byte, arg1 string, arg2 [32]byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OnChanCloseInit(&_UniversalChannelHandler.TransactOpts, channelId, arg1, arg2)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) OnChanOpenConfirm(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "onChanOpenConfirm", channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OnChanOpenConfirm(&_UniversalChannelHandler.TransactOpts, channelId)
}

// OnChanOpenConfirm is a paid mutator transaction binding the contract method 0xfad28a24.
//
// Solidity: function onChanOpenConfirm(bytes32 channelId) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) OnChanOpenConfirm(channelId [32]byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OnChanOpenConfirm(&_UniversalChannelHandler.TransactOpts, channelId)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket, bool skipAck)
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) OnRecvPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "onRecvPacket", packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket, bool skipAck)
func (_UniversalChannelHandler *UniversalChannelHandlerSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OnRecvPacket(&_UniversalChannelHandler.TransactOpts, packet)
}

// OnRecvPacket is a paid mutator transaction binding the contract method 0x4dcc0aa6.
//
// Solidity: function onRecvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns((bool,bytes) ackPacket, bool skipAck)
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) OnRecvPacket(packet IbcPacket) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OnRecvPacket(&_UniversalChannelHandler.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) OnTimeoutPacket(opts *bind.TransactOpts, packet IbcPacket) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "onTimeoutPacket", packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OnTimeoutPacket(&_UniversalChannelHandler.TransactOpts, packet)
}

// OnTimeoutPacket is a paid mutator transaction binding the contract method 0x602f9834.
//
// Solidity: function onTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) OnTimeoutPacket(packet IbcPacket) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OnTimeoutPacket(&_UniversalChannelHandler.TransactOpts, packet)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xace02de7.
//
// Solidity: function openChannel(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortIdentifier) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) OpenChannel(opts *bind.TransactOpts, version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortIdentifier string) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "openChannel", version, ordering, feeEnabled, connectionHops, counterpartyPortIdentifier)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xace02de7.
//
// Solidity: function openChannel(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortIdentifier) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) OpenChannel(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortIdentifier string) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OpenChannel(&_UniversalChannelHandler.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortIdentifier)
}

// OpenChannel is a paid mutator transaction binding the contract method 0xace02de7.
//
// Solidity: function openChannel(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortIdentifier) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) OpenChannel(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortIdentifier string) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.OpenChannel(&_UniversalChannelHandler.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortIdentifier)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.RenounceOwnership(&_UniversalChannelHandler.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.RenounceOwnership(&_UniversalChannelHandler.TransactOpts)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) SendUniversalPacket(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "sendUniversalPacket", channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_UniversalChannelHandler *UniversalChannelHandlerSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.SendUniversalPacket(&_UniversalChannelHandler.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacket is a paid mutator transaction binding the contract method 0x1f3a5830.
//
// Solidity: function sendUniversalPacket(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) SendUniversalPacket(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.SendUniversalPacket(&_UniversalChannelHandler.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) SendUniversalPacketWithFee(opts *bind.TransactOpts, channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "sendUniversalPacketWithFee", channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_UniversalChannelHandler *UniversalChannelHandlerSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.SendUniversalPacketWithFee(&_UniversalChannelHandler.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SendUniversalPacketWithFee is a paid mutator transaction binding the contract method 0x462fdf83.
//
// Solidity: function sendUniversalPacketWithFee(bytes32 channelId, bytes32 destPortAddr, bytes appData, uint64 timeoutTimestamp, uint256[2] gasLimits, uint256[2] gasPrices) payable returns(uint64 sequence)
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) SendUniversalPacketWithFee(channelId [32]byte, destPortAddr [32]byte, appData []byte, timeoutTimestamp uint64, gasLimits [2]*big.Int, gasPrices [2]*big.Int) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.SendUniversalPacketWithFee(&_UniversalChannelHandler.TransactOpts, channelId, destPortAddr, appData, timeoutTimestamp, gasLimits, gasPrices)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address _dispatcher) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) SetDispatcher(opts *bind.TransactOpts, _dispatcher common.Address) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "setDispatcher", _dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address _dispatcher) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) SetDispatcher(_dispatcher common.Address) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.SetDispatcher(&_UniversalChannelHandler.TransactOpts, _dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address _dispatcher) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) SetDispatcher(_dispatcher common.Address) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.SetDispatcher(&_UniversalChannelHandler.TransactOpts, _dispatcher)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.TransferOwnership(&_UniversalChannelHandler.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.TransferOwnership(&_UniversalChannelHandler.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.UpgradeTo(&_UniversalChannelHandler.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.UpgradeTo(&_UniversalChannelHandler.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.UpgradeToAndCall(&_UniversalChannelHandler.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.UpgradeToAndCall(&_UniversalChannelHandler.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniversalChannelHandler.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniversalChannelHandler *UniversalChannelHandlerSession) Receive() (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.Receive(&_UniversalChannelHandler.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniversalChannelHandler *UniversalChannelHandlerTransactorSession) Receive() (*types.Transaction, error) {
	return _UniversalChannelHandler.Contract.Receive(&_UniversalChannelHandler.TransactOpts)
}

// UniversalChannelHandlerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerAdminChangedIterator struct {
	Event *UniversalChannelHandlerAdminChanged // Event containing the contract specifics and raw log

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
func (it *UniversalChannelHandlerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalChannelHandlerAdminChanged)
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
		it.Event = new(UniversalChannelHandlerAdminChanged)
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
func (it *UniversalChannelHandlerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalChannelHandlerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalChannelHandlerAdminChanged represents a AdminChanged event raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*UniversalChannelHandlerAdminChangedIterator, error) {

	logs, sub, err := _UniversalChannelHandler.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &UniversalChannelHandlerAdminChangedIterator{contract: _UniversalChannelHandler.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *UniversalChannelHandlerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _UniversalChannelHandler.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalChannelHandlerAdminChanged)
				if err := _UniversalChannelHandler.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) ParseAdminChanged(log types.Log) (*UniversalChannelHandlerAdminChanged, error) {
	event := new(UniversalChannelHandlerAdminChanged)
	if err := _UniversalChannelHandler.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalChannelHandlerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerBeaconUpgradedIterator struct {
	Event *UniversalChannelHandlerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *UniversalChannelHandlerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalChannelHandlerBeaconUpgraded)
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
		it.Event = new(UniversalChannelHandlerBeaconUpgraded)
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
func (it *UniversalChannelHandlerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalChannelHandlerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalChannelHandlerBeaconUpgraded represents a BeaconUpgraded event raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*UniversalChannelHandlerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UniversalChannelHandler.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &UniversalChannelHandlerBeaconUpgradedIterator{contract: _UniversalChannelHandler.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *UniversalChannelHandlerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UniversalChannelHandler.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalChannelHandlerBeaconUpgraded)
				if err := _UniversalChannelHandler.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) ParseBeaconUpgraded(log types.Log) (*UniversalChannelHandlerBeaconUpgraded, error) {
	event := new(UniversalChannelHandlerBeaconUpgraded)
	if err := _UniversalChannelHandler.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalChannelHandlerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerInitializedIterator struct {
	Event *UniversalChannelHandlerInitialized // Event containing the contract specifics and raw log

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
func (it *UniversalChannelHandlerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalChannelHandlerInitialized)
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
		it.Event = new(UniversalChannelHandlerInitialized)
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
func (it *UniversalChannelHandlerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalChannelHandlerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalChannelHandlerInitialized represents a Initialized event raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) FilterInitialized(opts *bind.FilterOpts) (*UniversalChannelHandlerInitializedIterator, error) {

	logs, sub, err := _UniversalChannelHandler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UniversalChannelHandlerInitializedIterator{contract: _UniversalChannelHandler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UniversalChannelHandlerInitialized) (event.Subscription, error) {

	logs, sub, err := _UniversalChannelHandler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalChannelHandlerInitialized)
				if err := _UniversalChannelHandler.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) ParseInitialized(log types.Log) (*UniversalChannelHandlerInitialized, error) {
	event := new(UniversalChannelHandlerInitialized)
	if err := _UniversalChannelHandler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalChannelHandlerOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerOwnershipTransferStartedIterator struct {
	Event *UniversalChannelHandlerOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *UniversalChannelHandlerOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalChannelHandlerOwnershipTransferStarted)
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
		it.Event = new(UniversalChannelHandlerOwnershipTransferStarted)
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
func (it *UniversalChannelHandlerOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalChannelHandlerOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalChannelHandlerOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniversalChannelHandlerOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalChannelHandler.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalChannelHandlerOwnershipTransferStartedIterator{contract: _UniversalChannelHandler.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *UniversalChannelHandlerOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalChannelHandler.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalChannelHandlerOwnershipTransferStarted)
				if err := _UniversalChannelHandler.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) ParseOwnershipTransferStarted(log types.Log) (*UniversalChannelHandlerOwnershipTransferStarted, error) {
	event := new(UniversalChannelHandlerOwnershipTransferStarted)
	if err := _UniversalChannelHandler.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalChannelHandlerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerOwnershipTransferredIterator struct {
	Event *UniversalChannelHandlerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UniversalChannelHandlerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalChannelHandlerOwnershipTransferred)
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
		it.Event = new(UniversalChannelHandlerOwnershipTransferred)
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
func (it *UniversalChannelHandlerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalChannelHandlerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalChannelHandlerOwnershipTransferred represents a OwnershipTransferred event raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniversalChannelHandlerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalChannelHandler.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalChannelHandlerOwnershipTransferredIterator{contract: _UniversalChannelHandler.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UniversalChannelHandlerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UniversalChannelHandler.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalChannelHandlerOwnershipTransferred)
				if err := _UniversalChannelHandler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) ParseOwnershipTransferred(log types.Log) (*UniversalChannelHandlerOwnershipTransferred, error) {
	event := new(UniversalChannelHandlerOwnershipTransferred)
	if err := _UniversalChannelHandler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalChannelHandlerUCHPacketSentIterator is returned from FilterUCHPacketSent and is used to iterate over the raw logs and unpacked data for UCHPacketSent events raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerUCHPacketSentIterator struct {
	Event *UniversalChannelHandlerUCHPacketSent // Event containing the contract specifics and raw log

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
func (it *UniversalChannelHandlerUCHPacketSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalChannelHandlerUCHPacketSent)
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
		it.Event = new(UniversalChannelHandlerUCHPacketSent)
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
func (it *UniversalChannelHandlerUCHPacketSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalChannelHandlerUCHPacketSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalChannelHandlerUCHPacketSent represents a UCHPacketSent event raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerUCHPacketSent struct {
	Source      common.Address
	Destination [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUCHPacketSent is a free log retrieval operation binding the contract event 0x9831d8c66285bfd33de069ced58ad437d6bf08f63446bf06c3713e40b4b7e873.
//
// Solidity: event UCHPacketSent(address source, bytes32 destination)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) FilterUCHPacketSent(opts *bind.FilterOpts) (*UniversalChannelHandlerUCHPacketSentIterator, error) {

	logs, sub, err := _UniversalChannelHandler.contract.FilterLogs(opts, "UCHPacketSent")
	if err != nil {
		return nil, err
	}
	return &UniversalChannelHandlerUCHPacketSentIterator{contract: _UniversalChannelHandler.contract, event: "UCHPacketSent", logs: logs, sub: sub}, nil
}

// WatchUCHPacketSent is a free log subscription operation binding the contract event 0x9831d8c66285bfd33de069ced58ad437d6bf08f63446bf06c3713e40b4b7e873.
//
// Solidity: event UCHPacketSent(address source, bytes32 destination)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) WatchUCHPacketSent(opts *bind.WatchOpts, sink chan<- *UniversalChannelHandlerUCHPacketSent) (event.Subscription, error) {

	logs, sub, err := _UniversalChannelHandler.contract.WatchLogs(opts, "UCHPacketSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalChannelHandlerUCHPacketSent)
				if err := _UniversalChannelHandler.contract.UnpackLog(event, "UCHPacketSent", log); err != nil {
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
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) ParseUCHPacketSent(log types.Log) (*UniversalChannelHandlerUCHPacketSent, error) {
	event := new(UniversalChannelHandlerUCHPacketSent)
	if err := _UniversalChannelHandler.contract.UnpackLog(event, "UCHPacketSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalChannelHandlerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerUpgradedIterator struct {
	Event *UniversalChannelHandlerUpgraded // Event containing the contract specifics and raw log

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
func (it *UniversalChannelHandlerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalChannelHandlerUpgraded)
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
		it.Event = new(UniversalChannelHandlerUpgraded)
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
func (it *UniversalChannelHandlerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalChannelHandlerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalChannelHandlerUpgraded represents a Upgraded event raised by the UniversalChannelHandler contract.
type UniversalChannelHandlerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UniversalChannelHandlerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UniversalChannelHandler.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UniversalChannelHandlerUpgradedIterator{contract: _UniversalChannelHandler.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UniversalChannelHandlerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UniversalChannelHandler.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalChannelHandlerUpgraded)
				if err := _UniversalChannelHandler.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_UniversalChannelHandler *UniversalChannelHandlerFilterer) ParseUpgraded(log types.Log) (*UniversalChannelHandlerUpgraded, error) {
	event := new(UniversalChannelHandlerUpgraded)
	if err := _UniversalChannelHandler.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
