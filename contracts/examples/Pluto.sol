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

import {AckPacket} from "../libs/Ibc.sol";
import {Mars} from "./Mars.sol";
import {IbcPacket} from "../interfaces/IbcReceiver.sol";
import {IbcDispatcher} from "../interfaces/IbcDispatcher.sol";

/**
 * @title Pluto
 * @notice Pluto is a simple IBC receiver contract that receives packets without sending acks.
 * @dev This contract is used for only testing IBC functionality and as an example for dapp developers
 *  on how to integrate with the vibc protocol.
 */
contract Pluto is Mars {
    constructor(IbcDispatcher _dispatcher) Mars(_dispatcher) {}

    /**
     * @notice Callback for receiving a packet; triggered when a counterparty sends an an IBC packet
     * @param packet The IBC packet received
     * @return ackPacket The acknowledgement packet generated in response
     * @dev Make sure to validate packet's source and destiation channels and ports.
     */
    function onRecvPacket(IbcPacket memory packet)
        external
        override
        onlyIbcDispatcher
        returns (AckPacket memory ackPacket, bool skipAck)
    {
        recvedPackets.push(packet);

        // solhint-disable-next-line quotes
        return (AckPacket(true, abi.encodePacked('{ "account": "account", "reply": "got the message" }')), true);
    }
}
