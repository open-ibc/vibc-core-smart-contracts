//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "./Ibc.sol";
import "./IbcReceiver.sol";
import "./IbcDispatcher.sol";

contract Earth is IbcReceiverBase {
    constructor(IbcDispatcher _dispatcher) IbcReceiverBase(_dispatcher) {}

    function greet(
        string calldata portId,
        bytes32 channelId,
        bytes calldata message,
        uint64 timeoutTimestamp,
        PacketFee calldata fee
    ) external payable {
        dispatcher.sendPacketOverUniversalChannel{value: Ibc.calcEscrowFee(fee)}(
            portId, channelId, message, timeoutTimestamp, fee
        );
    }
}
