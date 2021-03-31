package exchanger

import (
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/types"
	"github.com/sirupsen/logrus"
)

// Service defines a service that exchanges deposits.
type Service struct {
	config config.Config
	log    *logrus.Logger

	ethTransferDetails   <-chan types.ETHTransferDetails
	erc20TransferDetails <-chan types.ERC20TransferDetails
	withdrawalDetails    chan<- types.WithdrawalDetails
}

// New creates a service that exchanges deposits.
func New(
	cfg config.Config,
	ethTransferDetails <-chan types.ETHTransferDetails,
	erc20TransferDetails <-chan types.ERC20TransferDetails,
	withdrawalDetails chan<- types.WithdrawalDetails,
) *Service {
	return &Service{
		config: cfg,
		log:    cfg.Logger(),

		ethTransferDetails:   ethTransferDetails,
		erc20TransferDetails: erc20TransferDetails,
		withdrawalDetails:    withdrawalDetails,
	}
}

// Run performs exchanging
func (s *Service) Run() {
	go s.exchangeETH()
	go s.exchangeERC20()
}

// exchangeETH exchanges ETH to odin
func (s *Service) exchangeETH() {
	for data := range s.ethTransferDetails {
		// TODO: Implement exchange logic

		s.log.WithFields(logrus.Fields{
			"odin_address": data.OdinAddress,
			"amount":       data.DepositAmount,
		}).Info("Exchange")

		s.withdrawalDetails <- types.WithdrawalDetails{
			OdinAddress:      data.OdinAddress,
			WithdrawalAmount: data.DepositAmount,
		}
	}
}

// exchangeERC20 exchanges ERC20 to odin
func (s *Service) exchangeERC20() {
	for data := range s.erc20TransferDetails {
		// TODO: Implement Exchange logic

		s.log.WithFields(logrus.Fields{
			"odin_address": data.OdinAddress,
			"amount":       data.DepositAmount,
		}).Info("Exchange")

		s.withdrawalDetails <- types.WithdrawalDetails{
			OdinAddress:      data.OdinAddress,
			WithdrawalAmount: data.DepositAmount,
		}
	}
}
