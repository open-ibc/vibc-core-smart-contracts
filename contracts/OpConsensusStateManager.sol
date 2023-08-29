// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import './IbcVerifier.sol';


// OptimisticConsensusStateManager manages the appHash at different
// heights and track the fraud proof end time for them.
library OptimisticConsensusStateManager {
    struct DataStore {
        uint32 fraudProofWindowSeconds;

        // consensusStates maps from the height to the appHash.
        mapping(uint256 => uint256) consensusStates;

        // fraudProofEndtime maps from the appHash to the fraud proof end time.
        mapping(uint256 => uint256) fraudProofEndtime;
    }

    // addOpConsensusState adds an appHash to internal store and
    // returns the fraud proof end time, and a bool flag indicating if
    // the fraud proof window has passed according to the block's
    // timestamp.
    function addOpConsensusState(DataStore storage self, uint256 height, uint256 appHash)
        external returns (uint256 fraudProofEndTime, bool ended) {
        uint256 hash = self.consensusStates[height];
        if (hash == 0) {
            // a new appHash
            self.consensusStates[height] = appHash;
            uint256 endTime = block.timestamp + self.fraudProofWindowSeconds;
            self.consensusStates[height] = appHash;
            self.fraudProofEndtime[appHash] = endTime;
            return (endTime, false);
        }

        if (hash == appHash) {
            uint256 endTime = self.fraudProofEndtime[hash];
            return (endTime, block.timestamp > endTime);
        }

        revert('cannot update a pending optimistic consensus state with a different appHash, please submit fraud proof instead');
    }

    // getState returns the appHash at the given height, and the fraud
    // proof end time.
    //
    // 0 is returned if there isn't an appHash with the given height.
    function getState(DataStore storage self, uint256 height) public
        returns (uint256 appHash, uint256 fraudProofEndTime, bool ended) {
        uint256 hash = self.consensusStates[height];
        return (hash, self.fraudProofEndtime[hash], hash != 0 &&  block.timestamp > self.fraudProofEndtime[hash]);
    }
}
