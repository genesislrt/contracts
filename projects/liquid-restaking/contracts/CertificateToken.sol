// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";

import "./libs/MathUtils.sol";

import "./interfaces/ICertificateToken.sol";
import "./interfaces/IStakingConfig.sol";
import "./interfaces/IRatioFeed.sol";
import "./interfaces/IPausable.sol";
import "./interfaces/IStakingPool.sol";
import "./interfaces/IPausable.sol";


contract CertificateToken is
OwnableUpgradeable,
ERC20Upgradeable,
ICertificateToken,
IPausable
{
    // earn config
    IStakingConfig internal _stakingConfig;

    // pausable
    bool private _paused;

    // reserve some gap for the future upgrades
    uint256[100 - 4] private __reserved;

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
        require(
            msg.sender == _stakingConfig.getGovernanceAddress(),
            "CertificateToken: only governance allowed"
        );
        _;
    }

    modifier onlyLiquidStakingPool() virtual {
        require(
            msg.sender == _stakingConfig.getStakingPoolAddress(),
            "CertificateToken: only liquid staking pool"
        );
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

    function transfer(
        address to,
        uint256 amount
    )
    public
    virtual
    override(ERC20Upgradeable, IERC20Upgradeable)
    whenNotPaused
    returns (bool)
    {
        address ownerAddress = _msgSender();
        _transfer(ownerAddress, to, amount);
        return true;
    }

    function transferFrom(
        address from,
        address to,
        uint256 amount
    )
    public
    virtual
    override(ERC20Upgradeable, IERC20Upgradeable)
    whenNotPaused
    returns (bool)
    {
        address spender = _msgSender();
        _spendAllowance(from, spender, amount);
        _transfer(from, to, amount);
        return true;
    }

    function approve(
        address spender,
        uint256 amount
    )
    public
    virtual
    override(ERC20Upgradeable, IERC20Upgradeable)
    whenNotPaused
    returns (bool)
    {
        address ownerAddress = _msgSender();
        _approve(ownerAddress, spender, amount);
        return true;
    }

    function mint(
        address account,
        uint256 amount
    ) external override whenNotPaused onlyLiquidStakingPool {
        _mint(account, amount);
    }

    function burn(
        address account,
        uint256 amount
    ) external override whenNotPaused onlyLiquidStakingPool {
        _burn(account, amount);
    }

    function sharesToBonds(
        uint256 amount
    ) public view override returns (uint256) {
        return MathUtils.multiplyAndDivideCeil(amount, 1 ether, ratio());
    }

    function bondsToShares(
        uint256 amount
    ) public view override returns (uint256) {
        return MathUtils.multiplyAndDivideFloor(amount, ratio(), 1 ether);
    }

    function ratio() public view override returns (uint256) {
        address feed = _stakingConfig.getRatioFeedAddress();
        uint256 ratio = IRatioFeed(feed).getRatioFor(address(this));
        return ratio;
    }

    function isRebasing() external pure override returns (bool) {
        return false;
    }

    function paused() public view returns (bool) {
        return _paused;
    }

    function pause() external whenNotPaused onlyGovernance {
        _paused = true;
        emit Paused(_msgSender());
    }

    function unpause() external whenPaused onlyGovernance {
        _paused = false;
        emit Unpaused(_msgSender());
    }
}
