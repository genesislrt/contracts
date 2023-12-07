// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "./IGovernable.sol";

interface IStakingConfig is IGovernable {
    event OperatorAddressChanged(address prevValue, address newValue);
    event TreasuryAddressChanged(address prevValue, address newValue);
    event RatioFeedAddressChanged(address prevValue, address newValue);
    event CertTokenAddressChanged(address prevValue, address newValue);
    event StakingPoolAddressChanged(address prevValue, address newValue);
    event EigenManagerAddressChanged(address prevValue, address newValue);
    event MinUnstakeChanged(uint256 prevValue, uint256 newValue);
    event MinStakeChanged(uint256 prevValue, uint256 newValue);

    function getOperatorAddress() external view returns (address);

    function getRatioFeedAddress() external view returns (address);

    function getStakingPoolAddress() external view returns (address);

    function getTreasuryAddress() external view returns (address);

    function getEigenPodManagerAddress() external view returns (address);

    function getCertTokenAddress() external view returns (address);

    function getMinStake() external view returns (uint256);

    function getMinUnstake() external view returns (uint256);

    function setStakingPoolAddress(address newValue) external;

    function setOperatorAddress(address newValue) external;

    function setRatioFeedAddress(address newValue) external;

    function setTreasuryAddress(address newValue) external;

    function setCertTokenAddress(address newValue) external;

    function setMinStake(uint256 newValue) external;

    function setMinUnstake(uint256 newValue) external;
}
