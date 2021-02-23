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

// OdinBridgeABI is the input ABI used to generate the binding from.
const OdinBridgeABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OdinBridgeBin is the compiled bytecode used for deploying new system-contracts.
var OdinBridgeBin = "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212205be7b290e8c910d5933af25fad5fffe9b8f039520913ba1d1c9d1b92758250c464736f6c63430007050033"

// DeployOdinBridge deploys a new Ethereum contract, binding an instance of OdinBridge to it.
func DeployOdinBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OdinBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(OdinBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OdinBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OdinBridge{OdinBridgeCaller: OdinBridgeCaller{contract: contract}, OdinBridgeTransactor: OdinBridgeTransactor{contract: contract}, OdinBridgeFilterer: OdinBridgeFilterer{contract: contract}}, nil
}

// OdinBridge is an auto generated Go binding around an Ethereum contract.
type OdinBridge struct {
	OdinBridgeCaller     // Read-only binding to the contract
	OdinBridgeTransactor // Write-only binding to the contract
	OdinBridgeFilterer   // Log filterer for contract events
}

// OdinBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type OdinBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OdinBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OdinBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OdinBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OdinBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OdinBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OdinBridgeSession struct {
	Contract     *OdinBridge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OdinBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OdinBridgeCallerSession struct {
	Contract *OdinBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OdinBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OdinBridgeTransactorSession struct {
	Contract     *OdinBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OdinBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type OdinBridgeRaw struct {
	Contract *OdinBridge // Generic contract binding to access the raw methods on
}

// OdinBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OdinBridgeCallerRaw struct {
	Contract *OdinBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// OdinBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OdinBridgeTransactorRaw struct {
	Contract *OdinBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOdinBridge creates a new instance of OdinBridge, bound to a specific deployed contract.
func NewOdinBridge(address common.Address, backend bind.ContractBackend) (*OdinBridge, error) {
	contract, err := bindOdinBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OdinBridge{OdinBridgeCaller: OdinBridgeCaller{contract: contract}, OdinBridgeTransactor: OdinBridgeTransactor{contract: contract}, OdinBridgeFilterer: OdinBridgeFilterer{contract: contract}}, nil
}

// NewOdinBridgeCaller creates a new read-only instance of OdinBridge, bound to a specific deployed contract.
func NewOdinBridgeCaller(address common.Address, caller bind.ContractCaller) (*OdinBridgeCaller, error) {
	contract, err := bindOdinBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OdinBridgeCaller{contract: contract}, nil
}

// NewOdinBridgeTransactor creates a new write-only instance of OdinBridge, bound to a specific deployed contract.
func NewOdinBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*OdinBridgeTransactor, error) {
	contract, err := bindOdinBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OdinBridgeTransactor{contract: contract}, nil
}

// NewOdinBridgeFilterer creates a new log filterer instance of OdinBridge, bound to a specific deployed contract.
func NewOdinBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*OdinBridgeFilterer, error) {
	contract, err := bindOdinBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OdinBridgeFilterer{contract: contract}, nil
}

// bindOdinBridge binds a generic wrapper to an already deployed contract.
func bindOdinBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OdinBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OdinBridge *OdinBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OdinBridge.Contract.OdinBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OdinBridge *OdinBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OdinBridge.Contract.OdinBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OdinBridge *OdinBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OdinBridge.Contract.OdinBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OdinBridge *OdinBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OdinBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OdinBridge *OdinBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OdinBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OdinBridge *OdinBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OdinBridge.Contract.contract.Transact(opts, method, params...)
}
