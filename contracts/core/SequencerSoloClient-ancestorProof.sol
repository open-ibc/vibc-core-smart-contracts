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

import {Ics23Proof} from "../interfaces/IProofVerifier.sol";
import {ISignatureVerifier} from "../interfaces/ISignatureVerifier.sol";
import {ILightClient, LightClientType} from "../interfaces/ILightClient.sol";
import {L1Block} from "optimism/L2/L1Block.sol";
import {console2} from "forge-std/console2.sol";

/**
 * @title SubfinalityLightClient
 * @author Polymer Labs
 * @dev This specific light client implementation uses the same client that is used in the op-stack
 */
contract SequencerSoloClientAncestorProof is ILightClient {
    error InvalidL1Origin();

    LightClientType public constant LIGHT_CLIENT_TYPE = LightClientType.SequencerLightClient; // Stored as a constant
        // for cheap on-chain use
    uint8 private _LightClientType = uint8(LIGHT_CLIENT_TYPE); // Also redundantly stored as a private mutable type in
        // case it needs to be accessed in any proofs

    // consensusStates maps from the height to the appHash.
    mapping(uint256 => uint256) public consensusStates;
    mapping(bytes32 => bool) public checkpoints;

    ISignatureVerifier public immutable verifier;
    L1Block public l1BlockProvider;

    /// Optimization: for the l1 block hash, we can use a ring buffer. reading from this can be done off chain,, so
    /// relayer provides the l1 index
    // SO we need a getter for that option.
    // And then we just look at the block index at this poitn. and then we can remove this/overwrite blockhashes for it.
    // to avoi
    // Always autoincrement the ring buffer, so that you can always auto-increment the prev hash.
    // Note; we need to think about the case where we have a griefer for the ring buffer.
    // Doing this in a hash map is probably nicer here.
    //
    error CannotUpdatePendingOptimisticConsensusState();
    error AppHashHasNotPassedFraudProofWindow();
    error NonMembershipProofsNotYetImplemented();
    error NoConsensusStateAtHeight(uint256 height);

    constructor(ISignatureVerifier verifier_, L1Block _l1BlockProvider) {
        verifier = verifier_;
        l1BlockProvider = _l1BlockProvider;
    }

    /**
     * @inheritdoc ILightClient
     * @param proof An array of bytes that contain the l1blockhash and the sequencer's signature. The first 32 bytes of
     * this argument should be
     * the l1BlockHash, and the remaining bytes should be the sequencer signature which attests to the peptide AppHash
     * for that l1BlockHash
     */
    function updateClient(bytes calldata proof, uint256 peptideHeight, uint256 peptideAppHash) external override {

        // Here proof will contain the signed header like in the previous one. Then we will need to pass in the  
        // the length of block heights from 1 to others. 

        console2.log("length");
        console2.log(proof.length);

        if (l1BlockProvider.hash() != bytes32(proof[:32])) {
            revert InvalidL1Origin();
        }
        _updateClient(proof, peptideHeight, peptideAppHash);
    }

    /* Use if you don't need to use a checkpointed block, to save a SLOAD and 5k gas */
    function updateClientFromCheckpoint(bytes calldata proof, uint256 peptideHeight, uint256 peptideAppHash) external {
        if (!checkpoints[bytes32(proof[:32])]) {
            revert InvalidL1Origin();
        }
        _updateClient(proof, peptideHeight, peptideAppHash);

        delete checkpoints[bytes32(proof[:32])];
    }

    function _updateClient(bytes calldata proof, uint256 peptideHeight, uint256 peptideAppHash) internal {
        if (consensusStates[peptideHeight] != 0) {
            return;
        }
        verifier.verifyStateUpdate(peptideHeight, bytes32(peptideAppHash), bytes32(proof[:32]), proof[32:]);
        consensusStates[peptideHeight] = peptideAppHash;
    }


    ///
   /// l1 bock info is at origin h1 .Peptide l1 origin is at blcok h0. 
   /// We want to prove that h0 is an ancestor of h1.  
   /// We get the h0 hash from the signed payload. 
   /// Now we need the blockdata from h0+1 to h1, such that we can concat each one
   // 
    ///

    /**
     * @inheritdoc ILightClient
     */
    function getState(uint256 height) external view returns (uint256 appHash) {
        return _getState(height);
    }

    /**
     * @dev store l1 block hash so that it can be used in a later transaction
     */
    function checkPoint() external {
        checkpoints[l1BlockProvider.hash()] = true;
    }

    /**
     * @inheritdoc ILightClient
     */
    function verifyMembership(Ics23Proof calldata proof, bytes calldata key, bytes calldata expectedValue)
        external
        view
    {
        // a proof generated at height H can only be verified against state root (app hash) from block H - 1.
        // this means the relayer must have updated the contract with the app hash from the previous block and
        // that is why we use proof.height - 1 here.
        verifier.verifyMembership(bytes32(_getState(proof.height - 1)), key, expectedValue, proof);
    }

    /**
     * @inheritdoc ILightClient
     */
    function verifyNonMembership(Ics23Proof calldata, bytes calldata) external pure {
        revert NonMembershipProofsNotYetImplemented();
    }

    /**
     * @dev Returns the internal state of the light client at a given height.
     */
    function _getState(uint256 height) internal view returns (uint256 appHash) {
        return consensusStates[height];
    }
}
