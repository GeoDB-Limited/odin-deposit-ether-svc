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
const BridgeABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_supportedTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"ERC20Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"ETHDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"}],\"name\":\"depositEther\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payBackEther\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"payBackToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x60806040523480156200001157600080fd5b5060405162001c8738038062001c8783398181016040528101906200003791906200026b565b600062000049620001aa60201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35060005b8151811015620001a257600180600084848151811062000133577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550808062000199906200037c565b915050620000ea565b505062000453565b600033905090565b6000620001c9620001c384620002d9565b620002b0565b90508083825260208201905082856020860282011115620001e957600080fd5b60005b858110156200021d578162000202888262000227565b845260208401935060208301925050600181019050620001ec565b5050509392505050565b600081519050620002388162000439565b92915050565b600082601f8301126200025057600080fd5b815162000262848260208601620001b2565b91505092915050565b6000602082840312156200027e57600080fd5b600082015167ffffffffffffffff8111156200029957600080fd5b620002a7848285016200023e565b91505092915050565b6000620002bc620002cf565b9050620002ca828262000346565b919050565b6000604051905090565b600067ffffffffffffffff821115620002f757620002f6620003f9565b5b602082029050602081019050919050565b600062000315826200031c565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b620003518262000428565b810181811067ffffffffffffffff82111715620003735762000372620003f9565b5b80604052505050565b600062000389826200033c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415620003bf57620003be620003ca565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b620004448162000308565b81146200045057600080fd5b50565b61182480620004636000396000f3fe6080604052600436106100915760003560e01c80638da5cb5b116100595780638da5cb5b14610194578063b5653ec6146101bf578063d48bfca7146101fc578063e83c847214610239578063f2fde38b1461027657610091565b8063428377b8146100965780635fa7b584146100d357806368c4ac2614610110578063715018a61461014d57806377c7632114610164575b600080fd5b3480156100a257600080fd5b506100bd60048036038101906100b89190610f33565b61029f565b6040516100ca9190611294565b60405180910390f35b3480156100df57600080fd5b506100fa60048036038101906100f59190610ebb565b6104d1565b6040516101079190611294565b60405180910390f35b34801561011c57600080fd5b5061013760048036038101906101329190610ebb565b6105f3565b6040516101449190611294565b60405180910390f35b34801561015957600080fd5b50610162610613565b005b61017e60048036038101906101799190610fff565b61074d565b60405161018b9190611294565b60405180910390f35b3480156101a057600080fd5b506101a96107ea565b6040516101b69190611219565b60405180910390f35b3480156101cb57600080fd5b506101e660048036038101906101e19190610ee4565b610813565b6040516101f39190611294565b60405180910390f35b34801561020857600080fd5b50610223600480360381019061021e9190610ebb565b6109fa565b6040516102309190611294565b60405180910390f35b34801561024557600080fd5b50610260600480360381019061025b9190610f9a565b610b1b565b60405161026d9190611294565b60405180910390f35b34801561028257600080fd5b5061029d60048036038101906102989190610ebb565b610c50565b005b60006102c08473ffffffffffffffffffffffffffffffffffffffff16610df9565b6102ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f6906113df565b60405180910390fd5b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661038b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103829061137f565b60405180910390fd5b60008473ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b81526004016103ca93929190611234565b602060405180830381600087803b1580156103e457600080fd5b505af11580156103f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061041c9190610fd6565b90508061045e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610455906113bf565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fec58c097aa640af3fa1dfb24771f7eff813e3fb228f984cf0d92672e68bffa4286866040516104bd9291906112af565b60405180910390a360019150509392505050565b60006104db610e0c565b73ffffffffffffffffffffffffffffffffffffffff166104f96107ea565b73ffffffffffffffffffffffffffffffffffffffff161461054f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105469061133f565b60405180910390fd5b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f4c910b69fe65a61f7531b9c5042b2329ca7179c77290aa7e2eb3afa3c8511fd360405160405180910390a260019050919050565b60016020528060005260406000206000915054906101000a900460ff1681565b61061b610e0c565b73ffffffffffffffffffffffffffffffffffffffff166106396107ea565b73ffffffffffffffffffffffffffffffffffffffff161461068f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106869061133f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000803411610791576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610788906112ff565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167fa5ee847b808915e440e213630058643951dfa817cd585254162b8e69b71fa43283346040516107d99291906112af565b60405180910390a260019050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600061081d610e0c565b73ffffffffffffffffffffffffffffffffffffffff1661083b6107ea565b73ffffffffffffffffffffffffffffffffffffffff1614610891576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108889061133f565b60405180910390fd5b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661091d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109149061139f565b60405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86856040518363ffffffff1660e01b815260040161095a92919061126b565b602060405180830381600087803b15801561097457600080fd5b505af1158015610988573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ac9190610fd6565b9050806109ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e59061131f565b60405180910390fd5b60019150509392505050565b6000610a04610e0c565b73ffffffffffffffffffffffffffffffffffffffff16610a226107ea565b73ffffffffffffffffffffffffffffffffffffffff1614610a78576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6f9061133f565b60405180910390fd5b60018060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a460405160405180910390a260019050919050565b6000610b25610e0c565b73ffffffffffffffffffffffffffffffffffffffff16610b436107ea565b73ffffffffffffffffffffffffffffffffffffffff1614610b99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b909061133f565b60405180910390fd5b60008373ffffffffffffffffffffffffffffffffffffffff1683604051610bbf90611204565b60006040518083038185875af1925050503d8060008114610bfc576040519150601f19603f3d011682016040523d82523d6000602084013e610c01565b606091505b5050905080610c45576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3c9061135f565b60405180910390fd5b600191505092915050565b610c58610e0c565b73ffffffffffffffffffffffffffffffffffffffff16610c766107ea565b73ffffffffffffffffffffffffffffffffffffffff1614610ccc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cc39061133f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610d3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d33906112df565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080823b905060008111915050919050565b600033905090565b6000610e27610e2284611424565b6113ff565b905082815260208101848484011115610e3f57600080fd5b610e4a8482856114c4565b509392505050565b600081359050610e61816117a9565b92915050565b600081519050610e76816117c0565b92915050565b600082601f830112610e8d57600080fd5b8135610e9d848260208601610e14565b91505092915050565b600081359050610eb5816117d7565b92915050565b600060208284031215610ecd57600080fd5b6000610edb84828501610e52565b91505092915050565b600080600060608486031215610ef957600080fd5b6000610f0786828701610e52565b9350506020610f1886828701610e52565b9250506040610f2986828701610ea6565b9150509250925092565b600080600060608486031215610f4857600080fd5b6000610f5686828701610e52565b935050602084013567ffffffffffffffff811115610f7357600080fd5b610f7f86828701610e7c565b9250506040610f9086828701610ea6565b9150509250925092565b60008060408385031215610fad57600080fd5b6000610fbb85828601610e52565b9250506020610fcc85828601610ea6565b9150509250929050565b600060208284031215610fe857600080fd5b6000610ff684828501610e67565b91505092915050565b60006020828403121561101157600080fd5b600082013567ffffffffffffffff81111561102b57600080fd5b61103784828501610e7c565b91505092915050565b6110498161147c565b82525050565b6110588161148e565b82525050565b600061106982611455565b611073818561146b565b93506110838185602086016114d3565b61108c81611566565b840191505092915050565b60006110a460268361146b565b91506110af82611577565b604082019050919050565b60006110c760458361146b565b91506110d2826115c6565b606082019050919050565b60006110ea60128361146b565b91506110f58261163b565b602082019050919050565b600061110d60208361146b565b915061111882611664565b602082019050919050565b600061113060268361146b565b915061113b8261168d565b604082019050919050565b600061115360258361146b565b915061115e826116dc565b604082019050919050565b600061117660128361146b565b91506111818261172b565b602082019050919050565b6000611199601a8361146b565b91506111a482611754565b602082019050919050565b60006111bc600083611460565b91506111c78261177d565b600082019050919050565b60006111df601d8361146b565b91506111ea82611780565b602082019050919050565b6111fe816114ba565b82525050565b600061120f826111af565b9150819050919050565b600060208201905061122e6000830184611040565b92915050565b60006060820190506112496000830186611040565b6112566020830185611040565b61126360408301846111f5565b949350505050565b60006040820190506112806000830185611040565b61128d60208301846111f5565b9392505050565b60006020820190506112a9600083018461104f565b92915050565b600060408201905081810360008301526112c9818561105e565b90506112d860208301846111f5565b9392505050565b600060208201905081810360008301526112f881611097565b9050919050565b60006020820190508181036000830152611318816110ba565b9050919050565b60006020820190508181036000830152611338816110dd565b9050919050565b6000602082019050818103600083015261135881611100565b9050919050565b6000602082019050818103600083015261137881611123565b9050919050565b6000602082019050818103600083015261139881611146565b9050919050565b600060208201905081810360008301526113b881611169565b9050919050565b600060208201905081810360008301526113d88161118c565b9050919050565b600060208201905081810360008301526113f8816111d2565b9050919050565b600061140961141a565b90506114158282611506565b919050565b6000604051905090565b600067ffffffffffffffff82111561143f5761143e611537565b5b61144882611566565b9050602081019050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b60006114878261149a565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156114f15780820151818401526020810190506114d6565b83811115611500576000848401525b50505050565b61150f82611566565b810181811067ffffffffffffffff8211171561152e5761152d611537565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f496e76616c69642076616c756520666f7220746865206465706f73697420616d60008201527f6f756e742c206661696c656420746f206465706f7369742061207a65726f207660208201527f616c75652e000000000000000000000000000000000000000000000000000000604082015250565b7f4661696c656420746f20706179206261636b0000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f4661696c656420746f20706179206261636b20746865206465706f736974206160008201527f6d6f756e742e0000000000000000000000000000000000000000000000000000602082015250565b7f556e737570706f7274656420746f6b656e2c206661696c656420746f2064657060008201527f6f7369742e000000000000000000000000000000000000000000000000000000602082015250565b7f556e737570706f7274656420746f6b656e2e0000000000000000000000000000600082015250565b7f4661696c656420746f207472616e7366657220746f6b656e732e000000000000600082015250565b50565b7f476976656e20746f6b656e206973206e6f74206120636f6e7472616374000000600082015250565b6117b28161147c565b81146117bd57600080fd5b50565b6117c98161148e565b81146117d457600080fd5b50565b6117e0816114ba565b81146117eb57600080fd5b5056fea2646970667358221220dab7bd0772e32ecfcf7ee227555681665037519f7366903375776fd23268e35764736f6c63430008030033"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _supportedTokens []common.Address) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend, _supportedTokens)
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

