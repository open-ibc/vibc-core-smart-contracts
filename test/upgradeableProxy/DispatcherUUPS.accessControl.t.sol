// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../../contracts/libs/Ibc.sol";
import {Dispatcher} from "../../contracts/core/Dispatcher.sol";
import {IbcEventsEmitter} from "../../contracts/interfaces/IbcDispatcher.sol";
import {IbcReceiver} from "../../contracts/interfaces/IbcReceiver.sol";
import "../../contracts/examples/Mars.sol";
import "../../contracts/core/OptimisticLightClient.sol";
import "../utils/Dispatcher.base.t.sol";
import {DispatcherV2} from "./upgrades/DispatcherV2.sol";
import {DispatcherV2Initializable} from "./upgrades/DispatcherV2Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {DummyLightClient} from "../../contracts/utils/DummyLightClient.sol";

contract DispatcherUUPSAccessControl is Base {
    string public portPrefix2 = "IIpolyibc.eth.";
    ILightClient lightClient2 = new DummyLightClient();
    address public notOwner = vm.addr(1);
    DispatcherV2 dispatcherImplementation2;
    DispatcherV2Initializable dispatcherImplementation3;

    function setUp() public override {
        (dispatcherProxy, dispatcherImplementation) = deployDispatcherProxyAndImpl(portPrefix);
        dispatcherProxy.setClientForConnection(connectionHops[0], dummyLightClient);
        dispatcherImplementation2 = new DispatcherV2();
        dispatcherImplementation3 = new DispatcherV2Initializable();
    }

    function test_Dispatcher_Allows_Upgrade_123() public {
        assertEq(address(dispatcherImplementation), getProxyImplementation(address(dispatcherProxy), vm));
        UUPSUpgradeable(address(dispatcherProxy)).upgradeTo(address(dispatcherImplementation2));
        assertEq(address(dispatcherImplementation2), getProxyImplementation(address(dispatcherProxy), vm));
        assertEq(dispatcherProxy.portPrefix(), portPrefix);
    }

    function test_Dispatcher_Allows_Upgrade_To_And_Call() public {
        assertEq(address(dispatcherImplementation), getProxyImplementation(address(dispatcherProxy), vm));
        bytes memory initData = abi.encodeWithSignature("initialize(string)", portPrefix2);
        UUPSUpgradeable(address(dispatcherProxy)).upgradeToAndCall(address(dispatcherImplementation3), initData);
        assertEq(address(dispatcherImplementation3), getProxyImplementation(address(dispatcherProxy), vm));
        assertEq(dispatcherProxy.portPrefix(), portPrefix2);
    }

    function test_Dispatcher_Prevents_Non_Owner_Updgrade() public {
        vm.startPrank(notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        UUPSUpgradeable(address(dispatcherProxy)).upgradeTo(address(dispatcherImplementation2));
    }

    function test_Dispatcher_Prevents_Non_Owner_UpdgradeToAndCall() public {
        vm.expectRevert("Function must be called through delegatecall");
        UUPSUpgradeable(address(dispatcherImplementation)).upgradeTo(address(dispatcherImplementation2));
    }

    function test_Dispatcher_Prevents_Reinit_Attacks() public {
        vm.expectRevert("Initializable: contract is already initialized");
        dispatcherImplementation.initialize("IIpolyibc.eth.");
    }
}
