// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../contracts/core/OptimisticLightClient.sol";
import "../contracts/utils/DummyProofVerifier.sol";
import "./utils/Proof.base.t.sol";

contract OptimisticLightClientTest is ProofBase {
    OptimisticLightClient manager;
    IProofVerifier verifier;

    constructor() {
        verifier = new DummyProofVerifier();
    }

    function setUp() public override {
        super.setUp();
        manager = new OptimisticLightClient(1, verifier, l1BlockProvider);
    }

    function test_addOpConsensusState_newOpConsensusStateCreatedWithPendingStatus() public {
        (, bool ended) = manager.updateClient(emptyl1header, abi.encode(invalidStateProof), 1, 1);
        assertEq(false, ended);
    }

    function test_addOpConsensusState_addingAlreadyTrustedOpConsensusStateIsNoop() public {
        manager.updateClient(emptyl1header, abi.encode(invalidStateProof), 1, 1);

        // fast forward time.
        vm.warp(block.timestamp + 100);

        // the fraud proof window has passed.
        manager.updateClient(emptyl1header, abi.encode(invalidStateProof), 1, 1);

        (,, bool ended) = manager.getState(1);
        assertEq(true, ended);
    }

    function test_addOpConsensusState_addingPendingOpConsensusStateWithDifferentValuesIsError() public {
        manager.updateClient(emptyl1header, abi.encode(invalidStateProof), 1, 1);

        vm.expectRevert(OptimisticLightClient.CannotUpdatePendingOptimisticConsensusState.selector);
        manager.updateClient(emptyl1header, abi.encode(invalidStateProof), 1, 2);
    }

    function test_addOpConsensusState_addingSameOpConsensusStateIsNoop() public {
        manager.updateClient(emptyl1header, abi.encode(invalidStateProof), 1, 1);

        (, uint256 originalFraudProofEndTime,) = manager.getState(1);

        vm.warp(block.timestamp + 1);

        // adding the same appHash later doesn't update the fraud
        // proof end time.
        manager.updateClient(emptyl1header, abi.encode(invalidStateProof), 1, 1);
        (, uint256 newFraudProofEndTime,) = manager.getState(1);
        assertEq(originalFraudProofEndTime, newFraudProofEndTime);
    }

    function test_zero_proof_window() public {
        manager = new OptimisticLightClient(0, verifier, l1BlockProvider);
        manager.updateClient(emptyl1header, abi.encode(invalidStateProof), 1, 1);
        (,, bool ended) = manager.getState(1);
        assertEq(true, ended);
    }

    function test_getState_nonExist() public {
        (uint256 appHash,, bool ended) = manager.getState(1);
        assertEq(0, appHash);
        assertEq(false, ended);
    }
}

contract OptimisticLightClientWithRealVerifierTest is ProofBase {
    OptimisticLightClient manager;

    function setUp() public override {
        super.setUp();
        manager = new OptimisticLightClient(1, opProofVerifier, l1BlockProvider);
    }

    function test_addOpConsensusState_newAppHashWithValidProof() public {
        // trick the L1Block contract into thinking it is updated with the right l1 header
        setL1BlockAttributes(keccak256(RLPWriter.writeList(l1header.header)), l1header.number);

        manager.updateClient(l1header, abi.encode(validStateProof), 1, uint256(apphash));

        // since we are setting using an already known apphash, the proof is ignored
        manager.updateClient(emptyl1header, abi.encode(invalidStateProof), 1, uint256(apphash));
    }

    function test_addOpConsensusState_newAppHashWithInvalidProof() public {
        setL1BlockAttributes(keccak256(RLPWriter.writeList(l1header.header)), l1header.number);
        vm.expectRevert("MerkleTrie: ran out of proof elements");
        manager.updateClient(l1header, abi.encode(invalidStateProof), 1, uint256(apphash));
    }
}
