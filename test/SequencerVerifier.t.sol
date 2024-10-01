// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../contracts/core/OptimisticProofVerifier.sol";
import "../contracts/libs/Ibc.sol";
import {IbcUtils} from "../contracts/libs/IbcUtils.sol";
import "forge-std/Test.sol";
import "./utils/SigningClient.base.t.sol";
import "../contracts/core/SequencerSoloClient.sol";
import "../contracts/core/SequencerSignatureVerifier.sol";

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
    // Valid signature but on the wrong header should revert

    function test_verify_state_update_real_data() public {
        address verifier = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;
        bytes32 peptideChainId = 0x0000000000000000000000000000000000000000000000000000000000000387;
        SequencerSignatureVerifier localVerifier =
            new SequencerSignatureVerifier(0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266, peptideChainId);
        SequencerSoloClient sc = new SequencerSoloClient(localVerifier);
        sc.updateClient(
            hex"32b62ebe733d3974f718e759a140a9656a1d474c334ecf36897868a970405133a21fd238da9de5fadb44765be29c4f65a41a48edbd41d42b2e3da39bd13d42e9423f297a0c73c0cb84c8135d4b0a7fafc6f08685a83ac19a9766d7f9bc19fe261c",
            uint256(133),
            uint256(
                29_131_683_131_418_862_121_097_353_289_911_432_830_775_670_159_092_440_431_404_531_363_980_419_005_200
            )
        );
    }
}

// Note: Membership verification tests are done in verifier.t.sol, as it is part of the AppStateVerifier abstract
// contract
