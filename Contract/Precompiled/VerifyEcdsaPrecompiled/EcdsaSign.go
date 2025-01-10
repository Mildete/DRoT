package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 生成ECDSA签名
func signMessage(privateKey *ecdsa.PrivateKey, message string) (string, string, string, string, error) {
	// 对消息进行哈希
	hash := sha256.Sum256([]byte(message))

	// 使用私钥对哈希进行签名
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", "", "", "", err
	}

	// 获取v值 (恢复ID，27或28)
	v := 0
	if s.Bit(0) == 0 {
		v = 27
	} else {
		v = 28
	}

	// 获取公钥
	pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	// 将r, s, v, 公钥转换为hex字符串
	rHex := hex.EncodeToString(r.Bytes())
	sHex := hex.EncodeToString(s.Bytes())
	vHex := fmt.Sprintf("%d", v)
	pubKeyHex := hex.EncodeToString(pubKey)

	return rHex, sHex, vHex, pubKeyHex, nil
}

func main() {
	// 生成ECDSA私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	message := "Hello, FISCO BCOS!"
	r, s, v, pubKey, err := signMessage(privateKey, message)
	if err != nil {
		fmt.Println("Error signing message:", err)
		return
	}

	fmt.Println("r:", r)
	fmt.Println("s:", s)
	fmt.Println("v:", v)
	fmt.Println("publicKey:", pubKey)
}
