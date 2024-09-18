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

pragma solidity 0.8.15;

import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {Ownable2StepUpgradeable} from "@openzeppelin-upgradeable/contracts/access/Ownable2StepUpgradeable.sol";
import {IbcChannelReceiver, IbcPacketReceiver} from "../interfaces/IbcReceiver.sol";
import {L1Header, OpL2StateProof, Ics23Proof} from "../interfaces/IProofVerifier.sol";
import {ILightClient} from "../interfaces/ILightClient.sol";
import {IDispatcher} from "../interfaces/IDispatcher.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin-upgradeable/contracts/security/ReentrancyGuardUpgradeable.sol";
import {Channel, ChannelEnd, ChannelOrder, IbcPacket, ChannelState, AckPacket, AckStatus, Ibc} from "../libs/Ibc.sol";
import {IBCErrors} from "../libs/IbcErrors.sol";
import {IbcUtils} from "../libs/IbcUtils.sol";
import {IFeeVault} from "../interfaces/IFeeVault.sol";

/**
 * @title Dispatcher
 * @author Polymer Labs
 * @notice This contract facilitates the 4-step channel opening process and IBC packet sending/receiving.
 * @notice Contract callers call this contract to send IBC-like messages, which can be relayed to a rollup module on the
 * Polymerase chain
 * @notice In addition to directly calling this contract, this contract can also be called by arbitrary middleware
 * contracts, such as the UniversalChannelHandler
 * @notice This contract is upgradeable and uses the UUPS pattern
 *
 */
