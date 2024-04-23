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

import {Ics23Proof, L1Header, OpL2StateProof} from "./ProofVerifier.sol";

interface LightClient {
    /**
     * addOpConsensusState adds an appHash to internal store and
     * returns the fraud proof end time, and a bool flag indicating if
     * the fraud proof window has passed according to the block's time stamp.
     */
    function addOpConsensusState(
        L1Header calldata l1header,
        OpL2StateProof calldata proof,
        uint256 height,
        uint256 appHash
    ) external returns (uint256 endTime, bool ended);

    /**
     *
     */
    function getFraudProofEndtime(uint256 height) external returns (uint256 endTime);

    /**
     * verifyMembership checks if the current state
     * can be used to perform the membership test and if so, it uses
     * the verifier to perform membership check.
     */
    function verifyMembership(Ics23Proof calldata proof, bytes memory key, bytes memory expectedValue) external;

    /**
     *
     */
    function verifyNonMembership(Ics23Proof calldata proof, bytes memory key) external;

    /**
     *
     */
    function getState(uint256 height) external view returns (uint256 appHash, uint256 fraudProofEndTime, bool ended);
}
