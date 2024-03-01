import {IbcDispatcher, IbcEventsEmitter} from "./IbcDispatcher.sol";
import {L1Header, OpL2StateProof, Ics23Proof} from "./ProofVerifier.sol";
import {IbcChannelReceiver, IbcPacketReceiver} from "./IbcReceiver.sol";
import {
    Channel,
    CounterParty,
    ChannelOrder,
    IbcPacket,
    ChannelState,
    AckPacket,
    IBCErrors,
    IbcUtils,
    Ibc
} from "../libs/Ibc.sol";

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IDispatcher is IbcDispatcher, IbcEventsEmitter {
    //
    // Utility functions
    //
    function portIdAddressMatch(address addr, string calldata portId) external view returns (bool);

    //
    // CoreSC maaintainer methods, only invoked by the owner
    //
    function setPortPrefix(string calldata _portPrefix) external;

    function updateClientWithOptimisticConsensusState(
        L1Header calldata l1header,
        OpL2StateProof calldata proof,
        uint256 height,
        uint256 appHash
    ) external returns (uint256 fraudProofEndTime, bool ended);

    function getOptimisticConsensusState(uint256 height)
        external
        view
        returns (uint256 appHash, uint256 fraudProofEndTime, bool ended);

    //
    // IBC Channel methods
    //
    function openIbcChannel(
        IbcChannelReceiver portAddress,
        CounterParty calldata local,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        CounterParty calldata counterparty,
        Ics23Proof calldata proof
    ) external;

    function connectIbcChannel(
        IbcChannelReceiver portAddress,
        CounterParty calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        bool isChanConfirm,
        CounterParty calldata counterparty,
        Ics23Proof calldata proof
    ) external;

    function getChannel(address portAddress, bytes32 channelId) external view returns (Channel memory);

    function closeIbcChannel(bytes32 channelId) external;

    //
    // IBC Packet methods
    //
    function sendPacket(bytes32 channelId, bytes calldata packet, uint64 timeoutTimestamp) external;
    function sendPacketCommitment(address portAddress, bytes32 channelId, uint64 sequence) external returns (bool);

    function acknowledgement(
        IbcPacketReceiver receiver,
        IbcPacket calldata packet,
        bytes calldata ack,
        Ics23Proof calldata proof
    ) external;

    function timeout(IbcPacketReceiver receiver, IbcPacket calldata packet, Ics23Proof calldata proof) external;

    function recvPacket(IbcPacketReceiver receiver, IbcPacket calldata packet, Ics23Proof calldata proof) external;
}
