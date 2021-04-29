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
const BridgeABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_supportedTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_depositCompensation\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_tokenPrecision\",\"type\":\"uint8\"}],\"name\":\"ERC20Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"ETHDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compensationDeposited\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositCompensation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"depositERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"}],\"name\":\"depositETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payBackERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payBackETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setDepositCompensation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x60806040523480156200001157600080fd5b50604051620026503803806200265083398181016040528101906200003791906200028b565b600062000049620001b360201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35060005b8251811015620001a35760016003600085848151811062000134577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555080806200019a90620003b1565b915050620000ea565b50806001819055505050620004a2565b600033905090565b6000620001d2620001cc846200030e565b620002e5565b90508083825260208201905082856020860282011115620001f257600080fd5b60005b858110156200022657816200020b888262000230565b845260208401935060208301925050600181019050620001f5565b5050509392505050565b60008151905062000241816200046e565b92915050565b600082601f8301126200025957600080fd5b81516200026b848260208601620001bb565b91505092915050565b600081519050620002858162000488565b92915050565b600080604083850312156200029f57600080fd5b600083015167ffffffffffffffff811115620002ba57600080fd5b620002c88582860162000247565b9250506020620002db8582860162000274565b9150509250929050565b6000620002f162000304565b9050620002ff82826200037b565b919050565b6000604051905090565b600067ffffffffffffffff8211156200032c576200032b6200042e565b5b602082029050602081019050919050565b60006200034a8262000351565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b62000386826200045d565b810181811067ffffffffffffffff82111715620003a857620003a76200042e565b5b80604052505050565b6000620003be8262000371565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415620003f457620003f3620003ff565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b62000479816200033d565b81146200048557600080fd5b50565b620004938162000371565b81146200049f57600080fd5b50565b61219e80620004b26000396000f3fe6080604052600436106100c25760003560e01c80639371388d1161007f578063be20e53911610059578063be20e53914610288578063c7b77eb0146102c5578063d48bfca714610302578063f2fde38b1461033f576100c2565b80639371388d146101f0578063994f5a771461022d5780639b1c48e614610258576100c2565b806334c15e28146100c75780635fa7b5841461010457806368c4ac2614610141578063715018a61461017e5780637288f11a146101955780638da5cb5b146101c5575b600080fd5b3480156100d357600080fd5b506100ee60048036038101906100e99190611554565b610368565b6040516100fb9190611a15565b60405180910390f35b34801561011057600080fd5b5061012b60048036038101906101269190611554565b610388565b6040516101389190611a15565b60405180910390f35b34801561014d57600080fd5b5061016860048036038101906101639190611554565b6104aa565b6040516101759190611a15565b60405180910390f35b34801561018a57600080fd5b506101936104ca565b005b6101af60048036038101906101aa91906115cc565b610604565b6040516101bc9190611a15565b60405180910390f35b3480156101d157600080fd5b506101da610a2e565b6040516101e7919061199a565b60405180910390f35b3480156101fc57600080fd5b506102176004803603810190610212919061157d565b610a57565b6040516102249190611a15565b60405180910390f35b34801561023957600080fd5b50610242610cb8565b60405161024f9190611c13565b60405180910390f35b610272600480360381019061026d9190611698565b610cbe565b60405161027f9190611a15565b60405180910390f35b34801561029457600080fd5b506102af60048036038101906102aa919061171a565b610e6b565b6040516102bc9190611a15565b60405180910390f35b3480156102d157600080fd5b506102ec60048036038101906102e79190611633565b610ef9565b6040516102f99190611a15565b60405180910390f35b34801561030e57600080fd5b5061032960048036038101906103249190611554565b611134565b6040516103369190611a15565b60405180910390f35b34801561034b57600080fd5b5061036660048036038101906103619190611554565b611256565b005b60026020528060005260406000206000915054906101000a900460ff1681565b60006103926113ff565b73ffffffffffffffffffffffffffffffffffffffff166103b0610a2e565b73ffffffffffffffffffffffffffffffffffffffff1614610406576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103fd90611b33565b60405180910390fd5b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd360405160405180910390a260019050919050565b60036020528060005260406000206000915054906101000a900460ff1681565b6104d26113ff565b73ffffffffffffffffffffffffffffffffffffffff166104f0610a2e565b73ffffffffffffffffffffffffffffffffffffffff1614610546576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053d90611b33565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60006106258473ffffffffffffffffffffffffffffffffffffffff16611407565b610664576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065b90611bd3565b60405180910390fd5b600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166106f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e790611b93565b60405180910390fd5b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166107df57600154341015610786576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077d90611b53565b60405180910390fd5b6001600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505b600084905060008173ffffffffffffffffffffffffffffffffffffffff166323b872dd3330876040518463ffffffff1660e01b8152600401610823939291906119b5565b602060405180830381600087803b15801561083d57600080fd5b505af1158015610851573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610875919061166f565b9050806108b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ae90611bb3565b60405180910390fd5b8573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f380818a686b3749c864e4fe5b344c0fd3e0681231881695473b42c2c84d7271d87878673ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b15801561094e57600080fd5b505afa158015610962573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061098b91906116d9565b8773ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156109d157600080fd5b505afa1580156109e5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a099190611743565b604051610a199493929190611a60565b60405180910390a36001925050509392505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000610a616113ff565b73ffffffffffffffffffffffffffffffffffffffff16610a7f610a2e565b73ffffffffffffffffffffffffffffffffffffffff1614610ad5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610acc90611b33565b60405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86856040518363ffffffff1660e01b8152600401610b129291906119ec565b602060405180830381600087803b158015610b2c57600080fd5b505af1158015610b40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b64919061166f565b905080610ba6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9d90611b13565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff16600154604051610bcc90611985565b60006040518083038185875af1925050503d8060008114610c09576040519150601f19603f3d011682016040523d82523d6000602084013e610c0e565b606091505b50508091505080610c54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c4b90611bf3565b60405180910390fd5b6000600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060019150509392505050565b60015481565b600080349050600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610dce576000600154905080821015610d5f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5690611ab3565b60405180910390fd5b6001600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550610dca818361141a90919063ffffffff16565b9150505b60008111610e11576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e0890611af3565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167fa5ee847b808915e440e213630058643951dfa817cd585254162b8e69b71fa4328483604051610e59929190611a30565b60405180910390a26001915050919050565b6000610e756113ff565b73ffffffffffffffffffffffffffffffffffffffff16610e93610a2e565b73ffffffffffffffffffffffffffffffffffffffff1614610ee9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ee090611b33565b60405180910390fd5b8160018190555060019050919050565b6000610f036113ff565b73ffffffffffffffffffffffffffffffffffffffff16610f21610a2e565b73ffffffffffffffffffffffffffffffffffffffff1614610f77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6e90611b33565b60405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff1683604051610f9d90611985565b60006040518083038185875af1925050503d8060008114610fda576040519150601f19603f3d011682016040523d82523d6000602084013e610fdf565b606091505b5050905080611023576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101a90611b73565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff1660015460405161104990611985565b60006040518083038185875af1925050503d8060008114611086576040519150601f19603f3d011682016040523d82523d6000602084013e61108b565b606091505b505080915050806110d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110c890611bf3565b60405180910390fd5b6000600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600191505092915050565b600061113e6113ff565b73ffffffffffffffffffffffffffffffffffffffff1661115c610a2e565b73ffffffffffffffffffffffffffffffffffffffff16146111b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111a990611b33565b60405180910390fd5b6001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a460405160405180910390a260019050919050565b61125e6113ff565b73ffffffffffffffffffffffffffffffffffffffff1661127c610a2e565b73ffffffffffffffffffffffffffffffffffffffff16146112d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112c990611b33565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611342576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133990611ad3565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b600080823b905060008111915050919050565b600081836114289190611cab565b905092915050565b600061144361143e84611c53565b611c2e565b90508281526020810184848401111561145b57600080fd5b611466848285611d34565b509392505050565b600061148161147c84611c53565b611c2e565b90508281526020810184848401111561149957600080fd5b6114a4848285611d43565b509392505050565b6000813590506114bb8161210c565b92915050565b6000815190506114d081612123565b92915050565b600082601f8301126114e757600080fd5b81356114f7848260208601611430565b91505092915050565b600082601f83011261151157600080fd5b815161152184826020860161146e565b91505092915050565b6000813590506115398161213a565b92915050565b60008151905061154e81612151565b92915050565b60006020828403121561156657600080fd5b6000611574848285016114ac565b91505092915050565b60008060006060848603121561159257600080fd5b60006115a0868287016114ac565b93505060206115b1868287016114ac565b92505060406115c28682870161152a565b9150509250925092565b6000806000606084860312156115e157600080fd5b60006115ef868287016114ac565b935050602084013567ffffffffffffffff81111561160c57600080fd5b611618868287016114d6565b92505060406116298682870161152a565b9150509250925092565b6000806040838503121561164657600080fd5b6000611654858286016114ac565b92505060206116658582860161152a565b9150509250929050565b60006020828403121561168157600080fd5b600061168f848285016114c1565b91505092915050565b6000602082840312156116aa57600080fd5b600082013567ffffffffffffffff8111156116c457600080fd5b6116d0848285016114d6565b91505092915050565b6000602082840312156116eb57600080fd5b600082015167ffffffffffffffff81111561170557600080fd5b61171184828501611500565b91505092915050565b60006020828403121561172c57600080fd5b600061173a8482850161152a565b91505092915050565b60006020828403121561175557600080fd5b60006117638482850161153f565b91505092915050565b61177581611cdf565b82525050565b61178481611cf1565b82525050565b600061179582611c84565b61179f8185611c9a565b93506117af818560208601611d43565b6117b881611e05565b840191505092915050565b60006117d0602a83611c9a565b91506117db82611e16565b604082019050919050565b60006117f3602683611c9a565b91506117fe82611e65565b604082019050919050565b6000611816604583611c9a565b915061182182611eb4565b606082019050919050565b6000611839601283611c9a565b915061184482611f29565b602082019050919050565b600061185c602083611c9a565b915061186782611f52565b602082019050919050565b600061187f602b83611c9a565b915061188a82611f7b565b604082019050919050565b60006118a2602683611c9a565b91506118ad82611fca565b604082019050919050565b60006118c5602583611c9a565b91506118d082612019565b604082019050919050565b60006118e8601a83611c9a565b91506118f382612068565b602082019050919050565b600061190b600083611c8f565b915061191682612091565b600082019050919050565b600061192e601d83611c9a565b915061193982612094565b602082019050919050565b6000611951602f83611c9a565b915061195c826120bd565b604082019050919050565b61197081611d1d565b82525050565b61197f81611d27565b82525050565b6000611990826118fe565b9150819050919050565b60006020820190506119af600083018461176c565b92915050565b60006060820190506119ca600083018661176c565b6119d7602083018561176c565b6119e46040830184611967565b949350505050565b6000604082019050611a01600083018561176c565b611a0e6020830184611967565b9392505050565b6000602082019050611a2a600083018461177b565b92915050565b60006040820190508181036000830152611a4a818561178a565b9050611a596020830184611967565b9392505050565b60006080820190508181036000830152611a7a818761178a565b9050611a896020830186611967565b8181036040830152611a9b818561178a565b9050611aaa6060830184611976565b95945050505050565b60006020820190508181036000830152611acc816117c3565b9050919050565b60006020820190508181036000830152611aec816117e6565b9050919050565b60006020820190508181036000830152611b0c81611809565b9050919050565b60006020820190508181036000830152611b2c8161182c565b9050919050565b60006020820190508181036000830152611b4c8161184f565b9050919050565b60006020820190508181036000830152611b6c81611872565b9050919050565b60006020820190508181036000830152611b8c81611895565b9050919050565b60006020820190508181036000830152611bac816118b8565b9050919050565b60006020820190508181036000830152611bcc816118db565b9050919050565b60006020820190508181036000830152611bec81611921565b9050919050565b60006020820190508181036000830152611c0c81611944565b9050919050565b6000602082019050611c286000830184611967565b92915050565b6000611c38611c49565b9050611c448282611d76565b919050565b6000604051905090565b600067ffffffffffffffff821115611c6e57611c6d611dd6565b5b611c7782611e05565b9050602081019050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b6000611cb682611d1d565b9150611cc183611d1d565b925082821015611cd457611cd3611da7565b5b828203905092915050565b6000611cea82611cfd565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015611d61578082015181840152602081019050611d46565b83811115611d70576000848401525b50505050565b611d7f82611e05565b810181811067ffffffffffffffff82111715611d9e57611d9d611dd6565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f496e73756666696369656e742066756e647320746f206465706f73697420636f60008201527f6d70656e736174696f6e00000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f496e76616c69642076616c756520666f7220746865206465706f73697420616d60008201527f6f756e742c206661696c656420746f206465706f7369742061207a65726f207660208201527f616c75652e000000000000000000000000000000000000000000000000000000604082015250565b7f4661696c656420746f20706179206261636b0000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f496e73756666696369656e742066756e647320666f72206465706f736974206360008201527f6f6d70656e736174696f6e000000000000000000000000000000000000000000602082015250565b7f4661696c656420746f20706179206261636b20746865206465706f736974206160008201527f6d6f756e742e0000000000000000000000000000000000000000000000000000602082015250565b7f556e737570706f7274656420746f6b656e2c206661696c656420746f2064657060008201527f6f7369742e000000000000000000000000000000000000000000000000000000602082015250565b7f4661696c656420746f207472616e7366657220746f6b656e732e000000000000600082015250565b50565b7f476976656e20746f6b656e206973206e6f74206120636f6e7472616374000000600082015250565b7f4661696c656420746f207061792074686520636f6d70656e736174696f6e206660008201527f6f7220706179696e67206261636b2e0000000000000000000000000000000000602082015250565b61211581611cdf565b811461212057600080fd5b50565b61212c81611cf1565b811461213757600080fd5b50565b61214381611d1d565b811461214e57600080fd5b50565b61215a81611d27565b811461216557600080fd5b5056fea2646970667358221220d2d98b9a5a373807cb4672e566670b676edb9ddca8e5801832ecd31c6175ce5364736f6c63430008030033"

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
	UserAddress    common.Address
	OdinAddress    string
	DepositAmount  *big.Int
	TokenAddress   common.Address
	Symbol         string
	TokenPrecision uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterERC20Deposited is a free log retrieval operation binding the contract event 0x380818a686b3749c864e4fe5b344c0fd3e0681231881695473b42c2c84d7271d.
//
// Solidity: event ERC20Deposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount, address indexed _tokenAddress, string _symbol, uint8 _tokenPrecision)
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

// WatchERC20Deposited is a free log subscription operation binding the contract event 0x380818a686b3749c864e4fe5b344c0fd3e0681231881695473b42c2c84d7271d.
//
// Solidity: event ERC20Deposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount, address indexed _tokenAddress, string _symbol, uint8 _tokenPrecision)
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

// ParseERC20Deposited is a log parse operation binding the contract event 0x380818a686b3749c864e4fe5b344c0fd3e0681231881695473b42c2c84d7271d.
//
// Solidity: event ERC20Deposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount, address indexed _tokenAddress, string _symbol, uint8 _tokenPrecision)
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
