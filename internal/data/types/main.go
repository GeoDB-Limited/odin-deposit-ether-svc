package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

// ETHTransfer defines a parsed event log
type ETHTransfer struct {
	UserAddress   common.Address
	OdinAddress   string
	DepositAmount *big.Int
}

// ETHTransferDetails defines unpacked data of the event of depositing ETH.
type ETHTransferDetails struct {
	DepositAmount   *big.Int
	UserAddress     common.Address
	OdinAddress     string
	TransactionHash string
	BlockTime       time.Time
}

// ERC20Transfer defines a parsed event log
type ERC20Transfer struct {
	UserAddress   common.Address
	OdinAddress   string
	DepositAmount *big.Int
	TokenAddress  common.Address
}

// ERC20TransferDetails defines unpacked data of the event of depositing ERC20 tokens.
type ERC20TransferDetails struct {
	DepositAmount   *big.Int
	UserAddress     common.Address
	OdinAddress     string
	TransactionHash string
	BlockTime       time.Time
	TokenAddress    common.Address
}

// WithdrawalDetails defines a data for querying the withdrawal
type WithdrawalDetails struct {
	WithdrawalAmount *big.Int
	OdinAddress      string
}
