// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// FantomAddressRegistryMetaData contains all meta data concerning the FantomAddressRegistry contract.
var FantomAddressRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"artFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"artion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bundleMarketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"privateArtFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"privateFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_artFactory\",\"type\":\"address\"}],\"name\":\"updateArtFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_privateArtFactory\",\"type\":\"address\"}],\"name\":\"updateArtFactoryPrivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_artion\",\"type\":\"address\"}],\"name\":\"updateArtion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auction\",\"type\":\"address\"}],\"name\":\"updateAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bundleMarketplace\",\"type\":\"address\"}],\"name\":\"updateBundleMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"}],\"name\":\"updateMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"updateNFTFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_privateFactory\",\"type\":\"address\"}],\"name\":\"updateNFTFactoryPrivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"updatePriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenRegistry\",\"type\":\"address\"}],\"name\":\"updateTokenRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0fb27d8c": "artFactory()",
		"18944e55": "artion()",
		"7d9f6db5": "auction()",
		"9e1f5b28": "bundleMarketplace()",
		"c45a0155": "factory()",
		"abc8c7af": "marketplace()",
		"8da5cb5b": "owner()",
		"741bef1a": "priceFeed()",
		"924db13f": "privateArtFactory()",
		"735db939": "privateFactory()",
		"715018a6": "renounceOwnership()",
		"9d23c4c7": "tokenRegistry()",
		"f2fde38b": "transferOwnership(address)",
		"8590f295": "updateArtFactory(address)",
		"b717bb22": "updateArtFactoryPrivate(address)",
		"ce35688a": "updateArtion(address)",
		"3402a036": "updateAuction(address)",
		"9d0e24a7": "updateBundleMarketplace(address)",
		"eb636317": "updateMarketplace(address)",
		"975a426a": "updateNFTFactory(address)",
		"e35bead5": "updateNFTFactoryPrivate(address)",
		"95877f78": "updatePriceFeed(address)",
		"1fa915a4": "updateTokenRegistry(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61089f8061007e6000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806395877f78116100c3578063b717bb221161007c578063b717bb22146102a6578063c45a0155146102b9578063ce35688a146102cc578063e35bead5146102df578063eb636317146102f2578063f2fde38b1461030557600080fd5b806395877f7814610234578063975a426a146102475780639d0e24a71461025a5780639d23c4c71461026d5780639e1f5b2814610280578063abc8c7af1461029357600080fd5b8063735db93911610115578063735db939146101c4578063741bef1a146101d75780637d9f6db5146101ea5780638590f295146101fd5780638da5cb5b14610210578063924db13f1461022157600080fd5b80630fb27d8c1461015257806318944e55146101815780631fa915a4146101945780633402a036146101a9578063715018a6146101bc575b600080fd5b600754610165906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f35b600154610165906001600160a01b031681565b6101a76101a23660046107e2565b610318565b005b6101a76101b73660046107e2565b61036d565b6101a76103b9565b600654610165906001600160a01b031681565b600a54610165906001600160a01b031681565b600254610165906001600160a01b031681565b6101a761020b3660046107e2565b6103ef565b6000546001600160a01b0316610165565b600854610165906001600160a01b031681565b6101a76102423660046107e2565b61043b565b6101a76102553660046107e2565b610487565b6101a76102683660046107e2565b6104d3565b600954610165906001600160a01b031681565b600454610165906001600160a01b031681565b600354610165906001600160a01b031681565b6101a76102b43660046107e2565b61051f565b600554610165906001600160a01b031681565b6101a76102da3660046107e2565b61056b565b6101a76102ed3660046107e2565b61065f565b6101a76103003660046107e2565b6106ab565b6101a76103133660046107e2565b6106f7565b6000546001600160a01b0316331461034b5760405162461bcd60e51b815260040161034290610812565b60405180910390fd5b600980546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146103975760405162461bcd60e51b815260040161034290610812565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146103e35760405162461bcd60e51b815260040161034290610812565b6103ed6000610792565b565b6000546001600160a01b031633146104195760405162461bcd60e51b815260040161034290610812565b600780546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146104655760405162461bcd60e51b815260040161034290610812565b600a80546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146104b15760405162461bcd60e51b815260040161034290610812565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146104fd5760405162461bcd60e51b815260040161034290610812565b600480546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146105495760405162461bcd60e51b815260040161034290610812565b600880546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146105955760405162461bcd60e51b815260040161034290610812565b6040516301ffc9a760e01b81526380ac58cd60e01b60048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa1580156105e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106049190610847565b61063d5760405162461bcd60e51b815260206004820152600a6024820152694e6f742045524337323160b01b6044820152606401610342565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146106895760405162461bcd60e51b815260040161034290610812565b600680546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146106d55760405162461bcd60e51b815260040161034290610812565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146107215760405162461bcd60e51b815260040161034290610812565b6001600160a01b0381166107865760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610342565b61078f81610792565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156107f457600080fd5b81356001600160a01b038116811461080b57600080fd5b9392505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60006020828403121561085957600080fd5b8151801515811461080b57600080fdfea2646970667358221220877477a2c591ad906446846d48fc2c825d158acdaba8ba6637f23c7808f7463f64736f6c634300080b0033",
}

// FantomAddressRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomAddressRegistryMetaData.ABI instead.
var FantomAddressRegistryABI = FantomAddressRegistryMetaData.ABI

// Deprecated: Use FantomAddressRegistryMetaData.Sigs instead.
// FantomAddressRegistryFuncSigs maps the 4-byte function signature to its string representation.
var FantomAddressRegistryFuncSigs = FantomAddressRegistryMetaData.Sigs

// FantomAddressRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FantomAddressRegistryMetaData.Bin instead.
var FantomAddressRegistryBin = FantomAddressRegistryMetaData.Bin

// DeployFantomAddressRegistry deploys a new Ethereum contract, binding an instance of FantomAddressRegistry to it.
func DeployFantomAddressRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FantomAddressRegistry, error) {
	parsed, err := FantomAddressRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FantomAddressRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FantomAddressRegistry{FantomAddressRegistryCaller: FantomAddressRegistryCaller{contract: contract}, FantomAddressRegistryTransactor: FantomAddressRegistryTransactor{contract: contract}, FantomAddressRegistryFilterer: FantomAddressRegistryFilterer{contract: contract}}, nil
}

// FantomAddressRegistry is an auto generated Go binding around an Ethereum contract.
type FantomAddressRegistry struct {
	FantomAddressRegistryCaller     // Read-only binding to the contract
	FantomAddressRegistryTransactor // Write-only binding to the contract
	FantomAddressRegistryFilterer   // Log filterer for contract events
}

// FantomAddressRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FantomAddressRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomAddressRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomAddressRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomAddressRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomAddressRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomAddressRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomAddressRegistrySession struct {
	Contract     *FantomAddressRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FantomAddressRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomAddressRegistryCallerSession struct {
	Contract *FantomAddressRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// FantomAddressRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomAddressRegistryTransactorSession struct {
	Contract     *FantomAddressRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// FantomAddressRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FantomAddressRegistryRaw struct {
	Contract *FantomAddressRegistry // Generic contract binding to access the raw methods on
}

// FantomAddressRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomAddressRegistryCallerRaw struct {
	Contract *FantomAddressRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// FantomAddressRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomAddressRegistryTransactorRaw struct {
	Contract *FantomAddressRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomAddressRegistry creates a new instance of FantomAddressRegistry, bound to a specific deployed contract.
func NewFantomAddressRegistry(address common.Address, backend bind.ContractBackend) (*FantomAddressRegistry, error) {
	contract, err := bindFantomAddressRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomAddressRegistry{FantomAddressRegistryCaller: FantomAddressRegistryCaller{contract: contract}, FantomAddressRegistryTransactor: FantomAddressRegistryTransactor{contract: contract}, FantomAddressRegistryFilterer: FantomAddressRegistryFilterer{contract: contract}}, nil
}

// NewFantomAddressRegistryCaller creates a new read-only instance of FantomAddressRegistry, bound to a specific deployed contract.
func NewFantomAddressRegistryCaller(address common.Address, caller bind.ContractCaller) (*FantomAddressRegistryCaller, error) {
	contract, err := bindFantomAddressRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomAddressRegistryCaller{contract: contract}, nil
}

// NewFantomAddressRegistryTransactor creates a new write-only instance of FantomAddressRegistry, bound to a specific deployed contract.
func NewFantomAddressRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*FantomAddressRegistryTransactor, error) {
	contract, err := bindFantomAddressRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomAddressRegistryTransactor{contract: contract}, nil
}

