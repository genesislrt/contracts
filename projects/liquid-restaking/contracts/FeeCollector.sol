// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import "./interfaces/IFeeCollector.sol";
import "./interfaces/IProtocolConfig.sol";

/**
 * @title MEV & Tips fee recipient
 * @author GenesisLRT
 * @notice Contract receives EL (tips/MEV) rewards and send them to RestakingPool
 */
contract FeeCollector is ReentrancyGuardUpgradeable, IFeeCollector {
    uint16 public constant MAX_COMMISSION = uint16(1e4); // 100.00

    IProtocolConfig internal _config;
    uint16 public commission;

    modifier onlyGovernance() virtual {
        require(
            msg.sender == _config.getGovernance(),
            "FeeCollector: only governance allowed"
        );
        _;
    }

    modifier onlyOperator() {
        require(
            msg.sender == _config.getOperator(),
            "FeeCollector: only consensus allowed"
        );
        _;
    }

    /*******************************************************************************
                        CONSTRUCTOR
    *******************************************************************************/

    function initialize(
        IProtocolConfig config,
        uint16 commission_
    ) public initializer {
        __ReentrancyGuard_init();
        _config = config;
        __FeeCollector_init(commission_);
    }

    function __FeeCollector_init(uint16 commission_) internal onlyInitializing {
        _setCommission(commission_);
    }

    /*******************************************************************************
                        WRITE FUNCTIONS
    *******************************************************************************/

    receive() external payable {
        emit Received(msg.sender, msg.value);
    }

    /**
     * @notice withdraw collected rewards to pool and treasury
     */
    function withdraw() external override nonReentrant {
        uint256 balance = address(this).balance;
        // min balance to withdraw is max commission
        if (balance >= MAX_COMMISSION) {
            (uint256 fee, uint256 rewardsWithoutCommission) = _takeFee(balance);

            address pool = address(_config.getRestakingPool());
            address treasury = _config.getTreasury();

            (bool success, ) = payable(pool).call{
                value: rewardsWithoutCommission
            }("");
            if (!success) {
                revert FeeCollectorTransferFailed(pool);
            }

            (success, ) = payable(treasury).call{value: fee}("");
            if (!success) {
                revert FeeCollectorTransferFailed(treasury);
            }

            emit Withdrawn(pool, treasury, rewardsWithoutCommission, fee);
        }
    }

    /*******************************************************************************
                        VIEW FUNCTIONS
    *******************************************************************************/

    /**
     * @notice get collected pool rewards w/ fee
     */
    function getRewards() external view returns (uint256 rewards) {
        (, rewards) = _takeFee(address(this).balance);
    }

    function _takeFee(
        uint256 amount
    ) internal view returns (uint256 fee, uint256 rewards) {
        fee = (amount * commission) / MAX_COMMISSION;
        rewards = amount - fee;
    }

    /*******************************************************************************
                        GOVERNANCE FUNCTIONS
    *******************************************************************************/

    function setCommission(uint16 newValue) external onlyGovernance {
        _setCommission(newValue);
    }

    function _setCommission(uint16 value) internal {
        if (value >= MAX_COMMISSION) {
            revert CommissionNotInRange();
        }

        emit CommissionChanged(commission, value);
        commission = value;
    }
}
