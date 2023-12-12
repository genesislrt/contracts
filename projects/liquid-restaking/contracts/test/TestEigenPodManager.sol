// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "../interfaces/IEigenPodManager.sol";

contract TestEigenPodManager is IEigenPodManager {
    function stake(
        bytes calldata pubkey,
        bytes calldata signature,
        bytes32 depositDataRoot
    ) external payable override {}

    function ownerToPod(
        address podOwner
    ) external view override returns (IEigenPod) {}

    function getPod(
        address podOwner
    ) external view override returns (IEigenPod) {}

    function ethPOS() external view override returns (IETHPOSDeposit) {}

    function eigenPodBeacon() external view override returns (IBeacon) {}

    function beaconChainOracle()
        external
        view
        override
        returns (IBeaconChainOracle)
    {}

    function getBlockRootAtTimestamp(
        uint64 timestamp
    ) external view override returns (bytes32) {}

    function strategyManager()
        external
        view
        override
        returns (IStrategyManager)
    {}

    function slasher() external view override returns (ISlasher) {}

    function hasPod(address podOwner) external view override returns (bool) {}

    function numPods() external view override returns (uint256) {}

    function maxPods() external view override returns (uint256) {}

    function podOwnerShares(
        address podOwner
    ) external view override returns (int256) {}

    function beaconChainETHStrategy()
        external
        view
        override
        returns (IStrategy)
    {}

    function createPod() external override returns (address) {}
}
