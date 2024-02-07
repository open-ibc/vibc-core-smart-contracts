// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import '../contracts/Ibc.sol';
import {Dispatcher} from '../contracts/Dispatcher.sol';
import {IbcEventsEmitter} from '../contracts/IbcDispatcher.sol';
import {IbcReceiver} from '../contracts/IbcReceiver.sol';
import '../contracts/Mars.sol';
import '../contracts/OpConsensusStateManager.sol';
import './Dispatcher.base.t.sol';

contract ClientTestBase is Base {
    function setUp() public virtual override {
        super.setUp();
        dispatcher = new Dispatcher(portPrefix, opConsensusStateManager);
    }
}

contract DispatcherUpdateClientTest is ClientTestBase {
    function setUp() public override {
        super.setUp();
    }

    function test_updateOptimisticConsensusState_success() public {
        // trick the L1Block contract into thinking it is updated with the right l1 header
        setL1BlockAttributes(keccak256(RLPWriter.writeList(l1header.header)), l1header.number);
        dispatcher.updateClientWithOptimisticConsensusState(l1header, validStateProof, 1, uint256(apphash));
    }

    function test_updateOptimisticConsensusState_failure() public {
        setL1BlockAttributes(keccak256(RLPWriter.writeList(l1header.header)), l1header.number);
        vm.expectRevert('MerkleTrie: ran out of proof elements');
        dispatcher.updateClientWithOptimisticConsensusState(l1header, invalidStateProof, 1, uint256(apphash));
    }
}
