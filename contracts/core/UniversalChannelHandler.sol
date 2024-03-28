//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IbcDispatcher} from "../interfaces/IbcDispatcher.sol";
import {
    IbcMiddleware,
    IbcUniversalChannelMW,
    IbcUniversalPacketReceiver,
    IbcMwPacketReceiver,
    IbcMwEventsEmitter
} from "../interfaces/IbcMiddleware.sol";
import {IbcReceiver} from "../interfaces/IbcReceiver.sol";
import {IbcReceiverBaseUpgradeable} from "../interfaces/IbcReceiverUpgradeable.sol";
import {ChannelOrder, ChannelEnd, IbcPacket, AckPacket, UniversalPacket, IbcUtils} from "../libs/Ibc.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";

/**
 * @title Universal Channel Handler
 * @author Polymer Labs
 * @notice Implements universal channels for virtual IBC. Universal channels prevent dapps from needing to do a 4-step
 * channel handshake to establish a channel.
 * @dev This contract can integrate directly with dapps, or a middleware stack for packet routing.
 */
contract UniversalChannelHandler is IbcReceiverBaseUpgradeable, UUPSUpgradeable, IbcUniversalChannelMW {
    uint256[49] private __gap;

    bytes32 private _UNUSED; // Storage placeholder to ensure upgrade from this version is backwards compatible

    string public constant VERSION = "1.0";
    uint256 public constant MW_ID = 1;

    // Key: middleware bitmap, Value: middleware address from receiver(chain B)'s perspective
    mapping(uint256 => address[]) public mwStackAddrs;

    event UCHPacketSent(address source, bytes32 destination);

    constructor() {
        _disableInitializers();
    }

    function initialize(IbcDispatcher _dispatcher) public initializer {
        __IbcReceiverBase_init(_dispatcher);
    }

    function onChanCloseInit(bytes32 channelId, string calldata, bytes32) external onlyIbcDispatcher {}

    function onChanCloseConfirm(bytes32 channelId, string calldata, bytes32) external onlyIbcDispatcher {}

    function openChannel(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortIdentifier
    ) external onlyOwner {
        dispatcher.channelOpenInit(version, ordering, feeEnabled, connectionHops, counterpartyPortIdentifier);
    }

    /**
     * @notice Sends a universal packet over an IBC channel
     * @param channelId The channel ID through which the packet is sent on the dispatcher
     * @param destPortAddr The destination port address
     * @param appData The packet data to be sent
     * @param timeoutTimestamp of when the packet can timeout
     */
    function sendUniversalPacket(
        bytes32 channelId,
        bytes32 destPortAddr,
        bytes calldata appData,
        uint64 timeoutTimestamp
    ) external {
        bytes memory packetData = IbcUtils.toUniversalPacketBytes(
            UniversalPacket(IbcUtils.toBytes32(msg.sender), MW_ID, destPortAddr, appData)
        );
        emit UCHPacketSent(msg.sender, destPortAddr);
        dispatcher.sendPacket(channelId, packetData, timeoutTimestamp);
    }

    /**
     * @notice Sends a middleware packet over an IBC channel. This is intended to be called by another middleware
     * contract, rather than an end Dapp itself.
     * @param channelId The channel ID through which the packet is sent on the dispatcher
     * @param destPortAddr The destination port address
     * @param srcMwIds The mwId bitmap of the middleware stack
     * @param appData The packet data to be sent
     * @param timeoutTimestamp of when the packet can timeout
     */
    function sendMWPacket(
        bytes32 channelId,
        // original source address of the packet
        bytes32 srcPortAddr,
        bytes32 destPortAddr,
        // source middleware ids bit AND
        uint256 srcMwIds,
        bytes calldata appData,
        uint64 timeoutTimestamp
    ) external {
        bytes memory packetData =
            IbcUtils.toUniversalPacketBytes(UniversalPacket(srcPortAddr, srcMwIds | MW_ID, destPortAddr, appData));
        dispatcher.sendPacket(channelId, packetData, timeoutTimestamp);
    }

    /**
     * @notice Handles the reception of an IBC packet from the counterparty
     * @param packet The received IBC packet
     * @return ackPacket The packet acknowledgement
     */
    function onRecvPacket(IbcPacket calldata packet)
        external
        override
        onlyIbcDispatcher
        returns (AckPacket memory ackPacket)
    {
        UniversalPacket memory ucPacket = IbcUtils.fromUniversalPacketBytes(packet.data);
        address[] storage mwAddrs = mwStackAddrs[ucPacket.mwBitmap];
        if (mwAddrs.length == 0) {
            // no other middleware stack registered for this packet. Deliver packet to dApp directly.
            return IbcUniversalPacketReceiver(IbcUtils.toAddress(ucPacket.destPortAddr)).onRecvUniversalPacket(
                packet.dest.channelId, ucPacket
            );
        } else {
            // send packet to first MW in the stack
            return IbcMwPacketReceiver(mwAddrs[0]).onRecvMWPacket(packet.dest.channelId, ucPacket, 0, mwAddrs);
        }
    }

    /**
     * @notice Handles acknowledging the reception of an acknowledgement packet by the counterparty
     * @param packet The IBC packet to be acknowledged
     * @param ack The packet acknowledgement
     */
    function onAcknowledgementPacket(IbcPacket calldata packet, AckPacket calldata ack)
        external
        override
        onlyIbcDispatcher
    {
        UniversalPacket memory ucPacket = IbcUtils.fromUniversalPacketBytes(packet.data);
        address[] storage mwAddrs = mwStackAddrs[ucPacket.mwBitmap];
        if (mwAddrs.length == 0) {
            // no other middleware stack registered for this packet. Deliver ack to dApp directly.
            IbcUniversalPacketReceiver(IbcUtils.toAddress(ucPacket.srcPortAddr)).onUniversalAcknowledgement(
                packet.src.channelId, ucPacket, ack
            );
        } else {
            // send ack to last MW in the stack
            IbcMwPacketReceiver(mwAddrs[0]).onRecvMWAck(packet.src.channelId, ucPacket, 0, mwAddrs, ack);
        }
    }

    /**
     * @notice Handles the timeout event for an IBC packet
     * @param packet The IBC packet that has timed out
     */
    function onTimeoutPacket(IbcPacket calldata packet) external override onlyIbcDispatcher {
        UniversalPacket memory ucPacketData = IbcUtils.fromUniversalPacketBytes(packet.data);
        address[] storage mwAddrs = mwStackAddrs[ucPacketData.mwBitmap];
        if (mwAddrs.length == 0) {
            // no other middleware stack registered for this packet. Deliver timeout to dApp directly.
            IbcUniversalPacketReceiver(IbcUtils.toAddress(ucPacketData.srcPortAddr)).onTimeoutUniversalPacket(
                packet.src.channelId, ucPacketData
            );
        } else {
            // send timeout to last MW in the stack
            IbcMwPacketReceiver(mwAddrs[0]).onRecvMWTimeout(packet.src.channelId, ucPacketData, 0, mwAddrs);
        }
    }

    /**
     * @dev Register a middleware stack for universal packet routing.
     * This is a temporary solution for testing only.
     * Polymer chain will maintain a global registry of middleware stacks.
     * @param mwBitmap Bit OR of all MW IDs in the stack, excluding this MW's ID
     * @param mwAddrs addresses in the stack, from the perspective of the receiver (chain B)
     * MW closer to UniversalChannel MW has smaller index. MW stack must be in the same order on both chains.
     */
    function registerMwStack(uint256 mwBitmap, address[] calldata mwAddrs) external onlyOwner {
        if (mwBitmap == 0) revert MwBitmpaCannotBeZero();
        mwStackAddrs[mwBitmap] = mwAddrs;
    }

    /**
     * @notice Handles the initialization of channel opening.
     * @param version The new channel's version
     */
    function onChanOpenInit(ChannelOrder, string[] calldata, string calldata, string calldata version)
        external
        view
        onlyIbcDispatcher
        returns (string memory selectedVersion)
    {
        return _openChannel(version);
    }

    // solhint-disable-next-line ordering
    function onChanOpenTry(
        ChannelOrder,
        string[] memory,
        bytes32 channelId,
        string memory,
        bytes32,
        string calldata counterpartyVersion
    ) external onlyIbcDispatcher returns (string memory selectedVersion) {
        return _connectChannel(channelId, counterpartyVersion);
    }

    // IBC callback functions
    /**
     * @notice Handles the acknowledgment of channel opening.
     * This function is accessible only by the IBC dispatcher.
     * @param channelId The channel ID of the opened channel
     * @param counterpartyVersion The version string provided by the counterparty
     */
    function onChanOpenAck(bytes32 channelId, bytes32, string calldata counterpartyVersion)
        external
        onlyIbcDispatcher
    {
        _connectChannel(channelId, counterpartyVersion);
    }

    function onChanOpenConfirm(bytes32 channelId) external onlyIbcDispatcher {}

    function setDispatcher(IbcDispatcher _dispatcher) external onlyOwner {
        dispatcher = _dispatcher;
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /**
     * @dev Internal function to connect a channel only if the version matches what is expected.
     * @param channelId The channel ID of the channel to connect
     * @param version The version string provided by the counterparty
     */
    function _connectChannel(bytes32 channelId, string calldata version)
        internal
        returns (string memory checkedVersion)
    {
        if (keccak256(abi.encodePacked(version)) != keccak256(abi.encodePacked(VERSION))) {
            revert UnsupportedVersion();
        }
        return version;
    }

    /**
     * @dev Internal function to open a channel only if the version matches what is expected.
     * @param version The version string provided by the counterparty
     * @return selectedVersion The selected version string
     */
    function _openChannel(string calldata version) internal pure returns (string memory selectedVersion) {
        if (keccak256(abi.encodePacked(version)) != keccak256(abi.encodePacked(VERSION))) {
            revert UnsupportedVersion();
        }
        return VERSION;
    }
}
