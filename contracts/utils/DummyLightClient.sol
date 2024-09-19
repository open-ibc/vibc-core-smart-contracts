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

import {ILightClient, Ics23Proof, LightClientType} from "../interfaces/ILightClient.sol";

/**
 * @title DummyLightClient
 * @dev This contract is a dummy implementation of a consensus state manager.
 *      It should only be used for testing purposes.
 *      The logic for checking if the proof length is greater than zero is naive.
 */
contract DummyLightClient is ILightClient {
    uint8 private _LightClientType = uint8(LIGHT_CLIENT_TYPE); // Also redundantly stored as a private mutable type in
    // cheap on-chain use case it needs to be accessed in any proofs
    LightClientType public constant LIGHT_CLIENT_TYPE = LightClientType.SimTestLightClient; // Stored as a constant for

    error InvalidDummyMembershipProof();
    error InvalidDummyNonMembershipProof();

    constructor() {}

    function updateClient(bytes calldata, uint256, uint256) external pure override {}

    function getState(uint256) external pure override returns (uint256 appHash) {
        return (0);
    }

    function verifyMembership(Ics23Proof calldata proof, bytes memory, bytes memory) external pure override {
        if (proof.height == 0) revert InvalidDummyMembershipProof();
    }

    function verifyNonMembership(Ics23Proof calldata proof, bytes memory) external pure override {
        if (proof.height == 0) revert InvalidDummyNonMembershipProof();
    }
}
