//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./IbcDispatcher.sol";
import "./Ibc.sol";

contract UniversalChannelHandler is IbcReceiverBase, IbcReceiver {
    constructor(IbcDispatcher _dispatcher) IbcReceiverBase(_dispatcher) {}

    bytes32[] public connectedChannels;
    string constant VERSION = "1.0";

    /**
     * @dev Close a universal channel.
     * Cannot send or receive packets after the channel is closed.
     * @param channelId The channel id of the channel to be closed.
     */
    function closeChannel(bytes32 channelId) external onlyOwner {
        dispatcher.closeIbcChannel(channelId);
    }

    // IBC callback functions

    function onRecvPacket(IbcPacket calldata packet) external onlyIbcDispatcher returns (AckPacket memory ackPacket) {
        return AckPacket(true, abi.encodePacked('{ "account": "account", "reply": "got the message" }'));
    }

    function onAcknowledgementPacket(IbcPacket calldata packet, AckPacket calldata ack) external onlyIbcDispatcher {}

    function onTimeoutPacket(IbcPacket calldata packet) external onlyIbcDispatcher {}

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
}
