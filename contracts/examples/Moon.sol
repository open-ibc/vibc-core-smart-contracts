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

import {Mars} from "./Mars.sol";
import {IbcDispatcher} from "../interfaces/IbcDispatcher.sol";

/**
 * @title Moon
 * @notice Moon is a simple IBC receiver contract that receives packets and sends acks.
 * For now, it is a copy of Mars.sol, but may be extended in the future.
 * @dev This contract is used for only testing IBC functionality and as an example for dapp developers
 *  on how to integrate with the vibc protocol.
 */
contract Moon is Mars {
    constructor(IbcDispatcher _dispatcher) Mars(_dispatcher) {}
}
