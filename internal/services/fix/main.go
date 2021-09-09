package fix

import (
	"context"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/data/system-contracts/generated"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/odin/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tendermint/tendermint/abci/types"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	"math/big"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	rpcclient "github.com/tendermint/tendermint/rpc/client"
	httpclient "github.com/tendermint/tendermint/rpc/client/http"
)

// Service defines a service that allows you to receive unprocessed transactions.
type Service struct {
	config   config.Config
	context  context.Context
	logger   *logrus.Logger
	ethereum *ethclient.Client
	contract *generated.Bridge
	odin     client.Client
}

type withdrawalEvent struct {
	withdrawalAmount string
	receiver         string
}

// New creates a service that allows you to receive unprocessed transactions.
func New(ctx context.Context, cfg config.Config) *Service {
	odinClient := client.New(ctx, cfg).WithSigner()
	bridgeAddr := cfg.FixConfig().BridgeAddress

	etherClient := cfg.EthereumClient()
	bridge, err := generated.NewBridge(bridgeAddr, etherClient)
	if err != nil {
		panic(errors.Wrap(err, "failed to create bridge instance"))
	}

	return &Service{
		config:   cfg,
		context:  ctx,
		logger:   cfg.Logger(),
		odin:     odinClient,
		ethereum: etherClient,
		contract: bridge,
	}
}

// Run performs events listening and querying the Odin minting module.
func (s *Service) Run() {
	s.logger.Info("Get unprocessed events...")
	unprocessedEvents := s.getUnprocessedEvents()

	for i, event := range unprocessedEvents {
		s.logger.WithFields(logrus.Fields{
			"event_index":		i,
			"odin_address":		event.OdinAddress,
			"deposit_amount":	event.DepositAmount.String(),
		}).Info("Unprocessed event")
	}

	s.logger.Info("Process events...")
	s.processEvents(unprocessedEvents)
}

// processEvents process unprocessed events from ethereum smart contract
func (s *Service) processEvents(unprocessedEvents []*generated.BridgeTokensDeposited) {
	for _, event := range unprocessedEvents {
		withdrawalAmount, err := s.exchangeTokens(event.DepositAmount, event.Symbol, int64(event.TokenPrecision))
		if err != nil {
			panic(errors.Wrap(err, "failed to exchange tokens"))
		}

		if err := s.odin.ClaimWithdrawal(event.OdinAddress, withdrawalAmount); err != nil {
			println(err.Error())
			s.logger.WithFields(logrus.Fields{
				"odin_address": 		event.OdinAddress,
				"withdrawal_amount": 	withdrawalAmount,
			}).Warn("Failed to send tokens")
		} else {
			s.logger.WithFields(logrus.Fields{
				"odin_address":			event.OdinAddress,
				"withdrawal_amount":	withdrawalAmount,
			}).Info("Tokens sent successfully")
		}
	}
}

// getUnprocessedEvents returns current unprocessed events
func (s *Service) getUnprocessedEvents() []*generated.BridgeTokensDeposited {
	unprocessedEvents := make([]*generated.BridgeTokensDeposited, 0, 10)

	contractEvents := s.getContractEvents()
	internalEvents := s.getInternalEvents()

	for _, event := range contractEvents {
		var isProcessed bool
		exchangedAmount, err := s.exchangeTokens(event.DepositAmount, event.Symbol, int64(event.TokenPrecision))
		if err != nil {
			panic(errors.Wrap(err, "failed to exchange tokens"))
		}

		for i, internalEvent := range internalEvents {
			if event.OdinAddress == internalEvent.receiver && exchangedAmount.String() == internalEvent.withdrawalAmount {
				internalEvents[i] = internalEvents[len(internalEvents)-1]
				internalEvents = internalEvents[:len(internalEvents)-1]

				isProcessed = true
				break
			}
		}

		if !isProcessed {
			unprocessedEvents = append(unprocessedEvents, event)
		}
	}

	return unprocessedEvents
}

