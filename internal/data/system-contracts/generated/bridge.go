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
const BridgeABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_supportedTokens\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"EtherDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"TokenDeposited\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x608060405234801561001057600080fd5b5060405162002567380380620025678339818101604052602081101561003557600080fd5b810190808051604051939291908464010000000082111561005557600080fd5b8382019150602082018581111561006b57600080fd5b825186602082028301116401000000008211171561008857600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100bf5780820151818401526020810190506100a4565b5050505090500160405250505060006100dc61023560201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a350806040516101879061023d565b8080602001828103825283818151815260200191508051906020019060200280838360005b838110156101c75780820151818401526020810190506101ac565b5050505090500192505050604051809103906000f0801580156101ee573d6000803e3d6000fd5b50600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061024b565b600033905090565b6113d8806200118f83390190565b610f34806200025b6000396000f3fe6080604052600436106100705760003560e01c8063a26e11861161004e578063a26e118614610134578063d48bfca714610205578063f2fde38b1461026c578063f3db9fe5146102bd57610070565b80635fa7b58414610075578063715018a6146100dc5780638da5cb5b146100f3575b600080fd5b34801561008157600080fd5b506100c46004803603602081101561009857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506103c5565b60405180821515815260200191505060405180910390f35b3480156100e857600080fd5b506100f1610522565b005b3480156100ff57600080fd5b5061010861068f565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101ed6004803603602081101561014a57600080fd5b810190808035906020019064010000000081111561016757600080fd5b82018360208201111561017957600080fd5b8035906020019184600183028401116401000000008311171561019b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506106b8565b60405180821515815260200191505060405180910390f35b34801561021157600080fd5b506102546004803603602081101561022857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107d6565b60405180821515815260200191505060405180910390f35b34801561027857600080fd5b506102bb6004803603602081101561028f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610933565b005b3480156102c957600080fd5b506103ad600480360360608110156102e057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561031d57600080fd5b82018360208201111561032f57600080fd5b8035906020019184600183028401116401000000008311171561035157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050610b25565b60405180821515815260200191505060405180910390f35b60006103cf610e66565b73ffffffffffffffffffffffffffffffffffffffff166103ed61068f565b73ffffffffffffffffffffffffffffffffffffffff1614610476576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c32d805a836040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561050157600080fd5b505af1158015610515573d6000803e3d6000fd5b5050505060019050919050565b61052a610e66565b73ffffffffffffffffffffffffffffffffffffffff1661054861068f565b73ffffffffffffffffffffffffffffffffffffffff16146105d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000803411610712576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526045815260200180610e956045913960600191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167ff93ea55a2b7738a220c1c386239e26cf4448d5724777f7d2ffd8c6524ca0c37f83346040518080602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610792578082015181840152602081019050610777565b50505050905090810190601f1680156107bf5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a260019050919050565b60006107e0610e66565b73ffffffffffffffffffffffffffffffffffffffff166107fe61068f565b73ffffffffffffffffffffffffffffffffffffffff1614610887576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639e9405c3836040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561091257600080fd5b505af1158015610926573d6000803e3d6000fd5b5050505060019050919050565b61093b610e66565b73ffffffffffffffffffffffffffffffffffffffff1661095961068f565b73ffffffffffffffffffffffffffffffffffffffff16146109e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a68576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180610e6f6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635dbe47e8856040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015610bb057600080fd5b505afa158015610bc4573d6000803e3d6000fd5b505050506040513d6020811015610bda57600080fd5b8101908080519060200190929190505050610c40576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180610eda6025913960400191505060405180910390fd5b60008473ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610cd157600080fd5b505af1158015610ce5573d6000803e3d6000fd5b505050506040513d6020811015610cfb57600080fd5b8101908080519060200190929190505050905080610d81576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f4661696c656420746f207472616e7366657220746f6b656e732e00000000000081525060200191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167fa4db3f02423c0041c8767e69f6d4ff1935d1a64182cff6348d0e6468176c72c285878660405180806020018473ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828103825285818151815260200191508051906020019080838360005b83811015610e1e578082015181840152602081019050610e03565b50505050905090810190601f168015610e4b5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a260019150509392505050565b60003390509056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e76616c69642076616c756520666f7220746865206465706f73697420616d6f756e742c206661696c656420746f206465706f7369742061207a65726f2076616c75652e556e737570706f7274656420746f6b656e2c206661696c656420746f206465706f7369742ea2646970667358221220e93789f7a84239b79a1f872b04d737a44cc9aadb09f62f0f9bd82b3873c70fdc64736f6c63430007020033608060405234801561001057600080fd5b50604051620013d8380380620013d88339818101604052602081101561003557600080fd5b810190808051604051939291908464010000000082111561005557600080fd5b8382019150602082018581111561006b57600080fd5b825186602082028301116401000000008211171561008857600080fd5b8083526020830192505050908051906020019060200280838360005b838110156100bf5780820151818401526020810190506100a4565b5050505090500160405250505060006100dc61026f60201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35060005b815181101561026857600282828151811061019457fe5b60200260200101519080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506002805490506001600084848151811061021257fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000181905550808060010191505061017d565b5050610277565b600033905090565b61115180620002876000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063949d225d11610066578063949d225d146101ef5780639e9405c31461020d578063a39fac1214610251578063c32d805a146102b0578063f2fde38b146102f45761009e565b80630a3b0a4f146100a357806329092d0e146100fd5780635dbe47e814610157578063715018a6146101b15780638da5cb5b146101bb575b600080fd5b6100e5600480360360208110156100b957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610338565b60405180821515815260200191505060405180910390f35b61013f6004803603602081101561011357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061044f565b60405180821515815260200191505060405180910390f35b6101996004803603602081101561016d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061057d565b60405180821515815260200191505060405180910390f35b6101b96105cb565b005b6101c3610738565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101f7610761565b6040518082815260200191505060405180910390f35b61024f6004803603602081101561022357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061076e565b005b6102596108c4565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561029c578082015181840152602081019050610281565b505050509050019250505060405180910390f35b6102f2600480360360208110156102c657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610952565b005b6103366004803603602081101561030a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b0d565b005b6000610342610cff565b73ffffffffffffffffffffffffffffffffffffffff16610360610738565b73ffffffffffffffffffffffffffffffffffffffff16146103e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541461043c576000905061044a565b61044582610d07565b600190505b919050565b6000610459610cff565b73ffffffffffffffffffffffffffffffffffffffff16610477610738565b73ffffffffffffffffffffffffffffffffffffffff1614610500576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001549050600081148061055a575060028054905081115b15610569576000915050610578565b61057283610db9565b60019150505b919050565b600080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154119050919050565b6105d3610cff565b73ffffffffffffffffffffffffffffffffffffffff166105f1610738565b73ffffffffffffffffffffffffffffffffffffffff161461067a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600280549050905090565b610776610cff565b73ffffffffffffffffffffffffffffffffffffffff16610794610738565b73ffffffffffffffffffffffffffffffffffffffff161461081d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154146108b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260648152602001806110666064913960800191505060405180910390fd5b6108c181610d07565b50565b6060600280548060200260200160405190810160405280929190818152602001828054801561094857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116108fe575b5050505050905090565b61095a610cff565b73ffffffffffffffffffffffffffffffffffffffff16610978610738565b73ffffffffffffffffffffffffffffffffffffffff1614610a01576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015490506000811415610aa2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260528152602001806110ca6052913960600191505060405180910390fd5b600280549050811115610b00576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260638152602001806110036063913960800191505060405180910390fd5b610b0982610db9565b5050565b610b15610cff565b73ffffffffffffffffffffffffffffffffffffffff16610b33610738565b73ffffffffffffffffffffffffffffffffffffffff1614610bbc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610c42576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180610fdd6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b6002819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600280549050600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555050565b610dc1610fc9565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180602001604052908160008201548152505090506000600160028054905003905081600001516001600060028481548110610e3a57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555060028181548110610eb357fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166002600184600001510381548110610ef257fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506002805480610f4557fe5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000808201600090555050505050565b604051806020016040528060008152509056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e76616c696420696e6465782076616c756520666f722074686520616464726573732073746f726167652c206661696c656420746f2072656d6f76652074686520616464726573732066726f6d2074686520616464726573732073746f726167652e54686520616464726573732068617320616c7265616479206265656e20616464656420746f207468652073746f726167652c206661696c656420746f2061646420746865206164647265737320746f2074686520616464726573732073746f726167652e546865206164647265737320646f6573206e6f742065786973742c206661696c656420746f2072656d6f76652074686520616464726573732066726f6d2074686520616464726573732073746f726167652ea264697066735822122050972d3f3a712cbfa735142c9da9d0b7dcbf23bb6c8cc2c226ceb8806c95631864736f6c63430007020033"

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

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string _odinAddress) payable returns(bool)
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, _odinAddress string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", _odinAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string _odinAddress) payable returns(bool)
func (_Bridge *BridgeSession) Deposit(_odinAddress string) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, _odinAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string _odinAddress) payable returns(bool)
func (_Bridge *BridgeTransactorSession) Deposit(_odinAddress string) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, _odinAddress)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf3db9fe5.
//
// Solidity: function deposit(address _tokenAddress, string _odinAddress, uint256 _depositAmount) returns(bool)
func (_Bridge *BridgeTransactor) Deposit0(opts *bind.TransactOpts, _tokenAddress common.Address, _odinAddress string, _depositAmount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit0", _tokenAddress, _odinAddress, _depositAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf3db9fe5.
//
// Solidity: function deposit(address _tokenAddress, string _odinAddress, uint256 _depositAmount) returns(bool)
func (_Bridge *BridgeSession) Deposit0(_tokenAddress common.Address, _odinAddress string, _depositAmount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit0(&_Bridge.TransactOpts, _tokenAddress, _odinAddress, _depositAmount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf3db9fe5.
//
// Solidity: function deposit(address _tokenAddress, string _odinAddress, uint256 _depositAmount) returns(bool)
func (_Bridge *BridgeTransactorSession) Deposit0(_tokenAddress common.Address, _odinAddress string, _depositAmount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit0(&_Bridge.TransactOpts, _tokenAddress, _odinAddress, _depositAmount)
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

// BridgeEtherDepositedIterator is returned from FilterEtherDeposited and is used to iterate over the raw logs and unpacked data for EtherDeposited events raised by the Bridge contract.
type BridgeEtherDepositedIterator struct {
	Event *BridgeEtherDeposited // Event containing the contract specifics and raw log

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
func (it *BridgeEtherDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEtherDeposited)
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
		it.Event = new(BridgeEtherDeposited)
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
func (it *BridgeEtherDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEtherDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEtherDeposited represents a EtherDeposited event raised by the Bridge contract.
type BridgeEtherDeposited struct {
	UserAddress   common.Address
	OdinAddress   string
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEtherDeposited is a free log retrieval operation binding the contract event 0xf93ea55a2b7738a220c1c386239e26cf4448d5724777f7d2ffd8c6524ca0c37f.
//
// Solidity: event EtherDeposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount)
func (_Bridge *BridgeFilterer) FilterEtherDeposited(opts *bind.FilterOpts, _userAddress []common.Address) (*BridgeEtherDepositedIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "EtherDeposited", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgeEtherDepositedIterator{contract: _Bridge.contract, event: "EtherDeposited", logs: logs, sub: sub}, nil
}

// WatchEtherDeposited is a free log subscription operation binding the contract event 0xf93ea55a2b7738a220c1c386239e26cf4448d5724777f7d2ffd8c6524ca0c37f.
//
// Solidity: event EtherDeposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount)
func (_Bridge *BridgeFilterer) WatchEtherDeposited(opts *bind.WatchOpts, sink chan<- *BridgeEtherDeposited, _userAddress []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "EtherDeposited", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEtherDeposited)
				if err := _Bridge.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
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

// ParseEtherDeposited is a log parse operation binding the contract event 0xf93ea55a2b7738a220c1c386239e26cf4448d5724777f7d2ffd8c6524ca0c37f.
//
// Solidity: event EtherDeposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount)
func (_Bridge *BridgeFilterer) ParseEtherDeposited(log types.Log) (*BridgeEtherDeposited, error) {
	event := new(BridgeEtherDeposited)
	if err := _Bridge.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
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

// BridgeTokenDepositedIterator is returned from FilterTokenDeposited and is used to iterate over the raw logs and unpacked data for TokenDeposited events raised by the Bridge contract.
type BridgeTokenDepositedIterator struct {
	Event *BridgeTokenDeposited // Event containing the contract specifics and raw log

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
func (it *BridgeTokenDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenDeposited)
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
		it.Event = new(BridgeTokenDeposited)
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
func (it *BridgeTokenDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenDeposited represents a TokenDeposited event raised by the Bridge contract.
type BridgeTokenDeposited struct {
	UserAddress   common.Address
	OdinAddress   string
	TokenAddress  common.Address
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposited is a free log retrieval operation binding the contract event 0xa4db3f02423c0041c8767e69f6d4ff1935d1a64182cff6348d0e6468176c72c2.
//
// Solidity: event TokenDeposited(address indexed _userAddress, string _odinAddress, address _tokenAddress, uint256 _depositAmount)
func (_Bridge *BridgeFilterer) FilterTokenDeposited(opts *bind.FilterOpts, _userAddress []common.Address) (*BridgeTokenDepositedIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenDeposited", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenDepositedIterator{contract: _Bridge.contract, event: "TokenDeposited", logs: logs, sub: sub}, nil
}

// WatchTokenDeposited is a free log subscription operation binding the contract event 0xa4db3f02423c0041c8767e69f6d4ff1935d1a64182cff6348d0e6468176c72c2.
//
// Solidity: event TokenDeposited(address indexed _userAddress, string _odinAddress, address _tokenAddress, uint256 _depositAmount)
func (_Bridge *BridgeFilterer) WatchTokenDeposited(opts *bind.WatchOpts, sink chan<- *BridgeTokenDeposited, _userAddress []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenDeposited", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenDeposited)
				if err := _Bridge.contract.UnpackLog(event, "TokenDeposited", log); err != nil {
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

// ParseTokenDeposited is a log parse operation binding the contract event 0xa4db3f02423c0041c8767e69f6d4ff1935d1a64182cff6348d0e6468176c72c2.
//
// Solidity: event TokenDeposited(address indexed _userAddress, string _odinAddress, address _tokenAddress, uint256 _depositAmount)
func (_Bridge *BridgeFilterer) ParseTokenDeposited(log types.Log) (*BridgeTokenDeposited, error) {
	event := new(BridgeTokenDeposited)
	if err := _Bridge.contract.UnpackLog(event, "TokenDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
