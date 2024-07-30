// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dispatcher

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

// Channel is an auto generated low-level Go binding around an user-defined struct.
type Channel struct {
	Version               string
	Ordering              uint8
	FeeEnabled            bool
	ConnectionHops        []string
	CounterpartyPortId    string
	CounterpartyChannelId [32]byte
	PortId                string
}

// ChannelEnd is an auto generated low-level Go binding around an user-defined struct.
type ChannelEnd struct {
	PortId    string
	ChannelId [32]byte
	Version   string
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

// DispatcherMetaData contains all meta data concerning the Dispatcher contract.
var DispatcherMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acknowledgement\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"ack\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelCloseConfirm\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelCloseInit\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenAck\",\"inputs\":[{\"name\":\"local\",\"type\":\"tuple\",\"internalType\":\"structChannelEnd\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelEnd\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenConfirm\",\"inputs\":[{\"name\":\"local\",\"type\":\"tuple\",\"internalType\":\"structChannelEnd\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelEnd\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenInit\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"channelOpenTry\",\"inputs\":[{\"name\":\"local\",\"type\":\"tuple\",\"internalType\":\"structChannelEnd\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterparty\",\"type\":\"tuple\",\"internalType\":\"structChannelEnd\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeVault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFeeVault\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChannel\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"channel\",\"type\":\"tuple\",\"internalType\":\"structChannel\",\"components\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOptimisticConsensusState\",\"inputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"connection\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fraudProofEndTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ended\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initPortPrefix\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_feeVault\",\"type\":\"address\",\"internalType\":\"contractIFeeVault\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"portPrefix\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"portPrefixLen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recvPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeConnection\",\"inputs\":[{\"name\":\"connection\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendPacket\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClientForConnection\",\"inputs\":[{\"name\":\"connection\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"lightClient\",\"type\":\"address\",\"internalType\":\"contractILightClient\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPortPrefix\",\"inputs\":[{\"name\":\"_portPrefix\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"timeout\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateClientWithOptimisticConsensusState\",\"inputs\":[{\"name\":\"l1header\",\"type\":\"tuple\",\"internalType\":\"structL1Header\",\"components\":[{\"name\":\"header\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"number\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structOpL2StateProof\",\"components\":[{\"name\":\"accountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"outputRootProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"l2OutputProposalKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"appHash\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"connection\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"fraudProofEndTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ended\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"writeTimeoutPacket\",\"inputs\":[{\"name\":\"packet\",\"type\":\"tuple\",\"internalType\":\"structIbcPacket\",\"components\":[{\"name\":\"src\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"dest\",\"type\":\"tuple\",\"internalType\":\"structIbcEndpoint\",\"components\":[{\"name\":\"portId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"sequence\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"tuple\",\"internalType\":\"structIcs23Proof\",\"components\":[{\"name\":\"proof\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23Proof[]\",\"components\":[{\"name\":\"path\",\"type\":\"tuple[]\",\"internalType\":\"structOpIcs23ProofPath[]\",\"components\":[{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"suffix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"key\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"prefix\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Acknowledgement\",\"inputs\":[{\"name\":\"sourcePortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AcknowledgementError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelCloseConfirm\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelCloseConfirmError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelCloseInit\",\"inputs\":[{\"name\":\"portAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelCloseInitError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenAck\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenAckError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenConfirm\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"channelId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenConfirmError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenInit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenInitError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenTry\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"ordering\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumChannelOrder\"},{\"name\":\"feeEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"connectionHops\",\"type\":\"string[]\",\"indexed\":false,\"internalType\":\"string[]\"},{\"name\":\"counterpartyPortId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"counterpartyChannelId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChannelOpenTryError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RecvPacket\",\"inputs\":[{\"name\":\"destPortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SendPacket\",\"inputs\":[{\"name\":\"sourcePortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"packet\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Timeout\",\"inputs\":[{\"name\":\"sourcePortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TimeoutError\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WriteAckPacket\",\"inputs\":[{\"name\":\"writerPortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"writerChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"ackPacket\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structAckPacket\",\"components\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WriteTimeoutPacket\",\"inputs\":[{\"name\":\"writerPortAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"writerChannelId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"sequence\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"timeoutHeight\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structHeight\",\"components\":[{\"name\":\"revision_number\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"revision_height\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"timeoutTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ackPacketCommitmentAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"channelIdNotFound\",\"inputs\":[{\"name\":\"channelId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"channelNotOwnedByPortAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"channelNotOwnedBySender\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidConnection\",\"inputs\":[{\"name\":\"connection\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"invalidConnectionHops\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidCounterParty\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidPacket\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidPacketSequence\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"invalidPortPrefix\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"lightClientNotFound\",\"inputs\":[{\"name\":\"connection\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"notEnoughGas\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"packetCommitmentNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"packetNotTimedOut\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"packetReceiptAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"unexpectedPacketSequence\",\"inputs\":[]}]",
}

// DispatcherABI is the input ABI used to generate the binding from.
// Deprecated: Use DispatcherMetaData.ABI instead.
var DispatcherABI = DispatcherMetaData.ABI

// Dispatcher is an auto generated Go binding around an Ethereum contract.
type Dispatcher struct {
	DispatcherCaller     // Read-only binding to the contract
	DispatcherTransactor // Write-only binding to the contract
	DispatcherFilterer   // Log filterer for contract events
}

// DispatcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type DispatcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DispatcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DispatcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DispatcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DispatcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DispatcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DispatcherSession struct {
	Contract     *Dispatcher       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DispatcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DispatcherCallerSession struct {
	Contract *DispatcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DispatcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DispatcherTransactorSession struct {
	Contract     *DispatcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DispatcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type DispatcherRaw struct {
	Contract *Dispatcher // Generic contract binding to access the raw methods on
}

// DispatcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DispatcherCallerRaw struct {
	Contract *DispatcherCaller // Generic read-only contract binding to access the raw methods on
}

// DispatcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DispatcherTransactorRaw struct {
	Contract *DispatcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDispatcher creates a new instance of Dispatcher, bound to a specific deployed contract.
func NewDispatcher(address common.Address, backend bind.ContractBackend) (*Dispatcher, error) {
	contract, err := bindDispatcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dispatcher{DispatcherCaller: DispatcherCaller{contract: contract}, DispatcherTransactor: DispatcherTransactor{contract: contract}, DispatcherFilterer: DispatcherFilterer{contract: contract}}, nil
}

// NewDispatcherCaller creates a new read-only instance of Dispatcher, bound to a specific deployed contract.
func NewDispatcherCaller(address common.Address, caller bind.ContractCaller) (*DispatcherCaller, error) {
	contract, err := bindDispatcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DispatcherCaller{contract: contract}, nil
}

// NewDispatcherTransactor creates a new write-only instance of Dispatcher, bound to a specific deployed contract.
func NewDispatcherTransactor(address common.Address, transactor bind.ContractTransactor) (*DispatcherTransactor, error) {
	contract, err := bindDispatcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DispatcherTransactor{contract: contract}, nil
}

// NewDispatcherFilterer creates a new log filterer instance of Dispatcher, bound to a specific deployed contract.
func NewDispatcherFilterer(address common.Address, filterer bind.ContractFilterer) (*DispatcherFilterer, error) {
	contract, err := bindDispatcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DispatcherFilterer{contract: contract}, nil
}

// bindDispatcher binds a generic wrapper to an already deployed contract.
func bindDispatcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DispatcherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dispatcher *DispatcherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dispatcher.Contract.DispatcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dispatcher *DispatcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dispatcher.Contract.DispatcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dispatcher *DispatcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dispatcher.Contract.DispatcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dispatcher *DispatcherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dispatcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dispatcher *DispatcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dispatcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dispatcher *DispatcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dispatcher.Contract.contract.Transact(opts, method, params...)
}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_Dispatcher *DispatcherCaller) FeeVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "feeVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_Dispatcher *DispatcherSession) FeeVault() (common.Address, error) {
	return _Dispatcher.Contract.FeeVault(&_Dispatcher.CallOpts)
}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_Dispatcher *DispatcherCallerSession) FeeVault() (common.Address, error) {
	return _Dispatcher.Contract.FeeVault(&_Dispatcher.CallOpts)
}

// GetChannel is a free data retrieval call binding the contract method 0x42852d24.
//
// Solidity: function getChannel(address portAddress, bytes32 channelId) view returns((string,uint8,bool,string[],string,bytes32,string) channel)
func (_Dispatcher *DispatcherCaller) GetChannel(opts *bind.CallOpts, portAddress common.Address, channelId [32]byte) (Channel, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "getChannel", portAddress, channelId)

	if err != nil {
		return *new(Channel), err
	}

	out0 := *abi.ConvertType(out[0], new(Channel)).(*Channel)

	return out0, err

}

// GetChannel is a free data retrieval call binding the contract method 0x42852d24.
//
// Solidity: function getChannel(address portAddress, bytes32 channelId) view returns((string,uint8,bool,string[],string,bytes32,string) channel)
func (_Dispatcher *DispatcherSession) GetChannel(portAddress common.Address, channelId [32]byte) (Channel, error) {
	return _Dispatcher.Contract.GetChannel(&_Dispatcher.CallOpts, portAddress, channelId)
}

// GetChannel is a free data retrieval call binding the contract method 0x42852d24.
//
// Solidity: function getChannel(address portAddress, bytes32 channelId) view returns((string,uint8,bool,string[],string,bytes32,string) channel)
func (_Dispatcher *DispatcherCallerSession) GetChannel(portAddress common.Address, channelId [32]byte) (Channel, error) {
	return _Dispatcher.Contract.GetChannel(&_Dispatcher.CallOpts, portAddress, channelId)
}

// GetOptimisticConsensusState is a free data retrieval call binding the contract method 0x8dd34bb4.
//
// Solidity: function getOptimisticConsensusState(uint256 height, string connection) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_Dispatcher *DispatcherCaller) GetOptimisticConsensusState(opts *bind.CallOpts, height *big.Int, connection string) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "getOptimisticConsensusState", height, connection)

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

// GetOptimisticConsensusState is a free data retrieval call binding the contract method 0x8dd34bb4.
//
// Solidity: function getOptimisticConsensusState(uint256 height, string connection) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_Dispatcher *DispatcherSession) GetOptimisticConsensusState(height *big.Int, connection string) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	return _Dispatcher.Contract.GetOptimisticConsensusState(&_Dispatcher.CallOpts, height, connection)
}