// NewFantomAddressRegistryFilterer creates a new log filterer instance of FantomAddressRegistry, bound to a specific deployed contract.
func NewFantomAddressRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*FantomAddressRegistryFilterer, error) {
	contract, err := bindFantomAddressRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomAddressRegistryFilterer{contract: contract}, nil
}

// bindFantomAddressRegistry binds a generic wrapper to an already deployed contract.
func bindFantomAddressRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomAddressRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomAddressRegistry *FantomAddressRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomAddressRegistry.Contract.FantomAddressRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomAddressRegistry *FantomAddressRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.FantomAddressRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomAddressRegistry *FantomAddressRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.FantomAddressRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomAddressRegistry *FantomAddressRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomAddressRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomAddressRegistry *FantomAddressRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomAddressRegistry *FantomAddressRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.contract.Transact(opts, method, params...)
}

// ArtFactory is a free data retrieval call binding the contract method 0x0fb27d8c.
//
// Solidity: function artFactory() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCaller) ArtFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAddressRegistry.contract.Call(opts, &out, "artFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ArtFactory is a free data retrieval call binding the contract method 0x0fb27d8c.
//
// Solidity: function artFactory() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistrySession) ArtFactory() (common.Address, error) {
	return _FantomAddressRegistry.Contract.ArtFactory(&_FantomAddressRegistry.CallOpts)
}

// ArtFactory is a free data retrieval call binding the contract method 0x0fb27d8c.
//
// Solidity: function artFactory() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCallerSession) ArtFactory() (common.Address, error) {
	return _FantomAddressRegistry.Contract.ArtFactory(&_FantomAddressRegistry.CallOpts)
}

// Artion is a free data retrieval call binding the contract method 0x18944e55.
//
// Solidity: function artion() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCaller) Artion(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAddressRegistry.contract.Call(opts, &out, "artion")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Artion is a free data retrieval call binding the contract method 0x18944e55.
//
// Solidity: function artion() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistrySession) Artion() (common.Address, error) {
	return _FantomAddressRegistry.Contract.Artion(&_FantomAddressRegistry.CallOpts)
}

