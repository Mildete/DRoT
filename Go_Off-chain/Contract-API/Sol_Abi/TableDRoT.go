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

// RegistrationABI is the input ABI used to generate the binding from.
const RegistrationABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"NodeID\",\"type\":\"string\"},{\"name\":\"PublicKeyAddress\",\"type\":\"string\"},{\"name\":\"AttestationScoreFunc\",\"type\":\"string\"},{\"name\":\"ReferenceMeasurement\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"NodeID\",\"type\":\"string\"},{\"name\":\"field\",\"type\":\"string\"}],\"name\":\"Select\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"NodeID\",\"type\":\"string\"},{\"name\":\"PublicKeyAddress\",\"type\":\"string\"}],\"name\":\"remove\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"nodeID\",\"type\":\"string\"},{\"name\":\"trustscore\",\"type\":\"uint256\"}],\"name\":\"updateTrustScore\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"ret\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"NodeID\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"PublicKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"AttestationScoreFunc\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"ReferenceMeasurement\",\"type\":\"string\"}],\"name\":\"RegisterEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"RemoveResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"UpdateStatusResult\",\"type\":\"event\"}]"

// RegistrationBin is the compiled bytecode used for deploying new contracts.
var RegistrationBin = "0x60806040523480156200001157600080fd5b506200002b62000031640100000000026401000000009004565b6200021e565b600061101090508073ffffffffffffffffffffffffffffffffffffffff166356004b6a6040805190810160405280600f81526020017f44526f545f5461626c655f5465737400000000000000000000000000000000008152506040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018060200180602001848103845285818151815260200191508051906020019080838360005b83811015620000ff578082015181840152602081019050620000e2565b50505050905090810190601f1680156200012d5780820380516001836020036101000a031916815260200191505b50848103835260068152602001807f4e6f6465494400000000000000000000000000000000000000000000000000008152506020018481038252603e8152602001807f5075626c69634b65792c4174746573746174696f6e53636f726546756e632c5281526020017f65666572656e63654d6561737572656d656e742c547275737453636f72650000815250604001945050505050602060405180830381600087803b158015620001dd57600080fd5b505af1158015620001f2573d6000803e3d6000fd5b505050506040513d60208110156200020957600080fd5b81019080805190602001909291905050505050565b611e56806200022e6000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630e24c52c1461006757806322cd843f146101b657806344590a7e146102e5578063d25a06f4146103a8575b600080fd5b34801561007357600080fd5b506101a0600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061042f565b6040518082815260200191505060405180910390f35b3480156101c257600080fd5b50610263600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610e33565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156102a957808201518184015260208101905061028e565b50505050905090810190601f1680156102d65780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b3480156102f157600080fd5b50610392600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611359565b6040518082815260200191505060405180910390f35b3480156103b457600080fd5b50610419600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001909291905050506117dc565b6040518082815260200191505060405180910390f35b600080600060606000806000809550600094506104818b6040805190810160405280600981526020017f5075626c69634b65790000000000000000000000000000000000000000000000815250610e33565b8095508196505050600085141515610c135761049b611cd2565b92508273ffffffffffffffffffffffffffffffffffffffff166313db93466040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561050157600080fd5b505af1158015610515573d6000803e3d6000fd5b505050506040513d602081101561052b57600080fd5b810190808051906020019092919050505091508173ffffffffffffffffffffffffffffffffffffffff1663e942b5168c6040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001838103835260068152602001807f4e6f646549440000000000000000000000000000000000000000000000000000815250602001838103825284818151815260200191508051906020019080838360005b838110156105fe5780820151818401526020810190506105e3565b50505050905090810190601f16801561062b5780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b15801561064b57600080fd5b505af115801561065f573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff1663e942b5168b6040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001838103835260098152602001807f5075626c69634b65790000000000000000000000000000000000000000000000815250602001838103825284818151815260200191508051906020019080838360005b83811015610723578082015181840152602081019050610708565b50505050905090810190601f1680156107505780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b15801561077057600080fd5b505af1158015610784573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff1663e942b5168a6040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001838103835260148152602001807f4174746573746174696f6e53636f726546756e63000000000000000000000000815250602001838103825284818151815260200191508051906020019080838360005b8381101561084857808201518184015260208101905061082d565b50505050905090810190601f1680156108755780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b15801561089557600080fd5b505af11580156108a9573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff1663e942b516896040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001838103835260148152602001807f5265666572656e63654d6561737572656d656e74000000000000000000000000815250602001838103825284818151815260200191508051906020019080838360005b8381101561096d578082015181840152602081019050610952565b50505050905090810190601f16801561099a5780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b1580156109ba57600080fd5b505af11580156109ce573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff16638a42ebe960006040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018381526020018281038252600a8152602001807f547275737453636f72650000000000000000000000000000000000000000000081525060200192505050600060405180830381600087803b158015610a7b57600080fd5b505af1158015610a8f573d6000803e3d6000fd5b505050508273ffffffffffffffffffffffffffffffffffffffff166331afac368c846040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b83811015610b4e578082015181840152602081019050610b33565b50505050905090810190601f168015610b7b5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015610b9b57600080fd5b505af1158015610baf573d6000803e3d6000fd5b505050506040513d6020811015610bc557600080fd5b810190808051906020019092919050505090506001811415610bea5760009550610c0e565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe95505b610c37565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff95505b7f42767499ffab62653527a1b68b7c4c2a6832d72c0daaa41eb3a4ece5f5dbf163868c8c8c8c6040518086815260200180602001806020018060200180602001858103855289818151815260200191508051906020019080838360005b83811015610caf578082015181840152602081019050610c94565b50505050905090810190601f168015610cdc5780820380516001836020036101000a031916815260200191505b50858103845288818151815260200191508051906020019080838360005b83811015610d15578082015181840152602081019050610cfa565b50505050905090810190601f168015610d425780820380516001836020036101000a031916815260200191505b50858103835287818151815260200191508051906020019080838360005b83811015610d7b578082015181840152602081019050610d60565b50505050905090810190601f168015610da85780820380516001836020036101000a031916815260200191505b50858103825286818151815260200191508051906020019080838360005b83811015610de1578082015181840152602081019050610dc6565b50505050905090810190601f168015610e0e5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390a1859650505050505050949350505050565b600060606000806000610e44611cd2565b92508273ffffffffffffffffffffffffffffffffffffffff1663e8434e39888573ffffffffffffffffffffffffffffffffffffffff16637857d7c96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610ec757600080fd5b505af1158015610edb573d6000803e3d6000fd5b505050506040513d6020811015610ef157600080fd5b81019080805190602001909291905050506040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b83811015610f9f578082015181840152602081019050610f84565b50505050905090810190601f168015610fcc5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b158015610fec57600080fd5b505af1158015611000573d6000803e3d6000fd5b505050506040513d602081101561101657600080fd5b810190808051906020019092919050505091508173ffffffffffffffffffffffffffffffffffffffff1663949d225d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561108d57600080fd5b505af11580156110a1573d6000803e3d6000fd5b505050506040513d60208110156110b757600080fd5b810190808051906020019092919050505060001415611133577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8090506040805190810160405280601681526020017f4e6f6465206973206e6f742072656769737465726564000000000000000000008152509450945061134f565b8173ffffffffffffffffffffffffffffffffffffffff1663846719e060006040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b1580156111a357600080fd5b505af11580156111b7573d6000803e3d6000fd5b505050506040513d60208110156111cd57600080fd5b8101908080519060200190929190505050905060008173ffffffffffffffffffffffffffffffffffffffff16639c981fcb886040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561126a57808201518184015260208101905061124f565b50505050905090810190601f1680156112975780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b1580156112b657600080fd5b505af11580156112ca573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060208110156112f457600080fd5b81019080805164010000000081111561130c57600080fd5b8281019050602081018481111561132257600080fd5b815185600182028301116401000000008211171561133f57600080fd5b5050929190505050819150945094505b5050509250929050565b600080600080611367611cd2565b92508273ffffffffffffffffffffffffffffffffffffffff16637857d7c96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156113cd57600080fd5b505af11580156113e1573d6000803e3d6000fd5b505050506040513d60208110156113f757600080fd5b810190808051906020019092919050505091508173ffffffffffffffffffffffffffffffffffffffff1663cd30a1d1876040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001838103835260068152602001807f4e6f646549440000000000000000000000000000000000000000000000000000815250602001838103825284818151815260200191508051906020019080838360005b838110156114ca5780820151818401526020810190506114af565b50505050905090810190601f1680156114f75780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b15801561151757600080fd5b505af115801561152b573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff1663cd30a1d1866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001838103835260098152602001807f5075626c69634b65790000000000000000000000000000000000000000000000815250602001838103825284818151815260200191508051906020019080838360005b838110156115ef5780820151818401526020810190506115d4565b50505050905090810190601f16801561161c5780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b15801561163c57600080fd5b505af1158015611650573d6000803e3d6000fd5b505050508273ffffffffffffffffffffffffffffffffffffffff166328bb211787846040518363ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b8381101561170f5780820151818401526020810190506116f4565b50505050905090810190601f16801561173c5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561175c57600080fd5b505af1158015611770573d6000803e3d6000fd5b505050506040513d602081101561178657600080fd5b810190808051906020019092919050505090507f4b930e280fe29620bdff00c88155d46d6d82a39f45dd5c3ea114dc3157358112816040518082815260200191505060405180910390a180935050505092915050565b60008060008060006117ec611cd2565b93508373ffffffffffffffffffffffffffffffffffffffff166313db93466040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561185257600080fd5b505af1158015611866573d6000803e3d6000fd5b505050506040513d602081101561187c57600080fd5b810190808051906020019092919050505092508273ffffffffffffffffffffffffffffffffffffffff16638a42ebe9876040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018381526020018281038252600a8152602001807f547275737453636f72650000000000000000000000000000000000000000000081525060200192505050600060405180830381600087803b15801561193757600080fd5b505af115801561194b573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff16637857d7c96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156119b357600080fd5b505af11580156119c7573d6000803e3d6000fd5b505050506040513d60208110156119dd57600080fd5b810190808051906020019092919050505091508173ffffffffffffffffffffffffffffffffffffffff1663cd30a1d1886040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808060200180602001838103835260068152602001807f4e6f646549440000000000000000000000000000000000000000000000000000815250602001838103825284818151815260200191508051906020019080838360005b83811015611ab0578082015181840152602081019050611a95565b50505050905090810190601f168015611add5780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b158015611afd57600080fd5b505af1158015611b11573d6000803e3d6000fd5b505050508373ffffffffffffffffffffffffffffffffffffffff1663bf2b70a18885856040518463ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825285818151815260200191508051906020019080838360005b83811015611c03578082015181840152602081019050611be8565b50505050905090810190601f168015611c305780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b158015611c5157600080fd5b505af1158015611c65573d6000803e3d6000fd5b505050506040513d6020811015611c7b57600080fd5b810190808051906020019092919050505090507f056d4ff851ef14b8d2a7611d6bcd01b76e5d5d03485176f7aa975576098c5d05816040518082815260200191505060405180910390a18094505050505092915050565b600080600061100191508173ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600f81526020017f44526f545f5461626c655f5465737400000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611d99578082015181840152602081019050611d7e565b50505050905090810190601f168015611dc65780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b158015611de557600080fd5b505af1158015611df9573d6000803e3d6000fd5b505050506040513d6020811015611e0f57600080fd5b810190808051906020019092919050505090508092505050905600a165627a7a7230582023fca316cbec157d25ad0afac8455ab7d8e7b482474efab7822e9ec5e74d740b0029"

