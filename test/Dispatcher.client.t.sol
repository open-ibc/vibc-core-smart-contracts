// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../contracts/libs/Ibc.sol";
import {Dispatcher} from "../contracts/core/Dispatcher.sol";
import {IDispatcher} from "../contracts/interfaces/IDispatcher.sol";
import {IbcEventsEmitter} from "../contracts/interfaces/IbcDispatcher.sol";
import {IbcReceiver} from "../contracts/interfaces/IbcReceiver.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "../contracts/examples/Mars.sol";
import "../contracts/core/OpLightClient.sol";
import "./Dispatcher.base.t.sol";
import {TestUtils} from "./TestUtils.t.sol";

abstract contract DispatcherUpdateClientTestSuite is Base {
    function test_updateOptimisticConsensusState_success() public {
        // trick the L1Block contract into thinking it is updated with the right l1 header
        setL1BlockAttributes(keccak256(RLPWriter.writeList(l1header.header)), l1header.number);
        dispatcherProxy.updateClientWithOptimisticConsensusState(l1header, validStateProof, 1, uint256(apphash));
    }

    function test_updateOptimisticConsensusState_failure() public {
        setL1BlockAttributes(keccak256(RLPWriter.writeList(l1header.header)), l1header.number);
        vm.expectRevert("MerkleTrie: ran out of proof elements");
        dispatcherProxy.updateClientWithOptimisticConsensusState(l1header, invalidStateProof, 1, uint256(apphash));
    }
}

contract DispatcherUpdateClientTest is DispatcherUpdateClientTestSuite {
    function setUp() public virtual override {
        (dispatcherProxy, dispatcherImplementation) = TestUtils.deployDispatcherProxyAndImpl(portPrefix, opLightClient);
    }
}
