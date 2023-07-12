//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import './IbcDispatcher.sol';

struct IbcEndpoint {
    string portId;
    bytes32 channelId;
}

/// In IBC each package must set at least one type of timeout:
/// the timestamp or the block height.
struct IbcTimeout {
    uint64 blockHeight;
    uint64 timestamp;
}

struct IbcPacket {
    /// identifies the channel and port on the sending chain.
    IbcEndpoint src;
    /// identifies the channel and port on the receiving chain.
    IbcEndpoint dest;
    /// The sequence number of the packet on the given channel
    uint64 sequence;
    bytes data;
    /// when packet times out, measured on remote chain
    IbcTimeout timeout;
}

/**
 * @title IbcReceiver
 * @author Polymer Labs
 * @notice IBC receiver interface must be implemented by a IBC-enabled contract.
 * The implementer, aka. dApp devs, should implement channel handshake and packet handling methods.
 */
interface IbcReceiver {
    //
    // Packet handling methods
    //

    function onRecvPacket(IbcPacket calldata packet) external returns (AckPacket memory ackPacket);

    function onAcknowledgementPacket(IbcPacket calldata packet) external;

    function onTimeoutPacket(IbcPacket calldata packet) external;

    //
    // Channel handshake methods
    //

    function onOpenIbcChannel(
        string calldata version,
        ChannelOrder ordering,
        string[] calldata connectionHops,
        string calldata counterpartyPortId,
        string calldata counterPartyChannelId,
        string calldata counterpartyVersion
    ) external returns (string memory selectedVersion);

    function onConnectIbcChannel(
        bytes32 channelId,
        bytes32 counterpartyChannelId,
        bytes32 counterpartyVersion
    ) external;

    function onCloseIbcChannel(
        bytes32 channelId,
        string calldata counterpartyPortId,
        bytes32 counterpartyChannelId
    ) external;
}
