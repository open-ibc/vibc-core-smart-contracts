//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./IbcDispatcher.sol";
import "./IbcMiddleware.sol";
import "./Ibc.sol";

contract UniversalChannelHandler is IbcReceiverBase, IbcUniversalChannelMW {
    constructor(IbcDispatcher _dispatcher) IbcReceiverBase(_dispatcher) {}

    bytes32[] public connectedChannels;
    string constant VERSION = "1.0";
    uint256 public constant MW_ID = 1;

    /**
     * @dev Close a universal channel.
     * Cannot send or receive packets after the channel is closed.
     * @param channelId The channel id of the channel to be closed.
     */
    function closeChannel(bytes32 channelId) external onlyOwner {
        dispatcher.closeIbcChannel(channelId);
    }

    // IBC callback functions

    function onOpenIbcChannel(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId,
        bytes32 counterpartyChannelId,
        string calldata counterpartyVersion
    ) external onlyIbcDispatcher returns (string memory selectedVersion) {
        if (counterpartyChannelId == bytes32(0)) {
            // ChanOpenInit
            require(keccak256(abi.encodePacked(version)) == keccak256(abi.encodePacked(VERSION)), "Unsupported version");
        } else {
            // ChanOpenTry
            require(
                keccak256(abi.encodePacked(counterpartyVersion)) == keccak256(abi.encodePacked(VERSION)),
                "Unsupported version"
            );
        }
        return VERSION;
    }

    function onConnectIbcChannel(bytes32 channelId, bytes32 counterpartyChannelId, string calldata counterpartyVersion)
        external
        onlyIbcDispatcher
    {
        require(
            keccak256(abi.encodePacked(counterpartyVersion)) == keccak256(abi.encodePacked(VERSION)),
            "Unsupported version"
        );
        connectedChannels.push(channelId);
    }

    function onCloseIbcChannel(bytes32 channelId, string calldata counterpartyPortId, bytes32 counterpartyChannelId)
        external
        onlyIbcDispatcher
    {
        // logic to determin if the channel should be closed
        bool channelFound = false;
        for (uint256 i = 0; i < connectedChannels.length; i++) {
            if (connectedChannels[i] == channelId) {
                delete connectedChannels[i];
                channelFound = true;
                break;
            }
        }
        require(channelFound, "Channel not found");
    }

    function sendUniversalPacket(
        bytes32 channelId,
        address destPortAddr,
        bytes calldata appData,
        uint64 timeoutTimestamp
    ) external {
        bytes memory packetData =
            Ibc.toUniversalPacketDataBytes(UniversalPacketData(msg.sender, MW_ID, destPortAddr, appData));
        dispatcher.sendPacket(channelId, packetData, timeoutTimestamp);
    }

    // called by another IBC middleware; pack packet and send over to Dispatcher
    function sendMWPacket(
        bytes32 channelId,
        // original source address of the packet
        address srcPortAddr,
        address destPortAddr,
        // source middleware ids bit AND
        uint256 srcMwIds,
        bytes calldata appData,
        uint64 timeoutTimestamp
    ) external {
        bytes memory packetData =
            Ibc.toUniversalPacketDataBytes(UniversalPacketData(srcPortAddr, srcMwIds & MW_ID, destPortAddr, appData));
        dispatcher.sendPacket(channelId, packetData, timeoutTimestamp);
    }

    function onRecvPacket(IbcPacket memory packet) external override returns (AckPacket memory ackPacket) {
        UniversalPacketData memory ucPacketData = Ibc.fromUniversalPacketDataBytes(packet.data);
    }

    function onAcknowledgementPacket(IbcPacket calldata packet, AckPacket calldata ack) external override {}

    function onTimeoutPacket(IbcPacket calldata packet) external override {}
}
