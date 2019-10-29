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

// DeliverNotaryABI is the input ABI used to generate the binding from.
const DeliverNotaryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x41c0e1b5\"},{\"constant\":false,\"inputs\":[{\"name\":\"hashes\",\"type\":\"bytes32[]\"}],\"name\":\"cancelTransferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x4bf87862\"},{\"constant\":true,\"inputs\":[],\"name\":\"deployedOn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x73fc8420\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingTransfers\",\"outputs\":[{\"name\":\"network\",\"type\":\"uint8\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"app\",\"type\":\"uint8\"},{\"name\":\"pathSpec\",\"type\":\"string\"},{\"name\":\"expiresAt\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8d4ab2c0\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8da5cb5b\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"assets\",\"outputs\":[{\"name\":\"network\",\"type\":\"uint8\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"app\",\"type\":\"uint8\"},{\"name\":\"pathSpec\",\"type\":\"string\"},{\"name\":\"expiresAt\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x9fda5b66\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingTransfersTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xed5decd6\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"signature\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"network\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AssetRegistered\",\"type\":\"event\",\"signature\":\"0x13b3af77244e191e093ea8049c1101ea5138cbfcef94b18a2ba38a5d4cf11a87\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"previousHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"network\",\"type\":\"uint8\"}],\"name\":\"AssetUpdated\",\"type\":\"event\",\"signature\":\"0xb5d4eba2eacb304e955df1375729746034777183757f10bdcb249c2eaa76038e\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"network\",\"type\":\"uint8\"}],\"name\":\"AssetNewClaim\",\"type\":\"event\",\"signature\":\"0x7997f98c4555f29a1b69504bcfc3c1c6ec6ed49580cc4264856a4bb6348137aa\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"AssetClaimRejected\",\"type\":\"event\",\"signature\":\"0x6137dc9923baf8633ade74cf103f8dd5f29a8948866faa456fd527e5aa232613\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"oldNetwork\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newNetwork\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"AssetOwnershipTransferred\",\"type\":\"event\",\"signature\":\"0x4e9328cb9fd3ff9c95b201e40c1741b959429f623a270159d664d7448d10e44c\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"AssetDeleted\",\"type\":\"event\",\"signature\":\"0x22ced05e2a8c120ab8dbb2d8bff75c90683795bba4268b3e549ed50542cc78c1\"},{\"constant\":false,\"inputs\":[{\"name\":\"myNetwork\",\"type\":\"uint8\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"app\",\"type\":\"uint8\"},{\"name\":\"pathSpec\",\"type\":\"string\"},{\"name\":\"expiresAt\",\"type\":\"uint32\"}],\"name\":\"registerAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x9c7331cd\"},{\"constant\":false,\"inputs\":[{\"name\":\"myNetwork\",\"type\":\"uint8\"},{\"name\":\"currentHash\",\"type\":\"bytes32\"},{\"name\":\"newHash\",\"type\":\"bytes32\"},{\"name\":\"app\",\"type\":\"uint8\"},{\"name\":\"pathSpec\",\"type\":\"string\"},{\"name\":\"expiresAt\",\"type\":\"uint32\"}],\"name\":\"updateAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xc63a4d42\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"myNetwork\",\"type\":\"uint8\"},{\"name\":\"newNetwork\",\"type\":\"uint8\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x2e2a9280\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"myNetwork\",\"type\":\"uint8\"},{\"name\":\"app\",\"type\":\"uint8\"},{\"name\":\"pathSpec\",\"type\":\"string\"},{\"name\":\"expiresAt\",\"type\":\"uint32\"}],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xb5726202\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"myNetwork\",\"type\":\"uint8\"}],\"name\":\"rejectOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x151b29d0\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"myNetwork\",\"type\":\"uint8\"},{\"name\":\"app\",\"type\":\"uint8\"},{\"name\":\"pathSpec\",\"type\":\"string\"},{\"name\":\"expiresAt\",\"type\":\"uint32\"}],\"name\":\"claimAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x9f5d6fc2\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"myNetwork\",\"type\":\"uint8\"}],\"name\":\"acceptClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x49f60a82\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"myNetwork\",\"type\":\"uint8\"}],\"name\":\"rejectClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xd278d864\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"myNetwork\",\"type\":\"uint8\"}],\"name\":\"deleteAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xec11498e\"}]"

// DeliverNotary is an auto generated Go binding around an Ethereum contract.
type DeliverNotary struct {
	DeliverNotaryCaller     // Read-only binding to the contract
	DeliverNotaryTransactor // Write-only binding to the contract
	DeliverNotaryFilterer   // Log filterer for contract events
}

// DeliverNotaryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeliverNotaryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeliverNotaryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeliverNotaryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeliverNotaryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeliverNotaryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeliverNotarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeliverNotarySession struct {
	Contract     *DeliverNotary    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeliverNotaryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeliverNotaryCallerSession struct {
	Contract *DeliverNotaryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DeliverNotaryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeliverNotaryTransactorSession struct {
	Contract     *DeliverNotaryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DeliverNotaryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeliverNotaryRaw struct {
	Contract *DeliverNotary // Generic contract binding to access the raw methods on
}

// DeliverNotaryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeliverNotaryCallerRaw struct {
	Contract *DeliverNotaryCaller // Generic read-only contract binding to access the raw methods on
}

// DeliverNotaryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeliverNotaryTransactorRaw struct {
	Contract *DeliverNotaryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeliverNotary creates a new instance of DeliverNotary, bound to a specific deployed contract.
func NewDeliverNotary(address common.Address, backend bind.ContractBackend) (*DeliverNotary, error) {
	contract, err := bindDeliverNotary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeliverNotary{DeliverNotaryCaller: DeliverNotaryCaller{contract: contract}, DeliverNotaryTransactor: DeliverNotaryTransactor{contract: contract}, DeliverNotaryFilterer: DeliverNotaryFilterer{contract: contract}}, nil
}

// NewDeliverNotaryCaller creates a new read-only instance of DeliverNotary, bound to a specific deployed contract.
func NewDeliverNotaryCaller(address common.Address, caller bind.ContractCaller) (*DeliverNotaryCaller, error) {
	contract, err := bindDeliverNotary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeliverNotaryCaller{contract: contract}, nil
}

// NewDeliverNotaryTransactor creates a new write-only instance of DeliverNotary, bound to a specific deployed contract.
func NewDeliverNotaryTransactor(address common.Address, transactor bind.ContractTransactor) (*DeliverNotaryTransactor, error) {
	contract, err := bindDeliverNotary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeliverNotaryTransactor{contract: contract}, nil
}

// NewDeliverNotaryFilterer creates a new log filterer instance of DeliverNotary, bound to a specific deployed contract.
func NewDeliverNotaryFilterer(address common.Address, filterer bind.ContractFilterer) (*DeliverNotaryFilterer, error) {
	contract, err := bindDeliverNotary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeliverNotaryFilterer{contract: contract}, nil
}

// bindDeliverNotary binds a generic wrapper to an already deployed contract.
func bindDeliverNotary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeliverNotaryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeliverNotary *DeliverNotaryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DeliverNotary.Contract.DeliverNotaryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeliverNotary *DeliverNotaryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeliverNotary.Contract.DeliverNotaryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeliverNotary *DeliverNotaryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeliverNotary.Contract.DeliverNotaryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeliverNotary *DeliverNotaryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DeliverNotary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeliverNotary *DeliverNotaryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeliverNotary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeliverNotary *DeliverNotaryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeliverNotary.Contract.contract.Transact(opts, method, params...)
}

// Assets is a free data retrieval call binding the contract method 0x9fda5b66.
//
// Solidity: function assets(bytes32 ) constant returns(uint8 network, address owner, uint8 app, string pathSpec, uint32 expiresAt)
func (_DeliverNotary *DeliverNotaryCaller) Assets(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Network   uint8
	Owner     common.Address
	App       uint8
	PathSpec  string
	ExpiresAt uint32
}, error) {
	ret := new(struct {
		Network   uint8
		Owner     common.Address
		App       uint8
		PathSpec  string
		ExpiresAt uint32
	})
	out := ret
	err := _DeliverNotary.contract.Call(opts, out, "assets", arg0)
	return *ret, err
}

