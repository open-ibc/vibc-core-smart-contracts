//SPDX-License-Identifier: UNLICENSED

pragma solidity 0.8.15;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgradeable/contracts/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {IbcChannelReceiver, IbcPacketReceiver} from "../interfaces/IbcReceiver.sol";
import {L1Header, OpL2StateProof, Ics23Proof} from "../interfaces/IProofVerifier.sol";
import {ILightClient} from "../interfaces/ILightClient.sol";
import {IDispatcher} from "../interfaces/IDispatcher.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import {
    Channel, CounterParty, ChannelOrder, IbcPacket, ChannelState, AckPacket, IBCErrors, Ibc
} from "../libs/Ibc.sol";

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
contract Dispatcher is OwnableUpgradeable, UUPSUpgradeable, ReentrancyGuard, IDispatcher {
    // Gap to allow for additional contract inheritance, similar to OpenZeppelin's Initializable contract
    uint256[49] private __gap;

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
    mapping(string => ILightClient) private _connectionToLightClient;
    uint256 private _numClients;

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initializes the Dispatcher contract with the provided port prefix.
     * @param initPortPrefix The initial port prefix to be set for the contract.
     * @dev This method should be called only once during contract deployment.
     * @dev For contract upgarades, which need to reinitialize the contract, use the reinitializer modifier.
     */
    function initialize(string memory initPortPrefix) public virtual initializer {
        if (bytes(initPortPrefix).length == 0) {
            revert IBCErrors.invalidPortPrefix();
        }
        __Ownable_init();
        portPrefix = initPortPrefix;
        portPrefixLen = uint32(bytes(initPortPrefix).length);
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
     * This function is called by a 'relayer' on behalf of a dApp. The dApp should implement IbcChannelHandler's
     * onChanOpenInit. If the callback succeeds, the dApp should return the selected version and the emitted event
     * will be relayed to the  IBC/VIBC hub chain.
     */
    /**
     * @notice Initializes the channel opening process with the specified parameters. This is the first step in the  channel
     * handshake, initiated directly by the dapp which wishes to establish a channel with the receiver.
     * @param receiver The address of receiver contract that will handle the channel opening process.
     * @param ordering The ordering of the channel (ORDERED or UNORDERED).
     * @param feeEnabled A boolean indicating whether fees are enabled for the channel.
     * @param connectionHops The list of connection hops associated with the channel, with the first channel in this
     * array always starting from the chain this contract is deployed on
     * @param counterpartyPortId The port ID of the counterparty.
     * @dev This function initializes the channel opening process by calling the onChanOpenInit function of the
     *      specified receiver contract.
     *      It can only be called by authorized parties.
     */
    function channelOpenInit(
        IbcChannelReceiver receiver,
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) external nonReentrant {
        if (address(receiver) == address(0)) {
            revert IBCErrors.invalidAddress();
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
     * @notice Initializes step 2 of the channel handshake process. This is called by a relayer on behalf of the
     * receiving dapp.
     * To minimize trust assumptions in proving that the counterparty has indeed initiated step 1, there must be a valid
     * proof of the counterparty being in the TRY_PENDING state from the Polymer Client
     * @notice   The receiving dApp should implement IbcChannelHandler's onChanOpenTry callback. If the callback
     * succeeds, the dApp should return the selected version and the emitted even will be relayed to the  IBC/VIBC hub
     * chain.
     * @param receiver The dapp address that the sender is attempting to initiate a channel with
     * @param local The counterparty information for the receiver.
     * @param ordering The ordering of the channel (ORDERED or UNORDERED).
     * @param feeEnabled Whether fees are enabled for the channel.
     * @param connectionHops The list of connection hops associated with the channel; with the first channel in this
     * array always starting from the chain this contract is deployed on
     * @param counterparty The counterparty information of the sender
     * @param proof The proof that the counterparty is in the TRY_PENDING state (i.e. that it is indeed intending to
     * initialize a channel with the given receiver)
     */
    function channelOpenTry(
        IbcChannelReceiver receiver,
        CounterParty calldata local,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        CounterParty calldata counterparty,
        Ics23Proof calldata proof
    ) external nonReentrant {
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
     * @notice Initializes step 3 of the channel handshake process. This method is called by a relayer on behalf of the
     * sending, dapp (i.e. the dapp which initiated the channelOpenInit call). This step happens after the IBC/VIBC hub
     * chain has processed the ChannelOpenTry event.
     * To minimize trust assumptions in proving that the counterparty had indeed responded successfully in step 2 of the
     * handshake, there must be a valid proof of the counterparty being in the ACK_PENDING state from the Polymer Client
     * @notice Completes the channel opening acknowledge process with the specified parameters.
     * @notice The dApp should implement the onChannelOpenAck method to handle the third channel handshake method:
     * ChanOpenAck process.
     * @param receiver The address of the IbcChannelReceiver contract that will handle the channel opening acknowledge
     * @param local The counterparty information for the local channel.
     * @param connectionHops The list of connection hops associated with the channel, with the first channel in this
     * array always starting from the chain this contract is deployed on.
     * @param ordering The ordering of the channel (ORDERED or UNORDERED).
     * @param feeEnabled A boolean indicating whether fees are enabled for the channel.
     * @param counterparty The counterparty information for the channel.
     * @param proof The proof that the counterparty is in the ACK_PENDING state (i.e. that it responded with a
     * successful channelOpenTry )
     */
    function channelOpenAck(
        IbcChannelReceiver receiver,
        CounterParty calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        CounterParty calldata counterparty,
        Ics23Proof calldata proof
    ) external nonReentrant {
        if (
            bytes(local.portId).length == 0 || bytes(counterparty.portId).length == 0 || local.channelId == bytes32(0)
                || counterparty.channelId == bytes32(0)
        ) {
            revert IBCErrors.invalidCounterParty();
        }
        if (address(receiver) == address(0)) {
            revert IBCErrors.invalidAddress();
        }
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
     * @notice Initializes step 4 of the channel handshake process. This method is called by a relayer on behalf of the
     * receiving dapp. This step happens after the IBC/VIBC hub chain has processed the ChannelOpenAck event.
     * To minimize trust assumptions in proving that the counterparty had indeed responded successfully in step 2 of the
     * handshake, there must be a valid proof of the counterparty being in the CONFIRM_PENDING state from the Polymer
     * Client
     *
     * @notice Initiates the channel opening confirm process with the specified parameters.
     * @param receiver The address of the IbcChannelReceiver contract that will handle the channel opening confirm
     * process. The receiver should implement the onChannelConnect method to handle the last channel handshake method:
     * ChannelOpenConfirm
     * @param local The counterparty information for the local channel.
     * @param ordering The ordering of the channel (ORDERED or UNORDERED).
     * @param connectionHops The list of connection hops associated with the channel, with the first channel in this
     * array always starting from the chain this contract is deployed on.
     * @param counterparty The counterparty information for the channel.
     * @param proof The proof of channel opening confirm.
     * @dev This function initiates the channel opening confirm process by calling the onChanOpenConfirm function of the
     * specified
     *      receiver contract.
     *      It can only be called by authorized parties.
     */
    function channelOpenConfirm(
        IbcChannelReceiver receiver,
        CounterParty calldata local,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        CounterParty calldata counterparty,
        Ics23Proof calldata proof
    ) external nonReentrant {
        if (
            bytes(local.portId).length == 0 || bytes(counterparty.portId).length == 0 || local.channelId == bytes32(0)
                || counterparty.channelId == bytes32(0)
        ) {
            revert IBCErrors.invalidCounterParty();
        }

        if (address(receiver) == address(0)) {
            revert IBCErrors.invalidAddress();
        }

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
     * @notice Initializes a close channel handshake process. It is directly called by the dapp which wants to close
     * the channel
     * @dev Emits a `CloseIbcChannel` event with the given `channelId` and the address of the message sender
     * @notice Close the specified IBC channel by channel ID
     * @notice Must be called by the channel owner, ie. _portChannelMap[msg.sender][channelId] must exist
     */
    function channelCloseInit(bytes32 channelId) external nonReentrant {
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

    /**
     * @notice Sends an IBC packet on a existing channel with the specified packet data and timeout block timestamp.
     * @dev Emits an `IbcPacketEvent` event containing the sender address, channel ID, packet data, and timeout block
     * timestamp (formatted as seconds after the unix epoch).
     * @param channelId The ID of the channel on which to send the packet.
     * @param packet The packet data to send.
     * @param timeoutTimestamp The timestamp, in seconds after the unix epoch, after which the packet times out if it
     * has not been received.
     */
    function sendPacket(bytes32 channelId, bytes calldata packet, uint64 timeoutTimestamp) external {
        // ensure port owns channel
        if (_portChannelMap[msg.sender][channelId].counterpartyChannelId == bytes32(0)) {
            revert IBCErrors.channelNotOwnedBySender();
        }

        _sendPacket(msg.sender, channelId, packet, timeoutTimestamp);
    }

    /**
     * @notice Handle the acknowledgement of an IBC packet by the counterparty, called by a relayer.
     * @dev To minimize trust assumptions, an inclusion proof must be given that proves the packet commitment is
     * in the IBC/VIBC hub chain.
     * @dev This method also calls the `onAcknowledgementPacket` function on
     * the given `receiver` contract,
     * @dev Prerequisite: The original packet must be committed and not ack'ed or timed out yet.
     * @param receiver The IbcPacketHandler contract that should handle the packet acknowledgement event
     * If the address doesn't satisfy the interface, the nextSequenceAck will not be updated
     * @param packet The IbcPacket data for the acknowledged packet
     * @param ack The acknowledgement receipt for the packet
     * @param proof The membership proof to verify the packet acknowledgement committed on Polymer chain
     */
    function acknowledgement(
        IbcPacketReceiver receiver,
        IbcPacket calldata packet,
        bytes calldata ack,
        Ics23Proof calldata proof
    ) external nonReentrant {
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
    function timeout(IbcPacketReceiver receiver, IbcPacket calldata packet, Ics23Proof calldata proof)
        external
        nonReentrant
    {
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
     * @dev To minimize trust assumptions, verifies the given proof
     * @dev calls the `onRecvPacket` function on the given receiver contract
     * If the address doesn't satisfy the interface, the nextSequenceRecv will not be updated
     * The receiver must be the intended packet destination, which is the same as packet.dest.portId.
     * @param packet The IbcPacket data for the received packet
     * @param proof The proof data needed to verify the packet receipt
     * @dev Emit an `RecvPacket` event with the details of the received packet;
     * Also emit a WriteAckPacket event, which can be relayed to Polymer chain by relayers
     */
    function recvPacket(IbcPacketReceiver receiver, IbcPacket calldata packet, Ics23Proof calldata proof)
        external
        nonReentrant
    {
        if (address(receiver) == address(0)) {
            revert IBCErrors.invalidAddress();
        }
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

    /**
     * @notice Writes a timeout packet to the specified receiver.
     * @param receiver The address of the receiver contract.
     * @param packet The timeout packet data.
     * @dev Requirements:
     * - The `receiver` address must not be the zero address.
     * - The `receiver` must be the original packet sender.
     * - The packet must not have a receipt.
     * - The packet must have timed out.
     */
    function writeTimeoutPacket(address receiver, IbcPacket calldata packet) external {
        if (address(receiver) == address(0)) {
            revert IBCErrors.invalidAddress();
        }
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
     * default values per EVM.
     */
    function getChannel(address portAddress, bytes32 channelId) external view returns (Channel memory channel) {
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

    // verify an EVM address matches an IBC portId.
    // IBC_PortID = portPrefix + address (hex string without 0x prefx, case-insensitive)
    function portIdAddressMatch(address addr, string calldata portId) public view returns (bool isMatch) {
        if (keccak256(abi.encodePacked(portPrefix)) != keccak256(abi.encodePacked(portId[0:portPrefixLen]))) {
            return false;
        }
        isMatch = Ibc._hexStrToAddress(portId[portPrefixLen:]) == addr;
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
     * @notice Sends a packet on the specified channel with the provided details.
     * @param sender The address of the sender.
     * @param channelId The ID of the channel.
     * @param packet The packet data to be sent.
     * @param timeoutTimestamp The timeout timestamp for the packet.
     */
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

    /**
     * @notice Connects a channel with the provided parameters.
     * @param portAddress The address of the IbcChannelReceiver contract.
     * @param local The details of the local counterparty.
     * @param connectionHops The connection hops associated with the channel, with the first channel in this
     * array always starting from the chain this contract is deployed on.
     * @param ordering The ordering of the channel.
     * @param feeEnabled A boolean indicating whether fees are enabled for the channel.
     * @param counterparty The details of the counterparty.
     */
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
        (success, message) = receiver.call(args);
    }

    /**
     * @inheritdoc UUPSUpgradeable
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

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
        lightClient = _connectionToLightClient[connection];
        if (address(lightClient) == address(0)) {
            revert IBCErrors.lightClientNotFound(connection);
        }
    }
}
