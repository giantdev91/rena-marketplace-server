// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace

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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220da7f2527387fc4e2fe10fa316278aa2e3c5805c8cf510684b0db0d036cd1c91864736f6c634300080b0033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220dfa53fcea62b27496a675d253bf7cabbcfb7c077e01f71075304e92d9c54f2b164736f6c634300080b0033",
}

// AddressUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressUpgradeableMetaData.ABI instead.
var AddressUpgradeableABI = AddressUpgradeableMetaData.ABI

// AddressUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressUpgradeableMetaData.Bin instead.
var AddressUpgradeableBin = AddressUpgradeableMetaData.Bin

// DeployAddressUpgradeable deploys a new Ethereum contract, binding an instance of AddressUpgradeable to it.
func DeployAddressUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUpgradeable, error) {
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// AddressUpgradeable is an auto generated Go binding around an Ethereum contract.
type AddressUpgradeable struct {
	AddressUpgradeableCaller     // Read-only binding to the contract
	AddressUpgradeableTransactor // Write-only binding to the contract
	AddressUpgradeableFilterer   // Log filterer for contract events
}

// AddressUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUpgradeableSession struct {
	Contract     *AddressUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AddressUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUpgradeableCallerSession struct {
	Contract *AddressUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AddressUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUpgradeableTransactorSession struct {
	Contract     *AddressUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AddressUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUpgradeableRaw struct {
	Contract *AddressUpgradeable // Generic contract binding to access the raw methods on
}

// AddressUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUpgradeableCallerRaw struct {
	Contract *AddressUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactorRaw struct {
	Contract *AddressUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUpgradeable creates a new instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeable(address common.Address, backend bind.ContractBackend) (*AddressUpgradeable, error) {
	contract, err := bindAddressUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// NewAddressUpgradeableCaller creates a new read-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AddressUpgradeableCaller, error) {
	contract, err := bindAddressUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableCaller{contract: contract}, nil
}

// NewAddressUpgradeableTransactor creates a new write-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUpgradeableTransactor, error) {
	contract, err := bindAddressUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableTransactor{contract: contract}, nil
}

// NewAddressUpgradeableFilterer creates a new log filterer instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUpgradeableFilterer, error) {
	contract, err := bindAddressUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableFilterer{contract: contract}, nil
}

// bindAddressUpgradeable binds a generic wrapper to an already deployed contract.
func bindAddressUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.AddressUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transact(opts, method, params...)
}

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

// FantomMarketplaceMetaData contains all meta data concerning the FantomMarketplace contract.
var FantomMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_platformFee\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ItemCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"name\":\"ItemListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"unitPrice\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"}],\"name\":\"ItemSold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"ItemUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OfferCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"OfferCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"platformFee\",\"type\":\"uint16\"}],\"name\":\"UpdatePlatformFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"UpdatePlatformFeeRecipient\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressRegistry\",\"outputs\":[{\"internalType\":\"contractIFantomAddressRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"buyItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collectionRoyalties\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"royalty\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"}],\"name\":\"createOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceipient\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingTime\",\"type\":\"uint256\"}],\"name\":\"listItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerItem\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_royalty\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"registerCollectionRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_royalty\",\"type\":\"uint16\"}],\"name\":\"registerRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"royalties\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"updateAddressRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_payToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newPrice\",\"type\":\"uint256\"}],\"name\":\"updateListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_platformFee\",\"type\":\"uint16\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"updatePlatformFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"validateItemSold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3bbb2806": "acceptOffer(address,uint256,address)",
		"f3ad65f4": "addressRegistry()",
		"259ca365": "buyItem(address,uint256,address)",
		"85f9d345": "buyItem(address,uint256,address,address)",
		"b2ddee06": "cancelListing(address,uint256)",
		"058a56ac": "cancelOffer(address,uint256)",
		"57b275b4": "collectionRoyalties(address)",
		"42cbb4b8": "createOffer(address,uint256,address,uint256,uint256,uint256)",
		"3740ebb3": "feeReceipient()",
		"41976e09": "getPrice(address)",
		"3fc1cc26": "listItem(address,uint256,uint256,address,uint256,uint256)",
		"6bd3a64b": "listings(address,uint256,address)",
		"78746d0d": "minters(address,uint256)",
		"d3f494cc": "offers(address,uint256,address)",
		"8da5cb5b": "owner()",
		"26232a2e": "platformFee()",
		"c1661a9a": "registerCollectionRoyalty(address,address,uint16,address)",
		"f3880b6e": "registerRoyalty(address,uint256,uint16)",
		"715018a6": "renounceOwnership()",
		"e1e549c4": "royalties(address,uint256)",
		"f2fde38b": "transferOwnership(address)",
		"68e79e89": "updateAddressRegistry(address)",
		"e984f2eb": "updateListing(address,uint256,address,uint256)",
		"afb06952": "updatePlatformFee(uint16)",
		"f5fe7f71": "updatePlatformFeeRecipient(address)",
		"e940ebeb": "validateItemSold(address,uint256,address,address)",
	},
	Bin: "0x60806040523480156200001157600080fd5b50604051620053c3380380620053c38339810160408190526200003491620000c9565b6200003f3362000079565b60018055600680546001600160a01b0390931662010000026001600160b01b031990931661ffff9092169190911791909117905562000119565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60008060408385031215620000dd57600080fd5b82516001600160a01b0381168114620000f557600080fd5b602084015190925061ffff811681146200010e57600080fd5b809150509250929050565b61529a80620001296000396000f3fe6080604052600436106101815760003560e01c806385f9d345116100d1578063e1e549c41161008a578063f2fde38b11610064578063f2fde38b146105cd578063f3880b6e146105ed578063f3ad65f41461060d578063f5fe7f711461062d57600080fd5b8063e1e549c414610551578063e940ebeb1461058d578063e984f2eb146105ad57600080fd5b806385f9d3451461042c5780638da5cb5b1461044c578063afb069521461046a578063b2ddee061461048a578063c1661a9a146104aa578063d3f494cc146104ca57600080fd5b806341976e091161013e57806368e79e891161011857806368e79e89146103315780636bd3a64b14610351578063715018a6146103d657806378746d0d146103eb57600080fd5b806341976e091461026c57806342cbb4b81461029a57806357b275b4146102ba57600080fd5b8063058a56ac14610186578063259ca365146101a857806326232a2e146101bb5780633740ebb3146101ee5780633bbb28061461022c5780633fc1cc261461024c575b600080fd5b34801561019257600080fd5b506101a66101a1366004614a1a565b61064d565b005b6101a66101b6366004614a46565b610799565b3480156101c757600080fd5b506006546101d69061ffff1681565b60405161ffff90911681526020015b60405180910390f35b3480156101fa57600080fd5b50600654610214906201000090046001600160a01b031681565b6040516001600160a01b0390911681526020016101e5565b34801561023857600080fd5b506101a6610247366004614a46565b610c16565b34801561025857600080fd5b506101a6610267366004614a88565b611582565b34801561027857600080fd5b5061028c610287366004614ae4565b611bbb565b6040519081526020016101e5565b3480156102a657600080fd5b506101a66102b5366004614b01565b611e66565b3480156102c657600080fd5b506103076102d5366004614ae4565b6007602052600090815260409020805460019091015461ffff8216916001600160a01b03620100009091048116911683565b6040805161ffff90941684526001600160a01b0392831660208501529116908201526060016101e5565b34801561033d57600080fd5b506101a661034c366004614ae4565b612470565b34801561035d57600080fd5b506103ae61036c366004614a46565b6004602090815260009384526040808520825292845282842090528252902080546001820154600283015460039093015491926001600160a01b039091169184565b604080519485526001600160a01b0390931660208501529183015260608201526080016101e5565b3480156103e257600080fd5b506101a66124bc565b3480156103f757600080fd5b50610214610406366004614a1a565b60026020908152600092835260408084209091529082529020546001600160a01b031681565b34801561043857600080fd5b506101a6610447366004614b5d565b6124f2565b34801561045857600080fd5b506000546001600160a01b0316610214565b34801561047657600080fd5b506101a6610485366004614bc7565b6128f9565b34801561049657600080fd5b506101a66104a5366004614a1a565b61296e565b3480156104b657600080fd5b506101a66104c5366004614be2565b612a2f565b3480156104d657600080fd5b506105276104e5366004614a46565b600560209081526000938452604080852082529284528284209052825290208054600182015460028301546003909301546001600160a01b0390921692909184565b604080516001600160a01b03909516855260208501939093529183015260608201526080016101e5565b34801561055d57600080fd5b506101d661056c366004614a1a565b600360209081526000928352604080842090915290825290205461ffff1681565b34801561059957600080fd5b506101a66105a8366004614b5d565b612c6a565b3480156105b957600080fd5b506101a66105c8366004614c21565b612e37565b3480156105d957600080fd5b506101a66105e8366004614ae4565b61333f565b3480156105f957600080fd5b506101a6610608366004614c69565b6133da565b34801561061957600080fd5b50600854610214906001600160a01b031681565b34801561063957600080fd5b506101a6610648366004614ae4565b61373c565b8181336001600160a01b0380841660009081526005602090815260408083208684528252808320858516845282529182902082516080810184528154909416845260018101549184018290526002810154928401929092526003909101546060830152158015906106c15750428160600151115b6107125760405162461bcd60e51b815260206004820152601b60248201527f6f66666572206e6f7420657869737473206f722065787069726564000000000060448201526064015b60405180910390fd5b6001600160a01b038616600081815260056020908152604080832089845282528083203380855290835281842080546001600160a01b0319168155600181018590556002810185905560030193909355518881527fc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7910160405180910390a3505050505050565b600260015414156107bc5760405162461bcd60e51b815260040161070990614ca7565b600260018181556001600160a01b0380861660009081526004602090815260408083208884528252808320848816845282529182902082516080810184528154808252958201549094169184019190915293840154908201526003909201546060830152849184918491906108435760405162461bcd60e51b815260040161070990614cde565b6001600160a01b0380881660008181526004602081815260408084208c85528252808420868c16855282529283902083516080810185528154815260018201549096169186019190915260028101548584015260030154606085015290516301ffc9a760e01b81528a938a938a93919290916301ffc9a7916108ce916380ac58cd60e01b9101614d07565b602060405180830381865afa1580156108eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061090f9190614d2c565b156109b0576040516331a9108f60e11b81526004810184905284906001600160a01b038481169190831690636352211e90602401602060405180830381865afa158015610960573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109849190614d47565b6001600160a01b0316146109aa5760405162461bcd60e51b815260040161070990614d64565b50610ad4565b6040516301ffc9a760e01b81526001600160a01b038516906301ffc9a7906109e390636cdb3d1360e11b90600401614d07565b602060405180830381865afa158015610a00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a249190614d2c565b15610abc578051604051627eeac760e11b81526001600160a01b038481166004830152602482018690528692919083169062fdd58e90604401602060405180830381865afa158015610a7a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a9e9190614d8d565b10156109aa5760405162461bcd60e51b815260040161070990614d64565b60405162461bcd60e51b815260040161070990614da6565b6060810151421015610b1b5760405162461bcd60e51b815260206004820152601060248201526f6974656d206e6f742062757961626c6560801b6044820152606401610709565b6001600160a01b03808c1660009081526004602090815260408083208e845282528083208d851684528252918290208251608081018452815481526001820154909416918401829052600281015492840192909252600390910154606083015215610b985760405162461bcd60e51b815260040161070990614dd3565b80516040820151610ba8916137be565b341015610bf75760405162461bcd60e51b815260206004820152601b60248201527f696e73756666696369656e742062616c616e636520746f2062757900000000006044820152606401610709565b610c048c8c60008d6137d3565b50506001805550505050505050505050565b60026001541415610c395760405162461bcd60e51b815260040161070990614ca7565b600260018181556001600160a01b03808616600090815260056020908152604080832088845282528083208488168452825291829020825160808101845281549094168452938401549083018190529383015490820152600390910154606082015284918491849115801590610cb25750428160600151115b610cfe5760405162461bcd60e51b815260206004820152601b60248201527f6f66666572206e6f7420657869737473206f72206578706972656400000000006044820152606401610709565b6001600160a01b0380881660008181526005602090815260408083208b845282528083208a86168452825291829020825160808101845281549095168552600181015491850191909152600281015484830152600301546060840152516301ffc9a760e01b81526301ffc9a790610d80906380ac58cd60e01b90600401614d07565b602060405180830381865afa158015610d9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc19190614d2c565b15610e615787336040516331a9108f60e11b8152600481018a90526001600160a01b0391821691831690636352211e90602401602060405180830381865afa158015610e11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e359190614d47565b6001600160a01b031614610e5b5760405162461bcd60e51b815260040161070990614d64565b50610f70565b6040516301ffc9a760e01b81526001600160a01b038916906301ffc9a790610e9490636cdb3d1360e11b90600401614d07565b602060405180830381865afa158015610eb1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ed59190614d2c565b15610abc5760208181015160408051627eeac760e11b8152336004820152602481018b905290518b936001600160a01b0385169262fdd58e92604480830193928290030181865afa158015610f2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f529190614d8d565b1015610e5b5760405162461bcd60e51b815260040161070990614d64565b6000610f8d826020015183604001516137be90919063ffffffff16565b600654909150600090610fb3906103e890610fad90859061ffff166137be565b90613fa7565b6006548451919250600091610fdd916001600160a01b03918216918c916201000090041685613fb3565b6001600160a01b03808c1660008181526002602090815260408083208f8452825280832054938352600382528083208f845290915290205491169061ffff16811580159061102e575061ffff811615155b1561107c57611051612710610fad61ffff841661104b8989614013565b906137be565b865190935061106b906001600160a01b03168c8486613fb3565b611075848461401f565b93506110fb565b50506001600160a01b03808c1660009081526007602052604090206001810154905491169061ffff1681158015906110b7575061ffff811615155b156110fb576110d4612710610fad61ffff841661104b8989614013565b86519093506110ee906001600160a01b03168c8486613fb3565b6110f8848461401f565b93505b61111d8b3361110a8888614013565b89516001600160a01b0316929190613fb3565b6040516301ffc9a760e01b81526001600160a01b038e16906301ffc9a790611150906380ac58cd60e01b90600401614d07565b602060405180830381865afa15801561116d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111919190614d2c565b15611213576001600160a01b038d166342842e0e336040516001600160e01b031960e084901b1681526001600160a01b039182166004820152908e166024820152604481018f9052606401600060405180830381600087803b1580156111f657600080fd5b505af115801561120a573d6000803e3d6000fd5b5050505061128d565b6001600160a01b038d1663f242432a338d8f8a60200151604051806020016040528060008152506040518663ffffffff1660e01b815260040161125a959493929190614e56565b600060405180830381600087803b15801561127457600080fd5b505af1158015611288573d6000803e3d6000fd5b505050505b600860009054906101000a90046001600160a01b03166001600160a01b0316639e1f5b286040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113049190614d47565b6001600160a01b03166327c1f4f48e8e89602001516040518463ffffffff1660e01b8152600401611353939291906001600160a01b039390931683526020830191909152604082015260600190565b600060405180830381600087803b15801561136d57600080fd5b505af1158015611381573d6000803e3d6000fd5b50505050600460008e6001600160a01b03166001600160a01b0316815260200190815260200160002060008d815260200190815260200160002060006113c43390565b6001600160a01b03166001600160a01b031681526020019081526020016000206000808201600090556001820160006101000a8154906001600160a01b030219169055600282016000905560038201600090555050600560008e6001600160a01b03166001600160a01b0316815260200190815260200160002060008d815260200190815260200160002060008c6001600160a01b03166001600160a01b03168152602001908152602001600020600080820160006101000a8154906001600160a01b03021916905560018201600090556002820160009055600382016000905550508c6001600160a01b03168b6001600160a01b03166114c23390565b6001600160a01b03167f949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d8f8a602001518b600001516115048d60000151611bbb565b8d6040015160405161151a959493929190614e90565b60405180910390a48c6001600160a01b03168b6001600160a01b03167fc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a78e60405161156791815260200190565b60405180910390a35050600180555050505050505050505050565b8585336001600160a01b0380841660009081526004602090815260408083208684528252808320858516845282529182902082516080810184528154808252600183015490951692810192909252600281015492820192909252600390910154606082015290156116265760405162461bcd60e51b815260206004820152600e60248201526d185b1c9958591e481b1a5cdd195960921b6044820152606401610709565b6040516301ffc9a760e01b81526001600160a01b038b16906301ffc9a790611659906380ac58cd60e01b90600401614d07565b602060405180830381865afa158015611676573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061169a9190614d2c565b156117f85789336040516331a9108f60e11b8152600481018c90526001600160a01b0391821691831690636352211e90602401602060405180830381865afa1580156116ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061170e9190614d47565b6001600160a01b0316146117345760405162461bcd60e51b815260040161070990614d64565b6001600160a01b03811663e985e9c5336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152306024820152604401602060405180830381865afa15801561178e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117b29190614d2c565b6117f25760405162461bcd60e51b81526020600482015260116024820152701a5d195b481b9bdd08185c1c1c9bdd9959607a1b6044820152606401610709565b50611937565b6040516301ffc9a760e01b81526001600160a01b038b16906301ffc9a79061182b90636cdb3d1360e11b90600401614d07565b602060405180830381865afa158015611848573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061186c9190614d2c565b15610abc5789886001600160a01b03821662fdd58e336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018e9052604401602060405180830381865afa1580156118cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118f19190614d8d565b10156117345760405162461bcd60e51b81526020600482015260156024820152746d75737420686f6c6420656e6f756768206e66747360581b6044820152606401610709565b6001600160a01b0387161580611aab575060085460408051639d23c4c760e01b815290516000926001600160a01b031691639d23c4c79160048083019260209291908290030181865afa158015611992573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119b69190614d47565b6001600160a01b031614158015611aab5750600860009054906101000a90046001600160a01b03166001600160a01b0316639d23c4c76040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a3f9190614d47565b6040516398afdfe360e01b81526001600160a01b03898116600483015291909116906398afdfe390602401602060405180830381865afa158015611a87573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aab9190614d2c565b611ac75760405162461bcd60e51b815260040161070990614dd3565b604080516080810182528981526001600160a01b038981166020808401919091528284018a905260608301899052908d166000908152600482528381208d82529091529182209091336001600160a01b039081168252602080830193909352604091820160002084518155928401516001840180546001600160a01b0319169183169190911790559083015160028301556060909201516003909101558a16336001600160a01b03167fa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa26298b8b8b8b8b604051611ba7959493929190614e90565b60405180910390a350505050505050505050565b600080806001600160a01b038416611d215760085460408051633a0df78d60e11b815290516000926001600160a01b03169163741bef1a9160048083019260209291908290030181865afa158015611c17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c3b9190614d47565b9050806001600160a01b03166341976e09826001600160a01b031663050d63ec6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cae9190614d47565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024016040805180830381865afa158015611cf1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d159190614ebc565b9093509150611e099050565b600860009054906101000a90046001600160a01b03166001600160a01b031663741bef1a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611d74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d989190614d47565b6040516341976e0960e01b81526001600160a01b03868116600483015291909116906341976e09906024016040805180830381865afa158015611ddf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e039190614ebc565b90925090505b60128160ff161015611e3c57611e20816012614f08565b611e2b90600a615060565b611e35908361506f565b9150611e5f565b611e47601282614f08565b611e5290600a615060565b611e5c908361510a565b91505b5092915050565b8585336001600160a01b03808416600090815260056020908152604080832086845282528083208585168452825291829020825160808101845281549094168452600181015491840182905260028101549284019290925260039091015460608301521580611ed9575042816060015111155b611f1d5760405162461bcd60e51b81526020600482015260156024820152741bd999995c88185b1c9958591e4818dc99585d1959605a1b6044820152606401610709565b6040516301ffc9a760e01b81526001600160a01b038b16906301ffc9a790611f50906380ac58cd60e01b90600401614d07565b602060405180830381865afa158015611f6d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f919190614d2c565b8061200b57506040516301ffc9a760e01b81526001600160a01b038b16906301ffc9a790611fca90636cdb3d1360e11b90600401614d07565b602060405180830381865afa158015611fe7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061200b9190614d2c565b6120275760405162461bcd60e51b815260040161070990614da6565b60085460408051637d9f6db560e01b815290516000926001600160a01b031691637d9f6db59160048083019260209291908290030181865afa158015612071573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120959190614d47565b60405163227c8e0f60e11b81526001600160a01b038d81166004830152602482018d90529192506000918291908416906344f91c1e9060440160c060405180830381865afa1580156120eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061210f9190615138565b9550509450505050816000148061212857506001811515145b6121895760405162461bcd60e51b815260206004820152602c60248201527f63616e6e6f7420706c61636520616e206f666665722069662061756374696f6e60448201526b1034b99033b7b4b7339037b760a11b6064820152608401610709565b4288116121cd5760405162461bcd60e51b815260206004820152601260248201527134b73b30b634b21032bc3834b930ba34b7b760711b6044820152606401610709565b6001600160a01b038b161580612341575060085460408051639d23c4c760e01b815290516000926001600160a01b031691639d23c4c79160048083019260209291908290030181865afa158015612228573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061224c9190614d47565b6001600160a01b0316141580156123415750600860009054906101000a90046001600160a01b03166001600160a01b0316639d23c4c76040518163ffffffff1660e01b8152600401602060405180830381865afa1580156122b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122d59190614d47565b6040516398afdfe360e01b81526001600160a01b038d8116600483015291909116906398afdfe390602401602060405180830381865afa15801561231d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123419190614d2c565b61235d5760405162461bcd60e51b815260040161070990614dd3565b60405180608001604052808c6001600160a01b031681526020018b81526020018a815260200189815250600560008f6001600160a01b03166001600160a01b0316815260200190815260200160002060008e815260200190815260200160002060006123c63390565b6001600160a01b0390811682526020808301939093526040918201600020845181546001600160a01b0319169083161781559284015160018401559083015160028301556060909201516003909101558d16336001600160a01b03167f89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f68e8d8f8e8e604051612459959493929190614e90565b60405180910390a350505050505050505050505050565b6000546001600160a01b0316331461249a5760405162461bcd60e51b81526004016107099061519e565b600880546001600160a01b0319166001600160a01b0392909216919091179055565b6000546001600160a01b031633146124e65760405162461bcd60e51b81526004016107099061519e565b6124f0600061402b565b565b600260015414156125155760405162461bcd60e51b815260040161070990614ca7565b600260018181556001600160a01b03808716600090815260046020908152604080832089845282528083208488168452825291829020825160808101845281548082529582015490941691840191909152938401549082015260039092015460608301528591859184919061259c5760405162461bcd60e51b815260040161070990614cde565b6001600160a01b0380891660008181526004602081815260408084208d85528252808420868c16855282529283902083516080810185528154815260018201549096169186019190915260028101548584015260030154606085015290516301ffc9a760e01b81528b938b938a93919290916301ffc9a791612627916380ac58cd60e01b9101614d07565b602060405180830381865afa158015612644573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126689190614d2c565b15612709576040516331a9108f60e11b81526004810184905284906001600160a01b038481169190831690636352211e90602401602060405180830381865afa1580156126b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126dd9190614d47565b6001600160a01b0316146127035760405162461bcd60e51b815260040161070990614d64565b50612815565b6040516301ffc9a760e01b81526001600160a01b038516906301ffc9a79061273c90636cdb3d1360e11b90600401614d07565b602060405180830381865afa158015612759573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061277d9190614d2c565b15610abc578051604051627eeac760e11b81526001600160a01b038481166004830152602482018690528692919083169062fdd58e90604401602060405180830381865afa1580156127d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127f79190614d8d565b10156127035760405162461bcd60e51b815260040161070990614d64565b606081015142101561285c5760405162461bcd60e51b815260206004820152601060248201526f6974656d206e6f742062757961626c6560801b6044820152606401610709565b6001600160a01b03808d1660009081526004602090815260408083208f845282528083208d851684528252918290208251608081018452815481526001820154851692810183905260028201549381019390935260030154606083015290918c16146128da5760405162461bcd60e51b815260040161070990614dd3565b6128e68d8d8d8d6137d3565b5050600180555050505050505050505050565b6000546001600160a01b031633146129235760405162461bcd60e51b81526004016107099061519e565b6006805461ffff191661ffff83169081179091556040519081527fcabead39256298cdf7d1228eee68b47bc54a2192283d8d93574da08d2d7ac303906020015b60405180910390a150565b600260015414156129915760405162461bcd60e51b815260040161070990614ca7565b60026001558181336001600160a01b0380841660009081526004602090815260408083208684528252808320858516845282529182902082516080810184528154808252600183015490951692810192909252600281015492820192909252600390910154606082015290612a185760405162461bcd60e51b815260040161070990614cde565b612a2386863361407b565b50506001805550505050565b6000546001600160a01b03163314612a595760405162461bcd60e51b81526004016107099061519e565b6001600160a01b038316612aaf5760405162461bcd60e51b815260206004820152601760248201527f696e76616c69642063726561746f7220616464726573730000000000000000006044820152606401610709565b6127108261ffff161115612af75760405162461bcd60e51b815260206004820152600f60248201526e696e76616c696420726f79616c747960881b6044820152606401610709565b61ffff82161580612b1057506001600160a01b03811615155b612b5c5760405162461bcd60e51b815260206004820152601d60248201527f696e76616c69642066656520726563697069656e7420616464726573730000006044820152606401610709565b612b6584614383565b15612b825760405162461bcd60e51b815260040161070990614da6565b6001600160a01b038481166000908152600760205260409020546201000090041615612be65760405162461bcd60e51b81526020600482015260136024820152721c9bde585b1d1e48185b1c9958591e481cd95d606a1b6044820152606401610709565b6040805160608101825261ffff93841681526001600160a01b039485166020808301918252938616828401908152968616600090815260079094529190922091518254915193166001600160b01b03199091161762010000928416929092029190911781559151600190920180546001600160a01b03191692909116919091179055565b600854604080516313c3eb6560e31b8152905133926001600160a01b031691639e1f5b289160048083019260209291908290030181865afa158015612cb3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cd79190614d47565b6001600160a01b031614612d375760405162461bcd60e51b815260206004820152602160248201527f73656e646572206d7573742062652062756e646c65206d61726b6574706c61636044820152606560f81b6064820152608401610709565b6001600160a01b038085166000908152600460209081526040808320878452825280832086851684528252918290208251608081018452815480825260018301549095169281019290925260028101549282019290925260039091015460608201529015612daa57612daa85858561407b565b6001600160a01b038581166000818152600560209081526040808320898452825280832094871680845294825280832080546001600160a01b031916815560018101849055600281018490556003019290925590518781529192917fc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7910160405180910390a35050505050565b60026001541415612e5a5760405162461bcd60e51b815260040161070990614ca7565b60026001558383336001600160a01b0380841660009081526004602090815260408083208684528252808320858516845282529182902082516080810184528154808252600183015490951692810192909252600281015492820192909252600390910154606082015290612ee15760405162461bcd60e51b815260040161070990614cde565b6001600160a01b03881660008181526004602081815260408084208c855282528084203385529091529182902091516301ffc9a760e01b81529192916301ffc9a791612f36916380ac58cd60e01b9101614d07565b602060405180830381865afa158015612f53573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f779190614d2c565b156130175788336040516331a9108f60e11b8152600481018b90526001600160a01b0391821691831690636352211e90602401602060405180830381865afa158015612fc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612feb9190614d47565b6001600160a01b0316146130115760405162461bcd60e51b815260040161070990614d64565b50613130565b6040516301ffc9a760e01b81526001600160a01b038a16906301ffc9a79061304a90636cdb3d1360e11b90600401614d07565b602060405180830381865afa158015613067573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061308b9190614d2c565b15610abc57805489906001600160a01b03821662fdd58e336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018d9052604401602060405180830381865afa1580156130ee573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131129190614d8d565b10156130115760405162461bcd60e51b815260040161070990614d64565b6001600160a01b03871615806132a4575060085460408051639d23c4c760e01b815290516000926001600160a01b031691639d23c4c79160048083019260209291908290030181865afa15801561318b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131af9190614d47565b6001600160a01b0316141580156132a45750600860009054906101000a90046001600160a01b03166001600160a01b0316639d23c4c76040518163ffffffff1660e01b8152600401602060405180830381865afa158015613214573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132389190614d47565b6040516398afdfe360e01b81526001600160a01b03898116600483015291909116906398afdfe390602401602060405180830381865afa158015613280573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132a49190614d2c565b6132c05760405162461bcd60e51b815260040161070990614dd3565b6001810180546001600160a01b038981166001600160a01b0319909216821790925560028301889055604080518b8152602081019290925281810189905251918b169133917f60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd919081900360600190a350506001805550505050505050565b6000546001600160a01b031633146133695760405162461bcd60e51b81526004016107099061519e565b6001600160a01b0381166133ce5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610709565b6133d78161402b565b50565b6127108161ffff1611156134225760405162461bcd60e51b815260206004820152600f60248201526e696e76616c696420726f79616c747960881b6044820152606401610709565b61342b83614383565b6134475760405162461bcd60e51b815260040161070990614da6565b6040516301ffc9a760e01b81526001600160a01b038416906301ffc9a79061347a906380ac58cd60e01b90600401614d07565b602060405180830381865afa158015613497573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134bb9190614d2c565b1561355b5782336040516331a9108f60e11b8152600481018590526001600160a01b0391821691831690636352211e90602401602060405180830381865afa15801561350b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061352f9190614d47565b6001600160a01b0316146135555760405162461bcd60e51b815260040161070990614d64565b50613674565b6040516301ffc9a760e01b81526001600160a01b038416906301ffc9a79061358e90636cdb3d1360e11b90600401614d07565b602060405180830381865afa1580156135ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135cf9190614d2c565b15613674578260006001600160a01b03821662fdd58e336040516001600160e01b031960e084901b1681526001600160a01b03909116600482015260248101879052604401602060405180830381865afa158015613631573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136559190614d8d565b116136725760405162461bcd60e51b815260040161070990614d64565b505b6001600160a01b03838116600090815260026020908152604080832086845290915290205416156136dd5760405162461bcd60e51b81526020600482015260136024820152721c9bde585b1d1e48185b1c9958591e481cd95d606a1b6044820152606401610709565b6001600160a01b03929092166000818152600260209081526040808320858452825280832080546001600160a01b0319163317905592825260038152828220938252929092529020805461ffff90921661ffff19909216919091179055565b6000546001600160a01b031633146137665760405162461bcd60e51b81526004016107099061519e565b6006805462010000600160b01b031916620100006001600160a01b038416908102919091179091556040519081527fe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca790602001612963565b60006137ca82846151d3565b90505b92915050565b6001600160a01b038085166000908152600460209081526040808320878452825280832085851684528252808320815160808101835281548082526001830154909616938101939093526002810154918301829052600301546060830152909261383d91906137be565b60065490915060009061385d906103e890610fad90859061ffff166137be565b90506001600160a01b038516613912576006546040516000916201000090046001600160a01b03169083908381818185875af1925050503d80600081146138c0576040519150601f19603f3d011682016040523d82523d6000602084013e6138c5565b606091505b505090508061390c5760405162461bcd60e51b8152602060048201526013602482015272199959481d1c985b9cd9995c8819985a5b1959606a1b6044820152606401610709565b50613933565b613933336006546001600160a01b0388811692916201000090041684613fb3565b6001600160a01b0380881660008181526002602090815260408083208b8452825280832054938352600382528083208b845290915290205491169061ffff168115801590613984575061ffff811615155b15613a855760006139a3612710610fad61ffff851661104b8989614013565b90506001600160a01b038816613a5c576000836001600160a01b03168260405160006040518083038185875af1925050503d8060008114613a00576040519150601f19603f3d011682016040523d82523d6000602084013e613a05565b606091505b5050905080613a565760405162461bcd60e51b815260206004820152601b60248201527f726f79616c747920666565207472616e73666572206661696c656400000000006044820152606401610709565b50613a73565b613a73335b6001600160a01b038a16908584613fb3565b613a7d848261401f565b935050613baf565b50506001600160a01b0380881660009081526007602052604090206001810154905491169061ffff168115801590613ac0575061ffff811615155b15613baf576000613adf612710610fad61ffff851661104b8989614013565b90506001600160a01b038816613b98576000836001600160a01b03168260405160006040518083038185875af1925050503d8060008114613b3c576040519150601f19603f3d011682016040523d82523d6000602084013e613b41565b606091505b5050905080613b925760405162461bcd60e51b815260206004820152601b60248201527f726f79616c747920666565207472616e73666572206661696c656400000000006044820152606401610709565b50613ba1565b613ba133613a61565b613bab848261401f565b9350505b6001600160a01b038716613c635760006001600160a01b038716613bd38686614013565b604051600081818185875af1925050503d8060008114613c0f576040519150601f19603f3d011682016040523d82523d6000602084013e613c14565b606091505b5050905080613c5d5760405162461bcd60e51b81526020600482015260156024820152741bdddb995c881d1c985b9cd9995c8819985a5b1959605a1b6044820152606401610709565b50613c84565b613c843387613c728787614013565b6001600160a01b038b16929190613fb3565b6040516301ffc9a760e01b81526001600160a01b038a16906301ffc9a790613cb7906380ac58cd60e01b90600401614d07565b602060405180830381865afa158015613cd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613cf89190614d2c565b15613d7a576001600160a01b0389166342842e0e87336040516001600160e01b031960e085901b1681526001600160a01b03928316600482015291166024820152604481018b9052606401600060405180830381600087803b158015613d5d57600080fd5b505af1158015613d71573d6000803e3d6000fd5b50505050613dec565b845160408051602081018252600081529051637921219560e11b81526001600160a01b038c169263f242432a92613db9928b9233928f92600401614e56565b600060405180830381600087803b158015613dd357600080fd5b505af1158015613de7573d6000803e3d6000fd5b505050505b600860009054906101000a90046001600160a01b03166001600160a01b0316639e1f5b286040518163ffffffff1660e01b8152600401602060405180830381865afa158015613e3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e639190614d47565b85516040516309f07d3d60e21b81526001600160a01b038c81166004830152602482018c905260448201929092529116906327c1f4f490606401600060405180830381600087803b158015613eb757600080fd5b505af1158015613ecb573d6000803e3d6000fd5b50505050886001600160a01b0316613ee03390565b6001600160a01b0316876001600160a01b03167f949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d8b89600001518c613f248e611bbb565b8c51613f31908d90613fa7565b604051613f42959493929190614e90565b60405180910390a45050506001600160a01b03958616600090815260046020908152604080832097835296815286822094909716815292909552509182208281556001810180546001600160a01b031916905560028101839055600301919091555050565b60006137ca82846151f2565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b17905261400d9085906147b4565b50505050565b60006137ca8284615206565b60006137ca828461521d565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b0380841660008181526004602081815260408084208885528252808420878716855282529283902083516080810185528154815260018201549096169186019190915260028101548584015260030154606085015290516301ffc9a760e01b81526301ffc9a7916140fc916380ac58cd60e01b9101614d07565b602060405180830381865afa158015614119573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061413d9190614d2c565b156141de576040516331a9108f60e11b81526004810184905284906001600160a01b038481169190831690636352211e90602401602060405180830381865afa15801561418e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141b29190614d47565b6001600160a01b0316146141d85760405162461bcd60e51b815260040161070990614d64565b506142f7565b6040516301ffc9a760e01b81526001600160a01b038516906301ffc9a79061421190636cdb3d1360e11b90600401614d07565b602060405180830381865afa15801561422e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906142529190614d2c565b15610abc57805184906001600160a01b03821662fdd58e336040516001600160e01b031960e084901b1681526001600160a01b03909116600482015260248101889052604401602060405180830381865afa1580156142b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906142d99190614d8d565b10156141d85760405162461bcd60e51b815260040161070990614d64565b6001600160a01b03848116600081815260046020908152604080832088845282528083209487168084529482528083208381556001810180546001600160a01b0319169055600281018490556003019290925590518681529192917f9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158910160405180910390a350505050565b6000816001600160a01b0316600860009054906101000a90046001600160a01b03166001600160a01b03166318944e556040518163ffffffff1660e01b8152600401602060405180830381865afa1580156143e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906144069190614d47565b6001600160a01b031614806144f95750600860009054906101000a90046001600160a01b03166001600160a01b031663c45a01556040518163ffffffff1660e01b8152600401602060405180830381865afa158015614469573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061448d9190614d47565b604051637b51e92760e11b81526001600160a01b038481166004830152919091169063f6a3d24e90602401602060405180830381865afa1580156144d5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906144f99190614d2c565b806145e25750600860009054906101000a90046001600160a01b03166001600160a01b031663735db9396040518163ffffffff1660e01b8152600401602060405180830381865afa158015614552573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906145769190614d47565b604051637b51e92760e11b81526001600160a01b038481166004830152919091169063f6a3d24e90602401602060405180830381865afa1580156145be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906145e29190614d2c565b806146cb5750600860009054906101000a90046001600160a01b03166001600160a01b0316630fb27d8c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561463b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061465f9190614d47565b604051637b51e92760e11b81526001600160a01b038481166004830152919091169063f6a3d24e90602401602060405180830381865afa1580156146a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906146cb9190614d2c565b806137cd5750600860009054906101000a90046001600160a01b03166001600160a01b031663924db13f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614724573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906147489190614d47565b604051637b51e92760e11b81526001600160a01b038481166004830152919091169063f6a3d24e90602401602060405180830381865afa158015614790573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137cd9190614d2c565b6000614809826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661488b9092919063ffffffff16565b80519091501561488657808060200190518101906148279190614d2c565b6148865760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610709565b505050565b606061489a84846000856148a4565b90505b9392505050565b6060824710156149055760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610709565b843b6149535760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610709565b600080866001600160a01b0316858760405161496f9190615235565b60006040518083038185875af1925050503d80600081146149ac576040519150601f19603f3d011682016040523d82523d6000602084013e6149b1565b606091505b50915091506149c18282866149cc565b979650505050505050565b606083156149db57508161489d565b8251156149eb5782518084602001fd5b8160405162461bcd60e51b81526004016107099190615251565b6001600160a01b03811681146133d757600080fd5b60008060408385031215614a2d57600080fd5b8235614a3881614a05565b946020939093013593505050565b600080600060608486031215614a5b57600080fd5b8335614a6681614a05565b9250602084013591506040840135614a7d81614a05565b809150509250925092565b60008060008060008060c08789031215614aa157600080fd5b8635614aac81614a05565b955060208701359450604087013593506060870135614aca81614a05565b9598949750929560808101359460a0909101359350915050565b600060208284031215614af657600080fd5b813561489d81614a05565b60008060008060008060c08789031215614b1a57600080fd5b8635614b2581614a05565b9550602087013594506040870135614b3c81614a05565b959894975094956060810135955060808101359460a0909101359350915050565b60008060008060808587031215614b7357600080fd5b8435614b7e81614a05565b9350602085013592506040850135614b9581614a05565b91506060850135614ba581614a05565b939692955090935050565b803561ffff81168114614bc257600080fd5b919050565b600060208284031215614bd957600080fd5b6137ca82614bb0565b60008060008060808587031215614bf857600080fd5b8435614c0381614a05565b93506020850135614c1381614a05565b9250614b9560408601614bb0565b60008060008060808587031215614c3757600080fd5b8435614c4281614a05565b9350602085013592506040850135614c5981614a05565b9396929550929360600135925050565b600080600060608486031215614c7e57600080fd5b8335614c8981614a05565b925060208401359150614c9e60408501614bb0565b90509250925092565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b6020808252600f908201526e6e6f74206c6973746564206974656d60881b604082015260600190565b6001600160e01b031991909116815260200190565b80518015158114614bc257600080fd5b600060208284031215614d3e57600080fd5b6137ca82614d1c565b600060208284031215614d5957600080fd5b815161489d81614a05565b6020808252600f908201526e6e6f74206f776e696e67206974656d60881b604082015260600190565b600060208284031215614d9f57600080fd5b5051919050565b602080825260139082015272696e76616c6964206e6674206164647265737360681b604082015260600190565b60208082526011908201527034b73b30b634b2103830bc903a37b5b2b760791b604082015260600190565b60005b83811015614e19578181015183820152602001614e01565b8381111561400d5750506000910152565b60008151808452614e42816020860160208601614dfe565b601f01601f19169290920160200192915050565b6001600160a01b03868116825285166020820152604081018490526060810183905260a0608082018190526000906149c190830184614e2a565b94855260208501939093526001600160a01b039190911660408401526060830152608082015260a00190565b60008060408385031215614ecf57600080fd5b82519150602083015160ff81168114614ee757600080fd5b809150509250929050565b634e487b7160e01b600052601160045260246000fd5b600060ff821660ff841680821015614f2257614f22614ef2565b90039392505050565b80825b6001808611614f3d5750614f70565b6001600160ff1b03829004821115614f5757614f57614ef2565b80861615614f6457918102915b9490941c938002614f2e565b935093915050565b6000828015614f8e5760018114614f9857614fa1565b60019150506137cd565b829150506137cd565b5081614faf575060006137cd565b50600160008213808214614fc8578015614fe757615001565b6001600160ff1b03839004831115614fe257614fe2614ef2565b615001565b6001600160ff1b0383900583121561500157615001614ef2565b508083161561500d5750805b61501d8360011c83840283614f2b565b600082136001600160ff1b038290048311161561503c5761503c614ef2565b60008212600160ff1b8290058312161561505857615058614ef2565b029392505050565b60006137ca60ff841683614f78565b60006001600160ff1b038184138284138082168684048611161561509557615095614ef2565b600160ff1b60008712828116878305891216156150b4576150b4614ef2565b600087129250878205871284841616156150d0576150d0614ef2565b878505871281841616156150e6576150e6614ef2565b505050929093029392505050565b634e487b7160e01b600052601260045260246000fd5b600082615119576151196150f4565b600160ff1b82146000198414161561513357615133614ef2565b500590565b60008060008060008060c0878903121561515157600080fd5b865161515c81614a05565b602088015190965061516d81614a05565b8095505060408701519350606087015192506080870151915061519260a08801614d1c565b90509295509295509295565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60008160001904831182151516156151ed576151ed614ef2565b500290565b600082615201576152016150f4565b500490565b60008282101561521857615218614ef2565b500390565b6000821982111561523057615230614ef2565b500190565b60008251615247818460208701614dfe565b9190910192915050565b6020815260006137ca6020830184614e2a56fea2646970667358221220dfc4e1744bef0c7892f2f557b0d8628c0b1af1b58a7b09a9cd3d019e606d045864736f6c634300080b0033",
}

// FantomMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use FantomMarketplaceMetaData.ABI instead.
var FantomMarketplaceABI = FantomMarketplaceMetaData.ABI

// Deprecated: Use FantomMarketplaceMetaData.Sigs instead.
// FantomMarketplaceFuncSigs maps the 4-byte function signature to its string representation.
var FantomMarketplaceFuncSigs = FantomMarketplaceMetaData.Sigs

// FantomMarketplaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FantomMarketplaceMetaData.Bin instead.
var FantomMarketplaceBin = FantomMarketplaceMetaData.Bin

// DeployFantomMarketplace deploys a new Ethereum contract, binding an instance of FantomMarketplace to it.
func DeployFantomMarketplace(auth *bind.TransactOpts, backend bind.ContractBackend, _feeRecipient common.Address, _platformFee uint16) (common.Address, *types.Transaction, *FantomMarketplace, error) {
	parsed, err := FantomMarketplaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FantomMarketplaceBin), backend, _feeRecipient, _platformFee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FantomMarketplace{FantomMarketplaceCaller: FantomMarketplaceCaller{contract: contract}, FantomMarketplaceTransactor: FantomMarketplaceTransactor{contract: contract}, FantomMarketplaceFilterer: FantomMarketplaceFilterer{contract: contract}}, nil
}

