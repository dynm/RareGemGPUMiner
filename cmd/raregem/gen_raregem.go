// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// ProvablyRareGemABI is the input ABI used to generate the binding from.
const ProvablyRareGemABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"kind\",\"type\":\"uint256\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"miner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"kind\",\"type\":\"uint256\"}],\"name\":\"Mine\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kinds\",\"type\":\"uint256[]\"}],\"name\":\"acceptManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kind\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"craft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"color\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gemsPerMine\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"crafter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gemCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"gems\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"color\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"entropy\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gemsPerMine\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"crafter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pendingManager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kind\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"luck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kind\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"mine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kinds\",\"type\":\"uint256[]\"}],\"name\":\"renounceManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kinds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"kinds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"crafter\",\"type\":\"address\"}],\"name\":\"updateCrafter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kind\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"entropy\",\"type\":\"bytes32\"}],\"name\":\"updateEntropy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kind\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"color\",\"type\":\"string\"}],\"name\":\"updateGemInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kind\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gemsPerMine\",\"type\":\"uint256\"}],\"name\":\"updateMiningData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"kind\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProvablyRareGem is an auto generated Go binding around an Ethereum contract.
type ProvablyRareGem struct {
	ProvablyRareGemCaller     // Read-only binding to the contract
	ProvablyRareGemTransactor // Write-only binding to the contract
	ProvablyRareGemFilterer   // Log filterer for contract events
}

// ProvablyRareGemCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProvablyRareGemCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvablyRareGemTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProvablyRareGemTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvablyRareGemFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProvablyRareGemFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvablyRareGemSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProvablyRareGemSession struct {
	Contract     *ProvablyRareGem  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProvablyRareGemCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProvablyRareGemCallerSession struct {
	Contract *ProvablyRareGemCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProvablyRareGemTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProvablyRareGemTransactorSession struct {
	Contract     *ProvablyRareGemTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProvablyRareGemRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProvablyRareGemRaw struct {
	Contract *ProvablyRareGem // Generic contract binding to access the raw methods on
}

// ProvablyRareGemCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProvablyRareGemCallerRaw struct {
	Contract *ProvablyRareGemCaller // Generic read-only contract binding to access the raw methods on
}

// ProvablyRareGemTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProvablyRareGemTransactorRaw struct {
	Contract *ProvablyRareGemTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProvablyRareGem creates a new instance of ProvablyRareGem, bound to a specific deployed contract.
func NewProvablyRareGem(address common.Address, backend bind.ContractBackend) (*ProvablyRareGem, error) {
	contract, err := bindProvablyRareGem(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProvablyRareGem{ProvablyRareGemCaller: ProvablyRareGemCaller{contract: contract}, ProvablyRareGemTransactor: ProvablyRareGemTransactor{contract: contract}, ProvablyRareGemFilterer: ProvablyRareGemFilterer{contract: contract}}, nil
}

// NewProvablyRareGemCaller creates a new read-only instance of ProvablyRareGem, bound to a specific deployed contract.
func NewProvablyRareGemCaller(address common.Address, caller bind.ContractCaller) (*ProvablyRareGemCaller, error) {
	contract, err := bindProvablyRareGem(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProvablyRareGemCaller{contract: contract}, nil
}

// NewProvablyRareGemTransactor creates a new write-only instance of ProvablyRareGem, bound to a specific deployed contract.
func NewProvablyRareGemTransactor(address common.Address, transactor bind.ContractTransactor) (*ProvablyRareGemTransactor, error) {
	contract, err := bindProvablyRareGem(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProvablyRareGemTransactor{contract: contract}, nil
}

// NewProvablyRareGemFilterer creates a new log filterer instance of ProvablyRareGem, bound to a specific deployed contract.
func NewProvablyRareGemFilterer(address common.Address, filterer bind.ContractFilterer) (*ProvablyRareGemFilterer, error) {
	contract, err := bindProvablyRareGem(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProvablyRareGemFilterer{contract: contract}, nil
}

// bindProvablyRareGem binds a generic wrapper to an already deployed contract.
func bindProvablyRareGem(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProvablyRareGemABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProvablyRareGem *ProvablyRareGemRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProvablyRareGem.Contract.ProvablyRareGemCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProvablyRareGem *ProvablyRareGemRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.ProvablyRareGemTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProvablyRareGem *ProvablyRareGemRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.ProvablyRareGemTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProvablyRareGem *ProvablyRareGemCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProvablyRareGem.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProvablyRareGem *ProvablyRareGemTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProvablyRareGem *ProvablyRareGemTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ProvablyRareGem.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ProvablyRareGem.Contract.BalanceOf(&_ProvablyRareGem.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ProvablyRareGem.Contract.BalanceOf(&_ProvablyRareGem.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ProvablyRareGem *ProvablyRareGemCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ProvablyRareGem.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ProvablyRareGem *ProvablyRareGemSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ProvablyRareGem.Contract.BalanceOfBatch(&_ProvablyRareGem.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ProvablyRareGem *ProvablyRareGemCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ProvablyRareGem.Contract.BalanceOfBatch(&_ProvablyRareGem.CallOpts, accounts, ids)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_ProvablyRareGem *ProvablyRareGemCaller) Exists(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _ProvablyRareGem.contract.Call(opts, &out, "exists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_ProvablyRareGem *ProvablyRareGemSession) Exists(id *big.Int) (bool, error) {
	return _ProvablyRareGem.Contract.Exists(&_ProvablyRareGem.CallOpts, id)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 id) view returns(bool)
func (_ProvablyRareGem *ProvablyRareGemCallerSession) Exists(id *big.Int) (bool, error) {
	return _ProvablyRareGem.Contract.Exists(&_ProvablyRareGem.CallOpts, id)
}

// GemCount is a free data retrieval call binding the contract method 0x5b8bbc24.
//
// Solidity: function gemCount() view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemCaller) GemCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProvablyRareGem.contract.Call(opts, &out, "gemCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GemCount is a free data retrieval call binding the contract method 0x5b8bbc24.
//
// Solidity: function gemCount() view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemSession) GemCount() (*big.Int, error) {
	return _ProvablyRareGem.Contract.GemCount(&_ProvablyRareGem.CallOpts)
}

// GemCount is a free data retrieval call binding the contract method 0x5b8bbc24.
//
// Solidity: function gemCount() view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemCallerSession) GemCount() (*big.Int, error) {
	return _ProvablyRareGem.Contract.GemCount(&_ProvablyRareGem.CallOpts)
}

// Gems is a free data retrieval call binding the contract method 0xa1f0406d.
//
// Solidity: function gems(uint256 ) view returns(string name, string color, bytes32 entropy, uint256 difficulty, uint256 gemsPerMine, uint256 multiplier, address crafter, address manager, address pendingManager)
func (_ProvablyRareGem *ProvablyRareGemCaller) Gems(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name           string
	Color          string
	Entropy        [32]byte
	Difficulty     *big.Int
	GemsPerMine    *big.Int
	Multiplier     *big.Int
	Crafter        common.Address
	Manager        common.Address
	PendingManager common.Address
}, error) {
	var out []interface{}
	err := _ProvablyRareGem.contract.Call(opts, &out, "gems", arg0)

	outstruct := new(struct {
		Name           string
		Color          string
		Entropy        [32]byte
		Difficulty     *big.Int
		GemsPerMine    *big.Int
		Multiplier     *big.Int
		Crafter        common.Address
		Manager        common.Address
		PendingManager common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Color = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Entropy = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.Difficulty = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.GemsPerMine = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Multiplier = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Crafter = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Manager = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.PendingManager = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Gems is a free data retrieval call binding the contract method 0xa1f0406d.
//
// Solidity: function gems(uint256 ) view returns(string name, string color, bytes32 entropy, uint256 difficulty, uint256 gemsPerMine, uint256 multiplier, address crafter, address manager, address pendingManager)
func (_ProvablyRareGem *ProvablyRareGemSession) Gems(arg0 *big.Int) (struct {
	Name           string
	Color          string
	Entropy        [32]byte
	Difficulty     *big.Int
	GemsPerMine    *big.Int
	Multiplier     *big.Int
	Crafter        common.Address
	Manager        common.Address
	PendingManager common.Address
}, error) {
	return _ProvablyRareGem.Contract.Gems(&_ProvablyRareGem.CallOpts, arg0)
}

// Gems is a free data retrieval call binding the contract method 0xa1f0406d.
//
// Solidity: function gems(uint256 ) view returns(string name, string color, bytes32 entropy, uint256 difficulty, uint256 gemsPerMine, uint256 multiplier, address crafter, address manager, address pendingManager)
func (_ProvablyRareGem *ProvablyRareGemCallerSession) Gems(arg0 *big.Int) (struct {
	Name           string
	Color          string
	Entropy        [32]byte
	Difficulty     *big.Int
	GemsPerMine    *big.Int
	Multiplier     *big.Int
	Crafter        common.Address
	Manager        common.Address
	PendingManager common.Address
}, error) {
	return _ProvablyRareGem.Contract.Gems(&_ProvablyRareGem.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ProvablyRareGem *ProvablyRareGemCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ProvablyRareGem.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ProvablyRareGem *ProvablyRareGemSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ProvablyRareGem.Contract.IsApprovedForAll(&_ProvablyRareGem.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_ProvablyRareGem *ProvablyRareGemCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _ProvablyRareGem.Contract.IsApprovedForAll(&_ProvablyRareGem.CallOpts, account, operator)
}

// Luck is a free data retrieval call binding the contract method 0xbe71abe1.
//
// Solidity: function luck(uint256 kind, uint256 salt) view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemCaller) Luck(opts *bind.CallOpts, kind *big.Int, salt *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ProvablyRareGem.contract.Call(opts, &out, "luck", kind, salt)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Luck is a free data retrieval call binding the contract method 0xbe71abe1.
//
// Solidity: function luck(uint256 kind, uint256 salt) view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemSession) Luck(kind *big.Int, salt *big.Int) (*big.Int, error) {
	return _ProvablyRareGem.Contract.Luck(&_ProvablyRareGem.CallOpts, kind, salt)
}

// Luck is a free data retrieval call binding the contract method 0xbe71abe1.
//
// Solidity: function luck(uint256 kind, uint256 salt) view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemCallerSession) Luck(kind *big.Int, salt *big.Int) (*big.Int, error) {
	return _ProvablyRareGem.Contract.Luck(&_ProvablyRareGem.CallOpts, kind, salt)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProvablyRareGem.contract.Call(opts, &out, "nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _ProvablyRareGem.Contract.Nonce(&_ProvablyRareGem.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _ProvablyRareGem.Contract.Nonce(&_ProvablyRareGem.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProvablyRareGem *ProvablyRareGemCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ProvablyRareGem.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProvablyRareGem *ProvablyRareGemSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ProvablyRareGem.Contract.SupportsInterface(&_ProvablyRareGem.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProvablyRareGem *ProvablyRareGemCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ProvablyRareGem.Contract.SupportsInterface(&_ProvablyRareGem.CallOpts, interfaceId)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ProvablyRareGem.contract.Call(opts, &out, "totalSupply", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _ProvablyRareGem.Contract.TotalSupply(&_ProvablyRareGem.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemCallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _ProvablyRareGem.Contract.TotalSupply(&_ProvablyRareGem.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 kind) view returns(string)
func (_ProvablyRareGem *ProvablyRareGemCaller) Uri(opts *bind.CallOpts, kind *big.Int) (string, error) {
	var out []interface{}
	err := _ProvablyRareGem.contract.Call(opts, &out, "uri", kind)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 kind) view returns(string)
func (_ProvablyRareGem *ProvablyRareGemSession) Uri(kind *big.Int) (string, error) {
	return _ProvablyRareGem.Contract.Uri(&_ProvablyRareGem.CallOpts, kind)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 kind) view returns(string)
func (_ProvablyRareGem *ProvablyRareGemCallerSession) Uri(kind *big.Int) (string, error) {
	return _ProvablyRareGem.Contract.Uri(&_ProvablyRareGem.CallOpts, kind)
}

// AcceptManager is a paid mutator transaction binding the contract method 0xa41b41d2.
//
// Solidity: function acceptManager(uint256[] kinds) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactor) AcceptManager(opts *bind.TransactOpts, kinds []*big.Int) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "acceptManager", kinds)
}

// AcceptManager is a paid mutator transaction binding the contract method 0xa41b41d2.
//
// Solidity: function acceptManager(uint256[] kinds) returns()
func (_ProvablyRareGem *ProvablyRareGemSession) AcceptManager(kinds []*big.Int) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.AcceptManager(&_ProvablyRareGem.TransactOpts, kinds)
}

// AcceptManager is a paid mutator transaction binding the contract method 0xa41b41d2.
//
// Solidity: function acceptManager(uint256[] kinds) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) AcceptManager(kinds []*big.Int) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.AcceptManager(&_ProvablyRareGem.TransactOpts, kinds)
}

// Craft is a paid mutator transaction binding the contract method 0x8797b196.
//
// Solidity: function craft(uint256 kind, uint256 amount, address to) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactor) Craft(opts *bind.TransactOpts, kind *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "craft", kind, amount, to)
}

// Craft is a paid mutator transaction binding the contract method 0x8797b196.
//
// Solidity: function craft(uint256 kind, uint256 amount, address to) returns()
func (_ProvablyRareGem *ProvablyRareGemSession) Craft(kind *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.Craft(&_ProvablyRareGem.TransactOpts, kind, amount, to)
}

// Craft is a paid mutator transaction binding the contract method 0x8797b196.
//
// Solidity: function craft(uint256 kind, uint256 amount, address to) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) Craft(kind *big.Int, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.Craft(&_ProvablyRareGem.TransactOpts, kind, amount, to)
}

// Create is a paid mutator transaction binding the contract method 0x626bc15b.
//
// Solidity: function create(string name, string color, uint256 difficulty, uint256 gemsPerMine, uint256 multiplier, address crafter, address manager) returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemTransactor) Create(opts *bind.TransactOpts, name string, color string, difficulty *big.Int, gemsPerMine *big.Int, multiplier *big.Int, crafter common.Address, manager common.Address) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "create", name, color, difficulty, gemsPerMine, multiplier, crafter, manager)
}

// Create is a paid mutator transaction binding the contract method 0x626bc15b.
//
// Solidity: function create(string name, string color, uint256 difficulty, uint256 gemsPerMine, uint256 multiplier, address crafter, address manager) returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemSession) Create(name string, color string, difficulty *big.Int, gemsPerMine *big.Int, multiplier *big.Int, crafter common.Address, manager common.Address) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.Create(&_ProvablyRareGem.TransactOpts, name, color, difficulty, gemsPerMine, multiplier, crafter, manager)
}

// Create is a paid mutator transaction binding the contract method 0x626bc15b.
//
// Solidity: function create(string name, string color, uint256 difficulty, uint256 gemsPerMine, uint256 multiplier, address crafter, address manager) returns(uint256)
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) Create(name string, color string, difficulty *big.Int, gemsPerMine *big.Int, multiplier *big.Int, crafter common.Address, manager common.Address) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.Create(&_ProvablyRareGem.TransactOpts, name, color, difficulty, gemsPerMine, multiplier, crafter, manager)
}

// Mine is a paid mutator transaction binding the contract method 0x071e9503.
//
// Solidity: function mine(uint256 kind, uint256 salt) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactor) Mine(opts *bind.TransactOpts, kind *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "mine", kind, salt)
}

// Mine is a paid mutator transaction binding the contract method 0x071e9503.
//
// Solidity: function mine(uint256 kind, uint256 salt) returns()
func (_ProvablyRareGem *ProvablyRareGemSession) Mine(kind *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.Mine(&_ProvablyRareGem.TransactOpts, kind, salt)
}

// Mine is a paid mutator transaction binding the contract method 0x071e9503.
//
// Solidity: function mine(uint256 kind, uint256 salt) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) Mine(kind *big.Int, salt *big.Int) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.Mine(&_ProvablyRareGem.TransactOpts, kind, salt)
}

