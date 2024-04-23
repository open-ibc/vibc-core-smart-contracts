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

import {IProofVerifier, L1Header, OpL2StateProof, Ics23Proof} from "../interfaces/IProofVerifier.sol";

/**
 * @title DummyProofVerifier
 * @dev A dummy implementation of the IProofVerifier interface for testing purposes. Does not actually verify any
 * proofs.
 */
contract DummyProofVerifier is IProofVerifier {
    function verifyStateUpdate(L1Header calldata, OpL2StateProof calldata, bytes32, bytes32, uint64) external pure {}

    function verifyMembership(bytes32, bytes calldata, bytes calldata, Ics23Proof calldata) external pure {}

    function verifyNonMembership(bytes32, bytes calldata, Ics23Proof calldata) external pure {}
}
