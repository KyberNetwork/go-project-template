// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multiutil

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MultiUtilMetaData contains all meta data concerning the MultiUtil contract.
var MultiUtilMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"balanceOfMultiTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOfMultiUsers\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610a11806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80635133ac4a14610046578063546028ce14610134578063ef5bfc3714610222575b600080fd5b6100dd6004803603604081101561005c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561009957600080fd5b8201836020820111156100ab57600080fd5b803590602001918460208302840111640100000000831117156100cd57600080fd5b9091929391929390505050610345565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610120578082015181840152602081019050610105565b505050509050019250505060405180910390f35b6101cb6004803603604081101561014a57600080fd5b810190808035906020019064010000000081111561016757600080fd5b82018360208201111561017957600080fd5b8035906020019184602083028401116401000000008311171561019b57600080fd5b9091929391929390803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610543565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561020e5780820151818401526020810190506101f3565b505050509050019250505060405180910390f35b6102ee6004803603604081101561023857600080fd5b810190808035906020019064010000000081111561025557600080fd5b82018360208201111561026757600080fd5b8035906020019184602083028401116401000000008311171561028957600080fd5b9091929391929390803590602001906401000000008111156102aa57600080fd5b8201836020820111156102bc57600080fd5b803590602001918460208302840111640100000000831117156102de57600080fd5b909192939192939050505061075d565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610331578082015181840152602081019050610316565b505050509050019250505060405180910390f35b6060808383905067ffffffffffffffff8111801561036257600080fd5b506040519080825280602002602001820160405280156103915781602001602082028036833780820191505090505b50905060005b848490508110156105375773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168585838181106103d957fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610447578573ffffffffffffffffffffffffffffffffffffffff163182828151811061043657fe5b60200260200101818152505061052a565b84848281811061045357fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231876040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156104d657600080fd5b505afa1580156104ea573d6000803e3d6000fd5b505050506040513d602081101561050057600080fd5b810190808051906020019092919050505082828151811061051d57fe5b6020026020010181815250505b8080600101915050610397565b50809150509392505050565b6060808484905067ffffffffffffffff8111801561056057600080fd5b5060405190808252806020026020018201604052801561058f5781602001602082028036833780820191505090505b50905073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156106535760005b8585905081101561064d578585828181106105f457fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163182828151811061063457fe5b60200260200101818152505080806001019150506105dd565b50610752565b60005b85859050811015610750578373ffffffffffffffffffffffffffffffffffffffff166370a0823187878481811061068957fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156106f057600080fd5b505afa158015610704573d6000803e3d6000fd5b505050506040513d602081101561071a57600080fd5b810190808051906020019092919050505082828151811061073757fe5b6020026020010181815250508080600101915050610656565b505b809150509392505050565b60608083839050868690500267ffffffffffffffff8111801561077f57600080fd5b506040519080825280602002602001820160405280156107ae5781602001602082028036833780820191505090505b50905060005b868690508110156109ce5760005b858590508110156109c05773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff1686868381811061080457fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156108a15787878381811061084957fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16318382888890508502018151811061089057fe5b6020026020010181815250506109b3565b8585828181106108ad57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a082318989858181106108f157fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561095857600080fd5b505afa15801561096c573d6000803e3d6000fd5b505050506040513d602081101561098257600080fd5b8101908080519060200190929190505050838288889050850201815181106109a657fe5b6020026020010181815250505b80806001019150506107c2565b5080806001019150506107b4565b508091505094935050505056fea26469706673582212209cc97551e5722c2bc04e6912996846d81fc36cbec8f29662efbb16d15e35bb0c64736f6c634300060c0033",
}

// MultiUtilABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiUtilMetaData.ABI instead.
var MultiUtilABI = MultiUtilMetaData.ABI

// MultiUtilBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultiUtilMetaData.Bin instead.
var MultiUtilBin = MultiUtilMetaData.Bin

// DeployMultiUtil deploys a new Ethereum contract, binding an instance of MultiUtil to it.
func DeployMultiUtil(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MultiUtil, error) {
	parsed, err := MultiUtilMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultiUtilBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MultiUtil{MultiUtilCaller: MultiUtilCaller{contract: contract}, MultiUtilTransactor: MultiUtilTransactor{contract: contract}, MultiUtilFilterer: MultiUtilFilterer{contract: contract}}, nil
}

// MultiUtil is an auto generated Go binding around an Ethereum contract.
type MultiUtil struct {
	MultiUtilCaller     // Read-only binding to the contract
	MultiUtilTransactor // Write-only binding to the contract
	MultiUtilFilterer   // Log filterer for contract events
}

// MultiUtilCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiUtilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiUtilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiUtilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiUtilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiUtilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiUtilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiUtilSession struct {
	Contract     *MultiUtil        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiUtilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiUtilCallerSession struct {
	Contract *MultiUtilCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MultiUtilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiUtilTransactorSession struct {
	Contract     *MultiUtilTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MultiUtilRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiUtilRaw struct {
	Contract *MultiUtil // Generic contract binding to access the raw methods on
}

// MultiUtilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiUtilCallerRaw struct {
	Contract *MultiUtilCaller // Generic read-only contract binding to access the raw methods on
}

// MultiUtilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiUtilTransactorRaw struct {
	Contract *MultiUtilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiUtil creates a new instance of MultiUtil, bound to a specific deployed contract.
func NewMultiUtil(address common.Address, backend bind.ContractBackend) (*MultiUtil, error) {
	contract, err := bindMultiUtil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiUtil{MultiUtilCaller: MultiUtilCaller{contract: contract}, MultiUtilTransactor: MultiUtilTransactor{contract: contract}, MultiUtilFilterer: MultiUtilFilterer{contract: contract}}, nil
}

// NewMultiUtilCaller creates a new read-only instance of MultiUtil, bound to a specific deployed contract.
func NewMultiUtilCaller(address common.Address, caller bind.ContractCaller) (*MultiUtilCaller, error) {
	contract, err := bindMultiUtil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiUtilCaller{contract: contract}, nil
}

// NewMultiUtilTransactor creates a new write-only instance of MultiUtil, bound to a specific deployed contract.
func NewMultiUtilTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiUtilTransactor, error) {
	contract, err := bindMultiUtil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiUtilTransactor{contract: contract}, nil
}

// NewMultiUtilFilterer creates a new log filterer instance of MultiUtil, bound to a specific deployed contract.
func NewMultiUtilFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiUtilFilterer, error) {
	contract, err := bindMultiUtil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiUtilFilterer{contract: contract}, nil
}

// bindMultiUtil binds a generic wrapper to an already deployed contract.
func bindMultiUtil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiUtilABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiUtil *MultiUtilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiUtil.Contract.MultiUtilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiUtil *MultiUtilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiUtil.Contract.MultiUtilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiUtil *MultiUtilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiUtil.Contract.MultiUtilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiUtil *MultiUtilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiUtil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiUtil *MultiUtilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiUtil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiUtil *MultiUtilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiUtil.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfMultiTokens is a free data retrieval call binding the contract method 0x5133ac4a.
//
// Solidity: function balanceOfMultiTokens(address user, address[] tokens) view returns(uint256[])
func (_MultiUtil *MultiUtilCaller) BalanceOfMultiTokens(opts *bind.CallOpts, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _MultiUtil.contract.Call(opts, &out, "balanceOfMultiTokens", user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfMultiTokens is a free data retrieval call binding the contract method 0x5133ac4a.
//
// Solidity: function balanceOfMultiTokens(address user, address[] tokens) view returns(uint256[])
func (_MultiUtil *MultiUtilSession) BalanceOfMultiTokens(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _MultiUtil.Contract.BalanceOfMultiTokens(&_MultiUtil.CallOpts, user, tokens)
}

// BalanceOfMultiTokens is a free data retrieval call binding the contract method 0x5133ac4a.
//
// Solidity: function balanceOfMultiTokens(address user, address[] tokens) view returns(uint256[])
func (_MultiUtil *MultiUtilCallerSession) BalanceOfMultiTokens(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _MultiUtil.Contract.BalanceOfMultiTokens(&_MultiUtil.CallOpts, user, tokens)
}

// BalanceOfMultiUsers is a free data retrieval call binding the contract method 0x546028ce.
//
// Solidity: function balanceOfMultiUsers(address[] users, address token) view returns(uint256[])
func (_MultiUtil *MultiUtilCaller) BalanceOfMultiUsers(opts *bind.CallOpts, users []common.Address, token common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _MultiUtil.contract.Call(opts, &out, "balanceOfMultiUsers", users, token)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfMultiUsers is a free data retrieval call binding the contract method 0x546028ce.
//
// Solidity: function balanceOfMultiUsers(address[] users, address token) view returns(uint256[])
func (_MultiUtil *MultiUtilSession) BalanceOfMultiUsers(users []common.Address, token common.Address) ([]*big.Int, error) {
	return _MultiUtil.Contract.BalanceOfMultiUsers(&_MultiUtil.CallOpts, users, token)
}

// BalanceOfMultiUsers is a free data retrieval call binding the contract method 0x546028ce.
//
// Solidity: function balanceOfMultiUsers(address[] users, address token) view returns(uint256[])
func (_MultiUtil *MultiUtilCallerSession) BalanceOfMultiUsers(users []common.Address, token common.Address) ([]*big.Int, error) {
	return _MultiUtil.Contract.BalanceOfMultiUsers(&_MultiUtil.CallOpts, users, token)
}

// GetBalances is a free data retrieval call binding the contract method 0xef5bfc37.
//
// Solidity: function getBalances(address[] users, address[] tokens) view returns(uint256[])
func (_MultiUtil *MultiUtilCaller) GetBalances(opts *bind.CallOpts, users []common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _MultiUtil.contract.Call(opts, &out, "getBalances", users, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0xef5bfc37.
//
// Solidity: function getBalances(address[] users, address[] tokens) view returns(uint256[])
func (_MultiUtil *MultiUtilSession) GetBalances(users []common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _MultiUtil.Contract.GetBalances(&_MultiUtil.CallOpts, users, tokens)
}

// GetBalances is a free data retrieval call binding the contract method 0xef5bfc37.
//
// Solidity: function getBalances(address[] users, address[] tokens) view returns(uint256[])
func (_MultiUtil *MultiUtilCallerSession) GetBalances(users []common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _MultiUtil.Contract.GetBalances(&_MultiUtil.CallOpts, users, tokens)
}
