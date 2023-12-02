//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "./Ibc.sol";
import "./IbcReceiver.sol";
import "./IbcDispatcher.sol";

contract Earth is IbcReceiverBase {
    constructor(IbcDispatcher _dispatcher) IbcReceiverBase(_dispatcher) {}

    function greet(string calldata portId, bytes32 channelId, string calldata message) external payable {
        PacketFee memory fee;
        dispatcher.sendPacketOverUniversalChannel{value: Ibc.calcEscrowFee(fee)}(
            portId, channelId, bytes(message), 1 days, fee
        );
    }
}
