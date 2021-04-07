# ETH & ERC20 Deposit service
ETH & ERC20 deposit service is a bridge between Odin and Ethereum blockchain which allows
to deposit ether and tokens into Odin from Ethereum blockchain. It listens for smart contract events.

## Usage

Environmental variable `CONFIG` must be set and contain path to desired config file.

### To deploy a new Solidity cotract:

```bash
bridge run deploy
```

### To run deposit service:
```bash
bridge run deposit
```
