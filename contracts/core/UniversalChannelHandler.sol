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

pragma solidity 0.8.15;

import {IbcDispatcher} from "../interfaces/IbcDispatcher.sol";
import {IbcUniversalChannelMW, IbcUniversalPacketReceiver} from "../interfaces/IbcMiddleware.sol";
import {IbcReceiverBaseUpgradeable} from "../implementation_templates/IbcReceiverUpgradeable.sol";
import {ChannelOrder, IbcPacket, AckPacket, UniversalPacket} from "../libs/Ibc.sol";
import {IbcUtils} from "../libs/IbcUtils.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {FeeSender} from "../implementation_templates/FeeSender.sol";

/**
 * @title Universal Channel Handler
 * @author Polymer Labs
 * @notice Implements universal channels for virtual IBC. Universal channels prevent dapps from needing to do a 4-step
 * channel handshake to establish a channel.
 * @dev This contract can integrate directly with dapps, or a middleware stack for packet routing.
 */
contract UniversalChannelHandler is IbcReceiverBaseUpgradeable, FeeSender, UUPSUpgradeable, IbcUniversalChannelMW {
    string public constant VERSION = "1.0";
    uint256 public constant MW_ID = 1;

    event UCHPacketSent(address source, bytes32 destination);

    constructor() {
        _disableInitializers();
    }

    function initialize(IbcDispatcher _dispatcher) public initializer {
        __IbcReceiverBase_init(_dispatcher);
    }

    /**
     * @dev Close a universal channel.
     * Cannot send or receive packets after the channel is closed.
     * @param channelId The channel id of the channel to be closed.
     */
    function closeChannel(bytes32 channelId) external onlyOwner {
        dispatcher.channelCloseInit(channelId);
    }

    function onChanCloseInit(bytes32 channelId, string calldata, bytes32) external onlyIbcDispatcher {}

    function onChanCloseConfirm(bytes32 channelId, string calldata, bytes32) external onlyIbcDispatcher {}

    function openChannel(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortIdentifier
    ) external onlyOwner {
        dispatcher.channelOpenInit(version, ordering, feeEnabled, connectionHops, counterpartyPortIdentifier);
    }

    /**
     * @notice Sends a universal packet over an IBC channel, without sending relaying fees to the FeeVault
     * @param channelId The channel ID through which the packet is sent on the dispatcher
     * @param destPortAddr The destination port address
     * @param appData The packet data to be sent
     * @param timeoutTimestamp of when the packet can timeout
     */
    function sendUniversalPacket(
        bytes32 channelId,
        bytes32 destPortAddr,
        bytes calldata appData,
        uint64 timeoutTimestamp
    ) external returns (uint64 sequence) {
        bytes memory packetData = IbcUtils.toUniversalPacketBytes(
            UniversalPacket(IbcUtils.toBytes32(msg.sender), MW_ID, destPortAddr, appData)
        );
        emit UCHPacketSent(msg.sender, destPortAddr);
        sequence = dispatcher.sendPacket(channelId, packetData, timeoutTimestamp);
    }

    /**
     * @notice Sends a universal packet over an IBC channel, and deposits relaying fees to a FeeVault within the same
     * tx.
     * @param channelId The channel ID through which the packet is sent on the dispatcher
     * @param destPortAddr The destination port address
     * @param appData The packet data to be sent
     * @param timeoutTimestamp of when the packet can timeout
     * @param gasLimits An array containing two gas limit values:
     *                  - gasLimits[0] for `recvPacket` fees
     *                  - gasLimits[1] for `ackPacket` fees.
     * @param gasPrices An array containing two gas price values:
     *                  - gasPrices[0] for `recvPacket` fees, for the dest chain
     *                  - gasPrices[1] for `ackPacket` fees, for the src chain
     * @notice The total fees sent in the msg.value should be equal to the total gasLimits[0] * gasPrices[0] +
     * gasLimits[1] * gasPrices[1]. The transaction will revert if a higher or lower value is sent
     * @notice if you are relaying your own transactions, you should not call this method, and instead call
     * sendUniversalPacket
     * @notice Use the Polymer fee estimation api to get the required fees to ensure that enough fees are sent.
     */
    function sendUniversalPacketWithFee(
        bytes32 channelId,
        bytes32 destPortAddr,
        bytes calldata appData,
        uint64 timeoutTimestamp,
        uint256[2] calldata gasLimits,
        uint256[2] calldata gasPrices
    ) external payable returns (uint64 sequence) {
        // Cache dispatcher for gas savings
        IbcDispatcher _dispatcher = dispatcher;

        bytes memory packetData = IbcUtils.toUniversalPacketBytes(
            UniversalPacket(IbcUtils.toBytes32(msg.sender), MW_ID, destPortAddr, appData)
        );
        emit UCHPacketSent(msg.sender, destPortAddr);
        sequence = _dispatcher.sendPacket(channelId, packetData, timeoutTimestamp);
        _depositSendPacketFee(dispatcher, channelId, sequence, gasLimits, gasPrices);
    }

    /**
     * @notice Handles the reception of an IBC packet from the counterparty
     * @param packet The received IBC packet
     * @return ackPacket The packet acknowledgement
     */
    function onRecvPacket(IbcPacket calldata packet)
        external
        override
        onlyIbcDispatcher
        returns (AckPacket memory ackPacket)
    {
        UniversalPacket memory ucPacket = IbcUtils.fromUniversalPacketBytes(packet.data);
        // no other middleware stack registered for this packet. Deliver packet to dApp directly.
        return IbcUniversalPacketReceiver(IbcUtils.toAddress(ucPacket.destPortAddr)).onRecvUniversalPacket(
            packet.dest.channelId, ucPacket
        );
    }

    /**
     * @notice Handles acknowledging the reception of an acknowledgement packet by the counterparty
     * @param packet The IBC packet to be acknowledged
     * @param ack The packet acknowledgement
     */
    function onAcknowledgementPacket(IbcPacket calldata packet, AckPacket calldata ack)
        external
        override
        onlyIbcDispatcher
    {
        UniversalPacket memory ucPacket = IbcUtils.fromUniversalPacketBytes(packet.data);
        // no other middleware stack registered for this packet. Deliver ack to dApp directly.
        IbcUniversalPacketReceiver(IbcUtils.toAddress(ucPacket.srcPortAddr)).onUniversalAcknowledgement(
            packet.src.channelId, ucPacket, ack
        );
    }

    /**
     * @notice Handles the timeout event for an IBC packet
     * @param packet The IBC packet that has timed out
     */
    function onTimeoutPacket(IbcPacket calldata packet) external override onlyIbcDispatcher {
        UniversalPacket memory ucPacketData = IbcUtils.fromUniversalPacketBytes(packet.data);
        // no other middleware stack registered for this packet. Deliver timeout to dApp directly.
        IbcUniversalPacketReceiver(IbcUtils.toAddress(ucPacketData.srcPortAddr)).onTimeoutUniversalPacket(
            packet.src.channelId, ucPacketData
        );
    }

    function onChanOpenConfirm(bytes32 channelId) external onlyIbcDispatcher {}

    function setDispatcher(IbcDispatcher _dispatcher) external onlyOwner {
        dispatcher = _dispatcher;
    }

    // solhint-disable-next-line ordering
    function onChanOpenTry(
        ChannelOrder,
        string[] memory,
        bytes32 channelId,
        string memory,
        bytes32,
        string calldata counterpartyVersion
    ) external view onlyIbcDispatcher returns (string memory selectedVersion) {
        return _connectChannel(channelId, counterpartyVersion);
    }

    // IBC callback functions
    /**
     * @notice Handles the acknowledgment of channel opening.
     * This function is accessible only by the IBC dispatcher.
     * @param channelId The channel ID of the opened channel
     * @param counterpartyVersion The version string provided by the counterparty
     */
    function onChanOpenAck(bytes32 channelId, bytes32, string calldata counterpartyVersion)
        external
        view
        onlyIbcDispatcher
    {
        _connectChannel(channelId, counterpartyVersion);
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /**
     * @dev Internal function to connect a channel only if the version matches what is expected.
     * @param version The version string provided by the counterparty
     */
    function _connectChannel(bytes32, string calldata version) internal pure returns (string memory checkedVersion) {
        if (keccak256(abi.encodePacked(version)) != keccak256(abi.encodePacked(VERSION))) {
            revert UnsupportedVersion();
        }
        return version;
    }
}