// RenounceManager is a paid mutator transaction binding the contract method 0x5d4b60aa.
//
// Solidity: function renounceManager(uint256[] kinds) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactor) RenounceManager(opts *bind.TransactOpts, kinds []*big.Int) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "renounceManager", kinds)
}

// RenounceManager is a paid mutator transaction binding the contract method 0x5d4b60aa.
//
// Solidity: function renounceManager(uint256[] kinds) returns()
func (_ProvablyRareGem *ProvablyRareGemSession) RenounceManager(kinds []*big.Int) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.RenounceManager(&_ProvablyRareGem.TransactOpts, kinds)
}

// RenounceManager is a paid mutator transaction binding the contract method 0x5d4b60aa.
//
// Solidity: function renounceManager(uint256[] kinds) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) RenounceManager(kinds []*big.Int) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.RenounceManager(&_ProvablyRareGem.TransactOpts, kinds)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ProvablyRareGem *ProvablyRareGemSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.SafeBatchTransferFrom(&_ProvablyRareGem.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.SafeBatchTransferFrom(&_ProvablyRareGem.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ProvablyRareGem *ProvablyRareGemSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.SafeTransferFrom(&_ProvablyRareGem.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.SafeTransferFrom(&_ProvablyRareGem.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ProvablyRareGem *ProvablyRareGemSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.SetApprovalForAll(&_ProvablyRareGem.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.SetApprovalForAll(&_ProvablyRareGem.TransactOpts, operator, approved)
}

