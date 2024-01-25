// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IRestakingPool {
    /* structs */

    struct Unstake {
        address recipient;
        uint256 amount;
    }

    /* errors */

    error OnlyGovernanceAllowed();
    error OnlyOperatorAllowed();

    error PoolZeroAmount();
    error PoolZeroAddress();
    error PoolRestakerExists();
    error PoolRestakerNotExists();
    error PoolInsufficientBalance();
    error PoolWrongInputLength();

    /**
     * @dev A call to an address target failed. The target may have reverted.
     */
    error PoolFailedInnerCall();

    error PoolDistributeGasLimitNotInRange(uint64 max);
    error PoolDistributeGasLimitNotSet();

    error PoolStakeAmLessThanMin();
    error PoolUnstakeAmLessThanMin();

    /* events */

    event Received(address indexed sender, uint256 amount);

    event Staked(address indexed staker, uint256 amount, uint256 shares);

    event Unstaked(
        address indexed from,
        address indexed to,
        uint256 amount,
        uint256 shares
    );

    event Deposited(string indexed provider, bytes[] pubkeys);

    event DistributeGasLimitChanged(uint32 prevValue, uint32 newValue);

    event PendingUnstake(
        address indexed ownerAddress,
        address indexed receiverAddress,
        uint256 amount,
        uint256 shares
    );

    /**
     * @dev Deprecated.
     */
    event RewardsDistributed(address[] claimers, uint256[] amounts);

    event UnstakesDistributed(Unstake[] unstakes);

    event ClaimExpected(address indexed claimer, uint256 value);

    event UnstakeClaimed(
        address indexed claimer,
        address indexed caller,
        uint256 value
    );

    event FeeClaimed(address indexed treasury, uint256 amount);

    /* functions */

    function getMinStake() external view returns (uint256);

    function getMinUnstake() external view returns (uint256);
}
