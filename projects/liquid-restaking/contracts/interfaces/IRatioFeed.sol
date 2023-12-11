// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IRatioFeed {
    /* errors */

    error OnlyGovernanceAllowed();
    error OnlyOperatorAllowed();

    error RatioThresholdNotSet();
    error RatioNotUpdated(string reason);
    error RatioThresholdNotInRange();

    struct HistoricalRatios {
        uint64[9] historicalRatios;
        uint40 lastUpdate;
    }

    event RatioThresholdChanged(uint256 oldValue, uint256 newValue);
    event RatioUpdated(
        address indexed tokenAddress,
        uint256 oldRatio,
        uint256 newRatio
    );

    function updateRatio(address token, uint256 ratio) external;
    function getRatio(address token) external view returns (uint256 ratio);
}
