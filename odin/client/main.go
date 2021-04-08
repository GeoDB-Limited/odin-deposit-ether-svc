package client

import (
	"context"
	app "github.com/GeoDB-Limited/odin-core/app"
	odinapp "github.com/GeoDB-Limited/odin-core/app"
	odinminttypes "github.com/GeoDB-Limited/odin-core/x/mint/types"
	"github.com/GeoDB-Limited/odin-deposit-ether-svc/internal/config"
	sdktxclient "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	sdkauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	sdkauthtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"io/ioutil"
	"math/big"
)

var encoding = odinapp.MakeEncodingConfig()

// Client defines an interface for the wrapped cosmos sdk service client.
type Client interface {
	WithSigner()

	SetBridgeAddress(common.Address) error
	GetBridgeAddress() (common.Address, error)
	ClaimWithdrawal(string, *big.Int) error
	GetExchangeRate(string) (*big.Int, error)
}

// client defines typed wrapper for the cosmos sdk service client.
type client struct {
	connection *grpc.ClientConn
	config     config.Config
	signer     *signer
}

// signer defines data to sign the transactions
type signer struct {
	address    sdk.AccAddress
	privateKey *secp256k1.PrivKey
	number     uint64
	sequence   uint64
}

// New creates a client that uses the given cosmos sdk service client.
func New(cfg config.Config) Client {
	odinConfig := cfg.OdinConfig()
	conn, err := grpc.Dial(odinConfig.Endpoint, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(errors.Wrapf(err, "failed to dial %s", odinConfig.Endpoint))
	}

	return &client{
		connection: conn,
		config:     cfg,
	}
}

// WithSigner initializes odin signer to broadcast transactions
func (c *client) WithSigner() {
	app.SetBech32AddressPrefixesAndBip44CoinType(sdk.GetConfig())
	address, pk := c.config.OdinSigner()

	authClient := sdkauthtypes.NewQueryClient(c.connection)
	response, err := authClient.Account(context.TODO(), &sdkauthtypes.QueryAccountRequest{Address: address.String()})
	if err != nil {
		panic(errors.Wrap(err, "failed to query account"))
	}

	var account sdkauthtypes.AccountI
	if err := encoding.Marshaler.UnpackAny(response.Account, &account); err != nil {
		panic(errors.Wrap(err, "failed to parse query response"))
	}

	c.signer = &signer{
		address:    address,
		privateKey: pk,
		number:     account.GetAccountNumber(),
		sequence:   account.GetSequence(),
	}
}

// SetBridgeAddress sets an address of the bridge contract to the storage.
func (c *client) SetBridgeAddress(address common.Address) error {
	if err := ioutil.WriteFile(c.config.BridgeAddressStorage(), address.Bytes(), 0777); err != nil {
		return errors.Wrap(err, "failed to add the address to the storage")
	}
	return nil
}

// GetBridgeAddress returns an address of the bridge contract.
func (c *client) GetBridgeAddress() (common.Address, error) {
	mintClient := odinminttypes.NewQueryClient(c.connection)
	response, err := mintClient.EthIntegrationAddress(context.TODO(), &odinminttypes.QueryEthIntegrationAddressRequest{})
	if err != nil {
		return common.Address{}, errors.Wrap(err, "failed to query ethereum integration address")
	}

	return common.HexToAddress(response.EthIntegrationAddress), nil
}

// ClaimWithdrawal claims minting from Odin
func (c *client) ClaimWithdrawal(address string, amount *big.Int) error {
	withdrawalAmount := sdk.NewCoins(sdk.NewCoin(c.config.OdinConfig().Denom, sdk.NewIntFromBigInt(amount)))
	receiverAddress, err := sdk.AccAddressFromHex(address)
	if err != nil {
		return errors.Wrapf(err, "failed to parse receiver address: %s", address)
	}

	msg := odinminttypes.NewMsgWithdrawCoinsToAccFromTreasury(withdrawalAmount, c.signer.address, receiverAddress)
	txBytes, err := c.signTx(&msg)
	if err != nil {
		return errors.Wrapf(err, "failed to sign the transaction to to claim withdrawal with message: %s", msg.String())
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

// signTx signs the transaction with the given message
func (c *client) signTx(msg sdk.Msg) ([]byte, error) {
	txBuilder := encoding.TxConfig.NewTxBuilder()
	odinConfig := c.config.OdinConfig()
	txBuilder.SetMemo(odinConfig.Memo)
	fee := sdk.NewCoins(sdk.NewCoin(odinConfig.Denom, sdk.NewIntFromBigInt(odinConfig.GasPrice)))
	txBuilder.SetFeeAmount(fee)
	txBuilder.SetGasLimit(odinConfig.GasLimit.Uint64())

	if err := txBuilder.SetMsgs(msg); err != nil {
		return nil, errors.Wrapf(err, "failed to set transaction builder message: %s", msg.String())
	}

	signV2 := signing.SignatureV2{
		PubKey: c.signer.privateKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode:  encoding.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: c.signer.sequence,
	}
	if err := txBuilder.SetSignatures(signV2); err != nil {
		return nil, errors.Wrap(err, "failed to set transaction builder signatures")
	}

	signerData := sdkauthsigning.SignerData{
		ChainID:       odinConfig.ChainId,
		AccountNumber: c.signer.number,
		Sequence:      c.signer.sequence,
	}

	signV2, err := sdktxclient.SignWithPrivKey(
		encoding.TxConfig.SignModeHandler().DefaultMode(),
		signerData,
		txBuilder,
		c.signer.privateKey,
		encoding.TxConfig,
		c.signer.sequence,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign with private key")
	}

	err = txBuilder.SetSignatures(signV2)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set transaction builder signatures")
	}

	txBytes, err := encoding.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction bytes")
	}

	return txBytes, nil
}

// GetExchangeRate returns rate of assets.
func (c *client) GetExchangeRate(key string) (*big.Int, error) {
	// TODO: implement logic
	return big.NewInt(1), nil
}