// Artion is a free data retrieval call binding the contract method 0x18944e55.
//
// Solidity: function artion() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCallerSession) Artion() (common.Address, error) {
	return _FantomAddressRegistry.Contract.Artion(&_FantomAddressRegistry.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCaller) Auction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAddressRegistry.contract.Call(opts, &out, "auction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistrySession) Auction() (common.Address, error) {
	return _FantomAddressRegistry.Contract.Auction(&_FantomAddressRegistry.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCallerSession) Auction() (common.Address, error) {
	return _FantomAddressRegistry.Contract.Auction(&_FantomAddressRegistry.CallOpts)
}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCaller) BundleMarketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAddressRegistry.contract.Call(opts, &out, "bundleMarketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistrySession) BundleMarketplace() (common.Address, error) {
	return _FantomAddressRegistry.Contract.BundleMarketplace(&_FantomAddressRegistry.CallOpts)
}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCallerSession) BundleMarketplace() (common.Address, error) {
	return _FantomAddressRegistry.Contract.BundleMarketplace(&_FantomAddressRegistry.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAddressRegistry.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistrySession) Factory() (common.Address, error) {
	return _FantomAddressRegistry.Contract.Factory(&_FantomAddressRegistry.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCallerSession) Factory() (common.Address, error) {
	return _FantomAddressRegistry.Contract.Factory(&_FantomAddressRegistry.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCaller) Marketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAddressRegistry.contract.Call(opts, &out, "marketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistrySession) Marketplace() (common.Address, error) {
	return _FantomAddressRegistry.Contract.Marketplace(&_FantomAddressRegistry.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCallerSession) Marketplace() (common.Address, error) {
	return _FantomAddressRegistry.Contract.Marketplace(&_FantomAddressRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAddressRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistrySession) Owner() (common.Address, error) {
	return _FantomAddressRegistry.Contract.Owner(&_FantomAddressRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCallerSession) Owner() (common.Address, error) {
	return _FantomAddressRegistry.Contract.Owner(&_FantomAddressRegistry.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAddressRegistry.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistrySession) PriceFeed() (common.Address, error) {
	return _FantomAddressRegistry.Contract.PriceFeed(&_FantomAddressRegistry.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCallerSession) PriceFeed() (common.Address, error) {
	return _FantomAddressRegistry.Contract.PriceFeed(&_FantomAddressRegistry.CallOpts)
}

// PrivateArtFactory is a free data retrieval call binding the contract method 0x924db13f.
//
// Solidity: function privateArtFactory() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCaller) PrivateArtFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAddressRegistry.contract.Call(opts, &out, "privateArtFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrivateArtFactory is a free data retrieval call binding the contract method 0x924db13f.
//
// Solidity: function privateArtFactory() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistrySession) PrivateArtFactory() (common.Address, error) {
	return _FantomAddressRegistry.Contract.PrivateArtFactory(&_FantomAddressRegistry.CallOpts)
}

// PrivateArtFactory is a free data retrieval call binding the contract method 0x924db13f.
//
// Solidity: function privateArtFactory() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCallerSession) PrivateArtFactory() (common.Address, error) {
	return _FantomAddressRegistry.Contract.PrivateArtFactory(&_FantomAddressRegistry.CallOpts)
}

// PrivateFactory is a free data retrieval call binding the contract method 0x735db939.
//
// Solidity: function privateFactory() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCaller) PrivateFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAddressRegistry.contract.Call(opts, &out, "privateFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrivateFactory is a free data retrieval call binding the contract method 0x735db939.
//
// Solidity: function privateFactory() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistrySession) PrivateFactory() (common.Address, error) {
	return _FantomAddressRegistry.Contract.PrivateFactory(&_FantomAddressRegistry.CallOpts)
}

// PrivateFactory is a free data retrieval call binding the contract method 0x735db939.
//
// Solidity: function privateFactory() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCallerSession) PrivateFactory() (common.Address, error) {
	return _FantomAddressRegistry.Contract.PrivateFactory(&_FantomAddressRegistry.CallOpts)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCaller) TokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomAddressRegistry.contract.Call(opts, &out, "tokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistrySession) TokenRegistry() (common.Address, error) {
	return _FantomAddressRegistry.Contract.TokenRegistry(&_FantomAddressRegistry.CallOpts)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_FantomAddressRegistry *FantomAddressRegistryCallerSession) TokenRegistry() (common.Address, error) {
	return _FantomAddressRegistry.Contract.TokenRegistry(&_FantomAddressRegistry.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomAddressRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomAddressRegistry *FantomAddressRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.RenounceOwnership(&_FantomAddressRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.RenounceOwnership(&_FantomAddressRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomAddressRegistry *FantomAddressRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.TransferOwnership(&_FantomAddressRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.TransferOwnership(&_FantomAddressRegistry.TransactOpts, newOwner)
}

// UpdateArtFactory is a paid mutator transaction binding the contract method 0x8590f295.
//
// Solidity: function updateArtFactory(address _artFactory) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactor) UpdateArtFactory(opts *bind.TransactOpts, _artFactory common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.contract.Transact(opts, "updateArtFactory", _artFactory)
}

// UpdateArtFactory is a paid mutator transaction binding the contract method 0x8590f295.
//
// Solidity: function updateArtFactory(address _artFactory) returns()
func (_FantomAddressRegistry *FantomAddressRegistrySession) UpdateArtFactory(_artFactory common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateArtFactory(&_FantomAddressRegistry.TransactOpts, _artFactory)
}

// UpdateArtFactory is a paid mutator transaction binding the contract method 0x8590f295.
//
// Solidity: function updateArtFactory(address _artFactory) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactorSession) UpdateArtFactory(_artFactory common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateArtFactory(&_FantomAddressRegistry.TransactOpts, _artFactory)
}

// UpdateArtFactoryPrivate is a paid mutator transaction binding the contract method 0xb717bb22.
//
// Solidity: function updateArtFactoryPrivate(address _privateArtFactory) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactor) UpdateArtFactoryPrivate(opts *bind.TransactOpts, _privateArtFactory common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.contract.Transact(opts, "updateArtFactoryPrivate", _privateArtFactory)
}

// UpdateArtFactoryPrivate is a paid mutator transaction binding the contract method 0xb717bb22.
//
// Solidity: function updateArtFactoryPrivate(address _privateArtFactory) returns()
func (_FantomAddressRegistry *FantomAddressRegistrySession) UpdateArtFactoryPrivate(_privateArtFactory common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateArtFactoryPrivate(&_FantomAddressRegistry.TransactOpts, _privateArtFactory)
}

// UpdateArtFactoryPrivate is a paid mutator transaction binding the contract method 0xb717bb22.
//
// Solidity: function updateArtFactoryPrivate(address _privateArtFactory) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactorSession) UpdateArtFactoryPrivate(_privateArtFactory common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateArtFactoryPrivate(&_FantomAddressRegistry.TransactOpts, _privateArtFactory)
}

// UpdateArtion is a paid mutator transaction binding the contract method 0xce35688a.
//
// Solidity: function updateArtion(address _artion) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactor) UpdateArtion(opts *bind.TransactOpts, _artion common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.contract.Transact(opts, "updateArtion", _artion)
}

// UpdateArtion is a paid mutator transaction binding the contract method 0xce35688a.
//
// Solidity: function updateArtion(address _artion) returns()
func (_FantomAddressRegistry *FantomAddressRegistrySession) UpdateArtion(_artion common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateArtion(&_FantomAddressRegistry.TransactOpts, _artion)
}

// UpdateArtion is a paid mutator transaction binding the contract method 0xce35688a.
//
// Solidity: function updateArtion(address _artion) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactorSession) UpdateArtion(_artion common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateArtion(&_FantomAddressRegistry.TransactOpts, _artion)
}

// UpdateAuction is a paid mutator transaction binding the contract method 0x3402a036.
//
// Solidity: function updateAuction(address _auction) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactor) UpdateAuction(opts *bind.TransactOpts, _auction common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.contract.Transact(opts, "updateAuction", _auction)
}

// UpdateAuction is a paid mutator transaction binding the contract method 0x3402a036.
//
// Solidity: function updateAuction(address _auction) returns()
func (_FantomAddressRegistry *FantomAddressRegistrySession) UpdateAuction(_auction common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateAuction(&_FantomAddressRegistry.TransactOpts, _auction)
}

// UpdateAuction is a paid mutator transaction binding the contract method 0x3402a036.
//
// Solidity: function updateAuction(address _auction) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactorSession) UpdateAuction(_auction common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateAuction(&_FantomAddressRegistry.TransactOpts, _auction)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactor) UpdateBundleMarketplace(opts *bind.TransactOpts, _bundleMarketplace common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.contract.Transact(opts, "updateBundleMarketplace", _bundleMarketplace)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_FantomAddressRegistry *FantomAddressRegistrySession) UpdateBundleMarketplace(_bundleMarketplace common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateBundleMarketplace(&_FantomAddressRegistry.TransactOpts, _bundleMarketplace)
}

