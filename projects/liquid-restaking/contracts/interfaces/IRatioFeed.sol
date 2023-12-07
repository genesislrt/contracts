// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

interface IRatioFeed {
    event OperatorAdded(address operator);
    event OperatorRemoved(address operator);
    event RatioThresholdChanged(uint256 oldValue, uint256 newValue);

    event RatioUpdated(
        address indexed tokenAddress,
        uint256 oldRatio,
        uint256 newRatio
    );
    event RatioNotUpdated(
        address indexed tokenAddress,
        uint256 failedRatio,
        string reason
    );

    function updateRatioBatch(
        address[] calldata addresses,
        uint256[] calldata ratios
    ) external;

    function getRatioFor(address) external view returns (uint256);
}
