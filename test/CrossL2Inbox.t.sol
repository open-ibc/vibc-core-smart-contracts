// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "forge-std/Test.sol";
import "./utils/CrossL2Inbox.base.t.sol";

contract CrossL2InboxTest is CrossL2InboxBase {
    function test_validate_message_success() public {
        bool valid = cross.validateMessageWithProof(receiptProof, rawLog);
        require(valid, "validate message should return successful");
    }
}
