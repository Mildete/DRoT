package Contract_API

import (
	"crypto/ecdsa"
	"fmt"
	DRoT_abi "github.com/FISCO-BCOS/go-sdk/Go_Off-chain/Contract-API/Sol_Abi"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"os"
	"strings"
)

func AR_Verify(nodeID string, attestationTime *big.Int, measurement, referenceMeasurement string, reliabilityThreshold *big.Int) (bool, *big.Int) {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]
	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}

	// load the contract
	addressfilePath := "./contract/ContractAddress/DRoTPrecompiled.txt"
	// 读取文件内容
	data, err := os.ReadFile(addressfilePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
	}
	// 去除可能的换行符或空白字符
	contractAddresshex := strings.TrimSpace(string(data))
	contractAddress := common.HexToAddress(contractAddresshex) // 此处填写部署后16进制的合约地址
	instance, err := DRoT_abi.NewDRoT(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	ARverifySession := &DRoT_abi.DRoTSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	// call AR API
	verificationResult, trustScore, verificationInfo, err := ARverifySession.VerifyAR(nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold) // call select API   返回值写在sol里，根据对应函数的返回值接数据
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("verificationResult:", verificationResult)
	//fmt.Println("trustScore:", trustScore)
	fmt.Println("verificationInfo:", verificationInfo)
	return verificationResult, trustScore
}
func Ecdsa_Verify(message string, r, s string, signer string) bool {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]
	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}

	// load the contract
	// 合约路径
	addressfilePath := "./contract/ContractAddress/DRoTPrecompiled.txt"
	// 读取文件内容
	data, err := os.ReadFile(addressfilePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
	}
	// 去除可能的换行符或空白字符
	contractAddresshex := strings.TrimSpace(string(data))
	contractAddress := common.HexToAddress(contractAddresshex) // 此处填写部署后16进制的合约地址

	// load the contract
	//contractAddress := common.HexToAddress("0x32CA0eEe658Acd8d92FF68EdAf991B818b6870b2") // 此处填写部署后16进制的合约地址

	instance, err := DRoT_abi.NewDRoT(contractAddress, client)
	//instance, err := registration.NewRegistration(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	ecdsaverifySession := &DRoT_abi.DRoTSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	// call verifySignature API
	ret, err := ecdsaverifySession.VerifyEcdsa(message, r, s, signer) // call select API   返回值写在sol里，根据对应函数的返回值接数据
	if err != nil {
		log.Fatal(err)
	}
	if ret {
		fmt.Println("API: On_chain Signature Verified successfully")
	} else {
		fmt.Println("API: On_chain Signature Verified Failed")
	}
	return ret
	//fmt.Printf("transaction hash of receipt: %s\n", receipt.GetTransactionHash())
}

func Ecdsa_VerifySol(messagehash string, v uint8, r, s string, signer common.Address) bool {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]
	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}

	// load the contract
	// 合约路径
	addressfilePath := "./contract/ContractAddress/DRoTPrecompiled.txt"
	// 读取文件内容
	data, err := os.ReadFile(addressfilePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
	}
	// 去除可能的换行符或空白字符
	contractAddresshex := strings.TrimSpace(string(data))
	contractAddress := common.HexToAddress(contractAddresshex) // 此处填写部署后16进制的合约地址

	instance, err := DRoT_abi.NewDRoT(contractAddress, client)
	ecdsaverifysolSession := &DRoT_abi.DRoTSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	ret, err := ecdsaverifysolSession.EcdsaVerifySol(messagehash, v, r, s, signer) // call select API   返回值写在sol里，根据对应函数的返回值接数据
	if err != nil {
		log.Fatal(err)
	}
	if ret {
		fmt.Println("On_chain Signature Verified successfully")
	} else {
		fmt.Println("On_chain Signature Verified Failed")
	}
	return ret
}

func Register(NodeID string, PublicKeyAddress *ecdsa.PublicKey, AttestationScoreFunc string, ReferenceMeasurement string) *big.Int {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]
	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}

	// load the contract
	// 合约路径
	addressfilePath := "./contract/ContractAddress/DRoTRegistration.txt"
	// 读取文件内容
	data, err := os.ReadFile(addressfilePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
	}
	// 去除可能的换行符或空白字符
	contractAddresshex := strings.TrimSpace(string(data))
	contractAddress := common.HexToAddress(contractAddresshex) // 此处填写部署后16进制的合约地址

	instance, err := DRoT_abi.NewRegistration(contractAddress, client)
	registrationSession := &DRoT_abi.RegistrationSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	ret, _, _, err := registrationSession.Register(NodeID, crypto.PubkeyToAddress(*PublicKeyAddress).Hex(), AttestationScoreFunc, ReferenceMeasurement) // call select API   返回值写在sol里，根据对应函数的返回值接数据
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("状态码", ret) //nil：成功 -1：节点已存在  -2：失败? 无权限或者其他错误
	return ret
}

func Select(NodeID, field string) (*big.Int, string) {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]
	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}

	// load the contract
	// 合约路径
	addressfilePath := "./contract/ContractAddress/DRoTRegistration.txt"
	// 读取文件内容
	data, err := os.ReadFile(addressfilePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
	}
	// 去除可能的换行符或空白字符
	contractAddresshex := strings.TrimSpace(string(data))
	contractAddress := common.HexToAddress(contractAddresshex) // 此处填写部署后16进制的合约地址

	instance, err := DRoT_abi.NewRegistration(contractAddress, client)
	registrationSession := &DRoT_abi.RegistrationSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	ret, result, err := registrationSession.Select(NodeID, field) // call select API   返回值写在sol里，根据对应函数的返回值接数据
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("状态码", ret) //nil：成功 -1：节点已存在  -2：失败? 无权限或者其他错误
	fmt.Println("查询id"+NodeID+"_字段 "+field+"：", result)
	return ret, result
}

func AR_Verify_Store(nodeID string, attestationTime *big.Int, measurement, referenceMeasurement string, reliabilityThreshold *big.Int) (bool, *big.Int) {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]
	client, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}

	// load the contract
	// 合约路径
	addressfilePath := "./contract/ContractAddress/DRoTPrecompiled.txt"
	// 读取文件内容
	data, err := os.ReadFile(addressfilePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
	}
	// 去除可能的换行符或空白字符
	contractAddresshex := strings.TrimSpace(string(data))
	contractAddress := common.HexToAddress(contractAddresshex) // 此处填写部署后16进制的合约地址
	instance, err := DRoT_abi.NewDRoT(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	ARverifySession := &DRoT_abi.DRoTSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}

	// call AR API
	verificationResult, trustScore, verificationInfo, _, _, err := ARverifySession.VerifyARandStore(nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold) // call select API   返回值写在sol里，根据对应函数的返回值接数据
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("verificationInfo:", verificationInfo)
	return verificationResult, trustScore
}