// UpdateBundleMarketplace is a paid mutator transaction binding the contract method 0x9d0e24a7.
//
// Solidity: function updateBundleMarketplace(address _bundleMarketplace) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactorSession) UpdateBundleMarketplace(_bundleMarketplace common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateBundleMarketplace(&_FantomAddressRegistry.TransactOpts, _bundleMarketplace)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactor) UpdateMarketplace(opts *bind.TransactOpts, _marketplace common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.contract.Transact(opts, "updateMarketplace", _marketplace)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_FantomAddressRegistry *FantomAddressRegistrySession) UpdateMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateMarketplace(&_FantomAddressRegistry.TransactOpts, _marketplace)
}

// UpdateMarketplace is a paid mutator transaction binding the contract method 0xeb636317.
//
// Solidity: function updateMarketplace(address _marketplace) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactorSession) UpdateMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateMarketplace(&_FantomAddressRegistry.TransactOpts, _marketplace)
}

// UpdateNFTFactory is a paid mutator transaction binding the contract method 0x975a426a.
//
// Solidity: function updateNFTFactory(address _factory) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactor) UpdateNFTFactory(opts *bind.TransactOpts, _factory common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.contract.Transact(opts, "updateNFTFactory", _factory)
}

// UpdateNFTFactory is a paid mutator transaction binding the contract method 0x975a426a.
//
// Solidity: function updateNFTFactory(address _factory) returns()
func (_FantomAddressRegistry *FantomAddressRegistrySession) UpdateNFTFactory(_factory common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateNFTFactory(&_FantomAddressRegistry.TransactOpts, _factory)
}

// UpdateNFTFactory is a paid mutator transaction binding the contract method 0x975a426a.
//
// Solidity: function updateNFTFactory(address _factory) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactorSession) UpdateNFTFactory(_factory common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateNFTFactory(&_FantomAddressRegistry.TransactOpts, _factory)
}

