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
    mapping(address => mapping(address => Refund)) public refund;
    mapping(address => mapping(address => uint256)) public lockedTokens;

    event TokensDeposited(
        address indexed _userAddress,
        string _odinAddress,
        uint256 _depositAmount,
        address indexed _tokenAddress,
        string _symbol,
        uint8 _tokenPrecision
    );
    event TokenAdded(address indexed _tokenAddress);
    event TokenRemoved(address indexed _tokenAddress);
    event RefundSet(address indexed _userAddress, address indexed _tokenAddress, uint256 _refundAmount);
    event RefundClaimed(address indexed _userAddress, address indexed _tokenAddress, uint256 _refundAmount);
    event TokensLocked(
        address indexed _userAddress,
        string _odinAddress,
        uint256 _depositAmount,
        address indexed _tokenAddress,
        string _symbol,
        uint8 _tokenPrecision
    );
    event LockedTokensClaimed(address indexed _userAddress, address indexed _tokenAddress, uint256 _lockedAmount);

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
    * @notice Deposits ERC20 compatible tokens
    * @param _tokenAddress Address of the ERC20 compatible token contract
    * @param _odinAddress Address in the Odin chain
    * @param _depositAmount Amount to deposit
    * @return True if everything went well
    */
    function deposit(string memory _odinAddress, address _tokenAddress, uint256 _depositAmount)
    external onlyDepositingAllowed returns (bool)
    {
        require(_tokenAddress.isContract(), "Given token is not a contract.");
        require(supportedTokens[_tokenAddress], "Unsupported token, failed to deposit.");

        IERC20Token _token = IERC20Token(_tokenAddress);
        address _bridgeAddress = address(this);
        uint256 _balanceBeforeTransfer = _token.balanceOf(_bridgeAddress);

        bool _success = _token.transferFrom(msg.sender, _bridgeAddress, _depositAmount);
        require(_success, "Failed to transfer tokens.");

        uint256 _balanceAfterTransfer = _token.balanceOf(_bridgeAddress);
        uint256 _actualDepositAmount = _balanceAfterTransfer.sub(_balanceBeforeTransfer);

        if (lockingFundsAllowed) {
            lockedTokens[msg.sender][_tokenAddress] = lockedTokens[msg.sender][_tokenAddress].add(_actualDepositAmount);
            emit TokensLocked(
                msg.sender,
                _odinAddress,
                _actualDepositAmount,
                _tokenAddress,
                _token.symbol(),
                _token.decimals()
            );
        }

        emit TokensDeposited(
            msg.sender,
            _odinAddress,
            _actualDepositAmount,
            _tokenAddress,
            _token.symbol(),
            _token.decimals()
        );

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
    * @notice Stores the amount of refund ERC20
    * @param _userAddress Address of user who deposited
    * @param _userAddress Address of refund token
    * @param _refundAmount Refund amount
    * @return True if everything went well
    */
    function setRefund(address _userAddress, address _tokenAddress, uint256 _refundAmount)
    external onlyOwner returns (bool)
    {
        Refund memory _refund = refund[_userAddress][_tokenAddress];
        _refund.amount = _refund.amount.add(_refundAmount);
        _refund.fee = _refund.fee.add(tx.gasprice.mul(refundGas));
        refund[_userAddress][_tokenAddress] = _refund;

        if (lockingFundsAllowed) {
            lockedTokens[_userAddress][_tokenAddress] = lockedTokens[_userAddress][_tokenAddress].sub(_refundAmount);
        }

        emit RefundSet(_userAddress, _tokenAddress, _refundAmount);
        return true;
    }

    /**
    * @notice Refunds ERC20 to msg sender
    * @param _tokenAddress Address of claimable refund token
    * @return True if everything went well
    */
    function claimRefund(address _tokenAddress) external payable returns (bool) {
        Refund memory _refund = refund[msg.sender][_tokenAddress];
        require(_refund.fee > 0, "Zero refund amount.");
        require(msg.value >= _refund.fee, "Insufficient refund fee.");

        bool _success = IERC20Token(_tokenAddress).transfer(msg.sender, _refund.amount);
        require(_success, "Failed to transfer locked ERC20.");

        (_success,) = payable(owner()).call{value : _refund.fee}("");
        require(_success, "Failed to pay the compensation for paying back.");

        emit RefundClaimed(msg.sender, _tokenAddress, _refund.amount);

        delete refund[msg.sender][_tokenAddress];

        return true;
    }

    /**
    * @notice Transfers locked ERC20 to msg sender
    * @param _tokenAddress Claimable token address
    * @return True if everything went well
    */
    function claimLockedTokens(address _tokenAddress)
    external onlyClaimingLockedFundsAllowed returns (bool)
    {
        uint256 _lockedAmount = lockedTokens[msg.sender][_tokenAddress];
        require(
            _lockedAmount > 0,
            "Zero locked amount."
        );

        bool _success = IERC20Token(_tokenAddress).transfer(msg.sender, _lockedAmount);
        require(_success, "Failed to transfer locked ERC20.");

        emit LockedTokensClaimed(msg.sender, _tokenAddress, _lockedAmount);

        delete lockedTokens[msg.sender][_tokenAddress];

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
    * @notice Transfers contract ERC20 to contract owner
    * @param _claimableAmount Claimable amount
    * @param _tokenAddress Claimable token address
    * @return True if everything went well
    */
    function claimContractTokens(uint256 _claimableAmount, address _tokenAddress) external onlyOwner returns (bool) {
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
