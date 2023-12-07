// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

interface IStakingPool {
    event Staked(address indexed staker, uint256 amount, uint256 shares);

    event PoolOnGoing(bytes pubkey);

    event DistributeGasLimitChanged(uint256 prevValue, uint256 newValue);

    event PendingUnstake(
        address indexed ownerAddress,
        address indexed receiverAddress,
        uint256 amount,
        uint256 shares
    );

    event Received(address indexed from, uint256 value);

    function stakeCerts() external payable;

    function getFreeBalance() external view returns (uint256);

    function getMinStake() external view returns (uint256);

    function getMinUnstake() external view returns (uint256);

    function pushToBeaconMulti(
        bytes[] calldata pubkeys,
        bytes[] calldata signatures,
        bytes32[] calldata deposit_data_roots
    ) external;
}
