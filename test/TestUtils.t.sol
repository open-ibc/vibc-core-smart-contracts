// SPDX-License-Identifier: UNLICENSED
import {IDispatcher} from "../contracts/interfaces/IDispatcher.sol";
import {IUniversalChannelHandler} from "../contracts/interfaces/IUniversalChannelHandler.sol";
import {LightClient} from "../contracts/interfaces/LightClient.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {Dispatcher} from "../contracts/core/Dispatcher.sol";
import {UniversalChannelHandler} from "../contracts/core/UniversalChannelHandler.sol";

import {Test} from "forge-std/Test.sol";
import {Vm} from "forge-std/Vm.sol";

pragma solidity ^0.8.0;

abstract contract TestUtilsTest {
    bytes32 internal constant _ADMIN_SLOT = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
    bytes32 private constant _ROLLBACK_SLOT = 0x4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143;
    bytes32 internal constant _IMPLEMENTATION_SLOT = 0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc;

    function deployDispatcherProxyAndImpl(string memory initPortPrefix, LightClient lightClient)
        public
        returns (IDispatcher proxy, Dispatcher dispatcherImplementation)
    {
        dispatcherImplementation = new Dispatcher();
        proxy = IDispatcher(
            address(
                new ERC1967Proxy(
                    address(dispatcherImplementation),
                    abi.encodeWithSelector(Dispatcher.initialize.selector, initPortPrefix, lightClient)
                )
            )
        );
    }

    function deployUCHProxyAndImpl(address dispatcherProxy)
        public
        returns (IUniversalChannelHandler proxy, UniversalChannelHandler uchImplementation)
    {
        uchImplementation = new UniversalChannelHandler();
        proxy = IUniversalChannelHandler(
            address(
                new ERC1967Proxy(
                    address(uchImplementation),
                    abi.encodeWithSelector(UniversalChannelHandler.initialize.selector, dispatcherProxy)
                )
            )
        );
    }

    function getProxyImplementation(address proxy, Vm vm) public view returns (address dispatcherImplementation) {
        dispatcherImplementation = address(uint160(uint256(vm.load(address(proxy), _IMPLEMENTATION_SLOT))));
    }
}
