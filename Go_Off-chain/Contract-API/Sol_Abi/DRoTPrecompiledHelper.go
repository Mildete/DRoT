// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DRoT

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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
)

// DRoTABI is the input ABI used to generate the binding from.
const DRoTABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"message\",\"type\":\"string\"},{\"name\":\"rHex\",\"type\":\"string\"},{\"name\":\"sHex\",\"type\":\"string\"},{\"name\":\"pubKeyHex\",\"type\":\"string\"}],\"name\":\"VerifyEcdsa\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"nodeID\",\"type\":\"string\"},{\"name\":\"attestationTime\",\"type\":\"uint256\"},{\"name\":\"measurement\",\"type\":\"string\"},{\"name\":\"referenceMeasurement\",\"type\":\"string\"},{\"name\":\"reliabilityThreshold\",\"type\":\"uint256\"}],\"name\":\"verifyAR\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hexString\",\"type\":\"string\"}],\"name\":\"hexStringToBytes32\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeID\",\"type\":\"string\"},{\"name\":\"attestationTime\",\"type\":\"uint256\"},{\"name\":\"measurement\",\"type\":\"string\"},{\"name\":\"referenceMeasurement\",\"type\":\"string\"},{\"name\":\"reliabilityThreshold\",\"type\":\"uint256\"}],\"name\":\"verifyARandStore\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"messageHash\",\"type\":\"string\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"string\"},{\"name\":\"s\",\"type\":\"string\"},{\"name\":\"expectedSigner\",\"type\":\"address\"}],\"name\":\"EcdsaVerifySol\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tableRegistrationAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"UpdateStatusResult\",\"type\":\"event\"}]"

