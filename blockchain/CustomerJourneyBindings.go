// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CustomerJourneyABI is the input ABI used to generate the binding from.
const CustomerJourneyABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deployedOn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"notaryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"initiator\",\"type\":\"address\"}],\"name\":\"JourneyStart\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"parent\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"typ\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"}],\"name\":\"AssetLinked\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"delUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"typ\",\"type\":\"uint8\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"Start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"parent\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"typ\",\"type\":\"uint8\"},{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"Link\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CustomerJourney is an auto generated Go binding around an Ethereum contract.
type CustomerJourney struct {
	CustomerJourneyCaller     // Read-only binding to the contract
	CustomerJourneyTransactor // Write-only binding to the contract
	CustomerJourneyFilterer   // Log filterer for contract events
}

// CustomerJourneyCaller is an auto generated read-only Go binding around an Ethereum contract.
type CustomerJourneyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomerJourneyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CustomerJourneyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomerJourneyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CustomerJourneyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomerJourneySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CustomerJourneySession struct {
	Contract     *CustomerJourney  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CustomerJourneyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CustomerJourneyCallerSession struct {
	Contract *CustomerJourneyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CustomerJourneyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CustomerJourneyTransactorSession struct {
	Contract     *CustomerJourneyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CustomerJourneyRaw is an auto generated low-level Go binding around an Ethereum contract.
type CustomerJourneyRaw struct {
	Contract *CustomerJourney // Generic contract binding to access the raw methods on
}

// CustomerJourneyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CustomerJourneyCallerRaw struct {
	Contract *CustomerJourneyCaller // Generic read-only contract binding to access the raw methods on
}

// CustomerJourneyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CustomerJourneyTransactorRaw struct {
	Contract *CustomerJourneyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCustomerJourney creates a new instance of CustomerJourney, bound to a specific deployed contract.
func NewCustomerJourney(address common.Address, backend bind.ContractBackend) (*CustomerJourney, error) {
	contract, err := bindCustomerJourney(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustomerJourney{CustomerJourneyCaller: CustomerJourneyCaller{contract: contract}, CustomerJourneyTransactor: CustomerJourneyTransactor{contract: contract}, CustomerJourneyFilterer: CustomerJourneyFilterer{contract: contract}}, nil
}

// NewCustomerJourneyCaller creates a new read-only instance of CustomerJourney, bound to a specific deployed contract.
func NewCustomerJourneyCaller(address common.Address, caller bind.ContractCaller) (*CustomerJourneyCaller, error) {
	contract, err := bindCustomerJourney(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustomerJourneyCaller{contract: contract}, nil
}

// NewCustomerJourneyTransactor creates a new write-only instance of CustomerJourney, bound to a specific deployed contract.
func NewCustomerJourneyTransactor(address common.Address, transactor bind.ContractTransactor) (*CustomerJourneyTransactor, error) {
	contract, err := bindCustomerJourney(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustomerJourneyTransactor{contract: contract}, nil
}

// NewCustomerJourneyFilterer creates a new log filterer instance of CustomerJourney, bound to a specific deployed contract.
func NewCustomerJourneyFilterer(address common.Address, filterer bind.ContractFilterer) (*CustomerJourneyFilterer, error) {
	contract, err := bindCustomerJourney(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustomerJourneyFilterer{contract: contract}, nil
}

// bindCustomerJourney binds a generic wrapper to an already deployed contract.
func bindCustomerJourney(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CustomerJourneyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustomerJourney *CustomerJourneyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CustomerJourney.Contract.CustomerJourneyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustomerJourney *CustomerJourneyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomerJourney.Contract.CustomerJourneyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustomerJourney *CustomerJourneyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomerJourney.Contract.CustomerJourneyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustomerJourney *CustomerJourneyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CustomerJourney.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustomerJourney *CustomerJourneyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomerJourney.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustomerJourney *CustomerJourneyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomerJourney.Contract.contract.Transact(opts, method, params...)
}

// Allowed is a free data retrieval call binding the contract method 0xd63a8e11.
//
// Solidity: function allowed(address ) constant returns(bool)
func (_CustomerJourney *CustomerJourneyCaller) Allowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _CustomerJourney.contract.Call(opts, out, "allowed", arg0)
	return *ret0, err
}

// Allowed is a free data retrieval call binding the contract method 0xd63a8e11.
//
// Solidity: function allowed(address ) constant returns(bool)
func (_CustomerJourney *CustomerJourneySession) Allowed(arg0 common.Address) (bool, error) {
	return _CustomerJourney.Contract.Allowed(&_CustomerJourney.CallOpts, arg0)
}

// Allowed is a free data retrieval call binding the contract method 0xd63a8e11.
//
// Solidity: function allowed(address ) constant returns(bool)
func (_CustomerJourney *CustomerJourneyCallerSession) Allowed(arg0 common.Address) (bool, error) {
	return _CustomerJourney.Contract.Allowed(&_CustomerJourney.CallOpts, arg0)
}

// DeployedOn is a free data retrieval call binding the contract method 0x73fc8420.
//
// Solidity: function deployedOn() constant returns(uint256)
func (_CustomerJourney *CustomerJourneyCaller) DeployedOn(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CustomerJourney.contract.Call(opts, out, "deployedOn")
	return *ret0, err
}

// DeployedOn is a free data retrieval call binding the contract method 0x73fc8420.
//
// Solidity: function deployedOn() constant returns(uint256)
func (_CustomerJourney *CustomerJourneySession) DeployedOn() (*big.Int, error) {
	return _CustomerJourney.Contract.DeployedOn(&_CustomerJourney.CallOpts)
}

// DeployedOn is a free data retrieval call binding the contract method 0x73fc8420.
//
// Solidity: function deployedOn() constant returns(uint256)
func (_CustomerJourney *CustomerJourneyCallerSession) DeployedOn() (*big.Int, error) {
	return _CustomerJourney.Contract.DeployedOn(&_CustomerJourney.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CustomerJourney *CustomerJourneyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CustomerJourney.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CustomerJourney *CustomerJourneySession) Owner() (common.Address, error) {
	return _CustomerJourney.Contract.Owner(&_CustomerJourney.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CustomerJourney *CustomerJourneyCallerSession) Owner() (common.Address, error) {
	return _CustomerJourney.Contract.Owner(&_CustomerJourney.CallOpts)
}

// Link is a paid mutator transaction binding the contract method 0xd16e024e.
//
// Solidity: function Link(bytes32 parent, bytes32 hash, uint8 typ, string key) returns()
func (_CustomerJourney *CustomerJourneyTransactor) Link(opts *bind.TransactOpts, parent [32]byte, hash [32]byte, typ uint8, key string) (*types.Transaction, error) {
	return _CustomerJourney.contract.Transact(opts, "Link", parent, hash, typ, key)
}

// Link is a paid mutator transaction binding the contract method 0xd16e024e.
//
// Solidity: function Link(bytes32 parent, bytes32 hash, uint8 typ, string key) returns()
func (_CustomerJourney *CustomerJourneySession) Link(parent [32]byte, hash [32]byte, typ uint8, key string) (*types.Transaction, error) {
	return _CustomerJourney.Contract.Link(&_CustomerJourney.TransactOpts, parent, hash, typ, key)
}

// Link is a paid mutator transaction binding the contract method 0xd16e024e.
//
// Solidity: function Link(bytes32 parent, bytes32 hash, uint8 typ, string key) returns()
func (_CustomerJourney *CustomerJourneyTransactorSession) Link(parent [32]byte, hash [32]byte, typ uint8, key string) (*types.Transaction, error) {
	return _CustomerJourney.Contract.Link(&_CustomerJourney.TransactOpts, parent, hash, typ, key)
}

// Start is a paid mutator transaction binding the contract method 0xba377a26.
//
// Solidity: function Start(bytes32 hash, uint8 typ, string key) returns()
func (_CustomerJourney *CustomerJourneyTransactor) Start(opts *bind.TransactOpts, hash [32]byte, typ uint8, key string) (*types.Transaction, error) {
	return _CustomerJourney.contract.Transact(opts, "Start", hash, typ, key)
}

// Start is a paid mutator transaction binding the contract method 0xba377a26.
//
// Solidity: function Start(bytes32 hash, uint8 typ, string key) returns()
func (_CustomerJourney *CustomerJourneySession) Start(hash [32]byte, typ uint8, key string) (*types.Transaction, error) {
	return _CustomerJourney.Contract.Start(&_CustomerJourney.TransactOpts, hash, typ, key)
}

// Start is a paid mutator transaction binding the contract method 0xba377a26.
//
// Solidity: function Start(bytes32 hash, uint8 typ, string key) returns()
func (_CustomerJourney *CustomerJourneyTransactorSession) Start(hash [32]byte, typ uint8, key string) (*types.Transaction, error) {
	return _CustomerJourney.Contract.Start(&_CustomerJourney.TransactOpts, hash, typ, key)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address user) returns()
func (_CustomerJourney *CustomerJourneyTransactor) AddUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _CustomerJourney.contract.Transact(opts, "addUser", user)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address user) returns()
func (_CustomerJourney *CustomerJourneySession) AddUser(user common.Address) (*types.Transaction, error) {
	return _CustomerJourney.Contract.AddUser(&_CustomerJourney.TransactOpts, user)
}

// AddUser is a paid mutator transaction binding the contract method 0x421b2d8b.
//
// Solidity: function addUser(address user) returns()
func (_CustomerJourney *CustomerJourneyTransactorSession) AddUser(user common.Address) (*types.Transaction, error) {
	return _CustomerJourney.Contract.AddUser(&_CustomerJourney.TransactOpts, user)
}

// DelUser is a paid mutator transaction binding the contract method 0xf849f379.
//
// Solidity: function delUser(address user) returns()
func (_CustomerJourney *CustomerJourneyTransactor) DelUser(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _CustomerJourney.contract.Transact(opts, "delUser", user)
}

// DelUser is a paid mutator transaction binding the contract method 0xf849f379.
//
// Solidity: function delUser(address user) returns()
func (_CustomerJourney *CustomerJourneySession) DelUser(user common.Address) (*types.Transaction, error) {
	return _CustomerJourney.Contract.DelUser(&_CustomerJourney.TransactOpts, user)
}

// DelUser is a paid mutator transaction binding the contract method 0xf849f379.
//
// Solidity: function delUser(address user) returns()
func (_CustomerJourney *CustomerJourneyTransactorSession) DelUser(user common.Address) (*types.Transaction, error) {
	return _CustomerJourney.Contract.DelUser(&_CustomerJourney.TransactOpts, user)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_CustomerJourney *CustomerJourneyTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomerJourney.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_CustomerJourney *CustomerJourneySession) Kill() (*types.Transaction, error) {
	return _CustomerJourney.Contract.Kill(&_CustomerJourney.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_CustomerJourney *CustomerJourneyTransactorSession) Kill() (*types.Transaction, error) {
	return _CustomerJourney.Contract.Kill(&_CustomerJourney.TransactOpts)
}

// CustomerJourneyAssetLinkedIterator is returned from FilterAssetLinked and is used to iterate over the raw logs and unpacked data for AssetLinked events raised by the CustomerJourney contract.
type CustomerJourneyAssetLinkedIterator struct {
	Event *CustomerJourneyAssetLinked // Event containing the contract specifics and raw log

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
func (it *CustomerJourneyAssetLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomerJourneyAssetLinked)
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
		it.Event = new(CustomerJourneyAssetLinked)
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
func (it *CustomerJourneyAssetLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomerJourneyAssetLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomerJourneyAssetLinked represents a AssetLinked event raised by the CustomerJourney contract.
type CustomerJourneyAssetLinked struct {
	Parent [32]byte
	Hash   [32]byte
	Typ    uint8
	Key    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAssetLinked is a free log retrieval operation binding the contract event 0x6f7f2cbbfe35924aacc98903ecaca5d09f1a3a29ba52182e6c5221b181c8b01c.
//
// Solidity: event AssetLinked(bytes32 indexed parent, bytes32 indexed hash, uint8 typ, string key)
func (_CustomerJourney *CustomerJourneyFilterer) FilterAssetLinked(opts *bind.FilterOpts, parent [][32]byte, hash [][32]byte) (*CustomerJourneyAssetLinkedIterator, error) {

	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _CustomerJourney.contract.FilterLogs(opts, "AssetLinked", parentRule, hashRule)
	if err != nil {
		return nil, err
	}
	return &CustomerJourneyAssetLinkedIterator{contract: _CustomerJourney.contract, event: "AssetLinked", logs: logs, sub: sub}, nil
}

// WatchAssetLinked is a free log subscription operation binding the contract event 0x6f7f2cbbfe35924aacc98903ecaca5d09f1a3a29ba52182e6c5221b181c8b01c.
//
// Solidity: event AssetLinked(bytes32 indexed parent, bytes32 indexed hash, uint8 typ, string key)
func (_CustomerJourney *CustomerJourneyFilterer) WatchAssetLinked(opts *bind.WatchOpts, sink chan<- *CustomerJourneyAssetLinked, parent [][32]byte, hash [][32]byte) (event.Subscription, error) {

	var parentRule []interface{}
	for _, parentItem := range parent {
		parentRule = append(parentRule, parentItem)
	}
	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _CustomerJourney.contract.WatchLogs(opts, "AssetLinked", parentRule, hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomerJourneyAssetLinked)
				if err := _CustomerJourney.contract.UnpackLog(event, "AssetLinked", log); err != nil {
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

// CustomerJourneyJourneyStartIterator is returned from FilterJourneyStart and is used to iterate over the raw logs and unpacked data for JourneyStart events raised by the CustomerJourney contract.
type CustomerJourneyJourneyStartIterator struct {
	Event *CustomerJourneyJourneyStart // Event containing the contract specifics and raw log

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
func (it *CustomerJourneyJourneyStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomerJourneyJourneyStart)
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
		it.Event = new(CustomerJourneyJourneyStart)
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
func (it *CustomerJourneyJourneyStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomerJourneyJourneyStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomerJourneyJourneyStart represents a JourneyStart event raised by the CustomerJourney contract.
type CustomerJourneyJourneyStart struct {
	Hash      [32]byte
	Typ       uint8
	Key       string
	Initiator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterJourneyStart is a free log retrieval operation binding the contract event 0x955de9d1da36fa142798faee4d3d3162c183e1bfe74cc06d3304c58e8771b1db.
//
// Solidity: event JourneyStart(bytes32 indexed hash, uint8 typ, string key, address indexed initiator)
func (_CustomerJourney *CustomerJourneyFilterer) FilterJourneyStart(opts *bind.FilterOpts, hash [][32]byte, initiator []common.Address) (*CustomerJourneyJourneyStartIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _CustomerJourney.contract.FilterLogs(opts, "JourneyStart", hashRule, initiatorRule)
	if err != nil {
		return nil, err
	}
	return &CustomerJourneyJourneyStartIterator{contract: _CustomerJourney.contract, event: "JourneyStart", logs: logs, sub: sub}, nil
}

// WatchJourneyStart is a free log subscription operation binding the contract event 0x955de9d1da36fa142798faee4d3d3162c183e1bfe74cc06d3304c58e8771b1db.
//
// Solidity: event JourneyStart(bytes32 indexed hash, uint8 typ, string key, address indexed initiator)
func (_CustomerJourney *CustomerJourneyFilterer) WatchJourneyStart(opts *bind.WatchOpts, sink chan<- *CustomerJourneyJourneyStart, hash [][32]byte, initiator []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _CustomerJourney.contract.WatchLogs(opts, "JourneyStart", hashRule, initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomerJourneyJourneyStart)
				if err := _CustomerJourney.contract.UnpackLog(event, "JourneyStart", log); err != nil {
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
