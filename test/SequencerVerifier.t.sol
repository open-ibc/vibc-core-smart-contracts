// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../contracts/core/OptimisticProofVerifier.sol";
import "../contracts/libs/Ibc.sol";
import {IbcUtils} from "../contracts/libs/IbcUtils.sol";
import "forge-std/Test.sol";
import "./utils/Signing.base.t.sol";

contract SequencerProofVerifierStateUpdate is SigningBase {
    function test_verify_signature_state_update_sucess() public view {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(sequencerPkey, hashToSign);

        bytes memory signature = abi.encodePacked(r, s, v); // Annoingly, v, r, s are in a different order than those
            // returned from vm.sign

        // This call should not revert as long as it's a valid sequencer signature
        sigVerifier.verifyStateUpdate(peptideBlockNumber, peptideAppHash, l1BlockHash, signature);
    }

    // Call should revert if invalid signature values are passed
    function test_verify_state_update_invalid_signature() public {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(sequencerPkey, hashToSign);

        bytes memory signatureTooShort = abi.encodePacked(r, s);
        bytes memory signatureTooLong = abi.encodePacked(r, s, v, uint256(123));

        // This call should not revert as long as it's a valid sequencer signature
        vm.expectRevert("ECDSA: invalid signature length");
        sigVerifier.verifyStateUpdate(peptideBlockNumber, peptideAppHash, l1BlockHash, signatureTooShort);
        vm.expectRevert("ECDSA: invalid signature length");
        sigVerifier.verifyStateUpdate(peptideBlockNumber, peptideAppHash, l1BlockHash, signatureTooLong);
    }

    // Valid signature but from the wrong signer should revert
    function test_verify_state_update_not_signer() public {
        (, uint256 notSequencerPkey) = makeAddrAndKey(unicode"😳");
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(notSequencerPkey, hashToSign);

        bytes memory wrongSignerSignature = abi.encodePacked(r, s, v);

        vm.expectRevert(abi.encodeWithSelector(ISignatureVerifier.InvalidSequencerSignature.selector));
        sigVerifier.verifyStateUpdate(peptideBlockNumber, peptideAppHash, l1BlockHash, wrongSignerSignature);
    }

    // Valid signature but on the wrong header should revert
    function test_verify_state_update_incorrect_header() public {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(sequencerPkey, hashToSign);

        bytes memory signature = abi.encodePacked(r, s, v);

        vm.expectRevert(abi.encodeWithSelector(ISignatureVerifier.InvalidSequencerSignature.selector));
        sigVerifier.verifyStateUpdate(peptideBlockNumber + 1, peptideAppHash, l1BlockHash, signature);
    }
}

// Note: Membership verification tests are done in verifier.t.sol, as it is part of the AppStateVerifier abstract
// contract
