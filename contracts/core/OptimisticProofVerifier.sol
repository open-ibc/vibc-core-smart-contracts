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

import {SecureMerkleTrie} from "optimism/libraries/trie/SecureMerkleTrie.sol";
import {RLPReader} from "optimism/libraries/rlp/RLPReader.sol";
import {RLPWriter} from "optimism/libraries/rlp/RLPWriter.sol";
import {L1Header, OpL2StateProof, IProofVerifier} from "../interfaces/IProofVerifier.sol";
import {AppStateVerifier} from "../base/AppStateVerifier.sol";
import {IProofVerifier} from "../interfaces/IProofVerifier.sol";

/**
 * @title OptimisticProofVerifier
 * @notice Verifies proofs related to Optimistic Rollup state updates
 * @author Polymer Labs
 */
contract OptimisticProofVerifier is AppStateVerifier, IProofVerifier {
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for bytes;

    // @notice index of the l1 state root in the "l1 header"
    uint256 internal constant _L1_STATE_ROOT_INDEX = 3;

    // @notice index of the l1 number in the "l1 header"
    uint256 internal constant _L1_NUMBER_INDEX = 8;

    // @notice known L2 Output Oracle contract address to verify state update proofs against
    address public l2OutputOracleAddress;

    constructor(address _l2OutputOracleAddress) {
        l2OutputOracleAddress = _l2OutputOracleAddress;
    }

    /**
     * @dev Prove that the provided app hash (L2 state root) is valid. Done so by proving that the L2OutputOracle
     * contains an output proposal within its state that we can derive from the given app hash. The high level
     * approach is:
     *
     * A. Prove the given L1 state root.
     * B. Prove the prescence of an output proposal in the L2OutputOracle contract.
     * C. Derive the peptide apphash from the output propsal.
     *
     * A more detailed explanation of the process goes as follows. All steps must be valid in order for the
     * app hash to be accepted. Otherwise, the function will revert.
     *
     * 1. Provided L1 header hash and number match that of the trusted ones. The trusted attributes must come
     *    from Optimism's L1Block contract. Given that this contract only holds the latest L1 attributes, there's
     *    a good chance for a race-condition to happen, so check them first.
     *
     * 2. Provided L1 header data includes the L1 state root. Compute the header hash and check it against the
     *    trusted one. In case of a match, the state root must be valid.
     *
     * 3. Based on the L1 state root and using the provided account proof and L2OutputOracle address, get the
     *    value stored in the MerkleTrie leaf. This is the state account.
     *
     * 4. With the state account root and using the provided storage proof and output proposal key, get the
     *    vlue stored in the MerkleTrie leaf. This is the output proposal root.
     *
     * 5. With the provided peptide apphash and L2 Block hash, try to compute a new output root and match it against
     *    the one we just proved to be valid.
     */
    function verifyStateUpdate(
        L1Header calldata l1header,
        OpL2StateProof calldata proof,
        bytes32 peptideAppHash,
        bytes32 trustedL1BlockHash,
        uint64 trustedL1BlockNumber
    ) external view {
        if (trustedL1BlockNumber != l1header.number) {
            revert InvalidL1BlockNumber();
        }

        // This computes the L1 header hash
        if (trustedL1BlockHash != keccak256(RLPWriter.writeList(l1header.header))) {
            revert InvalidL1BlockHash();
        }

        // These two checks are here to verify that the "plain" (i.e. not RLP encoded) values in the l1header are
        // the same ones found in l1header.header (i.e. RLP encoded). This is because it is cheaper to RLP
        // encode than decode
        if (keccak256(RLPWriter.writeUint(l1header.number)) != keccak256(l1header.header[_L1_NUMBER_INDEX])) {
            revert InvalidRLPEncodedL1BlockNumber();
        }
        if (
            keccak256(RLPWriter.writeBytes(abi.encode(l1header.stateRoot)))
                != keccak256(l1header.header[_L1_STATE_ROOT_INDEX])
        ) {
            revert InvalidRLPEncodedL1StateRoot();
        }

        //  stateAccount looks like this struct. We are interested in the Root field which is the one at index 2
        //  type StateAccount struct {
        //     Nonce    uint64       // index 0
        //     Balance  *big.Int     // index 1
        //     Root     common.Hash  // index 2
        //     CodeHash []byte       // index 3
        //  }
        RLPReader.RLPItem[] memory stateAccount = SecureMerkleTrie.get(
            abi.encodePacked(l2OutputOracleAddress), proof.accountProof, l1header.stateRoot
        ).toRLPItem().readList();

        bytes memory outputRoot = SecureMerkleTrie.get(
            abi.encode(proof.l2OutputProposalKey), proof.outputRootProof, bytes32(bytes(stateAccount[2].readBytes()))
        );

        // Now that the output root is verified, we need to verify the app hash. To do so we try to derive the
        // the output root the same way the proposer did.
        // See https://github.com/polymerdao/optimism/blob/polymer/v1.2.0/op-service/eth/output.go#L44
        if (
            keccak256(
                abi.encodePacked(
                    bytes32(0), // version
                    peptideAppHash,
                    bytes32(0), // message passer storage root.
                    proof.l2BlockHash
                )
            ) != bytes32(bytes(outputRoot.toRLPItem().readBytes()))
        ) {
            revert InvalidAppHash();
        }
    }
}
