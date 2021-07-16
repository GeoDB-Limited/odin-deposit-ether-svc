const HDWalletProvider = require('@truffle/hdwallet-provider');

module.exports = {
  compilers: {
    solc: {
      version: '0.8.3',
    },
  },
  plugins: ['truffle-plugin-verify'],
  api_keys: {
    bscscan: 'GMGBA1M5WIFX65V35YHEIREKUMKWDUEB5B',
  },
  networks: {
    bsc: {
      provider: () =>
        new HDWalletProvider(
          '1168dae9eddef3b839e4d83409021dcea5f0a98ae988c151c090ae570c1c0bda', // test account
          'wss://bsc-dataseed.binance.org',
        ),
      network_id: 56,
      gas: 7000000,
      gasPrice: 30000000000, // 30 gwei
      skipDryRun: true,
    },
  },
};
