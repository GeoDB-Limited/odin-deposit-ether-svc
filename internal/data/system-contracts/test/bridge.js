const Bridge = artifacts.require('Bridge.sol');

const BigNumber = require('bn.js');

contract('Bridge.sol', (accounts) => {
  const USER = accounts[0];

  let bridge;

  before(async () => {
    bridge = await Bridge.new([]);
  });

  describe('deposit()', async () => {
    it('should deposit ether successfully', async () => {
      const depositAmount = new BigNumber('10000000');
      const odinAddress = 'odin';
      const result = await bridge.depositEther(odinAddress, {from: USER, value: depositAmount});

      assert.equal(result.logs.length, 1);
      assert.equal(result.logs[0].event, 'EtherDeposited');
      assert.equal(result.logs[0].args._userAddress, USER);
      assert.equal(result.logs[0].args._odinAddress, odinAddress);
      assert.equal(result.logs[0].args._depositAmount.toString(), depositAmount.toString());
    });
  });
});
