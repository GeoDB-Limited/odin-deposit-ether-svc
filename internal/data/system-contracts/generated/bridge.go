// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_supportedTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_depositCompensation\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"ERC20Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"ETHDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compensationDeposited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCompensation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"}],\"name\":\"depositETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payBackERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payBackETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setDepositCompensation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x60806040523480156200001157600080fd5b50604051620025523803806200255283398181016040528101906200003791906200028b565b600062000049620001b360201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35060005b8251811015620001a35760016003600085848151811062000134577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080806200019a90620003b1565b915050620000ea565b50806001819055505050620004a2565b600033905090565b6000620001d2620001cc846200030e565b620002e5565b90508083825260208201905082856020860282011115620001f257600080fd5b60005b858110156200022657816200020b888262000230565b845260208401935060208301925050600181019050620001f5565b5050509392505050565b60008151905062000241816200046e565b92915050565b600082601f8301126200025957600080fd5b81516200026b848260208601620001bb565b91505092915050565b600081519050620002858162000488565b92915050565b600080604083850312156200029f57600080fd5b600083015167ffffffffffffffff811115620002ba57600080fd5b620002c88582860162000247565b9250506020620002db8582860162000274565b9150509250929050565b6000620002f162000304565b9050620002ff82826200037b565b919050565b6000604051905090565b600067ffffffffffffffff8211156200032c576200032b6200042e565b5b602082029050602081019050919050565b60006200034a8262000351565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b62000386826200045d565b810181811067ffffffffffffffff82111715620003a857620003a76200042e565b5b80604052505050565b6000620003be8262000371565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415620003f457620003f3620003ff565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b62000479816200033d565b81146200048557600080fd5b50565b620004938162000371565b81146200049f57600080fd5b50565b6120a080620004b26000396000f3fe6080604052600436106100c25760003560e01c80639371388d1161007f578063be20e53911610059578063be20e53914610288578063c7b77eb0146102c5578063d48bfca714610302578063f2fde38b1461033f576100c2565b80639371388d146101f0578063994f5a771461022d5780639b1c48e614610258576100c2565b806334c15e28146100c75780635fa7b5841461010457806368c4ac2614610141578063715018a61461017e5780637288f11a146101955780638da5cb5b146101c5575b600080fd5b3480156100d357600080fd5b506100ee60048036038101906100e991906114c0565b610368565b6040516100fb9190611949565b60405180910390f35b34801561011057600080fd5b5061012b600480360381019061012691906114c0565b610388565b6040516101389190611949565b60405180910390f35b34801561014d57600080fd5b50610168600480360381019061016391906114c0565b6104aa565b6040516101759190611949565b60405180910390f35b34801561018a57600080fd5b506101936104ca565b005b6101af60048036038101906101aa9190611538565b610604565b6040516101bc9190611949565b60405180910390f35b3480156101d157600080fd5b506101da6109af565b6040516101e791906118ce565b60405180910390f35b3480156101fc57600080fd5b50610217600480360381019061021291906114e9565b6109d8565b6040516102249190611949565b60405180910390f35b34801561023957600080fd5b50610242610c39565b60405161024f9190611b39565b60405180910390f35b610272600480360381019061026d9190611604565b610c3f565b60405161027f9190611949565b60405180910390f35b34801561029457600080fd5b506102af60048036038101906102aa9190611686565b610dec565b6040516102bc9190611949565b60405180910390f35b3480156102d157600080fd5b506102ec60048036038101906102e7919061159f565b610e7a565b6040516102f99190611949565b60405180910390f35b34801561030e57600080fd5b50610329600480360381019061032491906114c0565b6110b5565b6040516103369190611949565b60405180910390f35b34801561034b57600080fd5b50610366600480360381019061036191906114c0565b6111d7565b005b60026020528060005260406000206000915054906101000a900460ff1681565b6000610392611380565b73ffffffffffffffffffffffffffffffffffffffff166103b06109af565b73ffffffffffffffffffffffffffffffffffffffff1614610406576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103fd90611a59565b60405180910390fd5b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd360405160405180910390a260019050919050565b60036020528060005260406000206000915054906101000a900460ff1681565b6104d2611380565b73ffffffffffffffffffffffffffffffffffffffff166104f06109af565b73ffffffffffffffffffffffffffffffffffffffff1614610546576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053d90611a59565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60006106258473ffffffffffffffffffffffffffffffffffffffff16611388565b610664576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065b90611af9565b60405180910390fd5b600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166106f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e790611ab9565b60405180910390fd5b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166107df57600154341015610786576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077d90611a79565b60405180910390fd5b6001600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b600084905060008173ffffffffffffffffffffffffffffffffffffffff166323b872dd3330876040518463ffffffff1660e01b8152600401610823939291906118e9565b602060405180830381600087803b15801561083d57600080fd5b505af1158015610851573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087591906115db565b9050806108b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ae90611ad9565b60405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f86b8db4a61366e0537c773e8b41721cd695d4830de7ffd5c1229cf296d1565f887878673ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b15801561094e57600080fd5b505afa158015610962573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061098b9190611645565b60405161099a93929190611994565b60405180910390a36001925050509392505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006109e2611380565b73ffffffffffffffffffffffffffffffffffffffff16610a006109af565b73ffffffffffffffffffffffffffffffffffffffff1614610a56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4d90611a59565b60405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86856040518363ffffffff1660e01b8152600401610a93929190611920565b602060405180830381600087803b158015610aad57600080fd5b505af1158015610ac1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ae591906115db565b905080610b27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1e90611a39565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600154604051610b4d906118b9565b60006040518083038185875af1925050503d8060008114610b8a576040519150601f19603f3d011682016040523d82523d6000602084013e610b8f565b606091505b50508091505080610bd5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bcc90611b19565b60405180910390fd5b6000600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060019150509392505050565b60015481565b600080349050600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610d4f576000600154905080821015610ce0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cd7906119d9565b60405180910390fd5b6001600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550610d4b818361139b90919063ffffffff16565b9150505b60008111610d92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d8990611a19565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167fa5ee847b808915e440e213630058643951dfa817cd585254162b8e69b71fa4328483604051610dda929190611964565b60405180910390a26001915050919050565b6000610df6611380565b73ffffffffffffffffffffffffffffffffffffffff16610e146109af565b73ffffffffffffffffffffffffffffffffffffffff1614610e6a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6190611a59565b60405180910390fd5b8160018190555060019050919050565b6000610e84611380565b73ffffffffffffffffffffffffffffffffffffffff16610ea26109af565b73ffffffffffffffffffffffffffffffffffffffff1614610ef8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eef90611a59565b60405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff1683604051610f1e906118b9565b60006040518083038185875af1925050503d8060008114610f5b576040519150601f19603f3d011682016040523d82523d6000602084013e610f60565b606091505b5050905080610fa4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9b90611a99565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600154604051610fca906118b9565b60006040518083038185875af1925050503d8060008114611007576040519150601f19603f3d011682016040523d82523d6000602084013e61100c565b606091505b50508091505080611052576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161104990611b19565b60405180910390fd5b6000600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600191505092915050565b60006110bf611380565b73ffffffffffffffffffffffffffffffffffffffff166110dd6109af565b73ffffffffffffffffffffffffffffffffffffffff1614611133576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161112a90611a59565b60405180910390fd5b6001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a460405160405180910390a260019050919050565b6111df611380565b73ffffffffffffffffffffffffffffffffffffffff166111fd6109af565b73ffffffffffffffffffffffffffffffffffffffff1614611253576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161124a90611a59565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156112c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112ba906119f9565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b600080823b905060008111915050919050565b600081836113a99190611bd1565b905092915050565b60006113c46113bf84611b79565b611b54565b9050828152602081018484840111156113dc57600080fd5b6113e7848285611c4d565b509392505050565b60006114026113fd84611b79565b611b54565b90508281526020810184848401111561141a57600080fd5b611425848285611c5c565b509392505050565b60008135905061143c81612025565b92915050565b6000815190506114518161203c565b92915050565b600082601f83011261146857600080fd5b81356114788482602086016113b1565b91505092915050565b600082601f83011261149257600080fd5b81516114a28482602086016113ef565b91505092915050565b6000813590506114ba81612053565b92915050565b6000602082840312156114d257600080fd5b60006114e08482850161142d565b91505092915050565b6000806000606084860312156114fe57600080fd5b600061150c8682870161142d565b935050602061151d8682870161142d565b925050604061152e868287016114ab565b9150509250925092565b60008060006060848603121561154d57600080fd5b600061155b8682870161142d565b935050602084013567ffffffffffffffff81111561157857600080fd5b61158486828701611457565b9250506040611595868287016114ab565b9150509250925092565b600080604083850312156115b257600080fd5b60006115c08582860161142d565b92505060206115d1858286016114ab565b9150509250929050565b6000602082840312156115ed57600080fd5b60006115fb84828501611442565b91505092915050565b60006020828403121561161657600080fd5b600082013567ffffffffffffffff81111561163057600080fd5b61163c84828501611457565b91505092915050565b60006020828403121561165757600080fd5b600082015167ffffffffffffffff81111561167157600080fd5b61167d84828501611481565b91505092915050565b60006020828403121561169857600080fd5b60006116a6848285016114ab565b91505092915050565b6116b881611c05565b82525050565b6116c781611c17565b82525050565b60006116d882611baa565b6116e28185611bc0565b93506116f2818560208601611c5c565b6116fb81611d1e565b840191505092915050565b6000611713602a83611bc0565b915061171e82611d2f565b604082019050919050565b6000611736602683611bc0565b915061174182611d7e565b604082019050919050565b6000611759604583611bc0565b915061176482611dcd565b606082019050919050565b600061177c601283611bc0565b915061178782611e42565b602082019050919050565b600061179f602083611bc0565b91506117aa82611e6b565b602082019050919050565b60006117c2602b83611bc0565b91506117cd82611e94565b604082019050919050565b60006117e5602683611bc0565b91506117f082611ee3565b604082019050919050565b6000611808602583611bc0565b915061181382611f32565b604082019050919050565b600061182b601a83611bc0565b915061183682611f81565b602082019050919050565b600061184e600083611bb5565b915061185982611faa565b600082019050919050565b6000611871601d83611bc0565b915061187c82611fad565b602082019050919050565b6000611894602f83611bc0565b915061189f82611fd6565b604082019050919050565b6118b381611c43565b82525050565b60006118c482611841565b9150819050919050565b60006020820190506118e360008301846116af565b92915050565b60006060820190506118fe60008301866116af565b61190b60208301856116af565b61191860408301846118aa565b949350505050565b600060408201905061193560008301856116af565b61194260208301846118aa565b9392505050565b600060208201905061195e60008301846116be565b92915050565b6000604082019050818103600083015261197e81856116cd565b905061198d60208301846118aa565b9392505050565b600060608201905081810360008301526119ae81866116cd565b90506119bd60208301856118aa565b81810360408301526119cf81846116cd565b9050949350505050565b600060208201905081810360008301526119f281611706565b9050919050565b60006020820190508181036000830152611a1281611729565b9050919050565b60006020820190508181036000830152611a328161174c565b9050919050565b60006020820190508181036000830152611a528161176f565b9050919050565b60006020820190508181036000830152611a7281611792565b9050919050565b60006020820190508181036000830152611a92816117b5565b9050919050565b60006020820190508181036000830152611ab2816117d8565b9050919050565b60006020820190508181036000830152611ad2816117fb565b9050919050565b60006020820190508181036000830152611af28161181e565b9050919050565b60006020820190508181036000830152611b1281611864565b9050919050565b60006020820190508181036000830152611b3281611887565b9050919050565b6000602082019050611b4e60008301846118aa565b92915050565b6000611b5e611b6f565b9050611b6a8282611c8f565b919050565b6000604051905090565b600067ffffffffffffffff821115611b9457611b93611cef565b5b611b9d82611d1e565b9050602081019050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b6000611bdc82611c43565b9150611be783611c43565b925082821015611bfa57611bf9611cc0565b5b828203905092915050565b6000611c1082611c23565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611c7a578082015181840152602081019050611c5f565b83811115611c89576000848401525b50505050565b611c9882611d1e565b810181811067ffffffffffffffff82111715611cb757611cb6611cef565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f496e73756666696369656e742066756e647320746f206465706f73697420636f60008201527f6d70656e736174696f6e00000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f496e76616c69642076616c756520666f7220746865206465706f73697420616d60008201527f6f756e742c206661696c656420746f206465706f7369742061207a65726f207660208201527f616c75652e000000000000000000000000000000000000000000000000000000604082015250565b7f4661696c656420746f20706179206261636b0000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f496e73756666696369656e742066756e647320666f72206465706f736974206360008201527f6f6d70656e736174696f6e000000000000000000000000000000000000000000602082015250565b7f4661696c656420746f20706179206261636b20746865206465706f736974206160008201527f6d6f756e742e0000000000000000000000000000000000000000000000000000602082015250565b7f556e737570706f7274656420746f6b656e2c206661696c656420746f2064657060008201527f6f7369742e000000000000000000000000000000000000000000000000000000602082015250565b7f4661696c656420746f207472616e7366657220746f6b656e732e000000000000600082015250565b50565b7f476976656e20746f6b656e206973206e6f74206120636f6e7472616374000000600082015250565b7f4661696c656420746f207061792074686520636f6d70656e736174696f6e206660008201527f6f7220706179696e67206261636b2e0000000000000000000000000000000000602082015250565b61202e81611c05565b811461203957600080fd5b50565b61204581611c17565b811461205057600080fd5b50565b61205c81611c43565b811461206757600080fd5b5056fea26469706673582212205bdedc7871c9d5db7b2a2b96b562075b8da990e13781782e18071ca6cf19ff8064736f6c63430008030033"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _supportedTokens []common.Address, _depositCompensation *big.Int) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend, _supportedTokens, _depositCompensation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// CompensationDeposited is a free data retrieval call binding the contract method 0x34c15e28.
//
// Solidity: function compensationDeposited(address ) view returns(bool)
func (_Bridge *BridgeCaller) CompensationDeposited(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "compensationDeposited", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompensationDeposited is a free data retrieval call binding the contract method 0x34c15e28.
//
// Solidity: function compensationDeposited(address ) view returns(bool)
func (_Bridge *BridgeSession) CompensationDeposited(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.CompensationDeposited(&_Bridge.CallOpts, arg0)
}

// CompensationDeposited is a free data retrieval call binding the contract method 0x34c15e28.
//
// Solidity: function compensationDeposited(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) CompensationDeposited(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.CompensationDeposited(&_Bridge.CallOpts, arg0)
}

