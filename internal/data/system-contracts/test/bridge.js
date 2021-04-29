const Bridge = artifacts.require('Bridge.sol');
const WETH = artifacts.require('ERC20Mock.sol');

const BigNumber = require('bn.js');

contract('Bridge.sol', (accounts) => {
  const USER = accounts[0];

  let bridge;
  let weth;

  before(async () => {
    weth = await WETH.new('WETH', 'WETH', {from: USER});
    bridge = await Bridge.new([weth.address], new BigNumber(0), {from: USER});
  });

  describe('depositETH()', async () => {
    it('should deposit ether successfully', async () => {
      const depositAmount = new BigNumber('10000000');
      const odinAddress = 'odin';
      const result = await bridge.depositETH(odinAddress, {from: USER, value: depositAmount});

      assert.equal(result.logs.length, 1);
      assert.equal(result.logs[0].event, 'ETHDeposited');
      assert.equal(result.logs[0].args._userAddress, USER);
      assert.equal(result.logs[0].args._odinAddress, odinAddress);
      assert.equal(result.logs[0].args._depositAmount.toString(), depositAmount.toString());
    });
  });

  describe('depositToken()', async () => {
    it('should deposit erc20 compatible tokens successfully', async () => {
      const depositAmount = new BigNumber('10000000');
      const odinAddress = 'odin';
      await weth.approve(bridge.address, depositAmount);
      const result = await bridge.depositERC20(weth.address, odinAddress, depositAmount, {from: USER});

      assert.equal(result.logs.length, 1);
      assert.equal(result.logs[0].event, 'ERC20Deposited');
      assert.equal(result.logs[0].args._userAddress, USER);
      assert.equal(result.logs[0].args._tokenAddress, weth.address);
      assert.equal(result.logs[0].args._odinAddress, odinAddress);
      assert.equal(result.logs[0].args._depositAmount.toString(), depositAmount.toString());
    });
  });
});