// TransferManager is a paid mutator transaction binding the contract method 0xb88c1afb.
//
// Solidity: function transferManager(uint256[] kinds, address to) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactor) TransferManager(opts *bind.TransactOpts, kinds []*big.Int, to common.Address) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "transferManager", kinds, to)
}

// TransferManager is a paid mutator transaction binding the contract method 0xb88c1afb.
//
// Solidity: function transferManager(uint256[] kinds, address to) returns()
func (_ProvablyRareGem *ProvablyRareGemSession) TransferManager(kinds []*big.Int, to common.Address) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.TransferManager(&_ProvablyRareGem.TransactOpts, kinds, to)
}

// TransferManager is a paid mutator transaction binding the contract method 0xb88c1afb.
//
// Solidity: function transferManager(uint256[] kinds, address to) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) TransferManager(kinds []*big.Int, to common.Address) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.TransferManager(&_ProvablyRareGem.TransactOpts, kinds, to)
}

// UpdateCrafter is a paid mutator transaction binding the contract method 0xad1e8b37.
//
// Solidity: function updateCrafter(uint256[] kinds, address crafter) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactor) UpdateCrafter(opts *bind.TransactOpts, kinds []*big.Int, crafter common.Address) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "updateCrafter", kinds, crafter)
}

// UpdateCrafter is a paid mutator transaction binding the contract method 0xad1e8b37.
//
// Solidity: function updateCrafter(uint256[] kinds, address crafter) returns()
func (_ProvablyRareGem *ProvablyRareGemSession) UpdateCrafter(kinds []*big.Int, crafter common.Address) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.UpdateCrafter(&_ProvablyRareGem.TransactOpts, kinds, crafter)
}

