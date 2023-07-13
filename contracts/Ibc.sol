//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

struct IbcPacket {
    /// identifies the channel and port on the sending chain.
    IbcEndpoint src;
    /// identifies the channel and port on the receiving chain.
    IbcEndpoint dest;
    /// The sequence number of the packet on the given channel
    uint64 sequence;
    bytes data;
    /// when packet times out, measured on remote chain
    /// Only inbound packets support both timeout height and timestamp.
    IbcTimeout timeout;
}

struct AckPacket {
    // success indicates the dApp-level logic. Even when a dApp fails to process a packet per its dApp logic, the
    // delivery of packet and ack packet are still considered successful.
    bool success;
    bytes data;
}

/// In IBC each package must set at least one type of timeout:
/// the timestamp or the block height.
struct IbcTimeout {
    uint64 blockHeight;
    uint64 timestamp;
}

enum ChannelOrder {
    NONE,
    UNORDERED,
    ORDERED
}

struct IbcEndpoint {
    string portId;
    bytes32 channelId;
}

struct PacketFee {
    uint256 recvFee;
    uint256 ackFee;
    uint256 timeoutFee;
}
