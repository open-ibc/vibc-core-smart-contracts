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
pragma solidity ^0.8.0;

import {ILightClient} from "./ILightClient.sol";

/**
 * @title ILightClient
 * @author Polymer Labs
 * @notice A contract that manages the merkle root(appHash) at different block heights of a chain and tracks the fraud
 * proof end time for them.
 * @notice This is used for state inclusion proofs
 */
interface IOptimisticLightClient is ILightClient {
    /**
     * @dev Returns the fraud proof end time at a given block
     * 0 is returned if there isn't an appHash with the given l2 height.
     */
    function getFraudProofEndtime(uint256 height) external returns (uint256 endTime);

    /**
     * @dev Returns the apphash at a given block height, as well as the fraud proof end time and whether the fraud proof
     * has ended
     */
    function getStateAndEndTime(uint256 height)
        external
        view
        returns (uint256 appHash, uint256 fraudProofEndTime, bool ended);
}
