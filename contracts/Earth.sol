//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "./Ibc.sol";
import "./IbcReceiver.sol";
import "./IbcDispatcher.sol";
import "./IbcMiddleware.sol";

contract Earth is IbcMwReceiverBase, IbcUniversalPacketReceiver {
    struct UcPacket {
        bytes32 channelId;
        address srcPortId;
        bytes appData;
    }

    // received packet as chain B
    UcPacket[] public recvedPackets;
    // received ack packet as chain A
    AckPacket[] public ackPackets;
    // received timeout packet as chain A
    IbcPacket[] public timeoutPackets;

    constructor(IbcMiddleware _middleware) IbcMwReceiverBase(_middleware) {}

    function greet(address destPortAddr, bytes32 channelId, bytes calldata message, uint64 timeoutTimestamp) external {
        mw.sendUniversalPacket(channelId, destPortAddr, message, timeoutTimestamp);
    }

    // For testing only; real dApps should implment their own ack logic
    function generateAckPacket(bytes32 channelId, address srcPortAddr, bytes calldata appData)
        external
        view
        returns (AckPacket memory ackPacket)
    {
        return AckPacket(true, abi.encodePacked(this, srcPortAddr, "ack-", appData));
    }

    function onRecvUniversalPacket(bytes32 channelId, address srcPortAddr, bytes calldata appData)
        external
        onlyIbcMw
        returns (AckPacket memory ackPacket)
    {
        recvedPackets.push(UcPacket(channelId, srcPortAddr, appData));
        return this.generateAckPacket(channelId, srcPortAddr, appData);
    }

    function onAcknowledgementPacket(IbcPacket calldata packet, AckPacket calldata ack) external onlyIbcMw {
        ackPackets.push(ack);
    }

    function onTimeoutPacket(IbcPacket calldata packet) external onlyIbcMw {
        timeoutPackets.push(packet);
    }
}
