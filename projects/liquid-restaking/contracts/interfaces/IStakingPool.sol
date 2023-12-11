// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./IRestakingPool.sol";

interface IStakingPool is IRestakingPool {
    event PoolOnGoing(bytes pubkey);

    event DistributeGasLimitChanged(uint256 prevValue, uint256 newValue);

    function stakeCerts() external payable;

    function getMinStake() external view returns (uint256);

    function getMinUnstake() external view returns (uint256);
}
