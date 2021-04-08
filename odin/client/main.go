package client

import (
	"context"
	odinapp "github.com/GeoDB-Limited/odin-core/app"
	odinminttypes "github.com/GeoDB-Limited/odin-core/x/mint/types"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	sdktxclient "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	sdkauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	sdkauthtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io/ioutil"
	"math/big"
)

var encoding = odinapp.MakeEncodingConfig()

// Client defines an interface for the wrapped cosmos sdk service client.
type Client interface {
	SetBridgeAddress(common.Address) error
	GetBridgeAddress() (common.Address, error)
	ClaimWithdrawal(string, *big.Int) error
	GetExchangeRate(string) (*big.Int, error)
}

// client defines typed wrapper for the cosmos sdk service client.
type client struct {
	connection *grpc.ClientConn
	config     config.Config
	log        *logrus.Logger
}

// New creates a client that uses the given cosmos sdk service client.
func New(cfg config.Config) Client {
	odinConfig := cfg.OdinConfig()
	clientConn, err := grpc.Dial(odinConfig.Endpoint, grpc.WithInsecure())
	if err != nil {
		panic(errors.Wrapf(err, "failed to dial %s", odinConfig.Endpoint))
	}

	return &client{
		connection: clientConn,
		config:     cfg,
		log:        cfg.Logger(),
	}
}

// SetBridgeAddress sets an address of the bridge contract to the storage.
func (c client) SetBridgeAddress(address common.Address) error {
	if err := ioutil.WriteFile(c.config.BridgeAddressStorage(), address.Bytes(), 0777); err != nil {
		return errors.Wrap(err, "failed to add the address to the storage")
	}
	return nil
}

// GetBridgeAddress returns an address of the bridge contract.
func (c client) GetBridgeAddress() (common.Address, error) {
	mintClient := odinminttypes.NewQueryClient(c.connection)
	response, err := mintClient.EthIntegrationAddress(context.TODO(), &odinminttypes.QueryEthIntegrationAddressRequest{})
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to query ethereum integration address")
	}

	return common.HexToAddress(response.EthIntegrationAddress), nil
}

// ClaimWithdrawal claims minting from Odin
func (c client) ClaimWithdrawal(address string, amount *big.Int) error {
	txBuilder := encoding.TxConfig.NewTxBuilder()

	odinConfig := c.config.OdinConfig()
	txBuilder.SetMemo(odinConfig.Memo)
	fee := sdk.NewCoins(sdk.NewCoin(odinConfig.Denom, sdk.NewIntFromBigInt(odinConfig.GasPrice)))
	txBuilder.SetFeeAmount(fee)
	txBuilder.SetGasLimit(odinConfig.GasLimit.Uint64())

	signerAddress, signerPK := c.config.OdinSigner()
	signer, err := c.getAccount(signerAddress.String())
	if err != nil {
		return errors.Wrapf(err, "failed to get signer account: %s", signerAddress.String())
	}

	sequence := signer.GetSequence()
	accountNumber := signer.GetAccountNumber()

	withdrawalAmount := sdk.NewCoins(sdk.NewCoin(odinConfig.Denom, sdk.NewIntFromBigInt(amount)))
	receiverAddress, err := sdk.AccAddressFromHex(address)
	if err != nil {
		return errors.Wrapf(err, "failed to parse receiver address: %s", address)
	}

	msg := odinminttypes.NewMsgWithdrawCoinsToAccFromTreasury(withdrawalAmount, signerAddress, receiverAddress)

	if err := txBuilder.SetMsgs(&msg); err != nil {
		return errors.Wrapf(err, "failed to set transaction builder message: %s", msg.String())
	}

	signV2 := signing.SignatureV2{
		PubKey: signerPK.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encoding.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: signer.GetSequence(),
	}
	err = txBuilder.SetSignatures(signV2)
	if err != nil {
		return errors.Wrap(err, "failed to set transaction builder signatures")
	}

	signerData := sdkauthsigning.SignerData{
		ChainID:       odinConfig.ChainId,
		AccountNumber: accountNumber,
		Sequence:      sequence,
	}

	signV2, err = sdktxclient.SignWithPrivKey(
		encoding.TxConfig.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder,
		signerPK,
		encoding.TxConfig,
		sequence,
	)
	if err != nil {
		return errors.Wrap(err, "failed to sign with private key")
	}

	err = txBuilder.SetSignatures(signV2)
	if err != nil {
		return errors.Wrap(err, "failed to set transaction builder signatures")
	}

	txBytes, err := encoding.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return errors.Wrap(err, "failed to get transaction bytes")
	}

	serviceClient := tx.NewServiceClient(c.connection)
	resp, err := serviceClient.BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to broadcast transaction")
	}
	if resp.TxResponse.Code != 200 {
		return errors.Errorf("failed to withdraw coins from minting module: %s", resp.TxResponse.Info)
	}

	return nil
}

// GetExchangeRate returns rate of assets.
func (c client) GetExchangeRate(key string) (*big.Int, error) {
	// TODO: implement logic
	return big.NewInt(1), nil
}

// getAccount returns odin account address.
func (c client) getAccount(address string) (sdkauthtypes.AccountI, error) {
	authClient := sdkauthtypes.NewQueryClient(c.connection)
	response, err := authClient.Account(context.TODO(), &sdkauthtypes.QueryAccountRequest{Address: address})
	if err != nil {
		return nil, errors.Wrap(err, "failed to query account")
	}

	var account sdkauthtypes.AccountI
	if err := encoding.Marshaler.UnpackAny(response.Account, &account); err != nil {
		return nil, errors.Wrap(err, "failed to parse query response")
	}

	return account, nil
}
