//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "./Ibc.sol";
import "./IbcReceiver.sol";
import "./IbcDispatcher.sol";
import "./IbcMiddleware.sol";

contract Earth is IbcMwReceiverBase, IbcPacketReceiver {
    // received packet as chain B
    IbcPacket[] public recvedPackets;
    // received ack packet as chain A
    AckPacket[] public ackPackets;
    // received timeout packet as chain A
    IbcPacket[] public timeoutPackets;

    constructor(IbcMiddleware _middleware) IbcMwReceiverBase(_middleware) {}

    function greet(address destPortAddr, bytes32 channelId, bytes calldata message, uint64 timeoutTimestamp) external {
        mw.sendUniversalPacket(channelId, destPortAddr, message, timeoutTimestamp);
    }

    function onRecvPacket(IbcPacket memory packet) external onlyIbcMw returns (AckPacket memory ackPacket) {
        recvedPackets.push(packet);
        return AckPacket(true, abi.encodePacked("ack-", packet.data));
    }

    function onAcknowledgementPacket(IbcPacket calldata packet, AckPacket calldata ack) external onlyIbcMw {
        ackPackets.push(ack);
    }

    function onTimeoutPacket(IbcPacket calldata packet) external onlyIbcMw {
        timeoutPackets.push(packet);
    }
}