// DRoTBin is the compiled bytecode used for deploying new contracts.
var DRoTBin = "0x608060405234801561001057600080fd5b506040516020806115db833981018060405261002f919081019061010f565b6150096000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550615058600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610158565b60006101078251610138565b905092915050565b60006020828403121561012157600080fd5b600061012f848285016100fb565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b611474806101676000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806321d4f02d1461007257806386848a91146100af5780638b6bb00f146100ee5780638c67af271461012b578063ada76d581461016a575b600080fd5b34801561007e57600080fd5b5061009960048036036100949190810190610c5c565b6101a7565b6040516100a6919061107d565b60405180910390f35b3480156100bb57600080fd5b506100d660048036036100d19190810190610d1f565b610281565b6040516100e593929190611098565b60405180910390f35b3480156100fa57600080fd5b5061011560048036036101109190810190610c1b565b61036d565b60405161012291906110d6565b60405180910390f35b34801561013757600080fd5b50610152600480360361014d9190810190610d1f565b61050a565b60405161016193929190611098565b60405180910390f35b34801561017657600080fd5b50610191600480360361018c9190810190610dde565b610870565b60405161019e919061107d565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166321d4f02d868686866040518563ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016102259493929190611151565b602060405180830381600087803b15801561023f57600080fd5b505af1158015610253573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506102779190810190610b0e565b9050949350505050565b6000806060600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166386848a9189898989896040518663ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610305959493929190611217565b600060405180830381600087803b15801561031f57600080fd5b505af1158015610333573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525061035c9190810190610b37565b925092509250955095509592505050565b600060606000806000859350835192506042831415156103c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b99061127f565b60405180910390fd5b600290505b828110156104fe5761045684600183018151811015156103e357fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027f0100000000000000000000000000000000000000000000000000000000000000900461092a565b60106104dc868481518110151561046957fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f0100000000000000000000000000000000000000000000000000000000000000027f0100000000000000000000000000000000000000000000000000000000000000900461092a565b020160ff16600102600883600019169060020a021791506002810190506103c7565b81945050505050919050565b60008060606000606060008060606000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166322cd843f8f6040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161059191906111b2565b600060405180830381600087803b1580156105ab57600080fd5b505af11580156105bf573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506105e89190810190610bc7565b95509550600086141515610631576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610628906112bf565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166386848a918f8f8f8f8f6040518663ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016106b0959493929190611217565b600060405180830381600087803b1580156106ca57600080fd5b505af11580156106de573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506107079190810190610b37565b935093509350600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d25a06f48f856040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016107869291906111e7565b602060405180830381600087803b1580156107a057600080fd5b505af11580156107b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506107d89190810190610b9e565b905060008113151561081f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610816906112df565b60405180910390fd5b7f056d4ff851ef14b8d2a7611d6bcd01b76e5d5d03485176f7aa975576098c5d058160405161084e9190611136565b60405180910390a1838383985098509850505050505050955095509592505050565b60008060008060006108818a61036d565b935061088c8861036d565b92506108978761036d565b91506001848a8585604051600081526020016040526040516108bc94939291906110f1565b60206040516020810390808403906000865af11580156108e0573d6000803e3d6000fd5b5050506020604051035190508573ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161494505050505095945050505050565b600060308260ff1610158015610944575060398260ff1611155b15610954576030820390506109e5565b60618260ff161015801561096c575060668260ff1611155b1561097f57606182600a010390506109e5565b60418260ff1610158015610997575060468260ff1611155b156109aa57604182600a010390506109e5565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109dc9061129f565b60405180910390fd5b919050565b60006109f6823561139a565b905092915050565b6000610a0a82516113ba565b905092915050565b6000610a1e82516113c6565b905092915050565b600082601f8301121515610a3957600080fd5b8135610a4c610a478261132c565b6112ff565b91508082526020830160208301858383011115610a6857600080fd5b610a738382846113e7565b50505092915050565b600082601f8301121515610a8f57600080fd5b8151610aa2610a9d8261132c565b6112ff565b91508082526020830160208301858383011115610abe57600080fd5b610ac98382846113f6565b50505092915050565b6000610ade82356113d0565b905092915050565b6000610af282516113d0565b905092915050565b6000610b0682356113da565b905092915050565b600060208284031215610b2057600080fd5b6000610b2e848285016109fe565b91505092915050565b600080600060608486031215610b4c57600080fd5b6000610b5a868287016109fe565b9350506020610b6b86828701610ae6565b925050604084015167ffffffffffffffff811115610b8857600080fd5b610b9486828701610a7c565b9150509250925092565b600060208284031215610bb057600080fd5b6000610bbe84828501610a12565b91505092915050565b60008060408385031215610bda57600080fd5b6000610be885828601610a12565b925050602083015167ffffffffffffffff811115610c0557600080fd5b610c1185828601610a7c565b9150509250929050565b600060208284031215610c2d57600080fd5b600082013567ffffffffffffffff811115610c4757600080fd5b610c5384828501610a26565b91505092915050565b60008060008060808587031215610c7257600080fd5b600085013567ffffffffffffffff811115610c8c57600080fd5b610c9887828801610a26565b945050602085013567ffffffffffffffff811115610cb557600080fd5b610cc187828801610a26565b935050604085013567ffffffffffffffff811115610cde57600080fd5b610cea87828801610a26565b925050606085013567ffffffffffffffff811115610d0757600080fd5b610d1387828801610a26565b91505092959194509250565b600080600080600060a08688031215610d3757600080fd5b600086013567ffffffffffffffff811115610d5157600080fd5b610d5d88828901610a26565b9550506020610d6e88828901610ad2565b945050604086013567ffffffffffffffff811115610d8b57600080fd5b610d9788828901610a26565b935050606086013567ffffffffffffffff811115610db457600080fd5b610dc088828901610a26565b9250506080610dd188828901610ad2565b9150509295509295909350565b600080600080600060a08688031215610df657600080fd5b600086013567ffffffffffffffff811115610e1057600080fd5b610e1c88828901610a26565b9550506020610e2d88828901610afa565b945050604086013567ffffffffffffffff811115610e4a57600080fd5b610e5688828901610a26565b935050606086013567ffffffffffffffff811115610e7357600080fd5b610e7f88828901610a26565b9250506080610e90888289016109ea565b9150509295509295909350565b610ea681611363565b82525050565b610eb58161136f565b82525050565b610ec481611379565b82525050565b6000610ed582611358565b808452610ee98160208601602086016113f6565b610ef281611429565b602085010191505092915050565b6000601982527f496e76616c69642068657820737472696e67206c656e677468000000000000006020830152604082019050919050565b6000601582527f496e76616c6964206865782063686172616374657200000000000000000000006020830152604082019050919050565b6000602282527f4e6f64654944206e6f74207265676973746572656420696e207468652074616260208301527f6c650000000000000000000000000000000000000000000000000000000000006040830152606082019050919050565b6000600982527f5075626c69634b657900000000000000000000000000000000000000000000006020830152604082019050919050565b6000602982527f4661696c656420746f207570646174652074727573742073636f726520696e2060208301527f746865207461626c6500000000000000000000000000000000000000000000006040830152606082019050919050565b61106881611383565b82525050565b6110778161138d565b82525050565b60006020820190506110926000830184610e9d565b92915050565b60006060820190506110ad6000830186610e9d565b6110ba602083018561105f565b81810360408301526110cc8184610eca565b9050949350505050565b60006020820190506110eb6000830184610eac565b92915050565b60006080820190506111066000830187610eac565b611113602083018661106e565b6111206040830185610eac565b61112d6060830184610eac565b95945050505050565b600060208201905061114b6000830184610ebb565b92915050565b6000608082019050818103600083015261116b8187610eca565b9050818103602083015261117f8186610eca565b905081810360408301526111938185610eca565b905081810360608301526111a78184610eca565b905095945050505050565b600060408201905081810360008301526111cc8184610eca565b905081810360208301526111df81610fcb565b905092915050565b600060408201905081810360008301526112018185610eca565b9050611210602083018461105f565b9392505050565b600060a08201905081810360008301526112318188610eca565b9050611240602083018761105f565b81810360408301526112528186610eca565b905081810360608301526112668185610eca565b9050611275608083018461105f565b9695505050505050565b6000602082019050818103600083015261129881610f00565b9050919050565b600060208201905081810360008301526112b881610f37565b9050919050565b600060208201905081810360008301526112d881610f6e565b9050919050565b600060208201905081810360008301526112f881611002565b9050919050565b6000604051905081810181811067ffffffffffffffff8211171561132257600080fd5b8060405250919050565b600067ffffffffffffffff82111561134357600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b60008115159050919050565b6000819050919050565b6000819050919050565b6000819050919050565b600060ff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60008115159050919050565b6000819050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b838110156114145780820151818401526020810190506113f9565b83811115611423576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a72305820701a1030adf6f4bbe93617096dba0be74dbc582b274b8ff87d63d2276f8789936c6578706572696d656e74616cf50037"