// DeployRegistration deploys a new contract, binding an instance of Registration to it.
func DeployRegistration(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registration, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RegistrationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registration{RegistrationCaller: RegistrationCaller{contract: contract}, RegistrationTransactor: RegistrationTransactor{contract: contract}, RegistrationFilterer: RegistrationFilterer{contract: contract}}, nil
}

func AsyncDeployRegistration(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrationABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(RegistrationBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Registration is an auto generated Go binding around a Solidity contract.
type Registration struct {
	RegistrationCaller     // Read-only binding to the contract
	RegistrationTransactor // Write-only binding to the contract
	RegistrationFilterer   // Log filterer for contract events
}

// RegistrationCaller is an auto generated read-only Go binding around a Solidity contract.
type RegistrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationTransactor is an auto generated write-only Go binding around a Solidity contract.
type RegistrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type RegistrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type RegistrationSession struct {
	Contract     *Registration     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrationCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type RegistrationCallerSession struct {
	Contract *RegistrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RegistrationTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type RegistrationTransactorSession struct {
	Contract     *RegistrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RegistrationRaw is an auto generated low-level Go binding around a Solidity contract.
type RegistrationRaw struct {
	Contract *Registration // Generic contract binding to access the raw methods on
}

// RegistrationCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type RegistrationCallerRaw struct {
	Contract *RegistrationCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrationTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type RegistrationTransactorRaw struct {
	Contract *RegistrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistration creates a new instance of Registration, bound to a specific deployed contract.
func NewRegistration(address common.Address, backend bind.ContractBackend) (*Registration, error) {
	contract, err := bindRegistration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registration{RegistrationCaller: RegistrationCaller{contract: contract}, RegistrationTransactor: RegistrationTransactor{contract: contract}, RegistrationFilterer: RegistrationFilterer{contract: contract}}, nil
}

// NewRegistrationCaller creates a new read-only instance of Registration, bound to a specific deployed contract.
func NewRegistrationCaller(address common.Address, caller bind.ContractCaller) (*RegistrationCaller, error) {
	contract, err := bindRegistration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationCaller{contract: contract}, nil
}

// NewRegistrationTransactor creates a new write-only instance of Registration, bound to a specific deployed contract.
func NewRegistrationTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrationTransactor, error) {
	contract, err := bindRegistration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationTransactor{contract: contract}, nil
}

// NewRegistrationFilterer creates a new log filterer instance of Registration, bound to a specific deployed contract.
func NewRegistrationFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrationFilterer, error) {
	contract, err := bindRegistration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrationFilterer{contract: contract}, nil
}

// bindRegistration binds a generic wrapper to an already deployed contract.
func bindRegistration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registration *RegistrationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registration.Contract.RegistrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registration *RegistrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Registration.Contract.RegistrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registration *RegistrationRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Registration.Contract.RegistrationTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registration *RegistrationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Registration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registration *RegistrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Registration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registration *RegistrationTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Registration.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Select is a free data retrieval call binding the contract method 0x22cd843f.
//
// Solidity: function Select(string NodeID, string field) constant returns(int256, string)
func (_Registration *RegistrationCaller) Select(opts *bind.CallOpts, NodeID string, field string) (*big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Registration.contract.Call(opts, out, "Select", NodeID, field)
	return *ret0, *ret1, err
}

// Select is a free data retrieval call binding the contract method 0x22cd843f.
//
// Solidity: function Select(string NodeID, string field) constant returns(int256, string)
func (_Registration *RegistrationSession) Select(NodeID string, field string) (*big.Int, string, error) {
	return _Registration.Contract.Select(&_Registration.CallOpts, NodeID, field)
}

// Select is a free data retrieval call binding the contract method 0x22cd843f.
//
// Solidity: function Select(string NodeID, string field) constant returns(int256, string)
func (_Registration *RegistrationCallerSession) Select(NodeID string, field string) (*big.Int, string, error) {
	return _Registration.Contract.Select(&_Registration.CallOpts, NodeID, field)
}

// Register is a paid mutator transaction binding the contract method 0x0e24c52c.
//
// Solidity: function register(string NodeID, string PublicKeyAddress, string AttestationScoreFunc, string ReferenceMeasurement) returns(int256)
func (_Registration *RegistrationTransactor) Register(opts *bind.TransactOpts, NodeID string, PublicKeyAddress string, AttestationScoreFunc string, ReferenceMeasurement string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _Registration.contract.TransactWithResult(opts, out, "register", NodeID, PublicKeyAddress, AttestationScoreFunc, ReferenceMeasurement)
	return *ret0, transaction, receipt, err
}

func (_Registration *RegistrationTransactor) AsyncRegister(handler func(*types.Receipt, error), opts *bind.TransactOpts, NodeID string, PublicKeyAddress string, AttestationScoreFunc string, ReferenceMeasurement string) (*types.Transaction, error) {
	return _Registration.contract.AsyncTransact(opts, handler, "register", NodeID, PublicKeyAddress, AttestationScoreFunc, ReferenceMeasurement)
}

// Register is a paid mutator transaction binding the contract method 0x0e24c52c.
//
// Solidity: function register(string NodeID, string PublicKeyAddress, string AttestationScoreFunc, string ReferenceMeasurement) returns(int256)
func (_Registration *RegistrationSession) Register(NodeID string, PublicKeyAddress string, AttestationScoreFunc string, ReferenceMeasurement string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Registration.Contract.Register(&_Registration.TransactOpts, NodeID, PublicKeyAddress, AttestationScoreFunc, ReferenceMeasurement)
}

func (_Registration *RegistrationSession) AsyncRegister(handler func(*types.Receipt, error), NodeID string, PublicKeyAddress string, AttestationScoreFunc string, ReferenceMeasurement string) (*types.Transaction, error) {
	return _Registration.Contract.AsyncRegister(handler, &_Registration.TransactOpts, NodeID, PublicKeyAddress, AttestationScoreFunc, ReferenceMeasurement)
}

// Register is a paid mutator transaction binding the contract method 0x0e24c52c.
//
// Solidity: function register(string NodeID, string PublicKeyAddress, string AttestationScoreFunc, string ReferenceMeasurement) returns(int256)
func (_Registration *RegistrationTransactorSession) Register(NodeID string, PublicKeyAddress string, AttestationScoreFunc string, ReferenceMeasurement string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Registration.Contract.Register(&_Registration.TransactOpts, NodeID, PublicKeyAddress, AttestationScoreFunc, ReferenceMeasurement)
}

func (_Registration *RegistrationTransactorSession) AsyncRegister(handler func(*types.Receipt, error), NodeID string, PublicKeyAddress string, AttestationScoreFunc string, ReferenceMeasurement string) (*types.Transaction, error) {
	return _Registration.Contract.AsyncRegister(handler, &_Registration.TransactOpts, NodeID, PublicKeyAddress, AttestationScoreFunc, ReferenceMeasurement)
}

// Remove is a paid mutator transaction binding the contract method 0x44590a7e.
//
// Solidity: function remove(string NodeID, string PublicKeyAddress) returns(int256)
func (_Registration *RegistrationTransactor) Remove(opts *bind.TransactOpts, NodeID string, PublicKeyAddress string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _Registration.contract.TransactWithResult(opts, out, "remove", NodeID, PublicKeyAddress)
	return *ret0, transaction, receipt, err
}

func (_Registration *RegistrationTransactor) AsyncRemove(handler func(*types.Receipt, error), opts *bind.TransactOpts, NodeID string, PublicKeyAddress string) (*types.Transaction, error) {
	return _Registration.contract.AsyncTransact(opts, handler, "remove", NodeID, PublicKeyAddress)
}

// Remove is a paid mutator transaction binding the contract method 0x44590a7e.
//
// Solidity: function remove(string NodeID, string PublicKeyAddress) returns(int256)
func (_Registration *RegistrationSession) Remove(NodeID string, PublicKeyAddress string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Registration.Contract.Remove(&_Registration.TransactOpts, NodeID, PublicKeyAddress)
}

func (_Registration *RegistrationSession) AsyncRemove(handler func(*types.Receipt, error), NodeID string, PublicKeyAddress string) (*types.Transaction, error) {
	return _Registration.Contract.AsyncRemove(handler, &_Registration.TransactOpts, NodeID, PublicKeyAddress)
}

// Remove is a paid mutator transaction binding the contract method 0x44590a7e.
//
// Solidity: function remove(string NodeID, string PublicKeyAddress) returns(int256)
func (_Registration *RegistrationTransactorSession) Remove(NodeID string, PublicKeyAddress string) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Registration.Contract.Remove(&_Registration.TransactOpts, NodeID, PublicKeyAddress)
}

func (_Registration *RegistrationTransactorSession) AsyncRemove(handler func(*types.Receipt, error), NodeID string, PublicKeyAddress string) (*types.Transaction, error) {
	return _Registration.Contract.AsyncRemove(handler, &_Registration.TransactOpts, NodeID, PublicKeyAddress)
}

// UpdateTrustScore is a paid mutator transaction binding the contract method 0xd25a06f4.
//
// Solidity: function updateTrustScore(string nodeID, uint256 trustscore) returns(int256)
func (_Registration *RegistrationTransactor) UpdateTrustScore(opts *bind.TransactOpts, nodeID string, trustscore *big.Int) (*big.Int, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	transaction, receipt, err := _Registration.contract.TransactWithResult(opts, out, "updateTrustScore", nodeID, trustscore)
	return *ret0, transaction, receipt, err
}

func (_Registration *RegistrationTransactor) AsyncUpdateTrustScore(handler func(*types.Receipt, error), opts *bind.TransactOpts, nodeID string, trustscore *big.Int) (*types.Transaction, error) {
	return _Registration.contract.AsyncTransact(opts, handler, "updateTrustScore", nodeID, trustscore)
}

// UpdateTrustScore is a paid mutator transaction binding the contract method 0xd25a06f4.
//
// Solidity: function updateTrustScore(string nodeID, uint256 trustscore) returns(int256)
func (_Registration *RegistrationSession) UpdateTrustScore(nodeID string, trustscore *big.Int) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Registration.Contract.UpdateTrustScore(&_Registration.TransactOpts, nodeID, trustscore)
}

func (_Registration *RegistrationSession) AsyncUpdateTrustScore(handler func(*types.Receipt, error), nodeID string, trustscore *big.Int) (*types.Transaction, error) {
	return _Registration.Contract.AsyncUpdateTrustScore(handler, &_Registration.TransactOpts, nodeID, trustscore)
}

// UpdateTrustScore is a paid mutator transaction binding the contract method 0xd25a06f4.
//
// Solidity: function updateTrustScore(string nodeID, uint256 trustscore) returns(int256)
func (_Registration *RegistrationTransactorSession) UpdateTrustScore(nodeID string, trustscore *big.Int) (*big.Int, *types.Transaction, *types.Receipt, error) {
	return _Registration.Contract.UpdateTrustScore(&_Registration.TransactOpts, nodeID, trustscore)
}

func (_Registration *RegistrationTransactorSession) AsyncUpdateTrustScore(handler func(*types.Receipt, error), nodeID string, trustscore *big.Int) (*types.Transaction, error) {
	return _Registration.Contract.AsyncUpdateTrustScore(handler, &_Registration.TransactOpts, nodeID, trustscore)
}

// RegistrationRegisterEvent represents a RegisterEvent event raised by the Registration contract.
type RegistrationRegisterEvent struct {
	Ret                  *big.Int
	NodeID               string
	PublicKey            string
	AttestationScoreFunc string
	ReferenceMeasurement string
	Raw                  types.Log // Blockchain specific contextual infos
}

// WatchRegisterEvent is a free log subscription operation binding the contract event 0x42767499ffab62653527a1b68b7c4c2a6832d72c0daaa41eb3a4ece5f5dbf163.
//
// Solidity: event RegisterEvent(int256 ret, string NodeID, string PublicKey, string AttestationScoreFunc, string ReferenceMeasurement)
func (_Registration *RegistrationFilterer) WatchRegisterEvent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _Registration.contract.WatchLogs(fromBlock, handler, "RegisterEvent")
}

func (_Registration *RegistrationFilterer) WatchAllRegisterEvent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _Registration.contract.WatchLogs(fromBlock, handler, "RegisterEvent")
}

