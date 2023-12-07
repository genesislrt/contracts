// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

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

    modifier onlyGovernance() virtual {
        require(msg.sender == address(_governanceAddress), "StakingConfig: only governance");
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
        emit OperatorAddressChanged(address(0x00), operatorAddress);
        _certTokenAddress = certTokenAddress;
        emit CertTokenAddressChanged(address(0x00), certTokenAddress);
        _governanceAddress = governanceAddress;
        emit GovernanceAddressChanged(address(0x00), governanceAddress);
        _treasuryAddress = treasuryAddress;
        emit TreasuryAddressChanged(address(0x00), treasuryAddress);
        _stakingPoolAddress = stakingPoolAddress;
        emit StakingPoolAddressChanged(address(0x00), stakingPoolAddress);
        _eigenPodManagerAddress = eigenPodManager;
        emit EigenManagerAddressChanged(address(0x00), eigenPodManager);
        _ratioFeedAddress = ratioFeed;
        emit RatioFeedAddressChanged(address(0x00), ratioFeed);
    }

    function getOperatorAddress() external view override returns (address) {
        return _operatorAddress;
    }

    function setOperatorAddress(address newValue) external override onlyGovernance {
        require(newValue != address(0), "StakingConfig: address can't be nil");
        address prevValue = _operatorAddress;
        _operatorAddress = newValue;
        emit OperatorAddressChanged(prevValue, newValue);
    }

    function setGovernanceAddress(address newValue) external override onlyGovernance {
        require(newValue != address(0), "StakingConfig: address can't be nil");
        address prevValue = _governanceAddress;
        _governanceAddress = newValue;
        emit GovernanceAddressChanged(prevValue, newValue);
    }

    function setRatioFeedAddress(address newValue) external override onlyGovernance {
        require(newValue != address(0), "StakingConfig: address can't be nil");
        address prevValue = _ratioFeedAddress;
        _ratioFeedAddress = newValue;
        emit RatioFeedAddressChanged(prevValue, newValue);
    }

    function setTreasuryAddress(address newValue) external override onlyGovernance {
        require(newValue != address(0), "StakingConfig: address can't be nil");
        address prevValue = _treasuryAddress;
        _treasuryAddress = newValue;
        emit TreasuryAddressChanged(prevValue, newValue);
    }

    function setStakingPoolAddress(address newValue) external override onlyGovernance {
        require(newValue != address(0), "StakingConfig: address can't be nil");
        address prevValue = _stakingPoolAddress;
        _stakingPoolAddress = newValue;
        emit StakingPoolAddressChanged(prevValue, newValue);
    }


    function setCertTokenAddress(address newValue) external override onlyGovernance {
        require(newValue != address(0), "StakingConfig: address can't be nil");
        address prevValue = _certTokenAddress;
        _certTokenAddress = newValue;
        emit CertTokenAddressChanged(prevValue, newValue);
    }


    function getGovernanceAddress() external view override returns (address) {
        return _governanceAddress;
    }

    function getStakingPoolAddress() external view override returns (address) {
        return _stakingPoolAddress;
    }

    function getRatioFeedAddress() external view override returns (address) {
        return _ratioFeedAddress;
    }

    function getEigenPodManagerAddress() external view override returns (address) {
        return _eigenPodManagerAddress;
    }

    function getCertTokenAddress() external view override returns (address) {
        return _certTokenAddress;
    }

    function getTreasuryAddress() external view override returns (address) {
        return _treasuryAddress;
    }

    function getMinStake() public view virtual override returns (uint256) {
        return uint256(_minimumStake) * 1 gwei;
    }

    function getMinUnstake() public view virtual override returns (uint256) {
        return uint256(_minimumUnstake) * 1 gwei;
    }

    function setMinUnstake(
        uint256 newValue
    ) external virtual override onlyGovernance {
        _setMinUnstake(newValue);
    }

    function _setMinUnstake(
        uint256 newValue
    ) internal {
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


    function setMinStake(
        uint256 newValue
    ) external virtual override onlyGovernance {
        _setMinStake(newValue);
    }

    function _setMinStake(uint256 newValue) internal {
        require(
            newValue % 1 gwei == 0,
            "StakingConfig: value should be multiplied of gwei"
        );
        uint256 prevValue = getMinStake();
        _minimumStake = uint64(newValue / 1 gwei);
        require(
            _minimumStake * 1 gwei == newValue,
            "StakingConfig: overflow"
        );
        emit MinStakeChanged(prevValue, newValue);
    }
}