// DepositEther is a paid mutator transaction binding the contract method 0x77c76321.
//
// Solidity: function depositEther(string _odinAddress) payable returns(bool)
func (_Bridge *BridgeTransactor) DepositEther(opts *bind.TransactOpts, _odinAddress string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositEther", _odinAddress)
}

// DepositEther is a paid mutator transaction binding the contract method 0x77c76321.
//
// Solidity: function depositEther(string _odinAddress) payable returns(bool)
func (_Bridge *BridgeSession) DepositEther(_odinAddress string) (*types.Transaction, error) {
	return _Bridge.Contract.DepositEther(&_Bridge.TransactOpts, _odinAddress)
}

// DepositEther is a paid mutator transaction binding the contract method 0x77c76321.
//
// Solidity: function depositEther(string _odinAddress) payable returns(bool)
func (_Bridge *BridgeTransactorSession) DepositEther(_odinAddress string) (*types.Transaction, error) {
	return _Bridge.Contract.DepositEther(&_Bridge.TransactOpts, _odinAddress)
}

// DepositToken is a paid mutator transaction binding the contract method 0x428377b8.
//
// Solidity: function depositToken(address _tokenAddress, string _odinAddress, uint256 _depositAmount) returns(bool)
func (_Bridge *BridgeTransactor) DepositToken(opts *bind.TransactOpts, _tokenAddress common.Address, _odinAddress string, _depositAmount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "depositToken", _tokenAddress, _odinAddress, _depositAmount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x428377b8.
//
// Solidity: function depositToken(address _tokenAddress, string _odinAddress, uint256 _depositAmount) returns(bool)
func (_Bridge *BridgeSession) DepositToken(_tokenAddress common.Address, _odinAddress string, _depositAmount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DepositToken(&_Bridge.TransactOpts, _tokenAddress, _odinAddress, _depositAmount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x428377b8.
//
// Solidity: function depositToken(address _tokenAddress, string _odinAddress, uint256 _depositAmount) returns(bool)
func (_Bridge *BridgeTransactorSession) DepositToken(_tokenAddress common.Address, _odinAddress string, _depositAmount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.DepositToken(&_Bridge.TransactOpts, _tokenAddress, _odinAddress, _depositAmount)
}

// PayBackEther is a paid mutator transaction binding the contract method 0xe83c8472.
//
// Solidity: function payBackEther(address _user, uint256 _amount) returns(bool)
func (_Bridge *BridgeTransactor) PayBackEther(opts *bind.TransactOpts, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "payBackEther", _user, _amount)
}

// PayBackEther is a paid mutator transaction binding the contract method 0xe83c8472.
//
// Solidity: function payBackEther(address _user, uint256 _amount) returns(bool)
func (_Bridge *BridgeSession) PayBackEther(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.PayBackEther(&_Bridge.TransactOpts, _user, _amount)
}

// PayBackEther is a paid mutator transaction binding the contract method 0xe83c8472.
//
// Solidity: function payBackEther(address _user, uint256 _amount) returns(bool)
func (_Bridge *BridgeTransactorSession) PayBackEther(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.PayBackEther(&_Bridge.TransactOpts, _user, _amount)
}

// PayBackToken is a paid mutator transaction binding the contract method 0xb5653ec6.
//
// Solidity: function payBackToken(address _user, address _token, uint256 _amount) returns(bool)
func (_Bridge *BridgeTransactor) PayBackToken(opts *bind.TransactOpts, _user common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "payBackToken", _user, _token, _amount)
}

// PayBackToken is a paid mutator transaction binding the contract method 0xb5653ec6.
//
// Solidity: function payBackToken(address _user, address _token, uint256 _amount) returns(bool)
func (_Bridge *BridgeSession) PayBackToken(_user common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.PayBackToken(&_Bridge.TransactOpts, _user, _token, _amount)
}

// PayBackToken is a paid mutator transaction binding the contract method 0xb5653ec6.
//
// Solidity: function payBackToken(address _user, address _token, uint256 _amount) returns(bool)
func (_Bridge *BridgeTransactorSession) PayBackToken(_user common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.PayBackToken(&_Bridge.TransactOpts, _user, _token, _amount)
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
	TokenAddress  common.Address
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterERC20Deposited is a free log retrieval operation binding the contract event 0xec58c097aa640af3fa1dfb24771f7eff813e3fb228f984cf0d92672e68bffa42.
//
// Solidity: event ERC20Deposited(address indexed _userAddress, string _odinAddress, address indexed _tokenAddress, uint256 _depositAmount)
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

// WatchERC20Deposited is a free log subscription operation binding the contract event 0xec58c097aa640af3fa1dfb24771f7eff813e3fb228f984cf0d92672e68bffa42.
//
// Solidity: event ERC20Deposited(address indexed _userAddress, string _odinAddress, address indexed _tokenAddress, uint256 _depositAmount)
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

// ParseERC20Deposited is a log parse operation binding the contract event 0xec58c097aa640af3fa1dfb24771f7eff813e3fb228f984cf0d92672e68bffa42.
//
// Solidity: event ERC20Deposited(address indexed _userAddress, string _odinAddress, address indexed _tokenAddress, uint256 _depositAmount)
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