// ParseRegisterEvent is a log parse operation binding the contract event 0x42767499ffab62653527a1b68b7c4c2a6832d72c0daaa41eb3a4ece5f5dbf163.
//
// Solidity: event RegisterEvent(int256 ret, string NodeID, string PublicKey, string AttestationScoreFunc, string ReferenceMeasurement)
func (_Registration *RegistrationFilterer) ParseRegisterEvent(log types.Log) (*RegistrationRegisterEvent, error) {
	event := new(RegistrationRegisterEvent)
	if err := _Registration.contract.UnpackLog(event, "RegisterEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchRegisterEvent is a free log subscription operation binding the contract event 0x42767499ffab62653527a1b68b7c4c2a6832d72c0daaa41eb3a4ece5f5dbf163.
//
// Solidity: event RegisterEvent(int256 ret, string NodeID, string PublicKey, string AttestationScoreFunc, string ReferenceMeasurement)
func (_Registration *RegistrationSession) WatchRegisterEvent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _Registration.Contract.WatchRegisterEvent(fromBlock, handler)
}

func (_Registration *RegistrationSession) WatchAllRegisterEvent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _Registration.Contract.WatchAllRegisterEvent(fromBlock, handler)
}

// ParseRegisterEvent is a log parse operation binding the contract event 0x42767499ffab62653527a1b68b7c4c2a6832d72c0daaa41eb3a4ece5f5dbf163.
//
// Solidity: event RegisterEvent(int256 ret, string NodeID, string PublicKey, string AttestationScoreFunc, string ReferenceMeasurement)
func (_Registration *RegistrationSession) ParseRegisterEvent(log types.Log) (*RegistrationRegisterEvent, error) {
	return _Registration.Contract.ParseRegisterEvent(log)
}