// GetOptimisticConsensusState is a free data retrieval call binding the contract method 0x8dd34bb4.
//
// Solidity: function getOptimisticConsensusState(uint256 height, string connection) view returns(uint256 appHash, uint256 fraudProofEndTime, bool ended)
func (_Dispatcher *DispatcherCallerSession) GetOptimisticConsensusState(height *big.Int, connection string) (struct {
	AppHash           *big.Int
	FraudProofEndTime *big.Int
	Ended             bool
}, error) {
	return _Dispatcher.Contract.GetOptimisticConsensusState(&_Dispatcher.CallOpts, height, connection)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dispatcher *DispatcherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dispatcher *DispatcherSession) Owner() (common.Address, error) {
	return _Dispatcher.Contract.Owner(&_Dispatcher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dispatcher *DispatcherCallerSession) Owner() (common.Address, error) {
	return _Dispatcher.Contract.Owner(&_Dispatcher.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Dispatcher *DispatcherCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Dispatcher *DispatcherSession) PendingOwner() (common.Address, error) {
	return _Dispatcher.Contract.PendingOwner(&_Dispatcher.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Dispatcher *DispatcherCallerSession) PendingOwner() (common.Address, error) {
	return _Dispatcher.Contract.PendingOwner(&_Dispatcher.CallOpts)
}

// PortPrefix is a free data retrieval call binding the contract method 0x7774a6d3.
//
// Solidity: function portPrefix() view returns(string)
func (_Dispatcher *DispatcherCaller) PortPrefix(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "portPrefix")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PortPrefix is a free data retrieval call binding the contract method 0x7774a6d3.
//
// Solidity: function portPrefix() view returns(string)
func (_Dispatcher *DispatcherSession) PortPrefix() (string, error) {
	return _Dispatcher.Contract.PortPrefix(&_Dispatcher.CallOpts)
}

// PortPrefix is a free data retrieval call binding the contract method 0x7774a6d3.
//
// Solidity: function portPrefix() view returns(string)
func (_Dispatcher *DispatcherCallerSession) PortPrefix() (string, error) {
	return _Dispatcher.Contract.PortPrefix(&_Dispatcher.CallOpts)
}

// PortPrefixLen is a free data retrieval call binding the contract method 0x2494546b.
//
// Solidity: function portPrefixLen() view returns(uint32)
func (_Dispatcher *DispatcherCaller) PortPrefixLen(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "portPrefixLen")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PortPrefixLen is a free data retrieval call binding the contract method 0x2494546b.
//
// Solidity: function portPrefixLen() view returns(uint32)
func (_Dispatcher *DispatcherSession) PortPrefixLen() (uint32, error) {
	return _Dispatcher.Contract.PortPrefixLen(&_Dispatcher.CallOpts)
}

// PortPrefixLen is a free data retrieval call binding the contract method 0x2494546b.
//
// Solidity: function portPrefixLen() view returns(uint32)
func (_Dispatcher *DispatcherCallerSession) PortPrefixLen() (uint32, error) {
	return _Dispatcher.Contract.PortPrefixLen(&_Dispatcher.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Dispatcher *DispatcherCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Dispatcher *DispatcherSession) ProxiableUUID() ([32]byte, error) {
	return _Dispatcher.Contract.ProxiableUUID(&_Dispatcher.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Dispatcher *DispatcherCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Dispatcher.Contract.ProxiableUUID(&_Dispatcher.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Dispatcher *DispatcherTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Dispatcher *DispatcherSession) AcceptOwnership() (*types.Transaction, error) {
	return _Dispatcher.Contract.AcceptOwnership(&_Dispatcher.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Dispatcher *DispatcherTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Dispatcher.Contract.AcceptOwnership(&_Dispatcher.TransactOpts)
}

// Acknowledgement is a paid mutator transaction binding the contract method 0xba5a4d25.
//
// Solidity: function acknowledgement(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, bytes ack, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactor) Acknowledgement(opts *bind.TransactOpts, packet IbcPacket, ack []byte, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "acknowledgement", packet, ack, proof)
}

// Acknowledgement is a paid mutator transaction binding the contract method 0xba5a4d25.
//
// Solidity: function acknowledgement(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, bytes ack, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherSession) Acknowledgement(packet IbcPacket, ack []byte, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.Acknowledgement(&_Dispatcher.TransactOpts, packet, ack, proof)
}

// Acknowledgement is a paid mutator transaction binding the contract method 0xba5a4d25.
//
// Solidity: function acknowledgement(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, bytes ack, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactorSession) Acknowledgement(packet IbcPacket, ack []byte, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.Acknowledgement(&_Dispatcher.TransactOpts, packet, ack, proof)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0xf90b8e96.
//
// Solidity: function channelCloseConfirm(address portAddress, bytes32 channelId, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactor) ChannelCloseConfirm(opts *bind.TransactOpts, portAddress common.Address, channelId [32]byte, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "channelCloseConfirm", portAddress, channelId, proof)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0xf90b8e96.
//
// Solidity: function channelCloseConfirm(address portAddress, bytes32 channelId, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherSession) ChannelCloseConfirm(portAddress common.Address, channelId [32]byte, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChannelCloseConfirm(&_Dispatcher.TransactOpts, portAddress, channelId, proof)
}

// ChannelCloseConfirm is a paid mutator transaction binding the contract method 0xf90b8e96.
//
// Solidity: function channelCloseConfirm(address portAddress, bytes32 channelId, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactorSession) ChannelCloseConfirm(portAddress common.Address, channelId [32]byte, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChannelCloseConfirm(&_Dispatcher.TransactOpts, portAddress, channelId, proof)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0x81bc079b.
//
// Solidity: function channelCloseInit(bytes32 channelId) returns()
func (_Dispatcher *DispatcherTransactor) ChannelCloseInit(opts *bind.TransactOpts, channelId [32]byte) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "channelCloseInit", channelId)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0x81bc079b.
//
// Solidity: function channelCloseInit(bytes32 channelId) returns()
func (_Dispatcher *DispatcherSession) ChannelCloseInit(channelId [32]byte) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChannelCloseInit(&_Dispatcher.TransactOpts, channelId)
}

// ChannelCloseInit is a paid mutator transaction binding the contract method 0x81bc079b.
//
// Solidity: function channelCloseInit(bytes32 channelId) returns()
func (_Dispatcher *DispatcherTransactorSession) ChannelCloseInit(channelId [32]byte) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChannelCloseInit(&_Dispatcher.TransactOpts, channelId)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0x1eb9fc86.
//
// Solidity: function channelOpenAck((string,bytes32,string) local, string[] connectionHops, uint8 ordering, bool feeEnabled, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactor) ChannelOpenAck(opts *bind.TransactOpts, local ChannelEnd, connectionHops []string, ordering uint8, feeEnabled bool, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "channelOpenAck", local, connectionHops, ordering, feeEnabled, counterparty, proof)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0x1eb9fc86.
//
// Solidity: function channelOpenAck((string,bytes32,string) local, string[] connectionHops, uint8 ordering, bool feeEnabled, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherSession) ChannelOpenAck(local ChannelEnd, connectionHops []string, ordering uint8, feeEnabled bool, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChannelOpenAck(&_Dispatcher.TransactOpts, local, connectionHops, ordering, feeEnabled, counterparty, proof)
}

// ChannelOpenAck is a paid mutator transaction binding the contract method 0x1eb9fc86.
//
// Solidity: function channelOpenAck((string,bytes32,string) local, string[] connectionHops, uint8 ordering, bool feeEnabled, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactorSession) ChannelOpenAck(local ChannelEnd, connectionHops []string, ordering uint8, feeEnabled bool, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChannelOpenAck(&_Dispatcher.TransactOpts, local, connectionHops, ordering, feeEnabled, counterparty, proof)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x429446b6.
//
// Solidity: function channelOpenConfirm((string,bytes32,string) local, string[] connectionHops, uint8 ordering, bool feeEnabled, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactor) ChannelOpenConfirm(opts *bind.TransactOpts, local ChannelEnd, connectionHops []string, ordering uint8, feeEnabled bool, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "channelOpenConfirm", local, connectionHops, ordering, feeEnabled, counterparty, proof)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x429446b6.
//
// Solidity: function channelOpenConfirm((string,bytes32,string) local, string[] connectionHops, uint8 ordering, bool feeEnabled, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherSession) ChannelOpenConfirm(local ChannelEnd, connectionHops []string, ordering uint8, feeEnabled bool, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChannelOpenConfirm(&_Dispatcher.TransactOpts, local, connectionHops, ordering, feeEnabled, counterparty, proof)
}

// ChannelOpenConfirm is a paid mutator transaction binding the contract method 0x429446b6.
//
// Solidity: function channelOpenConfirm((string,bytes32,string) local, string[] connectionHops, uint8 ordering, bool feeEnabled, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactorSession) ChannelOpenConfirm(local ChannelEnd, connectionHops []string, ordering uint8, feeEnabled bool, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChannelOpenConfirm(&_Dispatcher.TransactOpts, local, connectionHops, ordering, feeEnabled, counterparty, proof)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0x418925b7.
//
// Solidity: function channelOpenInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_Dispatcher *DispatcherTransactor) ChannelOpenInit(opts *bind.TransactOpts, version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "channelOpenInit", version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0x418925b7.
//
// Solidity: function channelOpenInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_Dispatcher *DispatcherSession) ChannelOpenInit(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChannelOpenInit(&_Dispatcher.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// ChannelOpenInit is a paid mutator transaction binding the contract method 0x418925b7.
//
// Solidity: function channelOpenInit(string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId) returns()
func (_Dispatcher *DispatcherTransactorSession) ChannelOpenInit(version string, ordering uint8, feeEnabled bool, connectionHops []string, counterpartyPortId string) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChannelOpenInit(&_Dispatcher.TransactOpts, version, ordering, feeEnabled, connectionHops, counterpartyPortId)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x2bf5d19d.
//
// Solidity: function channelOpenTry((string,bytes32,string) local, uint8 ordering, bool feeEnabled, string[] connectionHops, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactor) ChannelOpenTry(opts *bind.TransactOpts, local ChannelEnd, ordering uint8, feeEnabled bool, connectionHops []string, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "channelOpenTry", local, ordering, feeEnabled, connectionHops, counterparty, proof)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x2bf5d19d.
//
// Solidity: function channelOpenTry((string,bytes32,string) local, uint8 ordering, bool feeEnabled, string[] connectionHops, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherSession) ChannelOpenTry(local ChannelEnd, ordering uint8, feeEnabled bool, connectionHops []string, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChannelOpenTry(&_Dispatcher.TransactOpts, local, ordering, feeEnabled, connectionHops, counterparty, proof)
}

// ChannelOpenTry is a paid mutator transaction binding the contract method 0x2bf5d19d.
//
// Solidity: function channelOpenTry((string,bytes32,string) local, uint8 ordering, bool feeEnabled, string[] connectionHops, (string,bytes32,string) counterparty, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactorSession) ChannelOpenTry(local ChannelEnd, ordering uint8, feeEnabled bool, connectionHops []string, counterparty ChannelEnd, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChannelOpenTry(&_Dispatcher.TransactOpts, local, ordering, feeEnabled, connectionHops, counterparty, proof)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ab4339d.
//
// Solidity: function initialize(string initPortPrefix, address _feeVault) returns()
func (_Dispatcher *DispatcherTransactor) Initialize(opts *bind.TransactOpts, initPortPrefix string, _feeVault common.Address) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "initialize", initPortPrefix, _feeVault)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ab4339d.
//
// Solidity: function initialize(string initPortPrefix, address _feeVault) returns()
func (_Dispatcher *DispatcherSession) Initialize(initPortPrefix string, _feeVault common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.Initialize(&_Dispatcher.TransactOpts, initPortPrefix, _feeVault)
}

