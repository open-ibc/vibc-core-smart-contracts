// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "../../contracts/libs/Ibc.sol";
import "../../contracts/interfaces/IProofVerifier.sol";
import "../../contracts/core/SequencerSignatureVerifier.sol";
import {RLPWriter} from "optimism/libraries/rlp/RLPWriter.sol";
import {L1Block} from "optimism/L2/L1Block.sol";

contract SigningBase is Test {
    using stdJson for string;

    string rootDir = vm.projectRoot();

    bytes32 l1BlockHash;
    bytes32 l1AncestorBlockHash;
    bytes32 peptideAppHash;
    uint256 public peptideBlockNumber;
    uint256 public sequencerPkey;
    uint256 public notSequencerPkey;
    address public sequencer;
    bytes32 hashToSign;
    bytes32 domain; // Domain will be empty so we can leave it as initialized to default 0x 32 bytes

    SequencerSignatureVerifier public sigVerifier;
    bytes32 PEPTIDE_CHAIN_ID = bytes32(uint256(444));

    L1Header childl1Block; // Child block, represents the l1 origin of dest chain when peptide catches up to ancestor L1
        // block
    L1Header ancestorL1Block; // Ancestor block, represents the l1 origin of dest chain when peptide wants to submit a

    // client update but is behind, so we need to checkpoint this block
    L1Block l1BlockProvider = new L1Block();

    constructor() {
        (sequencer, sequencerPkey) = makeAddrAndKey("alice");
        (, notSequencerPkey) = makeAddrAndKey(unicode"bob😈");
        sigVerifier = new SequencerSignatureVerifier(sequencer, PEPTIDE_CHAIN_ID);

        // generate the channel_proof.hex file with the following command:
        // cd test-data-generator && go run ./cmd/ --type l1 > ../test/payload/l1_block_0x4df537.hex
        // this is the "rlp" half-encoded header that would be sent by the relayer. this was produced
        // by the test-data-generator tool.
        L1Header memory l1header = abi.decode(
            vm.parseBytes(vm.readFile(string.concat(rootDir, "/test/payload/l1_block_0x4df537.hex"))), (L1Header)
        );

        l1BlockHash = keccak256(RLPWriter.writeList(l1header.header)); // Blockhash that will be signed by sequencer

        // Set l1 child block hash which is used to verify signature against
        l1AncestorBlockHash = readBytes32FromJson("/test/payload/l1_block_ancestor.json", ".result.hash");

        peptideBlockNumber = 101;

        // this happens to be the polymer height when the L2OO was updated with the output proposal
        // we are using in the test
        string memory l2BlockJson = vm.readFile(string.concat(rootDir, "/test/payload/l2_block_0x4b0.json"));
        peptideAppHash = abi.decode(l2BlockJson.parseRaw(".result.stateRoot"), (bytes32));

        bytes32 payloadHash = keccak256(abi.encodePacked(peptideBlockNumber, peptideAppHash, l1BlockHash));
        hashToSign = keccak256(bytes.concat(domain, PEPTIDE_CHAIN_ID, payloadHash));
    }

    // Read a specific bytes32 json property from a json file
    function readBytes32FromJson(string memory fileName, string memory property) public view returns (bytes32) {
        string memory blockJson = vm.readFile(string.concat(rootDir, fileName));
        return abi.decode(blockJson.parseRaw(property), (bytes32));
    }

    // Read a specific uint64 json property from a json file
    function readUint64FromJson(string memory fileName, string memory property) public view returns (uint64) {
        string memory blockJson = vm.readFile(string.concat(rootDir, fileName));
        return abi.decode(blockJson.parseRaw(property), (uint64));
    }

    // Set the block attributes on the l1 block info to that from a json file
    function setBlockAttributesFromJson(string memory fileName) public {
        bytes32 hash = readBytes32FromJson(fileName, ".result.hash");
        uint64 number = readUint64FromJson(fileName, ".result.number");
        setL1BlockAttributes(hash, number);
    }

    function setL1BlockAttributes(bytes32 hash, uint64 number) public {
        vm.prank(l1BlockProvider.DEPOSITOR_ACCOUNT());
        l1BlockProvider.setL1BlockValues(
            number,
            0, //          timestamp
            0, //          basefee
            hash,
            0, //          sequenceNumber
            bytes32(0), // batcherHash
            0, //          l1FeeOverhead
            0 //           l1FeeScalar
        );
    }
}
