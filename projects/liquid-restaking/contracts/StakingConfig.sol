// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "./interfaces/IStakingConfig.sol";

contract StakingConfig is Initializable, IStakingConfig {
    /**
     * @dev compact size, should be multiplied by 1 gwei
     */
    uint64 internal _minimumUnstake;
    uint64 internal _minimumStake;

    address internal _operatorAddress;
    address internal _governanceAddress;
    address internal _treasuryAddress;
    address internal _ratioFeedAddress;
    address internal _stakingPoolAddress;
    address internal _eigenPodManagerAddress;
    address internal _certTokenAddress;
    IRestakerDeployer internal _restakerDeployer;

    modifier onlyGovernance() virtual {
        require(
            msg.sender == address(_governanceAddress),
            "StakingConfig: only governance"
        );
        _;
    }

    function initialize(
        uint256 minimumUnstake,
        uint256 minimumStake,
        address operatorAddress,
        address governanceAddress,
        address treasuryAddress,
        address stakingPoolAddress,
        address certTokenAddress,
        address ratioFeed,
        address eigenPodManager
    ) external initializer {
        _setMinStake(minimumStake);
        _setMinUnstake(minimumUnstake);
        _operatorAddress = operatorAddress;
        emit OperatorChanged(address(0x00), operatorAddress);
        _certTokenAddress = certTokenAddress;
        emit CTokenChanged(ICToken(address(0)), ICToken(certTokenAddress));
        _governanceAddress = governanceAddress;
        emit GovernanceChanged(address(0), governanceAddress);
        _treasuryAddress = treasuryAddress;
        emit TreasuryChanged(address(0), treasuryAddress);
        _stakingPoolAddress = stakingPoolAddress;
        emit RestakingPoolChanged(
            IRestakingPool(address(0)),
            IRestakingPool(stakingPoolAddress)
        );
        _eigenPodManagerAddress = eigenPodManager;
        emit EigenManagerAddressChanged(address(0), eigenPodManager);
        _ratioFeedAddress = ratioFeed;
        emit RatioFeedChanged(IRatioFeed(address(0)), IRatioFeed(ratioFeed));
    }

    function getOperator() external view override returns (address) {
        return _operatorAddress;
    }

    function setOperatorAddress(address newValue) external onlyGovernance {
        require(newValue != address(0), "StakingConfig: address can't be nil");
        address prevValue = _operatorAddress;
        _operatorAddress = newValue;
        emit OperatorChanged(prevValue, newValue);
    }

    function setGovernanceAddress(address newValue) external onlyGovernance {
        require(newValue != address(0), "StakingConfig: address can't be nil");
        address prevValue = _governanceAddress;
        _governanceAddress = newValue;
        emit GovernanceChanged(prevValue, newValue);
    }

    function setRatioFeed(address newValue) external onlyGovernance {
        require(newValue != address(0), "StakingConfig: address can't be nil");
        address prevValue = _ratioFeedAddress;
        _ratioFeedAddress = newValue;
        emit RatioFeedChanged(IRatioFeed(prevValue), IRatioFeed(newValue));
    }

    function setTreasury(address newValue) external onlyGovernance {
        require(newValue != address(0), "StakingConfig: address can't be nil");
        address prevValue = _treasuryAddress;
        _treasuryAddress = newValue;
        emit TreasuryChanged(prevValue, newValue);
    }

    function setRestakingPool(address newValue) external onlyGovernance {
        require(newValue != address(0), "StakingConfig: address can't be nil");
        address prevValue = _stakingPoolAddress;
        _stakingPoolAddress = newValue;
        emit RestakingPoolChanged(
            IRestakingPool(prevValue),
            IRestakingPool(newValue)
        );
    }

    function setCToken(address newValue) external onlyGovernance {
        require(newValue != address(0), "StakingConfig: address can't be nil");
        address prevValue = _certTokenAddress;
        _certTokenAddress = newValue;
        emit CTokenChanged(ICToken(prevValue), ICToken(newValue));
    }

    function setRestakerDeployer(
        IRestakerDeployer newValue
    ) external onlyGovernance {
        require(
            address(newValue) != address(0),
            "StakingConfig: address can't be nil"
        );
        emit RestakerDeployerChanged(_restakerDeployer, newValue);
        _restakerDeployer = newValue;
    }

    function getGovernance() external view override returns (address) {
        return _governanceAddress;
    }

    function getStakingPoolAddress() public view override returns (address) {
        return _stakingPoolAddress;
    }

    function getRatioFeedAddress() public view override returns (address) {
        return _ratioFeedAddress;
    }

    function getEigenPodManagerAddress()
        external
        view
        override
        returns (address)
    {
        return _eigenPodManagerAddress;
    }

    function getCertTokenAddress() public view override returns (address) {
        return _certTokenAddress;
    }

    function getTreasury() external view override returns (address) {
        return _treasuryAddress;
    }

    function getMinStake() public view virtual override returns (uint256) {
        return uint256(_minimumStake) * 1 gwei;
    }

    function getMinUnstake() public view virtual override returns (uint256) {
        return uint256(_minimumUnstake) * 1 gwei;
    }

    function setMinUnstake(uint256 newValue) external virtual onlyGovernance {
        _setMinUnstake(newValue);
    }

    function _setMinUnstake(uint256 newValue) internal {
        require(
            newValue % 1 gwei == 0,
            "StakingConfig: value should be multiplied of gwei"
        );
        uint256 prevValue = getMinUnstake();
        _minimumUnstake = uint64(newValue / 1 gwei);
        require(
            _minimumUnstake * 1 gwei == newValue,
            "StakingConfig: overflow"
        );
        emit MinUnstakeChanged(prevValue, newValue);
    }

    function setMinStake(uint256 newValue) external virtual onlyGovernance {
        _setMinStake(newValue);
    }

    function _setMinStake(uint256 newValue) internal {
        require(
            newValue % 1 gwei == 0,
            "StakingConfig: value should be multiplied of gwei"
        );
        uint256 prevValue = getMinStake();
        _minimumStake = uint64(newValue / 1 gwei);
        require(_minimumStake * 1 gwei == newValue, "StakingConfig: overflow");
        emit MinStakeChanged(prevValue, newValue);
    }

    function getCToken() external view override returns (ICToken token) {
        return ICToken(getCertTokenAddress());
    }

    function getRatioFeed() external view override returns (IRatioFeed feed) {
        return IRatioFeed(getRatioFeedAddress());
    }

    function getRestakingPool()
        external
        view
        override
        returns (IRestakingPool pool)
    {
        return IRestakingPool(getStakingPoolAddress());
    }

    function getEigenPodManager()
        external
        view
        override
        returns (IEigenPodManager manager)
    {}

    function getRestakerDeployer()
        external
        view
        override
        returns (IRestakerDeployer deployer)
    {
        return _restakerDeployer;
    }
}
