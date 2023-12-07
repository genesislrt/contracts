// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

interface IGovernable {
    event GovernanceAddressChanged(address prevValue, address newValue);

    function getGovernanceAddress() external view returns (address);
    function setGovernanceAddress(address newValue) external;
}