// UpdateNFTFactoryPrivate is a paid mutator transaction binding the contract method 0xe35bead5.
//
// Solidity: function updateNFTFactoryPrivate(address _privateFactory) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactor) UpdateNFTFactoryPrivate(opts *bind.TransactOpts, _privateFactory common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.contract.Transact(opts, "updateNFTFactoryPrivate", _privateFactory)
}

// UpdateNFTFactoryPrivate is a paid mutator transaction binding the contract method 0xe35bead5.
//
// Solidity: function updateNFTFactoryPrivate(address _privateFactory) returns()
func (_FantomAddressRegistry *FantomAddressRegistrySession) UpdateNFTFactoryPrivate(_privateFactory common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateNFTFactoryPrivate(&_FantomAddressRegistry.TransactOpts, _privateFactory)
}

// UpdateNFTFactoryPrivate is a paid mutator transaction binding the contract method 0xe35bead5.
//
// Solidity: function updateNFTFactoryPrivate(address _privateFactory) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactorSession) UpdateNFTFactoryPrivate(_privateFactory common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateNFTFactoryPrivate(&_FantomAddressRegistry.TransactOpts, _privateFactory)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0x95877f78.
//
// Solidity: function updatePriceFeed(address _priceFeed) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactor) UpdatePriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.contract.Transact(opts, "updatePriceFeed", _priceFeed)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0x95877f78.
//
// Solidity: function updatePriceFeed(address _priceFeed) returns()
func (_FantomAddressRegistry *FantomAddressRegistrySession) UpdatePriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdatePriceFeed(&_FantomAddressRegistry.TransactOpts, _priceFeed)
}

// UpdatePriceFeed is a paid mutator transaction binding the contract method 0x95877f78.
//
// Solidity: function updatePriceFeed(address _priceFeed) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactorSession) UpdatePriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdatePriceFeed(&_FantomAddressRegistry.TransactOpts, _priceFeed)
}

// UpdateTokenRegistry is a paid mutator transaction binding the contract method 0x1fa915a4.
//
// Solidity: function updateTokenRegistry(address _tokenRegistry) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactor) UpdateTokenRegistry(opts *bind.TransactOpts, _tokenRegistry common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.contract.Transact(opts, "updateTokenRegistry", _tokenRegistry)
}

// UpdateTokenRegistry is a paid mutator transaction binding the contract method 0x1fa915a4.
//
// Solidity: function updateTokenRegistry(address _tokenRegistry) returns()
func (_FantomAddressRegistry *FantomAddressRegistrySession) UpdateTokenRegistry(_tokenRegistry common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateTokenRegistry(&_FantomAddressRegistry.TransactOpts, _tokenRegistry)
}

// UpdateTokenRegistry is a paid mutator transaction binding the contract method 0x1fa915a4.
//
// Solidity: function updateTokenRegistry(address _tokenRegistry) returns()
func (_FantomAddressRegistry *FantomAddressRegistryTransactorSession) UpdateTokenRegistry(_tokenRegistry common.Address) (*types.Transaction, error) {
	return _FantomAddressRegistry.Contract.UpdateTokenRegistry(&_FantomAddressRegistry.TransactOpts, _tokenRegistry)
}

// FantomAddressRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomAddressRegistry contract.
type FantomAddressRegistryOwnershipTransferredIterator struct {
	Event *FantomAddressRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomAddressRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomAddressRegistryOwnershipTransferred)
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
		it.Event = new(FantomAddressRegistryOwnershipTransferred)
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
func (it *FantomAddressRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomAddressRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomAddressRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the FantomAddressRegistry contract.
type FantomAddressRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomAddressRegistry *FantomAddressRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomAddressRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomAddressRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomAddressRegistryOwnershipTransferredIterator{contract: _FantomAddressRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomAddressRegistry *FantomAddressRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomAddressRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomAddressRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomAddressRegistryOwnershipTransferred)
				if err := _FantomAddressRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FantomAddressRegistry *FantomAddressRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*FantomAddressRegistryOwnershipTransferred, error) {
	event := new(FantomAddressRegistryOwnershipTransferred)
	if err := _FantomAddressRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// Deprecated: Use IERC165MetaData.Sigs instead.
// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = IERC165MetaData.Sigs

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
