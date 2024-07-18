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

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IFeeVault} from "../interfaces/IFeeVault.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import {ChannelOrder} from "../libs/Ibc.sol";

contract FeeVault is Ownable, ReentrancyGuard, IFeeVault {
    /**
     * @notice Deposits the send packet fee for a given channel and sequence that is used for relaying recieve and
     * acknowledge steps of a packet handhsake after a dapp has called the sendPacket on dispatcher.
     * @notice If you are relaying your own packets, you should not call this method.
     * @notice Use the Polymer fee estimation api to get the required fees to ensure that enough fees are sent.
     * @dev This function calculates the required fee based on provided gas limits and gas prices,
     *      and reverts if the sent value does not match the calculated fee.
     *      The first entry in `gasLimits` and `gasPrices` arrays corresponds to `recvPacket` fees,
     *      and the second entry corresponds to `ackPacket` fees.
     * @param channelId The identifier of the channel.
     * @param sequence The sequence number for the packet, returned from the dispatcher sendPacket call.
     * @param gasLimits An array containing two gas limit values:
     *                  - gasLimits[0] for `recvPacket` fees
     *                  - gasLimits[1] for `ackPacket` fees.
     * @param gasPrices An array containing two gas price values:
     *                  - gasPrices[0] for `recvPacket` fees, for the dest chain
     *                  - gasPrices[1] for `ackPacket` fees, for the src chain
     * @notice The total fees sent in the msg.value should be equal to the total gasLimits[0] * gasPrices[0] +
     * gasLimits[1] * gasPrices[1]. The transaction will revert if a higher or lower value is sent
     * @dev Note: if you're having trouble with your packet data being mysteriously lost, try passing in the gasLimits
     * and gasPrices as memory, solidity sometimes misbehaves when trying to pass in too much calldata.
     */
    function depositSendPacketFee(
        bytes32 channelId,
        uint64 sequence,
        uint256[2] calldata gasLimits,
        uint256[2] calldata gasPrices
    ) external payable nonReentrant {
        uint256 fee = gasLimits[0] * gasPrices[0] + gasLimits[1] * gasPrices[1];
        if ((fee) != msg.value) {
            revert IncorrectFeeSent(fee, msg.value);
        }
        emit SendPacketFeeDeposited(channelId, sequence, gasLimits, gasPrices);
    }

    /**
     * @notice Deposits the fee for a channel handshake, to pay a relayer for relaying the channelOpenTry,
     * channelOpenConfirm, and channelOpenAck steps after a dapp has called channelOpenInit
     * @notice If you are relaying your own channelHandshake transactions, you should not call this method.
     * @dev The fee amount that needs to be sent for Polymer to relay the whole channel handshake can be queried on the
     * web2 layer.
     * @param src The address of the sender, should be the address in the localportId.
     * @param version The version string of the channel, the same argument as that sent in the
     * dispatcher.channelOpenInit call
     * @param ordering The ordering of the channel, the same argument as that sent in the dispatcher.channelOpenInit
     * call
     * @param connectionHops An array of connection hops, the same argument as that sent in the
     * dispatcher.channelOpenInit call
     * @param counterpartyPortId The counterparty port identifier, the same argument as that sent in the
     * dispatcher.channelOpenInit call
     */
    function depositOpenChannelFee(
        address src,
        string memory version,
        ChannelOrder ordering,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) external payable nonReentrant {
        if (msg.value == 0) {
            revert NoFeeSent();
        }
        emit OpenChannelFeeDeposited(src, version, ordering, connectionHops, counterpartyPortId, msg.value);
    }

    /**
     * @notice Withdraws all collected fees to the contract owner's address.
     * @dev Transfers the entire balance of this contract to the owner.
     * @dev Anyone can call this, but it will always only be sent to the owner.
     */
    function withdrawFeesToOwner() external {
        payable(owner()).transfer(address(this).balance);
    }
}
