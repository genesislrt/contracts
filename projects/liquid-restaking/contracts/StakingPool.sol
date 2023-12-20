// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

import "./interfaces/ICToken.sol";
import "./interfaces/IEigenPodManager.sol";
import "./interfaces/IStakingConfig.sol";
import "./interfaces/IStakingPool.sol";

import "hardhat/console.sol";

contract StakingPool is
    IStakingPool,
    Initializable,
    ReentrancyGuardUpgradeable
{
    // @dev block gas limit
    uint64 internal constant MAX_GAS_LIMIT = 30_000_000;

    // @dev max gas allocated for {_sendValue}
    uint256 public constant CALL_GAS_LIMIT = 10_000;

    /**
     * @dev external contracts
     */
    IStakingConfig internal _stakingConfig;

    uint256 internal _distributeGasLimit;

    // @dev Current gap of {_pendingUnstakes}.
    uint256 internal _pendingGap;

    uint256 internal _pendingTotalUnstakes;
    address[] internal _pendingClaimers;
    mapping(address => uint256) internal _pendingClaimerUnstakes;

    uint256[] internal _pendingRequests;

    uint256 internal _totalClaimable;
    mapping(address => uint256) internal _claimable;
    // keccak256(provider) => Restaker
    mapping(bytes32 => address) internal _restakers;

    // reserve some gap for the future upgrades
    uint256[50 - 8] private __reserved;

    modifier onlyGovernance() virtual {
        require(
            msg.sender == _stakingConfig.getGovernance(),
            "StakingPool: only governance allowed"
        );
        _;
    }

    modifier onlyOperator() {
        require(
            msg.sender == _stakingConfig.getOperator(),
            "StakingPool: only consensus allowed"
        );
        _;
    }

    function initialize(
        IStakingConfig stakingConfig,
        uint256 distributeGasLimit
    ) external initializer {
        _stakingConfig = stakingConfig;
        IEigenPodManager(stakingConfig.getEigenPodManager()).createPod();
        __QueuePool_init(distributeGasLimit);
    }

    function __QueuePool_init(uint256 distributeGasLimit) internal {
        setDistributeGasLimit(distributeGasLimit);
    }

    function stake() external payable {
        stakeCerts();
    }

    /**
     * @dev Deprecated.
     */
    function stakeCerts() public payable override {
        uint256 amount = msg.value;
        require(
            amount >= _stakingConfig.getMinStake(),
            "StakingPool: value must be greater than min amount"
        );
        ICToken certificateToken = _stakingConfig.getCToken();
        uint256 shares = certificateToken.convertToShares(amount);
        certificateToken.mint(msg.sender, shares);
        emit Staked(msg.sender, amount, shares);
    }

    function unstake(address to, uint256 shares) external {
        unstakeCerts(to, shares);
    }

    function batchDeposit(
        string memory provider,
        bytes[] calldata pubkeys,
        bytes[] calldata signatures,
        bytes32[] calldata deposit_data_roots
    ) external onlyOperator nonReentrant {
        uint256 pubkeysLen = pubkeys.length;

        if (
            pubkeysLen != signatures.length ||
            pubkeysLen != deposit_data_roots.length
        ) {
            revert PoolWrongInputLength();
        }
        if (getPending() < 32 ether * pubkeysLen) {
            revert PoolInsufficientBalance();
        }

        address restaker = _restakers[_getProviderHash(provider)];
        if (restaker == address(0)) {
            revert PoolRestakerNotExists();
        }

        for (uint i; i < pubkeysLen; i++) {
            IEigenPodManager(restaker).stake{value: 32 ether}(
                pubkeys[i],
                signatures[i],
                deposit_data_roots[i]
            );
        }

        emit Deposited(provider, pubkeys);
    }

    /**
     * @dev Deprecated.
     */
    function unstakeCerts(address receiverAddress, uint256 shares) public {
        address ownerAddress = msg.sender;
        ICToken certificateToken = _stakingConfig.getCToken();

        uint256 amount = certificateToken.convertToAmount(shares);
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

    function distributeUnstakes() external nonReentrant {
        require(
            _distributeGasLimit > 0,
            "StakingPool: DISTRIBUTE_GAS_LIMIT is not set"
        );
        uint256 poolBalance = getPending();
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
            gasleft() > _distributeGasLimit
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

            bool success = _sendValue(claimer, toDistribute, true);
            if (!success) {
                _addClaimable(claimer, toDistribute);
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

    function _sendValue(
        address recipient,
        uint256 amount,
        bool limit
    ) internal returns (bool success) {
        if (address(this).balance < amount) {
            revert PoolInsufficientBalance();
        }

        address payable wallet = payable(recipient);
        if (limit) {
            assembly {
                success := call(CALL_GAS_LIMIT, wallet, amount, 0, 0, 0, 0)
            }
        } else {
            (success, ) = wallet.call{value: amount}("");
        }

        return success;
    }

    function _addClaimable(address account, uint256 amount) internal {
        _totalClaimable += amount;
        _claimable[account] += amount;
        emit ClaimExpected(account, amount);
    }

    function claimUnstake(address claimer) external nonReentrant {
        if (claimer == address(0)) {
            revert PoolZeroAddress();
        }

        uint256 amount = claimableOf(claimer);

        if (amount == 0) {
            revert PoolZeroAmount();
        }

        if (address(this).balance < getTotalClaimable()) {
            revert PoolInsufficientBalance();
        }
        _totalClaimable -= amount;
        _claimable[claimer] = 0;

        bool result = _sendValue(claimer, amount, false);
        if (!result) {
            revert PoolFailedInnerCall();
        }

        emit UnstakeClaimed(claimer, msg.sender, amount);
    }

    /*******************************************************************************
                        EIGEN POD OWNER WRITE FUNCTIONS
    *******************************************************************************/

    // @dev will be called only once for each restaker, because it activates restaking.
    function activateRestaking(string memory provider) external onlyOperator {
        address restaker = _getRestakerOrRevert(provider);
        // it withdraw ETH to restaker
        IEigenPod(restaker).activateRestaking();
        // claim withdrawn ETH to pool
        IRestaker(restaker).__claim();
    }

    // @dev withdraw not restaked ETH
    function withdrawBeforeRestaking(
        string memory provider
    ) external onlyOperator {
        address restaker = _getRestakerOrRevert(provider);
        // it withdraw ETH to restaker
        IEigenPod(restaker).withdrawBeforeRestaking();
        // claim withdrawn ETH to pool
        IRestaker(restaker).__claim();
    }

    function verifyWithdrawalCredentials(
        string memory provider,
        uint64 oracleTimestamp,
        BeaconChainProofs.StateRootProof calldata stateRootProof,
        uint40[] calldata validatorIndices,
        bytes[] calldata validatorFieldsProofs,
        bytes32[][] calldata validatorFields
    ) external onlyOperator {
        IEigenPod restaker = IEigenPod(_getRestakerOrRevert(provider));
        restaker.verifyWithdrawalCredentials(
            oracleTimestamp,
            stateRootProof,
            validatorIndices,
            validatorFieldsProofs,
            validatorFields
        );
    }

    function withdrawNonBeaconChainETHBalanceWei(
        string memory provider,
        uint256 amountToWithdraw
    ) external {
        IEigenPod restaker = IEigenPod(_getRestakerOrRevert(provider));
        restaker.withdrawNonBeaconChainETHBalanceWei(
            address(this),
            amountToWithdraw
        );
    }

    function recoverTokens(
        string memory provider,
        IERC20[] memory tokenList,
        uint256[] memory amountsToWithdraw
    ) external {
        IEigenPod restaker = IEigenPod(_getRestakerOrRevert(provider));
        restaker.recoverTokens(
            tokenList,
            amountsToWithdraw,
            _stakingConfig.getOperator()
        );
    }

    /*******************************************************************************
                        VIEW FUNCTIONS
    *******************************************************************************/

    function getMinStake() public view virtual returns (uint256 amount) {
        return _stakingConfig.getMinStake();
    }

    function getMinUnstake()
        public
        view
        virtual
        override
        returns (uint256 shares)
    {
        return _stakingConfig.getMinUnstake();
    }

    function getPending() public view returns (uint256) {
        uint256 balance = address(this).balance;
        uint256 claimable = getTotalClaimable();

        if (claimable > balance) {
            return 0;
        } else {
            return balance - claimable;
        }
    }

    function getTotalClaimable() public view returns (uint256) {
        return _totalClaimable;
    }

    /**
     * @dev Deprecated.
     */
    function getFreeBalance() external view virtual returns (uint256) {
        return getPending();
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

    function hasClaimable(address claimer) public view returns (bool) {
        return _claimable[claimer] != uint256(0);
    }

    function claimableOf(address claimer) public view returns (uint256) {
        return _claimable[claimer];
    }

    function _getRestakerOrRevert(
        string memory provider
    ) internal view returns (address restaker) {
        restaker = _restakers[_getProviderHash(provider)];
        if (restaker == address(0)) {
            revert PoolRestakerNotExists();
        }
    }

    function _getProviderHash(
        string memory providerName
    ) internal pure returns (bytes32) {
        return keccak256(bytes(providerName));
    }

    /*******************************************************************************
                        GOVERNANCE FUNCTIONS
    *******************************************************************************/

    function addRestaker(string memory provider) external onlyGovernance {
        bytes32 providerHash = _getProviderHash(provider);
        address restaker = _restakers[providerHash];
        if (restaker != address(0)) {
            revert PoolRestakerExists();
        }
        restaker = address(
            _stakingConfig.getRestakerDeployer().deployRestaker()
        );
        _restakers[providerHash] = restaker;
        emit RestakerAdded(provider, restaker);
    }

    function setDistributeGasLimit(uint256 newValue) public onlyGovernance {
        uint256 prevValue = _distributeGasLimit;
        _distributeGasLimit = newValue;

        emit DistributeGasLimitChanged(uint32(prevValue), uint32(newValue));
    }
}
