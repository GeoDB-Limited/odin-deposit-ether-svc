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
const OdinBridgeABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"}],\"name\":\"UserDeposited\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_odinAddress\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// OdinBridgeBin is the compiled bytecode used for deploying new contracts.
var OdinBridgeBin = "0x608060405234801561001057600080fd5b5061028d806100206000396000f3fe60806040526004361061001e5760003560e01c8063a26e118614610023575b600080fd5b6100dc6004803603602081101561003957600080fd5b810190808035906020019064010000000081111561005657600080fd5b82018360208201111561006857600080fd5b8035906020019184600183028401116401000000008311171561008a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506100f4565b60405180821515815260200191505060405180910390f35b600080341161014e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260458152602001806102136045913960600191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f87e400cb9eb94227f5b3f6cf51a344b844eb08705e81b1291f66bb61f675ff8434846040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101ce5780820151818401526020810190506101b3565b50505050905090810190601f1680156101fb5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a26001905091905056fe496e76616c69642076616c756520666f7220746865206465706f73697420616d6f756e742c206661696c656420746f206465706f7369742061207a65726f2076616c75652ea264697066735822122069f60065e9ab474523ac453d5e1c165900a87796b53a9ea1f95ef230b8d6a19164736f6c63430007050033"

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

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string _odinAddress) payable returns(bool)
func (_OdinBridge *OdinBridgeTransactor) Deposit(opts *bind.TransactOpts, _odinAddress string) (*types.Transaction, error) {
	return _OdinBridge.contract.Transact(opts, "deposit", _odinAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string _odinAddress) payable returns(bool)
func (_OdinBridge *OdinBridgeSession) Deposit(_odinAddress string) (*types.Transaction, error) {
	return _OdinBridge.Contract.Deposit(&_OdinBridge.TransactOpts, _odinAddress)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string _odinAddress) payable returns(bool)
func (_OdinBridge *OdinBridgeTransactorSession) Deposit(_odinAddress string) (*types.Transaction, error) {
	return _OdinBridge.Contract.Deposit(&_OdinBridge.TransactOpts, _odinAddress)
}

// OdinBridgeUserDepositedIterator is returned from FilterUserDeposited and is used to iterate over the raw logs and unpacked data for UserDeposited events raised by the OdinBridge contract.
type OdinBridgeUserDepositedIterator struct {
	Event *OdinBridgeUserDeposited // Event containing the contract specifics and raw log

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
func (it *OdinBridgeUserDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OdinBridgeUserDeposited)
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
		it.Event = new(OdinBridgeUserDeposited)
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
func (it *OdinBridgeUserDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OdinBridgeUserDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OdinBridgeUserDeposited represents a UserDeposited event raised by the OdinBridge contract.
type OdinBridgeUserDeposited struct {
	User          common.Address
	DepositAmount *big.Int
	OdinAddress   string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserDeposited is a free log retrieval operation binding the contract event 0x87e400cb9eb94227f5b3f6cf51a344b844eb08705e81b1291f66bb61f675ff84.
//
// Solidity: event UserDeposited(address indexed _user, uint256 _depositAmount, string _odinAddress)
func (_OdinBridge *OdinBridgeFilterer) FilterUserDeposited(opts *bind.FilterOpts, _user []common.Address) (*OdinBridgeUserDepositedIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _OdinBridge.contract.FilterLogs(opts, "UserDeposited", _userRule)
	if err != nil {
		return nil, err
	}
	return &OdinBridgeUserDepositedIterator{contract: _OdinBridge.contract, event: "UserDeposited", logs: logs, sub: sub}, nil
}

// WatchUserDeposited is a free log subscription operation binding the contract event 0x87e400cb9eb94227f5b3f6cf51a344b844eb08705e81b1291f66bb61f675ff84.
//
// Solidity: event UserDeposited(address indexed _user, uint256 _depositAmount, string _odinAddress)
func (_OdinBridge *OdinBridgeFilterer) WatchUserDeposited(opts *bind.WatchOpts, sink chan<- *OdinBridgeUserDeposited, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _OdinBridge.contract.WatchLogs(opts, "UserDeposited", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OdinBridgeUserDeposited)
				if err := _OdinBridge.contract.UnpackLog(event, "UserDeposited", log); err != nil {
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

// ParseUserDeposited is a log parse operation binding the contract event 0x87e400cb9eb94227f5b3f6cf51a344b844eb08705e81b1291f66bb61f675ff84.
//
// Solidity: event UserDeposited(address indexed _user, uint256 _depositAmount, string _odinAddress)
func (_OdinBridge *OdinBridgeFilterer) ParseUserDeposited(log types.Log) (*OdinBridgeUserDeposited, error) {
	event := new(OdinBridgeUserDeposited)
	if err := _OdinBridge.contract.UnpackLog(event, "UserDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
