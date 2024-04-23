// SPDX-License-Identifier: Apache-2.0
/*
 * Copyright 2024, Polymer Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity ^0.8.9;

import {UniversalPacket, AckPacket} from "../libs/Ibc.sol";
import {IbcUtils} from "../libs/IbcUtils.sol";
import {IbcMwUser, IbcUniversalPacketReceiver, IbcUniversalPacketSender} from "../interfaces/IbcMiddleware.sol";

/**
 * @title Earth
 * @notice Earth is a simple IBC receiver contract that receives packets and sends acks.
 * @dev This contract is used for only testing IBC functionality and as an example for dapp developers on how to
 * integrate with the vibc protocol.
 */
contract Earth is IbcMwUser, IbcUniversalPacketReceiver {
    struct UcPacketWithChannel {
        bytes32 channelId;
        UniversalPacket packet;
    }

    struct UcAckWithChannel {
        bytes32 channelId;
        UniversalPacket packet;
        AckPacket ack;
    }

    // received packet as chain B
    UcPacketWithChannel[] public recvedPackets;
    // received ack packet as chain A
    UcAckWithChannel[] public ackPackets;
    // received timeout packet as chain A
    UcPacketWithChannel[] public timeoutPackets;

    constructor(address _middleware) IbcMwUser(_middleware) {}

    function greet(address destPortAddr, bytes32 channelId, bytes calldata message, uint64 timeoutTimestamp) external {
        IbcUniversalPacketSender(mw).sendUniversalPacket(
            channelId, IbcUtils.toBytes32(destPortAddr), message, timeoutTimestamp
        );
    }

    function onRecvUniversalPacket(bytes32 channelId, UniversalPacket calldata packet)
        external
        onlyIbcMw
        returns (AckPacket memory ackPacket)
    {
        recvedPackets.push(UcPacketWithChannel(channelId, packet));
        return this.generateAckPacket(channelId, IbcUtils.toAddress(packet.srcPortAddr), packet.appData);
    }

    function onUniversalAcknowledgement(bytes32 channelId, UniversalPacket memory packet, AckPacket calldata ack)
        external
        onlyIbcMw
    {
        // verify packet's destPortAddr is the ack's first encoded address. See generateAckPacket())
        if (ack.data.length < 20) revert ackDataTooShort();
        address ackSender = address(bytes20(ack.data[0:20]));
        if (IbcUtils.toAddress(packet.destPortAddr) != ackSender) revert ackAddressMismatch();
        ackPackets.push(UcAckWithChannel(channelId, packet, ack));
    }

    function onTimeoutUniversalPacket(bytes32 channelId, UniversalPacket calldata packet) external onlyIbcMw {
        timeoutPackets.push(UcPacketWithChannel(channelId, packet));
    }

    // For testing only; real dApps should implment their own ack logic
    function generateAckPacket(bytes32, address srcPortAddr, bytes calldata appData)
        external
        view
        returns (AckPacket memory ackPacket)
    {
        return AckPacket(true, abi.encodePacked(this, srcPortAddr, "ack-", appData));
    }
}
