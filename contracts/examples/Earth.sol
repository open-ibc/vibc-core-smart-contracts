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
import {IbcUniversalPacketReceiverBase, IbcUniversalPacketSender} from "../interfaces/IbcMiddleware.sol";
import {IUniversalChannelHandler} from "../interfaces/IUniversalChannelHandler.sol";

/**
 * @title Earth
 * @notice Earth is a simple IBC receiver contract that receives packets and sends acks.
 * @dev This contract is used for only testing IBC functionality and as an example for dapp developers on how to
 * integrate with the vibc protocol.
 */
contract Earth is IbcUniversalPacketReceiverBase {
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

    constructor(address _middleware) IbcUniversalPacketReceiverBase(_middleware) {}

    /**
     * @notice Send a packet to a destination chain. without a fee
     * @notice this is useful for self-relaying apckets which don't rely on polymer to fund.
     * @param destPortAddr The destination chain's port address.
     * @param channelId The channel id to send the packet on.
     * @param message The message to send.
     * @param timeoutTimestamp The timeout timestamp for the packet.
     */
    function greet(address destPortAddr, bytes32 channelId, bytes calldata message, uint64 timeoutTimestamp) external {
        IbcUniversalPacketSender(uch).sendUniversalPacket(
            channelId, IbcUtils.toBytes32(destPortAddr), message, timeoutTimestamp
        );
    }

    /**
     * @notice Sends a universal packet with fee.
     * @param destPortAddr The destination port address.
     * @param channelId The channel ID.
     * @param message The message to send.
     * @param timeoutTimestamp The timeout timestamp.
     * @param gasLimits The gas limits for the packet and the ack.
     * @param gasPrices The gas prices for the packet and the ack.
     * @return sequence The sequence number of the packet.
     * @param gasLimits An array containing two gas limit values:
     *                  - gasLimits[0] for `recvPacket` fees
     *                  - gasLimits[1] for `ackPacket` fees.
     * @param gasPrices An array containing two gas price values:
     *                  - gasPrices[0] for `recvPacket` fees, for the dest chain
     *                  - gasPrices[1] for `ackPacket` fees, for the src chain
     * @notice If you are relaying your own packets, you should not call this method, and instead call greet.
     * @notice The total fees sent in the msg.value should be equal to the total gasLimits[0] * gasPrices[0] +
     * @notice Use the Polymer fee estimation api to get the required fees to ensure that enough fees are sent.
     * gasLimits[1] * gasPrices[1]. The transaction will revert if a higher or lower value is sent
     */
    function greetWithFee(
        address destPortAddr,
        bytes32 channelId,
        bytes calldata message,
        uint64 timeoutTimestamp,
        uint256[2] memory gasLimits,
        uint256[2] memory gasPrices
    ) external payable returns (uint64 sequence) {
        return IUniversalChannelHandler(uch).sendUniversalPacketWithFee{value: msg.value}(
            channelId, IbcUtils.toBytes32(destPortAddr), message, timeoutTimestamp, gasLimits, gasPrices
        );
    }

    /**
     * @notice Handles the recv callback of a universal packet.
     * @param channelId The channel ID.
     * @param packet The universal packet.
     * @return ackPacket The acknowledgement packet.
     * @return skipAck Whether to skip the writeAck event.
     * @dev It's recommended to always validate the authorized channel of any packet or channel using the
     * onlyAuthorizedChannel modifier.
     */
    function onRecvUniversalPacket(bytes32 channelId, UniversalPacket calldata packet)
        external
        onlyUCH
        onlyAuthorizedChannel(channelId)
        returns (AckPacket memory ackPacket, bool skipAck)
    {
        recvedPackets.push(UcPacketWithChannel(channelId, packet));
        return (this.generateAckPacket(channelId, IbcUtils.toAddress(packet.srcPortAddr), packet.appData), false);
    }

    /**
     * @notice Handles the acknowledgement of a universal packet.
     * @param channelId The channel ID.
     * @param packet The universal packet.
     * @param ack The acknowledgement packet.
     * @dev It's recommended to always validate the authorized channel of any packet or channel using the
     * onlyAuthorizedChannel modifier.
     */
    function onUniversalAcknowledgement(bytes32 channelId, UniversalPacket memory packet, AckPacket calldata ack)
        external
        onlyUCH
        onlyAuthorizedChannel(channelId)
    {
        // verify packet's destPortAddr is the ack's first encoded address. See generateAckPacket())
        if (ack.data.length < 20) revert ackDataTooShort();
        address ackSender = address(bytes20(ack.data[0:20]));
        if (IbcUtils.toAddress(packet.destPortAddr) != ackSender) revert ackAddressMismatch();
        ackPackets.push(UcAckWithChannel(channelId, packet, ack));
    }

    /**
     * @notice Handles the timeout of a universal packet. Timeouts are currently unimplemented, so this handler is
     * unused.
     * @param channelId The channel ID.
     * @param packet The universal packet.
     * @dev It's recommended to always validate the authorized channel of any packet or channel using the
     * onlyAuthorizedChannel modifier.
     */
    function onTimeoutUniversalPacket(bytes32 channelId, UniversalPacket calldata packet)
        external
        onlyUCH
        onlyAuthorizedChannel(channelId)
    {
        timeoutPackets.push(UcPacketWithChannel(channelId, packet));
    }

    /**
     * @notice Authorize a channel that can receive/ack packets to this contract. If channels are not validated, this
     * could allow any arbitrary dapps to trigger callbacks on this dapp.
     * @param channelId The channel id to authorize; should be packet.dest.channelId for recv packets &
     * packet.src.channelId for ack packets.
     */
    function authorizeChannel(bytes32 channelId) external onlyOwner {
        _authorizeChannel(channelId);
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