// DeployDRoT deploys a new contract, binding an instance of DRoT to it.
func DeployDRoT(auth *bind.TransactOpts, backend bind.ContractBackend, tableRegistrationAddress common.Address) (common.Address, *types.Transaction, *DRoT, error) {
	parsed, err := abi.JSON(strings.NewReader(DRoTABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DRoTBin), backend, tableRegistrationAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DRoT{DRoTCaller: DRoTCaller{contract: contract}, DRoTTransactor: DRoTTransactor{contract: contract}, DRoTFilterer: DRoTFilterer{contract: contract}}, nil
}

func AsyncDeployDRoT(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, tableRegistrationAddress common.Address) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(DRoTABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(DRoTBin), backend, tableRegistrationAddress)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// DRoT is an auto generated Go binding around a Solidity contract.
type DRoT struct {
	DRoTCaller     // Read-only binding to the contract
	DRoTTransactor // Write-only binding to the contract
	DRoTFilterer   // Log filterer for contract events
}

// DRoTCaller is an auto generated read-only Go binding around a Solidity contract.
type DRoTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DRoTTransactor is an auto generated write-only Go binding around a Solidity contract.
type DRoTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DRoTFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type DRoTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DRoTSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type DRoTSession struct {
	Contract     *DRoT             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DRoTCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type DRoTCallerSession struct {
	Contract *DRoTCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DRoTTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type DRoTTransactorSession struct {
	Contract     *DRoTTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DRoTRaw is an auto generated low-level Go binding around a Solidity contract.
type DRoTRaw struct {
	Contract *DRoT // Generic contract binding to access the raw methods on
}

// DRoTCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type DRoTCallerRaw struct {
	Contract *DRoTCaller // Generic read-only contract binding to access the raw methods on
}

// DRoTTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type DRoTTransactorRaw struct {
	Contract *DRoTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDRoT creates a new instance of DRoT, bound to a specific deployed contract.
func NewDRoT(address common.Address, backend bind.ContractBackend) (*DRoT, error) {
	contract, err := bindDRoT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DRoT{DRoTCaller: DRoTCaller{contract: contract}, DRoTTransactor: DRoTTransactor{contract: contract}, DRoTFilterer: DRoTFilterer{contract: contract}}, nil
}

// NewDRoTCaller creates a new read-only instance of DRoT, bound to a specific deployed contract.
func NewDRoTCaller(address common.Address, caller bind.ContractCaller) (*DRoTCaller, error) {
	contract, err := bindDRoT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DRoTCaller{contract: contract}, nil
}

// NewDRoTTransactor creates a new write-only instance of DRoT, bound to a specific deployed contract.
func NewDRoTTransactor(address common.Address, transactor bind.ContractTransactor) (*DRoTTransactor, error) {
	contract, err := bindDRoT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DRoTTransactor{contract: contract}, nil
}

// NewDRoTFilterer creates a new log filterer instance of DRoT, bound to a specific deployed contract.
func NewDRoTFilterer(address common.Address, filterer bind.ContractFilterer) (*DRoTFilterer, error) {
	contract, err := bindDRoT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DRoTFilterer{contract: contract}, nil
}

// bindDRoT binds a generic wrapper to an already deployed contract.
func bindDRoT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DRoTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DRoT *DRoTRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DRoT.Contract.DRoTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DRoT *DRoTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _DRoT.Contract.DRoTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DRoT *DRoTRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _DRoT.Contract.DRoTTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DRoT *DRoTCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DRoT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DRoT *DRoTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _DRoT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DRoT *DRoTTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _DRoT.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// EcdsaVerifySol is a free data retrieval call binding the contract method 0xada76d58.
//
// Solidity: function EcdsaVerifySol(string messageHash, uint8 v, string r, string s, address expectedSigner) constant returns(bool)
func (_DRoT *DRoTCaller) EcdsaVerifySol(opts *bind.CallOpts, messageHash string, v uint8, r string, s string, expectedSigner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DRoT.contract.Call(opts, out, "EcdsaVerifySol", messageHash, v, r, s, expectedSigner)
	return *ret0, err
}

// EcdsaVerifySol is a free data retrieval call binding the contract method 0xada76d58.
//
// Solidity: function EcdsaVerifySol(string messageHash, uint8 v, string r, string s, address expectedSigner) constant returns(bool)
func (_DRoT *DRoTSession) EcdsaVerifySol(messageHash string, v uint8, r string, s string, expectedSigner common.Address) (bool, error) {
	return _DRoT.Contract.EcdsaVerifySol(&_DRoT.CallOpts, messageHash, v, r, s, expectedSigner)
}

// EcdsaVerifySol is a free data retrieval call binding the contract method 0xada76d58.
//
// Solidity: function EcdsaVerifySol(string messageHash, uint8 v, string r, string s, address expectedSigner) constant returns(bool)
func (_DRoT *DRoTCallerSession) EcdsaVerifySol(messageHash string, v uint8, r string, s string, expectedSigner common.Address) (bool, error) {
	return _DRoT.Contract.EcdsaVerifySol(&_DRoT.CallOpts, messageHash, v, r, s, expectedSigner)
}

// VerifyEcdsa is a free data retrieval call binding the contract method 0x21d4f02d.
//
// Solidity: function VerifyEcdsa(string message, string rHex, string sHex, string pubKeyHex) constant returns(bool)
func (_DRoT *DRoTCaller) VerifyEcdsa(opts *bind.CallOpts, message string, rHex string, sHex string, pubKeyHex string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DRoT.contract.Call(opts, out, "VerifyEcdsa", message, rHex, sHex, pubKeyHex)
	return *ret0, err
}

// VerifyEcdsa is a free data retrieval call binding the contract method 0x21d4f02d.
//
// Solidity: function VerifyEcdsa(string message, string rHex, string sHex, string pubKeyHex) constant returns(bool)
func (_DRoT *DRoTSession) VerifyEcdsa(message string, rHex string, sHex string, pubKeyHex string) (bool, error) {
	return _DRoT.Contract.VerifyEcdsa(&_DRoT.CallOpts, message, rHex, sHex, pubKeyHex)
}

// VerifyEcdsa is a free data retrieval call binding the contract method 0x21d4f02d.
//
// Solidity: function VerifyEcdsa(string message, string rHex, string sHex, string pubKeyHex) constant returns(bool)
func (_DRoT *DRoTCallerSession) VerifyEcdsa(message string, rHex string, sHex string, pubKeyHex string) (bool, error) {
	return _DRoT.Contract.VerifyEcdsa(&_DRoT.CallOpts, message, rHex, sHex, pubKeyHex)
}

// HexStringToBytes32 is a free data retrieval call binding the contract method 0x8b6bb00f.
//
// Solidity: function hexStringToBytes32(string hexString) constant returns(bytes32)
func (_DRoT *DRoTCaller) HexStringToBytes32(opts *bind.CallOpts, hexString string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DRoT.contract.Call(opts, out, "hexStringToBytes32", hexString)
	return *ret0, err
}

// HexStringToBytes32 is a free data retrieval call binding the contract method 0x8b6bb00f.
//
// Solidity: function hexStringToBytes32(string hexString) constant returns(bytes32)
func (_DRoT *DRoTSession) HexStringToBytes32(hexString string) ([32]byte, error) {
	return _DRoT.Contract.HexStringToBytes32(&_DRoT.CallOpts, hexString)
}

// HexStringToBytes32 is a free data retrieval call binding the contract method 0x8b6bb00f.
//
// Solidity: function hexStringToBytes32(string hexString) constant returns(bytes32)
func (_DRoT *DRoTCallerSession) HexStringToBytes32(hexString string) ([32]byte, error) {
	return _DRoT.Contract.HexStringToBytes32(&_DRoT.CallOpts, hexString)
}

// VerifyAR is a free data retrieval call binding the contract method 0x86848a91.
//
// Solidity: function verifyAR(string nodeID, uint256 attestationTime, string measurement, string referenceMeasurement, uint256 reliabilityThreshold) constant returns(bool, uint256, string)
func (_DRoT *DRoTCaller) VerifyAR(opts *bind.CallOpts, nodeID string, attestationTime *big.Int, measurement string, referenceMeasurement string, reliabilityThreshold *big.Int) (bool, *big.Int, string, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _DRoT.contract.Call(opts, out, "verifyAR", nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold)
	return *ret0, *ret1, *ret2, err
}

// VerifyAR is a free data retrieval call binding the contract method 0x86848a91.
//
// Solidity: function verifyAR(string nodeID, uint256 attestationTime, string measurement, string referenceMeasurement, uint256 reliabilityThreshold) constant returns(bool, uint256, string)
func (_DRoT *DRoTSession) VerifyAR(nodeID string, attestationTime *big.Int, measurement string, referenceMeasurement string, reliabilityThreshold *big.Int) (bool, *big.Int, string, error) {
	return _DRoT.Contract.VerifyAR(&_DRoT.CallOpts, nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold)
}

// VerifyAR is a free data retrieval call binding the contract method 0x86848a91.
//
// Solidity: function verifyAR(string nodeID, uint256 attestationTime, string measurement, string referenceMeasurement, uint256 reliabilityThreshold) constant returns(bool, uint256, string)
func (_DRoT *DRoTCallerSession) VerifyAR(nodeID string, attestationTime *big.Int, measurement string, referenceMeasurement string, reliabilityThreshold *big.Int) (bool, *big.Int, string, error) {
	return _DRoT.Contract.VerifyAR(&_DRoT.CallOpts, nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold)
}

// VerifyARandStore is a paid mutator transaction binding the contract method 0x8c67af27.
//
// Solidity: function verifyARandStore(string nodeID, uint256 attestationTime, string measurement, string referenceMeasurement, uint256 reliabilityThreshold) returns(bool, uint256, string)
func (_DRoT *DRoTTransactor) VerifyARandStore(opts *bind.TransactOpts, nodeID string, attestationTime *big.Int, measurement string, referenceMeasurement string, reliabilityThreshold *big.Int) (bool, *big.Int, string, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	transaction, receipt, err := _DRoT.contract.TransactWithResult(opts, out, "verifyARandStore", nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold)
	return *ret0, *ret1, *ret2, transaction, receipt, err
}

func (_DRoT *DRoTTransactor) AsyncVerifyARandStore(handler func(*types.Receipt, error), opts *bind.TransactOpts, nodeID string, attestationTime *big.Int, measurement string, referenceMeasurement string, reliabilityThreshold *big.Int) (*types.Transaction, error) {
	return _DRoT.contract.AsyncTransact(opts, handler, "verifyARandStore", nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold)
}

// VerifyARandStore is a paid mutator transaction binding the contract method 0x8c67af27.
//
// Solidity: function verifyARandStore(string nodeID, uint256 attestationTime, string measurement, string referenceMeasurement, uint256 reliabilityThreshold) returns(bool, uint256, string)
func (_DRoT *DRoTSession) VerifyARandStore(nodeID string, attestationTime *big.Int, measurement string, referenceMeasurement string, reliabilityThreshold *big.Int) (bool, *big.Int, string, *types.Transaction, *types.Receipt, error) {
	return _DRoT.Contract.VerifyARandStore(&_DRoT.TransactOpts, nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold)
}

func (_DRoT *DRoTSession) AsyncVerifyARandStore(handler func(*types.Receipt, error), nodeID string, attestationTime *big.Int, measurement string, referenceMeasurement string, reliabilityThreshold *big.Int) (*types.Transaction, error) {
	return _DRoT.Contract.AsyncVerifyARandStore(handler, &_DRoT.TransactOpts, nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold)
}

// VerifyARandStore is a paid mutator transaction binding the contract method 0x8c67af27.
//
// Solidity: function verifyARandStore(string nodeID, uint256 attestationTime, string measurement, string referenceMeasurement, uint256 reliabilityThreshold) returns(bool, uint256, string)
func (_DRoT *DRoTTransactorSession) VerifyARandStore(nodeID string, attestationTime *big.Int, measurement string, referenceMeasurement string, reliabilityThreshold *big.Int) (bool, *big.Int, string, *types.Transaction, *types.Receipt, error) {
	return _DRoT.Contract.VerifyARandStore(&_DRoT.TransactOpts, nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold)
}

func (_DRoT *DRoTTransactorSession) AsyncVerifyARandStore(handler func(*types.Receipt, error), nodeID string, attestationTime *big.Int, measurement string, referenceMeasurement string, reliabilityThreshold *big.Int) (*types.Transaction, error) {
	return _DRoT.Contract.AsyncVerifyARandStore(handler, &_DRoT.TransactOpts, nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold)
}

// DRoTUpdateStatusResult represents a UpdateStatusResult event raised by the DRoT contract.
type DRoTUpdateStatusResult struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchUpdateStatusResult is a free log subscription operation binding the contract event 0x056d4ff851ef14b8d2a7611d6bcd01b76e5d5d03485176f7aa975576098c5d05.
//
// Solidity: event UpdateStatusResult(int256 count)
func (_DRoT *DRoTFilterer) WatchUpdateStatusResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _DRoT.contract.WatchLogs(fromBlock, handler, "UpdateStatusResult")
}

func (_DRoT *DRoTFilterer) WatchAllUpdateStatusResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _DRoT.contract.WatchLogs(fromBlock, handler, "UpdateStatusResult")
}

// ParseUpdateStatusResult is a log parse operation binding the contract event 0x056d4ff851ef14b8d2a7611d6bcd01b76e5d5d03485176f7aa975576098c5d05.
//
// Solidity: event UpdateStatusResult(int256 count)
func (_DRoT *DRoTFilterer) ParseUpdateStatusResult(log types.Log) (*DRoTUpdateStatusResult, error) {
	event := new(DRoTUpdateStatusResult)
	if err := _DRoT.contract.UnpackLog(event, "UpdateStatusResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchUpdateStatusResult is a free log subscription operation binding the contract event 0x056d4ff851ef14b8d2a7611d6bcd01b76e5d5d03485176f7aa975576098c5d05.
//
// Solidity: event UpdateStatusResult(int256 count)
func (_DRoT *DRoTSession) WatchUpdateStatusResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _DRoT.Contract.WatchUpdateStatusResult(fromBlock, handler)
}

func (_DRoT *DRoTSession) WatchAllUpdateStatusResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _DRoT.Contract.WatchAllUpdateStatusResult(fromBlock, handler)
}

// ParseUpdateStatusResult is a log parse operation binding the contract event 0x056d4ff851ef14b8d2a7611d6bcd01b76e5d5d03485176f7aa975576098c5d05.
//
// Solidity: event UpdateStatusResult(int256 count)
func (_DRoT *DRoTSession) ParseUpdateStatusResult(log types.Log) (*DRoTUpdateStatusResult, error) {
	return _DRoT.Contract.ParseUpdateStatusResult(log)
}
