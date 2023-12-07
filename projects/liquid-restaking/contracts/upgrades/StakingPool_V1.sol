// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import "../interfaces/ICertificateToken.sol";
import "../interfaces/IStakingConfig.sol";
import "../interfaces/IStakingPool.sol";
import "../interfaces/IEigenPodManager.sol";

contract StakingPool_V1 is
    IStakingPool,
    Initializable,
    ReentrancyGuardUpgradeable
{
    /**
     * @dev external contracts
     */
    IStakingConfig internal _stakingConfig;

    uint256 internal _DISTRIBUTE_GAS_LIMIT;

    uint256 internal _pendingGap;

    uint256 internal _pendingTotalUnstakes;
    address[] internal _pendingClaimers;
    mapping(address => uint256) internal _pendingClaimerUnstakes;

    uint256[] internal _pendingRequests;

    // reserve some gap for the future upgrades
    uint256[50 - 5] private __reserved;

    event RewardsDistributed(address[] claimers, uint256[] amounts);

    modifier onlyGovernance() virtual {
        require(
            msg.sender == _stakingConfig.getGovernanceAddress(),
            "StakingPool: only governance allowed"
        );
        _;
    }

    modifier onlyOperator() {
        require(
            msg.sender == _stakingConfig.getOperatorAddress(),
            "StakingPool: only consensus allowed"
        );
        _;
    }

    function initialize(
        IStakingConfig stakingConfig,
        uint256 distributeGasLimit
    ) external initializer {
        _stakingConfig = stakingConfig;
        IEigenPodManager(stakingConfig.getEigenPodManagerAddress()).createPod();
        __QueuePool_init(distributeGasLimit);
    }

    function __QueuePool_init(uint256 distributeGasLimit) internal {
        _DISTRIBUTE_GAS_LIMIT = distributeGasLimit;
        emit DistributeGasLimitChanged(0, distributeGasLimit);
    }

    function stakeCerts() external payable override nonReentrant {
        uint256 amount = msg.value;
        require(
            amount >= _stakingConfig.getMinStake(),
            "StakingPool: value must be greater than min amount"
        );
        ICertificateToken certificateToken = ICertificateToken(
            _stakingConfig.getCertTokenAddress()
        );
        uint256 shares = certificateToken.bondsToShares(amount);
        certificateToken.mint(msg.sender, shares);
        emit Staked(msg.sender, amount, shares);
    }

    function pushToBeaconMulti(
        bytes[] calldata pubkeys,
        bytes[] calldata signatures,
        bytes32[] calldata deposit_data_roots
    ) public onlyOperator {
        uint256 pubkeysLen = pubkeys.length;
        require(
            pubkeysLen == signatures.length &&
                signatures.length == deposit_data_roots.length,
            "StakingPool: length are not equal"
        );
        require(
            address(this).balance >= 32 ether * pubkeysLen,
            "pending ethers not enough"
        );
        IEigenPodManager eigenPodManager = IEigenPodManager(
            _stakingConfig.getEigenPodManagerAddress()
        );
        for (uint i = 0; i < pubkeysLen; i++) {
            eigenPodManager.stake(
                pubkeys[i],
                signatures[i],
                deposit_data_roots[i]
            );
            emit PoolOnGoing(pubkeys[i]);
        }
    }

    function pushToBeacon(
        bytes calldata pubkey,
        bytes calldata signature,
        bytes32 deposit_data_root
    ) public onlyOperator {
        require(address(this).balance >= 32 ether, "pending ethers not enough");
        IEigenPodManager(_stakingConfig.getEigenPodManagerAddress()).stake(
            pubkey,
            signature,
            deposit_data_root
        );
        emit PoolOnGoing(pubkey);
    }

    function getFreeBalance() public view virtual returns (uint256) {
        return address(this).balance;
    }

    /**
     * @notice Burns amount of certificate from msg.sender
     * @notice Returns native token immediately or via queue
     * @param receiverAddress address for receiving unstaked funds
     * @param shares amount of certificate token to unstake
     */
    function unstakeCerts(address receiverAddress, uint256 shares) external {
        address ownerAddress = msg.sender;
        ICertificateToken certificateToken = ICertificateToken(
            _stakingConfig.getCertTokenAddress()
        );

        uint256 amount = certificateToken.sharesToBonds(shares);
        require(
            amount >= _stakingConfig.getMinUnstake(),
            "StakingPool: value must be greater than min amount"
        );

        require(
            certificateToken.balanceOf(ownerAddress) >= shares,
            "StakingPool: cannot unstake more than have on address"
        );
        certificateToken.burn(ownerAddress, shares);
        _addIntoQueue(ownerAddress, receiverAddress, shares, amount);
    }

    /**
     * @return Certificate token address
     */
    function getCert() external view virtual returns (address) {
        return _stakingConfig.getCertTokenAddress();
    }

    function getEigenPodManager() external view virtual returns (address) {
        return _stakingConfig.getEigenPodManagerAddress();
    }

    function getMinStake() external view virtual returns (uint256) {
        return _stakingConfig.getMinStake();
    }

    function getMinUnstake() public view virtual override returns (uint256) {
        return _stakingConfig.getMinUnstake();
    }

    function setDistributeGasLimit(uint256 newValue) external onlyGovernance {
        uint256 prevValue = _DISTRIBUTE_GAS_LIMIT;
        _DISTRIBUTE_GAS_LIMIT = newValue;

        emit DistributeGasLimitChanged(prevValue, newValue);
    }

    function getDistributeGasLimit() public view returns (uint256) {
        return _DISTRIBUTE_GAS_LIMIT;
    }

    function _addIntoQueue(
        address owner,
        address claimer,
        uint256 shares,
        uint256 amount
    ) internal {
        require(
            amount != 0 && claimer != address(0),
            "LiquidTokenStakingPool: zero input values"
        );
        // each new request is placed at the end of the queue
        _pendingTotalUnstakes += amount;
        _pendingClaimers.push(claimer);
        _pendingRequests.push(amount);
        _pendingClaimerUnstakes[claimer] += amount;
        emit PendingUnstake(owner, claimer, amount, shares);
    }

    function getTotalPendingUnstakes() public view returns (uint256) {
        return _pendingTotalUnstakes;
    }

    function getPendingRequestsOf(
        address claimer
    ) public view returns (uint256[] memory) {
        uint256 j;
        uint256[] memory unstakes = new uint256[](
            _pendingClaimers.length - _pendingGap
        );
        for (uint256 i = _pendingGap; i < _pendingClaimers.length; i++) {
            if (_pendingClaimers[i] == claimer) {
                unstakes[j] = _pendingRequests[i];
                ++j;
            }
        }
        uint256 removeCells = unstakes.length - j;
        if (removeCells > 0) {
            assembly {
                mstore(unstakes, j)
            }
        }
        return unstakes;
    }

    function getPendingUnstakesOf(
        address claimer
    ) public view returns (uint256) {
        return _pendingClaimerUnstakes[claimer];
    }

    function distributePendingRewards() external {
        require(
            _DISTRIBUTE_GAS_LIMIT > 0,
            "StakingPool: DISTRIBUTE_GAS_LIMIT is not set"
        );
        uint256 poolBalance = getFreeBalance();
        address[] memory claimers = new address[](
            _pendingClaimers.length - _pendingGap
        );
        uint256[] memory amounts = new uint256[](
            _pendingClaimers.length - _pendingGap
        );
        uint256 j = 0;
        uint256 i = _pendingGap;

        while (
            i < _pendingClaimers.length &&
            poolBalance > 0 &&
            gasleft() > _DISTRIBUTE_GAS_LIMIT
        ) {
            address claimer = _pendingClaimers[i];
            uint256 toDistribute = _pendingRequests[i];
            if (claimer == address(0) || toDistribute == 0) {
                ++i;
                continue;
            }

            if (poolBalance < toDistribute) {
                break;
            }

            _pendingClaimerUnstakes[claimer] -= toDistribute;
            _pendingTotalUnstakes -= toDistribute;
            poolBalance -= toDistribute;
            delete _pendingClaimers[i];
            delete _pendingRequests[i];
            ++i;

            bool success = _unsafeTransfer(claimer, toDistribute, true);
            if (!success) {
                continue;
            }
            claimers[j] = claimer;
            amounts[j] = toDistribute;
            ++j;
        }
        _pendingGap = i;
        /* decrease arrays */
        uint256 removeCells = claimers.length - j;
        if (removeCells > 0) {
            assembly {
                mstore(claimers, j)
            }
            assembly {
                mstore(amounts, j)
            }
        }

        emit RewardsDistributed(claimers, amounts);
    }

    /**
     * @dev Unsafe transfer with gas limit if necessary
     * @notice The solution was received from eth bounty program
     */
    function _unsafeTransfer(
        address receiverAddress,
        uint256 amount,
        bool limit
    ) internal virtual returns (bool) {
        address payable wallet = payable(receiverAddress);
        bool success;
        if (limit) {
            assembly {
                success := call(10000, wallet, amount, 0, 0, 0, 0)
            }
            return success;
        }
        (success, ) = wallet.call{value: amount}("");
        return success;
    }
}
