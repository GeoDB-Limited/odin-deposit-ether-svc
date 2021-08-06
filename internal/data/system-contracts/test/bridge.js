const Bridge = artifacts.require('Bridge.sol');
const WETH = artifacts.require('ERC20Mock.sol');

const truffleAssert = require('truffle-assertions');
const Reverter = require('./helpers/reverter');
const BigNumber = require('bn.js');

contract('Bridge.sol', (accounts) => {
  const OWNER = accounts[0];
  const USER = accounts[1];

  let bridge;
  let weth;

  const reverter = new Reverter(web3);
  const ODIN_ADDRESS = 'odin';
  const REFUND_GAS_LIMIT = bn(1000000);

  before(async () => {
    weth = await WETH.new('WETH', 'WETH', {from: OWNER});
    bridge = await Bridge.new([weth.address], REFUND_GAS_LIMIT, true, true, false, {from: OWNER});

    await weth.transfer(USER, bn('10000000000000'), {from: OWNER});

    await reverter.snapshot();
  });

  afterEach('revert', reverter.revert);

  describe('deposit()', async () => {
    it('should deposit erc20 compatible tokens successfully', async () => {
      const depositAmount = bn('10000000');
      await weth.approve(bridge.address, depositAmount, {from: USER});
      const result = await bridge.deposit(
        ODIN_ADDRESS,
        weth.address,
        depositAmount,
        {from: USER},
      );

      assert.equal(result.logs.length, 2);
      assert.equal(result.logs[1].event, 'TokensDeposited');
      assert.equal(result.logs[1].args._userAddress, USER);
      assert.equal(result.logs[1].args._tokenAddress, weth.address);
      assert.equal(result.logs[1].args._odinAddress, ODIN_ADDRESS);
      assert.equal(result.logs[1].args._depositAmount.toString(), depositAmount.toString());
    });
  });

  describe('claimLockedTokens()', async () => {
    it('should not be possible to claim', async () => {
      const depositAmount = bn('10000000');
      await weth.approve(bridge.address, depositAmount, {from: USER});
      await bridge.deposit(
        ODIN_ADDRESS,
        weth.address,
        depositAmount,
        {from: USER},
      );
      await truffleAssert.reverts(
        bridge.claimLockedTokens(weth.address, {from: USER}),
        'It is not allowed to claim locked funds.',
      );
    });

    it('should be possible to claim locked ERC20', async () => {
      await bridge.setAllowanceToClaimLockedFunds(true, {from: OWNER});

      const userBalanceBeforeDeposit = await weth.balanceOf(USER);
      const depositAmount = bn('10000000');
      await weth.approve(bridge.address, depositAmount, {from: USER});
      await bridge.deposit(
        ODIN_ADDRESS,
        weth.address,
        depositAmount,
        {from: USER},
      );

      const userBalanceAfterDeposit = await weth.balanceOf(USER);

      assert.equal(
        userBalanceAfterDeposit.toString(),
        userBalanceBeforeDeposit.sub(depositAmount).toString(),
      );

      await bridge.claimLockedTokens(weth.address, {from: USER});

      const userBalanceAfterClaim = await weth.balanceOf(USER);
      assert.equal(
        userBalanceAfterClaim.toString(),
        userBalanceAfterDeposit.add(depositAmount).toString(),
      );
    });
  });

  describe('claimContractTokens()', async () => {
    it('should not be possible to claim contract ERC20', async () => {
      const depositAmount = bn('10000000');
      await weth.approve(bridge.address, depositAmount, {from: USER});
      await bridge.deposit(
        ODIN_ADDRESS,
        weth.address,
        depositAmount,
        {from: USER},
      );


      await truffleAssert.reverts(
        bridge.claimContractTokens(depositAmount, weth.address, {from: USER}),
        'Ownable: caller is not the owner.',
      );
    });

    it('should be possible to claim contract ERC20 by OWNER', async () => {
      const depositAmount = bn('10000000');
      await weth.approve(bridge.address, depositAmount, {from: USER});
      await bridge.deposit(
        ODIN_ADDRESS,
        weth.address,
        depositAmount,
        {from: USER},
      );

      const ownerBalanceBeforeClaim = bn(await weth.balanceOf(OWNER));
      await bridge.claimContractTokens(depositAmount, weth.address, {from: OWNER});

      const ownerBalanceAfterClaim = bn(await weth.balanceOf(OWNER));
      assert.equal(
        ownerBalanceAfterClaim.toString(),
        ownerBalanceBeforeClaim.add(depositAmount).toString(),
      );
    });
  });
});

function bn(value) {
  return new BigNumber(value);
}
