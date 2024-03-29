//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {ProtoChannel, ProtoCounterparty} from "proto/channel.sol";
import {Base64} from "base64/base64.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {IBCErrors} from "./IBCErrors.sol";

/**
 * Ibc.sol
 * Basic IBC data structures and utilities.
 */

/// IbcPacket represents the packet data structure received from a remote chain
/// over an IBC channel.
struct IbcPacket {
    /// identifies the channel and port on the sending chain.
    IbcEndpoint src;
    /// identifies the channel and port on the receiving chain.
    IbcEndpoint dest;
    /// The sequence number of the packet on the given channel
    uint64 sequence;
    bytes data;
    /// block height after which the packet times out
    Height timeoutHeight;
    /// block timestamp (in seconds after the unix epoch) after which the packet times out
    uint64 timeoutTimestamp;
}

// UniversalPacke represents the data field of an IbcPacket
struct UniversalPacket {
    bytes32 srcPortAddr;
    // source middleware ids bitmap, ie. logic OR of all MW IDs in the MW stack.
    uint256 mwBitmap;
    bytes32 destPortAddr;
    bytes appData;
}

/// Height is a monotonically increasing data type
/// that can be compared against another Height for the purposes of updating and
/// freezing clients
///
/// Normally the RevisionHeight is incremented at each height while keeping
/// RevisionNumber the same. However some consensus algorithms may choose to
/// reset the height in certain conditions e.g. hard forks, state-machine
/// breaking changes In these cases, the RevisionNumber is incremented so that
/// height continues to be monitonically increasing even as the RevisionHeight
/// gets reset
struct Height {
    uint64 revision_number;
    uint64 revision_height;
}

struct AckPacket {
    // success indicates the dApp-level logic. Even when a dApp fails to process a packet per its dApp logic, the
    // delivery of packet and ack packet are still considered successful.
    bool success;
    bytes data;
}

struct IncentivizedAckPacket {
    bool success;
    // Forward relayer's payee address, an EMV address registered on Polymer chain with `RegisterCounterpartyPayee`
    // endpoint.
    // In case of missing payee, zero address is used on Polymer.
    // The relayer payee address is set when incentivized ack is created on Polymer.
    bytes relayer;
    bytes data;
}

enum ChannelOrder {
    NONE,
    UNORDERED,
    ORDERED
}

enum ChannelState {
    // Default State
    UNINITIALIZED,
    // A channel has just started the opening handshake.
    INIT,
    // A channel has acknowledged the handshake step on the counterparty chain.
    TRYOPEN,
    // A channel has completed the handshake. Open channels are
    // ready to send and receive packets.
    OPEN,
    // A channel has been closed and can no longer be used to send or receive
    // packets.
    CLOSED,
    // A channel has been forced closed due to a frozen client in the connection
    // path.
    FROZEN,
    // A channel has acknowledged the handshake step on the counterparty chain, but not yet confirmed with a virtual
    // chain. Virtual channel end ONLY.
    TRY_PENDING,
    // A channel has finished the ChanOpenAck handshake step on chain A, but not yet confirmed with the corresponding
    // virtual chain. Virtual channel end ONLY.
    ACK_PENDING,
    // A channel has finished the ChanOpenConfirm handshake step on chain B, but not yet confirmed with the
    // corresponding
    // virtual chain. Virtual channel end ONLY.
    CONFIRM_PENDING,
    // A channel has finished the ChanCloseConfirm step on chainB, but not yet confirmed with the corresponding
    // virtual chain. Virtual channel end ONLY.
    CLOSE_CONFIRM_PENDING
}

struct Channel {
    string portId;
    string version;
    ChannelOrder ordering;
    bool feeEnabled;
    string[] connectionHops;
    string counterpartyPortId;
    bytes32 counterpartyChannelId;
}

struct ChannelEnd {
    string portId;
    bytes32 channelId;
    string version;
}

struct IbcEndpoint {
    string portId;
    bytes32 channelId;
}

struct Proof {
    // block height at which the proof is valid for a membership or non-membership at the given keyPath
    Height proofHeight;
    // ics23 merkle proof
    bytes proof;
}

