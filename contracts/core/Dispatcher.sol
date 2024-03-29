//SPDX-License-Identifier: UNLICENSED

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
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
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

/**
 * @title Dispatcher
 * @author Polymer Labs
 * @notice
 *     Contract callers call this contract to send IBC-like msg,
 *     which can be relayed to a rollup module on the Polymerase chain
 */
contract Dispatcher is OwnableUpgradeable, UUPSUpgradeable, IDispatcher {
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
    mapping(bytes32 => string) private _channelIdToConnection;
    mapping(string => LightClient) private _connectionToLightClient;
    uint256 private _numClients;

    //
    // methods
    //
    constructor() {
        _disableInitializers();
    }

    function initialize(string memory initPortPrefix) public virtual initializer {
        __Ownable_init();
        portPrefix = initPortPrefix;
        portPrefixLen = uint32(bytes(initPortPrefix).length);
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
        uint256 appHash,
        string calldata connection
    ) external returns (uint256 fraudProofEndTime, bool ended) {
        return _getLightClientFromConnection(connection).addOpConsensusState(l1header, proof, height, appHash);
    }

    function setNewConnection(string calldata connection, LightClient lightClient) external onlyOwner {
        _setNewConnection(connection, lightClient);
    }

    function removeConnection(string calldata connection) external onlyOwner {
        delete _connectionToLightClient[connection];
    }

    /**
     * This function is called by a 'relayer' on behalf of a dApp. The dApp should implement IbcChannelHandler's
     * onChanOpenInit. If the callback succeeds, the dApp should return the selected version and the emitted event
     * will be relayed to the  IBC/VIBC hub chain.
     */
    function channelOpenInit(
        IbcChannelReceiver receiver,
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) external {
        if (bytes(counterpartyPortId).length == 0) {
            revert IBCErrors.invalidCounterPartyPortId();
        }

        (bool success, bytes memory data) = _callIfContract(
            address(receiver), abi.encodeWithSelector(IbcChannelReceiver.onChanOpenInit.selector, version)
        );

        if (success) {
            emit ChannelOpenInit(
                address(receiver), abi.decode(data, (string)), ordering, feeEnabled, connectionHops, counterpartyPortId
            );
        } else {
            emit ChannelOpenInitError(address(receiver), data);
        }
    }

    /**
     * This function is called by a 'relayer' on behalf of a dApp. The dApp should implement IbcChannelHandler's
     * onChanOpenTry. If the callback succeeds, the dApp should return the selected version and the emitted event
     * will be relayed to the  IBC/VIBC hub chain.
     */
    function channelOpenTry(
        IbcChannelReceiver receiver,
        CounterParty calldata local,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        CounterParty calldata counterparty,
        Ics23Proof calldata proof
    ) external {
        if (bytes(counterparty.portId).length == 0) {
            revert IBCErrors.invalidCounterPartyPortId();
        }

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

        (bool success, bytes memory data) = _callIfContract(
            address(receiver), abi.encodeWithSelector(IbcChannelReceiver.onChanOpenTry.selector, counterparty.version)
        );

        if (success) {
            emit ChannelOpenTry(
                address(receiver),
                abi.decode(data, (string)),
                ordering,
                feeEnabled,
                connectionHops,
                counterparty.portId,
                counterparty.channelId
            );
        } else {
            emit ChannelOpenTryError(address(receiver), data);
        }
    }

    /**
     * This func is called by a 'relayer' after the IBC/VIBC hub chain has processed the ChannelOpenTry event.
     * The dApp should implement the onChannelConnect method to handle the third channel handshake method: ChanOpenAck
     */
    function channelOpenAck(
        IbcChannelReceiver receiver,
        CounterParty calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        CounterParty calldata counterparty,
        Ics23Proof calldata proof
    ) external {
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

        (bool success, bytes memory data) = _callIfContract(
            address(receiver),
            abi.encodeWithSelector(IbcChannelReceiver.onChanOpenAck.selector, local.channelId, counterparty.version)
        );

        if (success) {
            _connectChannel(receiver, local, connectionHops, ordering, feeEnabled, counterparty);
            emit ChannelOpenAck(address(receiver), local.channelId);
        } else {
            emit ChannelOpenAckError(address(receiver), data);
        }
    }

    /**
     * This func is called by a 'relayer' after the IBC/VIBC hub chain has processed the ChannelOpenTry event.
     * The dApp should implement the onChannelConnect method to handle the last channel handshake method:
     * ChannelOpenConfirm
     */
    function channelOpenConfirm(
        IbcChannelReceiver receiver,
        CounterParty calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        CounterParty calldata counterparty,
        Ics23Proof calldata proof
    ) external {
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

        (bool success, bytes memory data) = _callIfContract(
            address(receiver),
            abi.encodeWithSelector(IbcChannelReceiver.onChanOpenConfirm.selector, local.channelId, counterparty.version)
        );

        if (success) {
            _connectChannel(receiver, local, connectionHops, ordering, feeEnabled, counterparty);
            emit ChannelOpenConfirm(address(receiver), local.channelId);
        } else {
            emit ChannelOpenConfirmError(address(receiver), data);
        }
    }

    /**
     * @dev Emits a `CloseIbcChannel` event with the given `channelId` and the address of the message sender
     * @notice Close the specified IBC channel by channel ID
     * Must be called by the channel owner, ie. _portChannelMap[msg.sender][channelId] must exist
     */
    function channelCloseInit(bytes32 channelId) external {
        Channel memory channel = _portChannelMap[msg.sender][channelId];
        if (channel.counterpartyChannelId == bytes32(0)) {
            revert IBCErrors.channelNotOwnedBySender();
        }
        (bool success, bytes memory data) = _callIfContract(
            msg.sender,
            abi.encodeWithSelector(
                IbcChannelReceiver.onChanCloseInit.selector,
                channelId,
                channel.counterpartyPortId,
                channel.counterpartyChannelId
            )
        );

        delete _portChannelMap[msg.sender][channelId];
        if (success) {
            emit ChannelCloseInit(msg.sender, channelId);
        } else {
            emit ChannelCloseInitError(address(msg.sender), data);
        }
    }

    /**
     * This func is called by a 'relayer' after the IBC/VIBC hub chain has processed ChanCloseConfirm event.
     * The dApp's onChanCloseConfirm callback is invoked.
     * dApp should throw an error if the channel should not be closed.
     */
    function channelCloseConfirm(address portAddress, bytes32 channelId, Ics23Proof calldata proof) external {
        // ensure port owns channel
        Channel memory channel = _portChannelMap[portAddress][channelId];
        if (channel.counterpartyChannelId == bytes32(0)) {
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

        delete _portChannelMap[portAddress][channelId];
        if (success) {
            emit ChannelCloseConfirm(portAddress, channelId);
        } else {
            emit ChannelCloseConfirmError(address(portAddress), data);
        }
    }

    //
    // IBC Packet methods
    //

    /**
     * @notice Sends an IBC packet on a existing channel with the specified packet data and timeout block timestamp.
     * @notice Data should be encoded in a format defined by the channel version, and the module on the other side
     * should know how to parse this.
     * @dev Emits an `IbcPacketEvent` event containing the sender address, channel ID, packet data, and timeout block
     * timestamp (formatted as seconds after the unix epoch).
     * @param channelId The ID of the channel on which to send the packet.
     * @param packet The packet data to send.
     * @param timeoutTimestamp The timestamp, in seconds after the unix epoch, after which the packet times out if it
     * has not been
     * received.
     */
    function sendPacket(bytes32 channelId, bytes calldata packet, uint64 timeoutTimestamp) external {
        // ensure port owns channel
        if (_portChannelMap[msg.sender][channelId].counterpartyChannelId == bytes32(0)) {
            revert IBCErrors.channelNotOwnedBySender();
        }

        _sendPacket(msg.sender, channelId, packet, timeoutTimestamp);
    }
    /**
     * @notice Handle the acknowledgement of an IBC packet by the counterparty
     * @dev Verifies the given proof and calls the `onAcknowledgementPacket` function on the given `receiver` contract,
     *    ie. the IBC dApp.
     *    Prerequisite: the original packet is committed and not ack'ed or timed out yet.
     * @param receiver The IbcPacketHandler contract that should handle the packet acknowledgement event
     * If the address doesn't satisfy the interface, the transaction will be reverted.
     * @param packet The IbcPacket data for the acknowledged packet
     * @param ack The acknowledgement receipt for the packet
     * @param proof The membership proof to verify the packet acknowledgement committed on Polymer chain
     */

    function acknowledgement(
        IbcPacketReceiver receiver,
        IbcPacket calldata packet,
        bytes calldata ack,
        Ics23Proof calldata proof
    ) external {
        // verify `receiver` is the original packet sender
        if (!portIdAddressMatch(address(receiver), packet.src.portId)) {
            revert IBCErrors.receiverNotOriginPacketSender();
        }

        // prove ack packet is on Polymer chain
        _getLightClientFromChannelId(packet.src.channelId).verifyMembership(
            proof, Ibc.ackProofKey(packet), abi.encode(Ibc.ackProofValue(ack))
        );
        // verify packet has been committed and not yet ack'ed or timed out
        bool hasCommitment = _sendPacketCommitment[address(receiver)][packet.src.channelId][packet.sequence];
        if (!hasCommitment) {
            revert IBCErrors.packetCommitmentNotFound();
        }

        // enforce ack'ed packet sequences always increment by 1 for ordered channels
        (bool success, bytes memory data) = _callIfContract(
            address(receiver),
            abi.encodeWithSelector(IbcPacketReceiver.onAcknowledgementPacket.selector, packet, Ibc.parseAckData(ack))
        );

        if (success) {
            if (_portChannelMap[address(receiver)][packet.src.channelId].ordering == ChannelOrder.ORDERED) {
                if (packet.sequence != _nextSequenceAck[address(receiver)][packet.src.channelId]) {
                    revert IBCErrors.unexpectedPacketSequence();
                }

                _nextSequenceAck[address(receiver)][packet.src.channelId] = packet.sequence + 1;
            }

            // delete packet commitment to avoid double ack
            delete _sendPacketCommitment[address(receiver)][packet.src.channelId][packet.sequence];
            emit Acknowledgement(address(receiver), packet.src.channelId, packet.sequence);
        } else {
            emit AcknowledgementError(address(receiver), data);
        }
    }

    /**
     * @notice Timeout of an IBC packet
     * @dev Verifies the given proof and calls the `onTimeoutPacket` function on the given `receiver` contract, ie. the
     * IBC-dApp.
     * Prerequisite: the original packet is committed and not ack'ed or timed out yet.
     * @param receiver The IbcPacketHandler contract that should handle the packet timeout event
     * If the address doesn't satisfy the interface, the transaction will be reverted.
     * @param packet The IbcPacket data for the timed-out packet
     * @param proof The non-membership proof data needed to verify the packet timeout
     */
    function timeout(IbcPacketReceiver receiver, IbcPacket calldata packet, Ics23Proof calldata proof) external {
        // verify `receiver` is the original packet sender
        if (!portIdAddressMatch(address(receiver), packet.src.portId)) {
            revert IBCErrors.receiverNotIntendedPacketDestination();
        }

        // prove absence of packet receipt on Polymer chain
        // TODO: add non membership support
        _getLightClientFromChannelId(packet.src.channelId).verifyNonMembership(
            proof, Ibc.packetCommitmentProofKey(packet)
        );

        // verify packet has been committed and not yet ack'ed or timed out
        bool hasCommitment = _sendPacketCommitment[address(receiver)][packet.src.channelId][packet.sequence];
        if (!hasCommitment) {
            revert IBCErrors.packetCommitmentNotFound();
        }

        (bool success, bytes memory data) = _callIfContract(
            address(receiver), abi.encodeWithSelector(IbcPacketReceiver.onTimeoutPacket.selector, packet)
        );
        if (success) {
            // delete packet commitment to avoid double timeout
            delete _sendPacketCommitment[address(receiver)][packet.src.channelId][packet.sequence];
            emit Timeout(address(receiver), packet.src.channelId, packet.sequence);
        } else {
            emit TimeoutError(address(receiver), data);
        }
    }

    /**
     * @notice Receive an IBC packet and then pass it to the IBC-dApp for processing if verification succeeds.
     * @dev Verifies the given proof and calls the `onRecvPacket` function on the given `receiver` contract
     * @param receiver The IbcPacketHandler contract that should handle the packet receipt event
     * If the address doesn't satisfy the interface, the transaction will be reverted.
     * The receiver must be the intended packet destination, which is the same as packet.dest.portId.
     * @param packet The IbcPacket data for the received packet
     * @param proof The proof data needed to verify the packet receipt
     * @dev Emit an `RecvPacket` event with the details of the received packet;
     * Also emit a WriteAckPacket event, which can be relayed to Polymer chain by relayers
     */
    function recvPacket(IbcPacketReceiver receiver, IbcPacket calldata packet, Ics23Proof calldata proof) external {
        // verify `receiver` is the intended packet destination
        if (!portIdAddressMatch(address(receiver), packet.dest.portId)) {
            revert IBCErrors.receiverNotIntendedPacketDestination();
        }
        _getLightClientFromChannelId(packet.dest.channelId).verifyMembership(
            proof, Ibc.packetCommitmentProofKey(packet), abi.encode(Ibc.packetCommitmentProofValue(packet))
        );

        // verify packet has not been received yet
        bool hasReceipt = _recvPacketReceipt[address(receiver)][packet.dest.channelId][packet.sequence];
        if (hasReceipt) {
            revert IBCErrors.packetReceiptAlreadyExists();
        }

        _recvPacketReceipt[address(receiver)][packet.dest.channelId][packet.sequence] = true;

        // enforce recv'ed packet sequences always increment by 1 for ordered channels
        if (_portChannelMap[address(receiver)][packet.dest.channelId].ordering == ChannelOrder.ORDERED) {
            if (packet.sequence != _nextSequenceRecv[address(receiver)][packet.dest.channelId]) {
                revert IBCErrors.unexpectedPacketSequence();
            }

            _nextSequenceRecv[address(receiver)][packet.dest.channelId] = packet.sequence + 1;
        }

        // Emit recv packet event to prove the relayer did the correct job, and pkt is received.
        emit RecvPacket(address(receiver), packet.dest.channelId, packet.sequence);

        // If pkt is already timed out, then return early so dApps won't receive it.
        if (_isPacketTimeout(packet.timeoutTimestamp, packet.timeoutHeight.revision_height)) {
            address writerPortAddress = address(receiver);
            emit WriteTimeoutPacket(
                writerPortAddress, packet.dest.channelId, packet.sequence, packet.timeoutHeight, packet.timeoutTimestamp
            );
            return;
        }

        // Not timeout yet, then do normal handling
        AckPacket memory ack;
        (bool success, bytes memory data) =
            _callIfContract(address(receiver), abi.encodeWithSelector(IbcPacketReceiver.onRecvPacket.selector, packet));
        if (success) {
            (ack) = abi.decode(data, (AckPacket));
        } else {
            ack = AckPacket(false, data);
        }
        bool hasAckPacketCommitment = _ackPacketCommitment[address(receiver)][packet.dest.channelId][packet.sequence];
        // check is not necessary for sync-acks
        if (hasAckPacketCommitment) {
            revert IBCErrors.ackPacketCommitmentAlreadyExists();
        }

        _ackPacketCommitment[address(receiver)][packet.dest.channelId][packet.sequence] = true;

        emit WriteAckPacket(address(receiver), packet.dest.channelId, packet.sequence, ack);
    }

    // TODO: add async writeAckPacket
    // // this can be invoked sync or async by the IBC-dApp
    // function writeAckPacket(IbcPacket calldata packet, AckPacket calldata ackPacket) external {
    //     // verify `receiver` is the original packet sender
    //     require(
    //         portIdAddressMatch(address(msg.sender), packet.src.portId),
    //         'Receiver is not the original packet sender'
    //     );
    // }

    // TODO: remove below writeTimeoutPacket() function
    //       1. core SC is responsible to generate timeout packet
    //       2. user contract are not free to generate timeout with different criteria
    //       3. [optional]: we may wish relayer to trigger timeout process, but in this case, belowunction won't do
    // the job, as it doesn't have proofs.
    //          There is no strong reason to do this, as relayer can always do the regular `recvPacket` flow, which will
    // do proper timeout generation.
    /**
     * Generate a timeout packet for the given packet
     */
    function writeTimeoutPacket(address receiver, IbcPacket calldata packet) external {
        // verify `receiver` is the original packet sender
        if (!portIdAddressMatch(receiver, packet.src.portId)) {
            revert IBCErrors.receiverNotIntendedPacketDestination();
        }

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
     * default
     *    values per EVM.
     */
    function getChannel(address portAddress, bytes32 channelId) external view returns (Channel memory channel) {
        channel = _portChannelMap[portAddress][channelId];
    }

    // getOptimisticConsensusState
    function getOptimisticConsensusState(uint256 height, string calldata connection)
        external
        view
        returns (uint256 appHash, uint256 fraudProofEndTime, bool ended)
    {
        return _getLightClientFromConnection(connection).getState(height);
    }

    // verify an EVM address matches an IBC portId.
    // IBC_PortID = portPrefix + address (hex string without 0x prefix, case-insensitive)
    function portIdAddressMatch(address addr, string calldata portId) public view returns (bool isMatch) {
        if (keccak256(abi.encodePacked(portPrefix)) != keccak256(abi.encodePacked(portId[0:portPrefixLen]))) {
            return false;
        }
        isMatch = Ibc._hexStrToAddress(portId[portPrefixLen:]) == addr;
    }

    function _setNewConnection(string calldata connection, LightClient lightClient) internal {
        if (bytes(connection).length == 0) {
            revert IBCErrors.invalidConnection("");
        }
        if (address(lightClient) == address(0)) {
            revert IBCErrors.invalidAddress();
        }

        _connectionToLightClient[connection] = lightClient;
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
        CounterParty calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        CounterParty calldata counterparty
    ) internal {
        // Register port and channel mapping
        // TODO: check duplicated channel registration?
        // TODO: The call to `Channel` constructor MUST be move to `openIbcChannel` phase
        //       Then `connectIbcChannel` phase can use the `version` as part of `require` condition.
        _portChannelMap[address(portAddress)][local.channelId] = Channel(
            local.portId,
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
        _channelIdToConnection[local.channelId] = connectionHops[0]; // Set channel to connection mapping for finding
            // light client
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
    // timestamp, in seconds after the unix epoch.
    function _isPacketTimeout(uint64 timeoutTimestamp, uint64 revisionHeight) internal view returns (bool isTimeOut) {
        return (
            isTimeOut = (timeoutTimestamp != 0 && block.timestamp >= timeoutTimestamp)
            // TODO: check timeoutHeight.revision_number?
            || (revisionHeight != 0 && block.number >= revisionHeight)
        );
    }

    function _getLightClientFromChannelId(bytes32 channelId) internal view returns (LightClient lightClient) {
        string memory connection = _channelIdToConnection[channelId];
        if (bytes(connection).length == 0) {
            revert IBCErrors.channelIdNotFound(channelId);
        }
        return _getLightClientFromConnection(connection);
    }

    function _getLightClientFromConnection(string memory connection) internal view returns (LightClient lightClient) {
        lightClient = _connectionToLightClient[connection];
        if (address(lightClient) == address(0)) {
            revert IBCErrors.lightClientNotFound(connection);
        }
    }
}