// Initialize is a paid mutator transaction binding the contract method 0x7ab4339d.
//
// Solidity: function initialize(string initPortPrefix, address _feeVault) returns()
func (_Dispatcher *DispatcherTransactorSession) Initialize(initPortPrefix string, _feeVault common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.Initialize(&_Dispatcher.TransactOpts, initPortPrefix, _feeVault)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x6b67055e.
//
// Solidity: function recvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactor) RecvPacket(opts *bind.TransactOpts, packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "recvPacket", packet, proof)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x6b67055e.
//
// Solidity: function recvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherSession) RecvPacket(packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.RecvPacket(&_Dispatcher.TransactOpts, packet, proof)
}

// RecvPacket is a paid mutator transaction binding the contract method 0x6b67055e.
//
// Solidity: function recvPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactorSession) RecvPacket(packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.RecvPacket(&_Dispatcher.TransactOpts, packet, proof)
}

// RemoveConnection is a paid mutator transaction binding the contract method 0xc00fa7c0.
//
// Solidity: function removeConnection(string connection) returns()
func (_Dispatcher *DispatcherTransactor) RemoveConnection(opts *bind.TransactOpts, connection string) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "removeConnection", connection)
}

// RemoveConnection is a paid mutator transaction binding the contract method 0xc00fa7c0.
//
// Solidity: function removeConnection(string connection) returns()
func (_Dispatcher *DispatcherSession) RemoveConnection(connection string) (*types.Transaction, error) {
	return _Dispatcher.Contract.RemoveConnection(&_Dispatcher.TransactOpts, connection)
}

// RemoveConnection is a paid mutator transaction binding the contract method 0xc00fa7c0.
//
// Solidity: function removeConnection(string connection) returns()
func (_Dispatcher *DispatcherTransactorSession) RemoveConnection(connection string) (*types.Transaction, error) {
	return _Dispatcher.Contract.RemoveConnection(&_Dispatcher.TransactOpts, connection)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dispatcher *DispatcherTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dispatcher *DispatcherSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dispatcher.Contract.RenounceOwnership(&_Dispatcher.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dispatcher *DispatcherTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dispatcher.Contract.RenounceOwnership(&_Dispatcher.TransactOpts)
}

// SendPacket is a paid mutator transaction binding the contract method 0xc3e1155c.
//
// Solidity: function sendPacket(bytes32 channelId, bytes packet, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_Dispatcher *DispatcherTransactor) SendPacket(opts *bind.TransactOpts, channelId [32]byte, packet []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "sendPacket", channelId, packet, timeoutTimestamp)
}

// SendPacket is a paid mutator transaction binding the contract method 0xc3e1155c.
//
// Solidity: function sendPacket(bytes32 channelId, bytes packet, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_Dispatcher *DispatcherSession) SendPacket(channelId [32]byte, packet []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Dispatcher.Contract.SendPacket(&_Dispatcher.TransactOpts, channelId, packet, timeoutTimestamp)
}

// SendPacket is a paid mutator transaction binding the contract method 0xc3e1155c.
//
// Solidity: function sendPacket(bytes32 channelId, bytes packet, uint64 timeoutTimestamp) returns(uint64 sequence)
func (_Dispatcher *DispatcherTransactorSession) SendPacket(channelId [32]byte, packet []byte, timeoutTimestamp uint64) (*types.Transaction, error) {
	return _Dispatcher.Contract.SendPacket(&_Dispatcher.TransactOpts, channelId, packet, timeoutTimestamp)
}

// SetClientForConnection is a paid mutator transaction binding the contract method 0x556d5178.
//
// Solidity: function setClientForConnection(string connection, address lightClient) returns()
func (_Dispatcher *DispatcherTransactor) SetClientForConnection(opts *bind.TransactOpts, connection string, lightClient common.Address) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "setClientForConnection", connection, lightClient)
}

// SetClientForConnection is a paid mutator transaction binding the contract method 0x556d5178.
//
// Solidity: function setClientForConnection(string connection, address lightClient) returns()
func (_Dispatcher *DispatcherSession) SetClientForConnection(connection string, lightClient common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.SetClientForConnection(&_Dispatcher.TransactOpts, connection, lightClient)
}

// SetClientForConnection is a paid mutator transaction binding the contract method 0x556d5178.
//
// Solidity: function setClientForConnection(string connection, address lightClient) returns()
func (_Dispatcher *DispatcherTransactorSession) SetClientForConnection(connection string, lightClient common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.SetClientForConnection(&_Dispatcher.TransactOpts, connection, lightClient)
}

// SetPortPrefix is a paid mutator transaction binding the contract method 0x9f59ae71.
//
// Solidity: function setPortPrefix(string _portPrefix) returns()
func (_Dispatcher *DispatcherTransactor) SetPortPrefix(opts *bind.TransactOpts, _portPrefix string) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "setPortPrefix", _portPrefix)
}

// SetPortPrefix is a paid mutator transaction binding the contract method 0x9f59ae71.
//
// Solidity: function setPortPrefix(string _portPrefix) returns()
func (_Dispatcher *DispatcherSession) SetPortPrefix(_portPrefix string) (*types.Transaction, error) {
	return _Dispatcher.Contract.SetPortPrefix(&_Dispatcher.TransactOpts, _portPrefix)
}

// SetPortPrefix is a paid mutator transaction binding the contract method 0x9f59ae71.
//
// Solidity: function setPortPrefix(string _portPrefix) returns()
func (_Dispatcher *DispatcherTransactorSession) SetPortPrefix(_portPrefix string) (*types.Transaction, error) {
	return _Dispatcher.Contract.SetPortPrefix(&_Dispatcher.TransactOpts, _portPrefix)
}

// Timeout is a paid mutator transaction binding the contract method 0x6050b5f3.
//
// Solidity: function timeout(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactor) Timeout(opts *bind.TransactOpts, packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "timeout", packet, proof)
}

// Timeout is a paid mutator transaction binding the contract method 0x6050b5f3.
//
// Solidity: function timeout(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherSession) Timeout(packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.Timeout(&_Dispatcher.TransactOpts, packet, proof)
}

// Timeout is a paid mutator transaction binding the contract method 0x6050b5f3.
//
// Solidity: function timeout(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactorSession) Timeout(packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.Timeout(&_Dispatcher.TransactOpts, packet, proof)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dispatcher *DispatcherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dispatcher *DispatcherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.TransferOwnership(&_Dispatcher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dispatcher *DispatcherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.TransferOwnership(&_Dispatcher.TransactOpts, newOwner)
}

// UpdateClientWithOptimisticConsensusState is a paid mutator transaction binding the contract method 0x940265cb.
//
// Solidity: function updateClientWithOptimisticConsensusState((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, uint256 height, uint256 appHash, string connection) returns(uint256 fraudProofEndTime, bool ended)
func (_Dispatcher *DispatcherTransactor) UpdateClientWithOptimisticConsensusState(opts *bind.TransactOpts, l1header L1Header, proof OpL2StateProof, height *big.Int, appHash *big.Int, connection string) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "updateClientWithOptimisticConsensusState", l1header, proof, height, appHash, connection)
}

// UpdateClientWithOptimisticConsensusState is a paid mutator transaction binding the contract method 0x940265cb.
//
// Solidity: function updateClientWithOptimisticConsensusState((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, uint256 height, uint256 appHash, string connection) returns(uint256 fraudProofEndTime, bool ended)
func (_Dispatcher *DispatcherSession) UpdateClientWithOptimisticConsensusState(l1header L1Header, proof OpL2StateProof, height *big.Int, appHash *big.Int, connection string) (*types.Transaction, error) {
	return _Dispatcher.Contract.UpdateClientWithOptimisticConsensusState(&_Dispatcher.TransactOpts, l1header, proof, height, appHash, connection)
}

