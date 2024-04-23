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

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin-upgradeable/contracts/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {IbcChannelReceiver, IbcPacketReceiver} from "../../../contracts/interfaces/IbcReceiver.sol";
import {L1Header, OpL2StateProof, Ics23Proof} from "../../../contracts/interfaces/IProofVerifier.sol";
import {ILightClient} from "../../../contracts/interfaces/ILightClient.sol";
import {IDispatcher} from "../../../contracts/interfaces/IDispatcher.sol";
import {Dispatcher} from "../../../contracts/core/Dispatcher.sol";
import {
    Channel,
    ChannelEnd,
    ChannelOrder,
    IbcPacket,
    ChannelState,
    AckPacket,
    IBCErrors,
    Ibc
} from "../../../contracts/libs/Ibc.sol";
import {IbcUtils} from "../../../contracts/libs/IbcUtils.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @title Dispatcher
 * @author Polymer Labs
 * @notice
 *     Contract callers call this contract to send IBC-like msg,
 *     which can be relayed to a rollup module on the Polymerase chain
 */
contract DispatcherV2 is Dispatcher {}
