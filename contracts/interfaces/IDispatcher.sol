// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

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
import {LightClient} from "./LightClient.sol";

interface IDispatcher is IbcDispatcher, IbcEventsEmitter {
    function setPortPrefix(string calldata _portPrefix) external;

    function updateClientWithOptimisticConsensusState(
        L1Header calldata l1header,
        OpL2StateProof calldata proof,
        uint256 height,
        uint256 appHash,
        uint256 clientId
    ) external returns (uint256 fraudProofEndTime, bool ended);

    /**
     * This function is called by a 'relayer' on behalf of a dApp. The dApp should implement IbcChannelHandler's
     * onChanOpenInit. If the callback succeeds, the dApp should return the selected version and the emitted event
     * will be relayed to the  IBC/VIBC hub chain.
     */
    function channelOpenInit(
        IbcChannelReceiver portAddress,
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) external;

    function addNewConnection(string calldata connection, LightClient lightClient) external;

    /**
     * This function is called by a 'relayer' on behalf of a dApp. The dApp should implement IbcChannelHandler's
     * onChanOpenTry. If the callback succeeds, the dApp should return the selected version and the emitted event
     * will be relayed to the  IBC/VIBC hub chain.
     */
    function channelOpenTry(
        IbcChannelReceiver portAddress,
        CounterParty calldata local,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        CounterParty calldata counterparty,
        Ics23Proof calldata proof
    ) external;

    /**
     * This func is called by a 'relayer' after the IBC/VIBC hub chain has processed the ChannelOpenTry event.
     * The dApp should implement the onChannelConnect method to handle the third channel handshake method: ChanOpenAck
     */
    function channelOpenAck(
        IbcChannelReceiver portAddress,
        CounterParty calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        CounterParty calldata counterparty,
        Ics23Proof calldata proof
    ) external;

    /**
     * This func is called by a 'relayer' after the IBC/VIBC hub chain has processed the ChannelOpenTry event.
     * The dApp should implement the onChannelConnect method to handle the last channel handshake method:
     * ChannelOpenConfirm
     */
    function channelOpenConfirm(
        IbcChannelReceiver portAddress,
        CounterParty calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        CounterParty calldata counterparty,
        Ics23Proof calldata proof
    ) external;

    function channelCloseConfirm(address portAddress, bytes32 channelId, Ics23Proof calldata proof) external;
    function channelCloseInit(bytes32 channelId) external;

    function acknowledgement(
        IbcPacketReceiver receiver,
        IbcPacket calldata packet,
        bytes calldata ack,
        Ics23Proof calldata proof
    ) external;

    function timeout(IbcPacketReceiver receiver, IbcPacket calldata packet, Ics23Proof calldata proof) external;

    function recvPacket(IbcPacketReceiver receiver, IbcPacket calldata packet, Ics23Proof calldata proof) external;

    function getOptimisticConsensusState(uint256 height, uint256 clientId)
        external
        view
        returns (uint256 appHash, uint256 fraudProofEndTime, bool ended);

    function portIdAddressMatch(address addr, string calldata portId) external view returns (bool isMatch);

    function getChannel(address portAddress, bytes32 channelId) external view returns (Channel memory channel);
}