// DepositCompensation is a free data retrieval call binding the contract method 0x994f5a77.
//
// Solidity: function depositCompensation() view returns(uint256)
func (_Bridge *BridgeCaller) DepositCompensation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "depositCompensation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositCompensation is a free data retrieval call binding the contract method 0x994f5a77.
//
// Solidity: function depositCompensation() view returns(uint256)
func (_Bridge *BridgeSession) DepositCompensation() (*big.Int, error) {
	return _Bridge.Contract.DepositCompensation(&_Bridge.CallOpts)
}

// DepositCompensation is a free data retrieval call binding the contract method 0x994f5a77.
//
// Solidity: function depositCompensation() view returns(uint256)
func (_Bridge *BridgeCallerSession) DepositCompensation() (*big.Int, error) {
	return _Bridge.Contract.DepositCompensation(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bridge *BridgeCallerSession) Owner() (common.Address, error) {
	return _Bridge.Contract.Owner(&_Bridge.CallOpts)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_Bridge *BridgeCaller) SupportedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "supportedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_Bridge *BridgeSession) SupportedTokens(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.SupportedTokens(&_Bridge.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) SupportedTokens(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.SupportedTokens(&_Bridge.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address _tokenAddress) returns(bool)
func (_Bridge *BridgeTransactor) AddToken(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addToken", _tokenAddress)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address _tokenAddress) returns(bool)
func (_Bridge *BridgeSession) AddToken(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddToken(&_Bridge.TransactOpts, _tokenAddress)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address _tokenAddress) returns(bool)
func (_Bridge *BridgeTransactorSession) AddToken(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddToken(&_Bridge.TransactOpts, _tokenAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x7288f11a.
//
// Solidity: function depositERC20(address _tokenAddress, string _odinAddress, uint256 _depositAmount) payable returns(bool)
func (_Bridge *BridgeTransactor) DepositERC20(opts *bind.TransactOpts, _tokenAddress common.Address, _odinAddress string, _depositAmount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositERC20", _tokenAddress, _odinAddress, _depositAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x7288f11a.
//
// Solidity: function depositERC20(address _tokenAddress, string _odinAddress, uint256 _depositAmount) payable returns(bool)
func (_Bridge *BridgeSession) DepositERC20(_tokenAddress common.Address, _odinAddress string, _depositAmount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20(&_Bridge.TransactOpts, _tokenAddress, _odinAddress, _depositAmount)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x7288f11a.
//
// Solidity: function depositERC20(address _tokenAddress, string _odinAddress, uint256 _depositAmount) payable returns(bool)
func (_Bridge *BridgeTransactorSession) DepositERC20(_tokenAddress common.Address, _odinAddress string, _depositAmount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DepositERC20(&_Bridge.TransactOpts, _tokenAddress, _odinAddress, _depositAmount)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9b1c48e6.
//
// Solidity: function depositETH(string _odinAddress) payable returns(bool)
func (_Bridge *BridgeTransactor) DepositETH(opts *bind.TransactOpts, _odinAddress string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositETH", _odinAddress)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9b1c48e6.
//
// Solidity: function depositETH(string _odinAddress) payable returns(bool)
func (_Bridge *BridgeSession) DepositETH(_odinAddress string) (*types.Transaction, error) {
	return _Bridge.Contract.DepositETH(&_Bridge.TransactOpts, _odinAddress)
}

// DepositETH is a paid mutator transaction binding the contract method 0x9b1c48e6.
//
// Solidity: function depositETH(string _odinAddress) payable returns(bool)
func (_Bridge *BridgeTransactorSession) DepositETH(_odinAddress string) (*types.Transaction, error) {
	return _Bridge.Contract.DepositETH(&_Bridge.TransactOpts, _odinAddress)
}

// PayBackERC20 is a paid mutator transaction binding the contract method 0x9371388d.
//
// Solidity: function payBackERC20(address _user, address _token, uint256 _amount) returns(bool)
func (_Bridge *BridgeTransactor) PayBackERC20(opts *bind.TransactOpts, _user common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "payBackERC20", _user, _token, _amount)
}

// PayBackERC20 is a paid mutator transaction binding the contract method 0x9371388d.
//
// Solidity: function payBackERC20(address _user, address _token, uint256 _amount) returns(bool)
func (_Bridge *BridgeSession) PayBackERC20(_user common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.PayBackERC20(&_Bridge.TransactOpts, _user, _token, _amount)
}

// PayBackERC20 is a paid mutator transaction binding the contract method 0x9371388d.
//
// Solidity: function payBackERC20(address _user, address _token, uint256 _amount) returns(bool)
func (_Bridge *BridgeTransactorSession) PayBackERC20(_user common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.PayBackERC20(&_Bridge.TransactOpts, _user, _token, _amount)
}

// PayBackETH is a paid mutator transaction binding the contract method 0xc7b77eb0.
//
// Solidity: function payBackETH(address _user, uint256 _amount) returns(bool)
func (_Bridge *BridgeTransactor) PayBackETH(opts *bind.TransactOpts, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "payBackETH", _user, _amount)
}

// PayBackETH is a paid mutator transaction binding the contract method 0xc7b77eb0.
//
// Solidity: function payBackETH(address _user, uint256 _amount) returns(bool)
func (_Bridge *BridgeSession) PayBackETH(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.PayBackETH(&_Bridge.TransactOpts, _user, _amount)
}

// PayBackETH is a paid mutator transaction binding the contract method 0xc7b77eb0.
//
// Solidity: function payBackETH(address _user, uint256 _amount) returns(bool)
func (_Bridge *BridgeTransactorSession) PayBackETH(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.PayBackETH(&_Bridge.TransactOpts, _user, _amount)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _tokenAddress) returns(bool)
func (_Bridge *BridgeTransactor) RemoveToken(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeToken", _tokenAddress)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _tokenAddress) returns(bool)
func (_Bridge *BridgeSession) RemoveToken(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveToken(&_Bridge.TransactOpts, _tokenAddress)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address _tokenAddress) returns(bool)
func (_Bridge *BridgeTransactorSession) RemoveToken(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveToken(&_Bridge.TransactOpts, _tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bridge *BridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bridge.Contract.RenounceOwnership(&_Bridge.TransactOpts)
}

// SetDepositCompensation is a paid mutator transaction binding the contract method 0xbe20e539.
//
// Solidity: function setDepositCompensation(uint256 _amount) returns(bool)
func (_Bridge *BridgeTransactor) SetDepositCompensation(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setDepositCompensation", _amount)
}

// SetDepositCompensation is a paid mutator transaction binding the contract method 0xbe20e539.
//
// Solidity: function setDepositCompensation(uint256 _amount) returns(bool)
func (_Bridge *BridgeSession) SetDepositCompensation(_amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetDepositCompensation(&_Bridge.TransactOpts, _amount)
}

// SetDepositCompensation is a paid mutator transaction binding the contract method 0xbe20e539.
//
// Solidity: function setDepositCompensation(uint256 _amount) returns(bool)
func (_Bridge *BridgeTransactorSession) SetDepositCompensation(_amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetDepositCompensation(&_Bridge.TransactOpts, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bridge *BridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.TransferOwnership(&_Bridge.TransactOpts, newOwner)
}

// BridgeERC20DepositedIterator is returned from FilterERC20Deposited and is used to iterate over the raw logs and unpacked data for ERC20Deposited events raised by the Bridge contract.
type BridgeERC20DepositedIterator struct {
	Event *BridgeERC20Deposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeERC20DepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeERC20Deposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeERC20Deposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeERC20DepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeERC20DepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeERC20Deposited represents a ERC20Deposited event raised by the Bridge contract.
type BridgeERC20Deposited struct {
	UserAddress   common.Address
	OdinAddress   string
	DepositAmount *big.Int
	TokenAddress  common.Address
	Symbol        string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterERC20Deposited is a free log retrieval operation binding the contract event 0x86b8db4a61366e0537c773e8b41721cd695d4830de7ffd5c1229cf296d1565f8.
//
// Solidity: event ERC20Deposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount, address indexed _tokenAddress, string symbol)
func (_Bridge *BridgeFilterer) FilterERC20Deposited(opts *bind.FilterOpts, _userAddress []common.Address, _tokenAddress []common.Address) (*BridgeERC20DepositedIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ERC20Deposited", _userAddressRule, _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgeERC20DepositedIterator{contract: _Bridge.contract, event: "ERC20Deposited", logs: logs, sub: sub}, nil
}

// WatchERC20Deposited is a free log subscription operation binding the contract event 0x86b8db4a61366e0537c773e8b41721cd695d4830de7ffd5c1229cf296d1565f8.
//
// Solidity: event ERC20Deposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount, address indexed _tokenAddress, string symbol)
func (_Bridge *BridgeFilterer) WatchERC20Deposited(opts *bind.WatchOpts, sink chan<- *BridgeERC20Deposited, _userAddress []common.Address, _tokenAddress []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ERC20Deposited", _userAddressRule, _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeERC20Deposited)
				if err := _Bridge.contract.UnpackLog(event, "ERC20Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC20Deposited is a log parse operation binding the contract event 0x86b8db4a61366e0537c773e8b41721cd695d4830de7ffd5c1229cf296d1565f8.
//
// Solidity: event ERC20Deposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount, address indexed _tokenAddress, string symbol)
func (_Bridge *BridgeFilterer) ParseERC20Deposited(log types.Log) (*BridgeERC20Deposited, error) {
	event := new(BridgeERC20Deposited)
	if err := _Bridge.contract.UnpackLog(event, "ERC20Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeETHDepositedIterator is returned from FilterETHDeposited and is used to iterate over the raw logs and unpacked data for ETHDeposited events raised by the Bridge contract.
type BridgeETHDepositedIterator struct {
	Event *BridgeETHDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeETHDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeETHDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeETHDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeETHDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeETHDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeETHDeposited represents a ETHDeposited event raised by the Bridge contract.
type BridgeETHDeposited struct {
	UserAddress   common.Address
	OdinAddress   string
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterETHDeposited is a free log retrieval operation binding the contract event 0xa5ee847b808915e440e213630058643951dfa817cd585254162b8e69b71fa432.
//
// Solidity: event ETHDeposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount)
func (_Bridge *BridgeFilterer) FilterETHDeposited(opts *bind.FilterOpts, _userAddress []common.Address) (*BridgeETHDepositedIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ETHDeposited", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgeETHDepositedIterator{contract: _Bridge.contract, event: "ETHDeposited", logs: logs, sub: sub}, nil
}

// WatchETHDeposited is a free log subscription operation binding the contract event 0xa5ee847b808915e440e213630058643951dfa817cd585254162b8e69b71fa432.
//
// Solidity: event ETHDeposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount)
func (_Bridge *BridgeFilterer) WatchETHDeposited(opts *bind.WatchOpts, sink chan<- *BridgeETHDeposited, _userAddress []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ETHDeposited", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeETHDeposited)
				if err := _Bridge.contract.UnpackLog(event, "ETHDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseETHDeposited is a log parse operation binding the contract event 0xa5ee847b808915e440e213630058643951dfa817cd585254162b8e69b71fa432.
//
// Solidity: event ETHDeposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount)
func (_Bridge *BridgeFilterer) ParseETHDeposited(log types.Log) (*BridgeETHDeposited, error) {
	event := new(BridgeETHDeposited)
	if err := _Bridge.contract.UnpackLog(event, "ETHDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bridge contract.
type BridgeOwnershipTransferredIterator struct {
	Event *BridgeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Bridge contract.
type BridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeOwnershipTransferredIterator{contract: _Bridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOwnershipTransferred)
				if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bridge *BridgeFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeOwnershipTransferred, error) {
	event := new(BridgeOwnershipTransferred)
	if err := _Bridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the Bridge contract.
type BridgeTokenAddedIterator struct {
	Event *BridgeTokenAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeTokenAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenAdded represents a TokenAdded event raised by the Bridge contract.
type BridgeTokenAdded struct {
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed _tokenAddress)
func (_Bridge *BridgeFilterer) FilterTokenAdded(opts *bind.FilterOpts, _tokenAddress []common.Address) (*BridgeTokenAddedIterator, error) {

	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenAdded", _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenAddedIterator{contract: _Bridge.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed _tokenAddress)
func (_Bridge *BridgeFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *BridgeTokenAdded, _tokenAddress []common.Address) (event.Subscription, error) {

	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenAdded", _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenAdded)
				if err := _Bridge.contract.UnpackLog(event, "TokenAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenAdded is a log parse operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed _tokenAddress)
func (_Bridge *BridgeFilterer) ParseTokenAdded(log types.Log) (*BridgeTokenAdded, error) {
	event := new(BridgeTokenAdded)
	if err := _Bridge.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the Bridge contract.
type BridgeTokenRemovedIterator struct {
	Event *BridgeTokenRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeTokenRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenRemoved represents a TokenRemoved event raised by the Bridge contract.
type BridgeTokenRemoved struct {
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed _tokenAddress)
func (_Bridge *BridgeFilterer) FilterTokenRemoved(opts *bind.FilterOpts, _tokenAddress []common.Address) (*BridgeTokenRemovedIterator, error) {

	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenRemoved", _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenRemovedIterator{contract: _Bridge.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed _tokenAddress)
func (_Bridge *BridgeFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *BridgeTokenRemoved, _tokenAddress []common.Address) (event.Subscription, error) {

	var _tokenAddressRule []interface{}
	for _, _tokenAddressItem := range _tokenAddress {
		_tokenAddressRule = append(_tokenAddressRule, _tokenAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenRemoved", _tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenRemoved)
				if err := _Bridge.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenRemoved is a log parse operation binding the contract event 0x4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd3.
//
// Solidity: event TokenRemoved(address indexed _tokenAddress)
func (_Bridge *BridgeFilterer) ParseTokenRemoved(log types.Log) (*BridgeTokenRemoved, error) {
	event := new(BridgeTokenRemoved)
	if err := _Bridge.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
