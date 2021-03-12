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
var EtherBridgeBin = "0x608060405234801561001057600080fd5b5061028d806100206000396000f3fe60806040526004361061001e5760003560e01c806377c7632114610023575b600080fd5b6100dc6004803603602081101561003957600080fd5b810190808035906020019064010000000081111561005657600080fd5b82018360208201111561006857600080fd5b8035906020019184600183028401116401000000008311171561008a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506100f4565b60405180821515815260200191505060405180910390f35b600080341161014e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260458152602001806102136045913960600191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167ff93ea55a2b7738a220c1c386239e26cf4448d5724777f7d2ffd8c6524ca0c37f83346040518080602001838152602001828103825284818151815260200191508051906020019080838360005b838110156101ce5780820151818401526020810190506101b3565b50505050905090810190601f1680156101fb5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a26001905091905056fe496e76616c69642076616c756520666f7220746865206465706f73697420616d6f756e742c206661696c656420746f206465706f7369742061207a65726f2076616c75652ea2646970667358221220874aa430443c3a581a3771183c56dc548a0cc96994e7984bffee1f675180546464736f6c63430007020033"

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
