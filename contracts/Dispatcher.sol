//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import '@openzeppelin/contracts/utils/Strings.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import 'hardhat/console.sol';

import './IbcDispatcher.sol';
import './IbcReceiver.sol';
import './IbcVerifier.sol';

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

struct Channel {
    bytes32 version;
    ChannelOrder ordering;
    string[] connectionHops;
    string counterpartyPortId;
    bytes32 counterpartyChannelId;
}

/**
 * @title Dispatcher
 * @author Polymer Labs
 * @notice
 *     Contract callers call this contract to send IBC-like msg,
 *     which can be relayed to a rollup module on the Polymerase chain
 */
contract Dispatcher is IbcDispatcher, Ownable {
    //
    // channel events
    //

    event OpenIbcChannel(
        address indexed portAddress,
        string version,
        ChannelOrder ordering,
        string[] connectionHops,
        string counterpartyPortId,
        string counterpartyChannelId
    );

    event ConnectIbcChannel(
        address indexed portAddress,
        bytes32 indexed channelId,
        string counterpartyPortId,
        bytes32 indexed counterpartyChannelId,
        string[] connectionHops
    );

    event CloseIbcChannel(address indexed portAddress, bytes32 indexed channelId);

    //
    // packet events
    //

    event SendPacket(
        address indexed sourcePortAddress,
        bytes32 indexed sourceChannelId,
        bytes packet,
        uint64 sequence,
        // timeoutTimestamp is in UNIX nano seconds; packet will be rejected if
        // delivered after this timestamp on the receiving chain.
        // Timeout semantics is compliant to IBC spec and ibc-go implementation
        uint64 timeoutTimestamp,
        uint256 fee
    );

    event Acknowledgement(
        address indexed sourcePortAddress,
        bytes32 indexed sourceChannelId,
        AckPacket AckPacket,
        uint64 sequence
    );

    event Timeout(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 indexed sequence);

    event RecvPacket(
        address indexed destPortAddress,
        bytes32 indexed destChannelId,
        string srcPortId,
        bytes32 indexed srcChannelId,
        uint64 sequence
    );

    event WriteAckPacket(
        address indexed writerPortAddress,
        bytes32 indexed writerChannelId,
        uint64 sequence,
        AckPacket ackPacket
    );

    event WriteTimeoutPacket(address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence);

    //
    // fields
    //

    ZKMintVerifier verifier;
    address payable escrow;
    bool isClientCreated = false;
    ConsensusState public latestConsensusState;
    // IBC_PortID = portPrefix + address (hex string without 0x prefix, case insensitive)
    string portPrefix;
    uint32 portPrefixLen;

    mapping(address => mapping(bytes32 => Channel)) public portChannelMap;
    mapping(address => mapping(bytes32 => uint64)) nextSendPacketSequence;
    // only stores a bit to mark packet has not been ack'ed or timed out yet; actual IBC packet verification is done on
    // Polymer chain.
    // Keep track of sent packets
    mapping(address => mapping(bytes32 => mapping(uint64 => bool))) sendPacketCommitment;
    // keep track of received packets to prevent replay attack
    mapping(address => mapping(bytes32 => mapping(uint64 => bool))) recvPacketReceipt;
    // keep track of received packets' sequences to ensure channel ordering is enforced for ordered channels
    mapping(address => mapping(bytes32 => uint64)) nextRecvPacketSequence;
    // keep track of outbound ack packets to prevent replay attack
    mapping(address => mapping(bytes32 => mapping(uint64 => bool))) ackPacketCommitment;

    //
    // methods
    //

    constructor(ZKMintVerifier _verifier, address payable _escrow, string memory initPortPrefix) {
        verifier = _verifier;
        escrow = _escrow;
        require(escrow != address(0), 'Escrow cannot be zero address');

        // initialize portPrefix
        portPrefix = initPortPrefix;
        portPrefixLen = uint32(bytes(initPortPrefix).length);
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
        require(bytes(hexStr).length == 40, 'Invalid hex string length');

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
        require(!isClientCreated, 'Client already created');
        isClientCreated = true;
        latestConsensusState = initClientMsg.consensusState;
    }

    /**
     * @dev Updates the Polymer client.
     *
     * Requirements:
     * - The consensus state must pass zkProof verification.
     *
     * @param updateClientMsg The new consensus state for the client.
     */
    function updateClient(UpdateClientMsg calldata updateClientMsg) external {
        require(isClientCreated, 'Client not created');
        require(
            verifier.verifyUpdateClientMsg(latestConsensusState, updateClientMsg),
            'UpdateClientMsg proof verification failed'
        );
        latestConsensusState = updateClientMsg.consensusState;
    }

    /**
     * @dev Upgrades the Polymer client.
     * It can only be run by CoreSC maintainer with a social consensus.
     * @param upgradeClientMsg The new client state and consensus state.
     */
    function upgradeClient(UpgradeClientMsg calldata upgradeClientMsg) external onlyOwner {
        require(isClientCreated, 'Client not created');
        latestConsensusState = upgradeClientMsg.consensusState;
    }

    /**
     * @notice Verify the given proof data
     * @dev This function currently only checks if the proof length is non-zero
     * @param proof The proof data to be verified
     * @return A boolean value indicating if the proof is valid
     */
    function verify(Proof calldata proof) internal pure returns (bool) {
        // TODO: replace with real merkle verification logic
        if (proof.proof.length == 0) {
            return false;
        }
        return true;
    }

    //
    // IBC Channel methods
    //

    function concatStrings(string memory str1, string memory str2) private pure returns (bytes memory) {
        return abi.encodePacked(str1, str2);
    }

    /**
     * This func is called by a 'relayer' on behalf of a dApp. The dApp should be implements IbcReceiver.
     * The dApp should implement the onOpenIbcChannel method to handle one of the first two channel handshake methods,
     * ie. ChanOpenInit or ChanOpenTry.
     * If callback succeeds, the dApp should return the selected version, and an emitted event will be relayed to the
     * IBC/VIBC hub chain.
     */
    function openIbcChannel(
        IbcReceiver portAddress,
        string calldata version,
        ChannelOrder ordering,
        string[] calldata connectionHops,
        string calldata counterpartyPortId,
        string calldata counterpartyChannelId,
        string calldata counterpartyVersion
    ) external {
        string memory selectedVersion = portAddress.onOpenIbcChannel(
            version,
            ordering,
            connectionHops,
            counterpartyPortId,
            counterpartyChannelId,
            counterpartyVersion
        );

        emit OpenIbcChannel(
            address(portAddress),
            selectedVersion,
            ordering,
            connectionHops,
            counterpartyPortId,
            counterpartyChannelId
        );
    }

    /**
     * This func is called by a 'relayer' after the IBC/VIBC hub chain has processed the onOpenIbcChannel event.
     * The dApp should implement the onConnectIbcChannel method to handle the last two channel handshake methods, ie.
     * ChanOpenAck or ChanOpenConfirm.
     */
    function connectIbcChannel(
        IbcReceiver portAddress,
        bytes32 channelId,
        string[] calldata connectionHops,
        ChannelOrder ordering,
        string calldata counterpartyPortId,
        bytes32 counterpartyChannelId,
        bytes32 counterpartyVersion,
        Proof calldata proof
    ) external {
        require(
            verifier.verifyMembership(
                latestConsensusState,
                proof,
                'channel/path/to/be/added/here',
                bytes('expected channel bytes constructed from params. Channel.State = {Ack_Pending, Confirm_Pending}')
            ),
            'Fail to prove channel state'
        );

        portAddress.onConnectIbcChannel(channelId, counterpartyChannelId, counterpartyVersion);

        // Register port and channel mapping
        // TODO: check duplicated channel registration?
        portChannelMap[address(portAddress)][channelId] = Channel(
            counterpartyVersion,
            ordering,
            connectionHops,
            counterpartyPortId,
            counterpartyChannelId
        );
        emit ConnectIbcChannel(
            address(portAddress),
            channelId,
            counterpartyPortId,
            counterpartyChannelId,
            connectionHops
        );
    }

    /**
     * @notice Get the IBC channel with the specified port and channel ID
     * @param portAddress EVM address of the IBC port
     * @param channelId IBC channel ID from the port perspective
     * @return A channel struct is always returned. If it doesn't exists, the channel struct is populated with default
       values per EVM.
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
        require(channel.counterpartyChannelId != bytes32(0), 'Channel not owned by msg.sender');
        IbcReceiver reciever = IbcReceiver(msg.sender);
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
        require(
            verifier.verifyMembership(
                latestConsensusState,
                proof,
                'channel/path/to/be/added/here',
                bytes('expected channel bytes constructed from params. Channel.State = {Closed(_Pending?)}')
            ),
            'Fail to prove channel state'
        );
        // ensure port owns channel
        Channel memory channel = portChannelMap[portAddress][channelId];
        require(channel.counterpartyChannelId != bytes32(0), 'Channel not owned by portAddress');

        // confirm with dApp by calling its callback
        IbcReceiver reciever = IbcReceiver(portAddress);
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
     * @param fee The fee serves as the packet incentive for relayers. It's escrowed on the running chain and will be
       claimed by relayer later once the packet is delivered and ack'ed.
     */
    function sendPacket(
        bytes32 channelId,
        bytes calldata packet,
        uint64 timeoutTimestamp,
        uint256 fee
    ) external payable {
        // ensure port owns channel
        Channel memory channel = portChannelMap[msg.sender][channelId];
        require(channel.counterpartyChannelId != bytes32(0), 'Channel not owned by sender');
        // escrow packet fee
        // ignore returned data from `call`
        // (bool sent, bytes memory _data) = escrow.call{value: fee}('');
        (bool sent, ) = escrow.call{value: fee}('');
        require(sent, 'Failed to escrow packet fee');
        // packet sequence
        uint64 sequence = nextSendPacketSequence[msg.sender][channelId] + 1;
        nextSendPacketSequence[msg.sender][channelId] = sequence;
        // packet commitment
        sendPacketCommitment[msg.sender][channelId][sequence] = true;

        emit SendPacket(msg.sender, channelId, packet, sequence, timeoutTimestamp, fee);
    }

    /**
     * @notice Handle the acknowledgement of an IBC packet by the counterparty
     * @dev Verifies the given proof and calls the `onAcknowledgementPacket` function on the given `receiver` contract,
       ie. the IBC dApp.
       Prerequisite: the original packet is committed and not ack'ed or timed out yet.
     * @param receiver The IbcReceiver contract that should handle the packet acknowledgement event
     * If the address doesn't satisfy the interface, the transaction will be reverted.
     * @param packet The IbcPacket data for the acknowledged packet
     * @param ackPacket The acknowledgement receipt for the packet
     * @param proof The membership proof to verify the packet acknowledgement committed on Polymer chain
     */
    function acknowledgement(
        IbcReceiver receiver,
        IbcPacket calldata packet,
        AckPacket calldata ackPacket,
        Proof calldata proof
    ) external {
        // verify `receiver` is the original packet sender
        require(portIdAddressMatch(address(receiver), packet.src.portId), 'Receiver is not the original packet sender');
        // prove ack packet is on Polymer chain
        require(
            verifier.verifyMembership(
                latestConsensusState,
                proof,
                'ack/packet/path',
                'expected ack receipt hash on Polymer chain'
            ),
            'Fail to prove ack'
        );
        // verify packet has been committed and not yet ack'ed or timed out
        bool hasCommitment = sendPacketCommitment[address(receiver)][packet.src.channelId][packet.sequence];
        require(hasCommitment, 'Packet commitment not found');

        receiver.onAcknowledgementPacket(packet);

        // delete packet commitment to avoid double ack
        delete sendPacketCommitment[address(receiver)][packet.src.channelId][packet.sequence];

        emit Acknowledgement(address(receiver), packet.src.channelId, ackPacket, packet.sequence);
    }

    /**
     * @notice Timeout of an IBC packet
     * @dev Verifies the given proof and calls the `onTimeoutPacket` function on the given `receiver` contract, ie. the IBC-dApp.
     * Prerequisite: the original packet is committed and not ack'ed or timed out yet.
     * @param receiver The IbcReceiver contract that should handle the packet timeout event
     * If the address doesn't satisfy the interface, the transaction will be reverted.
     * @param packet The IbcPacket data for the timed-out packet
     * @param proof The non-membership proof data needed to verify the packet timeout
     */
    function timeout(IbcReceiver receiver, IbcPacket calldata packet, Proof calldata proof) external {
        // verify `receiver` is the original packet sender
        require(portIdAddressMatch(address(receiver), packet.src.portId), 'Receiver is not the original packet sender');
        // prove absence of packet receipt on Polymer chain
        require(
            verifier.verifyNonMembership(latestConsensusState, proof, 'packet/receipt/path'),
            'Fail to prove timeout'
        );
        // verify packet has been committed and not yet ack'ed or timed out
        bool hasCommitment = sendPacketCommitment[address(receiver)][packet.src.channelId][packet.sequence];
        require(hasCommitment, 'Packet commitment not found');

        receiver.onTimeoutPacket(packet);

        // delete packet commitment to avoid double timeout
        delete sendPacketCommitment[address(receiver)][packet.src.channelId][packet.sequence];

        emit Timeout(address(receiver), packet.src.channelId, packet.sequence);
    }

    /**
     * @notice Receive an IBC packet and then pass it to the IBC-dApp for processing if verification succeeds.
     * @dev Verifies the given proof and calls the `onRecvPacket` function on the given `receiver` contract
     * @param receiver The IbcReceiver contract that should handle the packet receipt event
     * If the address doesn't satisfy the interface, the transaction will be reverted.
     * @param packet The IbcPacket data for the received packet
     * @param proof The proof data needed to verify the packet receipt
     * @dev Emit an `RecvPacket` event with the details of the received packet;
     * Also emit a WriteAckPacket event, which can be relayed to Polymer chain by relayers
     */
    function recvPacket(IbcReceiver receiver, IbcPacket calldata packet, Proof calldata proof) external {
        // verify `receiver` is the intended packet destination
        require(
            portIdAddressMatch(address(receiver), packet.dest.portId),
            'Receiver is not the intended packet destination'
        );
        // prove packet is received on Polymer chain
        require(
            verifier.verifyMembership(
                latestConsensusState,
                proof,
                'packet/commitment/path',
                'expected virtual packet commitment hash on Polymer chain'
            ),
            'Fail to prove packet commitment'
        );
        // verify packet has not been received yet
        bool hasReceipt = recvPacketReceipt[address(receiver)][packet.dest.channelId][packet.sequence];
        require(!hasReceipt, 'Packet receipt already exists');

        // enforce recv'ed packet sequences always increment by 1 for ordered channels
        Channel memory channel = portChannelMap[address(receiver)][packet.dest.channelId];
        if (channel.ordering == ChannelOrder.ORDERED) {
            require(
                packet.sequence == nextRecvPacketSequence[address(receiver)][packet.dest.channelId],
                'Unexpected packet sequence'
            );
            nextRecvPacketSequence[address(receiver)][packet.dest.channelId] = packet.sequence + 1;
        }

        AckPacket memory ack = receiver.onRecvPacket(packet);
        bool hasAckPacketCommitment = ackPacketCommitment[address(receiver)][packet.dest.channelId][packet.sequence];
        // check is not necessary for sync-acks
        require(!hasAckPacketCommitment, 'Ack packet commitment already exists');
        ackPacketCommitment[address(receiver)][packet.dest.channelId][packet.sequence] = true;

        emit RecvPacket(
            address(receiver),
            packet.dest.channelId,
            packet.src.portId,
            packet.src.channelId,
            packet.sequence
        );
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

    /**
     * Generate a timeout packet for the given packet
     */
    function writeTimeoutPacket(address receiver, IbcPacket calldata packet) external {
        // verify `receiver` is the original packet sender
        require(portIdAddressMatch(receiver, packet.src.portId), 'Receiver is not the intended packet destination');
        // verify packet does not have a receipt
        bool hasReceipt = recvPacketReceipt[receiver][packet.dest.channelId][packet.sequence];
        require(!hasReceipt, 'Packet receipt already exists');
        // verify packet has timed out; zero-value in packet.timeout means no timeout set
        require(
            (packet.timeout.timestamp == 0 || block.timestamp >= packet.timeout.timestamp) &&
                (packet.timeout.blockHeight == 0 || block.number >= packet.timeout.blockHeight),
            'Packet not timed out yet'
        );

        emit WriteTimeoutPacket(receiver, packet.dest.channelId, packet.sequence);
    }
}
