const Bridge = artifacts.require('Bridge.sol');
const WETH = artifacts.require('ERC20Mock.sol');

const BigNumber = require('bn.js');

contract('Bridge.sol', (accounts) => {
    const USER = accounts[0];

    let bridge;
    let weth;

    const REFUND_FEE = new BigNumber(10)
    let gasPrice;

    before(async () => {
        weth = await WETH.new('WETH', 'WETH', {from: USER});
        bridge = await Bridge.new([weth.address], REFUND_FEE, true, true, false, {from: USER});
        gasPrice = new BigNumber(bridge.constructor.class_defaults.gasPrice);
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
            assert.equal(result.logs[0].args._depositAmount.toString(), depositAmount.sub(REFUND_FEE).toString());
        });
    });

    describe('depositToken()', async () => {
        it('should deposit erc20 compatible tokens successfully', async () => {
            const depositAmount = new BigNumber('10000000');
            const odinAddress = 'odin';
            await weth.approve(bridge.address, depositAmount);
            const result = await bridge.depositERC20(
                weth.address,
                odinAddress,
                depositAmount,
                {from: USER, value: REFUND_FEE}
            );

            assert.equal(result.logs.length, 1);
            assert.equal(result.logs[0].event, 'ERC20Deposited');
            assert.equal(result.logs[0].args._userAddress, USER);
            assert.equal(result.logs[0].args._tokenAddress, weth.address);
            assert.equal(result.logs[0].args._odinAddress, odinAddress);
            assert.equal(result.logs[0].args._depositAmount.toString(), depositAmount.toString());
        });
    });

    describe('claimLockedETH()', async () => {
        it('should be possible to claim locked ETH', async () => {
            await bridge.setAllowanceToClaimLockedFunds(true, {from: USER})

            const userBalanceBeforeDeposit = new BigNumber(await web3.eth.getBalance(USER));
            const depositAmount = new BigNumber('10000000');
            const odinAddress = 'odin';
            let result = await bridge.depositETH(odinAddress, {from: USER, value: depositAmount});

            let txFee = new BigNumber(result.receipt.gasUsed).mul(gasPrice);
            const userBalanceAfterDeposit = new BigNumber(await web3.eth.getBalance(USER));

            assert.equal(
                userBalanceAfterDeposit.toString(),
                userBalanceBeforeDeposit.sub(depositAmount).sub(txFee).toString(),
            );

            result = await bridge.claimLockedETH(depositAmount.sub(REFUND_FEE));
            txFee = new BigNumber(result.receipt.gasUsed).mul(gasPrice);

            const userBalanceAfterClaim = new BigNumber(await web3.eth.getBalance(USER));
            assert.equal(
                userBalanceAfterClaim.toString(),
                userBalanceAfterDeposit.add(depositAmount.sub(REFUND_FEE)).sub(txFee).toString(),
            );
        });
    });
});