// RegistrationRemoveResult represents a RemoveResult event raised by the Registration contract.
type RegistrationRemoveResult struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchRemoveResult is a free log subscription operation binding the contract event 0x4b930e280fe29620bdff00c88155d46d6d82a39f45dd5c3ea114dc3157358112.
//
// Solidity: event RemoveResult(int256 count)
func (_Registration *RegistrationFilterer) WatchRemoveResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _Registration.contract.WatchLogs(fromBlock, handler, "RemoveResult")
}

func (_Registration *RegistrationFilterer) WatchAllRemoveResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _Registration.contract.WatchLogs(fromBlock, handler, "RemoveResult")
}

// ParseRemoveResult is a log parse operation binding the contract event 0x4b930e280fe29620bdff00c88155d46d6d82a39f45dd5c3ea114dc3157358112.
//
// Solidity: event RemoveResult(int256 count)
func (_Registration *RegistrationFilterer) ParseRemoveResult(log types.Log) (*RegistrationRemoveResult, error) {
	event := new(RegistrationRemoveResult)
	if err := _Registration.contract.UnpackLog(event, "RemoveResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchRemoveResult is a free log subscription operation binding the contract event 0x4b930e280fe29620bdff00c88155d46d6d82a39f45dd5c3ea114dc3157358112.
//
// Solidity: event RemoveResult(int256 count)
func (_Registration *RegistrationSession) WatchRemoveResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _Registration.Contract.WatchRemoveResult(fromBlock, handler)
}

func (_Registration *RegistrationSession) WatchAllRemoveResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _Registration.Contract.WatchAllRemoveResult(fromBlock, handler)
}

