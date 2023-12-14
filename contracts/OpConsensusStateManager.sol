// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./IbcVerifier.sol";
import "./Ibc.sol";
import "./ConsensusStateManager.sol";

// OptimisticConsensusStateManager manages the appHash at different
// heights and track the fraud proof end time for them.
contract OptimisticConsensusStateManager is ConsensusStateManager {
    uint32 fraudProofWindowSeconds;
    ZKMintVerifier verifier;

    // consensusStates maps from the height to the appHash.
    mapping(uint256 => uint256) consensusStates;

    // fraudProofEndtime maps from the appHash to the fraud proof end time.
    mapping(uint256 => uint256) fraudProofEndtime;

    constructor(uint32 fraudProofWindowSeconds_, ZKMintVerifier verifier_) {
        fraudProofWindowSeconds = fraudProofWindowSeconds_;
        verifier = verifier_;
    }

    // addOpConsensusState adds an appHash to internal store and
    // returns the fraud proof end time, and a bool flag indicating if
    // the fraud proof window has passed according to the block's
    // timestamp.
    function addOpConsensusState(uint256 height, uint256 appHash)
        external
        override
        returns (uint256 fraudProofEndTime, bool ended)
    {
        uint256 hash = consensusStates[height];
        if (hash == 0) {
            // a new appHash
            consensusStates[height] = appHash;
            uint256 endTime = block.timestamp + fraudProofWindowSeconds;
            consensusStates[height] = appHash;
            fraudProofEndtime[appHash] = endTime;
            return (endTime, false);
        }

        if (hash == appHash) {
            uint256 endTime = fraudProofEndtime[hash];
            return (endTime, block.timestamp >= endTime);
        }

        revert(
            "cannot update a pending optimistic consensus state with a different appHash, please submit fraud proof instead"
        );
    }

    /**
     * getState returns the appHash at the given height, and the fraud
     * proof end time.
     * 0 is returned if there isn't an appHash with the given height.
     */
    function getState(uint256 height) external view returns (uint256 appHash, uint256 fraudProofEndTime, bool ended) {
        return getInternalState(height);
    }

    function getInternalState(uint256 height)
        public
        view
        returns (uint256 appHash, uint256 fraudProofEndTime, bool ended)
    {
        uint256 hash = consensusStates[height];
        return (hash, fraudProofEndtime[hash], hash != 0 && block.timestamp >= fraudProofEndtime[hash]);
    }

    function getFraudProofEndtime(uint256 height) external view returns (uint256 fraudProofEndTime) {
        uint256 hash = consensusStates[height];
        return fraudProofEndtime[hash];
    }

    /**
     * verifyMembership checks if the current trustedOptimisticConsensusState state
     * can be used to perform the membership test and if so, it uses
     * the verifier to perform membership check.
     */
    function verifyMembership(Proof calldata proof, bytes memory key, bytes memory expectedValue, string memory message)
        external
        view
    {
        (uint256 appHash,, bool ended) = getInternalState(proof.proofHeight.revision_height);
        require(ended, "appHash hasn't passed the fraud proof window");
        require(verifier.verifyMembership(appHash, proof, key, expectedValue), message);
    }

    function verifyNonMembership(Proof calldata proof, bytes memory key, string memory message) external view {
        (uint256 appHash,, bool ended) = getInternalState(proof.proofHeight.revision_height);
        require(ended, "appHash hasn't passed the fraud proof window");
        require(verifier.verifyNonMembership(appHash, proof, key), message);
    }
}
