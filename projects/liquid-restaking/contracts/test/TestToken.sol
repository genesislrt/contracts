// SPDX-License-Identifier: MIT
pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract TestToken is Context, ERC20 {
    constructor() ERC20('Test', 'TST'){}
}
