// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.6.0 <0.8.0;

contract EtherBridge {
    event EtherDeposited(
        address indexed _userAddress,
        string _odinAddress,
        uint256 _depositAmount
    );

    function depositEther(string memory _odinAddress) external payable returns (bool) {
        require(msg.value > 0, "Invalid value for the deposit amount, failed to deposit a zero value.");
        emit EtherDeposited(msg.sender, _odinAddress, msg.value);
        return true;
    }
}
