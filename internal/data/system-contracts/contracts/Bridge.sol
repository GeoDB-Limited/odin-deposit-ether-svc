// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./IERC20Token.sol";


contract Bridge is Ownable {
    using SafeMath for uint256;
    using Address for address;

    uint256 public refundFee;
    bool depositingAllowed;
    bool lockingFundsAllowed;
    bool claimingLockedFundsAllowed;


    mapping(address => uint256) public lockedETH;
    mapping(address => mapping(address => uint256)) public lockedERC20;
    mapping(address => bool) public refundFeeDeposited;
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

    constructor(
        address[] memory _supportedTokens,
        uint256 _refundFee,
        bool _depositingAllowed,
        bool _lockingFundsAllowed,
        bool _claimingLockedFundsAllowed
    ) {
        for (uint256 i = 0; i < _supportedTokens.length; i++) {
            supportedTokens[_supportedTokens[i]] = true;
        }

        refundFee = _refundFee;
        depositingAllowed = _depositingAllowed;
        lockingFundsAllowed = _lockingFundsAllowed;
        claimingLockedFundsAllowed = _claimingLockedFundsAllowed;
    }

    /**
    * @notice Deposits ETH
    * @param _odinAddress Address in the Odin chain
    * @return True if everything went well
    */
    function depositETH(string memory _odinAddress) external onlyDepositingAllowed payable returns (bool) {
        uint256 _depositAmount = msg.value;

        if (!refundFeeDeposited[msg.sender]) {
            uint256 _depositCompensation = refundFee;
            require(_depositAmount >= _depositCompensation, "Insufficient funds to deposit compensation");

            refundFeeDeposited[msg.sender] = true;
            _depositAmount = _depositAmount.sub(_depositCompensation);
        }

        require(_depositAmount > 0, "Invalid value for the deposit amount, failed to deposit a zero value.");

        if (lockingFundsAllowed) {
            lockedETH[msg.sender] = lockedETH[msg.sender].add(_depositAmount);
        }

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
    external onlyDepositingAllowed payable returns (bool)
    {
        require(_tokenAddress.isContract(), "Given token is not a contract");
        require(supportedTokens[_tokenAddress], "Unsupported token, failed to deposit.");

        if (!refundFeeDeposited[msg.sender]) {
            require(msg.value >= refundFee, "Insufficient funds for deposit compensation");
            refundFeeDeposited[msg.sender] = true;
        }

        IERC20Token _token = IERC20Token(_tokenAddress);

        bool _success = _token.transferFrom(msg.sender, address(this), _depositAmount);
        require(_success, "Failed to transfer tokens.");

        if (lockingFundsAllowed) {
            lockedERC20[msg.sender][_tokenAddress] = lockedERC20[msg.sender][_tokenAddress].add(_depositAmount);
        }

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
    * @param _refundFee Amount of refund fee
    * @return True if everything went well
    */
    function setRefundFee(uint256 _refundFee) external onlyOwner returns (bool) {
        refundFee = _refundFee;
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

        (_success,) = payable(msg.sender).call{value : refundFee}("");
        require(_success, "Failed to pay the compensation for paying back.");
        refundFeeDeposited[_user] = false;

        return true;
    }

    /**
    * @notice Transfers the amount of the deposit if an error occurred during the deposit
    * @param _user Depositor
    * @param _tokenAddress Token address
    * @param _amount Deposit amount
    * @return True if everything went well
    */
    function payBackERC20(address _user, address _tokenAddress, uint256 _amount) external onlyOwner returns (bool) {
        bool _success = IERC20Token(_tokenAddress).transfer(_user, _amount);
        require(_success, "Failed to pay back");

        (_success,) = payable(msg.sender).call{value : refundFee}("");
        require(_success, "Failed to pay the compensation for paying back.");
        refundFeeDeposited[_user] = false;

        return true;
    }

    /**
    * @notice Transfers locked ETH to msg sender
    * @param _claimableAmount Claimable amount
    * @return True if everything went well
    */
    function claimLockedETH(uint256 _claimableAmount) external onlyClaimingLockedFundsAllowed returns (bool) {
        uint256 _lockedAmount = lockedETH[msg.sender];
        require(
            _claimableAmount <= _lockedAmount,
            "Insufficient locked ETH."
        );

        (bool _success,) = payable(msg.sender).call{value : _claimableAmount}("");
        require(_success, "Failed to transfer claimed ETH.");

        lockedETH[msg.sender] = _lockedAmount.sub(_claimableAmount);

        return true;
    }

    /**
    * @notice Transfers locked ERC20 to msg sender
    * @param _claimableAmount Claimable amount
    * @param _tokenAddress Claimable token address
    * @return True if everything went well
    */
    function claimLockedERC20(uint256 _claimableAmount, address _tokenAddress)
    external onlyClaimingLockedFundsAllowed returns (bool)
    {
        uint256 _lockedAmount = lockedERC20[msg.sender][_tokenAddress];
        require(
            _claimableAmount <= _lockedAmount,
            "Insufficient locked ERC20."
        );

        bool _success = IERC20Token(_tokenAddress).transfer(msg.sender, _claimableAmount);
        require(_success, "Failed to transfer locked ERC20.");

        lockedERC20[msg.sender][_tokenAddress] = _lockedAmount.sub(_claimableAmount);

        return true;
    }

    /**
    * @notice Sets allowance to lock deposit assets
    * @param _allowed If it is allowed to lock deposit assets
    * @return True if everything went well
    */
    function setAllowanceToLock(bool _allowed) external onlyOwner returns (bool) {
        lockingFundsAllowed = _allowed;

        return true;
    }

    /**
    * @notice Sets allowance to claim locked funds
    * @param _allowed If it is allowed to claim locked funds
    * @return True if everything went well
    */
    function setAllowanceToClaimLockedFunds(bool _allowed) external onlyOwner returns (bool) {
        claimingLockedFundsAllowed = _allowed;

        return true;
    }

    /**
    * @notice Sets allowance to deposit
    * @param _allowed If it is allowed to deposit
    * @return True if everything went well
    */
    function setAllowanceToDeposit(bool _allowed) external onlyOwner returns (bool) {
        depositingAllowed = _allowed;

        return true;
    }

    /**
    * @notice Transfers contract ETH to contract owner
    * @param _claimableAmount Claimable amount
    * @return True if everything went well
    */
    function claimContractETH(uint256 _claimableAmount) external onlyOwner returns (bool) {
        (bool _success,) = payable(msg.sender).call{value : _claimableAmount}("");
        require(_success, "Failed to transfer claimed amount.");

        return true;
    }

    /**
    * @notice Transfers contract ERC20 to contract owner
    * @param _claimableAmount Claimable amount
    * @param _tokenAddress Claimable token address
    * @return True if everything went well
    */
    function claimContractERC20(uint256 _claimableAmount, address _tokenAddress) external onlyOwner returns (bool) {
        bool _success = IERC20Token(_tokenAddress).transfer(msg.sender, _claimableAmount);
        require(_success, "Failed to transfer claimed amount.");

        return true;
    }

    /**
    * @notice Requires allowance to deposit
    */
    modifier onlyDepositingAllowed() {
        require(depositingAllowed, "It is not allowed to deposit.");
        _;
    }

    /**
    * @notice Requires allowance to claim locked funds
    */
    modifier onlyClaimingLockedFundsAllowed() {
        require(claimingLockedFundsAllowed, "It is not allowed to claim locked funds.");
        _;
    }
}
