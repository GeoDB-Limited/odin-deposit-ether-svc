package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

// ETHTransferDetails defines unpacked data of the event of depositing ETH.
type ETHTransferDetails struct {
	EthAddress      common.Address
	OdinAddress     string
	DepositAmount   *big.Int
	TransactionHash string
	BlockTime       time.Time
}

// ERC20TransferDetails defines unpacked data of the event of depositing ERC20 tokens.
type ERC20TransferDetails struct {
	EthAddress      common.Address
	OdinAddress     string
	DepositAmount   *big.Int
	TokenAddress    common.Address
	TokenSymbol     string
	TransactionHash string
	BlockTime       time.Time
}
