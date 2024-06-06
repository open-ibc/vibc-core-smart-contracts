// SPDX-License-Identifier: Apache-2.0
/*
 * Copyright 2024, Polymer Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity 0.8.15;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract FeeVault is Ownable {
    error SenderNotDispatcher();
    error IncorrectFeeSent();

    event feeDeposited(address indexed from, uint256 ackFee, uint256 recvFee);

    address public dispatcher;

    modifier onlyDispatcher() {
        if (msg.sender != dispatcher) revert SenderNotDispatcher();
        _;
    }

    constructor(address _dispatcher) {
        dispatcher = _dispatcher;
    }

    function depositFee(uint256 ackFee, uint256 recvFee) external payable onlyDispatcher {
        if (ackFee + recvFee != msg.value) {
            revert IncorrectFeeSent();
        }

        emit feeDeposited(msg.sender, ackFee, recvFee);
    }

    function withdraw() external onlyOwner {
        payable(owner()).transfer(address(this).balance);
    }
}
