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
import {ChannelOrder} from "../libs/Ibc.sol";

/**
 */
/**
 * @title FeeSender
 * @notice This contract provides methods to deposit fees for sendPacket and channelOpenInit calls, so that they can be
 * relayed by polymer.
 * @notice Contracts don't need to inherit from this contract if they plan to relay their own packets and channel
 * handshake txs.
 * @notice Use the Polymer fee estimation api to get the required fees to ensure that enough fees are sent.
 */
abstract contract FeeSender {
    /**
     * @notice Deposits the send packet relayer fee for a given channel and sequence into the FeeVault.
     * @notice If you are relaying your own packets, you should not call this method.
     * @param dispatcher The dispatcher contract that the corresponding sendPacket will be called on.
     * @param channelId The channel id to deposit fees for that the packet is sent over.
     * @param sequence The sequence of the packet to deposit fees for. This is returned from the sendPacket call.
     * @param gasLimits An array containing two gas limit values:
     *                  - gasLimits[0] for `recvPacket` fees
     *                  - gasLimits[1] for `ackPacket` fees.
     * @param gasPrices An array containing two gas price values:
     *                  - gasPrices[0] for `recvPacket` fees, for the dest chain
     *                  - gasPrices[1] for `ackPacket` fees, for the src chain
     * @notice The total fees sent in the msg.value should be equal to the total gasLimits[0] * gasPrices[0] +
     * @notice Use the Polymer fee estimation api to get the required fees to ensure that enough fees are sent.
     * @dev Note: We have to have gasLimits and gasPrices as memory arrays. We cannot have them as calldata arrays
     * because solidity has weird behavior with using too much calldata in stacked calls
     * gasLimits[1] * gasPrices[1]. The transaction will revert if a higher or lower value is sent
     */
    function _depositSendPacketFee(
        IbcDispatcher dispatcher,
        bytes32 channelId,
        uint64 sequence,
        uint256[2] memory gasLimits,
        uint256[2] memory gasPrices
    ) internal {
        dispatcher.feeVault().depositSendPacketFee{value: msg.value}(channelId, sequence, gasLimits, gasPrices);
    }

    /**
     * @notice Deposits the open channel fee for a given channel and sequence into the FeeVault.
     * @notice If you are relaying your own channelHandshake transactions, you should not call this method.
     * @param dispatcher The dispatcher contract the channelOpenInit was called on.
     * @param version The version of the channelOpenInit call, used to identify which chanOpenInit call fees were
     * deposited for.
     * @param ordering The ordering of the channelOpenInit call, used to identify which chanOpenInit call fees were
     * deposited for.
     * @param connectionHops The connection hops for the channelOpenInit call, used to identify which channelOpenInit
     * call fees were deposited for .
     * @param counterpartyPortId The counterparty port ID for the channelOpenInit call, used to identify which
     * channelOpenInit call fees were deposited for .
     */
    function _depositOpenChannelFee(
        IbcDispatcher dispatcher,
        string memory version,
        ChannelOrder ordering,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) internal {
        dispatcher.feeVault().depositOpenChannelFee{value: msg.value}(
            address(this), version, ordering, connectionHops, counterpartyPortId
        );
    }
}