// getContractEvents returns events from the ethereum contracts since a certain block
func (s *Service) getContractEvents() []*generated.BridgeTokensDeposited {
	fromBlockNumber := s.config.FixConfig().FromBlockNumber.Uint64()

	it, err := s.contract.FilterTokensDeposited(&bind.FilterOpts{Start: fromBlockNumber, Context: s.context}, []common.Address{}, []common.Address{})
	if err != nil {
		panic(errors.Wrap(err, "failed to get contract events"))
	}

	resultArr := make([]*generated.BridgeTokensDeposited, 0, 10)

	for it.Next() {
		resultArr = append(resultArr, it.Event)
	}

	return resultArr
}

// getInternalEvents returns internal system events
func (s *Service) getInternalEvents() []*withdrawalEvent {
	resultArray := make([]*withdrawalEvent, 0, 10)
	currentPage := 1
	maxTxsPerPage := s.config.FixConfig().MaxTxsPerPage

	res, err := s.searchTxs(currentPage, 1) //Get transaction count
	if err != nil {
		panic(errors.Wrap(err, "failed to search txs"))
	}

	neededIteration := res.TotalCount / maxTxsPerPage
	if res.TotalCount%maxTxsPerPage != 0 {
		neededIteration++
	}

	var currentEvent types.Event

	for i := 0; i < neededIteration; i++ {
		res, err = s.searchTxs(currentPage, maxTxsPerPage)
		if err != nil {
			panic(errors.Wrap(err, "failed to search txs"))
		}

		for _, tx := range res.Txs {
			for _, event := range tx.TxResult.Events {
				if event.Type == "withdrawal" {
					currentEvent = event
				}
			}

			resultArray = append(resultArray, &withdrawalEvent{withdrawalAmount: string(currentEvent.Attributes[0].Value), receiver: string(currentEvent.Attributes[1].Value)})
		}

		currentPage++
	}

	return resultArray
}

// searchTxs searches txs with withdrawal event
func (s *Service) searchTxs(page, maxTxsPerPage int) (*coretypes.ResultTxSearch, error) {
	newClient, err := httpclient.New(s.config.FixConfig().ClientEndpoint, "/websocket")
	if err != nil {
		panic(errors.Wrap(err, "failed to create http client"))
	}

	err = newClient.Start()
	if err != nil {
		panic(errors.Wrap(err, "failed to start client"))
	}
	defer newClient.Stop()

	query := "tx.height>=0 AND withdrawal.withdrawal_amount EXISTS"

	resSearch, err := rpcclient.Client(newClient).TxSearch(s.context, query, false, &page, &maxTxsPerPage, "asc")

	if err != nil {
		return nil, sdkerrors.Wrap(err, "failed to find txs")
	}

	return resSearch, nil
}

// exchangeTokens exchanges deposited ERC20 to odin tokens
func (s *Service) exchangeTokens(
	amount *big.Int,
	tokenSymbol string,
	tokenPrecision int64,
) (sdk.Coin, error) {
	rate, err := s.odin.GetExchangeRate(tokenSymbol)
	if err != nil {
		return sdk.Coin{}, errors.Wrapf(err, "failed to get the exchange rate for %s", tokenSymbol)
	}

	withdrawalAmount, err := s.exchange(amount, rate, tokenPrecision)
	if err != nil {
		return sdk.Coin{}, errors.Wrapf(
			err,
			"failed to exchange the deposit: %s with rate: %s",
			amount.String(),
			rate.String(),
		)
	}

	if withdrawalAmount.Amount.Equal(sdk.NewIntFromUint64(0)) {
		return sdk.Coin{}, errors.New("insufficient deposit amount")
	}

	return withdrawalAmount, nil
}

// exchange calculates new coin with the given exchange rate
func (s *Service) exchange(amount *big.Int, rate sdk.Dec, precision int64) (sdk.Coin, error) {
	withdrawalAmount := sdk.NewDecFromBigIntWithPrec(amount, precision).Mul(rate)
	return sdk.NewCoin(s.config.OdinConfig().Denom, withdrawalAmount.TruncateInt()), nil
}
