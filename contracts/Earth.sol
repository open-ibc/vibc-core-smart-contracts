//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "./Ibc.sol";
import "./IbcReceiver.sol";
import "./IbcDispatcher.sol";

contract Earth is IbcReceiverBase, IbcPacketHandler {
    // received packet as chain B
    IbcPacket[] public recvedPackets;
    // received ack packet as chain A
    AckPacket[] public ackPackets;
    // received timeout packet as chain A
    IbcPacket[] public timeoutPackets;

    constructor(IbcDispatcher _dispatcher) IbcReceiverBase(_dispatcher) {}

    function greet(string calldata portId, bytes32 channelId, bytes calldata message, uint64 timeoutTimestamp)
        external
    {
        dispatcher.sendPacketOverUniversalChannel(portId, channelId, message, timeoutTimestamp);
    }

    function onRecvPacket(IbcPacket memory packet) external onlyIbcDispatcher returns (AckPacket memory ackPacket) {
        recvedPackets.push(packet);
        return AckPacket(true, abi.encodePacked("ack-", packet.data));
    }

    function onAcknowledgementPacket(IbcPacket calldata packet, AckPacket calldata ack) external onlyIbcDispatcher {
        ackPackets.push(ack);
    }

    function onTimeoutPacket(IbcPacket calldata packet) external onlyIbcDispatcher {
        timeoutPackets.push(packet);
    }
}
