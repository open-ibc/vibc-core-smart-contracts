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
    bytes[] receiptProof;
    bytes32 logHash;
    uint256 logIdx;
    string peptideClientId;

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

        receiptRoot = hex"91de58019987854102326f49d235578cec08f71ddd5333c9e340ed63149d24a4";
        receiptIdx = hex"01";

        rlpEncodedReceipt =
            hex"02f9038801834008fdb9010000000000000001000000000000000000000000000000000000040000000001000000000000000000000000100000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000008000000000010000000000000400000000000000000000000200000000000000000002000000000000000040000000000000000000000001002100000000000020000000000000000000000000000000010000000000400000000000000000000000000000000000000000000000000000000000000000180000200000000000000080000020000002000000000000000000000000000000000000400000000000000000f9027df87a944200000000000000000000000000000000000006f842a0e1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109ca00000000000000000000000006f26bf09b1c792e3228e5467807a900a503c0281a0000000000000000000000000000000000000000000000000051ef38b821e8000f901fe946f26bf09b1c792e3228e5467807a900a503c0281f884a0a123dc29aebf7d0c3322c8eeb5b999e859f39937950ed31056532713d0de396fa00000000000000000000000000000000000000000000000000000000000002105a00000000000000000000000000000000000000000000000000000000000203bfba00000000000000000000000009d3976c25f4defe584ed80bae5a7cef59ba07aa5b9016000000000000000000000000042000000000000000000000000000000000000060000000000000000000000004200000000000000000000000000000000000006000000000000000000000000000000000000000000000000051ef38b821e8000000000000000000000000000000000000000000000000000051ebff283cb65ba00000000000000000000000000000000000000000000000000000000670d39ab00000000000000000000000000000000000000000000000000000000670d8f1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009d3976c25f4defe584ed80bae5a7cef59ba07aa5000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000000";

        logIdx = 1;
        receiptProof.push(
            hex"f891a0022a3a4141798ccc5b3d76565985c487f80d33c6d84c50681cacee53b7a1178aa05ea06d2cf663de8c05b78213a955c261d00cadd02dc60e5705cf8d7a6301fa7ea0a2cc6ca974174d203b891af5e00478e5f0515a8d0d436bd364958c89ee2b7a9c8080808080a05fa8cd42f989c3e3f4284ed784084e0f76ae0177e0d9efcace1dc6581819a53b8080808080808080"
        );
        receiptProof.push(
            hex"f90211a0c6d0a721cb31a7891a8fe6523fc8584fccbfdb50be81c0fd01514a6957ed7a1ba0d1a3b1372c1f53ffc668d023c054a227f737d21c5bb6342b30b77768d49aee07a0b4bfeb55218638db2e47dd18adee34ae9e7671790b50cfa0c25e4c91d7063308a0258e54f2e2230e65b0c3bcf0e975c708fa47d0268a6edceb0a8c322a266ddae0a0941a0ad1b44e45c8172a64f4e3e5e19bf47ec9e658515a294e17a52cb06277e2a03ad24c4887b623835a1028075321979578f6f940c442f3a173dd3967d5265bd7a07d1c077612de47ddf1a529ba5121d37db2859152562c57200687e37ad824ca8da0c99732203600dd41c8c247fa02eb7fb36e07627e900503e18dfaa29e653ff41fa0f3fb8db5579d90247fe06e13eb005191c9f33a6c1b418b6158200e3ebb5d0b0da0fcb23b65537675bc4eda2665ff7c58a34674dbaa50a98a6fe96156eb3767789da0250c1b48df8b7f41248441bec7f15e52f4463d0e8f9ae2f309fc10055ed0bfeca08673eba5254e10ddc189815e9e6d2c39d34f7095f40799271422ac796d74171ca02ea80e05b1398bd61f3b2702821dc6bd1f23b421ee943f02feec5fd66f36ae84a0709c2c5a5cbc93c6aff26d96dd867fa7ec57cf6e7daaf39c0918a74fca7b296ba0492df7ac62065ddfdac36054c9dec14539b908a4be45388f6df89aa441df559ba05b66c22bbabcaf4b91e812bce87f23b5d4b2299c5418c6c6e82a4d38d3faad8780"
        );
        receiptProof.push(
            hex"f9039020b9038c02f9038801834008fdb9010000000000000001000000000000000000000000000000000000040000000001000000000000000000000000100000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000008000000000010000000000000400000000000000000000000200000000000000000002000000000000000040000000000000000000000001002100000000000020000000000000000000000000000000010000000000400000000000000000000000000000000000000000000000000000000000000000180000200000000000000080000020000002000000000000000000000000000000000000400000000000000000f9027df87a944200000000000000000000000000000000000006f842a0e1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109ca00000000000000000000000006f26bf09b1c792e3228e5467807a900a503c0281a0000000000000000000000000000000000000000000000000051ef38b821e8000f901fe946f26bf09b1c792e3228e5467807a900a503c0281f884a0a123dc29aebf7d0c3322c8eeb5b999e859f39937950ed31056532713d0de396fa00000000000000000000000000000000000000000000000000000000000002105a00000000000000000000000000000000000000000000000000000000000203bfba00000000000000000000000009d3976c25f4defe584ed80bae5a7cef59ba07aa5b9016000000000000000000000000042000000000000000000000000000000000000060000000000000000000000004200000000000000000000000000000000000006000000000000000000000000000000000000000000000000051ef38b821e8000000000000000000000000000000000000000000000000000051ebff283cb65ba00000000000000000000000000000000000000000000000000000000670d39ab00000000000000000000000000000000000000000000000000000000670d8f1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009d3976c25f4defe584ed80bae5a7cef59ba07aa5000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000000"
        );
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
        vm.skip(true);

        string memory input = vm.readFile(string.concat(rootDir, "/test/payload/packet_ack_proof.hex"));
        bytes memory encoded = vm.parseBytes(input);
        (, Ics23Proof memory peptideProof) = abi.decode(encoded, (bytes32, Ics23Proof));

        // proof will be struct EventProof
        bytes memory proof = abi.encode(peptideProof, receiptProof, receiptRoot, peptideClientId, peptideBlockNumber);

        crossProver.validateReceipt(receiptIdx, rlpEncodedReceipt, proof);
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
        bytes memory proof = abi.encode(peptideProof, receiptProof, receiptRoot, peptideClientId, peptideBlockNumber);

        vm.expectRevert();
        crossProver.validateReceipt(receiptIdx, rlpEncodedReceipt, proof);
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
}
