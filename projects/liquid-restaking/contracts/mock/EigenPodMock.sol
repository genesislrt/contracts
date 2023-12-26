// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/Create2.sol";
import "../interfaces/IEigenPod.sol";
import "hardhat/console.sol";

contract EigenPodMock is Initializable, IEigenPod {
    address public podOwner;
    bool public hasRestaked;

    function initialize(address _podOwner) external initializer {
        require(
            _podOwner != address(0),
            "EigenPod.initialize: podOwner cannot be zero address"
        );
        podOwner = _podOwner;
        hasRestaked = true;
        emit RestakingActivated(podOwner);
    }

    function MAX_RESTAKED_BALANCE_GWEI_PER_VALIDATOR()
        external
        view
        override
        returns (uint64)
    {}

    function withdrawableRestakedExecutionLayerGwei()
        external
        view
        override
        returns (uint64)
    {}

    function nonBeaconChainETHBalanceWei()
        external
        view
        override
        returns (uint256)
    {}

    function eigenPodManager()
        external
        view
        override
        returns (IEigenPodManager)
    {}

    function mostRecentWithdrawalTimestamp()
        external
        view
        override
        returns (uint64)
    {}

    function validatorPubkeyHashToInfo(
        bytes32 validatorPubkeyHash
    ) external view override returns (ValidatorInfo memory) {}

    function provenWithdrawal(
        bytes32 validatorPubkeyHash,
        uint64 slot
    ) external view override returns (bool) {}

    function validatorStatus(
        bytes32 pubkeyHash
    ) external view override returns (VALIDATOR_STATUS) {}

    function verifyWithdrawalCredentials(
        uint64 oracleTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        uint40[] calldata validatorIndices,
        bytes[] calldata withdrawalCredentialProofs,
        bytes32[][] calldata validatorFields
    ) external override {}

    function verifyBalanceUpdates(
        uint64 oracleTimestamp,
        uint40[] calldata validatorIndices,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        BeaconChainProofs.BalanceUpdateProof[] calldata balanceUpdateProofs,
        bytes32[][] calldata validatorFields
    ) external override {}

    function verifyAndProcessWithdrawals(
        uint64 oracleTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        BeaconChainProofs.WithdrawalProof[] calldata withdrawalProofs,
        bytes[] calldata validatorFieldsProofs,
        bytes32[][] calldata validatorFields,
        bytes32[][] calldata withdrawalFields
    ) external override {}

    function activateRestaking() external override {}

    function withdrawBeforeRestaking() external override {}

    function withdrawNonBeaconChainETHBalanceWei(
        address recipient,
        uint256 amountToWithdraw
    ) external override {}

    function recoverTokens(
        IERC20[] memory tokenList,
        uint256[] memory amountsToWithdraw,
        address recipient
    ) external override {}
}
