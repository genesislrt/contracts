// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

/**
 * @dev Interface of the cToken.
 */
interface ICToken is IERC20Upgradeable {
    /* errors */

    error OnlyGovernanceAllowed();
    error OnlyOperatorAllowed();
    error OnlyRestakingPoolAllowed();

    /* functions */

    function convertToShares(
        uint256 amount
    ) external view returns (uint256 shares);

    function convertToAmount(
        uint256 shares
    ) external view returns (uint256 amount);

    function mint(address account, uint256 amount) external;

    function burn(address account, uint256 amount) external;

    function ratio() external view returns (uint256);
}
