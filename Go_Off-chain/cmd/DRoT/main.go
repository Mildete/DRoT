package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	DRoTPrecompiled "github.com/FISCO-BCOS/go-sdk/Go_Off-chain/Contract-API"
	oracle "github.com/FISCO-BCOS/go-sdk/Go_Off-chain/Oracle"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"math/big"
	"strconv"
	"time"
)

// 定义通信开销的时延（毫秒）
const CommunicationDelay = 100 // 200 毫秒

func main() {

	// 1. Register Oracle Node
	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("Step 1: Register Oracle Node and generate/load keys")
	fmt.Println("************************************************************************************************************************************************")

	//生成ECDSA私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	//  将私钥和公钥保存到文件
	KeyPath := "blockchain-storage/Node_key/"

	//生成不一样的nodeID,以免因为注册问题导致measure不一致
	timestamp := time.Now().Unix()                   // 获取当前时间戳（秒级别）
	timestampStr := strconv.FormatInt(timestamp, 10) // 转换时间戳为字符串
	nodeID := "Node" + timestampStr                  // 拼接前缀和时间戳

	privateKeyPath := KeyPath + nodeID + "_private_key.txt"
	publicKeyPath := KeyPath + nodeID + "_public_key.txt"

	SavePrivateKey(privateKey, privateKeyPath)
	SavePublicKey(&privateKey.PublicKey, publicKeyPath)

	// 注册信息
	attestationScoreFun := "example-attestation-score-function"
	PublicKeyAddress := LoadPublicKey(publicKeyPath)
	initialMeasurement := oracle.GetMeasurement(nodeID, crypto.PubkeyToAddress(*PublicKeyAddress).Hex(), attestationScoreFun)

	// 计时开始
	start := time.Now() // 记录开始时间

	// ***模拟通信开销***请求注册
	time.Sleep(time.Duration(CommunicationDelay) * time.Millisecond)
	fmt.Println("注册请求，模拟通信延迟时间增加：", CommunicationDelay, "ms")

	//注册上链
	registrationRet := DRoTPrecompiled.Register(nodeID, PublicKeyAddress, attestationScoreFun, initialMeasurement)
	if registrationRet == nil { //nil：成功 -1：节点已存在  -2：失败? 无权限或者其他错误
		fmt.Println("registration succeeded!")
	} else {
		fmt.Println("registration failed!")
	}

	// ***模拟通信开销***注册合约返回
	time.Sleep(time.Duration(CommunicationDelay) * time.Millisecond)
	fmt.Println("注册结果返回，模拟通信延迟时间增加：", CommunicationDelay, "ms")

	//2. 生成 AR report

	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("Step 2: Generate an Attestation Report")
	fmt.Println("************************************************************************************************************************************************")

	arFilePath := "blockchain-storage/AR_Report/" + "oracle-" + "-AR.txt"
	loadedPrivateKey := LoadPrivateKey(KeyPath + nodeID + "_private_key.txt")

	referenceMeasurement := oracle.GetMeasurement(nodeID, crypto.PubkeyToAddress(*PublicKeyAddress).Hex(), attestationScoreFun)
	report, err := oracle.GenerateAttestationReport(nodeID, loadedPrivateKey, referenceMeasurement, arFilePath, 7)
	if err != nil {
		fmt.Printf("Error generating attestation report: %v\n", err)
		return
	}
	fmt.Println("Generated Attestation Report:")
	fmt.Printf("NodeID: %s\n", report.NodeID)
	fmt.Printf("RandomValue: %s\n", report.RandomValue)
	fmt.Printf("Proof: %s\n", report.Proof)
	fmt.Printf("AttestationTime: %d\n", report.AttestationTime)
	fmt.Printf("Measurement: %s\n", report.Measurement)
	//	fmt.Printf("Signature : %s\n", report.Signature)

	// 计时开始
	verifytimestart := time.Now() // 记录开始时间
	// ***模拟通信开销***发送AR报告
	time.Sleep(time.Duration(CommunicationDelay) * time.Millisecond)
	fmt.Println("发送AR报告，模拟通信延迟时间增加：", CommunicationDelay, "ms")

	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("Step 3: Verify Signature")
	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("ReportHash:", report.ReportHash)
	fmt.Println("V:", report.V)
	fmt.Println("R:", report.R)
	fmt.Println("S:", report.S)
	fmt.Println("PubkeyAddress:", report.PublickeyHex)

	if !DRoTPrecompiled.Ecdsa_VerifySol(report.ReportHash, report.V, report.R, report.S, report.PublickeyHex) {
		//if !DRoTPrecompiled.Ecdsa_VerifySol("0xb94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9", 27, "0x45f7a21f85010f395120c03b33881f2aa1b763b94b987db754f10db75fb17972", "0x48575c3132535166b900f07e9eabb36176da0c149a6b6f146b9260400066427d", common.HexToAddress("0x65e2aC1C561A5E5a71D953697737b486ED7B90Ab")) {
		fmt.Println("Signature is invalid!!")
		return
	} else {
		fmt.Println("Signature Verification Successful!")
	}

	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("Step 4: Verify Attestation Report and Get Trust Score")
	fmt.Println("************************************************************************************************************************************************")
	var num int64 = 80
	reliabilityThreshold := big.NewInt(num)

	verify_ret, trustScore := DRoTPrecompiled.AR_Verify_Store(report.NodeID, report.AttestationTime, report.Measurement, report.Measurement, reliabilityThreshold)
	fmt.Println("verificationResult:", verify_ret)
	fmt.Println("trustScore:", trustScore)

	fmt.Println("************************************************************************************************************************************************")
	fmt.Println("The End!!!")
	fmt.Println("************************************************************************************************************************************************")

	verifytimeend := time.Now()
	verifytimediff := verifytimeend.Sub(verifytimestart) // 计算开始和结束时间的差值
	fmt.Printf("verification时间: %s\n", verifytimediff)   // 打印运行时间

	end := time.Now()         // 记录结束时间
	elapsed := end.Sub(start) // 计算开始和结束时间的差值

	//查看链上信任分（操作时间未计入）
	DRoTPrecompiled.Select(nodeID, "TrustScore")

	fmt.Printf("DRoT运行时间: %s\n", elapsed)                                 // 打印运行时间
	fmt.Printf("除去verification（压测时间）外运行时间: %s\n", elapsed-verifytimediff) // 打印运行时间
}

// 保存私钥到文件
func SavePrivateKey(privateKey *ecdsa.PrivateKey, filename string) {
	privateKeyBytes := crypto.FromECDSA(privateKey)
	err := ioutil.WriteFile(filename, []byte(hex.EncodeToString(privateKeyBytes)), 0600)
	if err != nil {
		log.Fatal(err)
	}
}

// 从文件读取私钥
func LoadPrivateKey(filename string) *ecdsa.PrivateKey {
	privateKeyHex, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes, err := hex.DecodeString(string(privateKeyHex))
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}

// 保存公钥到文件
func SavePublicKey(publicKey *ecdsa.PublicKey, filename string) {
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	err := ioutil.WriteFile(filename, []byte(hex.EncodeToString(publicKeyBytes)), 0600)
	if err != nil {
		log.Fatal(err)
	}
}

// 从文件读取公钥
func LoadPublicKey(filename string) *ecdsa.PublicKey {
	publicKeyHex, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	publicKeyBytes, err := hex.DecodeString(string(publicKeyHex))
	if err != nil {
		log.Fatal(err)
	}
	publicKey, err := crypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	return publicKey
}