// UpdateCrafter is a paid mutator transaction binding the contract method 0xad1e8b37.
//
// Solidity: function updateCrafter(uint256[] kinds, address crafter) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) UpdateCrafter(kinds []*big.Int, crafter common.Address) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.UpdateCrafter(&_ProvablyRareGem.TransactOpts, kinds, crafter)
}

// UpdateEntropy is a paid mutator transaction binding the contract method 0xb314782c.
//
// Solidity: function updateEntropy(uint256 kind, bytes32 entropy) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactor) UpdateEntropy(opts *bind.TransactOpts, kind *big.Int, entropy [32]byte) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "updateEntropy", kind, entropy)
}

// UpdateEntropy is a paid mutator transaction binding the contract method 0xb314782c.
//
// Solidity: function updateEntropy(uint256 kind, bytes32 entropy) returns()
func (_ProvablyRareGem *ProvablyRareGemSession) UpdateEntropy(kind *big.Int, entropy [32]byte) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.UpdateEntropy(&_ProvablyRareGem.TransactOpts, kind, entropy)
}

// UpdateEntropy is a paid mutator transaction binding the contract method 0xb314782c.
//
// Solidity: function updateEntropy(uint256 kind, bytes32 entropy) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) UpdateEntropy(kind *big.Int, entropy [32]byte) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.UpdateEntropy(&_ProvablyRareGem.TransactOpts, kind, entropy)
}

