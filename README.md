# ETH & ERC20 Deposit service
ETH & ERC20 deposit service is a bridge between Odin and Ethereum blockchain which allows
to deposit ether and tokens into Odin from Ethereum blockchain. It listens for smart contract events.

## Usage

Environmental variable `KV_VIPER_FILE` must be set and contain path to desired config file.

### To deploy a new Solidity cotract:

```bash
odin-deposit-ether-svc run deploy
```

### To run deposit service:
```bash
odin-deposit-ether-svc run deposit
```