// FantomMarketplace is an auto generated Go binding around an Ethereum contract.
type FantomMarketplace struct {
	FantomMarketplaceCaller     // Read-only binding to the contract
	FantomMarketplaceTransactor // Write-only binding to the contract
	FantomMarketplaceFilterer   // Log filterer for contract events
}

// FantomMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type FantomMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FantomMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FantomMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FantomMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FantomMarketplaceSession struct {
	Contract     *FantomMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FantomMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FantomMarketplaceCallerSession struct {
	Contract *FantomMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// FantomMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FantomMarketplaceTransactorSession struct {
	Contract     *FantomMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// FantomMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type FantomMarketplaceRaw struct {
	Contract *FantomMarketplace // Generic contract binding to access the raw methods on
}

// FantomMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FantomMarketplaceCallerRaw struct {
	Contract *FantomMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// FantomMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FantomMarketplaceTransactorRaw struct {
	Contract *FantomMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFantomMarketplace creates a new instance of FantomMarketplace, bound to a specific deployed contract.
func NewFantomMarketplace(address common.Address, backend bind.ContractBackend) (*FantomMarketplace, error) {
	contract, err := bindFantomMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplace{FantomMarketplaceCaller: FantomMarketplaceCaller{contract: contract}, FantomMarketplaceTransactor: FantomMarketplaceTransactor{contract: contract}, FantomMarketplaceFilterer: FantomMarketplaceFilterer{contract: contract}}, nil
}

// NewFantomMarketplaceCaller creates a new read-only instance of FantomMarketplace, bound to a specific deployed contract.
func NewFantomMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*FantomMarketplaceCaller, error) {
	contract, err := bindFantomMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceCaller{contract: contract}, nil
}

// NewFantomMarketplaceTransactor creates a new write-only instance of FantomMarketplace, bound to a specific deployed contract.
func NewFantomMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*FantomMarketplaceTransactor, error) {
	contract, err := bindFantomMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceTransactor{contract: contract}, nil
}

// NewFantomMarketplaceFilterer creates a new log filterer instance of FantomMarketplace, bound to a specific deployed contract.
func NewFantomMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*FantomMarketplaceFilterer, error) {
	contract, err := bindFantomMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceFilterer{contract: contract}, nil
}

// bindFantomMarketplace binds a generic wrapper to an already deployed contract.
func bindFantomMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FantomMarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomMarketplace *FantomMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomMarketplace.Contract.FantomMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomMarketplace *FantomMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.FantomMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomMarketplace *FantomMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.FantomMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FantomMarketplace *FantomMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FantomMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FantomMarketplace *FantomMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FantomMarketplace *FantomMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.contract.Transact(opts, method, params...)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomMarketplace *FantomMarketplaceCaller) AddressRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "addressRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomMarketplace *FantomMarketplaceSession) AddressRegistry() (common.Address, error) {
	return _FantomMarketplace.Contract.AddressRegistry(&_FantomMarketplace.CallOpts)
}

// AddressRegistry is a free data retrieval call binding the contract method 0xf3ad65f4.
//
// Solidity: function addressRegistry() view returns(address)
func (_FantomMarketplace *FantomMarketplaceCallerSession) AddressRegistry() (common.Address, error) {
	return _FantomMarketplace.Contract.AddressRegistry(&_FantomMarketplace.CallOpts)
}

// CollectionRoyalties is a free data retrieval call binding the contract method 0x57b275b4.
//
// Solidity: function collectionRoyalties(address ) view returns(uint16 royalty, address creator, address feeRecipient)
func (_FantomMarketplace *FantomMarketplaceCaller) CollectionRoyalties(opts *bind.CallOpts, arg0 common.Address) (struct {
	Royalty      uint16
	Creator      common.Address
	FeeRecipient common.Address
}, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "collectionRoyalties", arg0)

	outstruct := new(struct {
		Royalty      uint16
		Creator      common.Address
		FeeRecipient common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Royalty = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FeeRecipient = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CollectionRoyalties is a free data retrieval call binding the contract method 0x57b275b4.
//
// Solidity: function collectionRoyalties(address ) view returns(uint16 royalty, address creator, address feeRecipient)
func (_FantomMarketplace *FantomMarketplaceSession) CollectionRoyalties(arg0 common.Address) (struct {
	Royalty      uint16
	Creator      common.Address
	FeeRecipient common.Address
}, error) {
	return _FantomMarketplace.Contract.CollectionRoyalties(&_FantomMarketplace.CallOpts, arg0)
}

// CollectionRoyalties is a free data retrieval call binding the contract method 0x57b275b4.
//
// Solidity: function collectionRoyalties(address ) view returns(uint16 royalty, address creator, address feeRecipient)
func (_FantomMarketplace *FantomMarketplaceCallerSession) CollectionRoyalties(arg0 common.Address) (struct {
	Royalty      uint16
	Creator      common.Address
	FeeRecipient common.Address
}, error) {
	return _FantomMarketplace.Contract.CollectionRoyalties(&_FantomMarketplace.CallOpts, arg0)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomMarketplace *FantomMarketplaceCaller) FeeReceipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "feeReceipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomMarketplace *FantomMarketplaceSession) FeeReceipient() (common.Address, error) {
	return _FantomMarketplace.Contract.FeeReceipient(&_FantomMarketplace.CallOpts)
}

// FeeReceipient is a free data retrieval call binding the contract method 0x3740ebb3.
//
// Solidity: function feeReceipient() view returns(address)
func (_FantomMarketplace *FantomMarketplaceCallerSession) FeeReceipient() (common.Address, error) {
	return _FantomMarketplace.Contract.FeeReceipient(&_FantomMarketplace.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _payToken) view returns(int256)
func (_FantomMarketplace *FantomMarketplaceCaller) GetPrice(opts *bind.CallOpts, _payToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "getPrice", _payToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _payToken) view returns(int256)
func (_FantomMarketplace *FantomMarketplaceSession) GetPrice(_payToken common.Address) (*big.Int, error) {
	return _FantomMarketplace.Contract.GetPrice(&_FantomMarketplace.CallOpts, _payToken)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address _payToken) view returns(int256)
func (_FantomMarketplace *FantomMarketplaceCallerSession) GetPrice(_payToken common.Address) (*big.Int, error) {
	return _FantomMarketplace.Contract.GetPrice(&_FantomMarketplace.CallOpts, _payToken)
}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address , uint256 , address ) view returns(uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplace *FantomMarketplaceCaller) Listings(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
}, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "listings", arg0, arg1, arg2)

	outstruct := new(struct {
		Quantity     *big.Int
		PayToken     common.Address
		PricePerItem *big.Int
		StartingTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Quantity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PayToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PricePerItem = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartingTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address , uint256 , address ) view returns(uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplace *FantomMarketplaceSession) Listings(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
}, error) {
	return _FantomMarketplace.Contract.Listings(&_FantomMarketplace.CallOpts, arg0, arg1, arg2)
}

