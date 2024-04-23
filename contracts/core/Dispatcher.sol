// SPDX-License-Identifier: Apache-2.0
/*
 * Copyright 2024, Polymer Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.8.9;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgradeable/contracts/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {IbcChannelReceiver, IbcPacketReceiver} from "../interfaces/IbcReceiver.sol";
import {L1Header, OpL2StateProof, Ics23Proof} from "../interfaces/ProofVerifier.sol";
import {LightClient} from "../interfaces/LightClient.sol";
import {IDispatcher} from "../interfaces/IDispatcher.sol";
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
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @title Dispatcher
 * @author Polymer Labs
 * @notice
 *     Contract callers call this contract to send IBC-like msg,
 *     which can be relayed to a rollup module on the Polymerase chain
 */
contract Dispatcher is OwnableUpgradeable, UUPSUpgradeable, IDispatcher {
    uint256[49] private __gap;

    //
    // fields
    //
    // IBC_PortID = portPrefix + address (hex string without 0x prefix, case insensitive)
    string public portPrefix;
    uint32 public portPrefixLen;

    mapping(address => mapping(bytes32 => Channel)) private _portChannelMap;
    mapping(address => mapping(bytes32 => uint64)) private _nextSequenceSend;
    // keep track of received packets' sequences to ensure channel ordering is enforced for ordered channels
    mapping(address => mapping(bytes32 => uint64)) private _nextSequenceRecv;
    mapping(address => mapping(bytes32 => uint64)) private _nextSequenceAck;
    // only stores a bit to mark packet has not been ack'ed or timed out yet; actual IBC packet verification is done on
    // Polymer chain.
    // Keep track of sent packets
    mapping(address => mapping(bytes32 => mapping(uint64 => bool))) private _sendPacketCommitment;
    // keep track of received packets to prevent replay attack
    mapping(address => mapping(bytes32 => mapping(uint64 => bool))) private _recvPacketReceipt;
    // keep track of outbound ack packets to prevent replay attack
    mapping(address => mapping(bytes32 => mapping(uint64 => bool))) private _ackPacketCommitment;

    LightClient _lightClient; // Can't be set to immutable since it needs to be called in the initializer; not the

    //////// NEW storage
    mapping(bytes32 => string) private _channelIdToConnection;

    // constructor

    //
    // methods
    //
    constructor() {
        _disableInitializers();
    }

    function initialize(string memory initPortPrefix, LightClient lightClient) public initializer {
        __Ownable_init();
        portPrefix = initPortPrefix;
        portPrefixLen = uint32(bytes(initPortPrefix).length);
        _lightClient = lightClient;
    }

    //
    // CoreSC maaintainer methods, only invoked by the owner
    //
    function setPortPrefix(string calldata _portPrefix) external onlyOwner {
        portPrefix = _portPrefix;
        portPrefixLen = uint32(bytes(_portPrefix).length);
    }

    // updateClientWithOptimisticConsensusState updates the client
    // with the optimistic consensus state. The optimistic consensus
    // is accepted and will be open for verify in the fraud proof
    // window.
    function updateClientWithOptimisticConsensusState(
        L1Header calldata l1header,
        OpL2StateProof calldata proof,
        uint256 height,
        uint256 appHash
    ) external returns (uint256 fraudProofEndTime, bool ended) {
        return _lightClient.addOpConsensusState(l1header, proof, height, appHash);
    }

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
    ) external {
        if (bytes(counterpartyPortId).length == 0) {
            revert IBCErrors.invalidCounterPartyPortId();
        }

        // Have to encode here to avoid stack-too-deep error
        bytes memory chanOpenInitArgs = abi.encode(ordering, connectionHops, counterpartyPortId, version);
        (bool success, bytes memory data) =
            _callIfContract(msg.sender, bytes.concat(IbcChannelReceiver.onChanOpenInit.selector, chanOpenInitArgs));

        if (success) {
            emit ChannelOpenInit(
                msg.sender, abi.decode(data, (string)), ordering, feeEnabled, connectionHops, counterpartyPortId
            );
        } else {
            emit ChannelOpenInitError(msg.sender, data);
        }
    }

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
    ) external {
        if (bytes(counterparty.portId).length == 0) {
            revert IBCErrors.invalidCounterPartyPortId();
        }

        _lightClient.verifyMembership(
            proof,
            Ibc.channelProofKey(local.portId, local.channelId),
            Ibc.channelProofValue(ChannelState.TRY_PENDING, ordering, local.version, connectionHops, counterparty)
        );

        address receiver = _getAddressFromPort(local.portId);
        (bool success, bytes memory data) = _callIfContract(
            receiver,
            abi.encodeWithSelector(
                IbcChannelReceiver.onChanOpenTry.selector,
                ordering,
                connectionHops,
                local.channelId,
                counterparty.portId,
                counterparty.channelId,
                counterparty.version
            )
        );

        if (success) {
            emit ChannelOpenTry(
                receiver,
                abi.decode(data, (string)),
                ordering,
                feeEnabled,
                connectionHops,
                counterparty.portId,
                counterparty.channelId
            );
        } else {
            emit ChannelOpenTryError(receiver, data);
        }
    }

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
    ) external {
        _lightClient.verifyMembership(
            proof,
            Ibc.channelProofKey(local.portId, local.channelId),
            Ibc.channelProofValue(ChannelState.ACK_PENDING, ordering, local.version, connectionHops, counterparty)
        );

        address receiver = _getAddressFromPort(local.portId);
        (bool success, bytes memory data) = _callIfContract(
            receiver,
            abi.encodeWithSelector(
                IbcChannelReceiver.onChanOpenAck.selector, local.channelId, counterparty.channelId, counterparty.version
            )
        );

        if (success) {
            _connectChannel(IbcChannelReceiver(receiver), local, connectionHops, ordering, feeEnabled, counterparty);
            emit ChannelOpenAck(receiver, local.channelId);
        } else {
            emit ChannelOpenAckError(receiver, data);
        }
    }

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
    ) external {
        _lightClient.verifyMembership(
            proof,
            Ibc.channelProofKey(local.portId, local.channelId),
            Ibc.channelProofValue(ChannelState.CONFIRM_PENDING, ordering, local.version, connectionHops, counterparty)
        );

        address receiver = _getAddressFromPort(local.portId);
        (bool success, bytes memory data) = _callIfContract(
            receiver,
            abi.encodeWithSelector(IbcChannelReceiver.onChanOpenConfirm.selector, local.channelId, counterparty.version)
        );

        if (success) {
            _connectChannel(IbcChannelReceiver(receiver), local, connectionHops, ordering, feeEnabled, counterparty);
            emit ChannelOpenConfirm(receiver, local.channelId);
        } else {
            emit ChannelOpenConfirmError(receiver, data);
        }
    }

    /**
     * @notice Initializes a close channel handshake process. It is directly called by the dapp which wants to close
     * the channel
     */
    function channelCloseInit(bytes32 channelId) external {}

    /**
     * @notice Confirms a close channel handshake process. It is called by a relayer on behalf of the dapp whhich
     * initializes the channel closefter after the IBC/VIBC hub chain has processed ChanCloseConfirm event.
     */
    function channelCloseConfirm(address portAddress, bytes32 channelId, Ics23Proof calldata proof) external {}

    /**
     * This func is called by a 'relayer' after the IBC/VIBC hub chain has processed ChanCloseConfirm event.
     * The dApp's onCloseIbcChannel callback is invoked.
     * dApp should throw an error if the channel should not be closed.
     */
    // FIXME this is commented out to make the contract size smaller. We need to optimise for size
    // function onCloseIbcChannel(address portAddress, bytes32 channelId, Ics23Proof calldata proof) external {
    //     // verify VIBC/IBC hub chain has processed ChanCloseConfirm event
    //     _lightClient.verifyMembership(
    //         proof,
    //         bytes('channel/path/to/be/added/here'),
    //         bytes('expected channel bytes constructed from params. Channel.State = {Closed(_Pending?)}')
    //     );
    //
    //     // ensure port owns channel
    //     Channel memory channel = _portChannelMap[portAddress][channelId];
    //     if (channel.counterpartyChannelId == bytes32(0)) {
    //         revert channelNotOwnedByPortAddress();
    //     }
    //
    //     // confirm with dApp by calling its callback
    //     IbcChannelReceiver reciever = IbcChannelReceiver(portAddress);
    //     reciever.onCloseIbcChannel(channelId, channel.counterpartyPortId,
    // channel.counterpartyChannelId);
    //     delete _portChannelMap[portAddress][channelId];
    //     emit CloseIbcChannel(portAddress, channelId);
    // }

    //
    // IBC Packet methods
    //

    /**
     * @notice Sends an IBC packet on a existing channel with the specified packet data and timeout block timestamp.
     * @notice Data should be encoded in a format defined by the channel version, and the module on the other side
     * should know how to parse this.
     * @dev Emits an `IbcPacketEvent` event containing the sender address, channel ID, packet data, and timeout block
     * timestamp.
     * @param channelId The ID of the channel on which to send the packet.
     * @param packet The packet data to send.
     * @param timeoutTimestamp The timestamp in nanoseconds after which the packet times out if it has not been
     * received.
     */
    function sendPacket(bytes32 channelId, bytes calldata packet, uint64 timeoutTimestamp) external {
        // ensure port owns channel
        Channel memory channel = _portChannelMap[msg.sender][channelId];
        if (channel.counterpartyChannelId == bytes32(0)) {
            revert IBCErrors.channelNotOwnedBySender();
        }

        _sendPacket(msg.sender, channelId, packet, timeoutTimestamp);
    }
    /**
     * @notice Handle the acknowledgement of an IBC packet by the counterparty
     * @dev Verifies the given proof and calls the `onAcknowledgementPacket` function on the given `receiver` contract,
     *    ie. the IBC dApp.
     *    Prerequisite: the original packet is committed and not ack'ed or timed out yet.
     *    Note: If the receiving dapp doesn't satisfy the interface, the transaction will be reverted.
     * @param packet The IbcPacket data for the acknowledged packet
     * @param ack The acknowledgement receipt for the packet
     * @param proof The membership proof to verify the packet acknowledgement committed on Polymer chain
     */

    function acknowledgement(IbcPacket calldata packet, bytes calldata ack, Ics23Proof calldata proof) external {
        address receiver = _getAddressFromPort(packet.src.portId);
        // prove ack packet is on Polymer chain
        _lightClient.verifyMembership(proof, Ibc.ackProofKey(packet), abi.encode(Ibc.ackProofValue(ack)));
        // verify packet has been committed and not yet ack'ed or timed out
        bool hasCommitment = _sendPacketCommitment[receiver][packet.src.channelId][packet.sequence];
        if (!hasCommitment) {
            revert IBCErrors.packetCommitmentNotFound();
        }

        // enforce ack'ed packet sequences always increment by 1 for ordered channels
        Channel memory channel = _portChannelMap[receiver][packet.src.channelId];
        (bool success, bytes memory data) = _callIfContract(
            receiver,
            abi.encodeWithSelector(IbcPacketReceiver.onAcknowledgementPacket.selector, packet, Ibc.parseAckData(ack))
        );

        if (success) {
            if (channel.ordering == ChannelOrder.ORDERED) {
                if (packet.sequence != _nextSequenceAck[receiver][packet.src.channelId]) {
                    revert IBCErrors.unexpectedPacketSequence();
                }

                _nextSequenceAck[receiver][packet.src.channelId] = packet.sequence + 1;
            }

            // delete packet commitment to avoid double ack
            delete _sendPacketCommitment[receiver][packet.src.channelId][packet.sequence];
            emit Acknowledgement(receiver, packet.src.channelId, packet.sequence);
        } else {
            emit AcknowledgementError(receiver, data);
        }
    }

    /**
     * @notice Timeout of an IBC packet
     * @dev Verifies the given proof and calls the `onTimeoutPacket` function on the given `receiver` contract, ie. the
     * IBC-dApp.
     * Prerequisite: the original packet is committed and not ack'ed or timed out yet.
     * If the receiving dapp doesn't satisfy the interface, the transaction will be reverted.
     * @param packet The IbcPacket data for the timed-out packet
     * @param proof The non-membership proof data needed to verify the packet timeout
     */
    function timeout(IbcPacket calldata packet, Ics23Proof calldata proof) external {
        // prove absence of packet receipt on Polymer chain
        // TODO: add non membership support
        _lightClient.verifyNonMembership(proof, "packet/receipt/path");

        address receiver = _getAddressFromPort(packet.src.portId);
        // verify packet has been committed and not yet ack'ed or timed out
        bool hasCommitment = _sendPacketCommitment[receiver][packet.src.channelId][packet.sequence];
        if (!hasCommitment) {
            revert IBCErrors.packetCommitmentNotFound();
        }

        (bool success, bytes memory data) =
            _callIfContract(receiver, abi.encodeWithSelector(IbcPacketReceiver.onTimeoutPacket.selector, packet));
        if (success) {
            // delete packet commitment to avoid double timeout
            delete _sendPacketCommitment[receiver][packet.src.channelId][packet.sequence];
            emit Timeout(receiver, packet.src.channelId, packet.sequence);
        } else {
            emit TimeoutError(receiver, data);
        }
    }

    /**
     * @notice Receive an IBC packet and then pass it to the IBC-dApp for processing if verification succeeds.
     * @dev Verifies the given proof and calls the `onRecvPacket` function on the given `receiver` contract
     *  If the address doesn't satisfy the interface, the transaction will be reverted.
     * The receiver must be the intended packet destination, which is the same as packet.dest.portId.
     * @param packet The IbcPacket data for the received packet
     * @param proof The proof data needed to verify the packet receipt
     * @dev Emit an `RecvPacket` event with the details of the received packet;
     * Also emit a WriteAckPacket event, which can be relayed to Polymer chain by relayers
     */
    function recvPacket(IbcPacket calldata packet, Ics23Proof calldata proof) external {
        address receiver = _getAddressFromPort(packet.dest.portId);
        _lightClient.verifyMembership(
            proof, Ibc.packetCommitmentProofKey(packet), abi.encode(Ibc.packetCommitmentProofValue(packet))
        );

        // verify packet has not been received yet
        bool hasReceipt = _recvPacketReceipt[receiver][packet.dest.channelId][packet.sequence];
        if (hasReceipt) {
            revert IBCErrors.packetReceiptAlreadyExists();
        }

        _recvPacketReceipt[receiver][packet.dest.channelId][packet.sequence] = true;

        // enforce recv'ed packet sequences always increment by 1 for ordered channels
        Channel memory channel = _portChannelMap[receiver][packet.dest.channelId];
        if (channel.ordering == ChannelOrder.ORDERED) {
            if (packet.sequence != _nextSequenceRecv[receiver][packet.dest.channelId]) {
                revert IBCErrors.unexpectedPacketSequence();
            }

            _nextSequenceRecv[receiver][packet.dest.channelId] = packet.sequence + 1;
        }

        // Emit recv packet event to prove the relayer did the correct job, and pkt is received.
        emit RecvPacket(receiver, packet.dest.channelId, packet.sequence);

        // If pkt is already timed out, then return early so dApps won't receive it.
        if (_isPacketTimeout(packet)) {
            emit WriteTimeoutPacket(
                receiver, packet.dest.channelId, packet.sequence, packet.timeoutHeight, packet.timeoutTimestamp
            );
            return;
        }

        // Not timeout yet, then do normal handling
        IbcPacket memory pkt = packet;
        AckPacket memory ack;
        (bool success, bytes memory data) =
            _callIfContract(receiver, abi.encodeWithSelector(IbcPacketReceiver.onRecvPacket.selector, pkt));
        if (success) {
            (ack) = abi.decode(data, (AckPacket));
        } else {
            ack = AckPacket(false, data);
        }
        bool hasAckPacketCommitment = _ackPacketCommitment[receiver][packet.dest.channelId][packet.sequence];
        // check is not necessary for sync-acks
        if (hasAckPacketCommitment) {
            revert IBCErrors.ackPacketCommitmentAlreadyExists();
        }

        _ackPacketCommitment[receiver][packet.dest.channelId][packet.sequence] = true;

        emit WriteAckPacket(receiver, packet.dest.channelId, packet.sequence, ack);
    }

    /**
     * Generate a timeout packet for the given packet
     */
    function writeTimeoutPacket(IbcPacket calldata packet, Ics23Proof calldata proof) external {
        _lightClient.verifyMembership(
            proof, Ibc.packetCommitmentProofKey(packet), abi.encode(Ibc.packetCommitmentProofValue(packet))
        );

        address receiver = _getAddressFromPort(packet.src.portId);
        // verify packet does not have a receipt
        bool hasReceipt = _recvPacketReceipt[receiver][packet.dest.channelId][packet.sequence];
        if (hasReceipt) {
            revert IBCErrors.packetReceiptAlreadyExists();
        }

        // verify packet has timed out; zero-value in packet.timeout means no timeout set
        if (!_isPacketTimeout(packet)) {
            revert IBCErrors.packetNotTimedOut();
        }

        emit WriteTimeoutPacket(
            receiver, packet.dest.channelId, packet.sequence, packet.timeoutHeight, packet.timeoutTimestamp
        );
    }

    /**
     * @notice Get the IBC channel with the specified port and channel ID
     * @param portAddress EVM address of the IBC port
     * @param channelId IBC channel ID from the port perspective
     * @return channel A channel struct is always returned. If it doesn't exists, the channel struct is populated with
     * default
     *    values per EVM.
     */
    function getChannel(address portAddress, bytes32 channelId) external view returns (Channel memory channel) {
        channel = _portChannelMap[portAddress][channelId];
    }

    // getOptimisticConsensusState
    function getOptimisticConsensusState(uint256 height)
        external
        view
        returns (uint256 appHash, uint256 fraudProofEndTime, bool ended)
    {
        return _lightClient.getState(height);
    }

    // Prerequisite: must verify sender is authorized to send packet on the channel
    function _sendPacket(address sender, bytes32 channelId, bytes memory packet, uint64 timeoutTimestamp) internal {
        // current packet sequence
        uint64 sequence = _nextSequenceSend[sender][channelId];
        if (sequence == 0) {
            revert IBCErrors.invalidPacketSequence();
        }

        // packet commitment
        _sendPacketCommitment[sender][channelId][sequence] = true;
        // increment nextSendPacketSequence
        _nextSequenceSend[sender][channelId] = sequence + 1;

        emit SendPacket(sender, channelId, packet, sequence, timeoutTimestamp);
    }

    function _connectChannel(
        IbcChannelReceiver portAddress,
        ChannelEnd calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        ChannelEnd calldata counterparty
    ) internal {
        // Register port and channel mapping
        // TODO: check duplicated channel registration?
        // TODO: The call to `Channel` constructor MUST be move to `openIbcChannel` phase
        //       Then `connectIbcChannel` phase can use the `version` as part of `require` condition.
        _portChannelMap[address(portAddress)][local.channelId] = Channel(
            counterparty.version, // TODO: this should be self version instead of counterparty version
            ordering,
            feeEnabled,
            connectionHops,
            counterparty.portId,
            counterparty.channelId
        );

        // initialize channel sequences
        _nextSequenceSend[address(portAddress)][local.channelId] = 1;
        _nextSequenceRecv[address(portAddress)][local.channelId] = 1;
        _nextSequenceAck[address(portAddress)][local.channelId] = 1;
        _channelIdToConnection[local.channelId] = connectionHops[0]; // Set channel to connection mapping for
            // finding
    }

    // Returns the result of the call if no revert, otherwise returns the error if thrown.
    function _callIfContract(address receiver, bytes memory args)
        internal
        returns (bool success, bytes memory message)
    {
        if (!Address.isContract(receiver)) {
            return (false, bytes("call to non-contract"));
        }
        // Only call if we are sure receiver is a contract
        // Note: This tx won't revert if the low-level call fails, see
        // https://docs.soliditylang.org/en/latest/cheatsheet.html#members-of-address
        (success, message) = receiver.call(args);
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    // _isPacketTimeout returns true if the given packet has timed out acoording to host chain's block height and
    // timestamp
    function _isPacketTimeout(IbcPacket calldata packet) internal view returns (bool isTimeOut) {
        return (
            isTimeOut = (packet.timeoutTimestamp != 0 && block.timestamp >= packet.timeoutTimestamp)
            // TODO: check timeoutHeight.revision_number?
            || (packet.timeoutHeight.revision_height != 0 && block.number >= packet.timeoutHeight.revision_height)
        );
    }

    function _getAddressFromPort(string calldata port) internal view returns (address addr) {
        addr = Ibc._hexStrToAddress(port[portPrefixLen:]);
    }
}
