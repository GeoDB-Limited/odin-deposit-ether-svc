module.exports = {
  compilers: {
    solc: {
      version: '0.7.1',
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
