// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./IProtocolConfig.sol";

interface IStakingConfig is IProtocolConfig {
    event EigenManagerAddressChanged(address prevValue, address newValue);
    event MinUnstakeChanged(uint256 prevValue, uint256 newValue);
    event MinStakeChanged(uint256 prevValue, uint256 newValue);

    function getRatioFeedAddress() external view returns (address);

    function getStakingPoolAddress() external view returns (address);

    function getEigenPodManagerAddress() external view returns (address);

    function getCertTokenAddress() external view returns (address);

    function getMinStake() external view returns (uint256);

    function getMinUnstake() external view returns (uint256);
}