// Listings is a free data retrieval call binding the contract method 0x6bd3a64b.
//
// Solidity: function listings(address , uint256 , address ) view returns(uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplace *FantomMarketplaceCallerSession) Listings(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
}, error) {
	return _FantomMarketplace.Contract.Listings(&_FantomMarketplace.CallOpts, arg0, arg1, arg2)
}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(address)
func (_FantomMarketplace *FantomMarketplaceCaller) Minters(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "minters", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(address)
func (_FantomMarketplace *FantomMarketplaceSession) Minters(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _FantomMarketplace.Contract.Minters(&_FantomMarketplace.CallOpts, arg0, arg1)
}

// Minters is a free data retrieval call binding the contract method 0x78746d0d.
//
// Solidity: function minters(address , uint256 ) view returns(address)
func (_FantomMarketplace *FantomMarketplaceCallerSession) Minters(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _FantomMarketplace.Contract.Minters(&_FantomMarketplace.CallOpts, arg0, arg1)
}

// Offers is a free data retrieval call binding the contract method 0xd3f494cc.
//
// Solidity: function offers(address , uint256 , address ) view returns(address payToken, uint256 quantity, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplace *FantomMarketplaceCaller) Offers(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	PayToken     common.Address
	Quantity     *big.Int
	PricePerItem *big.Int
	Deadline     *big.Int
}, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "offers", arg0, arg1, arg2)

	outstruct := new(struct {
		PayToken     common.Address
		Quantity     *big.Int
		PricePerItem *big.Int
		Deadline     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PayToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Quantity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PricePerItem = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Offers is a free data retrieval call binding the contract method 0xd3f494cc.
//
// Solidity: function offers(address , uint256 , address ) view returns(address payToken, uint256 quantity, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplace *FantomMarketplaceSession) Offers(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	PayToken     common.Address
	Quantity     *big.Int
	PricePerItem *big.Int
	Deadline     *big.Int
}, error) {
	return _FantomMarketplace.Contract.Offers(&_FantomMarketplace.CallOpts, arg0, arg1, arg2)
}

// Offers is a free data retrieval call binding the contract method 0xd3f494cc.
//
// Solidity: function offers(address , uint256 , address ) view returns(address payToken, uint256 quantity, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplace *FantomMarketplaceCallerSession) Offers(arg0 common.Address, arg1 *big.Int, arg2 common.Address) (struct {
	PayToken     common.Address
	Quantity     *big.Int
	PricePerItem *big.Int
	Deadline     *big.Int
}, error) {
	return _FantomMarketplace.Contract.Offers(&_FantomMarketplace.CallOpts, arg0, arg1, arg2)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomMarketplace *FantomMarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomMarketplace *FantomMarketplaceSession) Owner() (common.Address, error) {
	return _FantomMarketplace.Contract.Owner(&_FantomMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FantomMarketplace *FantomMarketplaceCallerSession) Owner() (common.Address, error) {
	return _FantomMarketplace.Contract.Owner(&_FantomMarketplace.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint16)
func (_FantomMarketplace *FantomMarketplaceCaller) PlatformFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint16)
func (_FantomMarketplace *FantomMarketplaceSession) PlatformFee() (uint16, error) {
	return _FantomMarketplace.Contract.PlatformFee(&_FantomMarketplace.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint16)
func (_FantomMarketplace *FantomMarketplaceCallerSession) PlatformFee() (uint16, error) {
	return _FantomMarketplace.Contract.PlatformFee(&_FantomMarketplace.CallOpts)
}

// Royalties is a free data retrieval call binding the contract method 0xe1e549c4.
//
// Solidity: function royalties(address , uint256 ) view returns(uint16)
func (_FantomMarketplace *FantomMarketplaceCaller) Royalties(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (uint16, error) {
	var out []interface{}
	err := _FantomMarketplace.contract.Call(opts, &out, "royalties", arg0, arg1)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Royalties is a free data retrieval call binding the contract method 0xe1e549c4.
//
// Solidity: function royalties(address , uint256 ) view returns(uint16)
func (_FantomMarketplace *FantomMarketplaceSession) Royalties(arg0 common.Address, arg1 *big.Int) (uint16, error) {
	return _FantomMarketplace.Contract.Royalties(&_FantomMarketplace.CallOpts, arg0, arg1)
}

// Royalties is a free data retrieval call binding the contract method 0xe1e549c4.
//
// Solidity: function royalties(address , uint256 ) view returns(uint16)
func (_FantomMarketplace *FantomMarketplaceCallerSession) Royalties(arg0 common.Address, arg1 *big.Int) (uint16, error) {
	return _FantomMarketplace.Contract.Royalties(&_FantomMarketplace.CallOpts, arg0, arg1)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x3bbb2806.
//
// Solidity: function acceptOffer(address _nftAddress, uint256 _tokenId, address _creator) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) AcceptOffer(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "acceptOffer", _nftAddress, _tokenId, _creator)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x3bbb2806.
//
// Solidity: function acceptOffer(address _nftAddress, uint256 _tokenId, address _creator) returns()
func (_FantomMarketplace *FantomMarketplaceSession) AcceptOffer(_nftAddress common.Address, _tokenId *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.AcceptOffer(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _creator)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x3bbb2806.
//
// Solidity: function acceptOffer(address _nftAddress, uint256 _tokenId, address _creator) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) AcceptOffer(_nftAddress common.Address, _tokenId *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.AcceptOffer(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _creator)
}

// BuyItem is a paid mutator transaction binding the contract method 0x259ca365.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _owner) payable returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) BuyItem(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "buyItem", _nftAddress, _tokenId, _owner)
}

// BuyItem is a paid mutator transaction binding the contract method 0x259ca365.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _owner) payable returns()
func (_FantomMarketplace *FantomMarketplaceSession) BuyItem(_nftAddress common.Address, _tokenId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.BuyItem(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _owner)
}

// BuyItem is a paid mutator transaction binding the contract method 0x259ca365.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _owner) payable returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) BuyItem(_nftAddress common.Address, _tokenId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.BuyItem(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _owner)
}

// BuyItem0 is a paid mutator transaction binding the contract method 0x85f9d345.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _payToken, address _owner) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) BuyItem0(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "buyItem0", _nftAddress, _tokenId, _payToken, _owner)
}

// BuyItem0 is a paid mutator transaction binding the contract method 0x85f9d345.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _payToken, address _owner) returns()
func (_FantomMarketplace *FantomMarketplaceSession) BuyItem0(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.BuyItem0(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _owner)
}

// BuyItem0 is a paid mutator transaction binding the contract method 0x85f9d345.
//
// Solidity: function buyItem(address _nftAddress, uint256 _tokenId, address _payToken, address _owner) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) BuyItem0(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _owner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.BuyItem0(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _owner)
}

// CancelListing is a paid mutator transaction binding the contract method 0xb2ddee06.
//
// Solidity: function cancelListing(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) CancelListing(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "cancelListing", _nftAddress, _tokenId)
}

// CancelListing is a paid mutator transaction binding the contract method 0xb2ddee06.
//
// Solidity: function cancelListing(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplace *FantomMarketplaceSession) CancelListing(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.CancelListing(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId)
}

// CancelListing is a paid mutator transaction binding the contract method 0xb2ddee06.
//
// Solidity: function cancelListing(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) CancelListing(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.CancelListing(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) CancelOffer(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "cancelOffer", _nftAddress, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplace *FantomMarketplaceSession) CancelOffer(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.CancelOffer(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId)
}

// CancelOffer is a paid mutator transaction binding the contract method 0x058a56ac.
//
// Solidity: function cancelOffer(address _nftAddress, uint256 _tokenId) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) CancelOffer(_nftAddress common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.CancelOffer(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x42cbb4b8.
//
// Solidity: function createOffer(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _quantity, uint256 _pricePerItem, uint256 _deadline) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) CreateOffer(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _quantity *big.Int, _pricePerItem *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "createOffer", _nftAddress, _tokenId, _payToken, _quantity, _pricePerItem, _deadline)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x42cbb4b8.
//
// Solidity: function createOffer(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _quantity, uint256 _pricePerItem, uint256 _deadline) returns()
func (_FantomMarketplace *FantomMarketplaceSession) CreateOffer(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _quantity *big.Int, _pricePerItem *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.CreateOffer(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _quantity, _pricePerItem, _deadline)
}

// CreateOffer is a paid mutator transaction binding the contract method 0x42cbb4b8.
//
// Solidity: function createOffer(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _quantity, uint256 _pricePerItem, uint256 _deadline) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) CreateOffer(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _quantity *big.Int, _pricePerItem *big.Int, _deadline *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.CreateOffer(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _quantity, _pricePerItem, _deadline)
}

// ListItem is a paid mutator transaction binding the contract method 0x3fc1cc26.
//
// Solidity: function listItem(address _nftAddress, uint256 _tokenId, uint256 _quantity, address _payToken, uint256 _pricePerItem, uint256 _startingTime) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) ListItem(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int, _payToken common.Address, _pricePerItem *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "listItem", _nftAddress, _tokenId, _quantity, _payToken, _pricePerItem, _startingTime)
}

// ListItem is a paid mutator transaction binding the contract method 0x3fc1cc26.
//
// Solidity: function listItem(address _nftAddress, uint256 _tokenId, uint256 _quantity, address _payToken, uint256 _pricePerItem, uint256 _startingTime) returns()
func (_FantomMarketplace *FantomMarketplaceSession) ListItem(_nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int, _payToken common.Address, _pricePerItem *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.ListItem(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _quantity, _payToken, _pricePerItem, _startingTime)
}

// ListItem is a paid mutator transaction binding the contract method 0x3fc1cc26.
//
// Solidity: function listItem(address _nftAddress, uint256 _tokenId, uint256 _quantity, address _payToken, uint256 _pricePerItem, uint256 _startingTime) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) ListItem(_nftAddress common.Address, _tokenId *big.Int, _quantity *big.Int, _payToken common.Address, _pricePerItem *big.Int, _startingTime *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.ListItem(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _quantity, _payToken, _pricePerItem, _startingTime)
}

// RegisterCollectionRoyalty is a paid mutator transaction binding the contract method 0xc1661a9a.
//
// Solidity: function registerCollectionRoyalty(address _nftAddress, address _creator, uint16 _royalty, address _feeRecipient) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) RegisterCollectionRoyalty(opts *bind.TransactOpts, _nftAddress common.Address, _creator common.Address, _royalty uint16, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "registerCollectionRoyalty", _nftAddress, _creator, _royalty, _feeRecipient)
}

// RegisterCollectionRoyalty is a paid mutator transaction binding the contract method 0xc1661a9a.
//
// Solidity: function registerCollectionRoyalty(address _nftAddress, address _creator, uint16 _royalty, address _feeRecipient) returns()
func (_FantomMarketplace *FantomMarketplaceSession) RegisterCollectionRoyalty(_nftAddress common.Address, _creator common.Address, _royalty uint16, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.RegisterCollectionRoyalty(&_FantomMarketplace.TransactOpts, _nftAddress, _creator, _royalty, _feeRecipient)
}

// RegisterCollectionRoyalty is a paid mutator transaction binding the contract method 0xc1661a9a.
//
// Solidity: function registerCollectionRoyalty(address _nftAddress, address _creator, uint16 _royalty, address _feeRecipient) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) RegisterCollectionRoyalty(_nftAddress common.Address, _creator common.Address, _royalty uint16, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.RegisterCollectionRoyalty(&_FantomMarketplace.TransactOpts, _nftAddress, _creator, _royalty, _feeRecipient)
}

// RegisterRoyalty is a paid mutator transaction binding the contract method 0xf3880b6e.
//
// Solidity: function registerRoyalty(address _nftAddress, uint256 _tokenId, uint16 _royalty) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) RegisterRoyalty(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _royalty uint16) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "registerRoyalty", _nftAddress, _tokenId, _royalty)
}

// RegisterRoyalty is a paid mutator transaction binding the contract method 0xf3880b6e.
//
// Solidity: function registerRoyalty(address _nftAddress, uint256 _tokenId, uint16 _royalty) returns()
func (_FantomMarketplace *FantomMarketplaceSession) RegisterRoyalty(_nftAddress common.Address, _tokenId *big.Int, _royalty uint16) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.RegisterRoyalty(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _royalty)
}

// RegisterRoyalty is a paid mutator transaction binding the contract method 0xf3880b6e.
//
// Solidity: function registerRoyalty(address _nftAddress, uint256 _tokenId, uint16 _royalty) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) RegisterRoyalty(_nftAddress common.Address, _tokenId *big.Int, _royalty uint16) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.RegisterRoyalty(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _royalty)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomMarketplace *FantomMarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomMarketplace.Contract.RenounceOwnership(&_FantomMarketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FantomMarketplace.Contract.RenounceOwnership(&_FantomMarketplace.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomMarketplace *FantomMarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.TransferOwnership(&_FantomMarketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.TransferOwnership(&_FantomMarketplace.TransactOpts, newOwner)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) UpdateAddressRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "updateAddressRegistry", _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomMarketplace *FantomMarketplaceSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdateAddressRegistry(&_FantomMarketplace.TransactOpts, _registry)
}

// UpdateAddressRegistry is a paid mutator transaction binding the contract method 0x68e79e89.
//
// Solidity: function updateAddressRegistry(address _registry) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) UpdateAddressRegistry(_registry common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdateAddressRegistry(&_FantomMarketplace.TransactOpts, _registry)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xe984f2eb.
//
// Solidity: function updateListing(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _newPrice) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) UpdateListing(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "updateListing", _nftAddress, _tokenId, _payToken, _newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xe984f2eb.
//
// Solidity: function updateListing(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _newPrice) returns()
func (_FantomMarketplace *FantomMarketplaceSession) UpdateListing(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdateListing(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _newPrice)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xe984f2eb.
//
// Solidity: function updateListing(address _nftAddress, uint256 _tokenId, address _payToken, uint256 _newPrice) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) UpdateListing(_nftAddress common.Address, _tokenId *big.Int, _payToken common.Address, _newPrice *big.Int) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdateListing(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _payToken, _newPrice)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 _platformFee) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) UpdatePlatformFee(opts *bind.TransactOpts, _platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "updatePlatformFee", _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 _platformFee) returns()
func (_FantomMarketplace *FantomMarketplaceSession) UpdatePlatformFee(_platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdatePlatformFee(&_FantomMarketplace.TransactOpts, _platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 _platformFee) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) UpdatePlatformFee(_platformFee uint16) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdatePlatformFee(&_FantomMarketplace.TransactOpts, _platformFee)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) UpdatePlatformFeeRecipient(opts *bind.TransactOpts, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "updatePlatformFeeRecipient", _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomMarketplace *FantomMarketplaceSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdatePlatformFeeRecipient(&_FantomMarketplace.TransactOpts, _platformFeeRecipient)
}

