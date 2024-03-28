// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IProofVerifier, L1Header, OpL2StateProof, Ics23Proof} from "../interfaces/IProofVerifier.sol";

/**
 * @title DummyProofVerifier
 * @dev A dummy implementation of the IProofVerifier interface for testing purposes. Does not actually verify any proofs.
 */
contract DummyProofVerifier is IProofVerifier {
    function verifyStateUpdate(L1Header calldata, OpL2StateProof calldata, bytes32, bytes32, uint64) external pure {}

    function verifyMembership(bytes32, bytes calldata, bytes calldata, Ics23Proof calldata) external pure {}

    function verifyNonMembership(bytes32, bytes calldata, Ics23Proof calldata) external pure {}
}
