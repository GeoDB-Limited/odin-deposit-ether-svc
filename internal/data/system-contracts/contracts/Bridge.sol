// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./IERC20Token.sol";


contract Bridge is Ownable {
    using SafeMath for uint256;
    using Address for address;

    struct Refund {
        uint256 fee;
        uint256 amount;
    }

    uint256 refundGas;
    bool public depositingAllowed;
    bool public lockingFundsAllowed;
    bool public claimingLockedFundsAllowed;

    mapping(address => bool) public supportedTokens;

    mapping(address => Refund) public refundETH;
    mapping(address => mapping(address => Refund)) public refundERC20;

    mapping(address => uint256) public lockedETH;
    mapping(address => mapping(address => uint256)) public lockedERC20;

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
    event RefundETHSet(address indexed _userAddress, uint256 _refundAmount);
    event RefundERC20Set(address indexed _userAddress, address indexed _tokenAddress, uint256 _refundAmount);
    event RefundETHClaimed(address indexed _userAddress, uint256 _refundAmount);
    event RefundERC20Claimed(address indexed _userAddress, address indexed _tokenAddress, uint256 _refundAmount);

    constructor(
        address[] memory _supportedTokens,
        uint256 _refundGasLimit,
        bool _depositingAllowed,
        bool _lockingFundsAllowed,
        bool _claimingLockedFundsAllowed
    ) {
        for (uint256 i = 0; i < _supportedTokens.length; i++) {
            supportedTokens[_supportedTokens[i]] = true;
        }

        refundGas = _refundGasLimit;
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
        require(msg.value > 0, "Invalid value for the deposit amount, failed to deposit a zero value.");

        if (lockingFundsAllowed) {
            lockedETH[msg.sender] = lockedETH[msg.sender].add(msg.value);
        }

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
    function depositERC20(address _tokenAddress, string memory _odinAddress, uint256 _depositAmount)
    external onlyDepositingAllowed payable returns (bool)
    {
        require(_tokenAddress.isContract(), "Given token is not a contract.");
        require(supportedTokens[_tokenAddress], "Unsupported token, failed to deposit.");

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
    * @notice Stores the amount of refund ETH
    * @param _userAddress Address of user who deposited
    * @param _refundAmount Refund amount
    * @return True if everything went well
    */
    function setRefundETH(address _userAddress, uint256 _refundAmount) external onlyOwner returns (bool) {
        Refund memory _refund = refundETH[_userAddress];
        _refund.amount = _refund.amount.add(_refundAmount);
        _refund.fee = _refund.fee.add(tx.gasprice.mul(refundGas));
        refundETH[_userAddress] = _refund;

        if (lockingFundsAllowed) {
            lockedETH[_userAddress] = lockedETH[_userAddress].sub(_refundAmount);
        }

        emit RefundETHSet(_userAddress, _refundAmount);
        return true;
    }

    /**
    * @notice Stores the amount of refund ERC20
    * @param _userAddress Address of user who deposited
    * @param _userAddress Address of refund token
    * @param _refundAmount Refund amount
    * @return True if everything went well
    */
    function setRefundERC20(address _userAddress, address _tokenAddress, uint256 _refundAmount)
    external onlyOwner returns (bool)
    {
        Refund memory _refund = refundERC20[_userAddress][_tokenAddress];
        _refund.amount = _refund.amount.add(_refundAmount);
        _refund.fee = _refund.fee.add(tx.gasprice.mul(refundGas));
        refundERC20[_userAddress][_tokenAddress] = _refund;

        if (lockingFundsAllowed) {
            lockedERC20[_userAddress][_tokenAddress] = lockedERC20[_userAddress][_tokenAddress].sub(_refundAmount);
        }

        emit RefundERC20Set(_userAddress, _tokenAddress, _refundAmount);
        return true;
    }

    /**
    * @notice Refunds ETH to msg sender
    * @return True if everything went well
    */
    function claimRefundETH() external payable returns (bool) {
        Refund memory _refund = refundETH[msg.sender];
        require(_refund.fee > 0, "Zero refund amount.");
        require(msg.value >= _refund.fee, "Insufficient refund fee.");

        (bool _success,) = payable(msg.sender).call{value : _refund.amount}("");
        require(_success, "Failed to transfer claimed ETH.");

        (_success,) = payable(owner()).call{value : _refund.fee}("");
        require(_success, "Failed to pay the compensation for paying back.");

        emit RefundETHClaimed(msg.sender, _refund.amount);

        delete refundETH[msg.sender];

        return true;
    }

    /**
    * @notice Refunds ERC20 to msg sender
    * @param _tokenAddress Address of claimable refund token
    * @return True if everything went well
    */
    function claimRefundERC20(address _tokenAddress)
    external payable returns (bool)
    {
        Refund memory _refund = refundERC20[msg.sender][_tokenAddress];
        require(_refund.fee > 0, "Zero refund amount.");
        require(msg.value >= _refund.fee, "Insufficient refund fee.");

        bool _success = IERC20Token(_tokenAddress).transfer(msg.sender, _refund.amount);
        require(_success, "Failed to transfer locked ERC20.");

        (_success,) = payable(owner()).call{value : _refund.fee}("");
        require(_success, "Failed to pay the compensation for paying back.");

        emit RefundERC20Claimed(msg.sender, _tokenAddress, _refund.amount);

        delete refundERC20[msg.sender][_tokenAddress];

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
    * @notice Sets refund gas limit
    * @param _gas Amount of gas
    * @return True if everything went well
    */
    function setRefundGas(uint256 _gas) external onlyOwner returns (bool) {
        refundGas = _gas;
        return true;
    }


    /**
    * @notice Sets allowance to lock deposit assets
    * @param _allowed If it is allowed to lock deposit assets
    * @return True if everything went well
    */
    function setAllowanceToLock(bool _allowed) external onlyOwner returns (bool) {
        require(lockingFundsAllowed != _allowed, "Trying to set the same parameter value.");
        lockingFundsAllowed = _allowed;

        return true;
    }

    /**
    * @notice Sets allowance to claim locked funds
    * @param _allowed If it is allowed to claim locked funds
    * @return True if everything went well
    */
    function setAllowanceToClaimLockedFunds(bool _allowed) external onlyOwner returns (bool) {
        require(claimingLockedFundsAllowed != _allowed, "Trying to set the same parameter value.");
        claimingLockedFundsAllowed = _allowed;

        return true;
    }

    /**
    * @notice Sets allowance to deposit
    * @param _allowed If it is allowed to deposit
    * @return True if everything went well
    */
    function setAllowanceToDeposit(bool _allowed) external onlyOwner returns (bool) {
        require(depositingAllowed != _allowed, "Trying to set the same parameter value.");
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
