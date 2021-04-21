module.exports = {
  compilers: {
    solc: {
      version: '0.8.3',
      docker: false,
      settings: {
        optimizer: {
          enabled: true,
          runs: 100000,
        },
        evmVersion: 'istanbul',
      },
    },
  },
};
