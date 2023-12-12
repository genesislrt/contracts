// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/math/MathUpgradeable.sol";

import "./interfaces/ICToken.sol";
import "./interfaces/IPausable.sol";
import "./interfaces/IStakingConfig.sol";

contract CertificateToken is
    OwnableUpgradeable,
    ERC20Upgradeable,
    ICToken,
    IPausable
{
    using MathUpgradeable for uint256;

    // earn config
    IStakingConfig internal _stakingConfig;

    // pausable
    bool private _paused;

    // reserve some gap for the future upgrades
    uint256[100 - 4] private __reserved;

    /*******************************************************************************
                        CONSTRUCTOR
    *******************************************************************************/

    function initialize(
        IStakingConfig stakingConfig,
        string memory name,
        string memory symbol
    ) external initializer {
        __Ownable_init();
        __ERC20_init(name, symbol);
        __CertificateToken_init(stakingConfig);
    }

    function __CertificateToken_init(IStakingConfig stakingConfig) internal {
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

    modifier onlyRestakingPool() virtual {
        if (_msgSender() != address(_stakingConfig.getRestakingPool())) {
            revert OnlyRestakingPoolAllowed();
        }
        _;
    }

    modifier whenNotPaused() virtual {
        require(!paused(), "CertificateToken: paused");
        _;
    }

    modifier whenPaused() virtual {
        require(paused(), "CertificateToken: not paused");
        _;
    }

    /*******************************************************************************
                        WRITE FUNCTIONS
    *******************************************************************************/

    function _transfer(
        address from,
        address to,
        uint256 amount
    ) internal override whenNotPaused {
        super._transfer(from, to, amount);
    }

    /**
     * @notice Mints exactly `shares` to `account`.
     * - emit the Transfer event from zero address.
     */
    function mint(
        address account,
        uint256 shares
    ) external override whenNotPaused onlyRestakingPool {
        _mint(account, shares);
    }

    /**
     * @notice Burns exactly `shares` from `account`.
     * - emit the Transfer event to zero address.
     * - revert if all of shares cannot be burned (due to the owner not having enough shares).
     */
    function burn(
        address account,
        uint256 shares
    ) external override whenNotPaused onlyRestakingPool {
        _burn(account, shares);
    }

    /*******************************************************************************
                        READ FUNCTIONS
    *******************************************************************************/

    /**
     * @dev Deprecated.
     */
    function sharesToBonds(uint256 shares) external view returns (uint256) {
        return convertToAmount(shares);
    }

    /**
     * @notice Returns the `amount` of ETH that the cToken would exchange for the amount of `shares` provided, in an ideal
     * scenario where all the conditions are met.
     */
    function convertToAmount(
        uint256 shares
    ) public view override returns (uint256) {
        return shares.mulDiv(1 ether, ratio(), MathUpgradeable.Rounding.Up);
    }

    /**
     * @dev Deprecated.
     */
    function bondsToShares(uint256 amount) external view returns (uint256) {
        return convertToShares(amount);
    }

    /**
     * @notice Returns the amount of `shares` that the cToken would exchange for the `amount` of ETH provided, in an ideal
     * scenario where all the conditions are met.
     */
    function convertToShares(
        uint256 amount
    ) public view override returns (uint256) {
        return amount.mulDiv(ratio(), 1 ether, MathUpgradeable.Rounding.Down);
    }

    function ratio() public view override returns (uint256) {
        return _stakingConfig.getRatioFeed().getRatio(address(this));
    }

    function isRebasing() external pure returns (bool) {
        return false;
    }

    /**
     * @dev Returns the total amount of the ETH that is “managed” by Genesis.
     */
    function totalAssets() external view returns (uint256 totalManagedEth) {
        return convertToAmount(totalSupply());
    }

    function paused() public view returns (bool) {
        return _paused;
    }

    /*******************************************************************************
                        GOVERNANCE FUNCTIONS
    *******************************************************************************/

    function pause() external whenNotPaused onlyGovernance {
        _paused = true;
        emit Paused(_msgSender());
    }

    function unpause() external whenPaused onlyGovernance {
        _paused = false;
        emit Unpaused(_msgSender());
    }
}