// ParseRemoveResult is a log parse operation binding the contract event 0x4b930e280fe29620bdff00c88155d46d6d82a39f45dd5c3ea114dc3157358112.
//
// Solidity: event RemoveResult(int256 count)
func (_Registration *RegistrationSession) ParseRemoveResult(log types.Log) (*RegistrationRemoveResult, error) {
	return _Registration.Contract.ParseRemoveResult(log)
}

// RegistrationUpdateStatusResult represents a UpdateStatusResult event raised by the Registration contract.
type RegistrationUpdateStatusResult struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// WatchUpdateStatusResult is a free log subscription operation binding the contract event 0x056d4ff851ef14b8d2a7611d6bcd01b76e5d5d03485176f7aa975576098c5d05.
//
// Solidity: event UpdateStatusResult(int256 count)
func (_Registration *RegistrationFilterer) WatchUpdateStatusResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _Registration.contract.WatchLogs(fromBlock, handler, "UpdateStatusResult")
}

func (_Registration *RegistrationFilterer) WatchAllUpdateStatusResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _Registration.contract.WatchLogs(fromBlock, handler, "UpdateStatusResult")
}

// ParseUpdateStatusResult is a log parse operation binding the contract event 0x056d4ff851ef14b8d2a7611d6bcd01b76e5d5d03485176f7aa975576098c5d05.
//
// Solidity: event UpdateStatusResult(int256 count)
func (_Registration *RegistrationFilterer) ParseUpdateStatusResult(log types.Log) (*RegistrationUpdateStatusResult, error) {
	event := new(RegistrationUpdateStatusResult)
	if err := _Registration.contract.UnpackLog(event, "UpdateStatusResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchUpdateStatusResult is a free log subscription operation binding the contract event 0x056d4ff851ef14b8d2a7611d6bcd01b76e5d5d03485176f7aa975576098c5d05.
//
// Solidity: event UpdateStatusResult(int256 count)
func (_Registration *RegistrationSession) WatchUpdateStatusResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _Registration.Contract.WatchUpdateStatusResult(fromBlock, handler)
}

func (_Registration *RegistrationSession) WatchAllUpdateStatusResult(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _Registration.Contract.WatchAllUpdateStatusResult(fromBlock, handler)
}

// ParseUpdateStatusResult is a log parse operation binding the contract event 0x056d4ff851ef14b8d2a7611d6bcd01b76e5d5d03485176f7aa975576098c5d05.
//
// Solidity: event UpdateStatusResult(int256 count)
func (_Registration *RegistrationSession) ParseUpdateStatusResult(log types.Log) (*RegistrationUpdateStatusResult, error) {
	return _Registration.Contract.ParseUpdateStatusResult(log)
}