library Ibc {
    /**
     * Convert a non-0x-prefixed hex string to an address
     * @param hexStr hex string to convert to address. Note that the hex string must not include a 0x prefix.
     * hexStr is case-insensitive.
     */
    function _hexStrToAddress(string memory hexStr) external pure returns (address addr) {
        if (bytes(hexStr).length != 40) {
            revert IBCErrors.invalidHexStringLength();
        }

        bytes memory strBytes = bytes(hexStr);
        bytes memory addrBytes = new bytes(20);

        for (uint256 i = 0; i < 20; i++) {
            uint8 high = uint8(strBytes[i * 2]);
            uint8 low = uint8(strBytes[1 + i * 2]);
            // Convert to lowercase if the character is in uppercase
            if (high >= 65 && high <= 90) {
                high += 32;
            }
            if (low >= 65 && low <= 90) {
                low += 32;
            }
            uint8 digit = (high - (high >= 97 ? 87 : 48)) * 16 + (low - (low >= 97 ? 87 : 48));
            addrBytes[i] = bytes1(digit);
        }

        assembly {
            addr := mload(add(addrBytes, 20))
        }
    }

    // https://github.com/open-ibc/ibcx-go/blob/ef80dd6784fd/modules/core/24-host/keys.go#L135
    function channelProofKey(string calldata portId, bytes32 channelId) external pure returns (bytes memory proofKey) {
        proofKey = abi.encodePacked("channelEnds/ports/", portId, "/channels/", toStr(channelId));
    }

    function channelProofKeyMemory(string memory portId, bytes32 channelId)
        external
        pure
        returns (bytes memory proofKey)
    {
        proofKey = abi.encodePacked("channelEnds/ports/", portId, "/channels/", toStr(channelId));
    }

    // protobuf encoding of a channel object
    // https://github.com/open-ibc/ibcx-go/blob/ef80dd6784fd/modules/core/04-channel/keeper/keeper.go#L92
    function channelProofValue(
        ChannelState state,
        ChannelOrder ordering,
        string calldata version,
        string[] calldata connectionHops,
        string calldata counterpartyPortId,
        bytes32 counterpartyChannelId
    ) external pure returns (bytes memory proofValue) {
        proofValue = ProtoChannel.encode(
            ProtoChannel.Data(
                int32(uint32(state)),
                int32(uint32(ordering)),
                ProtoCounterparty.Data(counterpartyPortId, toStr(counterpartyChannelId)),
                connectionHops,
                version
            )
        );
    }

    function channelProofValueMemory(
        ChannelState state,
        ChannelOrder ordering,
        string memory version,
        string[] memory connectionHops,
        string memory counterpartyPortId,
        bytes32 counterpartyChannelId
    ) external pure returns (bytes memory proofValue) {
        proofValue = ProtoChannel.encode(
            ProtoChannel.Data(
                int32(uint32(state)),
                int32(uint32(ordering)),
                ProtoCounterparty.Data(counterpartyPortId, toStr(counterpartyChannelId)),
                connectionHops,
                version
            )
        );
    }

    // https://github.com/open-ibc/ibcx-go/blob/ef80dd6784fd/modules/core/24-host/keys.go#L185
    function packetCommitmentProofKey(IbcPacket calldata packet) external pure returns (bytes memory proofKey) {
        proofKey = abi.encodePacked(
            "commitments/ports/",
            packet.src.portId,
            "/channels/",
            toStr(packet.src.channelId),
            "/sequences/",
            toStr(packet.sequence)
        );
    }

    // https://github.com/open-ibc/ibcx-go/blob/ef80dd6784fd/modules/core/04-channel/types/packet.go#L19
    function packetCommitmentProofValue(IbcPacket calldata packet) external pure returns (bytes32 proofValue) {
        proofValue = sha256(
            abi.encodePacked(
                packet.timeoutTimestamp,
                packet.timeoutHeight.revision_number,
                packet.timeoutHeight.revision_height,
                sha256(packet.data)
            )
        );
    }

    // https://github.com/open-ibc/ibcx-go/blob/ef80dd6784fd/modules/core/24-host/keys.go#L201
    function ackProofKey(IbcPacket calldata packet) external pure returns (bytes memory proofKey) {
        proofKey = abi.encodePacked(
            "acks/ports/",
            packet.dest.portId,
            "/channels/",
            toStr(packet.dest.channelId),
            "/sequences/",
            toStr(packet.sequence)
        );
    }

    // https://github.com/open-ibc/ibcx-go/blob/ef80dd6784fd/modules/core/04-channel/types/packet.go#L38
    function ackProofValue(bytes calldata ack) external pure returns (bytes32 proofValue) {
        proofValue = sha256(ack);
    }

    function parseAckData(bytes calldata ack) external pure returns (AckPacket memory ackData) {
        // this hex value is '"result"'
        ackData = (keccak256(ack[1:9]) == keccak256(hex"22726573756c7422"))
            ? AckPacket(true, Base64.decode(string(ack[11:ack.length - 2]))) // result success
            : AckPacket(false, ack[10:ack.length - 2]); // this is an error
    }

    function toStr(bytes32 b) public pure returns (string memory outStr) {
        uint8 i = 0;
        while (i < 32 && b[i] != 0) {
            i++;
        }
        bytes memory bytesArray = new bytes(i);
        for (uint8 j = 0; j < i; j++) {
            bytesArray[j] = b[j];
        }
        outStr = string(bytesArray);
    }

    function toStr(uint256 _number) public pure returns (string memory outStr) {
        if (_number == 0) {
            return "0";
        }

        uint256 length;
        uint256 number = _number;

        // Determine the length of the string
        while (number != 0) {
            length++;
            number /= 10;
        }

        bytes memory buffer = new bytes(length);

        // Convert each digit to its ASCII representation
        for (uint256 i = length; i > 0; i--) {
            buffer[i - 1] = bytes1(uint8(48 + (_number % 10)));
            _number /= 10;
        }

        outStr = string(buffer);
    }
}
