// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "./interfaces/IRatioFeed.sol";
import "./interfaces/IStakingConfig.sol";

contract RatioFeed is
Initializable,
IRatioFeed
{
    IStakingConfig _stakingConfig;

    struct HistoricalRatios {
        uint64[9] historicalRatios;
        uint40 lastUpdate;
    }

    mapping(address => uint256) private _ratios;
    mapping(address => HistoricalRatios) public historicalRatios;

    uint32 public constant MAX_THRESHOLD = uint32(1e8); // 100000000

    /// @dev diff between the current ratio and a new one in %(0.000001 ... 100%)
    uint256 private _ratioThreshold;

    /// @dev use this instead of HistoricalRatios.lastUpdate to check for 12hr ratio update timeout
    mapping(address => uint256) private _ratioUpdates;

    function initialize(IStakingConfig stakingConfig) public initializer {
        _stakingConfig = stakingConfig;
    }

    modifier onlyGovernance() virtual {
        require(
            msg.sender == _stakingConfig.getGovernanceAddress(),
            "RatioFeed: only governance allowed"
        );
        _;
    }

    modifier onlyOperator() {
        require(
            msg.sender == _stakingConfig.getOperatorAddress(),
            "RatioFeed: only operator allowed"
        );
        _;
    }

    function updateRatioBatch(
        address[] calldata addresses,
        uint256[] calldata ratios
    ) public override onlyOperator {
        require(addresses.length == ratios.length, "corrupted ratio data");
        require(_ratioThreshold > 0, "ratio threshold is not set");

        for (uint256 i = 0; i < addresses.length; i++) {
            address tokenAddr = addresses[i];
            uint256 lastUpdate = _ratioUpdates[tokenAddr];
            uint256 oldRatio = _ratios[tokenAddr];
            uint256 newRatio = ratios[i];

            (bool valid, string memory reason) = _checkRatioRules(
                lastUpdate,
                newRatio,
                oldRatio
            );

            if (!valid) {
                emit RatioNotUpdated(tokenAddr, newRatio, reason);
                // continue to other ratios
                continue;
            }

            _ratios[tokenAddr] = newRatio;
            emit RatioUpdated(tokenAddr, oldRatio, newRatio);

            _ratioUpdates[tokenAddr] = uint40(block.timestamp);

            // let's compare with a new ratio
            HistoricalRatios storage hisRatio = historicalRatios[tokenAddr];
            if (block.timestamp - hisRatio.lastUpdate > 1 days - 1 minutes) {
                uint64 latestOffset = hisRatio.historicalRatios[0];
                hisRatio.historicalRatios[
                ((latestOffset + 1) % 8) + 1
                ] = uint64(newRatio);
                hisRatio.historicalRatios[0] = latestOffset + 1;
                hisRatio.lastUpdate = uint40(block.timestamp);
            }
        }
    }

    function getRatioThreshold() public view returns (uint256) {
        return _ratioThreshold;
    }

    function _checkRatioRules(
        uint256 lastUpdated,
        uint256 newRatio,
        uint256 oldRatio
    ) internal view returns (bool valid, string memory reason) {
        // initialization of the first ratio -> skip checks
        if (oldRatio == 0) {
            return (valid = true, reason);
        }

        if (block.timestamp - lastUpdated < 12 hours) {
            // valid == false
            return (valid, reason = "ratio was updated less than 12 hours ago");
        }
        // new ratio should be not greater than a previous one
        if (newRatio > oldRatio) {
            // valid == false
            return (valid, reason = "new ratio cannot be greater than old");
        }
        // new ratio should be in the range (oldRatio - threshold , oldRatio]
        uint256 threshold = (oldRatio * _ratioThreshold) / MAX_THRESHOLD;
        if (newRatio < oldRatio - threshold) {
            // valid == false
            return (valid, reason = "new ratio too low, not in threshold range");
        }

        return (valid = true, reason);
    }

    function averagePercentageRate(
        address addr,
        uint256 day
    ) external view returns (uint256) {
        require(day > 0 && day < 8, "day should be from 1 to 7");

        HistoricalRatios storage hisRatio = historicalRatios[addr];
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

    function repairRatioFor(address token, uint256 ratio) public onlyGovernance {
        require(ratio != 0, "ratio is zero");
        uint256 oldRatio = _ratios[token];
        _ratios[token] = ratio;
        emit RatioUpdated(token, oldRatio, ratio);
    }

    function getRatioFor(address token) public view override returns (uint256) {
        return _ratios[token];
    }

    function setRatioThreshold(uint256 newValue) public onlyGovernance {
        require(
            newValue < MAX_THRESHOLD && newValue > 0,
            "wrong value for ratio threshold"
        );
        uint256 oldValue = _ratioThreshold;
        _ratioThreshold = newValue;
        emit RatioThresholdChanged(oldValue, newValue);
    }
}
