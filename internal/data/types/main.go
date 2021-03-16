package types

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

// TransferDetails defines unpacked data of the event.
type TransferDetails struct {
	DepositAmount   *big.Int
	UserAddress     common.Address
	OdinAddress     string
	TransactionHash string
	BlockTime       time.Time
}

// Transfer defines a parsed event log
type Transfer struct {
	UserAddress   common.Address
	OdinAddress   string
	DepositAmount *big.Int
}
