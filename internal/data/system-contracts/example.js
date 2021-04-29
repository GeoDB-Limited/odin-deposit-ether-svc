/* eslint-disable indent */
const Web3 = require('web3');
const abi = require('./abi.json');
const BigNumber = require('bn.js');

const web3 = new Web3(new Web3.providers.HttpProvider('https://ropsten.infura.io/v3/d93dde4baecb4c0d9ddb0e8b4ca2fa9e'));

const depositAmount = web3.utils.toHex(web3.utils.toWei('0.01', 'ether'));
const senderAddress = '0xe11a6E1eb2D61590fEBD8d694e39C4428e681815';
const contractAddress = '0x5D9a9340b8249b7b1bbDFa124897117595b942aE';
const odinAddress = 'odin1nnfeguq30x6nwxjhaypxymx3nulyspsuja4a2x';

const gasPrice = web3.utils.toHex(web3.utils.toWei('1', 'gwei'));
const gasLimit = new BigNumber('6000000');

const Bridge = new web3.eth.Contract(abi, contractAddress, {
    from: senderAddress, // default from address
    gasPrice: gasPrice, // default gas price in wei, 20 gwei in this case
    gasLimit: gasLimit,
});

const privateKey = '444bec714882856cb695ddba7098862594ecafe69494d88c370deaab425cd6de';
const account = web3.eth.accounts.privateKeyToAccount('0x' + privateKey);
web3.eth.accounts.wallet.add(account);
web3.eth.defaultAccount = account.address;

async function main() {
    await Bridge.methods.depositETH(odinAddress).send({
        from: senderAddress,
        value: depositAmount,
        gasPrice: gasPrice,
        gasLimit: gasLimit,
    }).then(function(receipt) {
        console.log(receipt);
    });
}

main();
