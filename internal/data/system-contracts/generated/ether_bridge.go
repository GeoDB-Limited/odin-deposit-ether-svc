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

// EtherBridgeABI is the input ABI used to generate the binding from.
const EtherBridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"EtherDeposited\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"}],\"name\":\"depositEther\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// EtherBridgeBin is the compiled bytecode used for deploying new contracts.
var EtherBridgeBin = "0x608060405234801561001057600080fd5b50610464806100206000396000f3fe60806040526004361061001e5760003560e01c806377c7632114610023575b600080fd5b61003d60048036038101906100389190610158565b610053565b60405161004a9190610213565b60405180910390f35b6000803411610097576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161008e9061025e565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167ff93ea55a2b7738a220c1c386239e26cf4448d5724777f7d2ffd8c6524ca0c37f83346040516100df92919061022e565b60405180910390a260019050919050565b60006101036100fe846102a3565b61027e565b90508281526020810184848401111561011b57600080fd5b610126848285610306565b509392505050565b600082601f83011261013f57600080fd5b813561014f8482602086016100f0565b91505092915050565b60006020828403121561016a57600080fd5b600082013567ffffffffffffffff81111561018457600080fd5b6101908482850161012e565b91505092915050565b6101a2816102f0565b82525050565b60006101b3826102d4565b6101bd81856102df565b93506101cd818560208601610315565b6101d6816103a8565b840191505092915050565b60006101ee6045836102df565b91506101f9826103b9565b606082019050919050565b61020d816102fc565b82525050565b60006020820190506102286000830184610199565b92915050565b6000604082019050818103600083015261024881856101a8565b90506102576020830184610204565b9392505050565b60006020820190508181036000830152610277816101e1565b9050919050565b6000610288610299565b90506102948282610348565b919050565b6000604051905090565b600067ffffffffffffffff8211156102be576102bd610379565b5b6102c7826103a8565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b60008115159050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610333578082015181840152602081019050610318565b83811115610342576000848401525b50505050565b610351826103a8565b810181811067ffffffffffffffff821117156103705761036f610379565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f496e76616c69642076616c756520666f7220746865206465706f73697420616d60008201527f6f756e742c206661696c656420746f206465706f7369742061207a65726f207660208201527f616c75652e00000000000000000000000000000000000000000000000000000060408201525056fea26469706673582212200c70e3e3db35519d1797f4ddf32aff1f09524883ebc473f0b3c9ba6e7dbee36164736f6c63430008030033"

// DeployEtherBridge deploys a new Ethereum contract, binding an instance of EtherBridge to it.
func DeployEtherBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EtherBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EtherBridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherBridge{EtherBridgeCaller: EtherBridgeCaller{contract: contract}, EtherBridgeTransactor: EtherBridgeTransactor{contract: contract}, EtherBridgeFilterer: EtherBridgeFilterer{contract: contract}}, nil
}

// EtherBridge is an auto generated Go binding around an Ethereum contract.
type EtherBridge struct {
	EtherBridgeCaller     // Read-only binding to the contract
	EtherBridgeTransactor // Write-only binding to the contract
	EtherBridgeFilterer   // Log filterer for contract events
}

// EtherBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherBridgeSession struct {
	Contract     *EtherBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherBridgeCallerSession struct {
	Contract *EtherBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EtherBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherBridgeTransactorSession struct {
	Contract     *EtherBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EtherBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherBridgeRaw struct {
	Contract *EtherBridge // Generic contract binding to access the raw methods on
}

// EtherBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherBridgeCallerRaw struct {
	Contract *EtherBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// EtherBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherBridgeTransactorRaw struct {
	Contract *EtherBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherBridge creates a new instance of EtherBridge, bound to a specific deployed contract.
func NewEtherBridge(address common.Address, backend bind.ContractBackend) (*EtherBridge, error) {
	contract, err := bindEtherBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherBridge{EtherBridgeCaller: EtherBridgeCaller{contract: contract}, EtherBridgeTransactor: EtherBridgeTransactor{contract: contract}, EtherBridgeFilterer: EtherBridgeFilterer{contract: contract}}, nil
}

// NewEtherBridgeCaller creates a new read-only instance of EtherBridge, bound to a specific deployed contract.
func NewEtherBridgeCaller(address common.Address, caller bind.ContractCaller) (*EtherBridgeCaller, error) {
	contract, err := bindEtherBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherBridgeCaller{contract: contract}, nil
}

// NewEtherBridgeTransactor creates a new write-only instance of EtherBridge, bound to a specific deployed contract.
func NewEtherBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherBridgeTransactor, error) {
	contract, err := bindEtherBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherBridgeTransactor{contract: contract}, nil
}

// NewEtherBridgeFilterer creates a new log filterer instance of EtherBridge, bound to a specific deployed contract.
func NewEtherBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherBridgeFilterer, error) {
	contract, err := bindEtherBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherBridgeFilterer{contract: contract}, nil
}

