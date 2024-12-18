// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import {SigningBase} from "../utils/Signing.base.t.sol";
import {CrossL2Prover} from "contracts/core/proofAPI/CrossL2Prover.sol";

import {ISignatureVerifier} from "contracts/interfaces/ISignatureVerifier.sol";
import {ICrossL2Prover} from "contracts/interfaces/ICrossL2Prover.sol";
import {Ics23Proof} from "contracts/interfaces/IProofVerifier.sol";
import {MerkleTrie} from "optimism/libraries/trie/MerkleTrie.sol";

contract CrossL2InboxBase is SigningBase {
    using stdJson for string;

    CrossL2Prover crossProver;

    // Sequencer signing related
    bytes invalidSignature;
    bytes validAncestorSignature;

    // Proof related
    bytes32 receiptRoot;
    bytes receiptIdx;
    bytes rlpEncodedReceipt;
    bytes proof; // Opaque proof bytes returned by polymer proof api
    bytes[] receiptMMPTProof;

    bytes32 logHash;
    uint256 logIdx;
    string peptideClientId = "proof_api-0";
    bytes32 proofPeptideAppHash = hex"59c8995c4fc9c3f2f7bec810578235d1c0ffa3ba20fac3aa084c84fa229b314f";

    function setUp() public {
        crossProver = new CrossL2Prover(sigVerifier, peptideClientId);

        // Used for l1 origin check tests, represents the child block.
        bytes32 ancestorPayloadHash =
            keccak256(bytes.concat(bytes32(peptideBlockNumber), peptideAppHash, l1AncestorBlockHash));
        bytes32 ancestorHashToSign = keccak256(bytes.concat(domain, PEPTIDE_CHAIN_ID, ancestorPayloadHash));
        (uint8 v0, bytes32 r0, bytes32 s0) = vm.sign(sequencerPkey, ancestorHashToSign);
        validAncestorSignature = abi.encodePacked(r0, s0, v0); // Annoingly, v, r, s are in a different order than those

        (uint8 v1, bytes32 r1, bytes32 s1) = vm.sign(notSequencerPkey, ancestorHashToSign);
        invalidSignature = abi.encodePacked(r1, s1, v1); // Annoingly, v, r, s are in a different order than those

        string memory input = vm.readFile(string.concat(rootDir, "/test/payload/proof-api.hex"));
        proof = vm.parseBytes(input);
        // (proof) = abi.decode(encoded, (bytes));

        // (receiptIdx, rlpEncodedReceipt, proof) = (receiptIdxM, rlpEncodedReceiptM, receiptProofM);

        (Ics23Proof memory iavlProofM, bytes[] memory receiptMMPTProofM,, uint64 height,,) =
            abi.decode(proof, (Ics23Proof, bytes[], bytes32, uint64, bytes32, bytes));

        console2.log("height", height);
        console2.logBytes(proof);
        // console2.logBytes(receiptIdxM);

        do_client_update(proofPeptideAppHash, iavlProofM.height - 1);
    }

    // Happy path for CrossEventProver.updateClient()
    function test_clientUpdate_success() public {
        crossProver.updateClient(
            abi.encodePacked(l1AncestorBlockHash, validAncestorSignature), peptideBlockNumber, uint256(peptideAppHash)
        );
        assertEq(crossProver.getState(peptideBlockNumber), uint256(peptideAppHash));
    }

    // Happy path for CrossEventProver.validateReceipt()
    function test_validate_receipt_success() public {
        // TODO: Fill this in once we have relaying to peptide working and that we can generate peptide state proof
        // using this tool in relayer code at polymerase/
        // 1.) Run the oprelayer/pkg/encoder/cmd/main.go command to generate hex files  (note: this will need to be
        // modified to work with receipt proofs instead of channel/packet proofs)
        // 2.) load these proof files into the test contracts using the load_proof util, e.g. using something like
        // Ics23Proof memory testProof = load_proof("/test/payload/channel_confirm_pending_proof.hex",
        // address(dummyLightClient));
        // 3.) Run the validate_receipt call below

        // (, Ics23Proof memory peptideProof) = abi.decode(encoded, (bytes32, Ics23Proof));

        // // proof will be struct EventProof
        // bytes memory proof = abi.encode(peptideProof, proof, receiptRoot, peptideClientId, peptideBlockNumber);

        (string memory chainId, bytes memory rlpBytes) = crossProver.validateReceipt(proof);
        bytes memory expectedReceipt =
            hex"b9042802f904240183028f07b9010000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000200000000000000000100000000400000040000100001002000000000000001000000000000000000000000000000020000000000000000000800000000000000000000000000000000000000000000000000000000000000000000000480000000000000000000000000000000001000000000000000000000000000000000000000800000000000000000000000000000020000004000020000024000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000f90319f901dc9456a43eb56da12c0dc1d972acb089c06a5def8e69f842a0f6a97944f31ea060dfde0566e4167c1a1082551e64b60ecb14d599a9d023d451a00000000000000000000000000000000000000000000000000000000000018fccb90180000000000000000000000000000000000000000000000000000000042dd526ab000000000000000000000000ea1c859797599b47b7c6cff59e1a3cbe9fd0c2bd00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000001400000000000000000000000fd26636b88f82af90289a759d4cf414e00067844060000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000042d865ed0000000000000000000000000000000000000000000000000000000042da19acb000000000000000000000000000000000000000000000000000000042dd526ab000000000000000000000000000000000000000000000000000000042e157f7000000000000000000000000000000000000000000000000000000000000000040002030100000000000000000000000000000000000000000000000000000000f89b9456a43eb56da12c0dc1d972acb089c06a5def8e69f863a00109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271a00000000000000000000000000000000000000000000000000000000000018fcca00000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000006744c8a6f89b9456a43eb56da12c0dc1d972acb089c06a5def8e69f863a00559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5fa0000000000000000000000000000000000000000000000000000000042dd526aba00000000000000000000000000000000000000000000000000000000000018fcca0000000000000000000000000000000000000000000000000000000006744c8a6";

        console2.log("expected receipt:");
        console2.logBytes(expectedReceipt);
        console2.log("received receipt:");
        console2.logBytes(rlpBytes);
        assertEq(keccak256(abi.encode(rlpBytes)), keccak256(abi.encode(expectedReceipt)));
    }

    // Happy path for CrossEventProver.validateEvent()
    function test_validate_event_success() public {
        // TODO: Fill this in once we have relaying to peptide working and that we can generate peptide state proof
        // using this tool in relayer code at polymerase/
        // 1.) Run the oprelayer/pkg/encoder/cmd/main.go command to generate hex files  (note: this will need to be
        // modified to work with receipt proofs instead of channel/packet proofs)
        // 2.) load these proof files into the test contracts using the load_proof util, e.g. using something like
        // Ics23Proof memory testProof = load_proof("/test/payload/channel_confirm_pending_proof.hex",
        // address(dummyLightClient));
        // 3.) Call the crossProver.validateEvent method to test

        vm.skip(true);
    }

    // Test valid peptide proof but invalid MMPT receipt proof
    function test_revert_invalidReceiptProof() public {
        // TODO: Fill this in once we have relaying to peptide working and that we can generate peptide state proof
        // using this tool in relayer code at polymerase/
        // 1.) Run the oprelayer/pkg/encoder/cmd/main.go command to generate hex files  (note: this will need to be
        // modified to work with receipt proofs instead of channel/packet proofs)
        // 2.) load these proof files into the test contracts using the load_proof util, e.g. using something like
        // Ics23Proof memory testProof = load_proof("/test/payload/channel_confirm_pending_proof.hex",
        // address(dummyLightClient));
        // 3.) Call the crossProver.validateReceipt method to test

        vm.skip(true);
    }

    // Test trying to prove with a non-existent peptideApphash that hasn't yet been seen
    function test_revert_nonexistingPeptideAppHash() public {
        string memory input = vm.readFile(string.concat(rootDir, "/test/payload/packet_ack_proof.hex"));
        bytes memory encoded = vm.parseBytes(input);
        (, Ics23Proof memory peptideProof) = abi.decode(encoded, (bytes32, Ics23Proof));

        // proof will be struct EventProof
        // bytes memory proof = abi.encode(peptideProof, proof, receiptRoot, peptideClientId, peptideBlockNumber);

        vm.expectRevert();
        crossProver.validateReceipt(proof);
    }

    // Test revert to prove a peptide apphash which has been seen but which doesn't prove the MPT receipt root given
    function test_revert_invalidReceiptRoot() public {
        // TODO: Fill this in once we have relaying to peptide working and that we can generate peptide state proof
        // using this tool in relayer code at polymerase/
        // 1.) Run the oprelayer/pkg/encoder/cmd/main.go command to generate hex files  (note: this will need to be
        // modified to work with receipt proofs instead of channel/packet proofs)
        // 2.) load these proof files into the test contracts using the load_proof util, e.g. using something like
        // Ics23Proof memory testProof = load_proof("/test/payload/channel_confirm_pending_proof.hex",
        // address(dummyLightClient));
        // 3.) Call the crossProver.validateReceipt method to test revert
    }

    // Test that a client update will fail if it doesn't have a valid sequencer signature
    function test_revert_invalidClentUpdateSignature() public {
        vm.expectRevert(ISignatureVerifier.InvalidSequencerSignature.selector);
        crossProver.updateClient(
            abi.encodePacked(l1AncestorBlockHash, invalidSignature), peptideBlockNumber, uint256(peptideAppHash)
        );
        assertEq(crossProver.getState(peptideBlockNumber), 0);
    }

    // Test valid receipt proof but invalid event in receipt
    function test_invalidEvent() public {
        // TODO: Fill this in once we have relaying to peptide working and that we can generate peptide state proof
        // using this tool in relayer code at polymerase/
        // 1.) Run the oprelayer/pkg/encoder/cmd/main.go command to generate hex files  (note: this will need to be
        // modified to work with receipt proofs instead of channel/packet proofs)
        // 2.) load these proof files into the test contracts using the load_proof util, e.g. using something like
        // Ics23Proof memory testProof = load_proof("/test/payload/channel_confirm_pending_proof.hex",
        // address(dummyLightClient));
        // 3.) Call the crossProver.validateEvent method to test revert

        vm.skip(true);
    }

    // Do a client update to our cross l2 prover contract through calculating who the signer should be for the deployed
    // contract
    function do_client_update(bytes32 appHash, uint256 height) internal {
        bytes32 ancestorPayloadHash = keccak256(bytes.concat(bytes32(height), appHash, l1AncestorBlockHash));
        bytes32 ancestorHashToSign = keccak256(bytes.concat(domain, PEPTIDE_CHAIN_ID, ancestorPayloadHash));
        (uint8 v0, bytes32 r0, bytes32 s0) = vm.sign(sequencerPkey, ancestorHashToSign);
        validAncestorSignature = abi.encodePacked(r0, s0, v0); // Annoingly, v, r, s are in a different order than those

        crossProver.updateClient(
            abi.encodePacked(l1AncestorBlockHash, validAncestorSignature), height, uint256(appHash)
        );
    }
}