contract Dispatcher is Ownable2StepUpgradeable, UUPSUpgradeable, ReentrancyGuardUpgradeable, IDispatcher {
    // Gap to allow for additional contract inheritance, similar to OpenZeppelin's Initializable contract
    uint256[48] private __gap;

    string public portPrefix;
    uint32 public portPrefixLen;

    mapping(address => mapping(bytes32 => Channel)) private _portChannelMap; // Mapping used to check if a channel is
        // owned by a port
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

    ILightClient _UNUSED; // From previous dispatcher version
    mapping(bytes32 => string) private _channelIdToConnection;
    mapping(string => ILightClient) private _connectionToLightClient;
    IFeeVault public feeVault;

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initializes the Dispatcher contract with the provided port prefix.
     * @param initPortPrefix The initial port prefix to be set for the contract.
     * @dev This method should be called only once during contract deployment.
     * @dev For contract upgarades, which need to reinitialize the contract, use the reinitializer modifier.
     */
    function initialize(string memory initPortPrefix, IFeeVault _feeVault)
        external
        virtual
        reinitializer(2)
        nonReentrant
    {
        if (bytes(initPortPrefix).length == 0) {
            revert IBCErrors.invalidPortPrefix();
        }
        if (address(_feeVault) == address(0)) {
            revert IBCErrors.invalidAddress();
        }
        __Ownable2Step_init();
        __ReentrancyGuard_init();
        portPrefix = initPortPrefix;
        portPrefixLen = uint32(bytes(initPortPrefix).length);
        feeVault = _feeVault;
    }

    /**
     * @notice Sets the port prefix for the Dispatcher contract.
     * @param _portPrefix The new port prefix to be set.
     * @dev It can only be called by the contract owner.
     */
    function setPortPrefix(string calldata _portPrefix) external onlyOwner {
        if (bytes(_portPrefix).length == 0) {
            revert IBCErrors.invalidPortPrefix();
        }
        portPrefix = _portPrefix;
        portPrefixLen = uint32(bytes(_portPrefix).length);
    }

    // updateClientWithOptimisticConsensusState updates the client
    // with the optimistic consensus state. The optimistic consensus
    // is accepted and will be open for verify in the fraud proof
    // window.
    /**
     * @notice Updates the client with optimistic consensus state. The optimistic consensus is accepted and will be open
     * for verification in the fraud proof window
     * @dev Calls lightClient.addOpConsensusState; See Lightclient natspec for params information
     * @dev This function updates the client with optimistic consensus state.
     *      It should be called after verifying the optimistic consensus state on the main chain.
     */
    function updateClientWithOptimisticConsensusState(
        L1Header calldata l1header,
        OpL2StateProof calldata proof,
        uint256 height,
        uint256 appHash,
        string calldata connection
    ) external returns (uint256 fraudProofEndTime, bool ended) {
        return _getLightClientFromConnection(connection).addOpConsensusState(l1header, proof, height, appHash);
    }

    /**
     * @notice Adds a new mapping between a light client and a connection string.
     * @param connection The string representing the new connection, as per ICS03 (should always be first connection in
     * connectionHop array)
     * @param lightClient The ILightClient contract address used by the connection
     * @dev This method can only be called by the contract owner.
     */
    function setClientForConnection(string calldata connection, ILightClient lightClient) external onlyOwner {
        _setClientForConnection(connection, lightClient);
    }

    /**
     * @notice Remove's the given connection's light client.
     * @notice The conneciton will be unuseable after this operation until the client is set again.
     * @dev Only callable by the contract owner.
     * @param connection The connection string to remove the light client from.
     */
    function removeConnection(string calldata connection) external onlyOwner {
        delete _connectionToLightClient[connection];
    }

    /**
     * This function is called by a dApp to initiate a channel handshake.
     */
    /**
     * @notice Initializes the channel opening process with the specified parameters. This is the first step in the  channel
     * handshake, initiated directly by the dapp which wishes to establish a channel with the receiver.
     * @param ordering The ordering of the channel (ORDERED or UNORDERED). Note: ORDERED channels are not currently
     * supported
     * @param feeEnabled A boolean indicating whether fees are enabled for the channel. Note: This value isn't currently
     * used, and is not being verified in light client proofs.
     * @param connectionHops The list of connection hops associated with the channel, with the first channel in this
     * array always starting from the chain this contract is deployed on
     * @param counterpartyPortId The port ID of the counterparty.
     */
    function channelOpenInit(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) external nonReentrant {
        // We need to validate connectionHops & counterpartyPortId since they aren't validated in an internal
        // function like all other instances of these arguments. ConnectionHops can only be of length 2 for now since we
        // don't yet support multihop, but this will be updated in the future
        if (connectionHops.length != 2 || bytes(counterpartyPortId).length == 0) {
            revert IBCErrors.invalidCounterParty();
        }
        if (bytes(version).length == 0) {
            revert IBCErrors.invalidVersion();
        }

        emit ChannelOpenInit(msg.sender, version, ordering, feeEnabled, connectionHops, counterpartyPortId);
    }

    /**
     * @notice Initializes step 2 of the channel handshake process. This is called by a relayer on behalf of the
     * receiving dapp.
     * To minimize trust assumptions in proving that the counterparty has indeed initiated step 1, there must be a valid
     * proof of the counterparty being in the TRY_PENDING state from the Polymer Client
     * @notice   The receiving dApp should implement IbcChannelHandler's onChanOpenTry callback. If the callback
     * succeeds, the dApp should return the selected version and the emitted even will be relayed to the  IBC/VIBC hub
     * chain.
     * @param local The counterparty information for the receiver.
     * @param ordering The ordering of the channel (ORDERED or UNORDERED). Note: ORDERED channels are not currently
     * supported
     * @param feeEnabled Whether fees are enabled for the channel. Note: This value isn't currently used, and is not
     * being verified in light client proofs
     * @param connectionHops The list of connection hops associated with the channel; with the first channel in this
     * array always starting from the chain this contract is deployed on
     * @param counterparty The counterparty information of the sender
     * @param proof The proof that the counterparty is in the TRY_PENDING state (i.e. that it is indeed intending to
     * initialize a channel with the given receiver)
     */
    function channelOpenTry(
        ChannelEnd calldata local,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        ChannelEnd calldata counterparty,
        Ics23Proof calldata proof
    ) external nonReentrant {
        // Note: ConnectionHops can only be of length 2 for now since we don't yet support multihop, but this will be
        // updated in the future
        if (connectionHops.length != 2) {
            revert IBCErrors.invalidConnectionHops();
        }
        _checkInvalidChannelEnds(local.portId, local.channelId, counterparty.portId, counterparty.channelId);
        _getLightClientFromConnection(connectionHops[0]).verifyMembership(
            proof,
            Ibc.channelProofKey(local.portId, local.channelId),
            Ibc.channelProofValue(
                ChannelState.TRY_PENDING,
                ordering,
                local.version,
                connectionHops,
                counterparty.portId,
                counterparty.channelId
            )
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
     * @notice Initializes step 3 of the channel handshake process. This method is called by a relayer on behalf of the
     * sending, dapp (i.e. the dapp which initiated the channelOpenInit call). This step happens after the IBC/VIBC hub
     * chain has processed the ChannelOpenTry event.
     * To minimize trust assumptions in proving that the counterparty had indeed responded successfully in step 2 of the
     * handshake, there must be a valid proof of the counterparty being in the ACK_PENDING state from the Polymer Client
     * @notice Completes the channel opening acknowledge process with the specified parameters.
     * @notice The dApp should implement the onChannelOpenAck method to handle the third channel handshake method:
     * ChanOpenAck process.
     * @param local The counterparty information for the local channel.
     * @param connectionHops The list of connection hops associated with the channel, with the first channel in this
     * array always starting from the chain this contract is deployed on.
     * @param ordering The ordering of the channel (ORDERED or UNORDERED). Note: ORDERED channels are not currently
     * supported
     * @param feeEnabled A boolean indicating whether fees are enabled for the channel. Note: This value isn't currently
     * used and is not being verified in light client proofs.
     * @param counterparty The counterparty information for the channel.
     * @param proof The proof that the counterparty is in the ACK_PENDING state (i.e. that it responded with a
     * successful channelOpenTry )
     */
    function channelOpenAck(
        ChannelEnd calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        ChannelEnd calldata counterparty,
        Ics23Proof calldata proof
    ) external nonReentrant {
        if (connectionHops.length < 2) {
            revert IBCErrors.invalidConnectionHops();
        }

        _checkInvalidChannelEnds(local.portId, local.channelId, counterparty.portId, counterparty.channelId);
        _getLightClientFromConnection(connectionHops[0]).verifyMembership(
            proof,
            Ibc.channelProofKey(local.portId, local.channelId),
            Ibc.channelProofValue(
                ChannelState.ACK_PENDING,
                ordering,
                local.version,
                connectionHops,
                counterparty.portId,
                counterparty.channelId
            )
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
     * @notice Initializes step 4 of the channel handshake process. This method is called by a relayer on behalf of the
     * receiving dapp. This step happens after the IBC/VIBC hub chain has processed the ChannelOpenAck event.
     * To minimize trust assumptions in proving that the counterparty had indeed responded successfully in step 2 of the
     * handshake, there must be a valid proof of the counterparty being in the CONFIRM_PENDING state from the Polymer
     * Client
     *
     * @notice Initiates the channel opening confirm process with the specified parameters.
     * The receiver should implement the onChannelConnect method to handle the last channel handshake method:
     * ChannelOpenConfirm
     * @param local The counterparty information for the local channel.
     * @param ordering The ordering of the channel (ORDERED or UNORDERED). Note: ORDERED channels are not currently
     * supported
     * @param connectionHops The list of connection hops associated with the channel, with the first channel in this
     * array always starting from the chain this contract is deployed on.
     * @param feeEnabled A boolean indicating whether fees are enabled for the channel. Note: This value isn't currently
     * used and is not being verified in light client proofs.
     * @param counterparty The counterparty information for the channel.
     * @param proof The proof of channel opening confirm.
     * @dev This function initiates the channel opening confirm process by calling the onChanOpenConfirm function of the
     * specified
     *      receiver contract.
     *      It can only be called by authorized parties.
     */
    function channelOpenConfirm(
        ChannelEnd calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        ChannelEnd calldata counterparty,
        Ics23Proof calldata proof
    ) external nonReentrant {
        if (connectionHops.length < 2) {
            revert IBCErrors.invalidConnectionHops();
        }
        _checkInvalidChannelEnds(local.portId, local.channelId, counterparty.portId, counterparty.channelId);
        _getLightClientFromConnection(connectionHops[0]).verifyMembership(
            proof,
            Ibc.channelProofKey(local.portId, local.channelId),
            Ibc.channelProofValue(
                ChannelState.CONFIRM_PENDING,
                ordering,
                local.version,
                connectionHops,
                counterparty.portId,
                counterparty.channelId
            )
        );

        address receiver = _getAddressFromPort(local.portId);
        (bool success, bytes memory data) = _callIfContract(
            receiver, abi.encodeWithSelector(IbcChannelReceiver.onChanOpenConfirm.selector, local.channelId)
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
     * @dev Emits a `CloseIbcChannel` event with the given `channelId` and the address of the message sender
     * @notice Close the specified IBC channel by channel ID
     * @notice Must be called by the channel owner, ie. _portChannelMap[msg.sender][channelId] must exist
     */
    function channelCloseInit(bytes32 channelId) external nonReentrant {
        _checkChannelOwner(channelId);

        // Note: We delete the portChannelMap here even on Dapp revert to avoid having a case where a dapp deployed with
        // a faulty callback cannot close a channel (as is done on channelCloseConfirm)
        delete _portChannelMap[msg.sender][channelId];

        // Note: We also delete the channelId to connection mapping to brick any proof-dependent calls for this channel.
        // This is to keep this dispatcher's view of the channel consistent across local and proof-dependent logic. This
        // is useful in ensuring this channel end behaves consistently even if there are no events relayed to peptide
        delete _channelIdToConnection[channelId];
        emit ChannelCloseInit(msg.sender, channelId);
    }

    /**
     * @notice Confirms a close channel handshake process. It is called by a relayer on behalf of the dapp whhich
     * initializes the channel closefter after the IBC/VIBC hub chain has processed ChanCloseConfirm event.
     * @dev Emits a `CloseIbcChannel` event with the given `channelId` and the address of the message sender
     * The dApp's onChanCloseConfirm callback is invoked. dApp should throw an error if the channel should not be
     * closed.
     * @notice Close the specified IBC channel by channel ID
     * @notice Must be called by the channel owner, ie. _portChannelMap[msg.sender][channelId] must exist
     * @param portAddress The address of the receiver port of the channel to close
     * @param channelId The ID of the channel to close
     */
    function channelCloseConfirm(address portAddress, bytes32 channelId, Ics23Proof calldata proof)
        external
        nonReentrant
    {
        if (portAddress == address(0)) {
            revert IBCErrors.invalidAddress();
        }
        // ensure port owns channel - note: we don't use _checkChannelOwner here since we need the channel aftewards and
        // we save an SLOAD here.
        Channel memory channel = _portChannelMap[portAddress][channelId];
        if (channel.counterpartyChannelId == bytes32(0)) {
            // _portChannelMap is used to check ownership of a channel by a port
            revert IBCErrors.channelNotOwnedByPortAddress();
        }

        // verify VIBC/IBC hub chain has processed ChanCloseConfirm event
        _getLightClientFromChannelId(channelId).verifyMembership(
            proof,
            Ibc.channelProofKeyMemory(channel.portId, channelId),
            Ibc.channelProofValueMemory(
                ChannelState.CLOSE_CONFIRM_PENDING,
                channel.ordering,
                channel.version,
                channel.connectionHops,
                channel.counterpartyPortId,
                channel.counterpartyChannelId
            )
        );

        (bool success, bytes memory data) = _callIfContract(
            portAddress,
            abi.encodeWithSelector(
                IbcChannelReceiver.onChanCloseConfirm.selector,
                channelId,
                channel.counterpartyPortId,
                channel.counterpartyChannelId
            )
        );

        // Note: We delete the portChannelMap here even on Dapp revert to avoid having a case where a dapp deployed with
        // a faulty callback cannot close a channel (as is done on channelCloseInit)
        delete _portChannelMap[portAddress][channelId];
        // Note: We also delete the channelId to connection mapping to brick any proof-dependent calls for this channel.
        // This is to keep this dispatcher's view of the channel consistent across local and proof-dependent logic. This
        // is useful in ensuring this channel end behaves consistently even if there are no events relayed to peptide
        delete _channelIdToConnection[channelId];
        if (success) {
            emit ChannelCloseConfirm(portAddress, channelId);
        } else {
            emit ChannelCloseConfirmError(address(portAddress), data);
        }
    }

    /**
     * @notice Sends an IBC packet on a existing channel with the specified packet data and timeout block timestamp.
     * @dev Emits an `IbcPacketEvent` event containing the sender address, channel ID, packet data, and timeout block
     * timestamp (formatted as seconds after the unix epoch).
     * @param channelId The ID of the channel on which to send the packet.
     * @param packet The packet data to send.
     * @param timeoutTimestamp The timestamp, in seconds after the unix epoch, after which the packet times out if it
     * has not been received.
     * @return sequence The sequence number of the packet, starting from 1 and incremented with each sendPacket. This
     * sequence number is used to link with depositing packet fees in the fee vault
     */
    function sendPacket(bytes32 channelId, bytes calldata packet, uint64 timeoutTimestamp)
        external
        nonReentrant
        returns (uint64 sequence)
    {
        _checkChannelOwner(channelId);

        if (timeoutTimestamp <= block.timestamp) {
            revert IBCErrors.invalidPacket();
        }

        // current packet sequence
        sequence = _nextSequenceSend[msg.sender][channelId];
        if (sequence == 0) {
            revert IBCErrors.invalidPacketSequence();
        }

        // packet commitment
        _sendPacketCommitment[msg.sender][channelId][sequence] = true;
        // increment nextSendPacketSequence
        _nextSequenceSend[msg.sender][channelId] = sequence + 1;

        emit SendPacket(msg.sender, channelId, packet, sequence, timeoutTimestamp);
    }

    /**
     * @notice Handle the acknowledgement of an IBC packet by the counterparty, called by a relayer.
     * @dev To minimize trust assumptions, an inclusion proof must be given that proves the packet commitment is
     * in the IBC/VIBC hub chain.
     * @dev This method also calls the `onAcknowledgementPacket` function on
     * the given `receiver` contract,
     * @dev Prerequisite: The original packet must be committed and not ack'ed or timed out yet.
     *    Note: If the receiving dapp doesn't satisfy the interface, the transaction will be reverted.
     * @param packet The IbcPacket data for the acknowledged packet
     * @param ack The acknowledgement receipt for the packet
     * @param proof The membership proof to verify the packet acknowledgement committed on Polymer chain
     */
    function acknowledgement(IbcPacket calldata packet, bytes calldata ack, Ics23Proof calldata proof)
        external
        nonReentrant
    {
        address receiver = _getAddressFromPort(packet.src.portId);
        // prove ack packet is on Polymer chain
        _getLightClientFromChannelId(packet.src.channelId).verifyMembership(
            proof, Ibc.ackProofKey(packet), abi.encode(Ibc.ackProofValue(ack))
        );
        // verify packet has been committed and not yet ack'ed or timed out
        bool hasCommitment = _sendPacketCommitment[receiver][packet.src.channelId][packet.sequence];
        if (!hasCommitment) {
            revert IBCErrors.packetCommitmentNotFound();
        }

        // enforce ack'ed packet sequences always increment by 1 for ordered channels
        (bool success, bytes memory data) = _callIfContract(
            receiver,
            abi.encodeWithSelector(IbcPacketReceiver.onAcknowledgementPacket.selector, packet, Ibc.parseAckData(ack))
        );

        if (success) {
            if (_portChannelMap[address(receiver)][packet.src.channelId].ordering == ChannelOrder.ORDERED) {
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
     * @notice Timeout of an IBC packet. This method is not yet implemented but meant as a placeholder for any
     * integrating dapps.
     * @dev Verifies the given proof and calls the `onTimeoutPacket` function on the given `receiver` contract, ie. the
     * IBC-dApp.
     * Prerequisite: the original packet is committed and not ack'ed or timed out yet.
     * If the receiving dapp doesn't satisfy the interface, the transaction will be reverted.
     * @param packet The IbcPacket data for the timed-out packet
     * @param proof The non-membership proof data needed to verify the packet timeout
     */
    function timeout(IbcPacket calldata packet, Ics23Proof calldata proof) external nonReentrant {
        // Note: This method isn't needed for full vibc functionality, since packets can still be timed out in
        // recvPacket and writeTimeoutPacket, so this method will only eventually be implemented to comply with IbcSpec,
        // and is meant to only be a placeholder for integrating dapps.
    }

    /**
     * @notice Receive an IBC packet and then pass it to the IBC-dApp for processing if verification succeeds.
     * @dev To minimize trust assumptions, verifies the given proof
     * @dev calls the `onRecvPacket` function on the given receiver contract
     * If the address doesn't satisfy the interface, the nextSequenceRecv will not be updated
     * The receiver must be the intended packet destination, which is the same as packet.dest.portId.
     * @param packet The IbcPacket data for the received packet
     * @param proof The proof data needed to verify the packet receipt
     * @dev Emit an `RecvPacket` event with the details of the received packet;
     * Also emit a WriteAckPacket event, which can be relayed to Polymer chain by relayers
     */
    function recvPacket(IbcPacket calldata packet, Ics23Proof calldata proof) external nonReentrant {
        _getLightClientFromChannelId(packet.dest.channelId).verifyMembership(
            proof, Ibc.packetCommitmentProofKey(packet), abi.encode(Ibc.packetCommitmentProofValue(packet))
        );
        address receiver = _getAddressFromPort(packet.dest.portId);

        // verify packet has not been received yet
        bool hasReceipt = _recvPacketReceipt[receiver][packet.dest.channelId][packet.sequence];
        if (hasReceipt) {
            revert IBCErrors.packetReceiptAlreadyExists();
        }

        _recvPacketReceipt[receiver][packet.dest.channelId][packet.sequence] = true;

        // enforce recv'ed packet sequences always increment by 1 for ordered channels
        if (_portChannelMap[address(receiver)][packet.dest.channelId].ordering == ChannelOrder.ORDERED) {
            if (packet.sequence != _nextSequenceRecv[receiver][packet.dest.channelId]) {
                revert IBCErrors.unexpectedPacketSequence();
            }

            _nextSequenceRecv[receiver][packet.dest.channelId] = packet.sequence + 1;
        }

        // Emit recv packet event to prove the relayer did the correct job, and pkt is received.
        emit RecvPacket(receiver, packet.dest.channelId, packet.sequence);

        // If pkt is already timed out, then return early so dApps won't receive it.
        if (_isPacketTimeout(packet.timeoutTimestamp, packet.timeoutHeight.revision_height)) {
            emit WriteTimeoutPacket(
                receiver, packet.dest.channelId, packet.sequence, packet.timeoutHeight, packet.timeoutTimestamp
            );
            return;
        }

        AckPacket memory ack;
        // Not timeout yet, then do normal handling
        (bool success, bytes memory data) =
            _callIfContract(receiver, abi.encodeWithSelector(IbcPacketReceiver.onRecvPacket.selector, packet));
        if (success) {
            (ack) = abi.decode(data, (AckPacket));
            if (ack.status == AckStatus.SKIP) {
                return;
            }
        } else {
            ack = AckPacket(AckStatus.FAILURE, data);
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
     * @notice Writes a timeout packet to the specified receiver.
     * @param packet The timeout packet data.
     * @dev Requirements:
     * - The `receiver` address must not be the zero address.
     * - The `receiver` must be the original packet sender.
     * - The packet must not have a receipt.
     * - The packet must have timed out.
     */
    function writeTimeoutPacket(IbcPacket calldata packet, Ics23Proof calldata proof) external nonReentrant {
        _getLightClientFromChannelId(packet.dest.channelId).verifyMembership(
            proof, Ibc.packetCommitmentProofKey(packet), abi.encode(Ibc.packetCommitmentProofValue(packet))
        );

        address receiver = _getAddressFromPort(packet.dest.portId);
        // verify packet does not have a receipt
        bool hasReceipt = _recvPacketReceipt[receiver][packet.dest.channelId][packet.sequence];
        if (hasReceipt) {
            revert IBCErrors.packetReceiptAlreadyExists();
        }

        // verify packet has timed out; zero-value in packet.timeout means no timeout set
        if (!_isPacketTimeout(packet.timeoutTimestamp, packet.timeoutHeight.revision_height)) {
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
     * default values per EVM.
     */
    function getChannel(address portAddress, bytes32 channelId) external view returns (Channel memory channel) {
        // _portChannelMap is used to check ownership of a channel by a port
        channel = _portChannelMap[portAddress][channelId];
    }

    /**
     * @notice Retrieves the optimistic consensus state for the specified height and client ID.
     * @param height The height of the consensus state.
     * @param connection The connection the client corresponds to; used to find the light client.
     * @return appHash The application hash of the consensus state.
     * @return fraudProofEndTime The end time of the fraud proof.
     * @return ended A boolean indicating whether the consensus state has ended.
     */
    function getOptimisticConsensusState(uint256 height, string calldata connection)
        external
        view
        returns (uint256 appHash, uint256 fraudProofEndTime, bool ended)
    {
        return _getLightClientFromConnection(connection).getState(height);
    }

    function _setClientForConnection(string calldata connection, ILightClient lightClient) internal {
        if (bytes(connection).length == 0) {
            revert IBCErrors.invalidConnection("");
        }
        if (address(lightClient) == address(0)) {
            revert IBCErrors.invalidAddress();
        }

        _connectionToLightClient[connection] = lightClient;
    }

    /**
     * @notice Connects a channel with the provided parameters.
     * @param portAddress The address of the IbcChannelReceiver contract.
     * @param local The details of the local counterparty.
     * @param connectionHops The connection hops associated with the channel, with the first channel in this
     * array always starting from the chain this contract is deployed on.
     * @param ordering The ordering of the channel. Note: ORDERED channels are not currently supported
     * @param feeEnabled A boolean indicating whether fees are enabled for the channel. Note: This value isn't currently
     * used and is not being verified in light client proofs.
     * @param counterparty The details of the counterparty.
     */
    function _connectChannel(
        IbcChannelReceiver portAddress,
        ChannelEnd calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        ChannelEnd calldata counterparty
    ) internal {
        // We don't need to check that a channel isn't already present between a portAddress and a chanenlId since
        // polymer chain verification should prevent double registration of the same channel (with the execption of the
        // feeEnabled state, though this isn't currently used anywhere so change a channel's feeEnabled state shouldn't
        // have any outcome)
        _portChannelMap[address(portAddress)][local.channelId] = Channel(
            local.version,
            ordering,
            feeEnabled,
            connectionHops,
            counterparty.portId,
            counterparty.channelId,
            local.portId
        );

        // initialize channel sequences
        _nextSequenceSend[address(portAddress)][local.channelId] = 1;
        _nextSequenceRecv[address(portAddress)][local.channelId] = 1;
        _nextSequenceAck[address(portAddress)][local.channelId] = 1;
        _channelIdToConnection[local.channelId] = connectionHops[0]; // Set channel to connection mapping for finding
            // light client
    }

    /**
     * @notice Calls the specified contract with the provided arguments and returns the success status and
     *         message.
     * @notice This functions similar to a try/catch statement - if the low-level call reverts, this transaction will
     * not revert, but instead the error message will be returned
     * @param receiver The address of the contract to call.
     * @param args The abi-encoded low-level call
     * @return success A boolean indicating whether the call was successful and made to a contract address.
     * @return message The return message from the contract call; or error for a failed call
     *  @dev Requirements for the success return value to be true:
     * - The `receiver` address must correspond to a contract.
     * - The receiver contract must have the called implemented
     * - The receiver contract must not revert in the low-level call
     */
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
        uint256 gasBeforeCall = gasleft();
        (success, message) = receiver.call(args);

        // x is the amount of gas left right before the low level call. Solidity will allocate x*63/64 to the low level
        // call
        // and x/64 for the remaining execution after the low-level call. If this low-level call runs out of gas, then
        // gasLeft will be equal to x/64 at the start of the remaining execution. so we should check gasBefore < 1/64
        if (!success && gasleft() <= gasBeforeCall / 64) {
            // Only check for out of gas if the call failed; if it was a successful call then it was gauranteed to not
            // run out of gas
            revert IBCErrors.notEnoughGas();
        }
    }

    /**
     * @inheritdoc UUPSUpgradeable
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    function _checkChannelOwner(bytes32 channelId) internal view {
        if (_portChannelMap[msg.sender][channelId].counterpartyChannelId == bytes32(0)) {
            // _portChannelMap is used to check ownership of a channel by a port
            revert IBCErrors.channelNotOwnedBySender();
        }
    }

    /**
     * @notice Checks if the given packet has timed out according to the host chain's block height and timestamp.
     * @param timeoutTimestamp The timeout timestamp of the packet.
     * @param revisionHeight The block height of the packet's timeout revision.
     * @return isTimeOut Returns true if the given packet has timed out acoording to host chain's block height and
     * timestamp, in seconds after the unix epoch.
     */
    function _isPacketTimeout(uint64 timeoutTimestamp, uint64 revisionHeight) internal view returns (bool isTimeOut) {
        return (
            isTimeOut = (timeoutTimestamp != 0 && block.timestamp >= timeoutTimestamp)
            // TODO: check timeoutHeight.revision_number?
            || (revisionHeight != 0 && block.number >= revisionHeight)
        );
    }

    function _getAddressFromPort(string calldata port) internal view returns (address addr) {
        addr = IbcUtils.hexStrToAddress(port[portPrefixLen:]);
    }

    /**
     * @notice Retrieves the light client associated with the specified channel ID.
     * @param channelId The ID of the channel.
     * @return lightClient The light client associated with the channel.
     */
    function _getLightClientFromChannelId(bytes32 channelId) internal view returns (ILightClient lightClient) {
        string memory connection = _channelIdToConnection[channelId];
        if (bytes(connection).length == 0) {
            revert IBCErrors.channelIdNotFound(channelId);
        }
        return _getLightClientFromConnection(connection);
    }

    /**
     * @notice Retrieves the light client associated with the specified connection.
     * @param connection The connection string(i.e. the first connection in the connectionHops array)
     * @return lightClient The light client associated with the connection.
     */
    function _getLightClientFromConnection(string memory connection) internal view returns (ILightClient lightClient) {
        if (bytes(connection).length == 0) {
            revert IBCErrors.invalidConnection("");
        }
        lightClient = _connectionToLightClient[connection];
        if (address(lightClient) == address(0)) {
            revert IBCErrors.lightClientNotFound(connection);
        }
    }

    function _checkInvalidChannelEnds(
        string calldata localPortId,
        bytes32 localChannelId,
        string calldata counterPartyPortId,
        bytes32 counterPartyChannelId
    ) internal pure {
        if (
            bytes(localPortId).length == 0 || bytes(counterPartyPortId).length == 0 || localChannelId == bytes32(0)
                || counterPartyChannelId == bytes32(0)
        ) {
            revert IBCErrors.invalidCounterParty();
        }
    }
}