// UpdateClientWithOptimisticConsensusState is a paid mutator transaction binding the contract method 0x940265cb.
//
// Solidity: function updateClientWithOptimisticConsensusState((bytes[],bytes32,uint64) l1header, (bytes[],bytes[],bytes32,bytes32) proof, uint256 height, uint256 appHash, string connection) returns(uint256 fraudProofEndTime, bool ended)
func (_Dispatcher *DispatcherTransactorSession) UpdateClientWithOptimisticConsensusState(l1header L1Header, proof OpL2StateProof, height *big.Int, appHash *big.Int, connection string) (*types.Transaction, error) {
	return _Dispatcher.Contract.UpdateClientWithOptimisticConsensusState(&_Dispatcher.TransactOpts, l1header, proof, height, appHash, connection)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Dispatcher *DispatcherTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Dispatcher *DispatcherSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.UpgradeTo(&_Dispatcher.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Dispatcher *DispatcherTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.UpgradeTo(&_Dispatcher.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Dispatcher *DispatcherTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Dispatcher *DispatcherSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Dispatcher.Contract.UpgradeToAndCall(&_Dispatcher.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Dispatcher *DispatcherTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Dispatcher.Contract.UpgradeToAndCall(&_Dispatcher.TransactOpts, newImplementation, data)
}

// WriteTimeoutPacket is a paid mutator transaction binding the contract method 0x5d7adf96.
//
// Solidity: function writeTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactor) WriteTimeoutPacket(opts *bind.TransactOpts, packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "writeTimeoutPacket", packet, proof)
}

// WriteTimeoutPacket is a paid mutator transaction binding the contract method 0x5d7adf96.
//
// Solidity: function writeTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherSession) WriteTimeoutPacket(packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.WriteTimeoutPacket(&_Dispatcher.TransactOpts, packet, proof)
}

// WriteTimeoutPacket is a paid mutator transaction binding the contract method 0x5d7adf96.
//
// Solidity: function writeTimeoutPacket(((string,bytes32),(string,bytes32),uint64,bytes,(uint64,uint64),uint64) packet, (((bytes,bytes)[],bytes,bytes,bytes)[],uint256) proof) returns()
func (_Dispatcher *DispatcherTransactorSession) WriteTimeoutPacket(packet IbcPacket, proof Ics23Proof) (*types.Transaction, error) {
	return _Dispatcher.Contract.WriteTimeoutPacket(&_Dispatcher.TransactOpts, packet, proof)
}

// DispatcherAcknowledgementIterator is returned from FilterAcknowledgement and is used to iterate over the raw logs and unpacked data for Acknowledgement events raised by the Dispatcher contract.
type DispatcherAcknowledgementIterator struct {
	Event *DispatcherAcknowledgement // Event containing the contract specifics and raw log

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
func (it *DispatcherAcknowledgementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherAcknowledgement)
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
		it.Event = new(DispatcherAcknowledgement)
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
func (it *DispatcherAcknowledgementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherAcknowledgementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherAcknowledgement represents a Acknowledgement event raised by the Dispatcher contract.
type DispatcherAcknowledgement struct {
	SourcePortAddress common.Address
	SourceChannelId   [32]byte
	Sequence          uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAcknowledgement is a free log retrieval operation binding the contract event 0xe46f6591236abe528fe47a3b281fb002524dadd3e62b1f317ed285d07273c3b1.
//
// Solidity: event Acknowledgement(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 sequence)
func (_Dispatcher *DispatcherFilterer) FilterAcknowledgement(opts *bind.FilterOpts, sourcePortAddress []common.Address, sourceChannelId [][32]byte) (*DispatcherAcknowledgementIterator, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "Acknowledgement", sourcePortAddressRule, sourceChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherAcknowledgementIterator{contract: _Dispatcher.contract, event: "Acknowledgement", logs: logs, sub: sub}, nil
}

// WatchAcknowledgement is a free log subscription operation binding the contract event 0xe46f6591236abe528fe47a3b281fb002524dadd3e62b1f317ed285d07273c3b1.
//
// Solidity: event Acknowledgement(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 sequence)
func (_Dispatcher *DispatcherFilterer) WatchAcknowledgement(opts *bind.WatchOpts, sink chan<- *DispatcherAcknowledgement, sourcePortAddress []common.Address, sourceChannelId [][32]byte) (event.Subscription, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "Acknowledgement", sourcePortAddressRule, sourceChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherAcknowledgement)
				if err := _Dispatcher.contract.UnpackLog(event, "Acknowledgement", log); err != nil {
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

// ParseAcknowledgement is a log parse operation binding the contract event 0xe46f6591236abe528fe47a3b281fb002524dadd3e62b1f317ed285d07273c3b1.
//
// Solidity: event Acknowledgement(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 sequence)
func (_Dispatcher *DispatcherFilterer) ParseAcknowledgement(log types.Log) (*DispatcherAcknowledgement, error) {
	event := new(DispatcherAcknowledgement)
	if err := _Dispatcher.contract.UnpackLog(event, "Acknowledgement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherAcknowledgementErrorIterator is returned from FilterAcknowledgementError and is used to iterate over the raw logs and unpacked data for AcknowledgementError events raised by the Dispatcher contract.
type DispatcherAcknowledgementErrorIterator struct {
	Event *DispatcherAcknowledgementError // Event containing the contract specifics and raw log

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
func (it *DispatcherAcknowledgementErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherAcknowledgementError)
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
		it.Event = new(DispatcherAcknowledgementError)
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
func (it *DispatcherAcknowledgementErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherAcknowledgementErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherAcknowledgementError represents a AcknowledgementError event raised by the Dispatcher contract.
type DispatcherAcknowledgementError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcknowledgementError is a free log retrieval operation binding the contract event 0x625eea143c9dae6915c809da47016c22d9cd006c3ace7c345c5cbcf57d3aefbc.
//
// Solidity: event AcknowledgementError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) FilterAcknowledgementError(opts *bind.FilterOpts, receiver []common.Address) (*DispatcherAcknowledgementErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "AcknowledgementError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherAcknowledgementErrorIterator{contract: _Dispatcher.contract, event: "AcknowledgementError", logs: logs, sub: sub}, nil
}

// WatchAcknowledgementError is a free log subscription operation binding the contract event 0x625eea143c9dae6915c809da47016c22d9cd006c3ace7c345c5cbcf57d3aefbc.
//
// Solidity: event AcknowledgementError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) WatchAcknowledgementError(opts *bind.WatchOpts, sink chan<- *DispatcherAcknowledgementError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "AcknowledgementError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherAcknowledgementError)
				if err := _Dispatcher.contract.UnpackLog(event, "AcknowledgementError", log); err != nil {
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

// ParseAcknowledgementError is a log parse operation binding the contract event 0x625eea143c9dae6915c809da47016c22d9cd006c3ace7c345c5cbcf57d3aefbc.
//
// Solidity: event AcknowledgementError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) ParseAcknowledgementError(log types.Log) (*DispatcherAcknowledgementError, error) {
	event := new(DispatcherAcknowledgementError)
	if err := _Dispatcher.contract.UnpackLog(event, "AcknowledgementError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Dispatcher contract.
type DispatcherAdminChangedIterator struct {
	Event *DispatcherAdminChanged // Event containing the contract specifics and raw log

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
func (it *DispatcherAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherAdminChanged)
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
		it.Event = new(DispatcherAdminChanged)
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
func (it *DispatcherAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherAdminChanged represents a AdminChanged event raised by the Dispatcher contract.
type DispatcherAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dispatcher *DispatcherFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*DispatcherAdminChangedIterator, error) {

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &DispatcherAdminChangedIterator{contract: _Dispatcher.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Dispatcher *DispatcherFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *DispatcherAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherAdminChanged)
				if err := _Dispatcher.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Dispatcher *DispatcherFilterer) ParseAdminChanged(log types.Log) (*DispatcherAdminChanged, error) {
	event := new(DispatcherAdminChanged)
	if err := _Dispatcher.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Dispatcher contract.
type DispatcherBeaconUpgradedIterator struct {
	Event *DispatcherBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *DispatcherBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherBeaconUpgraded)
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
		it.Event = new(DispatcherBeaconUpgraded)
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
func (it *DispatcherBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherBeaconUpgraded represents a BeaconUpgraded event raised by the Dispatcher contract.
type DispatcherBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Dispatcher *DispatcherFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*DispatcherBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherBeaconUpgradedIterator{contract: _Dispatcher.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Dispatcher *DispatcherFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *DispatcherBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherBeaconUpgraded)
				if err := _Dispatcher.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Dispatcher *DispatcherFilterer) ParseBeaconUpgraded(log types.Log) (*DispatcherBeaconUpgraded, error) {
	event := new(DispatcherBeaconUpgraded)
	if err := _Dispatcher.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherChannelCloseConfirmIterator is returned from FilterChannelCloseConfirm and is used to iterate over the raw logs and unpacked data for ChannelCloseConfirm events raised by the Dispatcher contract.
type DispatcherChannelCloseConfirmIterator struct {
	Event *DispatcherChannelCloseConfirm // Event containing the contract specifics and raw log

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
func (it *DispatcherChannelCloseConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherChannelCloseConfirm)
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
		it.Event = new(DispatcherChannelCloseConfirm)
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
func (it *DispatcherChannelCloseConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherChannelCloseConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherChannelCloseConfirm represents a ChannelCloseConfirm event raised by the Dispatcher contract.
type DispatcherChannelCloseConfirm struct {
	PortAddress common.Address
	ChannelId   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseConfirm is a free log retrieval operation binding the contract event 0x5f010dbbd6bf46aec8131c5456301a75cd688d3cca9bc8380c9e26301be20729.
//
// Solidity: event ChannelCloseConfirm(address indexed portAddress, bytes32 indexed channelId)
func (_Dispatcher *DispatcherFilterer) FilterChannelCloseConfirm(opts *bind.FilterOpts, portAddress []common.Address, channelId [][32]byte) (*DispatcherChannelCloseConfirmIterator, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "ChannelCloseConfirm", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherChannelCloseConfirmIterator{contract: _Dispatcher.contract, event: "ChannelCloseConfirm", logs: logs, sub: sub}, nil
}

// WatchChannelCloseConfirm is a free log subscription operation binding the contract event 0x5f010dbbd6bf46aec8131c5456301a75cd688d3cca9bc8380c9e26301be20729.
//
// Solidity: event ChannelCloseConfirm(address indexed portAddress, bytes32 indexed channelId)
func (_Dispatcher *DispatcherFilterer) WatchChannelCloseConfirm(opts *bind.WatchOpts, sink chan<- *DispatcherChannelCloseConfirm, portAddress []common.Address, channelId [][32]byte) (event.Subscription, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "ChannelCloseConfirm", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherChannelCloseConfirm)
				if err := _Dispatcher.contract.UnpackLog(event, "ChannelCloseConfirm", log); err != nil {
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

// ParseChannelCloseConfirm is a log parse operation binding the contract event 0x5f010dbbd6bf46aec8131c5456301a75cd688d3cca9bc8380c9e26301be20729.
//
// Solidity: event ChannelCloseConfirm(address indexed portAddress, bytes32 indexed channelId)
func (_Dispatcher *DispatcherFilterer) ParseChannelCloseConfirm(log types.Log) (*DispatcherChannelCloseConfirm, error) {
	event := new(DispatcherChannelCloseConfirm)
	if err := _Dispatcher.contract.UnpackLog(event, "ChannelCloseConfirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherChannelCloseConfirmErrorIterator is returned from FilterChannelCloseConfirmError and is used to iterate over the raw logs and unpacked data for ChannelCloseConfirmError events raised by the Dispatcher contract.
type DispatcherChannelCloseConfirmErrorIterator struct {
	Event *DispatcherChannelCloseConfirmError // Event containing the contract specifics and raw log

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
func (it *DispatcherChannelCloseConfirmErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherChannelCloseConfirmError)
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
		it.Event = new(DispatcherChannelCloseConfirmError)
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
func (it *DispatcherChannelCloseConfirmErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherChannelCloseConfirmErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherChannelCloseConfirmError represents a ChannelCloseConfirmError event raised by the Dispatcher contract.
type DispatcherChannelCloseConfirmError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseConfirmError is a free log retrieval operation binding the contract event 0xc9d36d7a317cb116925d5cb66f0069fe825822fe23e9cf3f421c38cf444caa30.
//
// Solidity: event ChannelCloseConfirmError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) FilterChannelCloseConfirmError(opts *bind.FilterOpts, receiver []common.Address) (*DispatcherChannelCloseConfirmErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "ChannelCloseConfirmError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherChannelCloseConfirmErrorIterator{contract: _Dispatcher.contract, event: "ChannelCloseConfirmError", logs: logs, sub: sub}, nil
}

// WatchChannelCloseConfirmError is a free log subscription operation binding the contract event 0xc9d36d7a317cb116925d5cb66f0069fe825822fe23e9cf3f421c38cf444caa30.
//
// Solidity: event ChannelCloseConfirmError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) WatchChannelCloseConfirmError(opts *bind.WatchOpts, sink chan<- *DispatcherChannelCloseConfirmError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "ChannelCloseConfirmError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherChannelCloseConfirmError)
				if err := _Dispatcher.contract.UnpackLog(event, "ChannelCloseConfirmError", log); err != nil {
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

// ParseChannelCloseConfirmError is a log parse operation binding the contract event 0xc9d36d7a317cb116925d5cb66f0069fe825822fe23e9cf3f421c38cf444caa30.
//
// Solidity: event ChannelCloseConfirmError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) ParseChannelCloseConfirmError(log types.Log) (*DispatcherChannelCloseConfirmError, error) {
	event := new(DispatcherChannelCloseConfirmError)
	if err := _Dispatcher.contract.UnpackLog(event, "ChannelCloseConfirmError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherChannelCloseInitIterator is returned from FilterChannelCloseInit and is used to iterate over the raw logs and unpacked data for ChannelCloseInit events raised by the Dispatcher contract.
type DispatcherChannelCloseInitIterator struct {
	Event *DispatcherChannelCloseInit // Event containing the contract specifics and raw log

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
func (it *DispatcherChannelCloseInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherChannelCloseInit)
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
		it.Event = new(DispatcherChannelCloseInit)
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
func (it *DispatcherChannelCloseInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherChannelCloseInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherChannelCloseInit represents a ChannelCloseInit event raised by the Dispatcher contract.
type DispatcherChannelCloseInit struct {
	PortAddress common.Address
	ChannelId   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseInit is a free log retrieval operation binding the contract event 0x21372e37743553ba8e5f61239803174a827c7732d53ec8adcb76c6b3bb2c13f1.
//
// Solidity: event ChannelCloseInit(address indexed portAddress, bytes32 indexed channelId)
func (_Dispatcher *DispatcherFilterer) FilterChannelCloseInit(opts *bind.FilterOpts, portAddress []common.Address, channelId [][32]byte) (*DispatcherChannelCloseInitIterator, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "ChannelCloseInit", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherChannelCloseInitIterator{contract: _Dispatcher.contract, event: "ChannelCloseInit", logs: logs, sub: sub}, nil
}

// WatchChannelCloseInit is a free log subscription operation binding the contract event 0x21372e37743553ba8e5f61239803174a827c7732d53ec8adcb76c6b3bb2c13f1.
//
// Solidity: event ChannelCloseInit(address indexed portAddress, bytes32 indexed channelId)
func (_Dispatcher *DispatcherFilterer) WatchChannelCloseInit(opts *bind.WatchOpts, sink chan<- *DispatcherChannelCloseInit, portAddress []common.Address, channelId [][32]byte) (event.Subscription, error) {

	var portAddressRule []interface{}
	for _, portAddressItem := range portAddress {
		portAddressRule = append(portAddressRule, portAddressItem)
	}
	var channelIdRule []interface{}
	for _, channelIdItem := range channelId {
		channelIdRule = append(channelIdRule, channelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "ChannelCloseInit", portAddressRule, channelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherChannelCloseInit)
				if err := _Dispatcher.contract.UnpackLog(event, "ChannelCloseInit", log); err != nil {
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

// ParseChannelCloseInit is a log parse operation binding the contract event 0x21372e37743553ba8e5f61239803174a827c7732d53ec8adcb76c6b3bb2c13f1.
//
// Solidity: event ChannelCloseInit(address indexed portAddress, bytes32 indexed channelId)
func (_Dispatcher *DispatcherFilterer) ParseChannelCloseInit(log types.Log) (*DispatcherChannelCloseInit, error) {
	event := new(DispatcherChannelCloseInit)
	if err := _Dispatcher.contract.UnpackLog(event, "ChannelCloseInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherChannelCloseInitErrorIterator is returned from FilterChannelCloseInitError and is used to iterate over the raw logs and unpacked data for ChannelCloseInitError events raised by the Dispatcher contract.
type DispatcherChannelCloseInitErrorIterator struct {
	Event *DispatcherChannelCloseInitError // Event containing the contract specifics and raw log

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
func (it *DispatcherChannelCloseInitErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherChannelCloseInitError)
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
		it.Event = new(DispatcherChannelCloseInitError)
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
func (it *DispatcherChannelCloseInitErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherChannelCloseInitErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherChannelCloseInitError represents a ChannelCloseInitError event raised by the Dispatcher contract.
type DispatcherChannelCloseInitError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelCloseInitError is a free log retrieval operation binding the contract event 0xb1be59c1bcd39c54c7132a8e0d321af5db427575ddb3265560d8862804f4381b.
//
// Solidity: event ChannelCloseInitError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) FilterChannelCloseInitError(opts *bind.FilterOpts, receiver []common.Address) (*DispatcherChannelCloseInitErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "ChannelCloseInitError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherChannelCloseInitErrorIterator{contract: _Dispatcher.contract, event: "ChannelCloseInitError", logs: logs, sub: sub}, nil
}

// WatchChannelCloseInitError is a free log subscription operation binding the contract event 0xb1be59c1bcd39c54c7132a8e0d321af5db427575ddb3265560d8862804f4381b.
//
// Solidity: event ChannelCloseInitError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) WatchChannelCloseInitError(opts *bind.WatchOpts, sink chan<- *DispatcherChannelCloseInitError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "ChannelCloseInitError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherChannelCloseInitError)
				if err := _Dispatcher.contract.UnpackLog(event, "ChannelCloseInitError", log); err != nil {
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

// ParseChannelCloseInitError is a log parse operation binding the contract event 0xb1be59c1bcd39c54c7132a8e0d321af5db427575ddb3265560d8862804f4381b.
//
// Solidity: event ChannelCloseInitError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) ParseChannelCloseInitError(log types.Log) (*DispatcherChannelCloseInitError, error) {
	event := new(DispatcherChannelCloseInitError)
	if err := _Dispatcher.contract.UnpackLog(event, "ChannelCloseInitError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherChannelOpenAckIterator is returned from FilterChannelOpenAck and is used to iterate over the raw logs and unpacked data for ChannelOpenAck events raised by the Dispatcher contract.
type DispatcherChannelOpenAckIterator struct {
	Event *DispatcherChannelOpenAck // Event containing the contract specifics and raw log

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
func (it *DispatcherChannelOpenAckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherChannelOpenAck)
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
		it.Event = new(DispatcherChannelOpenAck)
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
func (it *DispatcherChannelOpenAckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherChannelOpenAckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherChannelOpenAck represents a ChannelOpenAck event raised by the Dispatcher contract.
type DispatcherChannelOpenAck struct {
	Receiver  common.Address
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenAck is a free log retrieval operation binding the contract event 0xcf8be9ab2b5edf8beb2c45abe8e0cc7646318ac19f6c3164ba2e19e93a8a32af.
//
// Solidity: event ChannelOpenAck(address indexed receiver, bytes32 channelId)
func (_Dispatcher *DispatcherFilterer) FilterChannelOpenAck(opts *bind.FilterOpts, receiver []common.Address) (*DispatcherChannelOpenAckIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "ChannelOpenAck", receiverRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherChannelOpenAckIterator{contract: _Dispatcher.contract, event: "ChannelOpenAck", logs: logs, sub: sub}, nil
}

// WatchChannelOpenAck is a free log subscription operation binding the contract event 0xcf8be9ab2b5edf8beb2c45abe8e0cc7646318ac19f6c3164ba2e19e93a8a32af.
//
// Solidity: event ChannelOpenAck(address indexed receiver, bytes32 channelId)
func (_Dispatcher *DispatcherFilterer) WatchChannelOpenAck(opts *bind.WatchOpts, sink chan<- *DispatcherChannelOpenAck, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "ChannelOpenAck", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherChannelOpenAck)
				if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenAck", log); err != nil {
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

// ParseChannelOpenAck is a log parse operation binding the contract event 0xcf8be9ab2b5edf8beb2c45abe8e0cc7646318ac19f6c3164ba2e19e93a8a32af.
//
// Solidity: event ChannelOpenAck(address indexed receiver, bytes32 channelId)
func (_Dispatcher *DispatcherFilterer) ParseChannelOpenAck(log types.Log) (*DispatcherChannelOpenAck, error) {
	event := new(DispatcherChannelOpenAck)
	if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenAck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherChannelOpenAckErrorIterator is returned from FilterChannelOpenAckError and is used to iterate over the raw logs and unpacked data for ChannelOpenAckError events raised by the Dispatcher contract.
type DispatcherChannelOpenAckErrorIterator struct {
	Event *DispatcherChannelOpenAckError // Event containing the contract specifics and raw log

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
func (it *DispatcherChannelOpenAckErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherChannelOpenAckError)
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
		it.Event = new(DispatcherChannelOpenAckError)
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
func (it *DispatcherChannelOpenAckErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherChannelOpenAckErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherChannelOpenAckError represents a ChannelOpenAckError event raised by the Dispatcher contract.
type DispatcherChannelOpenAckError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenAckError is a free log retrieval operation binding the contract event 0x971a4433f5bff5f011728a4123aeeca4b5275ac20b013cf276e65510491ac26f.
//
// Solidity: event ChannelOpenAckError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) FilterChannelOpenAckError(opts *bind.FilterOpts, receiver []common.Address) (*DispatcherChannelOpenAckErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "ChannelOpenAckError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherChannelOpenAckErrorIterator{contract: _Dispatcher.contract, event: "ChannelOpenAckError", logs: logs, sub: sub}, nil
}

// WatchChannelOpenAckError is a free log subscription operation binding the contract event 0x971a4433f5bff5f011728a4123aeeca4b5275ac20b013cf276e65510491ac26f.
//
// Solidity: event ChannelOpenAckError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) WatchChannelOpenAckError(opts *bind.WatchOpts, sink chan<- *DispatcherChannelOpenAckError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "ChannelOpenAckError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherChannelOpenAckError)
				if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenAckError", log); err != nil {
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

// ParseChannelOpenAckError is a log parse operation binding the contract event 0x971a4433f5bff5f011728a4123aeeca4b5275ac20b013cf276e65510491ac26f.
//
// Solidity: event ChannelOpenAckError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) ParseChannelOpenAckError(log types.Log) (*DispatcherChannelOpenAckError, error) {
	event := new(DispatcherChannelOpenAckError)
	if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenAckError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherChannelOpenConfirmIterator is returned from FilterChannelOpenConfirm and is used to iterate over the raw logs and unpacked data for ChannelOpenConfirm events raised by the Dispatcher contract.
type DispatcherChannelOpenConfirmIterator struct {
	Event *DispatcherChannelOpenConfirm // Event containing the contract specifics and raw log

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
func (it *DispatcherChannelOpenConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherChannelOpenConfirm)
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
		it.Event = new(DispatcherChannelOpenConfirm)
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
func (it *DispatcherChannelOpenConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherChannelOpenConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherChannelOpenConfirm represents a ChannelOpenConfirm event raised by the Dispatcher contract.
type DispatcherChannelOpenConfirm struct {
	Receiver  common.Address
	ChannelId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenConfirm is a free log retrieval operation binding the contract event 0xe80f571f70f7cabf9d7ac60ece08421be374117776c311c327a083ca398f802f.
//
// Solidity: event ChannelOpenConfirm(address indexed receiver, bytes32 channelId)
func (_Dispatcher *DispatcherFilterer) FilterChannelOpenConfirm(opts *bind.FilterOpts, receiver []common.Address) (*DispatcherChannelOpenConfirmIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "ChannelOpenConfirm", receiverRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherChannelOpenConfirmIterator{contract: _Dispatcher.contract, event: "ChannelOpenConfirm", logs: logs, sub: sub}, nil
}

// WatchChannelOpenConfirm is a free log subscription operation binding the contract event 0xe80f571f70f7cabf9d7ac60ece08421be374117776c311c327a083ca398f802f.
//
// Solidity: event ChannelOpenConfirm(address indexed receiver, bytes32 channelId)
func (_Dispatcher *DispatcherFilterer) WatchChannelOpenConfirm(opts *bind.WatchOpts, sink chan<- *DispatcherChannelOpenConfirm, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "ChannelOpenConfirm", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherChannelOpenConfirm)
				if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenConfirm", log); err != nil {
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

// ParseChannelOpenConfirm is a log parse operation binding the contract event 0xe80f571f70f7cabf9d7ac60ece08421be374117776c311c327a083ca398f802f.
//
// Solidity: event ChannelOpenConfirm(address indexed receiver, bytes32 channelId)
func (_Dispatcher *DispatcherFilterer) ParseChannelOpenConfirm(log types.Log) (*DispatcherChannelOpenConfirm, error) {
	event := new(DispatcherChannelOpenConfirm)
	if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenConfirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherChannelOpenConfirmErrorIterator is returned from FilterChannelOpenConfirmError and is used to iterate over the raw logs and unpacked data for ChannelOpenConfirmError events raised by the Dispatcher contract.
type DispatcherChannelOpenConfirmErrorIterator struct {
	Event *DispatcherChannelOpenConfirmError // Event containing the contract specifics and raw log

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
func (it *DispatcherChannelOpenConfirmErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherChannelOpenConfirmError)
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
		it.Event = new(DispatcherChannelOpenConfirmError)
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
func (it *DispatcherChannelOpenConfirmErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherChannelOpenConfirmErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherChannelOpenConfirmError represents a ChannelOpenConfirmError event raised by the Dispatcher contract.
type DispatcherChannelOpenConfirmError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenConfirmError is a free log retrieval operation binding the contract event 0xf6a58ef30f66943749e8c29c661c84da143a1c8ed017f5faa92b509e0000875a.
//
// Solidity: event ChannelOpenConfirmError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) FilterChannelOpenConfirmError(opts *bind.FilterOpts, receiver []common.Address) (*DispatcherChannelOpenConfirmErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "ChannelOpenConfirmError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherChannelOpenConfirmErrorIterator{contract: _Dispatcher.contract, event: "ChannelOpenConfirmError", logs: logs, sub: sub}, nil
}

// WatchChannelOpenConfirmError is a free log subscription operation binding the contract event 0xf6a58ef30f66943749e8c29c661c84da143a1c8ed017f5faa92b509e0000875a.
//
// Solidity: event ChannelOpenConfirmError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) WatchChannelOpenConfirmError(opts *bind.WatchOpts, sink chan<- *DispatcherChannelOpenConfirmError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "ChannelOpenConfirmError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherChannelOpenConfirmError)
				if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenConfirmError", log); err != nil {
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

// ParseChannelOpenConfirmError is a log parse operation binding the contract event 0xf6a58ef30f66943749e8c29c661c84da143a1c8ed017f5faa92b509e0000875a.
//
// Solidity: event ChannelOpenConfirmError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) ParseChannelOpenConfirmError(log types.Log) (*DispatcherChannelOpenConfirmError, error) {
	event := new(DispatcherChannelOpenConfirmError)
	if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenConfirmError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherChannelOpenInitIterator is returned from FilterChannelOpenInit and is used to iterate over the raw logs and unpacked data for ChannelOpenInit events raised by the Dispatcher contract.
type DispatcherChannelOpenInitIterator struct {
	Event *DispatcherChannelOpenInit // Event containing the contract specifics and raw log

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
func (it *DispatcherChannelOpenInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherChannelOpenInit)
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
		it.Event = new(DispatcherChannelOpenInit)
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
func (it *DispatcherChannelOpenInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherChannelOpenInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherChannelOpenInit represents a ChannelOpenInit event raised by the Dispatcher contract.
type DispatcherChannelOpenInit struct {
	Receiver           common.Address
	Version            string
	Ordering           uint8
	FeeEnabled         bool
	ConnectionHops     []string
	CounterpartyPortId string
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenInit is a free log retrieval operation binding the contract event 0x20fd8a5856711b18d00def4aa6abafbe00ce6d60795e015cc1cad35eb9b46359.
//
// Solidity: event ChannelOpenInit(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId)
func (_Dispatcher *DispatcherFilterer) FilterChannelOpenInit(opts *bind.FilterOpts, receiver []common.Address) (*DispatcherChannelOpenInitIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "ChannelOpenInit", receiverRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherChannelOpenInitIterator{contract: _Dispatcher.contract, event: "ChannelOpenInit", logs: logs, sub: sub}, nil
}

// WatchChannelOpenInit is a free log subscription operation binding the contract event 0x20fd8a5856711b18d00def4aa6abafbe00ce6d60795e015cc1cad35eb9b46359.
//
// Solidity: event ChannelOpenInit(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId)
func (_Dispatcher *DispatcherFilterer) WatchChannelOpenInit(opts *bind.WatchOpts, sink chan<- *DispatcherChannelOpenInit, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "ChannelOpenInit", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherChannelOpenInit)
				if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenInit", log); err != nil {
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

// ParseChannelOpenInit is a log parse operation binding the contract event 0x20fd8a5856711b18d00def4aa6abafbe00ce6d60795e015cc1cad35eb9b46359.
//
// Solidity: event ChannelOpenInit(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId)
func (_Dispatcher *DispatcherFilterer) ParseChannelOpenInit(log types.Log) (*DispatcherChannelOpenInit, error) {
	event := new(DispatcherChannelOpenInit)
	if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherChannelOpenInitErrorIterator is returned from FilterChannelOpenInitError and is used to iterate over the raw logs and unpacked data for ChannelOpenInitError events raised by the Dispatcher contract.
type DispatcherChannelOpenInitErrorIterator struct {
	Event *DispatcherChannelOpenInitError // Event containing the contract specifics and raw log

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
func (it *DispatcherChannelOpenInitErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherChannelOpenInitError)
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
		it.Event = new(DispatcherChannelOpenInitError)
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
func (it *DispatcherChannelOpenInitErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherChannelOpenInitErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherChannelOpenInitError represents a ChannelOpenInitError event raised by the Dispatcher contract.
type DispatcherChannelOpenInitError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenInitError is a free log retrieval operation binding the contract event 0x69c1283cce89382f0f9ddf19b7c4f05b4d9b3c30c84fc148b1ec800284be58d5.
//
// Solidity: event ChannelOpenInitError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) FilterChannelOpenInitError(opts *bind.FilterOpts, receiver []common.Address) (*DispatcherChannelOpenInitErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "ChannelOpenInitError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherChannelOpenInitErrorIterator{contract: _Dispatcher.contract, event: "ChannelOpenInitError", logs: logs, sub: sub}, nil
}

// WatchChannelOpenInitError is a free log subscription operation binding the contract event 0x69c1283cce89382f0f9ddf19b7c4f05b4d9b3c30c84fc148b1ec800284be58d5.
//
// Solidity: event ChannelOpenInitError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) WatchChannelOpenInitError(opts *bind.WatchOpts, sink chan<- *DispatcherChannelOpenInitError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "ChannelOpenInitError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherChannelOpenInitError)
				if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenInitError", log); err != nil {
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

// ParseChannelOpenInitError is a log parse operation binding the contract event 0x69c1283cce89382f0f9ddf19b7c4f05b4d9b3c30c84fc148b1ec800284be58d5.
//
// Solidity: event ChannelOpenInitError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) ParseChannelOpenInitError(log types.Log) (*DispatcherChannelOpenInitError, error) {
	event := new(DispatcherChannelOpenInitError)
	if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenInitError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherChannelOpenTryIterator is returned from FilterChannelOpenTry and is used to iterate over the raw logs and unpacked data for ChannelOpenTry events raised by the Dispatcher contract.
type DispatcherChannelOpenTryIterator struct {
	Event *DispatcherChannelOpenTry // Event containing the contract specifics and raw log

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
func (it *DispatcherChannelOpenTryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherChannelOpenTry)
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
		it.Event = new(DispatcherChannelOpenTry)
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
func (it *DispatcherChannelOpenTryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherChannelOpenTryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherChannelOpenTry represents a ChannelOpenTry event raised by the Dispatcher contract.
type DispatcherChannelOpenTry struct {
	Receiver              common.Address
	Version               string
	Ordering              uint8
	FeeEnabled            bool
	ConnectionHops        []string
	CounterpartyPortId    string
	CounterpartyChannelId [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenTry is a free log retrieval operation binding the contract event 0xf910705a7a768eb5958f281a5f84cae8bffc5dd811ca5cd303dda140a423698c.
//
// Solidity: event ChannelOpenTry(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId, bytes32 counterpartyChannelId)
func (_Dispatcher *DispatcherFilterer) FilterChannelOpenTry(opts *bind.FilterOpts, receiver []common.Address) (*DispatcherChannelOpenTryIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "ChannelOpenTry", receiverRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherChannelOpenTryIterator{contract: _Dispatcher.contract, event: "ChannelOpenTry", logs: logs, sub: sub}, nil
}

// WatchChannelOpenTry is a free log subscription operation binding the contract event 0xf910705a7a768eb5958f281a5f84cae8bffc5dd811ca5cd303dda140a423698c.
//
// Solidity: event ChannelOpenTry(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId, bytes32 counterpartyChannelId)
func (_Dispatcher *DispatcherFilterer) WatchChannelOpenTry(opts *bind.WatchOpts, sink chan<- *DispatcherChannelOpenTry, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "ChannelOpenTry", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherChannelOpenTry)
				if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenTry", log); err != nil {
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

// ParseChannelOpenTry is a log parse operation binding the contract event 0xf910705a7a768eb5958f281a5f84cae8bffc5dd811ca5cd303dda140a423698c.
//
// Solidity: event ChannelOpenTry(address indexed receiver, string version, uint8 ordering, bool feeEnabled, string[] connectionHops, string counterpartyPortId, bytes32 counterpartyChannelId)
func (_Dispatcher *DispatcherFilterer) ParseChannelOpenTry(log types.Log) (*DispatcherChannelOpenTry, error) {
	event := new(DispatcherChannelOpenTry)
	if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenTry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherChannelOpenTryErrorIterator is returned from FilterChannelOpenTryError and is used to iterate over the raw logs and unpacked data for ChannelOpenTryError events raised by the Dispatcher contract.
type DispatcherChannelOpenTryErrorIterator struct {
	Event *DispatcherChannelOpenTryError // Event containing the contract specifics and raw log

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
func (it *DispatcherChannelOpenTryErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherChannelOpenTryError)
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
		it.Event = new(DispatcherChannelOpenTryError)
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
func (it *DispatcherChannelOpenTryErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherChannelOpenTryErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherChannelOpenTryError represents a ChannelOpenTryError event raised by the Dispatcher contract.
type DispatcherChannelOpenTryError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChannelOpenTryError is a free log retrieval operation binding the contract event 0x9e2fe55a3b54b57f82334c273f8d048cd7f05ad19c16cf334276a8c1fec4b6fd.
//
// Solidity: event ChannelOpenTryError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) FilterChannelOpenTryError(opts *bind.FilterOpts, receiver []common.Address) (*DispatcherChannelOpenTryErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "ChannelOpenTryError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherChannelOpenTryErrorIterator{contract: _Dispatcher.contract, event: "ChannelOpenTryError", logs: logs, sub: sub}, nil
}

// WatchChannelOpenTryError is a free log subscription operation binding the contract event 0x9e2fe55a3b54b57f82334c273f8d048cd7f05ad19c16cf334276a8c1fec4b6fd.
//
// Solidity: event ChannelOpenTryError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) WatchChannelOpenTryError(opts *bind.WatchOpts, sink chan<- *DispatcherChannelOpenTryError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "ChannelOpenTryError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherChannelOpenTryError)
				if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenTryError", log); err != nil {
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

// ParseChannelOpenTryError is a log parse operation binding the contract event 0x9e2fe55a3b54b57f82334c273f8d048cd7f05ad19c16cf334276a8c1fec4b6fd.
//
// Solidity: event ChannelOpenTryError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) ParseChannelOpenTryError(log types.Log) (*DispatcherChannelOpenTryError, error) {
	event := new(DispatcherChannelOpenTryError)
	if err := _Dispatcher.contract.UnpackLog(event, "ChannelOpenTryError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Dispatcher contract.
type DispatcherInitializedIterator struct {
	Event *DispatcherInitialized // Event containing the contract specifics and raw log

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
func (it *DispatcherInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherInitialized)
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
		it.Event = new(DispatcherInitialized)
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
func (it *DispatcherInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherInitialized represents a Initialized event raised by the Dispatcher contract.
type DispatcherInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dispatcher *DispatcherFilterer) FilterInitialized(opts *bind.FilterOpts) (*DispatcherInitializedIterator, error) {

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DispatcherInitializedIterator{contract: _Dispatcher.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Dispatcher *DispatcherFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DispatcherInitialized) (event.Subscription, error) {

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherInitialized)
				if err := _Dispatcher.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Dispatcher *DispatcherFilterer) ParseInitialized(log types.Log) (*DispatcherInitialized, error) {
	event := new(DispatcherInitialized)
	if err := _Dispatcher.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Dispatcher contract.
type DispatcherOwnershipTransferStartedIterator struct {
	Event *DispatcherOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *DispatcherOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherOwnershipTransferStarted)
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
		it.Event = new(DispatcherOwnershipTransferStarted)
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
func (it *DispatcherOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Dispatcher contract.
type DispatcherOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Dispatcher *DispatcherFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DispatcherOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherOwnershipTransferStartedIterator{contract: _Dispatcher.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Dispatcher *DispatcherFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *DispatcherOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherOwnershipTransferStarted)
				if err := _Dispatcher.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_Dispatcher *DispatcherFilterer) ParseOwnershipTransferStarted(log types.Log) (*DispatcherOwnershipTransferStarted, error) {
	event := new(DispatcherOwnershipTransferStarted)
	if err := _Dispatcher.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dispatcher contract.
type DispatcherOwnershipTransferredIterator struct {
	Event *DispatcherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DispatcherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherOwnershipTransferred)
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
		it.Event = new(DispatcherOwnershipTransferred)
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
func (it *DispatcherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherOwnershipTransferred represents a OwnershipTransferred event raised by the Dispatcher contract.
type DispatcherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dispatcher *DispatcherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DispatcherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherOwnershipTransferredIterator{contract: _Dispatcher.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dispatcher *DispatcherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DispatcherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherOwnershipTransferred)
				if err := _Dispatcher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dispatcher *DispatcherFilterer) ParseOwnershipTransferred(log types.Log) (*DispatcherOwnershipTransferred, error) {
	event := new(DispatcherOwnershipTransferred)
	if err := _Dispatcher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherRecvPacketIterator is returned from FilterRecvPacket and is used to iterate over the raw logs and unpacked data for RecvPacket events raised by the Dispatcher contract.
type DispatcherRecvPacketIterator struct {
	Event *DispatcherRecvPacket // Event containing the contract specifics and raw log

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
func (it *DispatcherRecvPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherRecvPacket)
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
		it.Event = new(DispatcherRecvPacket)
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
func (it *DispatcherRecvPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherRecvPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherRecvPacket represents a RecvPacket event raised by the Dispatcher contract.
type DispatcherRecvPacket struct {
	DestPortAddress common.Address
	DestChannelId   [32]byte
	Sequence        uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRecvPacket is a free log retrieval operation binding the contract event 0xde5b57e6566d68a30b0979431df3d5df6db3b9aa89f8820f595b9315bf86d067.
//
// Solidity: event RecvPacket(address indexed destPortAddress, bytes32 indexed destChannelId, uint64 sequence)
func (_Dispatcher *DispatcherFilterer) FilterRecvPacket(opts *bind.FilterOpts, destPortAddress []common.Address, destChannelId [][32]byte) (*DispatcherRecvPacketIterator, error) {

	var destPortAddressRule []interface{}
	for _, destPortAddressItem := range destPortAddress {
		destPortAddressRule = append(destPortAddressRule, destPortAddressItem)
	}
	var destChannelIdRule []interface{}
	for _, destChannelIdItem := range destChannelId {
		destChannelIdRule = append(destChannelIdRule, destChannelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "RecvPacket", destPortAddressRule, destChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherRecvPacketIterator{contract: _Dispatcher.contract, event: "RecvPacket", logs: logs, sub: sub}, nil
}

// WatchRecvPacket is a free log subscription operation binding the contract event 0xde5b57e6566d68a30b0979431df3d5df6db3b9aa89f8820f595b9315bf86d067.
//
// Solidity: event RecvPacket(address indexed destPortAddress, bytes32 indexed destChannelId, uint64 sequence)
func (_Dispatcher *DispatcherFilterer) WatchRecvPacket(opts *bind.WatchOpts, sink chan<- *DispatcherRecvPacket, destPortAddress []common.Address, destChannelId [][32]byte) (event.Subscription, error) {

	var destPortAddressRule []interface{}
	for _, destPortAddressItem := range destPortAddress {
		destPortAddressRule = append(destPortAddressRule, destPortAddressItem)
	}
	var destChannelIdRule []interface{}
	for _, destChannelIdItem := range destChannelId {
		destChannelIdRule = append(destChannelIdRule, destChannelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "RecvPacket", destPortAddressRule, destChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherRecvPacket)
				if err := _Dispatcher.contract.UnpackLog(event, "RecvPacket", log); err != nil {
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

// ParseRecvPacket is a log parse operation binding the contract event 0xde5b57e6566d68a30b0979431df3d5df6db3b9aa89f8820f595b9315bf86d067.
//
// Solidity: event RecvPacket(address indexed destPortAddress, bytes32 indexed destChannelId, uint64 sequence)
func (_Dispatcher *DispatcherFilterer) ParseRecvPacket(log types.Log) (*DispatcherRecvPacket, error) {
	event := new(DispatcherRecvPacket)
	if err := _Dispatcher.contract.UnpackLog(event, "RecvPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherSendPacketIterator is returned from FilterSendPacket and is used to iterate over the raw logs and unpacked data for SendPacket events raised by the Dispatcher contract.
type DispatcherSendPacketIterator struct {
	Event *DispatcherSendPacket // Event containing the contract specifics and raw log

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
func (it *DispatcherSendPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherSendPacket)
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
		it.Event = new(DispatcherSendPacket)
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
func (it *DispatcherSendPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherSendPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherSendPacket represents a SendPacket event raised by the Dispatcher contract.
type DispatcherSendPacket struct {
	SourcePortAddress common.Address
	SourceChannelId   [32]byte
	Packet            []byte
	Sequence          uint64
	TimeoutTimestamp  uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSendPacket is a free log retrieval operation binding the contract event 0xb5bff96e18da044e4e34510d16df9053b9f1920f6a960732e5aaf22fe9b80136.
//
// Solidity: event SendPacket(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, bytes packet, uint64 sequence, uint64 timeoutTimestamp)
func (_Dispatcher *DispatcherFilterer) FilterSendPacket(opts *bind.FilterOpts, sourcePortAddress []common.Address, sourceChannelId [][32]byte) (*DispatcherSendPacketIterator, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "SendPacket", sourcePortAddressRule, sourceChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherSendPacketIterator{contract: _Dispatcher.contract, event: "SendPacket", logs: logs, sub: sub}, nil
}

// WatchSendPacket is a free log subscription operation binding the contract event 0xb5bff96e18da044e4e34510d16df9053b9f1920f6a960732e5aaf22fe9b80136.
//
// Solidity: event SendPacket(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, bytes packet, uint64 sequence, uint64 timeoutTimestamp)
func (_Dispatcher *DispatcherFilterer) WatchSendPacket(opts *bind.WatchOpts, sink chan<- *DispatcherSendPacket, sourcePortAddress []common.Address, sourceChannelId [][32]byte) (event.Subscription, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "SendPacket", sourcePortAddressRule, sourceChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherSendPacket)
				if err := _Dispatcher.contract.UnpackLog(event, "SendPacket", log); err != nil {
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

// ParseSendPacket is a log parse operation binding the contract event 0xb5bff96e18da044e4e34510d16df9053b9f1920f6a960732e5aaf22fe9b80136.
//
// Solidity: event SendPacket(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, bytes packet, uint64 sequence, uint64 timeoutTimestamp)
func (_Dispatcher *DispatcherFilterer) ParseSendPacket(log types.Log) (*DispatcherSendPacket, error) {
	event := new(DispatcherSendPacket)
	if err := _Dispatcher.contract.UnpackLog(event, "SendPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherTimeoutIterator is returned from FilterTimeout and is used to iterate over the raw logs and unpacked data for Timeout events raised by the Dispatcher contract.
type DispatcherTimeoutIterator struct {
	Event *DispatcherTimeout // Event containing the contract specifics and raw log

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
func (it *DispatcherTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherTimeout)
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
		it.Event = new(DispatcherTimeout)
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
func (it *DispatcherTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherTimeout represents a Timeout event raised by the Dispatcher contract.
type DispatcherTimeout struct {
	SourcePortAddress common.Address
	SourceChannelId   [32]byte
	Sequence          uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTimeout is a free log retrieval operation binding the contract event 0x19ac40c4084d9bfb5b43f819a94bf01c70789b0d579871f59e4f86def04d9344.
//
// Solidity: event Timeout(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 indexed sequence)
func (_Dispatcher *DispatcherFilterer) FilterTimeout(opts *bind.FilterOpts, sourcePortAddress []common.Address, sourceChannelId [][32]byte, sequence []uint64) (*DispatcherTimeoutIterator, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "Timeout", sourcePortAddressRule, sourceChannelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherTimeoutIterator{contract: _Dispatcher.contract, event: "Timeout", logs: logs, sub: sub}, nil
}

// WatchTimeout is a free log subscription operation binding the contract event 0x19ac40c4084d9bfb5b43f819a94bf01c70789b0d579871f59e4f86def04d9344.
//
// Solidity: event Timeout(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 indexed sequence)
func (_Dispatcher *DispatcherFilterer) WatchTimeout(opts *bind.WatchOpts, sink chan<- *DispatcherTimeout, sourcePortAddress []common.Address, sourceChannelId [][32]byte, sequence []uint64) (event.Subscription, error) {

	var sourcePortAddressRule []interface{}
	for _, sourcePortAddressItem := range sourcePortAddress {
		sourcePortAddressRule = append(sourcePortAddressRule, sourcePortAddressItem)
	}
	var sourceChannelIdRule []interface{}
	for _, sourceChannelIdItem := range sourceChannelId {
		sourceChannelIdRule = append(sourceChannelIdRule, sourceChannelIdItem)
	}
	var sequenceRule []interface{}
	for _, sequenceItem := range sequence {
		sequenceRule = append(sequenceRule, sequenceItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "Timeout", sourcePortAddressRule, sourceChannelIdRule, sequenceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherTimeout)
				if err := _Dispatcher.contract.UnpackLog(event, "Timeout", log); err != nil {
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

// ParseTimeout is a log parse operation binding the contract event 0x19ac40c4084d9bfb5b43f819a94bf01c70789b0d579871f59e4f86def04d9344.
//
// Solidity: event Timeout(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 indexed sequence)
func (_Dispatcher *DispatcherFilterer) ParseTimeout(log types.Log) (*DispatcherTimeout, error) {
	event := new(DispatcherTimeout)
	if err := _Dispatcher.contract.UnpackLog(event, "Timeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherTimeoutErrorIterator is returned from FilterTimeoutError and is used to iterate over the raw logs and unpacked data for TimeoutError events raised by the Dispatcher contract.
type DispatcherTimeoutErrorIterator struct {
	Event *DispatcherTimeoutError // Event containing the contract specifics and raw log

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
func (it *DispatcherTimeoutErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherTimeoutError)
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
		it.Event = new(DispatcherTimeoutError)
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
func (it *DispatcherTimeoutErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherTimeoutErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherTimeoutError represents a TimeoutError event raised by the Dispatcher contract.
type DispatcherTimeoutError struct {
	Receiver common.Address
	Error    []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTimeoutError is a free log retrieval operation binding the contract event 0x83adb31803bee4e18cda1d04a781d77f6f271718a61b25e3a06f319b5103a330.
//
// Solidity: event TimeoutError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) FilterTimeoutError(opts *bind.FilterOpts, receiver []common.Address) (*DispatcherTimeoutErrorIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "TimeoutError", receiverRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherTimeoutErrorIterator{contract: _Dispatcher.contract, event: "TimeoutError", logs: logs, sub: sub}, nil
}

// WatchTimeoutError is a free log subscription operation binding the contract event 0x83adb31803bee4e18cda1d04a781d77f6f271718a61b25e3a06f319b5103a330.
//
// Solidity: event TimeoutError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) WatchTimeoutError(opts *bind.WatchOpts, sink chan<- *DispatcherTimeoutError, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "TimeoutError", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherTimeoutError)
				if err := _Dispatcher.contract.UnpackLog(event, "TimeoutError", log); err != nil {
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

// ParseTimeoutError is a log parse operation binding the contract event 0x83adb31803bee4e18cda1d04a781d77f6f271718a61b25e3a06f319b5103a330.
//
// Solidity: event TimeoutError(address indexed receiver, bytes error)
func (_Dispatcher *DispatcherFilterer) ParseTimeoutError(log types.Log) (*DispatcherTimeoutError, error) {
	event := new(DispatcherTimeoutError)
	if err := _Dispatcher.contract.UnpackLog(event, "TimeoutError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Dispatcher contract.
type DispatcherUpgradedIterator struct {
	Event *DispatcherUpgraded // Event containing the contract specifics and raw log

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
func (it *DispatcherUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherUpgraded)
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
		it.Event = new(DispatcherUpgraded)
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
func (it *DispatcherUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherUpgraded represents a Upgraded event raised by the Dispatcher contract.
type DispatcherUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dispatcher *DispatcherFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*DispatcherUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherUpgradedIterator{contract: _Dispatcher.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Dispatcher *DispatcherFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *DispatcherUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherUpgraded)
				if err := _Dispatcher.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Dispatcher *DispatcherFilterer) ParseUpgraded(log types.Log) (*DispatcherUpgraded, error) {
	event := new(DispatcherUpgraded)
	if err := _Dispatcher.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherWriteAckPacketIterator is returned from FilterWriteAckPacket and is used to iterate over the raw logs and unpacked data for WriteAckPacket events raised by the Dispatcher contract.
type DispatcherWriteAckPacketIterator struct {
	Event *DispatcherWriteAckPacket // Event containing the contract specifics and raw log

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
func (it *DispatcherWriteAckPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherWriteAckPacket)
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
		it.Event = new(DispatcherWriteAckPacket)
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
func (it *DispatcherWriteAckPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherWriteAckPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherWriteAckPacket represents a WriteAckPacket event raised by the Dispatcher contract.
type DispatcherWriteAckPacket struct {
	WriterPortAddress common.Address
	WriterChannelId   [32]byte
	Sequence          uint64
	AckPacket         AckPacket
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWriteAckPacket is a free log retrieval operation binding the contract event 0xa32e6f42b1d63fb83ad73b009a6dbb9413d1da02e95b1bb08f081815eea8db20.
//
// Solidity: event WriteAckPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (bool,bytes) ackPacket)
func (_Dispatcher *DispatcherFilterer) FilterWriteAckPacket(opts *bind.FilterOpts, writerPortAddress []common.Address, writerChannelId [][32]byte) (*DispatcherWriteAckPacketIterator, error) {

	var writerPortAddressRule []interface{}
	for _, writerPortAddressItem := range writerPortAddress {
		writerPortAddressRule = append(writerPortAddressRule, writerPortAddressItem)
	}
	var writerChannelIdRule []interface{}
	for _, writerChannelIdItem := range writerChannelId {
		writerChannelIdRule = append(writerChannelIdRule, writerChannelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "WriteAckPacket", writerPortAddressRule, writerChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherWriteAckPacketIterator{contract: _Dispatcher.contract, event: "WriteAckPacket", logs: logs, sub: sub}, nil
}

// WatchWriteAckPacket is a free log subscription operation binding the contract event 0xa32e6f42b1d63fb83ad73b009a6dbb9413d1da02e95b1bb08f081815eea8db20.
//
// Solidity: event WriteAckPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (bool,bytes) ackPacket)
func (_Dispatcher *DispatcherFilterer) WatchWriteAckPacket(opts *bind.WatchOpts, sink chan<- *DispatcherWriteAckPacket, writerPortAddress []common.Address, writerChannelId [][32]byte) (event.Subscription, error) {

	var writerPortAddressRule []interface{}
	for _, writerPortAddressItem := range writerPortAddress {
		writerPortAddressRule = append(writerPortAddressRule, writerPortAddressItem)
	}
	var writerChannelIdRule []interface{}
	for _, writerChannelIdItem := range writerChannelId {
		writerChannelIdRule = append(writerChannelIdRule, writerChannelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "WriteAckPacket", writerPortAddressRule, writerChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherWriteAckPacket)
				if err := _Dispatcher.contract.UnpackLog(event, "WriteAckPacket", log); err != nil {
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

// ParseWriteAckPacket is a log parse operation binding the contract event 0xa32e6f42b1d63fb83ad73b009a6dbb9413d1da02e95b1bb08f081815eea8db20.
//
// Solidity: event WriteAckPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (bool,bytes) ackPacket)
func (_Dispatcher *DispatcherFilterer) ParseWriteAckPacket(log types.Log) (*DispatcherWriteAckPacket, error) {
	event := new(DispatcherWriteAckPacket)
	if err := _Dispatcher.contract.UnpackLog(event, "WriteAckPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DispatcherWriteTimeoutPacketIterator is returned from FilterWriteTimeoutPacket and is used to iterate over the raw logs and unpacked data for WriteTimeoutPacket events raised by the Dispatcher contract.
type DispatcherWriteTimeoutPacketIterator struct {
	Event *DispatcherWriteTimeoutPacket // Event containing the contract specifics and raw log

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
func (it *DispatcherWriteTimeoutPacketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherWriteTimeoutPacket)
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
		it.Event = new(DispatcherWriteTimeoutPacket)
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
func (it *DispatcherWriteTimeoutPacketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherWriteTimeoutPacketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherWriteTimeoutPacket represents a WriteTimeoutPacket event raised by the Dispatcher contract.
type DispatcherWriteTimeoutPacket struct {
	WriterPortAddress common.Address
	WriterChannelId   [32]byte
	Sequence          uint64
	TimeoutHeight     Height
	TimeoutTimestamp  uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWriteTimeoutPacket is a free log retrieval operation binding the contract event 0xedbcd9eeb09d85c3ea1b5bf002c04478059cb261cab82c903885cefccae374bc.
//
// Solidity: event WriteTimeoutPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_Dispatcher *DispatcherFilterer) FilterWriteTimeoutPacket(opts *bind.FilterOpts, writerPortAddress []common.Address, writerChannelId [][32]byte) (*DispatcherWriteTimeoutPacketIterator, error) {

	var writerPortAddressRule []interface{}
	for _, writerPortAddressItem := range writerPortAddress {
		writerPortAddressRule = append(writerPortAddressRule, writerPortAddressItem)
	}
	var writerChannelIdRule []interface{}
	for _, writerChannelIdItem := range writerChannelId {
		writerChannelIdRule = append(writerChannelIdRule, writerChannelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "WriteTimeoutPacket", writerPortAddressRule, writerChannelIdRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherWriteTimeoutPacketIterator{contract: _Dispatcher.contract, event: "WriteTimeoutPacket", logs: logs, sub: sub}, nil
}

// WatchWriteTimeoutPacket is a free log subscription operation binding the contract event 0xedbcd9eeb09d85c3ea1b5bf002c04478059cb261cab82c903885cefccae374bc.
//
// Solidity: event WriteTimeoutPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_Dispatcher *DispatcherFilterer) WatchWriteTimeoutPacket(opts *bind.WatchOpts, sink chan<- *DispatcherWriteTimeoutPacket, writerPortAddress []common.Address, writerChannelId [][32]byte) (event.Subscription, error) {

	var writerPortAddressRule []interface{}
	for _, writerPortAddressItem := range writerPortAddress {
		writerPortAddressRule = append(writerPortAddressRule, writerPortAddressItem)
	}
	var writerChannelIdRule []interface{}
	for _, writerChannelIdItem := range writerChannelId {
		writerChannelIdRule = append(writerChannelIdRule, writerChannelIdItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "WriteTimeoutPacket", writerPortAddressRule, writerChannelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherWriteTimeoutPacket)
				if err := _Dispatcher.contract.UnpackLog(event, "WriteTimeoutPacket", log); err != nil {
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

// ParseWriteTimeoutPacket is a log parse operation binding the contract event 0xedbcd9eeb09d85c3ea1b5bf002c04478059cb261cab82c903885cefccae374bc.
//
// Solidity: event WriteTimeoutPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, (uint64,uint64) timeoutHeight, uint64 timeoutTimestamp)
func (_Dispatcher *DispatcherFilterer) ParseWriteTimeoutPacket(log types.Log) (*DispatcherWriteTimeoutPacket, error) {
	event := new(DispatcherWriteTimeoutPacket)
	if err := _Dispatcher.contract.UnpackLog(event, "WriteTimeoutPacket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
