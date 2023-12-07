// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../interfaces/IEigenPodManager.sol";

contract TestEigenPodManager is IEigenPodManager {
    function createPod() external override {}

    function stake(
        bytes calldata pubkey,
        bytes calldata signature,
        bytes32 depositDataRoot
    ) external payable override {}
}
