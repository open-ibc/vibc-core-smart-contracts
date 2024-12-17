// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crossl2prover

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

// CrossL2ProverMetaData contains all meta data concerning the CrossL2Prover contract.
var CrossL2ProverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"verifier_\",\"type\":\"address\",\"internalType\":\"contractISignatureVerifier\"},{\"name\":\"clientType_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"LIGHT_CLIENT_TYPE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumLightClientType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clientType\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getState\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"peptideAppHashes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateClient\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"peptideHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"peptideAppHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"validateLog\",\"inputs\":[{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"emittingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"topics\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"unindexedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateReceipt\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"srcChainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"receiptRLP\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISignatureVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMembership\",\"inputs\":[{\"name\":\"appHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proofs\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"verifyNonMembership\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"CannotUpdateClientWithDifferentAppHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAppHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidIbcStateProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1BlockHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidL1BlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPacketProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedL1BlockNumber\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRLPEncodedL1StateRoot\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MethodNotImplemented\",\"inputs\":[]}]",
}

// CrossL2ProverABI is the input ABI used to generate the binding from.
// Deprecated: Use CrossL2ProverMetaData.ABI instead.
var CrossL2ProverABI = CrossL2ProverMetaData.ABI

// CrossL2Prover is an auto generated Go binding around an Ethereum contract.
type CrossL2Prover struct {
	CrossL2ProverCaller     // Read-only binding to the contract
	CrossL2ProverTransactor // Write-only binding to the contract
	CrossL2ProverFilterer   // Log filterer for contract events
}