// Assets is a free data retrieval call binding the contract method 0x9fda5b66.
//
// Solidity: function assets(bytes32 ) constant returns(uint8 network, address owner, uint8 app, string pathSpec, uint32 expiresAt)
func (_DeliverNotary *DeliverNotarySession) Assets(arg0 [32]byte) (struct {
	Network   uint8
	Owner     common.Address
	App       uint8
	PathSpec  string
	ExpiresAt uint32
}, error) {
	return _DeliverNotary.Contract.Assets(&_DeliverNotary.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0x9fda5b66.
//
// Solidity: function assets(bytes32 ) constant returns(uint8 network, address owner, uint8 app, string pathSpec, uint32 expiresAt)
func (_DeliverNotary *DeliverNotaryCallerSession) Assets(arg0 [32]byte) (struct {
	Network   uint8
	Owner     common.Address
	App       uint8
	PathSpec  string
	ExpiresAt uint32
}, error) {
	return _DeliverNotary.Contract.Assets(&_DeliverNotary.CallOpts, arg0)
}

// DeployedOn is a free data retrieval call binding the contract method 0x73fc8420.
//
// Solidity: function deployedOn() constant returns(uint256)
func (_DeliverNotary *DeliverNotaryCaller) DeployedOn(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DeliverNotary.contract.Call(opts, out, "deployedOn")
	return *ret0, err
}

// DeployedOn is a free data retrieval call binding the contract method 0x73fc8420.
//
// Solidity: function deployedOn() constant returns(uint256)
func (_DeliverNotary *DeliverNotarySession) DeployedOn() (*big.Int, error) {
	return _DeliverNotary.Contract.DeployedOn(&_DeliverNotary.CallOpts)
}

// DeployedOn is a free data retrieval call binding the contract method 0x73fc8420.
//
// Solidity: function deployedOn() constant returns(uint256)
func (_DeliverNotary *DeliverNotaryCallerSession) DeployedOn() (*big.Int, error) {
	return _DeliverNotary.Contract.DeployedOn(&_DeliverNotary.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DeliverNotary *DeliverNotaryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DeliverNotary.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DeliverNotary *DeliverNotarySession) Owner() (common.Address, error) {
	return _DeliverNotary.Contract.Owner(&_DeliverNotary.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DeliverNotary *DeliverNotaryCallerSession) Owner() (common.Address, error) {
	return _DeliverNotary.Contract.Owner(&_DeliverNotary.CallOpts)
}

// PendingTransfers is a free data retrieval call binding the contract method 0x8d4ab2c0.
//
// Solidity: function pendingTransfers(bytes32 ) constant returns(uint8 network, address owner, uint8 app, string pathSpec, uint32 expiresAt)
func (_DeliverNotary *DeliverNotaryCaller) PendingTransfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Network   uint8
	Owner     common.Address
	App       uint8
	PathSpec  string
	ExpiresAt uint32
}, error) {
	ret := new(struct {
		Network   uint8
		Owner     common.Address
		App       uint8
		PathSpec  string
		ExpiresAt uint32
	})
	out := ret
	err := _DeliverNotary.contract.Call(opts, out, "pendingTransfers", arg0)
	return *ret, err
}

// PendingTransfers is a free data retrieval call binding the contract method 0x8d4ab2c0.
//
// Solidity: function pendingTransfers(bytes32 ) constant returns(uint8 network, address owner, uint8 app, string pathSpec, uint32 expiresAt)
func (_DeliverNotary *DeliverNotarySession) PendingTransfers(arg0 [32]byte) (struct {
	Network   uint8
	Owner     common.Address
	App       uint8
	PathSpec  string
	ExpiresAt uint32
}, error) {
	return _DeliverNotary.Contract.PendingTransfers(&_DeliverNotary.CallOpts, arg0)
}

// PendingTransfers is a free data retrieval call binding the contract method 0x8d4ab2c0.
//
// Solidity: function pendingTransfers(bytes32 ) constant returns(uint8 network, address owner, uint8 app, string pathSpec, uint32 expiresAt)
func (_DeliverNotary *DeliverNotaryCallerSession) PendingTransfers(arg0 [32]byte) (struct {
	Network   uint8
	Owner     common.Address
	App       uint8
	PathSpec  string
	ExpiresAt uint32
}, error) {
	return _DeliverNotary.Contract.PendingTransfers(&_DeliverNotary.CallOpts, arg0)
}

// PendingTransfersTimestamp is a free data retrieval call binding the contract method 0xed5decd6.
//
// Solidity: function pendingTransfersTimestamp(bytes32 ) constant returns(uint256)
func (_DeliverNotary *DeliverNotaryCaller) PendingTransfersTimestamp(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DeliverNotary.contract.Call(opts, out, "pendingTransfersTimestamp", arg0)
	return *ret0, err
}

// PendingTransfersTimestamp is a free data retrieval call binding the contract method 0xed5decd6.
//
// Solidity: function pendingTransfersTimestamp(bytes32 ) constant returns(uint256)
func (_DeliverNotary *DeliverNotarySession) PendingTransfersTimestamp(arg0 [32]byte) (*big.Int, error) {
	return _DeliverNotary.Contract.PendingTransfersTimestamp(&_DeliverNotary.CallOpts, arg0)
}

// PendingTransfersTimestamp is a free data retrieval call binding the contract method 0xed5decd6.
//
// Solidity: function pendingTransfersTimestamp(bytes32 ) constant returns(uint256)
func (_DeliverNotary *DeliverNotaryCallerSession) PendingTransfersTimestamp(arg0 [32]byte) (*big.Int, error) {
	return _DeliverNotary.Contract.PendingTransfersTimestamp(&_DeliverNotary.CallOpts, arg0)
}

// AcceptClaim is a paid mutator transaction binding the contract method 0x49f60a82.
//
// Solidity: function acceptClaim(bytes32 hash, uint8 myNetwork) returns()
func (_DeliverNotary *DeliverNotaryTransactor) AcceptClaim(opts *bind.TransactOpts, hash [32]byte, myNetwork uint8) (*types.Transaction, error) {
	return _DeliverNotary.contract.Transact(opts, "acceptClaim", hash, myNetwork)
}

// AcceptClaim is a paid mutator transaction binding the contract method 0x49f60a82.
//
// Solidity: function acceptClaim(bytes32 hash, uint8 myNetwork) returns()
func (_DeliverNotary *DeliverNotarySession) AcceptClaim(hash [32]byte, myNetwork uint8) (*types.Transaction, error) {
	return _DeliverNotary.Contract.AcceptClaim(&_DeliverNotary.TransactOpts, hash, myNetwork)
}

// AcceptClaim is a paid mutator transaction binding the contract method 0x49f60a82.
//
// Solidity: function acceptClaim(bytes32 hash, uint8 myNetwork) returns()
func (_DeliverNotary *DeliverNotaryTransactorSession) AcceptClaim(hash [32]byte, myNetwork uint8) (*types.Transaction, error) {
	return _DeliverNotary.Contract.AcceptClaim(&_DeliverNotary.TransactOpts, hash, myNetwork)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0xb5726202.
//
// Solidity: function acceptOwnership(bytes32 hash, uint8 myNetwork, uint8 app, string pathSpec, uint32 expiresAt) returns()
func (_DeliverNotary *DeliverNotaryTransactor) AcceptOwnership(opts *bind.TransactOpts, hash [32]byte, myNetwork uint8, app uint8, pathSpec string, expiresAt uint32) (*types.Transaction, error) {
	return _DeliverNotary.contract.Transact(opts, "acceptOwnership", hash, myNetwork, app, pathSpec, expiresAt)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0xb5726202.
//
// Solidity: function acceptOwnership(bytes32 hash, uint8 myNetwork, uint8 app, string pathSpec, uint32 expiresAt) returns()
func (_DeliverNotary *DeliverNotarySession) AcceptOwnership(hash [32]byte, myNetwork uint8, app uint8, pathSpec string, expiresAt uint32) (*types.Transaction, error) {
	return _DeliverNotary.Contract.AcceptOwnership(&_DeliverNotary.TransactOpts, hash, myNetwork, app, pathSpec, expiresAt)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0xb5726202.
//
// Solidity: function acceptOwnership(bytes32 hash, uint8 myNetwork, uint8 app, string pathSpec, uint32 expiresAt) returns()
func (_DeliverNotary *DeliverNotaryTransactorSession) AcceptOwnership(hash [32]byte, myNetwork uint8, app uint8, pathSpec string, expiresAt uint32) (*types.Transaction, error) {
	return _DeliverNotary.Contract.AcceptOwnership(&_DeliverNotary.TransactOpts, hash, myNetwork, app, pathSpec, expiresAt)
}

// CancelTransferOwnership is a paid mutator transaction binding the contract method 0x4bf87862.
//
// Solidity: function cancelTransferOwnership(bytes32[] hashes) returns()
func (_DeliverNotary *DeliverNotaryTransactor) CancelTransferOwnership(opts *bind.TransactOpts, hashes [][32]byte) (*types.Transaction, error) {
	return _DeliverNotary.contract.Transact(opts, "cancelTransferOwnership", hashes)
}

// CancelTransferOwnership is a paid mutator transaction binding the contract method 0x4bf87862.
//
// Solidity: function cancelTransferOwnership(bytes32[] hashes) returns()
func (_DeliverNotary *DeliverNotarySession) CancelTransferOwnership(hashes [][32]byte) (*types.Transaction, error) {
	return _DeliverNotary.Contract.CancelTransferOwnership(&_DeliverNotary.TransactOpts, hashes)
}

// CancelTransferOwnership is a paid mutator transaction binding the contract method 0x4bf87862.
//
// Solidity: function cancelTransferOwnership(bytes32[] hashes) returns()
func (_DeliverNotary *DeliverNotaryTransactorSession) CancelTransferOwnership(hashes [][32]byte) (*types.Transaction, error) {
	return _DeliverNotary.Contract.CancelTransferOwnership(&_DeliverNotary.TransactOpts, hashes)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0x9f5d6fc2.
//
// Solidity: function claimAsset(bytes32 hash, uint8 myNetwork, uint8 app, string pathSpec, uint32 expiresAt) returns()
func (_DeliverNotary *DeliverNotaryTransactor) ClaimAsset(opts *bind.TransactOpts, hash [32]byte, myNetwork uint8, app uint8, pathSpec string, expiresAt uint32) (*types.Transaction, error) {
	return _DeliverNotary.contract.Transact(opts, "claimAsset", hash, myNetwork, app, pathSpec, expiresAt)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0x9f5d6fc2.
//
// Solidity: function claimAsset(bytes32 hash, uint8 myNetwork, uint8 app, string pathSpec, uint32 expiresAt) returns()
func (_DeliverNotary *DeliverNotarySession) ClaimAsset(hash [32]byte, myNetwork uint8, app uint8, pathSpec string, expiresAt uint32) (*types.Transaction, error) {
	return _DeliverNotary.Contract.ClaimAsset(&_DeliverNotary.TransactOpts, hash, myNetwork, app, pathSpec, expiresAt)
}

// ClaimAsset is a paid mutator transaction binding the contract method 0x9f5d6fc2.
//
// Solidity: function claimAsset(bytes32 hash, uint8 myNetwork, uint8 app, string pathSpec, uint32 expiresAt) returns()
func (_DeliverNotary *DeliverNotaryTransactorSession) ClaimAsset(hash [32]byte, myNetwork uint8, app uint8, pathSpec string, expiresAt uint32) (*types.Transaction, error) {
	return _DeliverNotary.Contract.ClaimAsset(&_DeliverNotary.TransactOpts, hash, myNetwork, app, pathSpec, expiresAt)
}

// DeleteAsset is a paid mutator transaction binding the contract method 0xec11498e.
//
// Solidity: function deleteAsset(bytes32 hash, uint8 myNetwork) returns()
func (_DeliverNotary *DeliverNotaryTransactor) DeleteAsset(opts *bind.TransactOpts, hash [32]byte, myNetwork uint8) (*types.Transaction, error) {
	return _DeliverNotary.contract.Transact(opts, "deleteAsset", hash, myNetwork)
}

// DeleteAsset is a paid mutator transaction binding the contract method 0xec11498e.
//
// Solidity: function deleteAsset(bytes32 hash, uint8 myNetwork) returns()
func (_DeliverNotary *DeliverNotarySession) DeleteAsset(hash [32]byte, myNetwork uint8) (*types.Transaction, error) {
	return _DeliverNotary.Contract.DeleteAsset(&_DeliverNotary.TransactOpts, hash, myNetwork)
}

// DeleteAsset is a paid mutator transaction binding the contract method 0xec11498e.
//
// Solidity: function deleteAsset(bytes32 hash, uint8 myNetwork) returns()
func (_DeliverNotary *DeliverNotaryTransactorSession) DeleteAsset(hash [32]byte, myNetwork uint8) (*types.Transaction, error) {
	return _DeliverNotary.Contract.DeleteAsset(&_DeliverNotary.TransactOpts, hash, myNetwork)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_DeliverNotary *DeliverNotaryTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeliverNotary.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_DeliverNotary *DeliverNotarySession) Kill() (*types.Transaction, error) {
	return _DeliverNotary.Contract.Kill(&_DeliverNotary.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_DeliverNotary *DeliverNotaryTransactorSession) Kill() (*types.Transaction, error) {
	return _DeliverNotary.Contract.Kill(&_DeliverNotary.TransactOpts)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0x9c7331cd.
//
// Solidity: function registerAsset(uint8 myNetwork, bytes32 hash, uint8 app, string pathSpec, uint32 expiresAt) returns()
func (_DeliverNotary *DeliverNotaryTransactor) RegisterAsset(opts *bind.TransactOpts, myNetwork uint8, hash [32]byte, app uint8, pathSpec string, expiresAt uint32) (*types.Transaction, error) {
	return _DeliverNotary.contract.Transact(opts, "registerAsset", myNetwork, hash, app, pathSpec, expiresAt)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0x9c7331cd.
//
// Solidity: function registerAsset(uint8 myNetwork, bytes32 hash, uint8 app, string pathSpec, uint32 expiresAt) returns()
func (_DeliverNotary *DeliverNotarySession) RegisterAsset(myNetwork uint8, hash [32]byte, app uint8, pathSpec string, expiresAt uint32) (*types.Transaction, error) {
	return _DeliverNotary.Contract.RegisterAsset(&_DeliverNotary.TransactOpts, myNetwork, hash, app, pathSpec, expiresAt)
}

// RegisterAsset is a paid mutator transaction binding the contract method 0x9c7331cd.
//
// Solidity: function registerAsset(uint8 myNetwork, bytes32 hash, uint8 app, string pathSpec, uint32 expiresAt) returns()
func (_DeliverNotary *DeliverNotaryTransactorSession) RegisterAsset(myNetwork uint8, hash [32]byte, app uint8, pathSpec string, expiresAt uint32) (*types.Transaction, error) {
	return _DeliverNotary.Contract.RegisterAsset(&_DeliverNotary.TransactOpts, myNetwork, hash, app, pathSpec, expiresAt)
}

// RejectClaim is a paid mutator transaction binding the contract method 0xd278d864.
//
// Solidity: function rejectClaim(bytes32 hash, uint8 myNetwork) returns()
func (_DeliverNotary *DeliverNotaryTransactor) RejectClaim(opts *bind.TransactOpts, hash [32]byte, myNetwork uint8) (*types.Transaction, error) {
	return _DeliverNotary.contract.Transact(opts, "rejectClaim", hash, myNetwork)
}

// RejectClaim is a paid mutator transaction binding the contract method 0xd278d864.
//
// Solidity: function rejectClaim(bytes32 hash, uint8 myNetwork) returns()
func (_DeliverNotary *DeliverNotarySession) RejectClaim(hash [32]byte, myNetwork uint8) (*types.Transaction, error) {
	return _DeliverNotary.Contract.RejectClaim(&_DeliverNotary.TransactOpts, hash, myNetwork)
}

// RejectClaim is a paid mutator transaction binding the contract method 0xd278d864.
//
// Solidity: function rejectClaim(bytes32 hash, uint8 myNetwork) returns()
func (_DeliverNotary *DeliverNotaryTransactorSession) RejectClaim(hash [32]byte, myNetwork uint8) (*types.Transaction, error) {
	return _DeliverNotary.Contract.RejectClaim(&_DeliverNotary.TransactOpts, hash, myNetwork)
}

// RejectOwnership is a paid mutator transaction binding the contract method 0x151b29d0.
//
// Solidity: function rejectOwnership(bytes32 hash, uint8 myNetwork) returns()
func (_DeliverNotary *DeliverNotaryTransactor) RejectOwnership(opts *bind.TransactOpts, hash [32]byte, myNetwork uint8) (*types.Transaction, error) {
	return _DeliverNotary.contract.Transact(opts, "rejectOwnership", hash, myNetwork)
}

// RejectOwnership is a paid mutator transaction binding the contract method 0x151b29d0.
//
// Solidity: function rejectOwnership(bytes32 hash, uint8 myNetwork) returns()
func (_DeliverNotary *DeliverNotarySession) RejectOwnership(hash [32]byte, myNetwork uint8) (*types.Transaction, error) {
	return _DeliverNotary.Contract.RejectOwnership(&_DeliverNotary.TransactOpts, hash, myNetwork)
}

// RejectOwnership is a paid mutator transaction binding the contract method 0x151b29d0.
//
// Solidity: function rejectOwnership(bytes32 hash, uint8 myNetwork) returns()
func (_DeliverNotary *DeliverNotaryTransactorSession) RejectOwnership(hash [32]byte, myNetwork uint8) (*types.Transaction, error) {
	return _DeliverNotary.Contract.RejectOwnership(&_DeliverNotary.TransactOpts, hash, myNetwork)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x2e2a9280.
//
// Solidity: function transferOwnership(bytes32 hash, uint8 myNetwork, uint8 newNetwork, address newOwner) returns()
func (_DeliverNotary *DeliverNotaryTransactor) TransferOwnership(opts *bind.TransactOpts, hash [32]byte, myNetwork uint8, newNetwork uint8, newOwner common.Address) (*types.Transaction, error) {
	return _DeliverNotary.contract.Transact(opts, "transferOwnership", hash, myNetwork, newNetwork, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x2e2a9280.
//
// Solidity: function transferOwnership(bytes32 hash, uint8 myNetwork, uint8 newNetwork, address newOwner) returns()
func (_DeliverNotary *DeliverNotarySession) TransferOwnership(hash [32]byte, myNetwork uint8, newNetwork uint8, newOwner common.Address) (*types.Transaction, error) {
	return _DeliverNotary.Contract.TransferOwnership(&_DeliverNotary.TransactOpts, hash, myNetwork, newNetwork, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x2e2a9280.
//
// Solidity: function transferOwnership(bytes32 hash, uint8 myNetwork, uint8 newNetwork, address newOwner) returns()
func (_DeliverNotary *DeliverNotaryTransactorSession) TransferOwnership(hash [32]byte, myNetwork uint8, newNetwork uint8, newOwner common.Address) (*types.Transaction, error) {
	return _DeliverNotary.Contract.TransferOwnership(&_DeliverNotary.TransactOpts, hash, myNetwork, newNetwork, newOwner)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0xc63a4d42.
//
// Solidity: function updateAsset(uint8 myNetwork, bytes32 currentHash, bytes32 newHash, uint8 app, string pathSpec, uint32 expiresAt) returns()
func (_DeliverNotary *DeliverNotaryTransactor) UpdateAsset(opts *bind.TransactOpts, myNetwork uint8, currentHash [32]byte, newHash [32]byte, app uint8, pathSpec string, expiresAt uint32) (*types.Transaction, error) {
	return _DeliverNotary.contract.Transact(opts, "updateAsset", myNetwork, currentHash, newHash, app, pathSpec, expiresAt)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0xc63a4d42.
//
// Solidity: function updateAsset(uint8 myNetwork, bytes32 currentHash, bytes32 newHash, uint8 app, string pathSpec, uint32 expiresAt) returns()
func (_DeliverNotary *DeliverNotarySession) UpdateAsset(myNetwork uint8, currentHash [32]byte, newHash [32]byte, app uint8, pathSpec string, expiresAt uint32) (*types.Transaction, error) {
	return _DeliverNotary.Contract.UpdateAsset(&_DeliverNotary.TransactOpts, myNetwork, currentHash, newHash, app, pathSpec, expiresAt)
}

// UpdateAsset is a paid mutator transaction binding the contract method 0xc63a4d42.
//
// Solidity: function updateAsset(uint8 myNetwork, bytes32 currentHash, bytes32 newHash, uint8 app, string pathSpec, uint32 expiresAt) returns()
func (_DeliverNotary *DeliverNotaryTransactorSession) UpdateAsset(myNetwork uint8, currentHash [32]byte, newHash [32]byte, app uint8, pathSpec string, expiresAt uint32) (*types.Transaction, error) {
	return _DeliverNotary.Contract.UpdateAsset(&_DeliverNotary.TransactOpts, myNetwork, currentHash, newHash, app, pathSpec, expiresAt)
}

// DeliverNotaryAssetClaimRejectedIterator is returned from FilterAssetClaimRejected and is used to iterate over the raw logs and unpacked data for AssetClaimRejected events raised by the DeliverNotary contract.
type DeliverNotaryAssetClaimRejectedIterator struct {
	Event *DeliverNotaryAssetClaimRejected // Event containing the contract specifics and raw log

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
func (it *DeliverNotaryAssetClaimRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeliverNotaryAssetClaimRejected)
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
		it.Event = new(DeliverNotaryAssetClaimRejected)
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
func (it *DeliverNotaryAssetClaimRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeliverNotaryAssetClaimRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeliverNotaryAssetClaimRejected represents a AssetClaimRejected event raised by the DeliverNotary contract.
type DeliverNotaryAssetClaimRejected struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAssetClaimRejected is a free log retrieval operation binding the contract event 0x6137dc9923baf8633ade74cf103f8dd5f29a8948866faa456fd527e5aa232613.
//
// Solidity: event AssetClaimRejected(bytes32 indexed hash)
func (_DeliverNotary *DeliverNotaryFilterer) FilterAssetClaimRejected(opts *bind.FilterOpts, hash [][32]byte) (*DeliverNotaryAssetClaimRejectedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _DeliverNotary.contract.FilterLogs(opts, "AssetClaimRejected", hashRule)
	if err != nil {
		return nil, err
	}
	return &DeliverNotaryAssetClaimRejectedIterator{contract: _DeliverNotary.contract, event: "AssetClaimRejected", logs: logs, sub: sub}, nil
}

// WatchAssetClaimRejected is a free log subscription operation binding the contract event 0x6137dc9923baf8633ade74cf103f8dd5f29a8948866faa456fd527e5aa232613.
//
// Solidity: event AssetClaimRejected(bytes32 indexed hash)
func (_DeliverNotary *DeliverNotaryFilterer) WatchAssetClaimRejected(opts *bind.WatchOpts, sink chan<- *DeliverNotaryAssetClaimRejected, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _DeliverNotary.contract.WatchLogs(opts, "AssetClaimRejected", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeliverNotaryAssetClaimRejected)
				if err := _DeliverNotary.contract.UnpackLog(event, "AssetClaimRejected", log); err != nil {
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

// DeliverNotaryAssetDeletedIterator is returned from FilterAssetDeleted and is used to iterate over the raw logs and unpacked data for AssetDeleted events raised by the DeliverNotary contract.
type DeliverNotaryAssetDeletedIterator struct {
	Event *DeliverNotaryAssetDeleted // Event containing the contract specifics and raw log

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
func (it *DeliverNotaryAssetDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeliverNotaryAssetDeleted)
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
		it.Event = new(DeliverNotaryAssetDeleted)
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
func (it *DeliverNotaryAssetDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeliverNotaryAssetDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeliverNotaryAssetDeleted represents a AssetDeleted event raised by the DeliverNotary contract.
type DeliverNotaryAssetDeleted struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAssetDeleted is a free log retrieval operation binding the contract event 0x22ced05e2a8c120ab8dbb2d8bff75c90683795bba4268b3e549ed50542cc78c1.
//
// Solidity: event AssetDeleted(bytes32 indexed hash)
func (_DeliverNotary *DeliverNotaryFilterer) FilterAssetDeleted(opts *bind.FilterOpts, hash [][32]byte) (*DeliverNotaryAssetDeletedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _DeliverNotary.contract.FilterLogs(opts, "AssetDeleted", hashRule)
	if err != nil {
		return nil, err
	}
	return &DeliverNotaryAssetDeletedIterator{contract: _DeliverNotary.contract, event: "AssetDeleted", logs: logs, sub: sub}, nil
}

// WatchAssetDeleted is a free log subscription operation binding the contract event 0x22ced05e2a8c120ab8dbb2d8bff75c90683795bba4268b3e549ed50542cc78c1.
//
// Solidity: event AssetDeleted(bytes32 indexed hash)
func (_DeliverNotary *DeliverNotaryFilterer) WatchAssetDeleted(opts *bind.WatchOpts, sink chan<- *DeliverNotaryAssetDeleted, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _DeliverNotary.contract.WatchLogs(opts, "AssetDeleted", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeliverNotaryAssetDeleted)
				if err := _DeliverNotary.contract.UnpackLog(event, "AssetDeleted", log); err != nil {
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

// DeliverNotaryAssetNewClaimIterator is returned from FilterAssetNewClaim and is used to iterate over the raw logs and unpacked data for AssetNewClaim events raised by the DeliverNotary contract.
type DeliverNotaryAssetNewClaimIterator struct {
	Event *DeliverNotaryAssetNewClaim // Event containing the contract specifics and raw log

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
func (it *DeliverNotaryAssetNewClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeliverNotaryAssetNewClaim)
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
		it.Event = new(DeliverNotaryAssetNewClaim)
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
func (it *DeliverNotaryAssetNewClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeliverNotaryAssetNewClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeliverNotaryAssetNewClaim represents a AssetNewClaim event raised by the DeliverNotary contract.
type DeliverNotaryAssetNewClaim struct {
	Hash    [32]byte
	Addr    common.Address
	Network uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetNewClaim is a free log retrieval operation binding the contract event 0x7997f98c4555f29a1b69504bcfc3c1c6ec6ed49580cc4264856a4bb6348137aa.
//
// Solidity: event AssetNewClaim(bytes32 hash, address indexed addr, uint8 indexed network)
func (_DeliverNotary *DeliverNotaryFilterer) FilterAssetNewClaim(opts *bind.FilterOpts, addr []common.Address, network []uint8) (*DeliverNotaryAssetNewClaimIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}

	logs, sub, err := _DeliverNotary.contract.FilterLogs(opts, "AssetNewClaim", addrRule, networkRule)
	if err != nil {
		return nil, err
	}
	return &DeliverNotaryAssetNewClaimIterator{contract: _DeliverNotary.contract, event: "AssetNewClaim", logs: logs, sub: sub}, nil
}

// WatchAssetNewClaim is a free log subscription operation binding the contract event 0x7997f98c4555f29a1b69504bcfc3c1c6ec6ed49580cc4264856a4bb6348137aa.
//
// Solidity: event AssetNewClaim(bytes32 hash, address indexed addr, uint8 indexed network)
func (_DeliverNotary *DeliverNotaryFilterer) WatchAssetNewClaim(opts *bind.WatchOpts, sink chan<- *DeliverNotaryAssetNewClaim, addr []common.Address, network []uint8) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var networkRule []interface{}
	for _, networkItem := range network {
		networkRule = append(networkRule, networkItem)
	}

	logs, sub, err := _DeliverNotary.contract.WatchLogs(opts, "AssetNewClaim", addrRule, networkRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeliverNotaryAssetNewClaim)
				if err := _DeliverNotary.contract.UnpackLog(event, "AssetNewClaim", log); err != nil {
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

// DeliverNotaryAssetOwnershipTransferredIterator is returned from FilterAssetOwnershipTransferred and is used to iterate over the raw logs and unpacked data for AssetOwnershipTransferred events raised by the DeliverNotary contract.
type DeliverNotaryAssetOwnershipTransferredIterator struct {
	Event *DeliverNotaryAssetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DeliverNotaryAssetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeliverNotaryAssetOwnershipTransferred)
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
		it.Event = new(DeliverNotaryAssetOwnershipTransferred)
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
func (it *DeliverNotaryAssetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeliverNotaryAssetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeliverNotaryAssetOwnershipTransferred represents a AssetOwnershipTransferred event raised by the DeliverNotary contract.
type DeliverNotaryAssetOwnershipTransferred struct {
	Hash       [32]byte
	OldNetwork uint8
	OldOwner   common.Address
	NewNetwork uint8
	NewOwner   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAssetOwnershipTransferred is a free log retrieval operation binding the contract event 0x4e9328cb9fd3ff9c95b201e40c1741b959429f623a270159d664d7448d10e44c.
//
// Solidity: event AssetOwnershipTransferred(bytes32 indexed hash, uint8 oldNetwork, address oldOwner, uint8 newNetwork, address newOwner)
func (_DeliverNotary *DeliverNotaryFilterer) FilterAssetOwnershipTransferred(opts *bind.FilterOpts, hash [][32]byte) (*DeliverNotaryAssetOwnershipTransferredIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _DeliverNotary.contract.FilterLogs(opts, "AssetOwnershipTransferred", hashRule)
	if err != nil {
		return nil, err
	}
	return &DeliverNotaryAssetOwnershipTransferredIterator{contract: _DeliverNotary.contract, event: "AssetOwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchAssetOwnershipTransferred is a free log subscription operation binding the contract event 0x4e9328cb9fd3ff9c95b201e40c1741b959429f623a270159d664d7448d10e44c.
//
// Solidity: event AssetOwnershipTransferred(bytes32 indexed hash, uint8 oldNetwork, address oldOwner, uint8 newNetwork, address newOwner)
func (_DeliverNotary *DeliverNotaryFilterer) WatchAssetOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DeliverNotaryAssetOwnershipTransferred, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _DeliverNotary.contract.WatchLogs(opts, "AssetOwnershipTransferred", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeliverNotaryAssetOwnershipTransferred)
				if err := _DeliverNotary.contract.UnpackLog(event, "AssetOwnershipTransferred", log); err != nil {
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

// DeliverNotaryAssetRegisteredIterator is returned from FilterAssetRegistered and is used to iterate over the raw logs and unpacked data for AssetRegistered events raised by the DeliverNotary contract.
type DeliverNotaryAssetRegisteredIterator struct {
	Event *DeliverNotaryAssetRegistered // Event containing the contract specifics and raw log

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
func (it *DeliverNotaryAssetRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeliverNotaryAssetRegistered)
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
		it.Event = new(DeliverNotaryAssetRegistered)
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
func (it *DeliverNotaryAssetRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeliverNotaryAssetRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeliverNotaryAssetRegistered represents a AssetRegistered event raised by the DeliverNotary contract.
type DeliverNotaryAssetRegistered struct {
	Hash    [32]byte
	Network uint8
	Owner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetRegistered is a free log retrieval operation binding the contract event 0x13b3af77244e191e093ea8049c1101ea5138cbfcef94b18a2ba38a5d4cf11a87.
//
// Solidity: event AssetRegistered(bytes32 indexed hash, uint8 network, address owner)
func (_DeliverNotary *DeliverNotaryFilterer) FilterAssetRegistered(opts *bind.FilterOpts, hash [][32]byte) (*DeliverNotaryAssetRegisteredIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _DeliverNotary.contract.FilterLogs(opts, "AssetRegistered", hashRule)
	if err != nil {
		return nil, err
	}
	return &DeliverNotaryAssetRegisteredIterator{contract: _DeliverNotary.contract, event: "AssetRegistered", logs: logs, sub: sub}, nil
}

// WatchAssetRegistered is a free log subscription operation binding the contract event 0x13b3af77244e191e093ea8049c1101ea5138cbfcef94b18a2ba38a5d4cf11a87.
//
// Solidity: event AssetRegistered(bytes32 indexed hash, uint8 network, address owner)
func (_DeliverNotary *DeliverNotaryFilterer) WatchAssetRegistered(opts *bind.WatchOpts, sink chan<- *DeliverNotaryAssetRegistered, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _DeliverNotary.contract.WatchLogs(opts, "AssetRegistered", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeliverNotaryAssetRegistered)
				if err := _DeliverNotary.contract.UnpackLog(event, "AssetRegistered", log); err != nil {
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

// DeliverNotaryAssetUpdatedIterator is returned from FilterAssetUpdated and is used to iterate over the raw logs and unpacked data for AssetUpdated events raised by the DeliverNotary contract.
type DeliverNotaryAssetUpdatedIterator struct {
	Event *DeliverNotaryAssetUpdated // Event containing the contract specifics and raw log

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
func (it *DeliverNotaryAssetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeliverNotaryAssetUpdated)
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
		it.Event = new(DeliverNotaryAssetUpdated)
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
func (it *DeliverNotaryAssetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeliverNotaryAssetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeliverNotaryAssetUpdated represents a AssetUpdated event raised by the DeliverNotary contract.
type DeliverNotaryAssetUpdated struct {
	NewHash      [32]byte
	PreviousHash [32]byte
	Network      uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAssetUpdated is a free log retrieval operation binding the contract event 0xb5d4eba2eacb304e955df1375729746034777183757f10bdcb249c2eaa76038e.
//
// Solidity: event AssetUpdated(bytes32 indexed newHash, bytes32 indexed previousHash, uint8 network)
func (_DeliverNotary *DeliverNotaryFilterer) FilterAssetUpdated(opts *bind.FilterOpts, newHash [][32]byte, previousHash [][32]byte) (*DeliverNotaryAssetUpdatedIterator, error) {

	var newHashRule []interface{}
	for _, newHashItem := range newHash {
		newHashRule = append(newHashRule, newHashItem)
	}
	var previousHashRule []interface{}
	for _, previousHashItem := range previousHash {
		previousHashRule = append(previousHashRule, previousHashItem)
	}

	logs, sub, err := _DeliverNotary.contract.FilterLogs(opts, "AssetUpdated", newHashRule, previousHashRule)
	if err != nil {
		return nil, err
	}
	return &DeliverNotaryAssetUpdatedIterator{contract: _DeliverNotary.contract, event: "AssetUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetUpdated is a free log subscription operation binding the contract event 0xb5d4eba2eacb304e955df1375729746034777183757f10bdcb249c2eaa76038e.
//
// Solidity: event AssetUpdated(bytes32 indexed newHash, bytes32 indexed previousHash, uint8 network)
func (_DeliverNotary *DeliverNotaryFilterer) WatchAssetUpdated(opts *bind.WatchOpts, sink chan<- *DeliverNotaryAssetUpdated, newHash [][32]byte, previousHash [][32]byte) (event.Subscription, error) {

	var newHashRule []interface{}
	for _, newHashItem := range newHash {
		newHashRule = append(newHashRule, newHashItem)
	}
	var previousHashRule []interface{}
	for _, previousHashItem := range previousHash {
		previousHashRule = append(previousHashRule, previousHashItem)
	}

	logs, sub, err := _DeliverNotary.contract.WatchLogs(opts, "AssetUpdated", newHashRule, previousHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeliverNotaryAssetUpdated)
				if err := _DeliverNotary.contract.UnpackLog(event, "AssetUpdated", log); err != nil {
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
