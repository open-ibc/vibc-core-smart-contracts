//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./IbcDispatcher.sol";
import {IbcChannelReceiver, IbcPacketReceiver} from "./IbcReceiver.sol";
import "./IbcVerifier.sol";
import "./ConsensusStateManager.sol";

// InitClientMsg is used to create a new Polymer client on an EVM chain
// TODO: replace bytes with explictly typed fields for gas cost saving
struct InitClientMsg {
    bytes clientState;
    ConsensusState consensusState;
}

// UpgradeClientMsg is used to upgrade an existing Polymer client on an EVM chain.
// It should only be run by CoreSC maintainer with a social consensus.
// TODO: replace bytes with explictly typed fields for gas cost saving
struct UpgradeClientMsg {
    bytes clientState;
    ConsensusState consensusState;
}

/**
 * @title Dispatcher
 * @author Polymer Labs
 * @notice
 *     Contract callers call this contract to send IBC-like msg,
 *     which can be relayed to a rollup module on the Polymerase chain
 */
contract Dispatcher is IbcDispatcher, IbcEventsEmitter, Ownable {
    //
    // fields
    //
    ZKMintVerifier verifier;
    bool isClientCreated = false;
    ConsensusState public latestConsensusState;
    // IBC_PortID = portPrefix + address (hex string without 0x prefix, case insensitive)
    string public portPrefix;
    uint32 portPrefixLen;

    mapping(address => mapping(bytes32 => Channel)) public portChannelMap;
    mapping(address => mapping(bytes32 => uint64)) nextSequenceSend;
    // keep track of received packets' sequences to ensure channel ordering is enforced for ordered channels
    mapping(address => mapping(bytes32 => uint64)) nextSequenceRecv;
    mapping(address => mapping(bytes32 => uint64)) nextSequenceAck;
    // only stores a bit to mark packet has not been ack'ed or timed out yet; actual IBC packet verification is done on
    // Polymer chain.
    // Keep track of sent packets
    mapping(address => mapping(bytes32 => mapping(uint64 => bool))) sendPacketCommitment;
    // keep track of received packets to prevent replay attack
    mapping(address => mapping(bytes32 => mapping(uint64 => bool))) recvPacketReceipt;
    // keep track of outbound ack packets to prevent replay attack
    mapping(address => mapping(bytes32 => mapping(uint64 => bool))) ackPacketCommitment;

    ConsensusStateManager consensusStateManager;

    //
    // methods
    //
    constructor(ZKMintVerifier _verifier, string memory initPortPrefix, ConsensusStateManager _consensusStateManager) {
        verifier = _verifier;
        portPrefix = initPortPrefix;
        portPrefixLen = uint32(bytes(initPortPrefix).length);
        consensusStateManager = _consensusStateManager;
    }

    //
    // Utility functions
    //

    /**
     * Convert a non-0x-prefixed hex string to an address
     * @param hexStr hex string to convert to address. Note that the hex string must not include a 0x prefix.
     * hexStr is case-insensitive.
     */
    function hexStrToAddress(string memory hexStr) internal pure returns (address) {
        if (bytes(hexStr).length != 40) {
            revert invalidHexStringLength();
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

        address addr;
        assembly {
            addr := mload(add(addrBytes, 20))
        }

        return addr;
    }

    // verify an EVM address matches an IBC portId.
    // IBC_PortID = portPrefix + address (hex string without 0x prefix, case-insensitive)
    function portIdAddressMatch(address addr, string calldata portId) public view returns (bool) {
        if (keccak256(abi.encodePacked(portPrefix)) != keccak256(abi.encodePacked(portId[0:portPrefixLen]))) {
            return false;
        }
        string memory portSuffix = portId[portPrefixLen:];
        return hexStrToAddress(portSuffix) == addr;
    }

    //
    // CoreSC maaintainer methods, only invoked by the owner
    //
    function setPortPrefix(string calldata _portPrefix) external onlyOwner {
        portPrefix = _portPrefix;
        portPrefixLen = uint32(bytes(_portPrefix).length);
    }

    //
    // Client methods
    //

    /**
     * @dev Creates a new Polymer client.
     * @param initClientMsg The initial client state and consensus state.
     */
    function createClient(InitClientMsg calldata initClientMsg) external onlyOwner {
        if (isClientCreated) {
            revert clientAlreadyCreated();
        }

        isClientCreated = true;
        latestConsensusState = initClientMsg.consensusState;
    }

    // updateClientWithConsensusState updates the client with the
    // latest consensus state. The zkproof related to this consensus
    // state is used to verify the consensus state.
    function updateClientWithConsensusState(ConsensusState calldata consensusState, ZkProof calldata zkProof)
        external
    {
        if (!isClientCreated) {
            revert clientNotCreated();
        }
        if (!verifier.verifyConsensusState(latestConsensusState, consensusState, zkProof)) {
            revert consensusStateVerificationFailed();
        }

        latestConsensusState = consensusState;
        return;
    }

    // updateClientWithOptimisticConsensusState updates the client
    // with the optimistic consensus state. The optimistic consensus
    // is accepted and will be open for verify in the fraud proof
    // window.
    function updateClientWithOptimisticConsensusState(uint256 height, uint256 appHash)
        external
        returns (uint256 fraudProofEndTime, bool ended)
    {
        if (!isClientCreated) {
            revert clientNotCreated();
        }

        return consensusStateManager.addOpConsensusState(height, appHash);
    }

    // getOptimisticConsensusState
    function getOptimisticConsensusState(uint256 height)
        external
        view
        returns (uint256 appHash, uint256 fraudProofEndTime, bool ended)
    {
        if (!isClientCreated) {
            revert clientNotCreated();
        }

        return consensusStateManager.getState(height);
    }

    /**
     * @dev Upgrades the Polymer client.
     * It can only be run by CoreSC maintainer with a social consensus.
     * @param upgradeClientMsg The new client state and consensus state.
     */
    function upgradeClient(UpgradeClientMsg calldata upgradeClientMsg) external onlyOwner {
        if (!isClientCreated) {
            revert clientNotCreated();
        }

        latestConsensusState = upgradeClientMsg.consensusState;
    }

    //
    // Utility functions
    //

    // return the concatenation of two strings in bytes
    function concatStrings(string memory str1, string memory str2) internal pure returns (bytes memory) {
        return abi.encodePacked(str1, str2);
    }

    function max(uint256 a, uint256 b) internal pure returns (uint256) {
        return a > b ? a : b;
    }

    //
    // IBC Channel methods
    //

    /**
     * This func is called by a 'relayer' on behalf of a dApp. The dApp should be implements IbcChannelHandler.
     * The dApp should implement the onOpenIbcChannel method to handle one of the first two channel handshake methods,
     * ie. ChanOpenInit or ChanOpenTry.
     * If callback succeeds, the dApp should return the selected version, and an emitted event will be relayed to the
     * IBC/VIBC hub chain.
     */
    function openIbcChannel(
        IbcChannelReceiver portAddress,
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        CounterParty calldata counterparty,
        Proof calldata proof
    ) external {
        if (bytes(counterparty.portId).length == 0) {
            revert invalidCounterPartyPortId();
        }

        // For XXXX => vIBC direction, SC needs to verify the proof of membership of TRY_PENDING
        // For vIBC initiated channel, SC doesn't need to verify any proof, and these should be all empty
        bool isChanOpenTry = false;
        if (counterparty.channelId == bytes32(0) && bytes(counterparty.version).length == 0) {
            // ChanOpenInit with unknow conterparty
        } else if (counterparty.channelId != bytes32(0) && bytes(counterparty.version).length != 0) {
            // this is the ChanOpenTry; counterparty must not be zero-value
            isChanOpenTry = true;
        } else {
            revert invalidCounterParty();
        }

        if (isChanOpenTry) {
            // TODO: fill below proof path
            consensusStateManager.verifyMembership(
                proof,
                "channel/path/to/be/added/here",
                bytes("expected channel bytes constructed from params. Channel.State = {Try_Pending}"),
                "Fail to prove channel state"
            );
        }

        string memory selectedVersion = portAddress.onOpenIbcChannel(
            version,
            ordering,
            feeEnabled,
            connectionHops,
            counterparty.portId,
            counterparty.channelId,
            counterparty.version
        );

        emit OpenIbcChannel(
            address(portAddress),
            selectedVersion,
            ordering,
            feeEnabled,
            connectionHops,
            counterparty.portId,
            counterparty.channelId
        );
    }

    /**
     * This func is called by a 'relayer' after the IBC/VIBC hub chain has processed the onOpenIbcChannel event.
     * The dApp should implement the onConnectIbcChannel method to handle the last two channel handshake methods, ie.
     * ChanOpenAck or ChanOpenConfirm.
     */
    function connectIbcChannel(
        IbcChannelReceiver portAddress,
        bytes32 channelId,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        bool feeEnabled,
        string calldata counterpartyPortId,
        bytes32 counterpartyChannelId,
        string calldata counterpartyVersion,
        Proof calldata proof
    ) external {
        consensusStateManager.verifyMembership(
            proof,
            bytes("channel/path/to/be/added/here"),
            bytes("expected channel bytes constructed from params. Channel.State = {Ack_Pending, Confirm_Pending}"),
            "Fail to prove channel state"
        );

        portAddress.onConnectIbcChannel(channelId, counterpartyChannelId, counterpartyVersion);

        // Register port and channel mapping
        // TODO: check duplicated channel registration?
        // TODO: The call to `Channel` constructor MUST be move to `openIbcChannel` phase
        //       Then `connectIbcChannel` phase can use the `version` as part of `require` condition.
        portChannelMap[address(portAddress)][channelId] = Channel(
            counterpartyVersion, // TODO: this should be self version instead of counterparty version
            ordering,
            feeEnabled,
            connectionHops,
            counterpartyPortId,
            counterpartyChannelId
        );

        // initialize channel sequences
        nextSequenceSend[address(portAddress)][channelId] = 1;
        nextSequenceRecv[address(portAddress)][channelId] = 1;
        nextSequenceAck[address(portAddress)][channelId] = 1;

        emit ConnectIbcChannel(address(portAddress), channelId);
    }

    /**
     * @notice Get the IBC channel with the specified port and channel ID
     * @param portAddress EVM address of the IBC port
     * @param channelId IBC channel ID from the port perspective
     * @return A channel struct is always returned. If it doesn't exists, the channel struct is populated with default
     *    values per EVM.
     */
    function getChannel(address portAddress, bytes32 channelId) external view returns (Channel memory) {
        return portChannelMap[portAddress][channelId];
    }

    /**
     * @dev Emits a `CloseIbcChannel` event with the given `channelId` and the address of the message sender
     * @notice Close the specified IBC channel by channel ID
     * Must be called by the channel owner, ie. portChannelMap[msg.sender][channelId] must exist
     */
    function closeIbcChannel(bytes32 channelId) external {
        Channel memory channel = portChannelMap[msg.sender][channelId];
        if (channel.counterpartyChannelId == bytes32(0)) {
            revert channelNotOwnedBySender();
        }

        IbcChannelReceiver reciever = IbcChannelReceiver(msg.sender);
        reciever.onCloseIbcChannel(channelId, channel.counterpartyPortId, channel.counterpartyChannelId);
        emit CloseIbcChannel(msg.sender, channelId);
    }

    /**
     * This func is called by a 'relayer' after the IBC/VIBC hub chain has processed ChanCloseConfirm event.
     * The dApp's onCloseIbcChannel callback is invoked.
     * dApp should throw an error if the channel should not be closed.
     */
    function onCloseIbcChannel(address portAddress, bytes32 channelId, Proof calldata proof) external {
        // verify VIBC/IBC hub chain has processed ChanCloseConfirm event
        consensusStateManager.verifyMembership(
            proof,
            bytes("channel/path/to/be/added/here"),
            bytes("expected channel bytes constructed from params. Channel.State = {Closed(_Pending?)}"),
            "Fail to prove channel state"
        );

        // ensure port owns channel
        Channel memory channel = portChannelMap[portAddress][channelId];
        if (channel.counterpartyChannelId == bytes32(0)) {
            revert channelNotOwnedByPortAddress();
        }

        // confirm with dApp by calling its callback
        IbcChannelReceiver reciever = IbcChannelReceiver(portAddress);
        reciever.onCloseIbcChannel(channelId, channel.counterpartyPortId, channel.counterpartyChannelId);
        delete portChannelMap[portAddress][channelId];
        emit CloseIbcChannel(portAddress, channelId);
    }

    //
    // IBC Packet methods
    //

    /**
     * @notice Sends an IBC packet on a existing channel with the specified packet data and timeout block timestamp.
     * @notice Data should be encoded in a format defined by the channel version, and the module on the other side should know how to parse this.
     * @dev Emits an `IbcPacketEvent` event containing the sender address, channel ID, packet data, and timeout block timestamp.
     * @param channelId The ID of the channel on which to send the packet.
     * @param packet The packet data to send.
     * @param timeoutTimestamp The timestamp in nanoseconds after which the packet times out if it has not been received.
     */
    function sendPacket(bytes32 channelId, bytes calldata packet, uint64 timeoutTimestamp) external {
        // ensure port owns channel
        Channel memory channel = portChannelMap[msg.sender][channelId];
        if (channel.counterpartyChannelId == bytes32(0)) {
            revert channelNotOwnedBySender();
        }

        _sendPacket(msg.sender, channelId, packet, timeoutTimestamp);
    }

    // Prerequisite: must verify sender is authorized to send packet on the channel
    function _sendPacket(address sender, bytes32 channelId, bytes memory packet, uint64 timeoutTimestamp) internal {
        // current packet sequence
        uint64 sequence = nextSequenceSend[sender][channelId];
        if (sequence == 0) {
            revert invalidPacketSequence();
        }

        // packet commitment
        sendPacketCommitment[sender][channelId][sequence] = true;
        // increment nextSendPacketSequence
        nextSequenceSend[sender][channelId] = sequence + 1;

        emit SendPacket(sender, channelId, packet, sequence, timeoutTimestamp);
    }

    /**
     * @notice Handle the acknowledgement of an IBC packet by the counterparty
     * @dev Verifies the given proof and calls the `onAcknowledgementPacket` function on the given `receiver` contract,
     *    ie. the IBC dApp.
     *    Prerequisite: the original packet is committed and not ack'ed or timed out yet.
     * @param receiver The IbcPacketHandler contract that should handle the packet acknowledgement event
     * If the address doesn't satisfy the interface, the transaction will be reverted.
     * @param packet The IbcPacket data for the acknowledged packet
     * @param ackPacket The acknowledgement receipt for the packet
     * @param proof The membership proof to verify the packet acknowledgement committed on Polymer chain
     */
    function acknowledgement(
        IbcPacketReceiver receiver,
        IbcPacket calldata packet,
        AckPacket calldata ackPacket,
        Proof calldata proof
    ) external {
        // verify `receiver` is the original packet sender
        if (!portIdAddressMatch(address(receiver), packet.src.portId)) {
            revert receiverNotOriginPacketSender();
        }

        // prove ack packet is on Polymer chain
        consensusStateManager.verifyMembership(
            proof, bytes("ack/packet/path"), bytes("expected ack receipt hash on Polymer chain"), "Fail to prove ack"
        );
        // verify packet has been committed and not yet ack'ed or timed out
        bool hasCommitment = sendPacketCommitment[address(receiver)][packet.src.channelId][packet.sequence];
        if (!hasCommitment) {
            revert packetCommitmentNotFound();
        }

        // enforce ack'ed packet sequences always increment by 1 for ordered channels
        Channel memory channel = portChannelMap[address(receiver)][packet.src.channelId];

        if (channel.ordering == ChannelOrder.ORDERED) {
            if (packet.sequence != nextSequenceAck[address(receiver)][packet.src.channelId]) {
                revert unexpectedPacketSequence();
            }

            nextSequenceAck[address(receiver)][packet.src.channelId] = packet.sequence + 1;
        }

        receiver.onAcknowledgementPacket(packet, ackPacket);

        // delete packet commitment to avoid double ack
        delete sendPacketCommitment[address(receiver)][packet.src.channelId][packet.sequence];

        emit Acknowledgement(address(receiver), packet.src.channelId, packet.sequence);
    }

    /**
     * @notice Timeout of an IBC packet
     * @dev Verifies the given proof and calls the `onTimeoutPacket` function on the given `receiver` contract, ie. the IBC-dApp.
     * Prerequisite: the original packet is committed and not ack'ed or timed out yet.
     * @param receiver The IbcPacketHandler contract that should handle the packet timeout event
     * If the address doesn't satisfy the interface, the transaction will be reverted.
     * @param packet The IbcPacket data for the timed-out packet
     * @param proof The non-membership proof data needed to verify the packet timeout
     */
    function timeout(IbcPacketReceiver receiver, IbcPacket calldata packet, Proof calldata proof) external {
        // verify `receiver` is the original packet sender
        if (!portIdAddressMatch(address(receiver), packet.src.portId)) {
            revert receiverNotIndtendedPacketDestination();
        }

        // prove absence of packet receipt on Polymer chain
        consensusStateManager.verifyNonMembership(proof, "packet/receipt/path", "Fail to prove timeout");

        // verify packet has been committed and not yet ack'ed or timed out
        bool hasCommitment = sendPacketCommitment[address(receiver)][packet.src.channelId][packet.sequence];
        if (!hasCommitment) {
            revert packetCommitmentNotFound();
        }

        receiver.onTimeoutPacket(packet);

        // delete packet commitment to avoid double timeout
        delete sendPacketCommitment[address(receiver)][packet.src.channelId][packet.sequence];

        emit Timeout(address(receiver), packet.src.channelId, packet.sequence);
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
    function recvPacket(IbcPacketReceiver receiver, IbcPacket calldata packet, Proof calldata proof) external {
        // verify `receiver` is the intended packet destination
        if (!portIdAddressMatch(address(receiver), packet.dest.portId)) {
            revert receiverNotIndtendedPacketDestination();
        }
        consensusStateManager.verifyMembership(
            proof,
            "packet/commitment/path",
            "expected virtual packet commitment hash on Polymer chain",
            "Fail to prove packet commitment"
        );

        // verify packet has not been received yet
        bool hasReceipt = recvPacketReceipt[address(receiver)][packet.dest.channelId][packet.sequence];
        if (hasReceipt) {
            revert packetReceiptAlreadyExists();
        }

        recvPacketReceipt[address(receiver)][packet.dest.channelId][packet.sequence] = true;

        // enforce recv'ed packet sequences always increment by 1 for ordered channels
        Channel memory channel = portChannelMap[address(receiver)][packet.dest.channelId];
        if (channel.ordering == ChannelOrder.ORDERED) {
            if (packet.sequence != nextSequenceRecv[address(receiver)][packet.dest.channelId]) {
                revert unexpectedPacketSequence();
            }

            nextSequenceRecv[address(receiver)][packet.dest.channelId] = packet.sequence + 1;
        }

        // Emit recv packet event to prove the relayer did the correct job, and pkt is received.
        emit RecvPacket(address(receiver), packet.dest.channelId, packet.sequence);

        // If pkt is already timed out, then return early so dApps won't receive it.
        if (isPacketTimeout(packet)) {
            address writerPortAddress = address(receiver);
            emit WriteTimeoutPacket(
                writerPortAddress, packet.dest.channelId, packet.sequence, packet.timeoutHeight, packet.timeoutTimestamp
            );
            return;
        }

        // Not timeout yet, then do normal handling
        IbcPacket memory pkt = packet;
        AckPacket memory ack = receiver.onRecvPacket(pkt);
        bool hasAckPacketCommitment = ackPacketCommitment[address(receiver)][packet.dest.channelId][packet.sequence];
        // check is not necessary for sync-acks
        if (hasAckPacketCommitment) {
            revert ackPacketCommitmentAlreadyExists();
        }

        ackPacketCommitment[address(receiver)][packet.dest.channelId][packet.sequence] = true;

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

    // isPacketTimeout returns true if the given packet has timed out acoording to host chain's block height and timestamp
    function isPacketTimeout(IbcPacket calldata packet) internal view returns (bool) {
        return (
            (packet.timeoutTimestamp != 0 && block.timestamp >= packet.timeoutTimestamp)
            // TODO: check timeoutHeight.revision_number?
            || (packet.timeoutHeight.revision_height != 0 && block.number >= packet.timeoutHeight.revision_height)
        );
    }

    // TODO: remove below writeTimeoutPacket() function
    //       1. core SC is responsible to generate timeout packet
    //       2. user contract are not free to generate timeout with different criteria
    //       3. [optional]: we may wish relayer to trigger timeout process, but in this case, below function won't do the job, as it doesn't have proofs.
    //          There is no strong reason to do this, as relayer can always do the regular `recvPacket` flow, which will do proper timeout generation.
    /**
     * Generate a timeout packet for the given packet
     */
    function writeTimeoutPacket(address receiver, IbcPacket calldata packet) external {
        // verify `receiver` is the original packet sender
        if (!portIdAddressMatch(receiver, packet.src.portId)) {
            revert receiverNotIndtendedPacketDestination();
        }

        // verify packet does not have a receipt
        bool hasReceipt = recvPacketReceipt[receiver][packet.dest.channelId][packet.sequence];
        if (hasReceipt) {
            revert packetReceiptAlreadyExists();
        }

        // verify packet has timed out; zero-value in packet.timeout means no timeout set
        if (!isPacketTimeout(packet)) {
            revert packetNotTimedOut();
        }

        emit WriteTimeoutPacket(
            receiver, packet.dest.channelId, packet.sequence, packet.timeoutHeight, packet.timeoutTimestamp
        );
    }
}
