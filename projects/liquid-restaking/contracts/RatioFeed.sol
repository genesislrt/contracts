// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "./interfaces/IRatioFeed.sol";
import "./interfaces/IProtocolConfig.sol";
import "./interfaces/IStakingConfig.sol";

contract RatioFeed is Initializable, IRatioFeed {
    IStakingConfig _stakingConfig;

    mapping(address => uint256) private _ratios;
    mapping(address => HistoricalRatios) public historicalRatios;

    uint32 public constant MAX_THRESHOLD = uint32(1e8); // 100000000

    /// @dev diff between the current ratio and a new one in %(0.000001 ... 100%)
    uint256 private _ratioThreshold;

    /// @dev use this instead of HistoricalRatios.lastUpdate to check for 12hr ratio update timeout
    mapping(address => uint256) private _ratioUpdates;

    /*******************************************************************************
                        CONSTRUCTOR
    *******************************************************************************/

    function initialize(IStakingConfig stakingConfig) public initializer {
        _stakingConfig = stakingConfig;
    }

    modifier onlyGovernance() virtual {
        if (msg.sender != _stakingConfig.getGovernance()) {
            revert OnlyGovernanceAllowed();
        }
        _;
    }

    modifier onlyOperator() virtual {
        if (msg.sender != _stakingConfig.getOperator()) {
            revert OnlyOperatorAllowed();
        }
        _;
    }

    /*******************************************************************************
                        WRITE FUNCTIONS
    *******************************************************************************/

    function updateRatio(
        address token,
        uint256 newRatio
    ) public override onlyOperator {
        if (_ratioThreshold == 0) {
            revert RatioThresholdNotSet();
        }

        uint256 lastUpdate = _ratioUpdates[token];
        uint256 oldRatio = _ratios[token];

        (bool valid, string memory reason) = _checkRatioRules(
            lastUpdate,
            newRatio,
            oldRatio
        );

        if (!valid) {
            revert RatioNotUpdated(reason);
        }

        _ratios[token] = newRatio;
        emit RatioUpdated(token, oldRatio, newRatio);

        _ratioUpdates[token] = uint40(block.timestamp);

        HistoricalRatios storage hisRatio = historicalRatios[token];
        if (block.timestamp - hisRatio.lastUpdate > 1 days - 1 minutes) {
            uint64 latestOffset = hisRatio.historicalRatios[0];
            hisRatio.historicalRatios[((latestOffset + 1) % 8) + 1] = uint64(
                newRatio
            );
            hisRatio.historicalRatios[0] = latestOffset + 1;
            hisRatio.lastUpdate = uint40(block.timestamp);
        }
    }

    function _checkRatioRules(
        uint256 lastUpdated,
        uint256 newRatio,
        uint256 oldRatio
    ) internal view returns (bool valid, string memory reason) {
        if (oldRatio == 0) {
            return (valid = true, reason);
        }

        if (block.timestamp - lastUpdated < 12 hours) {
            return (valid, reason = "ratio was updated less than 12 hours ago");
        }
        if (newRatio > oldRatio) {
            return (valid, reason = "new ratio cannot be greater than old");
        }
        uint256 threshold = (oldRatio * _ratioThreshold) / MAX_THRESHOLD;
        if (newRatio < oldRatio - threshold) {
            return (
                valid,
                reason = "new ratio too low, not in threshold range"
            );
        }

        return (valid = true, reason);
    }

    function repairRatio(
        address token,
        uint256 newRatio
    ) public onlyGovernance {
        if (newRatio > 1e18 || newRatio == 0) {
            revert RatioNotUpdated("not in range");
        }
        emit RatioUpdated(token, _ratios[token], newRatio);
        _ratios[token] = newRatio;
    }

    function setRatioThreshold(uint256 newValue) external onlyGovernance {
        _setRatioThreshold(newValue);
    }

    function _setRatioThreshold(uint256 value) internal {
        if (value >= MAX_THRESHOLD || value == 0) {
            revert RatioThresholdNotInRange();
        }
        emit RatioThresholdChanged(_ratioThreshold, value);
        _ratioThreshold = value;
    }

    /*******************************************************************************
                        READ FUNCTIONS
    *******************************************************************************/

    function ratioThreshold() external view returns (uint256) {
        return _ratioThreshold;
    }

    function getRatio(address token) public view override returns (uint256) {
        return _ratios[token];
    }

    /**
     * @notice Returns APR based on ratio changes for `day`s
     */
    function averagePercentageRate(
        address token,
        uint8 day
    ) external view returns (uint256) {
        require(day > 0 && day < 8, "day should be from 1 to 7");

        HistoricalRatios storage hisRatio = historicalRatios[token];
        uint64 latestOffset = hisRatio.historicalRatios[0];

        uint256 oldestRatio = hisRatio.historicalRatios[
            ((latestOffset - day) % 8) + 1
        ];
        uint256 newestRatio = hisRatio.historicalRatios[
            ((latestOffset) % 8) + 1
        ];

        if (oldestRatio < newestRatio) {
            return 0;
        }

        return
            ((oldestRatio - newestRatio) * 10 ** 20 * 365) /
            (oldestRatio * (day));
    }
}