// CrossL2ProverCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossL2ProverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossL2ProverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossL2ProverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossL2ProverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossL2ProverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossL2ProverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossL2ProverSession struct {
	Contract     *CrossL2Prover    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossL2ProverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossL2ProverCallerSession struct {
	Contract *CrossL2ProverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CrossL2ProverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossL2ProverTransactorSession struct {
	Contract     *CrossL2ProverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CrossL2ProverRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossL2ProverRaw struct {
	Contract *CrossL2Prover // Generic contract binding to access the raw methods on
}

// CrossL2ProverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossL2ProverCallerRaw struct {
	Contract *CrossL2ProverCaller // Generic read-only contract binding to access the raw methods on
}

// CrossL2ProverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossL2ProverTransactorRaw struct {
	Contract *CrossL2ProverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossL2Prover creates a new instance of CrossL2Prover, bound to a specific deployed contract.
func NewCrossL2Prover(address common.Address, backend bind.ContractBackend) (*CrossL2Prover, error) {
	contract, err := bindCrossL2Prover(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossL2Prover{CrossL2ProverCaller: CrossL2ProverCaller{contract: contract}, CrossL2ProverTransactor: CrossL2ProverTransactor{contract: contract}, CrossL2ProverFilterer: CrossL2ProverFilterer{contract: contract}}, nil
}

// NewCrossL2ProverCaller creates a new read-only instance of CrossL2Prover, bound to a specific deployed contract.
func NewCrossL2ProverCaller(address common.Address, caller bind.ContractCaller) (*CrossL2ProverCaller, error) {
	contract, err := bindCrossL2Prover(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossL2ProverCaller{contract: contract}, nil
}

// NewCrossL2ProverTransactor creates a new write-only instance of CrossL2Prover, bound to a specific deployed contract.
func NewCrossL2ProverTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossL2ProverTransactor, error) {
	contract, err := bindCrossL2Prover(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossL2ProverTransactor{contract: contract}, nil
}

// NewCrossL2ProverFilterer creates a new log filterer instance of CrossL2Prover, bound to a specific deployed contract.
func NewCrossL2ProverFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossL2ProverFilterer, error) {
	contract, err := bindCrossL2Prover(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossL2ProverFilterer{contract: contract}, nil
}

// bindCrossL2Prover binds a generic wrapper to an already deployed contract.
func bindCrossL2Prover(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrossL2ProverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossL2Prover *CrossL2ProverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossL2Prover.Contract.CrossL2ProverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossL2Prover *CrossL2ProverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossL2Prover.Contract.CrossL2ProverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossL2Prover *CrossL2ProverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossL2Prover.Contract.CrossL2ProverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossL2Prover *CrossL2ProverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossL2Prover.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossL2Prover *CrossL2ProverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossL2Prover.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossL2Prover *CrossL2ProverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossL2Prover.Contract.contract.Transact(opts, method, params...)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_CrossL2Prover *CrossL2ProverCaller) LIGHTCLIENTTYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrossL2Prover.contract.Call(opts, &out, "LIGHT_CLIENT_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_CrossL2Prover *CrossL2ProverSession) LIGHTCLIENTTYPE() (uint8, error) {
	return _CrossL2Prover.Contract.LIGHTCLIENTTYPE(&_CrossL2Prover.CallOpts)
}

// LIGHTCLIENTTYPE is a free data retrieval call binding the contract method 0x57c1c5f4.
//
// Solidity: function LIGHT_CLIENT_TYPE() view returns(uint8)
func (_CrossL2Prover *CrossL2ProverCallerSession) LIGHTCLIENTTYPE() (uint8, error) {
	return _CrossL2Prover.Contract.LIGHTCLIENTTYPE(&_CrossL2Prover.CallOpts)
}

// ClientType is a free data retrieval call binding the contract method 0xb3768f0d.
//
// Solidity: function clientType() view returns(string)
func (_CrossL2Prover *CrossL2ProverCaller) ClientType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrossL2Prover.contract.Call(opts, &out, "clientType")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ClientType is a free data retrieval call binding the contract method 0xb3768f0d.
//
// Solidity: function clientType() view returns(string)
func (_CrossL2Prover *CrossL2ProverSession) ClientType() (string, error) {
	return _CrossL2Prover.Contract.ClientType(&_CrossL2Prover.CallOpts)
}

// ClientType is a free data retrieval call binding the contract method 0xb3768f0d.
//
// Solidity: function clientType() view returns(string)
func (_CrossL2Prover *CrossL2ProverCallerSession) ClientType() (string, error) {
	return _CrossL2Prover.Contract.ClientType(&_CrossL2Prover.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256)
func (_CrossL2Prover *CrossL2ProverCaller) GetState(opts *bind.CallOpts, height *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossL2Prover.contract.Call(opts, &out, "getState", height)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256)
func (_CrossL2Prover *CrossL2ProverSession) GetState(height *big.Int) (*big.Int, error) {
	return _CrossL2Prover.Contract.GetState(&_CrossL2Prover.CallOpts, height)
}

// GetState is a free data retrieval call binding the contract method 0x44c9af28.
//
// Solidity: function getState(uint256 height) view returns(uint256)
func (_CrossL2Prover *CrossL2ProverCallerSession) GetState(height *big.Int) (*big.Int, error) {
	return _CrossL2Prover.Contract.GetState(&_CrossL2Prover.CallOpts, height)
}

// PeptideAppHashes is a free data retrieval call binding the contract method 0xc67e15f7.
//
// Solidity: function peptideAppHashes(uint256 ) view returns(uint256)
func (_CrossL2Prover *CrossL2ProverCaller) PeptideAppHashes(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrossL2Prover.contract.Call(opts, &out, "peptideAppHashes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeptideAppHashes is a free data retrieval call binding the contract method 0xc67e15f7.
//
// Solidity: function peptideAppHashes(uint256 ) view returns(uint256)
func (_CrossL2Prover *CrossL2ProverSession) PeptideAppHashes(arg0 *big.Int) (*big.Int, error) {
	return _CrossL2Prover.Contract.PeptideAppHashes(&_CrossL2Prover.CallOpts, arg0)
}

// PeptideAppHashes is a free data retrieval call binding the contract method 0xc67e15f7.
//
// Solidity: function peptideAppHashes(uint256 ) view returns(uint256)
func (_CrossL2Prover *CrossL2ProverCallerSession) PeptideAppHashes(arg0 *big.Int) (*big.Int, error) {
	return _CrossL2Prover.Contract.PeptideAppHashes(&_CrossL2Prover.CallOpts, arg0)
}

// ValidateLog is a free data retrieval call binding the contract method 0xfc396ddb.
//
// Solidity: function validateLog(uint256 logIndex, bytes proof) view returns(bytes32 chainId, address emittingContract, bytes[] topics, bytes unindexedData)
func (_CrossL2Prover *CrossL2ProverCaller) ValidateLog(opts *bind.CallOpts, logIndex *big.Int, proof []byte) (struct {
	ChainId          [32]byte
	EmittingContract common.Address
	Topics           [][]byte
	UnindexedData    []byte
}, error) {
	var out []interface{}
	err := _CrossL2Prover.contract.Call(opts, &out, "validateLog", logIndex, proof)

	outstruct := new(struct {
		ChainId          [32]byte
		EmittingContract common.Address
		Topics           [][]byte
		UnindexedData    []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ChainId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.EmittingContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Topics = *abi.ConvertType(out[2], new([][]byte)).(*[][]byte)
	outstruct.UnindexedData = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ValidateLog is a free data retrieval call binding the contract method 0xfc396ddb.
//
// Solidity: function validateLog(uint256 logIndex, bytes proof) view returns(bytes32 chainId, address emittingContract, bytes[] topics, bytes unindexedData)
func (_CrossL2Prover *CrossL2ProverSession) ValidateLog(logIndex *big.Int, proof []byte) (struct {
	ChainId          [32]byte
	EmittingContract common.Address
	Topics           [][]byte
	UnindexedData    []byte
}, error) {
	return _CrossL2Prover.Contract.ValidateLog(&_CrossL2Prover.CallOpts, logIndex, proof)
}

// ValidateLog is a free data retrieval call binding the contract method 0xfc396ddb.
//
// Solidity: function validateLog(uint256 logIndex, bytes proof) view returns(bytes32 chainId, address emittingContract, bytes[] topics, bytes unindexedData)
func (_CrossL2Prover *CrossL2ProverCallerSession) ValidateLog(logIndex *big.Int, proof []byte) (struct {
	ChainId          [32]byte
	EmittingContract common.Address
	Topics           [][]byte
	UnindexedData    []byte
}, error) {
	return _CrossL2Prover.Contract.ValidateLog(&_CrossL2Prover.CallOpts, logIndex, proof)
}

// ValidateReceipt is a free data retrieval call binding the contract method 0x2cd78e77.
//
// Solidity: function validateReceipt(bytes proof) view returns(bytes32 srcChainID, bytes receiptRLP)
func (_CrossL2Prover *CrossL2ProverCaller) ValidateReceipt(opts *bind.CallOpts, proof []byte) (struct {
	SrcChainID [32]byte
	ReceiptRLP []byte
}, error) {
	var out []interface{}
	err := _CrossL2Prover.contract.Call(opts, &out, "validateReceipt", proof)

	outstruct := new(struct {
		SrcChainID [32]byte
		ReceiptRLP []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SrcChainID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ReceiptRLP = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// ValidateReceipt is a free data retrieval call binding the contract method 0x2cd78e77.
//
// Solidity: function validateReceipt(bytes proof) view returns(bytes32 srcChainID, bytes receiptRLP)
func (_CrossL2Prover *CrossL2ProverSession) ValidateReceipt(proof []byte) (struct {
	SrcChainID [32]byte
	ReceiptRLP []byte
}, error) {
	return _CrossL2Prover.Contract.ValidateReceipt(&_CrossL2Prover.CallOpts, proof)
}

// ValidateReceipt is a free data retrieval call binding the contract method 0x2cd78e77.
//
// Solidity: function validateReceipt(bytes proof) view returns(bytes32 srcChainID, bytes receiptRLP)
func (_CrossL2Prover *CrossL2ProverCallerSession) ValidateReceipt(proof []byte) (struct {
	SrcChainID [32]byte
	ReceiptRLP []byte
}, error) {
	return _CrossL2Prover.Contract.ValidateReceipt(&_CrossL2Prover.CallOpts, proof)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_CrossL2Prover *CrossL2ProverCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrossL2Prover.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_CrossL2Prover *CrossL2ProverSession) Verifier() (common.Address, error) {
	return _CrossL2Prover.Contract.Verifier(&_CrossL2Prover.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_CrossL2Prover *CrossL2ProverCallerSession) Verifier() (common.Address, error) {
	return _CrossL2Prover.Contract.Verifier(&_CrossL2Prover.CallOpts)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proofs) pure returns()
func (_CrossL2Prover *CrossL2ProverCaller) VerifyMembership(opts *bind.CallOpts, appHash [32]byte, key []byte, value []byte, proofs Ics23Proof) error {
	var out []interface{}
	err := _CrossL2Prover.contract.Call(opts, &out, "verifyMembership", appHash, key, value, proofs)

	if err != nil {
		return err
	}

	return err

}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proofs) pure returns()
func (_CrossL2Prover *CrossL2ProverSession) VerifyMembership(appHash [32]byte, key []byte, value []byte, proofs Ics23Proof) error {
	return _CrossL2Prover.Contract.VerifyMembership(&_CrossL2Prover.CallOpts, appHash, key, value, proofs)
}

// VerifyMembership is a free data retrieval call binding the contract method 0xc2f0329f.
//
// Solidity: function verifyMembership(bytes32 appHash, bytes key, bytes value, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proofs) pure returns()
func (_CrossL2Prover *CrossL2ProverCallerSession) VerifyMembership(appHash [32]byte, key []byte, value []byte, proofs Ics23Proof) error {
	return _CrossL2Prover.Contract.VerifyMembership(&_CrossL2Prover.CallOpts, appHash, key, value, proofs)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_CrossL2Prover *CrossL2ProverCaller) VerifyNonMembership(opts *bind.CallOpts, arg0 [32]byte, arg1 []byte, arg2 Ics23Proof) error {
	var out []interface{}
	err := _CrossL2Prover.contract.Call(opts, &out, "verifyNonMembership", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_CrossL2Prover *CrossL2ProverSession) VerifyNonMembership(arg0 [32]byte, arg1 []byte, arg2 Ics23Proof) error {
	return _CrossL2Prover.Contract.VerifyNonMembership(&_CrossL2Prover.CallOpts, arg0, arg1, arg2)
}

// VerifyNonMembership is a free data retrieval call binding the contract method 0x2a6ded74.
//
// Solidity: function verifyNonMembership(bytes32 , bytes , (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) ) pure returns()
func (_CrossL2Prover *CrossL2ProverCallerSession) VerifyNonMembership(arg0 [32]byte, arg1 []byte, arg2 Ics23Proof) error {
	return _CrossL2Prover.Contract.VerifyNonMembership(&_CrossL2Prover.CallOpts, arg0, arg1, arg2)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes proof, uint256 peptideHeight, uint256 peptideAppHash) returns()
func (_CrossL2Prover *CrossL2ProverTransactor) UpdateClient(opts *bind.TransactOpts, proof []byte, peptideHeight *big.Int, peptideAppHash *big.Int) (*types.Transaction, error) {
	return _CrossL2Prover.contract.Transact(opts, "updateClient", proof, peptideHeight, peptideAppHash)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes proof, uint256 peptideHeight, uint256 peptideAppHash) returns()
func (_CrossL2Prover *CrossL2ProverSession) UpdateClient(proof []byte, peptideHeight *big.Int, peptideAppHash *big.Int) (*types.Transaction, error) {
	return _CrossL2Prover.Contract.UpdateClient(&_CrossL2Prover.TransactOpts, proof, peptideHeight, peptideAppHash)
}

// UpdateClient is a paid mutator transaction binding the contract method 0x49ff245e.
//
// Solidity: function updateClient(bytes proof, uint256 peptideHeight, uint256 peptideAppHash) returns()
func (_CrossL2Prover *CrossL2ProverTransactorSession) UpdateClient(proof []byte, peptideHeight *big.Int, peptideAppHash *big.Int) (*types.Transaction, error) {
	return _CrossL2Prover.Contract.UpdateClient(&_CrossL2Prover.TransactOpts, proof, peptideHeight, peptideAppHash)
}