// UpdatePlatformFeeRecipient is a paid mutator transaction binding the contract method 0xf5fe7f71.
//
// Solidity: function updatePlatformFeeRecipient(address _platformFeeRecipient) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) UpdatePlatformFeeRecipient(_platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.UpdatePlatformFeeRecipient(&_FantomMarketplace.TransactOpts, _platformFeeRecipient)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0xe940ebeb.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, address _seller, address _buyer) returns()
func (_FantomMarketplace *FantomMarketplaceTransactor) ValidateItemSold(opts *bind.TransactOpts, _nftAddress common.Address, _tokenId *big.Int, _seller common.Address, _buyer common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.contract.Transact(opts, "validateItemSold", _nftAddress, _tokenId, _seller, _buyer)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0xe940ebeb.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, address _seller, address _buyer) returns()
func (_FantomMarketplace *FantomMarketplaceSession) ValidateItemSold(_nftAddress common.Address, _tokenId *big.Int, _seller common.Address, _buyer common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.ValidateItemSold(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _seller, _buyer)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0xe940ebeb.
//
// Solidity: function validateItemSold(address _nftAddress, uint256 _tokenId, address _seller, address _buyer) returns()
func (_FantomMarketplace *FantomMarketplaceTransactorSession) ValidateItemSold(_nftAddress common.Address, _tokenId *big.Int, _seller common.Address, _buyer common.Address) (*types.Transaction, error) {
	return _FantomMarketplace.Contract.ValidateItemSold(&_FantomMarketplace.TransactOpts, _nftAddress, _tokenId, _seller, _buyer)
}

// FantomMarketplaceItemCanceledIterator is returned from FilterItemCanceled and is used to iterate over the raw logs and unpacked data for ItemCanceled events raised by the FantomMarketplace contract.
type FantomMarketplaceItemCanceledIterator struct {
	Event *FantomMarketplaceItemCanceled // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceItemCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceItemCanceled)
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
		it.Event = new(FantomMarketplaceItemCanceled)
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
func (it *FantomMarketplaceItemCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceItemCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceItemCanceled represents a ItemCanceled event raised by the FantomMarketplace contract.
type FantomMarketplaceItemCanceled struct {
	Owner   common.Address
	Nft     common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterItemCanceled is a free log retrieval operation binding the contract event 0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158.
//
// Solidity: event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterItemCanceled(opts *bind.FilterOpts, owner []common.Address, nft []common.Address) (*FantomMarketplaceItemCanceledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "ItemCanceled", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceItemCanceledIterator{contract: _FantomMarketplace.contract, event: "ItemCanceled", logs: logs, sub: sub}, nil
}

// WatchItemCanceled is a free log subscription operation binding the contract event 0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158.
//
// Solidity: event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchItemCanceled(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceItemCanceled, owner []common.Address, nft []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "ItemCanceled", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceItemCanceled)
				if err := _FantomMarketplace.contract.UnpackLog(event, "ItemCanceled", log); err != nil {
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

// ParseItemCanceled is a log parse operation binding the contract event 0x9ba1a3cb55ce8d63d072a886f94d2a744f50cddf82128e897d0661f5ec623158.
//
// Solidity: event ItemCanceled(address indexed owner, address indexed nft, uint256 tokenId)
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseItemCanceled(log types.Log) (*FantomMarketplaceItemCanceled, error) {
	event := new(FantomMarketplaceItemCanceled)
	if err := _FantomMarketplace.contract.UnpackLog(event, "ItemCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceItemListedIterator is returned from FilterItemListed and is used to iterate over the raw logs and unpacked data for ItemListed events raised by the FantomMarketplace contract.
type FantomMarketplaceItemListedIterator struct {
	Event *FantomMarketplaceItemListed // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceItemListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceItemListed)
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
		it.Event = new(FantomMarketplaceItemListed)
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
func (it *FantomMarketplaceItemListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceItemListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceItemListed represents a ItemListed event raised by the FantomMarketplace contract.
type FantomMarketplaceItemListed struct {
	Owner        common.Address
	Nft          common.Address
	TokenId      *big.Int
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	StartingTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterItemListed is a free log retrieval operation binding the contract event 0xa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa2629.
//
// Solidity: event ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterItemListed(opts *bind.FilterOpts, owner []common.Address, nft []common.Address) (*FantomMarketplaceItemListedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "ItemListed", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceItemListedIterator{contract: _FantomMarketplace.contract, event: "ItemListed", logs: logs, sub: sub}, nil
}

// WatchItemListed is a free log subscription operation binding the contract event 0xa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa2629.
//
// Solidity: event ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchItemListed(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceItemListed, owner []common.Address, nft []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "ItemListed", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceItemListed)
				if err := _FantomMarketplace.contract.UnpackLog(event, "ItemListed", log); err != nil {
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

// ParseItemListed is a log parse operation binding the contract event 0xa0294f02f8ad82fe4744717b0f953a105547196cd3c67056200c1a4ae3aa2629.
//
// Solidity: event ItemListed(address indexed owner, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 startingTime)
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseItemListed(log types.Log) (*FantomMarketplaceItemListed, error) {
	event := new(FantomMarketplaceItemListed)
	if err := _FantomMarketplace.contract.UnpackLog(event, "ItemListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceItemSoldIterator is returned from FilterItemSold and is used to iterate over the raw logs and unpacked data for ItemSold events raised by the FantomMarketplace contract.
type FantomMarketplaceItemSoldIterator struct {
	Event *FantomMarketplaceItemSold // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceItemSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceItemSold)
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
		it.Event = new(FantomMarketplaceItemSold)
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
func (it *FantomMarketplaceItemSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceItemSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceItemSold represents a ItemSold event raised by the FantomMarketplace contract.
type FantomMarketplaceItemSold struct {
	Seller       common.Address
	Buyer        common.Address
	Nft          common.Address
	TokenId      *big.Int
	Quantity     *big.Int
	PayToken     common.Address
	UnitPrice    *big.Int
	PricePerItem *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterItemSold is a free log retrieval operation binding the contract event 0x949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 unitPrice, uint256 pricePerItem)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterItemSold(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, nft []common.Address) (*FantomMarketplaceItemSoldIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "ItemSold", sellerRule, buyerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceItemSoldIterator{contract: _FantomMarketplace.contract, event: "ItemSold", logs: logs, sub: sub}, nil
}

// WatchItemSold is a free log subscription operation binding the contract event 0x949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 unitPrice, uint256 pricePerItem)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchItemSold(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceItemSold, seller []common.Address, buyer []common.Address, nft []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "ItemSold", sellerRule, buyerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceItemSold)
				if err := _FantomMarketplace.contract.UnpackLog(event, "ItemSold", log); err != nil {
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

// ParseItemSold is a log parse operation binding the contract event 0x949d1413baca5c0e4ab96b0198d536cac8cdcc17cb909b9ea24594f42ed9fa0d.
//
// Solidity: event ItemSold(address indexed seller, address indexed buyer, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, int256 unitPrice, uint256 pricePerItem)
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseItemSold(log types.Log) (*FantomMarketplaceItemSold, error) {
	event := new(FantomMarketplaceItemSold)
	if err := _FantomMarketplace.contract.UnpackLog(event, "ItemSold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceItemUpdatedIterator is returned from FilterItemUpdated and is used to iterate over the raw logs and unpacked data for ItemUpdated events raised by the FantomMarketplace contract.
type FantomMarketplaceItemUpdatedIterator struct {
	Event *FantomMarketplaceItemUpdated // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceItemUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceItemUpdated)
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
		it.Event = new(FantomMarketplaceItemUpdated)
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
func (it *FantomMarketplaceItemUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceItemUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceItemUpdated represents a ItemUpdated event raised by the FantomMarketplace contract.
type FantomMarketplaceItemUpdated struct {
	Owner    common.Address
	Nft      common.Address
	TokenId  *big.Int
	PayToken common.Address
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterItemUpdated is a free log retrieval operation binding the contract event 0x60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd.
//
// Solidity: event ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterItemUpdated(opts *bind.FilterOpts, owner []common.Address, nft []common.Address) (*FantomMarketplaceItemUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "ItemUpdated", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceItemUpdatedIterator{contract: _FantomMarketplace.contract, event: "ItemUpdated", logs: logs, sub: sub}, nil
}

// WatchItemUpdated is a free log subscription operation binding the contract event 0x60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd.
//
// Solidity: event ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchItemUpdated(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceItemUpdated, owner []common.Address, nft []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "ItemUpdated", ownerRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceItemUpdated)
				if err := _FantomMarketplace.contract.UnpackLog(event, "ItemUpdated", log); err != nil {
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

// ParseItemUpdated is a log parse operation binding the contract event 0x60a11f1619b1716bc2857bf610d4bc631336e14d197025fd5875c1aca1ac7cbd.
//
// Solidity: event ItemUpdated(address indexed owner, address indexed nft, uint256 tokenId, address payToken, uint256 newPrice)
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseItemUpdated(log types.Log) (*FantomMarketplaceItemUpdated, error) {
	event := new(FantomMarketplaceItemUpdated)
	if err := _FantomMarketplace.contract.UnpackLog(event, "ItemUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceOfferCanceledIterator is returned from FilterOfferCanceled and is used to iterate over the raw logs and unpacked data for OfferCanceled events raised by the FantomMarketplace contract.
type FantomMarketplaceOfferCanceledIterator struct {
	Event *FantomMarketplaceOfferCanceled // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceOfferCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceOfferCanceled)
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
		it.Event = new(FantomMarketplaceOfferCanceled)
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
func (it *FantomMarketplaceOfferCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceOfferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceOfferCanceled represents a OfferCanceled event raised by the FantomMarketplace contract.
type FantomMarketplaceOfferCanceled struct {
	Creator common.Address
	Nft     common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOfferCanceled is a free log retrieval operation binding the contract event 0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7.
//
// Solidity: event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterOfferCanceled(opts *bind.FilterOpts, creator []common.Address, nft []common.Address) (*FantomMarketplaceOfferCanceledIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "OfferCanceled", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceOfferCanceledIterator{contract: _FantomMarketplace.contract, event: "OfferCanceled", logs: logs, sub: sub}, nil
}

// WatchOfferCanceled is a free log subscription operation binding the contract event 0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7.
//
// Solidity: event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchOfferCanceled(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceOfferCanceled, creator []common.Address, nft []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "OfferCanceled", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceOfferCanceled)
				if err := _FantomMarketplace.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
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

// ParseOfferCanceled is a log parse operation binding the contract event 0xc6e24dcedb16cc237925b586889d0a38102c719734d6cc56acb89b013099b3a7.
//
// Solidity: event OfferCanceled(address indexed creator, address indexed nft, uint256 tokenId)
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseOfferCanceled(log types.Log) (*FantomMarketplaceOfferCanceled, error) {
	event := new(FantomMarketplaceOfferCanceled)
	if err := _FantomMarketplace.contract.UnpackLog(event, "OfferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceOfferCreatedIterator is returned from FilterOfferCreated and is used to iterate over the raw logs and unpacked data for OfferCreated events raised by the FantomMarketplace contract.
type FantomMarketplaceOfferCreatedIterator struct {
	Event *FantomMarketplaceOfferCreated // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceOfferCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceOfferCreated)
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
		it.Event = new(FantomMarketplaceOfferCreated)
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
func (it *FantomMarketplaceOfferCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceOfferCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceOfferCreated represents a OfferCreated event raised by the FantomMarketplace contract.
type FantomMarketplaceOfferCreated struct {
	Creator      common.Address
	Nft          common.Address
	TokenId      *big.Int
	Quantity     *big.Int
	PayToken     common.Address
	PricePerItem *big.Int
	Deadline     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOfferCreated is a free log retrieval operation binding the contract event 0x89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f6.
//
// Solidity: event OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterOfferCreated(opts *bind.FilterOpts, creator []common.Address, nft []common.Address) (*FantomMarketplaceOfferCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "OfferCreated", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceOfferCreatedIterator{contract: _FantomMarketplace.contract, event: "OfferCreated", logs: logs, sub: sub}, nil
}

// WatchOfferCreated is a free log subscription operation binding the contract event 0x89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f6.
//
// Solidity: event OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchOfferCreated(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceOfferCreated, creator []common.Address, nft []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "OfferCreated", creatorRule, nftRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceOfferCreated)
				if err := _FantomMarketplace.contract.UnpackLog(event, "OfferCreated", log); err != nil {
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

// ParseOfferCreated is a log parse operation binding the contract event 0x89f255157c655b5155655107b77c620998e5ad4e7485d749e4e6d7ddb63e70f6.
//
// Solidity: event OfferCreated(address indexed creator, address indexed nft, uint256 tokenId, uint256 quantity, address payToken, uint256 pricePerItem, uint256 deadline)
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseOfferCreated(log types.Log) (*FantomMarketplaceOfferCreated, error) {
	event := new(FantomMarketplaceOfferCreated)
	if err := _FantomMarketplace.contract.UnpackLog(event, "OfferCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FantomMarketplace contract.
type FantomMarketplaceOwnershipTransferredIterator struct {
	Event *FantomMarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceOwnershipTransferred)
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
		it.Event = new(FantomMarketplaceOwnershipTransferred)
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
func (it *FantomMarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the FantomMarketplace contract.
type FantomMarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FantomMarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceOwnershipTransferredIterator{contract: _FantomMarketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceOwnershipTransferred)
				if err := _FantomMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*FantomMarketplaceOwnershipTransferred, error) {
	event := new(FantomMarketplaceOwnershipTransferred)
	if err := _FantomMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceUpdatePlatformFeeIterator is returned from FilterUpdatePlatformFee and is used to iterate over the raw logs and unpacked data for UpdatePlatformFee events raised by the FantomMarketplace contract.
type FantomMarketplaceUpdatePlatformFeeIterator struct {
	Event *FantomMarketplaceUpdatePlatformFee // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceUpdatePlatformFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceUpdatePlatformFee)
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
		it.Event = new(FantomMarketplaceUpdatePlatformFee)
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
func (it *FantomMarketplaceUpdatePlatformFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceUpdatePlatformFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceUpdatePlatformFee represents a UpdatePlatformFee event raised by the FantomMarketplace contract.
type FantomMarketplaceUpdatePlatformFee struct {
	PlatformFee uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFee is a free log retrieval operation binding the contract event 0xcabead39256298cdf7d1228eee68b47bc54a2192283d8d93574da08d2d7ac303.
//
// Solidity: event UpdatePlatformFee(uint16 platformFee)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterUpdatePlatformFee(opts *bind.FilterOpts) (*FantomMarketplaceUpdatePlatformFeeIterator, error) {

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceUpdatePlatformFeeIterator{contract: _FantomMarketplace.contract, event: "UpdatePlatformFee", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFee is a free log subscription operation binding the contract event 0xcabead39256298cdf7d1228eee68b47bc54a2192283d8d93574da08d2d7ac303.
//
// Solidity: event UpdatePlatformFee(uint16 platformFee)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchUpdatePlatformFee(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceUpdatePlatformFee) (event.Subscription, error) {

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "UpdatePlatformFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceUpdatePlatformFee)
				if err := _FantomMarketplace.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
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

// ParseUpdatePlatformFee is a log parse operation binding the contract event 0xcabead39256298cdf7d1228eee68b47bc54a2192283d8d93574da08d2d7ac303.
//
// Solidity: event UpdatePlatformFee(uint16 platformFee)
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseUpdatePlatformFee(log types.Log) (*FantomMarketplaceUpdatePlatformFee, error) {
	event := new(FantomMarketplaceUpdatePlatformFee)
	if err := _FantomMarketplace.contract.UnpackLog(event, "UpdatePlatformFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FantomMarketplaceUpdatePlatformFeeRecipientIterator is returned from FilterUpdatePlatformFeeRecipient and is used to iterate over the raw logs and unpacked data for UpdatePlatformFeeRecipient events raised by the FantomMarketplace contract.
type FantomMarketplaceUpdatePlatformFeeRecipientIterator struct {
	Event *FantomMarketplaceUpdatePlatformFeeRecipient // Event containing the contract specifics and raw log

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
func (it *FantomMarketplaceUpdatePlatformFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FantomMarketplaceUpdatePlatformFeeRecipient)
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
		it.Event = new(FantomMarketplaceUpdatePlatformFeeRecipient)
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
func (it *FantomMarketplaceUpdatePlatformFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FantomMarketplaceUpdatePlatformFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FantomMarketplaceUpdatePlatformFeeRecipient represents a UpdatePlatformFeeRecipient event raised by the FantomMarketplace contract.
type FantomMarketplaceUpdatePlatformFeeRecipient struct {
	PlatformFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdatePlatformFeeRecipient is a free log retrieval operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomMarketplace *FantomMarketplaceFilterer) FilterUpdatePlatformFeeRecipient(opts *bind.FilterOpts) (*FantomMarketplaceUpdatePlatformFeeRecipientIterator, error) {

	logs, sub, err := _FantomMarketplace.contract.FilterLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &FantomMarketplaceUpdatePlatformFeeRecipientIterator{contract: _FantomMarketplace.contract, event: "UpdatePlatformFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdatePlatformFeeRecipient is a free log subscription operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomMarketplace *FantomMarketplaceFilterer) WatchUpdatePlatformFeeRecipient(opts *bind.WatchOpts, sink chan<- *FantomMarketplaceUpdatePlatformFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _FantomMarketplace.contract.WatchLogs(opts, "UpdatePlatformFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FantomMarketplaceUpdatePlatformFeeRecipient)
				if err := _FantomMarketplace.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
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

// ParseUpdatePlatformFeeRecipient is a log parse operation binding the contract event 0xe57e7c1f36cc83fade34e32351e6eee7eb9da532662b1b5da10c631e8222aca7.
//
// Solidity: event UpdatePlatformFeeRecipient(address platformFeeRecipient)
func (_FantomMarketplace *FantomMarketplaceFilterer) ParseUpdatePlatformFeeRecipient(log types.Log) (*FantomMarketplaceUpdatePlatformFeeRecipient, error) {
	event := new(FantomMarketplaceUpdatePlatformFeeRecipient)
	if err := _FantomMarketplace.contract.UnpackLog(event, "UpdatePlatformFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155MetaData contains all meta data concerning the IERC1155 contract.
var IERC1155MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"00fdd58e": "balanceOf(address,uint256)",
		"4e1273f4": "balanceOfBatch(address[],uint256[])",
		"e985e9c5": "isApprovedForAll(address,address)",
		"2eb2c2d6": "safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)",
		"f242432a": "safeTransferFrom(address,address,uint256,uint256,bytes)",
		"a22cb465": "setApprovalForAll(address,bool)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// IERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1155MetaData.ABI instead.
var IERC1155ABI = IERC1155MetaData.ABI

// Deprecated: Use IERC1155MetaData.Sigs instead.
// IERC1155FuncSigs maps the 4-byte function signature to its string representation.
var IERC1155FuncSigs = IERC1155MetaData.Sigs

// IERC1155 is an auto generated Go binding around an Ethereum contract.
type IERC1155 struct {
	IERC1155Caller     // Read-only binding to the contract
	IERC1155Transactor // Write-only binding to the contract
	IERC1155Filterer   // Log filterer for contract events
}

// IERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1155Session struct {
	Contract     *IERC1155         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1155CallerSession struct {
	Contract *IERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1155TransactorSession struct {
	Contract     *IERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1155Raw struct {
	Contract *IERC1155 // Generic contract binding to access the raw methods on
}

// IERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1155CallerRaw struct {
	Contract *IERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// IERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1155TransactorRaw struct {
	Contract *IERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1155 creates a new instance of IERC1155, bound to a specific deployed contract.
func NewIERC1155(address common.Address, backend bind.ContractBackend) (*IERC1155, error) {
	contract, err := bindIERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1155{IERC1155Caller: IERC1155Caller{contract: contract}, IERC1155Transactor: IERC1155Transactor{contract: contract}, IERC1155Filterer: IERC1155Filterer{contract: contract}}, nil
}

// NewIERC1155Caller creates a new read-only instance of IERC1155, bound to a specific deployed contract.
func NewIERC1155Caller(address common.Address, caller bind.ContractCaller) (*IERC1155Caller, error) {
	contract, err := bindIERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155Caller{contract: contract}, nil
}

// NewIERC1155Transactor creates a new write-only instance of IERC1155, bound to a specific deployed contract.
func NewIERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC1155Transactor, error) {
	contract, err := bindIERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1155Transactor{contract: contract}, nil
}

// NewIERC1155Filterer creates a new log filterer instance of IERC1155, bound to a specific deployed contract.
func NewIERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC1155Filterer, error) {
	contract, err := bindIERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1155Filterer{contract: contract}, nil
}

// bindIERC1155 binds a generic wrapper to an already deployed contract.
func bindIERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155 *IERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155.Contract.IERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155 *IERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155.Contract.IERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155 *IERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155.Contract.IERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1155 *IERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1155 *IERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1155 *IERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155 *IERC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IERC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155 *IERC1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _IERC1155.Contract.BalanceOf(&_IERC1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_IERC1155 *IERC1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _IERC1155.Contract.BalanceOf(&_IERC1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155 *IERC1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _IERC1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155 *IERC1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _IERC1155.Contract.BalanceOfBatch(&_IERC1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_IERC1155 *IERC1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _IERC1155.Contract.BalanceOfBatch(&_IERC1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155 *IERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155 *IERC1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _IERC1155.Contract.IsApprovedForAll(&_IERC1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_IERC1155 *IERC1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _IERC1155.Contract.IsApprovedForAll(&_IERC1155.CallOpts, account, operator)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155 *IERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155 *IERC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1155.Contract.SupportsInterface(&_IERC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC1155 *IERC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC1155.Contract.SupportsInterface(&_IERC1155.CallOpts, interfaceId)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155 *IERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155 *IERC1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.Contract.SafeBatchTransferFrom(&_IERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_IERC1155 *IERC1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.Contract.SafeBatchTransferFrom(&_IERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155 *IERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155 *IERC1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.Contract.SafeTransferFrom(&_IERC1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_IERC1155 *IERC1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC1155.Contract.SafeTransferFrom(&_IERC1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155 *IERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155 *IERC1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155.Contract.SetApprovalForAll(&_IERC1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IERC1155 *IERC1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IERC1155.Contract.SetApprovalForAll(&_IERC1155.TransactOpts, operator, approved)
}

// IERC1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC1155 contract.
type IERC1155ApprovalForAllIterator struct {
	Event *IERC1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155ApprovalForAll)
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
		it.Event = new(IERC1155ApprovalForAll)
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
func (it *IERC1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155ApprovalForAll represents a ApprovalForAll event raised by the IERC1155 contract.
type IERC1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_IERC1155 *IERC1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*IERC1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155ApprovalForAllIterator{contract: _IERC1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_IERC1155 *IERC1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155ApprovalForAll)
				if err := _IERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_IERC1155 *IERC1155Filterer) ParseApprovalForAll(log types.Log) (*IERC1155ApprovalForAll, error) {
	event := new(IERC1155ApprovalForAll)
	if err := _IERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the IERC1155 contract.
type IERC1155TransferBatchIterator struct {
	Event *IERC1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *IERC1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155TransferBatch)
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
		it.Event = new(IERC1155TransferBatch)
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
func (it *IERC1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155TransferBatch represents a TransferBatch event raised by the IERC1155 contract.
type IERC1155TransferBatch struct {
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
func (_IERC1155 *IERC1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*IERC1155TransferBatchIterator, error) {

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

	logs, sub, err := _IERC1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155TransferBatchIterator{contract: _IERC1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_IERC1155 *IERC1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *IERC1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IERC1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155TransferBatch)
				if err := _IERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
func (_IERC1155 *IERC1155Filterer) ParseTransferBatch(log types.Log) (*IERC1155TransferBatch, error) {
	event := new(IERC1155TransferBatch)
	if err := _IERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the IERC1155 contract.
type IERC1155TransferSingleIterator struct {
	Event *IERC1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *IERC1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155TransferSingle)
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
		it.Event = new(IERC1155TransferSingle)
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
func (it *IERC1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155TransferSingle represents a TransferSingle event raised by the IERC1155 contract.
type IERC1155TransferSingle struct {
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
func (_IERC1155 *IERC1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*IERC1155TransferSingleIterator, error) {

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

	logs, sub, err := _IERC1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155TransferSingleIterator{contract: _IERC1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_IERC1155 *IERC1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *IERC1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IERC1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155TransferSingle)
				if err := _IERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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
func (_IERC1155 *IERC1155Filterer) ParseTransferSingle(log types.Log) (*IERC1155TransferSingle, error) {
	event := new(IERC1155TransferSingle)
	if err := _IERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the IERC1155 contract.
type IERC1155URIIterator struct {
	Event *IERC1155URI // Event containing the contract specifics and raw log

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
func (it *IERC1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC1155URI)
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
		it.Event = new(IERC1155URI)
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
func (it *IERC1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC1155URI represents a URI event raised by the IERC1155 contract.
type IERC1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_IERC1155 *IERC1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*IERC1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &IERC1155URIIterator{contract: _IERC1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_IERC1155 *IERC1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *IERC1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IERC1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC1155URI)
				if err := _IERC1155.contract.UnpackLog(event, "URI", log); err != nil {
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
func (_IERC1155 *IERC1155Filterer) ParseURI(log types.Log) (*IERC1155URI, error) {
	event := new(IERC1155URI)
	if err := _IERC1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721MetaData contains all meta data concerning the IERC721 contract.
var IERC721MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"081812fc": "getApproved(uint256)",
		"e985e9c5": "isApprovedForAll(address,address)",
		"6352211e": "ownerOf(uint256)",
		"42842e0e": "safeTransferFrom(address,address,uint256)",
		"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
		"a22cb465": "setApprovalForAll(address,bool)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC721MetaData.ABI instead.
var IERC721ABI = IERC721MetaData.ABI

// Deprecated: Use IERC721MetaData.Sigs instead.
// IERC721FuncSigs maps the 4-byte function signature to its string representation.
var IERC721FuncSigs = IERC721MetaData.Sigs

// IERC721 is an auto generated Go binding around an Ethereum contract.
type IERC721 struct {
	IERC721Caller     // Read-only binding to the contract
	IERC721Transactor // Write-only binding to the contract
	IERC721Filterer   // Log filterer for contract events
}

// IERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC721Session struct {
	Contract     *IERC721          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC721CallerSession struct {
	Contract *IERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC721TransactorSession struct {
	Contract     *IERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC721Raw struct {
	Contract *IERC721 // Generic contract binding to access the raw methods on
}

// IERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC721CallerRaw struct {
	Contract *IERC721Caller // Generic read-only contract binding to access the raw methods on
}

// IERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC721TransactorRaw struct {
	Contract *IERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC721 creates a new instance of IERC721, bound to a specific deployed contract.
func NewIERC721(address common.Address, backend bind.ContractBackend) (*IERC721, error) {
	contract, err := bindIERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC721{IERC721Caller: IERC721Caller{contract: contract}, IERC721Transactor: IERC721Transactor{contract: contract}, IERC721Filterer: IERC721Filterer{contract: contract}}, nil
}

// NewIERC721Caller creates a new read-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Caller(address common.Address, caller bind.ContractCaller) (*IERC721Caller, error) {
	contract, err := bindIERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Caller{contract: contract}, nil
}

// NewIERC721Transactor creates a new write-only instance of IERC721, bound to a specific deployed contract.
func NewIERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC721Transactor, error) {
	contract, err := bindIERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC721Transactor{contract: contract}, nil
}

// NewIERC721Filterer creates a new log filterer instance of IERC721, bound to a specific deployed contract.
func NewIERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC721Filterer, error) {
	contract, err := bindIERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC721Filterer{contract: contract}, nil
}

// bindIERC721 binds a generic wrapper to an already deployed contract.
func bindIERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.IERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.IERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC721 *IERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC721 *IERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC721 *IERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_IERC721 *IERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC721.Contract.BalanceOf(&_IERC721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address operator)
func (_IERC721 *IERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.GetApproved(&_IERC721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IERC721 *IERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IERC721.Contract.IsApprovedForAll(&_IERC721.CallOpts, owner, operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_IERC721 *IERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IERC721.Contract.OwnerOf(&_IERC721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC721 *IERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC721.Contract.SupportsInterface(&_IERC721.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.Approve(&_IERC721.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_IERC721 *IERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IERC721.Contract.SafeTransferFrom0(&_IERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "setApprovalForAll", operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721Session) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool _approved) returns()
func (_IERC721 *IERC721TransactorSession) SetApprovalForAll(operator common.Address, _approved bool) (*types.Transaction, error) {
	return _IERC721.Contract.SetApprovalForAll(&_IERC721.TransactOpts, operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IERC721 *IERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IERC721.Contract.TransferFrom(&_IERC721.TransactOpts, from, to, tokenId)
}

// IERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC721 contract.
type IERC721ApprovalIterator struct {
	Event *IERC721Approval // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Approval)
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
		it.Event = new(IERC721Approval)
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
func (it *IERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Approval represents a Approval event raised by the IERC721 contract.
type IERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*IERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalIterator{contract: _IERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Approval)
				if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseApproval(log types.Log) (*IERC721Approval, error) {
	event := new(IERC721Approval)
	if err := _IERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the IERC721 contract.
type IERC721ApprovalForAllIterator struct {
	Event *IERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *IERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721ApprovalForAll)
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
		it.Event = new(IERC721ApprovalForAll)
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
func (it *IERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721ApprovalForAll represents a ApprovalForAll event raised by the IERC721 contract.
type IERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*IERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &IERC721ApprovalForAllIterator{contract: _IERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *IERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721ApprovalForAll)
				if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_IERC721 *IERC721Filterer) ParseApprovalForAll(log types.Log) (*IERC721ApprovalForAll, error) {
	event := new(IERC721ApprovalForAll)
	if err := _IERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC721 contract.
type IERC721TransferIterator struct {
	Event *IERC721Transfer // Event containing the contract specifics and raw log

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
func (it *IERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC721Transfer)
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
		it.Event = new(IERC721Transfer)
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
func (it *IERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC721Transfer represents a Transfer event raised by the IERC721 contract.
type IERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*IERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &IERC721TransferIterator{contract: _IERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _IERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC721Transfer)
				if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_IERC721 *IERC721Filterer) ParseTransfer(log types.Log) (*IERC721Transfer, error) {
	event := new(IERC721Transfer)
	if err := _IERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFantomAddressRegistryMetaData contains all meta data concerning the IFantomAddressRegistry contract.
var IFantomAddressRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"artFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"artion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bundleMarketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"privateArtFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"privateFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0fb27d8c": "artFactory()",
		"18944e55": "artion()",
		"7d9f6db5": "auction()",
		"9e1f5b28": "bundleMarketplace()",
		"c45a0155": "factory()",
		"741bef1a": "priceFeed()",
		"924db13f": "privateArtFactory()",
		"735db939": "privateFactory()",
		"9d23c4c7": "tokenRegistry()",
	},
}

// IFantomAddressRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IFantomAddressRegistryMetaData.ABI instead.
var IFantomAddressRegistryABI = IFantomAddressRegistryMetaData.ABI

// Deprecated: Use IFantomAddressRegistryMetaData.Sigs instead.
// IFantomAddressRegistryFuncSigs maps the 4-byte function signature to its string representation.
var IFantomAddressRegistryFuncSigs = IFantomAddressRegistryMetaData.Sigs

// IFantomAddressRegistry is an auto generated Go binding around an Ethereum contract.
type IFantomAddressRegistry struct {
	IFantomAddressRegistryCaller     // Read-only binding to the contract
	IFantomAddressRegistryTransactor // Write-only binding to the contract
	IFantomAddressRegistryFilterer   // Log filterer for contract events
}

// IFantomAddressRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFantomAddressRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomAddressRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFantomAddressRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomAddressRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFantomAddressRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomAddressRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFantomAddressRegistrySession struct {
	Contract     *IFantomAddressRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IFantomAddressRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFantomAddressRegistryCallerSession struct {
	Contract *IFantomAddressRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IFantomAddressRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFantomAddressRegistryTransactorSession struct {
	Contract     *IFantomAddressRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IFantomAddressRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFantomAddressRegistryRaw struct {
	Contract *IFantomAddressRegistry // Generic contract binding to access the raw methods on
}

// IFantomAddressRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFantomAddressRegistryCallerRaw struct {
	Contract *IFantomAddressRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IFantomAddressRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFantomAddressRegistryTransactorRaw struct {
	Contract *IFantomAddressRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFantomAddressRegistry creates a new instance of IFantomAddressRegistry, bound to a specific deployed contract.
func NewIFantomAddressRegistry(address common.Address, backend bind.ContractBackend) (*IFantomAddressRegistry, error) {
	contract, err := bindIFantomAddressRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFantomAddressRegistry{IFantomAddressRegistryCaller: IFantomAddressRegistryCaller{contract: contract}, IFantomAddressRegistryTransactor: IFantomAddressRegistryTransactor{contract: contract}, IFantomAddressRegistryFilterer: IFantomAddressRegistryFilterer{contract: contract}}, nil
}

// NewIFantomAddressRegistryCaller creates a new read-only instance of IFantomAddressRegistry, bound to a specific deployed contract.
func NewIFantomAddressRegistryCaller(address common.Address, caller bind.ContractCaller) (*IFantomAddressRegistryCaller, error) {
	contract, err := bindIFantomAddressRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFantomAddressRegistryCaller{contract: contract}, nil
}

// NewIFantomAddressRegistryTransactor creates a new write-only instance of IFantomAddressRegistry, bound to a specific deployed contract.
func NewIFantomAddressRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IFantomAddressRegistryTransactor, error) {
	contract, err := bindIFantomAddressRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFantomAddressRegistryTransactor{contract: contract}, nil
}

// NewIFantomAddressRegistryFilterer creates a new log filterer instance of IFantomAddressRegistry, bound to a specific deployed contract.
func NewIFantomAddressRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IFantomAddressRegistryFilterer, error) {
	contract, err := bindIFantomAddressRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFantomAddressRegistryFilterer{contract: contract}, nil
}

// bindIFantomAddressRegistry binds a generic wrapper to an already deployed contract.
func bindIFantomAddressRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFantomAddressRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFantomAddressRegistry *IFantomAddressRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFantomAddressRegistry.Contract.IFantomAddressRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFantomAddressRegistry *IFantomAddressRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFantomAddressRegistry.Contract.IFantomAddressRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFantomAddressRegistry *IFantomAddressRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFantomAddressRegistry.Contract.IFantomAddressRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFantomAddressRegistry *IFantomAddressRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFantomAddressRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFantomAddressRegistry *IFantomAddressRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFantomAddressRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFantomAddressRegistry *IFantomAddressRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFantomAddressRegistry.Contract.contract.Transact(opts, method, params...)
}

// ArtFactory is a free data retrieval call binding the contract method 0x0fb27d8c.
//
// Solidity: function artFactory() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCaller) ArtFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFantomAddressRegistry.contract.Call(opts, &out, "artFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ArtFactory is a free data retrieval call binding the contract method 0x0fb27d8c.
//
// Solidity: function artFactory() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistrySession) ArtFactory() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.ArtFactory(&_IFantomAddressRegistry.CallOpts)
}

// ArtFactory is a free data retrieval call binding the contract method 0x0fb27d8c.
//
// Solidity: function artFactory() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCallerSession) ArtFactory() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.ArtFactory(&_IFantomAddressRegistry.CallOpts)
}

// Artion is a free data retrieval call binding the contract method 0x18944e55.
//
// Solidity: function artion() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCaller) Artion(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFantomAddressRegistry.contract.Call(opts, &out, "artion")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Artion is a free data retrieval call binding the contract method 0x18944e55.
//
// Solidity: function artion() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistrySession) Artion() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.Artion(&_IFantomAddressRegistry.CallOpts)
}

// Artion is a free data retrieval call binding the contract method 0x18944e55.
//
// Solidity: function artion() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCallerSession) Artion() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.Artion(&_IFantomAddressRegistry.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCaller) Auction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFantomAddressRegistry.contract.Call(opts, &out, "auction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistrySession) Auction() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.Auction(&_IFantomAddressRegistry.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCallerSession) Auction() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.Auction(&_IFantomAddressRegistry.CallOpts)
}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCaller) BundleMarketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFantomAddressRegistry.contract.Call(opts, &out, "bundleMarketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistrySession) BundleMarketplace() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.BundleMarketplace(&_IFantomAddressRegistry.CallOpts)
}

// BundleMarketplace is a free data retrieval call binding the contract method 0x9e1f5b28.
//
// Solidity: function bundleMarketplace() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCallerSession) BundleMarketplace() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.BundleMarketplace(&_IFantomAddressRegistry.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFantomAddressRegistry.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistrySession) Factory() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.Factory(&_IFantomAddressRegistry.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCallerSession) Factory() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.Factory(&_IFantomAddressRegistry.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFantomAddressRegistry.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistrySession) PriceFeed() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.PriceFeed(&_IFantomAddressRegistry.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCallerSession) PriceFeed() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.PriceFeed(&_IFantomAddressRegistry.CallOpts)
}

// PrivateArtFactory is a free data retrieval call binding the contract method 0x924db13f.
//
// Solidity: function privateArtFactory() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCaller) PrivateArtFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFantomAddressRegistry.contract.Call(opts, &out, "privateArtFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrivateArtFactory is a free data retrieval call binding the contract method 0x924db13f.
//
// Solidity: function privateArtFactory() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistrySession) PrivateArtFactory() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.PrivateArtFactory(&_IFantomAddressRegistry.CallOpts)
}

// PrivateArtFactory is a free data retrieval call binding the contract method 0x924db13f.
//
// Solidity: function privateArtFactory() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCallerSession) PrivateArtFactory() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.PrivateArtFactory(&_IFantomAddressRegistry.CallOpts)
}

// PrivateFactory is a free data retrieval call binding the contract method 0x735db939.
//
// Solidity: function privateFactory() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCaller) PrivateFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFantomAddressRegistry.contract.Call(opts, &out, "privateFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrivateFactory is a free data retrieval call binding the contract method 0x735db939.
//
// Solidity: function privateFactory() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistrySession) PrivateFactory() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.PrivateFactory(&_IFantomAddressRegistry.CallOpts)
}

// PrivateFactory is a free data retrieval call binding the contract method 0x735db939.
//
// Solidity: function privateFactory() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCallerSession) PrivateFactory() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.PrivateFactory(&_IFantomAddressRegistry.CallOpts)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCaller) TokenRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFantomAddressRegistry.contract.Call(opts, &out, "tokenRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistrySession) TokenRegistry() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.TokenRegistry(&_IFantomAddressRegistry.CallOpts)
}

// TokenRegistry is a free data retrieval call binding the contract method 0x9d23c4c7.
//
// Solidity: function tokenRegistry() view returns(address)
func (_IFantomAddressRegistry *IFantomAddressRegistryCallerSession) TokenRegistry() (common.Address, error) {
	return _IFantomAddressRegistry.Contract.TokenRegistry(&_IFantomAddressRegistry.CallOpts)
}

// IFantomAuctionMetaData contains all meta data concerning the IFantomAuction contract.
var IFantomAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"44f91c1e": "auctions(address,uint256)",
	},
}

// IFantomAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use IFantomAuctionMetaData.ABI instead.
var IFantomAuctionABI = IFantomAuctionMetaData.ABI

// Deprecated: Use IFantomAuctionMetaData.Sigs instead.
// IFantomAuctionFuncSigs maps the 4-byte function signature to its string representation.
var IFantomAuctionFuncSigs = IFantomAuctionMetaData.Sigs

// IFantomAuction is an auto generated Go binding around an Ethereum contract.
type IFantomAuction struct {
	IFantomAuctionCaller     // Read-only binding to the contract
	IFantomAuctionTransactor // Write-only binding to the contract
	IFantomAuctionFilterer   // Log filterer for contract events
}

// IFantomAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFantomAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFantomAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFantomAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFantomAuctionSession struct {
	Contract     *IFantomAuction   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFantomAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFantomAuctionCallerSession struct {
	Contract *IFantomAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IFantomAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFantomAuctionTransactorSession struct {
	Contract     *IFantomAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IFantomAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFantomAuctionRaw struct {
	Contract *IFantomAuction // Generic contract binding to access the raw methods on
}

// IFantomAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFantomAuctionCallerRaw struct {
	Contract *IFantomAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// IFantomAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFantomAuctionTransactorRaw struct {
	Contract *IFantomAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFantomAuction creates a new instance of IFantomAuction, bound to a specific deployed contract.
func NewIFantomAuction(address common.Address, backend bind.ContractBackend) (*IFantomAuction, error) {
	contract, err := bindIFantomAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFantomAuction{IFantomAuctionCaller: IFantomAuctionCaller{contract: contract}, IFantomAuctionTransactor: IFantomAuctionTransactor{contract: contract}, IFantomAuctionFilterer: IFantomAuctionFilterer{contract: contract}}, nil
}

// NewIFantomAuctionCaller creates a new read-only instance of IFantomAuction, bound to a specific deployed contract.
func NewIFantomAuctionCaller(address common.Address, caller bind.ContractCaller) (*IFantomAuctionCaller, error) {
	contract, err := bindIFantomAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFantomAuctionCaller{contract: contract}, nil
}

// NewIFantomAuctionTransactor creates a new write-only instance of IFantomAuction, bound to a specific deployed contract.
func NewIFantomAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*IFantomAuctionTransactor, error) {
	contract, err := bindIFantomAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFantomAuctionTransactor{contract: contract}, nil
}

// NewIFantomAuctionFilterer creates a new log filterer instance of IFantomAuction, bound to a specific deployed contract.
func NewIFantomAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*IFantomAuctionFilterer, error) {
	contract, err := bindIFantomAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFantomAuctionFilterer{contract: contract}, nil
}

// bindIFantomAuction binds a generic wrapper to an already deployed contract.
func bindIFantomAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFantomAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFantomAuction *IFantomAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFantomAuction.Contract.IFantomAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFantomAuction *IFantomAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFantomAuction.Contract.IFantomAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFantomAuction *IFantomAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFantomAuction.Contract.IFantomAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFantomAuction *IFantomAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFantomAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFantomAuction *IFantomAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFantomAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFantomAuction *IFantomAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFantomAuction.Contract.contract.Transact(opts, method, params...)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address, address, uint256, uint256, uint256, bool)
func (_IFantomAuction *IFantomAuctionCaller) Auctions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _IFantomAuction.contract.Call(opts, &out, "auctions", arg0, arg1)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, err

}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address, address, uint256, uint256, uint256, bool)
func (_IFantomAuction *IFantomAuctionSession) Auctions(arg0 common.Address, arg1 *big.Int) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, bool, error) {
	return _IFantomAuction.Contract.Auctions(&_IFantomAuction.CallOpts, arg0, arg1)
}

// Auctions is a free data retrieval call binding the contract method 0x44f91c1e.
//
// Solidity: function auctions(address , uint256 ) view returns(address, address, uint256, uint256, uint256, bool)
func (_IFantomAuction *IFantomAuctionCallerSession) Auctions(arg0 common.Address, arg1 *big.Int) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, bool, error) {
	return _IFantomAuction.Contract.Auctions(&_IFantomAuction.CallOpts, arg0, arg1)
}

// IFantomBundleMarketplaceMetaData contains all meta data concerning the IFantomBundleMarketplace contract.
var IFantomBundleMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validateItemSold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"27c1f4f4": "validateItemSold(address,uint256,uint256)",
	},
}

// IFantomBundleMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use IFantomBundleMarketplaceMetaData.ABI instead.
var IFantomBundleMarketplaceABI = IFantomBundleMarketplaceMetaData.ABI

// Deprecated: Use IFantomBundleMarketplaceMetaData.Sigs instead.
// IFantomBundleMarketplaceFuncSigs maps the 4-byte function signature to its string representation.
var IFantomBundleMarketplaceFuncSigs = IFantomBundleMarketplaceMetaData.Sigs

// IFantomBundleMarketplace is an auto generated Go binding around an Ethereum contract.
type IFantomBundleMarketplace struct {
	IFantomBundleMarketplaceCaller     // Read-only binding to the contract
	IFantomBundleMarketplaceTransactor // Write-only binding to the contract
	IFantomBundleMarketplaceFilterer   // Log filterer for contract events
}

// IFantomBundleMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFantomBundleMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomBundleMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFantomBundleMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomBundleMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFantomBundleMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomBundleMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFantomBundleMarketplaceSession struct {
	Contract     *IFantomBundleMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IFantomBundleMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFantomBundleMarketplaceCallerSession struct {
	Contract *IFantomBundleMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IFantomBundleMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFantomBundleMarketplaceTransactorSession struct {
	Contract     *IFantomBundleMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IFantomBundleMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFantomBundleMarketplaceRaw struct {
	Contract *IFantomBundleMarketplace // Generic contract binding to access the raw methods on
}

// IFantomBundleMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFantomBundleMarketplaceCallerRaw struct {
	Contract *IFantomBundleMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// IFantomBundleMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFantomBundleMarketplaceTransactorRaw struct {
	Contract *IFantomBundleMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFantomBundleMarketplace creates a new instance of IFantomBundleMarketplace, bound to a specific deployed contract.
func NewIFantomBundleMarketplace(address common.Address, backend bind.ContractBackend) (*IFantomBundleMarketplace, error) {
	contract, err := bindIFantomBundleMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFantomBundleMarketplace{IFantomBundleMarketplaceCaller: IFantomBundleMarketplaceCaller{contract: contract}, IFantomBundleMarketplaceTransactor: IFantomBundleMarketplaceTransactor{contract: contract}, IFantomBundleMarketplaceFilterer: IFantomBundleMarketplaceFilterer{contract: contract}}, nil
}

// NewIFantomBundleMarketplaceCaller creates a new read-only instance of IFantomBundleMarketplace, bound to a specific deployed contract.
func NewIFantomBundleMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*IFantomBundleMarketplaceCaller, error) {
	contract, err := bindIFantomBundleMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFantomBundleMarketplaceCaller{contract: contract}, nil
}

// NewIFantomBundleMarketplaceTransactor creates a new write-only instance of IFantomBundleMarketplace, bound to a specific deployed contract.
func NewIFantomBundleMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*IFantomBundleMarketplaceTransactor, error) {
	contract, err := bindIFantomBundleMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFantomBundleMarketplaceTransactor{contract: contract}, nil
}

// NewIFantomBundleMarketplaceFilterer creates a new log filterer instance of IFantomBundleMarketplace, bound to a specific deployed contract.
func NewIFantomBundleMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*IFantomBundleMarketplaceFilterer, error) {
	contract, err := bindIFantomBundleMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFantomBundleMarketplaceFilterer{contract: contract}, nil
}

// bindIFantomBundleMarketplace binds a generic wrapper to an already deployed contract.
func bindIFantomBundleMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFantomBundleMarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFantomBundleMarketplace *IFantomBundleMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFantomBundleMarketplace.Contract.IFantomBundleMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFantomBundleMarketplace *IFantomBundleMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFantomBundleMarketplace.Contract.IFantomBundleMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFantomBundleMarketplace *IFantomBundleMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFantomBundleMarketplace.Contract.IFantomBundleMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFantomBundleMarketplace *IFantomBundleMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFantomBundleMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFantomBundleMarketplace *IFantomBundleMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFantomBundleMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFantomBundleMarketplace *IFantomBundleMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFantomBundleMarketplace.Contract.contract.Transact(opts, method, params...)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0x27c1f4f4.
//
// Solidity: function validateItemSold(address , uint256 , uint256 ) returns()
func (_IFantomBundleMarketplace *IFantomBundleMarketplaceTransactor) ValidateItemSold(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _IFantomBundleMarketplace.contract.Transact(opts, "validateItemSold", arg0, arg1, arg2)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0x27c1f4f4.
//
// Solidity: function validateItemSold(address , uint256 , uint256 ) returns()
func (_IFantomBundleMarketplace *IFantomBundleMarketplaceSession) ValidateItemSold(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _IFantomBundleMarketplace.Contract.ValidateItemSold(&_IFantomBundleMarketplace.TransactOpts, arg0, arg1, arg2)
}

// ValidateItemSold is a paid mutator transaction binding the contract method 0x27c1f4f4.
//
// Solidity: function validateItemSold(address , uint256 , uint256 ) returns()
func (_IFantomBundleMarketplace *IFantomBundleMarketplaceTransactorSession) ValidateItemSold(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (*types.Transaction, error) {
	return _IFantomBundleMarketplace.Contract.ValidateItemSold(&_IFantomBundleMarketplace.TransactOpts, arg0, arg1, arg2)
}

// IFantomNFTFactoryMetaData contains all meta data concerning the IFantomNFTFactory contract.
var IFantomNFTFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f6a3d24e": "exists(address)",
	},
}

// IFantomNFTFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IFantomNFTFactoryMetaData.ABI instead.
var IFantomNFTFactoryABI = IFantomNFTFactoryMetaData.ABI

// Deprecated: Use IFantomNFTFactoryMetaData.Sigs instead.
// IFantomNFTFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IFantomNFTFactoryFuncSigs = IFantomNFTFactoryMetaData.Sigs

// IFantomNFTFactory is an auto generated Go binding around an Ethereum contract.
type IFantomNFTFactory struct {
	IFantomNFTFactoryCaller     // Read-only binding to the contract
	IFantomNFTFactoryTransactor // Write-only binding to the contract
	IFantomNFTFactoryFilterer   // Log filterer for contract events
}

// IFantomNFTFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFantomNFTFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomNFTFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFantomNFTFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomNFTFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFantomNFTFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomNFTFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFantomNFTFactorySession struct {
	Contract     *IFantomNFTFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IFantomNFTFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFantomNFTFactoryCallerSession struct {
	Contract *IFantomNFTFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IFantomNFTFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFantomNFTFactoryTransactorSession struct {
	Contract     *IFantomNFTFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IFantomNFTFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFantomNFTFactoryRaw struct {
	Contract *IFantomNFTFactory // Generic contract binding to access the raw methods on
}

// IFantomNFTFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFantomNFTFactoryCallerRaw struct {
	Contract *IFantomNFTFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IFantomNFTFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFantomNFTFactoryTransactorRaw struct {
	Contract *IFantomNFTFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFantomNFTFactory creates a new instance of IFantomNFTFactory, bound to a specific deployed contract.
func NewIFantomNFTFactory(address common.Address, backend bind.ContractBackend) (*IFantomNFTFactory, error) {
	contract, err := bindIFantomNFTFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFantomNFTFactory{IFantomNFTFactoryCaller: IFantomNFTFactoryCaller{contract: contract}, IFantomNFTFactoryTransactor: IFantomNFTFactoryTransactor{contract: contract}, IFantomNFTFactoryFilterer: IFantomNFTFactoryFilterer{contract: contract}}, nil
}

// NewIFantomNFTFactoryCaller creates a new read-only instance of IFantomNFTFactory, bound to a specific deployed contract.
func NewIFantomNFTFactoryCaller(address common.Address, caller bind.ContractCaller) (*IFantomNFTFactoryCaller, error) {
	contract, err := bindIFantomNFTFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFantomNFTFactoryCaller{contract: contract}, nil
}

// NewIFantomNFTFactoryTransactor creates a new write-only instance of IFantomNFTFactory, bound to a specific deployed contract.
func NewIFantomNFTFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IFantomNFTFactoryTransactor, error) {
	contract, err := bindIFantomNFTFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFantomNFTFactoryTransactor{contract: contract}, nil
}

// NewIFantomNFTFactoryFilterer creates a new log filterer instance of IFantomNFTFactory, bound to a specific deployed contract.
func NewIFantomNFTFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IFantomNFTFactoryFilterer, error) {
	contract, err := bindIFantomNFTFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFantomNFTFactoryFilterer{contract: contract}, nil
}

// bindIFantomNFTFactory binds a generic wrapper to an already deployed contract.
func bindIFantomNFTFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFantomNFTFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFantomNFTFactory *IFantomNFTFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFantomNFTFactory.Contract.IFantomNFTFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFantomNFTFactory *IFantomNFTFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFantomNFTFactory.Contract.IFantomNFTFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFantomNFTFactory *IFantomNFTFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFantomNFTFactory.Contract.IFantomNFTFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFantomNFTFactory *IFantomNFTFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFantomNFTFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFantomNFTFactory *IFantomNFTFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFantomNFTFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFantomNFTFactory *IFantomNFTFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFantomNFTFactory.Contract.contract.Transact(opts, method, params...)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_IFantomNFTFactory *IFantomNFTFactoryCaller) Exists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IFantomNFTFactory.contract.Call(opts, &out, "exists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_IFantomNFTFactory *IFantomNFTFactorySession) Exists(arg0 common.Address) (bool, error) {
	return _IFantomNFTFactory.Contract.Exists(&_IFantomNFTFactory.CallOpts, arg0)
}

// Exists is a free data retrieval call binding the contract method 0xf6a3d24e.
//
// Solidity: function exists(address ) view returns(bool)
func (_IFantomNFTFactory *IFantomNFTFactoryCallerSession) Exists(arg0 common.Address) (bool, error) {
	return _IFantomNFTFactory.Contract.Exists(&_IFantomNFTFactory.CallOpts, arg0)
}

// IFantomPriceFeedMetaData contains all meta data concerning the IFantomPriceFeed contract.
var IFantomPriceFeedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wFTM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"41976e09": "getPrice(address)",
		"050d63ec": "wFTM()",
	},
}

// IFantomPriceFeedABI is the input ABI used to generate the binding from.
// Deprecated: Use IFantomPriceFeedMetaData.ABI instead.
var IFantomPriceFeedABI = IFantomPriceFeedMetaData.ABI

// Deprecated: Use IFantomPriceFeedMetaData.Sigs instead.
// IFantomPriceFeedFuncSigs maps the 4-byte function signature to its string representation.
var IFantomPriceFeedFuncSigs = IFantomPriceFeedMetaData.Sigs

// IFantomPriceFeed is an auto generated Go binding around an Ethereum contract.
type IFantomPriceFeed struct {
	IFantomPriceFeedCaller     // Read-only binding to the contract
	IFantomPriceFeedTransactor // Write-only binding to the contract
	IFantomPriceFeedFilterer   // Log filterer for contract events
}

// IFantomPriceFeedCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFantomPriceFeedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomPriceFeedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFantomPriceFeedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomPriceFeedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFantomPriceFeedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomPriceFeedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFantomPriceFeedSession struct {
	Contract     *IFantomPriceFeed // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFantomPriceFeedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFantomPriceFeedCallerSession struct {
	Contract *IFantomPriceFeedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IFantomPriceFeedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFantomPriceFeedTransactorSession struct {
	Contract     *IFantomPriceFeedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IFantomPriceFeedRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFantomPriceFeedRaw struct {
	Contract *IFantomPriceFeed // Generic contract binding to access the raw methods on
}

// IFantomPriceFeedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFantomPriceFeedCallerRaw struct {
	Contract *IFantomPriceFeedCaller // Generic read-only contract binding to access the raw methods on
}

// IFantomPriceFeedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFantomPriceFeedTransactorRaw struct {
	Contract *IFantomPriceFeedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFantomPriceFeed creates a new instance of IFantomPriceFeed, bound to a specific deployed contract.
func NewIFantomPriceFeed(address common.Address, backend bind.ContractBackend) (*IFantomPriceFeed, error) {
	contract, err := bindIFantomPriceFeed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFantomPriceFeed{IFantomPriceFeedCaller: IFantomPriceFeedCaller{contract: contract}, IFantomPriceFeedTransactor: IFantomPriceFeedTransactor{contract: contract}, IFantomPriceFeedFilterer: IFantomPriceFeedFilterer{contract: contract}}, nil
}

// NewIFantomPriceFeedCaller creates a new read-only instance of IFantomPriceFeed, bound to a specific deployed contract.
func NewIFantomPriceFeedCaller(address common.Address, caller bind.ContractCaller) (*IFantomPriceFeedCaller, error) {
	contract, err := bindIFantomPriceFeed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFantomPriceFeedCaller{contract: contract}, nil
}

// NewIFantomPriceFeedTransactor creates a new write-only instance of IFantomPriceFeed, bound to a specific deployed contract.
func NewIFantomPriceFeedTransactor(address common.Address, transactor bind.ContractTransactor) (*IFantomPriceFeedTransactor, error) {
	contract, err := bindIFantomPriceFeed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFantomPriceFeedTransactor{contract: contract}, nil
}

// NewIFantomPriceFeedFilterer creates a new log filterer instance of IFantomPriceFeed, bound to a specific deployed contract.
func NewIFantomPriceFeedFilterer(address common.Address, filterer bind.ContractFilterer) (*IFantomPriceFeedFilterer, error) {
	contract, err := bindIFantomPriceFeed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFantomPriceFeedFilterer{contract: contract}, nil
}

// bindIFantomPriceFeed binds a generic wrapper to an already deployed contract.
func bindIFantomPriceFeed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFantomPriceFeedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFantomPriceFeed *IFantomPriceFeedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFantomPriceFeed.Contract.IFantomPriceFeedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFantomPriceFeed *IFantomPriceFeedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFantomPriceFeed.Contract.IFantomPriceFeedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFantomPriceFeed *IFantomPriceFeedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFantomPriceFeed.Contract.IFantomPriceFeedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFantomPriceFeed *IFantomPriceFeedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFantomPriceFeed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFantomPriceFeed *IFantomPriceFeedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFantomPriceFeed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFantomPriceFeed *IFantomPriceFeedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFantomPriceFeed.Contract.contract.Transact(opts, method, params...)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address ) view returns(int256, uint8)
func (_IFantomPriceFeed *IFantomPriceFeedCaller) GetPrice(opts *bind.CallOpts, arg0 common.Address) (*big.Int, uint8, error) {
	var out []interface{}
	err := _IFantomPriceFeed.contract.Call(opts, &out, "getPrice", arg0)

	if err != nil {
		return *new(*big.Int), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address ) view returns(int256, uint8)
func (_IFantomPriceFeed *IFantomPriceFeedSession) GetPrice(arg0 common.Address) (*big.Int, uint8, error) {
	return _IFantomPriceFeed.Contract.GetPrice(&_IFantomPriceFeed.CallOpts, arg0)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address ) view returns(int256, uint8)
func (_IFantomPriceFeed *IFantomPriceFeedCallerSession) GetPrice(arg0 common.Address) (*big.Int, uint8, error) {
	return _IFantomPriceFeed.Contract.GetPrice(&_IFantomPriceFeed.CallOpts, arg0)
}

// WFTM is a free data retrieval call binding the contract method 0x050d63ec.
//
// Solidity: function wFTM() view returns(address)
func (_IFantomPriceFeed *IFantomPriceFeedCaller) WFTM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IFantomPriceFeed.contract.Call(opts, &out, "wFTM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WFTM is a free data retrieval call binding the contract method 0x050d63ec.
//
// Solidity: function wFTM() view returns(address)
func (_IFantomPriceFeed *IFantomPriceFeedSession) WFTM() (common.Address, error) {
	return _IFantomPriceFeed.Contract.WFTM(&_IFantomPriceFeed.CallOpts)
}

// WFTM is a free data retrieval call binding the contract method 0x050d63ec.
//
// Solidity: function wFTM() view returns(address)
func (_IFantomPriceFeed *IFantomPriceFeedCallerSession) WFTM() (common.Address, error) {
	return _IFantomPriceFeed.Contract.WFTM(&_IFantomPriceFeed.CallOpts)
}

// IFantomTokenRegistryMetaData contains all meta data concerning the IFantomTokenRegistry contract.
var IFantomTokenRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"enabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"98afdfe3": "enabled(address)",
	},
}

// IFantomTokenRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IFantomTokenRegistryMetaData.ABI instead.
var IFantomTokenRegistryABI = IFantomTokenRegistryMetaData.ABI

// Deprecated: Use IFantomTokenRegistryMetaData.Sigs instead.
// IFantomTokenRegistryFuncSigs maps the 4-byte function signature to its string representation.
var IFantomTokenRegistryFuncSigs = IFantomTokenRegistryMetaData.Sigs

// IFantomTokenRegistry is an auto generated Go binding around an Ethereum contract.
type IFantomTokenRegistry struct {
	IFantomTokenRegistryCaller     // Read-only binding to the contract
	IFantomTokenRegistryTransactor // Write-only binding to the contract
	IFantomTokenRegistryFilterer   // Log filterer for contract events
}

// IFantomTokenRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFantomTokenRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomTokenRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFantomTokenRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomTokenRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFantomTokenRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFantomTokenRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFantomTokenRegistrySession struct {
	Contract     *IFantomTokenRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IFantomTokenRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFantomTokenRegistryCallerSession struct {
	Contract *IFantomTokenRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IFantomTokenRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFantomTokenRegistryTransactorSession struct {
	Contract     *IFantomTokenRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IFantomTokenRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFantomTokenRegistryRaw struct {
	Contract *IFantomTokenRegistry // Generic contract binding to access the raw methods on
}

// IFantomTokenRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFantomTokenRegistryCallerRaw struct {
	Contract *IFantomTokenRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IFantomTokenRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFantomTokenRegistryTransactorRaw struct {
	Contract *IFantomTokenRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFantomTokenRegistry creates a new instance of IFantomTokenRegistry, bound to a specific deployed contract.
func NewIFantomTokenRegistry(address common.Address, backend bind.ContractBackend) (*IFantomTokenRegistry, error) {
	contract, err := bindIFantomTokenRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFantomTokenRegistry{IFantomTokenRegistryCaller: IFantomTokenRegistryCaller{contract: contract}, IFantomTokenRegistryTransactor: IFantomTokenRegistryTransactor{contract: contract}, IFantomTokenRegistryFilterer: IFantomTokenRegistryFilterer{contract: contract}}, nil
}

// NewIFantomTokenRegistryCaller creates a new read-only instance of IFantomTokenRegistry, bound to a specific deployed contract.
func NewIFantomTokenRegistryCaller(address common.Address, caller bind.ContractCaller) (*IFantomTokenRegistryCaller, error) {
	contract, err := bindIFantomTokenRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFantomTokenRegistryCaller{contract: contract}, nil
}

// NewIFantomTokenRegistryTransactor creates a new write-only instance of IFantomTokenRegistry, bound to a specific deployed contract.
func NewIFantomTokenRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IFantomTokenRegistryTransactor, error) {
	contract, err := bindIFantomTokenRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFantomTokenRegistryTransactor{contract: contract}, nil
}

// NewIFantomTokenRegistryFilterer creates a new log filterer instance of IFantomTokenRegistry, bound to a specific deployed contract.
func NewIFantomTokenRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IFantomTokenRegistryFilterer, error) {
	contract, err := bindIFantomTokenRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFantomTokenRegistryFilterer{contract: contract}, nil
}

// bindIFantomTokenRegistry binds a generic wrapper to an already deployed contract.
func bindIFantomTokenRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFantomTokenRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFantomTokenRegistry *IFantomTokenRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFantomTokenRegistry.Contract.IFantomTokenRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFantomTokenRegistry *IFantomTokenRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFantomTokenRegistry.Contract.IFantomTokenRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFantomTokenRegistry *IFantomTokenRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFantomTokenRegistry.Contract.IFantomTokenRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFantomTokenRegistry *IFantomTokenRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFantomTokenRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFantomTokenRegistry *IFantomTokenRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFantomTokenRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFantomTokenRegistry *IFantomTokenRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFantomTokenRegistry.Contract.contract.Transact(opts, method, params...)
}

// Enabled is a free data retrieval call binding the contract method 0x98afdfe3.
//
// Solidity: function enabled(address ) view returns(bool)
func (_IFantomTokenRegistry *IFantomTokenRegistryCaller) Enabled(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _IFantomTokenRegistry.contract.Call(opts, &out, "enabled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Enabled is a free data retrieval call binding the contract method 0x98afdfe3.
//
// Solidity: function enabled(address ) view returns(bool)
func (_IFantomTokenRegistry *IFantomTokenRegistrySession) Enabled(arg0 common.Address) (bool, error) {
	return _IFantomTokenRegistry.Contract.Enabled(&_IFantomTokenRegistry.CallOpts, arg0)
}

// Enabled is a free data retrieval call binding the contract method 0x98afdfe3.
//
// Solidity: function enabled(address ) view returns(bool)
func (_IFantomTokenRegistry *IFantomTokenRegistryCallerSession) Enabled(arg0 common.Address) (bool, error) {
	return _IFantomTokenRegistry.Contract.Enabled(&_IFantomTokenRegistry.CallOpts, arg0)
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

// ReentrancyGuardMetaData contains all meta data concerning the ReentrancyGuard contract.
var ReentrancyGuardMetaData = &bind.MetaData{
	ABI: "[]",
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardMetaData.ABI instead.
var ReentrancyGuardABI = ReentrancyGuardMetaData.ABI

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fe09875e3fedadcded4baa0b81bc508106fbbcdbc2cde21ae01fcbd2d627438d64736f6c634300080b0033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200f1328b21ce5dfa07614cff7e8fc2dfeba3c51b8284303bbd6a5af612cc056b964736f6c634300080b0033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
