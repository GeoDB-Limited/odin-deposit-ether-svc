// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.7.0;

import "./AddressStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";


contract Bridge is Ownable {
    AddressStorage supportedTokens;

    event EtherDeposited(
        address indexed _userAddress,
        string _odinAddress,
        uint256 _depositAmount
    );
    event TokenDeposited(
        address indexed _userAddress,
        string _odinAddress,
        address indexed _tokenAddress,
        uint256 _depositAmount
    );
    event TokenAdded(address indexed _tokenAddress);
    event TokenRemoved(address indexed _tokenAddress);

    constructor(address[] memory _supportedTokens) {
        supportedTokens = new AddressStorage(_supportedTokens);
    }

    /**
    * @notice Deposits ether
    * @param _odinAddress Address in the Odin chain
    * @return True if everything went well
    */
    function depositEther(string memory _odinAddress) external payable returns (bool) {
        require(msg.value > 0, "Invalid value for the deposit amount, failed to deposit a zero value.");
        emit EtherDeposited(msg.sender, _odinAddress, msg.value);
        return true;
    }

    /**
    * @notice Deposits ERC20 compatible tokens
    * @param _tokenAddress Address of the ERC20 compatible token contract
    * @param _odinAddress Address in the Odin chain
    * @param _depositAmount Amount to deposit
    * @return True if everything went well
    */
    function depositToken(address _tokenAddress, string memory _odinAddress, uint256 _depositAmount)
    external returns (bool)
    {
        require(supportedTokens.contains(_tokenAddress), "Unsupported token, failed to deposit.");

        bool _ok = IERC20(_tokenAddress).transferFrom(msg.sender, address(this), _depositAmount);
        require(_ok, "Failed to transfer tokens.");

        emit TokenDeposited(msg.sender, _odinAddress, _tokenAddress, _depositAmount);
        return true;
    }

    /**
    * @notice Adds ERC20 compatible token to supported tokens
    * @param _tokenAddress Address of the ERC20 compatible token contract
    * @return True if everything went well
    */
    function addToken(address _tokenAddress) onlyOwner() external returns (bool) {
        supportedTokens.mustAdd(_tokenAddress);
        emit TokenAdded(_tokenAddress);
        return true;
    }

    /**
    * @notice Removes ERC20 compatible token from supported tokens
    * @param _tokenAddress Address of the ERC20 compatible token contract
    * @return True if everything went well
    */
    function removeToken(address _tokenAddress) onlyOwner() external returns (bool) {
        supportedTokens.mustRemove(_tokenAddress);
        emit TokenRemoved(_tokenAddress);
        return true;
    }
}
