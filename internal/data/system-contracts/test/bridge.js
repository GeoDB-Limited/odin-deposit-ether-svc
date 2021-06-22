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
    const REFUND_FEE = bn(10);

    before(async () => {
        weth = await WETH.new('WETH', 'WETH', {from: OWNER});
        bridge = await Bridge.new([weth.address], REFUND_FEE, true, true, false, {from: OWNER});

        await weth.transfer(USER, bn('10000000000000'), {from: OWNER});

        await reverter.snapshot();
    });

    afterEach('revert', reverter.revert);

    describe('depositETH()', async () => {
        it('should deposit ether successfully', async () => {
            const depositAmount = bn('10000000');
            const result = await bridge.depositETH(ODIN_ADDRESS, {from: USER, value: depositAmount});

            assert.equal(result.logs.length, 1);
            assert.equal(result.logs[0].event, 'ETHDeposited');
            assert.equal(result.logs[0].args._userAddress, USER);
            assert.equal(result.logs[0].args._odinAddress, ODIN_ADDRESS);
            assert.equal(result.logs[0].args._depositAmount.toString(), depositAmount.sub(REFUND_FEE).toString());
        });
    });

    describe('depositERC20()', async () => {
        it('should deposit erc20 compatible tokens successfully', async () => {
            const depositAmount = bn('10000000');
            await weth.approve(bridge.address, depositAmount, {from: USER});
            const result = await bridge.depositERC20(
                weth.address,
                ODIN_ADDRESS,
                depositAmount,
                {from: USER, value: REFUND_FEE}
            );

            assert.equal(result.logs.length, 1);
            assert.equal(result.logs[0].event, 'ERC20Deposited');
            assert.equal(result.logs[0].args._userAddress, USER);
            assert.equal(result.logs[0].args._tokenAddress, weth.address);
            assert.equal(result.logs[0].args._odinAddress, ODIN_ADDRESS);
            assert.equal(result.logs[0].args._depositAmount.toString(), depositAmount.toString());
        });
    });

    describe('claimLockedETH()', async () => {
        it('should not be possible to claim', async () => {
            const depositAmount = bn('10000000');
            await bridge.depositETH(ODIN_ADDRESS, {from: USER, value: depositAmount});
            await truffleAssert.reverts(
                bridge.claimLockedETH(depositAmount.sub(REFUND_FEE), {from: USER}),
                'It is not allowed to claim locked funds.'
            );
        });

        it('should be possible to claim locked ETH', async () => {
            await bridge.setAllowanceToClaimLockedFunds(true, {from: OWNER})

            const userBalanceBeforeDeposit = bn(await web3.eth.getBalance(USER));

            const depositAmount = bn('10000000');
            let result = await bridge.depositETH(ODIN_ADDRESS, {from: USER, value: depositAmount});

            let txFee = await calculateTxPrice(result)
            const userBalanceAfterDeposit = bn(await web3.eth.getBalance(USER));

            assert.equal(
                userBalanceBeforeDeposit.toString(),
                userBalanceAfterDeposit.add(depositAmount).add(txFee).toString(),
            );

            result = await bridge.claimLockedETH(depositAmount.sub(REFUND_FEE), {from: USER});
            txFee = await calculateTxPrice(result)

            const userBalanceAfterClaim = bn(await web3.eth.getBalance(USER));
            assert.equal(
                userBalanceAfterClaim.toString(),
                userBalanceAfterDeposit.add(depositAmount).sub(REFUND_FEE).sub(txFee).toString(),
            );
        });
    });

    describe('claimLockedERC20()', async () => {
        it('should not be possible to claim', async () => {
            const depositAmount = bn('10000000');
            await weth.approve(bridge.address, depositAmount, {from: USER});
            await bridge.depositERC20(
                weth.address,
                ODIN_ADDRESS,
                depositAmount,
                {from: USER, value: REFUND_FEE}
            );
            await truffleAssert.reverts(
                bridge.claimLockedERC20(depositAmount, weth.address, {from: USER}),
                'It is not allowed to claim locked funds.'
            );
        });

        it('should be possible to claim locked ERC20', async () => {
            await bridge.setAllowanceToClaimLockedFunds(true, {from: OWNER})

            const userBalanceBeforeDeposit = await weth.balanceOf(USER);
            const depositAmount = bn('10000000');
            await weth.approve(bridge.address, depositAmount, {from: USER});
            await bridge.depositERC20(
                weth.address,
                ODIN_ADDRESS,
                depositAmount,
                {from: USER, value: REFUND_FEE}
            );

            const userBalanceAfterDeposit = await weth.balanceOf(USER);

            assert.equal(
                userBalanceAfterDeposit.toString(),
                userBalanceBeforeDeposit.sub(depositAmount).toString(),
            );

            await bridge.claimLockedERC20(depositAmount, weth.address, {from: USER});

            const userBalanceAfterClaim = await weth.balanceOf(USER);
            assert.equal(
                userBalanceAfterClaim.toString(),
                userBalanceAfterDeposit.add(depositAmount).toString(),
            );
        });
    });

    describe('claimContractETH()', async () => {
        it('should not be possible to claim contract ETH', async () => {
            const depositAmount = bn('10000000');
            await bridge.depositETH(ODIN_ADDRESS, {from: USER, value: depositAmount});

            await truffleAssert.reverts(
                bridge.claimContractETH(depositAmount.sub(REFUND_FEE), {from: USER}),
                'Ownable: caller is not the owner.',
            );
        });

        it('should be possible to claim contract ETH by OWNER', async () => {
            const depositAmount = bn('10000000');
            await bridge.depositETH(ODIN_ADDRESS, {from: USER, value: depositAmount});

            const ownerBalanceBeforeClaim = bn(await web3.eth.getBalance(OWNER));
            const result = await bridge.claimContractETH(depositAmount.sub(REFUND_FEE), {from: OWNER});
            const txFee = await calculateTxPrice(result);

            const ownerBalanceAfterClaim = bn(await web3.eth.getBalance(OWNER));
            assert.equal(
                ownerBalanceAfterClaim.toString(),
                ownerBalanceBeforeClaim.add(depositAmount).sub(txFee).sub(REFUND_FEE).toString(),
            );
        });
    });

    describe('claimContractERC20()', async () => {
        it('should not be possible to claim contract ERC20', async () => {
            const depositAmount = bn('10000000');
            await weth.approve(bridge.address, depositAmount, {from: USER});
            await bridge.depositERC20(
                weth.address,
                ODIN_ADDRESS,
                depositAmount,
                {from: USER, value: REFUND_FEE}
            );


            await truffleAssert.reverts(
                bridge.claimContractERC20(depositAmount, weth.address, {from: USER}),
                'Ownable: caller is not the owner.',
            );
        });

        it('should be possible to claim contract ERC20 by OWNER', async () => {
            const depositAmount = bn('10000000');
            await weth.approve(bridge.address, depositAmount, {from: USER});
            await bridge.depositERC20(
                weth.address,
                ODIN_ADDRESS,
                depositAmount,
                {from: USER, value: REFUND_FEE}
            );

            const ownerBalanceBeforeClaim = bn(await weth.balanceOf(OWNER));
            await bridge.claimContractERC20(depositAmount, weth.address, {from: OWNER});

            const ownerBalanceAfterClaim = bn(await weth.balanceOf(OWNER));
            assert.equal(
                ownerBalanceAfterClaim.toString(),
                ownerBalanceBeforeClaim.add(depositAmount).toString(),
            );
        });
    });
});

async function calculateTxPrice(result) {
    const tx = await web3.eth.getTransaction(result.tx)
    return bn(tx.gasPrice).mul(bn(result.receipt.gasUsed));
}

function bn(value) {
    return new BigNumber(value);
}
