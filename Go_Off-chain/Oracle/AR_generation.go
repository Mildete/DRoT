package ARtrust

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"os"
	"time"
)

type AttestationReport struct {
	NodeID          string         `json:"node_id"`
	RandomValue     string         `json:"random_value"`
	Proof           [4]*big.Int    `json:"proof"`
	AttestationTime *big.Int       `json:"attestation_time"`
	Measurement     string         `json:"measurement"`
	R               string         `json:"r"`
	S               string         `json:"s"`
	V               uint8          `json:"v"`
	ReportMessage   string         `json:"data_tosign"`
	ReportHash      string         `json:"report_hash"`
	Signature       []byte         `json:"signature"`
	PublickeyHex    common.Address `json:"public_key_hex"`
}

// SignData signs the provided data using the given private key.
func SignData(message_to_sign string, privateKey *ecdsa.PrivateKey) (string, []byte, uint8, string, string, common.Address) {
	// 1. 计算消息哈希
	hash := sha256.Sum256([]byte(message_to_sign))
	fmt.Printf("Message_to_sign Hash: 0x%x\n", hash)
	// 2. 使用私钥对消息进行签名
	signature, err := crypto.Sign(hash[:], privateKey)
	if err != nil {
		log.Fatal(err)
	}
	//3.返回签名
	fmt.Printf("Signature: ", signature)

	// 4. 从签名中提取 r, s, v
	r := new(big.Int).SetBytes(signature[:32])
	s := new(big.Int).SetBytes(signature[32:64])
	v := signature[64] + 27 // Ethereum的v值通常为27或28
	//fmt.Printf("Signature: \nr = %s\ns = %s\nv = %d\n", r, s, v)

	// 5. 确保 r 和 s 是32字节的定长字节数组
	// 确保 r 和 s 是32字节的定长数组（bytes32）
	rBytes := [32]byte{}
	sBytes := [32]byte{}
	copy(rBytes[32-len(r.Bytes()):], r.Bytes())
	copy(sBytes[32-len(s.Bytes()):], s.Bytes())
	//r_forsol := "0x" + rBytesPadded

	//fmt.Printf("Signature: \n Padded r (for Solidity): 0x%x\n", rBytes)
	//fmt.Printf("Padded s (for Solidity): 0x%x\n", sBytes)
	//fmt.Printf("v = %d\n", v)

	// 6. 输出公钥地址 (发送给Solidity)
	publicKey := privateKey.PublicKey
	signerAddress := crypto.PubkeyToAddress(publicKey)
	fmt.Printf("Public Key Address (Signer): %s\n", signerAddress.Hex())
	rstr := "0x" + hex.EncodeToString(rBytes[:])
	sstr := "0x" + hex.EncodeToString(sBytes[:])
	hashstr := "0x" + hex.EncodeToString(hash[:])
	return hashstr, signature, v, rstr, sstr, signerAddress
}

// GenerateAttestationReport generates an attestation report for the oracle node using the provided key pair.
func GenerateAttestationReport(nodeID string, privateKey *ecdsa.PrivateKey, measurement, filePath string, test_interval int) (*AttestationReport, error) {

	// Record the current time as the attestation time
	attestationTime64 := time.Now().Unix()
	attestationTime := big.NewInt(attestationTime64)
	attestationTime.Sub(attestationTime, big.NewInt(int64(10*test_interval))) //增加休眠时间，通过时间戳控制信任分使用

	//GetMeasure()  TPM执行

	// Generate a random value and proof
	vrfpk, proof := VRF_Gen([]byte(measurement), privateKey)
	fmt.Printf("Public Key: %s\n", vrfpk)
	fmt.Printf("Proof: %s\n", proof)

	randomValue := fmt.Sprintf("%x", sha256.Sum256([]byte(measurement+time.Now().String())))
	//proof := fmt.Sprintf("%x", sha256.Sum256([]byte(randomValue)))

	// Create the data to be signed, which includes node ID, random value, proof, time, and measurement
	// dataToSign := []byte(nodeID + randomValue + proof + fmt.Sprintf("%d", attestationTime) + measurement)
	dataToSign := nodeID + randomValue + fmt.Sprintf("%d", proof) + fmt.Sprintf("%d", attestationTime) + measurement
	//fmt.Println(dataToSign)
	// Sign the data using the private key
	reporthash, signature, v, r, s, pkhex := SignData(dataToSign, privateKey)

	// Create the attestation report
	report := &AttestationReport{
		NodeID:          nodeID,
		RandomValue:     randomValue,
		Proof:           proof,
		AttestationTime: attestationTime,
		Measurement:     measurement,
		ReportHash:      reporthash,
		Signature:       signature,
		ReportMessage:   dataToSign,
		V:               v,
		R:               r,
		S:               s,
		PublickeyHex:    pkhex,
	}

	// Save the attestation report to a file
	if err := SaveAttestationReport(report, filePath); err != nil {
		return nil, err
	}

	return report, nil
}

// SaveAttestationReport saves the attestation report to a file.
func SaveAttestationReport(report *AttestationReport, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	data, err := json.Marshal(report)
	if err != nil {
		return fmt.Errorf("failed to marshal attestation report: %v", err)
	}

	if _, err := file.WriteString(string(data) + "\n"); err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	return nil
}

//// ECDSA-P256 Version-Precompiled Verify
//func signMessage(privateKey *ecdsa.PrivateKey, message string) (string, string, string, string, error) {
//	// 对消息进行哈希
//	hash := sha256.Sum256([]byte(message))
//
//	// 使用私钥对哈希进行签名
//	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
//	if err != nil {
//		return "", "", "", "", err
//	}
//
//	// 获取v值 (恢复ID，27或28)
//	v := 0
//	if s.Bit(0) == 0 {
//		v = 27
//	} else {
//		v = 28
//	}
//
//	// 获取公钥
//	pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
//
//	// 将r, s, v, 公钥转换为hex字符串
//	rHex := hex.EncodeToString(r.Bytes())
//	sHex := hex.EncodeToString(s.Bytes())
//	vHex := fmt.Sprintf("%d", v)
//	pubKeyHex := hex.EncodeToString(pubKey)
//
//	return rHex, sHex, vHex, pubKeyHex, nil
//}

func GetMeasurement(NodeID string, pkhex string, attestationfun string) string {
	str := NodeID + pkhex + attestationfun
	measurehash := sha256.Sum256([]byte(str))
	measurement := hex.EncodeToString(measurehash[:])
	return measurement
}

func VRF_Gen(message []byte, privateKey *ecdsa.PrivateKey) ([2]*big.Int, [4]*big.Int) {
	publicKey := privateKey.PublicKey

	// 2. 消息
	hash := sha256.Sum256(message)

	// 3. 生成伪随机点 (Gamma)
	gammaX, gammaY := publicKey.Curve.ScalarBaseMult(hash[:])

	// 4. 生成挑战值 (c) 和响应值 (s)
	c := new(big.Int).SetBytes(hash[:16]) // 模拟挑战值 (取哈希前 16 字节)
	s := new(big.Int).Set(privateKey.D)   // 响应值 (私钥的倍数)

	vrfpk := [2]*big.Int{publicKey.X, publicKey.Y}
	proof := [4]*big.Int{gammaX, gammaY, c, s}
	return vrfpk, proof
}
