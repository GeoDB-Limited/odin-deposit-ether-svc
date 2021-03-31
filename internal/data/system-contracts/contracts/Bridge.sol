// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";


contract Bridge is Ownable {
    using Address for address;

    mapping(address => bool) public supportedTokens;

    event ETHDeposited(
        address indexed _userAddress,
        string _odinAddress,
        uint256 _depositAmount
    );
    event ERC20Deposited(
        address indexed _userAddress,
        string _odinAddress,
        address indexed _tokenAddress,
        uint256 _depositAmount
    );
    event TokenAdded(address indexed _tokenAddress);
    event TokenRemoved(address indexed _tokenAddress);

    constructor(address[] memory _supportedTokens) {
        for (uint256 i = 0; i < _supportedTokens.length; i++) {
            supportedTokens[_supportedTokens[i]] = true;
        }
    }

    /**
    * @notice Deposits ETH
    * @param _odinAddress Address in the Odin chain
    * @return True if everything went well
    */
    function depositEther(string memory _odinAddress) external payable returns (bool) {
        require(msg.value > 0, "Invalid value for the deposit amount, failed to deposit a zero value.");
        emit ETHDeposited(msg.sender, _odinAddress, msg.value);
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
        require(_tokenAddress.isContract(), "Given token is not a contract");
        require(supportedTokens[_tokenAddress], "Unsupported token, failed to deposit.");

        bool _success = IERC20(_tokenAddress).transferFrom(msg.sender, address(this), _depositAmount);
        require(_success, "Failed to transfer tokens.");

        emit ERC20Deposited(msg.sender, _odinAddress, _tokenAddress, _depositAmount);
        return true;
    }

    /**
    * @notice Adds ERC20 compatible token to supported tokens
    * @param _tokenAddress Address of the ERC20 compatible token contract
    * @return True if everything went well
    */
    function addToken(address _tokenAddress) external onlyOwner returns (bool) {
        supportedTokens[_tokenAddress] = true;
        emit TokenAdded(_tokenAddress);
        return true;
    }

    /**
    * @notice Removes ERC20 compatible token from supported tokens
    * @param _tokenAddress Address of the ERC20 compatible token contract
    * @return True if everything went well
    */
    function removeToken(address _tokenAddress) external onlyOwner returns (bool) {
        supportedTokens[_tokenAddress] = false;
        emit TokenRemoved(_tokenAddress);
        return true;
    }

    /**
    * @notice Transfers the amount of the deposit if an error occurred during the deposit
    * @param _user Depositor
    * @param _amount Deposit amount
    * @return True if everything went well
    */
    function payBackEther(address _user, uint256 _amount) external onlyOwner returns (bool) {
        (bool _success,) = _user.call{value : _amount}("");
        require(_success, "Failed to pay back the deposit amount.");

        return true;
    }

    /**
    * @notice Transfers the amount of the deposit if an error occurred during the deposit
    * @param _user Depositor
    * @param _amount Deposit amount
    * @return True if everything went well
    */
    function payBackToken(address _user, address _token, uint256 _amount) external onlyOwner returns (bool) {
        require(supportedTokens[_token], "Unsupported token.");
        bool _success = IERC20(_token).transfer(_user, _amount);
        require(_success, "Failed to pay back");

        return true;
    }
}