// bindEtherBridge binds a generic wrapper to an already deployed contract.
func bindEtherBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherBridge *EtherBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherBridge.Contract.EtherBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherBridge *EtherBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherBridge.Contract.EtherBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherBridge *EtherBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherBridge.Contract.EtherBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherBridge *EtherBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherBridge *EtherBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherBridge *EtherBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherBridge.Contract.contract.Transact(opts, method, params...)
}

// DepositEther is a paid mutator transaction binding the contract method 0x77c76321.
//
// Solidity: function depositEther(string _odinAddress) payable returns(bool)
func (_EtherBridge *EtherBridgeTransactor) DepositEther(opts *bind.TransactOpts, _odinAddress string) (*types.Transaction, error) {
	return _EtherBridge.contract.Transact(opts, "depositEther", _odinAddress)
}

// DepositEther is a paid mutator transaction binding the contract method 0x77c76321.
//
// Solidity: function depositEther(string _odinAddress) payable returns(bool)
func (_EtherBridge *EtherBridgeSession) DepositEther(_odinAddress string) (*types.Transaction, error) {
	return _EtherBridge.Contract.DepositEther(&_EtherBridge.TransactOpts, _odinAddress)
}

// DepositEther is a paid mutator transaction binding the contract method 0x77c76321.
//
// Solidity: function depositEther(string _odinAddress) payable returns(bool)
func (_EtherBridge *EtherBridgeTransactorSession) DepositEther(_odinAddress string) (*types.Transaction, error) {
	return _EtherBridge.Contract.DepositEther(&_EtherBridge.TransactOpts, _odinAddress)
}

// EtherBridgeEtherDepositedIterator is returned from FilterEtherDeposited and is used to iterate over the raw logs and unpacked data for EtherDeposited events raised by the EtherBridge contract.
type EtherBridgeEtherDepositedIterator struct {
	Event *EtherBridgeEtherDeposited // Event containing the contract specifics and raw log

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
func (it *EtherBridgeEtherDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherBridgeEtherDeposited)
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
		it.Event = new(EtherBridgeEtherDeposited)
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
func (it *EtherBridgeEtherDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherBridgeEtherDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherBridgeEtherDeposited represents a EtherDeposited event raised by the EtherBridge contract.
type EtherBridgeEtherDeposited struct {
	UserAddress   common.Address
	OdinAddress   string
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEtherDeposited is a free log retrieval operation binding the contract event 0xf93ea55a2b7738a220c1c386239e26cf4448d5724777f7d2ffd8c6524ca0c37f.
//
// Solidity: event EtherDeposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount)
func (_EtherBridge *EtherBridgeFilterer) FilterEtherDeposited(opts *bind.FilterOpts, _userAddress []common.Address) (*EtherBridgeEtherDepositedIterator, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _EtherBridge.contract.FilterLogs(opts, "EtherDeposited", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return &EtherBridgeEtherDepositedIterator{contract: _EtherBridge.contract, event: "EtherDeposited", logs: logs, sub: sub}, nil
}

// WatchEtherDeposited is a free log subscription operation binding the contract event 0xf93ea55a2b7738a220c1c386239e26cf4448d5724777f7d2ffd8c6524ca0c37f.
//
// Solidity: event EtherDeposited(address indexed _userAddress, string _odinAddress, uint256 _depositAmount)
func (_EtherBridge *EtherBridgeFilterer) WatchEtherDeposited(opts *bind.WatchOpts, sink chan<- *EtherBridgeEtherDeposited, _userAddress []common.Address) (event.Subscription, error) {

	var _userAddressRule []interface{}
	for _, _userAddressItem := range _userAddress {
		_userAddressRule = append(_userAddressRule, _userAddressItem)
	}

	logs, sub, err := _EtherBridge.contract.WatchLogs(opts, "EtherDeposited", _userAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherBridgeEtherDeposited)
				if err := _EtherBridge.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
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
func (_EtherBridge *EtherBridgeFilterer) ParseEtherDeposited(log types.Log) (*EtherBridgeEtherDeposited, error) {
	event := new(EtherBridgeEtherDeposited)
	if err := _EtherBridge.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
