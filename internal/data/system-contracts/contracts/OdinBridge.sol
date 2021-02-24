// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.7.0;

contract OdinBridge {
    constructor() {}

    event UserDeposited(address indexed _user, uint256 _depositAmount, string _odinAddress);

    function deposit(string memory _odinAddress) external payable returns (bool) {
        require(msg.value > 0, "Invalid value for the deposit amount, failed to deposit a zero value.");

        emit UserDeposited(msg.sender, msg.value, _odinAddress);
        return true;
    }
}
