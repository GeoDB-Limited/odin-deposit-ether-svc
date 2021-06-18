// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./IERC20Token.sol";


contract Bridge is Ownable {
    using SafeMath for uint256;
    using Address for address;

    uint256 public depositCompensation;
    mapping(address => bool) public compensationDeposited;
    mapping(address => bool) public supportedTokens;

    event ETHDeposited(
        address indexed _userAddress,
        string _odinAddress,
        uint256 _depositAmount
    );
    event ERC20Deposited(
        address indexed _userAddress,
        string _odinAddress,
        uint256 _depositAmount,
        address indexed _tokenAddress,
        string _symbol,
        uint8 _tokenPrecision
    );
    event TokenAdded(address indexed _tokenAddress);
    event TokenRemoved(address indexed _tokenAddress);

    constructor(address[] memory _supportedTokens, uint256 _depositCompensation) {
        for (uint256 i = 0; i < _supportedTokens.length; i++) {
            supportedTokens[_supportedTokens[i]] = true;
        }

        depositCompensation = _depositCompensation;
    }

    /**
    * @notice Deposits ETH
    * @param _odinAddress Address in the Odin chain
    * @return True if everything went well
    */
    function depositETH(string memory _odinAddress) external payable returns (bool) {
        uint256 _depositAmount = msg.value;

        if (!compensationDeposited[msg.sender]) {
            uint256 _depositCompensation = depositCompensation;
            require(_depositAmount >= _depositCompensation, "Insufficient funds to deposit compensation");

            compensationDeposited[msg.sender] = true;
            _depositAmount = _depositAmount.sub(_depositCompensation);
        }

        require(_depositAmount > 0, "Invalid value for the deposit amount, failed to deposit a zero value.");

        emit ETHDeposited(msg.sender, _odinAddress, _depositAmount);
        return true;
    }

    /**
    * @notice Deposits ERC20 compatible tokens
    * @param _tokenAddress Address of the ERC20 compatible token contract
    * @param _odinAddress Address in the Odin chain
    * @param _depositAmount Amount to deposit
    * @return True if everything went well
    */
    function depositERC20(address _tokenAddress, string memory _odinAddress, uint256 _depositAmount)
    external payable returns (bool)
    {
        require(_tokenAddress.isContract(), "Given token is not a contract");
        require(supportedTokens[_tokenAddress], "Unsupported token, failed to deposit.");

        if (!compensationDeposited[msg.sender]) {
            require(msg.value >= depositCompensation, "Insufficient funds for deposit compensation");
            compensationDeposited[msg.sender] = true;
        }

        IERC20Token _token = IERC20Token(_tokenAddress);

        bool _success = _token.transferFrom(msg.sender, address(this), _depositAmount);
        require(_success, "Failed to transfer tokens.");

        emit ERC20Deposited(msg.sender, _odinAddress, _depositAmount, _tokenAddress, _token.symbol(), _token.decimals());
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
    * @notice Sets a new compensation amount for paying back
    * @param _amount Amount of compensation
    * @return True if everything went well
    */
    function setDepositCompensation(uint256 _amount) external onlyOwner returns (bool) {
        depositCompensation = _amount;
        return true;
    }

    /**
    * @notice Transfers the amount of the deposit if an error occurred during the deposit
    * @param _user Depositor
    * @param _amount Deposit amount
    * @return True if everything went well
    */
    function payBackETH(address _user, uint256 _amount) external onlyOwner returns (bool) {
        (bool _success,) = payable(_user).call{value : _amount}("");
        require(_success, "Failed to pay back the deposit amount.");

        (_success,) = payable(msg.sender).call{value : depositCompensation}("");
        require(_success, "Failed to pay the compensation for paying back.");
        compensationDeposited[_user] = false;

        return true;
    }

    /**
    * @notice Transfers the amount of the deposit if an error occurred during the deposit
    * @param _user Depositor
    * @param _amount Deposit amount
    * @return True if everything went well
    */
    function payBackERC20(address _user, address _token, uint256 _amount) external onlyOwner returns (bool) {
        bool _success = IERC20Token(_token).transfer(_user, _amount);
        require(_success, "Failed to pay back");

        (_success,) = payable(msg.sender).call{value : depositCompensation}("");
        require(_success, "Failed to pay the compensation for paying back.");
        compensationDeposited[_user] = false;

        return true;
    }

    /**
    * @notice Transfers contract ETH funds to owner
    * @param _amount Claimable amount
    * @return True if everything went well
    */
    function claimETH(uint256 _amount) external onlyOwner returns (bool) {
        (bool _success,) = payable(msg.sender).call{value : _amount}("");
        require(_success, "Failed to claim contract ETH funds.");
        return true;
    }

    /**
    * @notice Transfers contract ERC20 funds to owner
    * @param _amount Claimable amount
    * @param _token Claimable token address
    * @return True if everything went well
    */
    function claimERC20(uint256 _amount, address _token) external onlyOwner returns (bool) {
        bool _success = IERC20Token(_token).transfer(msg.sender, _amount);
        require(_success, "Failed to claim contract ERC20 funds.");
        return true;
    }
}
