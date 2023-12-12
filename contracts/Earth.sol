//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "./Ibc.sol";
import "./IbcReceiver.sol";
import "./IbcDispatcher.sol";
import "./IbcMiddleware.sol";

contract Earth is IbcMwUser, IbcUniversalPacketReceiver {
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
    UniversalPacket[] public timeoutPackets;

    constructor(address _middleware) IbcMwUser(_middleware) {}

    function greet(address destPortAddr, bytes32 channelId, bytes calldata message, uint64 timeoutTimestamp) external {
        IbcUniversalPacketSender(mw).sendUniversalPacket(channelId, destPortAddr, message, timeoutTimestamp);
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

    function onUniversalAcknowledgement(UniversalPacket memory packet, AckPacket calldata ack) external onlyIbcMw {
        // verify packet's destPortAddr is the ack's first encoded address. See generateAckPacket())
        require(ack.data.length >= 20, "ack data too short");
        address ackSender = address(bytes20(ack.data[0:20]));
        require(packet.destPortAddr == ackSender, "ack address mismatch");
        ackPackets.push(ack);
    }

    function onTimeoutUniversalPacket(UniversalPacket calldata packet) external onlyIbcMw {
        timeoutPackets.push(packet);
    }
}