// UpdateGemInfo is a paid mutator transaction binding the contract method 0x089b1367.
//
// Solidity: function updateGemInfo(uint256 kind, string name, string color) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactor) UpdateGemInfo(opts *bind.TransactOpts, kind *big.Int, name string, color string) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "updateGemInfo", kind, name, color)
}

// UpdateGemInfo is a paid mutator transaction binding the contract method 0x089b1367.
//
// Solidity: function updateGemInfo(uint256 kind, string name, string color) returns()
func (_ProvablyRareGem *ProvablyRareGemSession) UpdateGemInfo(kind *big.Int, name string, color string) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.UpdateGemInfo(&_ProvablyRareGem.TransactOpts, kind, name, color)
}

// UpdateGemInfo is a paid mutator transaction binding the contract method 0x089b1367.
//
// Solidity: function updateGemInfo(uint256 kind, string name, string color) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) UpdateGemInfo(kind *big.Int, name string, color string) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.UpdateGemInfo(&_ProvablyRareGem.TransactOpts, kind, name, color)
}

// UpdateMiningData is a paid mutator transaction binding the contract method 0xbb03bfdd.
//
// Solidity: function updateMiningData(uint256 kind, uint256 difficulty, uint256 multiplier, uint256 gemsPerMine) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactor) UpdateMiningData(opts *bind.TransactOpts, kind *big.Int, difficulty *big.Int, multiplier *big.Int, gemsPerMine *big.Int) (*types.Transaction, error) {
	return _ProvablyRareGem.contract.Transact(opts, "updateMiningData", kind, difficulty, multiplier, gemsPerMine)
}

// UpdateMiningData is a paid mutator transaction binding the contract method 0xbb03bfdd.
//
// Solidity: function updateMiningData(uint256 kind, uint256 difficulty, uint256 multiplier, uint256 gemsPerMine) returns()
func (_ProvablyRareGem *ProvablyRareGemSession) UpdateMiningData(kind *big.Int, difficulty *big.Int, multiplier *big.Int, gemsPerMine *big.Int) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.UpdateMiningData(&_ProvablyRareGem.TransactOpts, kind, difficulty, multiplier, gemsPerMine)
}

// UpdateMiningData is a paid mutator transaction binding the contract method 0xbb03bfdd.
//
// Solidity: function updateMiningData(uint256 kind, uint256 difficulty, uint256 multiplier, uint256 gemsPerMine) returns()
func (_ProvablyRareGem *ProvablyRareGemTransactorSession) UpdateMiningData(kind *big.Int, difficulty *big.Int, multiplier *big.Int, gemsPerMine *big.Int) (*types.Transaction, error) {
	return _ProvablyRareGem.Contract.UpdateMiningData(&_ProvablyRareGem.TransactOpts, kind, difficulty, multiplier, gemsPerMine)
}

// ProvablyRareGemApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ProvablyRareGem contract.
type ProvablyRareGemApprovalForAllIterator struct {
	Event *ProvablyRareGemApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ProvablyRareGemApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvablyRareGemApprovalForAll)
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
		it.Event = new(ProvablyRareGemApprovalForAll)
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
func (it *ProvablyRareGemApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvablyRareGemApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvablyRareGemApprovalForAll represents a ApprovalForAll event raised by the ProvablyRareGem contract.
type ProvablyRareGemApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ProvablyRareGem *ProvablyRareGemFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*ProvablyRareGemApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ProvablyRareGem.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ProvablyRareGemApprovalForAllIterator{contract: _ProvablyRareGem.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ProvablyRareGem *ProvablyRareGemFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ProvablyRareGemApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ProvablyRareGem.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvablyRareGemApprovalForAll)
				if err := _ProvablyRareGem.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_ProvablyRareGem *ProvablyRareGemFilterer) ParseApprovalForAll(log types.Log) (*ProvablyRareGemApprovalForAll, error) {
	event := new(ProvablyRareGemApprovalForAll)
	if err := _ProvablyRareGem.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvablyRareGemCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the ProvablyRareGem contract.
type ProvablyRareGemCreateIterator struct {
	Event *ProvablyRareGemCreate // Event containing the contract specifics and raw log

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
func (it *ProvablyRareGemCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvablyRareGemCreate)
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
		it.Event = new(ProvablyRareGemCreate)
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
func (it *ProvablyRareGemCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvablyRareGemCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvablyRareGemCreate represents a Create event raised by the ProvablyRareGem contract.
type ProvablyRareGemCreate struct {
	Kind *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0x07eac9a0695a188fe9d6fd680bcbbbe39041fb114d5d7ac11252401391f79307.
//
// Solidity: event Create(uint256 indexed kind)
func (_ProvablyRareGem *ProvablyRareGemFilterer) FilterCreate(opts *bind.FilterOpts, kind []*big.Int) (*ProvablyRareGemCreateIterator, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}

	logs, sub, err := _ProvablyRareGem.contract.FilterLogs(opts, "Create", kindRule)
	if err != nil {
		return nil, err
	}
	return &ProvablyRareGemCreateIterator{contract: _ProvablyRareGem.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0x07eac9a0695a188fe9d6fd680bcbbbe39041fb114d5d7ac11252401391f79307.
//
// Solidity: event Create(uint256 indexed kind)
func (_ProvablyRareGem *ProvablyRareGemFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *ProvablyRareGemCreate, kind []*big.Int) (event.Subscription, error) {

	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}

	logs, sub, err := _ProvablyRareGem.contract.WatchLogs(opts, "Create", kindRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvablyRareGemCreate)
				if err := _ProvablyRareGem.contract.UnpackLog(event, "Create", log); err != nil {
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

// ParseCreate is a log parse operation binding the contract event 0x07eac9a0695a188fe9d6fd680bcbbbe39041fb114d5d7ac11252401391f79307.
//
// Solidity: event Create(uint256 indexed kind)
func (_ProvablyRareGem *ProvablyRareGemFilterer) ParseCreate(log types.Log) (*ProvablyRareGemCreate, error) {
	event := new(ProvablyRareGemCreate)
	if err := _ProvablyRareGem.contract.UnpackLog(event, "Create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvablyRareGemMineIterator is returned from FilterMine and is used to iterate over the raw logs and unpacked data for Mine events raised by the ProvablyRareGem contract.
type ProvablyRareGemMineIterator struct {
	Event *ProvablyRareGemMine // Event containing the contract specifics and raw log

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
func (it *ProvablyRareGemMineIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvablyRareGemMine)
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
		it.Event = new(ProvablyRareGemMine)
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
func (it *ProvablyRareGemMineIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvablyRareGemMineIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvablyRareGemMine represents a Mine event raised by the ProvablyRareGem contract.
type ProvablyRareGemMine struct {
	Miner common.Address
	Kind  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMine is a free log retrieval operation binding the contract event 0xf23a961744a760027f8811c59a0eaef0d29cf965578b17412bcc375b52fa39d1.
//
// Solidity: event Mine(address indexed miner, uint256 indexed kind)
func (_ProvablyRareGem *ProvablyRareGemFilterer) FilterMine(opts *bind.FilterOpts, miner []common.Address, kind []*big.Int) (*ProvablyRareGemMineIterator, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}

	logs, sub, err := _ProvablyRareGem.contract.FilterLogs(opts, "Mine", minerRule, kindRule)
	if err != nil {
		return nil, err
	}
	return &ProvablyRareGemMineIterator{contract: _ProvablyRareGem.contract, event: "Mine", logs: logs, sub: sub}, nil
}

// WatchMine is a free log subscription operation binding the contract event 0xf23a961744a760027f8811c59a0eaef0d29cf965578b17412bcc375b52fa39d1.
//
// Solidity: event Mine(address indexed miner, uint256 indexed kind)
func (_ProvablyRareGem *ProvablyRareGemFilterer) WatchMine(opts *bind.WatchOpts, sink chan<- *ProvablyRareGemMine, miner []common.Address, kind []*big.Int) (event.Subscription, error) {

	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}

	logs, sub, err := _ProvablyRareGem.contract.WatchLogs(opts, "Mine", minerRule, kindRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvablyRareGemMine)
				if err := _ProvablyRareGem.contract.UnpackLog(event, "Mine", log); err != nil {
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

// ParseMine is a log parse operation binding the contract event 0xf23a961744a760027f8811c59a0eaef0d29cf965578b17412bcc375b52fa39d1.
//
// Solidity: event Mine(address indexed miner, uint256 indexed kind)
func (_ProvablyRareGem *ProvablyRareGemFilterer) ParseMine(log types.Log) (*ProvablyRareGemMine, error) {
	event := new(ProvablyRareGemMine)
	if err := _ProvablyRareGem.contract.UnpackLog(event, "Mine", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvablyRareGemTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ProvablyRareGem contract.
type ProvablyRareGemTransferBatchIterator struct {
	Event *ProvablyRareGemTransferBatch // Event containing the contract specifics and raw log

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
func (it *ProvablyRareGemTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvablyRareGemTransferBatch)
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
		it.Event = new(ProvablyRareGemTransferBatch)
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
func (it *ProvablyRareGemTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvablyRareGemTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvablyRareGemTransferBatch represents a TransferBatch event raised by the ProvablyRareGem contract.
type ProvablyRareGemTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ProvablyRareGem *ProvablyRareGemFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ProvablyRareGemTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ProvablyRareGem.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ProvablyRareGemTransferBatchIterator{contract: _ProvablyRareGem.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ProvablyRareGem *ProvablyRareGemFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ProvablyRareGemTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ProvablyRareGem.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvablyRareGemTransferBatch)
				if err := _ProvablyRareGem.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ProvablyRareGem *ProvablyRareGemFilterer) ParseTransferBatch(log types.Log) (*ProvablyRareGemTransferBatch, error) {
	event := new(ProvablyRareGemTransferBatch)
	if err := _ProvablyRareGem.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvablyRareGemTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ProvablyRareGem contract.
type ProvablyRareGemTransferSingleIterator struct {
	Event *ProvablyRareGemTransferSingle // Event containing the contract specifics and raw log

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
func (it *ProvablyRareGemTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvablyRareGemTransferSingle)
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
		it.Event = new(ProvablyRareGemTransferSingle)
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
func (it *ProvablyRareGemTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvablyRareGemTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvablyRareGemTransferSingle represents a TransferSingle event raised by the ProvablyRareGem contract.
type ProvablyRareGemTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ProvablyRareGem *ProvablyRareGemFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ProvablyRareGemTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ProvablyRareGem.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ProvablyRareGemTransferSingleIterator{contract: _ProvablyRareGem.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ProvablyRareGem *ProvablyRareGemFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ProvablyRareGemTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ProvablyRareGem.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvablyRareGemTransferSingle)
				if err := _ProvablyRareGem.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ProvablyRareGem *ProvablyRareGemFilterer) ParseTransferSingle(log types.Log) (*ProvablyRareGemTransferSingle, error) {
	event := new(ProvablyRareGemTransferSingle)
	if err := _ProvablyRareGem.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProvablyRareGemURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ProvablyRareGem contract.
type ProvablyRareGemURIIterator struct {
	Event *ProvablyRareGemURI // Event containing the contract specifics and raw log

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
func (it *ProvablyRareGemURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProvablyRareGemURI)
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
		it.Event = new(ProvablyRareGemURI)
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
func (it *ProvablyRareGemURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProvablyRareGemURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProvablyRareGemURI represents a URI event raised by the ProvablyRareGem contract.
type ProvablyRareGemURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ProvablyRareGem *ProvablyRareGemFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ProvablyRareGemURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ProvablyRareGem.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ProvablyRareGemURIIterator{contract: _ProvablyRareGem.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ProvablyRareGem *ProvablyRareGemFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ProvablyRareGemURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ProvablyRareGem.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProvablyRareGemURI)
				if err := _ProvablyRareGem.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ProvablyRareGem *ProvablyRareGemFilterer) ParseURI(log types.Log) (*ProvablyRareGemURI, error) {
	event := new(ProvablyRareGemURI)
	if err := _ProvablyRareGem.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
