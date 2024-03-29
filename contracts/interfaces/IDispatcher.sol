// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IbcDispatcher, IbcEventsEmitter} from "./IbcDispatcher.sol";

import {L1Header, OpL2StateProof, Ics23Proof} from "./ProofVerifier.sol";
import {IbcChannelReceiver, IbcPacketReceiver} from "./IbcReceiver.sol";
import {
    Channel,
    ChannelEnd,
    ChannelOrder,
    IbcPacket,
    ChannelState,
    AckPacket,
    IBCErrors,
    IbcUtils,
    Ibc
} from "../libs/Ibc.sol";
import {LightClient} from "./LightClient.sol";

interface IDispatcher is IbcDispatcher, IbcEventsEmitter {
    function setPortPrefix(string calldata _portPrefix) external;

    function updateClientWithOptimisticConsensusState(
        L1Header calldata l1header,
        OpL2StateProof calldata proof,
        uint256 height,
        uint256 appHash,
        string calldata connection
    ) external returns (uint256 fraudProofEndTime, bool ended);

    /**
     * This function is called by a 'relayer' on behalf of a dApp. The dApp should implement IbcChannelHandler's
     * onChanOpenInit. If the callback succeeds, the dApp should return the selected version and the emitted event
     * will be relayed to the  IBC/VIBC hub chain.
     */
    function channelOpenInit(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) external;

    function setNewConnection(string calldata connection, LightClient lightClient) external;

    function removeConnection(string calldata connection) external;

    /**
     * This function is called by a 'relayer' on behalf of a dApp. The dApp should implement IbcChannelHandler's
     * onChanOpenTry. If the callback succeeds, the dApp should return the selected version and the emitted event
     * will be relayed to the  IBC/VIBC hub chain.
     */
    function channelOpenTry(
        ChannelEnd calldata local,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        ChannelEnd calldata counterparty,
        Ics23Proof calldata proof
    ) external;

    /**
     * This func is called by a 'relayer' after the IBC/VIBC hub chain has processed the ChannelOpenTry event.
     * The dApp should implement the onChannelConnect method to handle the third channel handshake method: ChanOpenAck
     */
    function channelOpenAck(
        ChannelEnd calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        ChannelEnd calldata counterparty,
        Ics23Proof calldata proof
    ) external;

    /**
     * This func is called by a 'relayer' after the IBC/VIBC hub chain has processed the ChannelOpenTry event.
     * The dApp should implement the onChannelConnect method to handle the last channel handshake method:
     * ChannelOpenConfirm
     */
    function channelOpenConfirm(
        ChannelEnd calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        ChannelEnd calldata counterparty,
        Ics23Proof calldata proof
    ) external;

    function channelCloseConfirm(address portAddress, bytes32 channelId, Ics23Proof calldata proof) external;
    function channelCloseInit(bytes32 channelId) external;

    function sendPacket(bytes32 channelId, bytes calldata packet, uint64 timeoutTimestamp) external;

    function acknowledgement(IbcPacket calldata packet, bytes calldata ack, Ics23Proof calldata proof) external;

    function timeout(IbcPacket calldata packet, Ics23Proof calldata proof) external;

    function writeTimeoutPacket(IbcPacket calldata packet, Ics23Proof calldata proof) external;

    function recvPacket(IbcPacket calldata packet, Ics23Proof calldata proof) external;

    function getOptimisticConsensusState(uint256 height, string calldata connection)
        external
        view
        returns (uint256 appHash, uint256 fraudProofEndTime, bool ended);

    function getChannel(address portAddress, bytes32 channelId) external view returns (Channel memory channel);
}
