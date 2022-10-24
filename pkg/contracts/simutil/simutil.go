// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simutil

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

// SimUtilABI is the input ABI used to generate the binding from.
const SimUtilABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SimUtilBin is the compiled bytecode used for deploying new contracts.
var SimUtilBin = "0x608060405234801561001057600080fd5b50610407806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063ef5bfc3714610030575b600080fd5b6100fc6004803603604081101561004657600080fd5b810190808035906020019064010000000081111561006357600080fd5b82018360208201111561007557600080fd5b8035906020019184602083028401116401000000008311171561009757600080fd5b9091929391929390803590602001906401000000008111156100b857600080fd5b8201836020820111156100ca57600080fd5b803590602001918460208302840111640100000000831117156100ec57600080fd5b9091929391929390505050610153565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561013f578082015181840152602081019050610124565b505050509050019250505060405180910390f35b60608083839050868690500267ffffffffffffffff8111801561017557600080fd5b506040519080825280602002602001820160405280156101a45781602001602082028036833780820191505090505b50905060005b868690508110156103c45760005b858590508110156103b65773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168686838181106101fa57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156102975787878381811061023f57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16318382888890508502018151811061028657fe5b6020026020010181815250506103a9565b8585828181106102a357fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a082318989858181106102e757fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561034e57600080fd5b505afa158015610362573d6000803e3d6000fd5b505050506040513d602081101561037857600080fd5b81019080805190602001909291905050508382888890508502018151811061039c57fe5b6020026020010181815250505b80806001019150506101b8565b5080806001019150506101aa565b508091505094935050505056fea26469706673582212200b8738e9ce84a5af761886e464d14ce62b0d8bce0316bc15a43ba74267a846ea64736f6c634300060c0033"

// DeploySimUtil deploys a new Ethereum contract, binding an instance of SimUtil to it.
func DeploySimUtil(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SimUtil, error) {
	parsed, err := abi.JSON(strings.NewReader(SimUtilABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SimUtilBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimUtil{SimUtilCaller: SimUtilCaller{contract: contract}, SimUtilTransactor: SimUtilTransactor{contract: contract}, SimUtilFilterer: SimUtilFilterer{contract: contract}}, nil
}

// SimUtil is an auto generated Go binding around an Ethereum contract.
type SimUtil struct {
	SimUtilCaller     // Read-only binding to the contract
	SimUtilTransactor // Write-only binding to the contract
	SimUtilFilterer   // Log filterer for contract events
}

// SimUtilCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimUtilCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimUtilTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimUtilTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimUtilFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimUtilFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimUtilSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimUtilSession struct {
	Contract     *SimUtil          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimUtilCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimUtilCallerSession struct {
	Contract *SimUtilCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SimUtilTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimUtilTransactorSession struct {
	Contract     *SimUtilTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SimUtilRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimUtilRaw struct {
	Contract *SimUtil // Generic contract binding to access the raw methods on
}

// SimUtilCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimUtilCallerRaw struct {
	Contract *SimUtilCaller // Generic read-only contract binding to access the raw methods on
}

// SimUtilTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimUtilTransactorRaw struct {
	Contract *SimUtilTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimUtil creates a new instance of SimUtil, bound to a specific deployed contract.
func NewSimUtil(address common.Address, backend bind.ContractBackend) (*SimUtil, error) {
	contract, err := bindSimUtil(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimUtil{SimUtilCaller: SimUtilCaller{contract: contract}, SimUtilTransactor: SimUtilTransactor{contract: contract}, SimUtilFilterer: SimUtilFilterer{contract: contract}}, nil
}

// NewSimUtilCaller creates a new read-only instance of SimUtil, bound to a specific deployed contract.
func NewSimUtilCaller(address common.Address, caller bind.ContractCaller) (*SimUtilCaller, error) {
	contract, err := bindSimUtil(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimUtilCaller{contract: contract}, nil
}

// NewSimUtilTransactor creates a new write-only instance of SimUtil, bound to a specific deployed contract.
func NewSimUtilTransactor(address common.Address, transactor bind.ContractTransactor) (*SimUtilTransactor, error) {
	contract, err := bindSimUtil(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimUtilTransactor{contract: contract}, nil
}

// NewSimUtilFilterer creates a new log filterer instance of SimUtil, bound to a specific deployed contract.
func NewSimUtilFilterer(address common.Address, filterer bind.ContractFilterer) (*SimUtilFilterer, error) {
	contract, err := bindSimUtil(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimUtilFilterer{contract: contract}, nil
}

// bindSimUtil binds a generic wrapper to an already deployed contract.
func bindSimUtil(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimUtilABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimUtil *SimUtilRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimUtil.Contract.SimUtilCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimUtil *SimUtilRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimUtil.Contract.SimUtilTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimUtil *SimUtilRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimUtil.Contract.SimUtilTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimUtil *SimUtilCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimUtil.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimUtil *SimUtilTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimUtil.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimUtil *SimUtilTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimUtil.Contract.contract.Transact(opts, method, params...)
}

// GetBalances is a free data retrieval call binding the contract method 0xef5bfc37.
//
// Solidity: function getBalances(address[] users, address[] tokens) view returns(uint256[])
func (_SimUtil *SimUtilCaller) GetBalances(opts *bind.CallOpts, users []common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _SimUtil.contract.Call(opts, &out, "getBalances", users, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBalances is a free data retrieval call binding the contract method 0xef5bfc37.
//
// Solidity: function getBalances(address[] users, address[] tokens) view returns(uint256[])
func (_SimUtil *SimUtilSession) GetBalances(users []common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _SimUtil.Contract.GetBalances(&_SimUtil.CallOpts, users, tokens)
}

// GetBalances is a free data retrieval call binding the contract method 0xef5bfc37.
//
// Solidity: function getBalances(address[] users, address[] tokens) view returns(uint256[])
func (_SimUtil *SimUtilCallerSession) GetBalances(users []common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _SimUtil.Contract.GetBalances(&_SimUtil.CallOpts, users, tokens)
}
