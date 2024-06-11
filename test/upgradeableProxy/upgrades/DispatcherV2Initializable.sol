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

import {DispatcherV2} from "./DispatcherV2.sol";
import {ILightClient} from "../../../contracts/interfaces/ILightClient.sol";
import {IFeeVault} from "../../../contracts/interfaces/IFeeVault.sol";
import {IBCErrors} from "../../../contracts/libs/IbcErrors.sol";

/**
 * @title Dispatcher
 * @author Polymer Labs
 * @notice
 *     Contract callers call this contract to send IBC-like msg,
 *     which can be relayed to a rollup module on the Polymerase chain
 */
contract DispatcherV2Initializable is DispatcherV2 {
    function initialize(string memory initPortPrefix, IFeeVault _feeVault) public override reinitializer(2) {
        __Ownable_init();
        portPrefix = initPortPrefix;
        portPrefixLen = uint32(bytes(initPortPrefix).length);
    }
}
