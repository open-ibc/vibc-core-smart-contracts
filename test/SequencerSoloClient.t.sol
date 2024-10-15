// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import {SequencerSoloClient} from "../contracts/core/SequencerSoloClient.sol";
import {SequencerSignatureVerifier, ISignatureVerifier} from "../contracts/core/SequencerSignatureVerifier.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {SigningBase} from "./utils/Signing.base.t.sol";

contract OptimisticLightClientWithRealVerifierTest is SigningBase {
    SequencerSoloClient sequencerClient;
    bytes32 ancestorBlockPayloadHash;
    bytes validChildSignature;
    bytes validAncestorSignature;
    bytes invalidSignature;

    function setUp() public {
        sequencerClient = new SequencerSoloClient(sigVerifier, l1BlockProvider);

        // Used for l1 origin check tests, represents the child block.
        bytes32 ancestorPayloadHash =
            keccak256(bytes.concat(bytes32(peptideBlockNumber), peptideAppHash, l1AncestorBlockHash));
        bytes32 ancestorHashToSign = keccak256(bytes.concat(domain, PEPTIDE_CHAIN_ID, ancestorPayloadHash));
        (uint8 v0, bytes32 r0, bytes32 s0) = vm.sign(sequencerPkey, ancestorHashToSign);
        validAncestorSignature = abi.encodePacked(r0, s0, v0); // Annoingly, v, r, s are in a different order than those

        (uint8 v1, bytes32 r1, bytes32 s1) = vm.sign(notSequencerPkey, ancestorHashToSign);
        invalidSignature = abi.encodePacked(r1, s1, v1); // Annoingly, v, r, s are in a different order than those
    }

    function test_UpdateValidSequencerClientState() public {
        // In this case, peptide l1 origin is ahead of dest chain's l1 origin. Relayer queries the
        // signed block from sequencer and waits for dest chain to come to same l1 origin.
        // We don't need a checkpointing solution in this case, so we call the update method directly.

        setBlockAttributesFromJson("/test/payload/l1_block_ancestor.json");
        sequencerClient.updateClient(
            abi.encodePacked(l1AncestorBlockHash, validAncestorSignature), peptideBlockNumber, uint256(peptideAppHash)
        );
        assertEq(sequencerClient.getState(peptideBlockNumber), uint256(peptideAppHash));

        // A subsequent client update should be a noop and wouldn't revert even with an invalid sig
        sequencerClient.updateClient(
            abi.encodePacked(l1AncestorBlockHash, invalidSignature), peptideBlockNumber, uint256(peptideAppHash)
        );
        assertEq(sequencerClient.getState(peptideBlockNumber), uint256(peptideAppHash));
    }

    // 1.) test that hte block update without the l1b lock header is invalid, or perhaps signatuer is invalid as well
    function test_revert_invalidSequencerClientState() public {
        // This mocks the dest chain moving ahead of the signed l1 origin, or a re-org occuring which means that the
        // blockhash returned by l1 block info isn't equal to that of the signed blockhash. i.e. We'd expect this update
        // client tx to revert.
        setBlockAttributesFromJson("/test/payload/l1_block_child.json");
        console.logBytes32(l1BlockProvider.hash());
        vm.expectRevert(SequencerSoloClient.InvalidL1Origin.selector);
        sequencerClient.updateClient(
            abi.encodePacked(l1AncestorBlockHash, validAncestorSignature), peptideBlockNumber, uint256(peptideAppHash)
        );
        assertEq(sequencerClient.getState(peptideBlockNumber), 0);

        setBlockAttributesFromJson("/test/payload/l1_block_ancestor.json");
        // Peptide l1 origin being the same as dest chain l1 origin is necessary but not sufficient. We also need to
        // check that the signature is valid as well, and it will revert if it isn't valid.
        vm.expectRevert(ISignatureVerifier.InvalidSequencerSignature.selector);
        sequencerClient.updateClient(
            abi.encodePacked(l1AncestorBlockHash, invalidSignature), peptideBlockNumber, uint256(peptideAppHash)
        );
        assertEq(sequencerClient.getState(peptideBlockNumber), 0);
    }
}
