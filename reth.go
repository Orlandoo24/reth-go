package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 定义常量
const (
	account       = "地址"        // Ethereum账户地址
	privateKeyHex = "私钥"        // 私钥的十六进制表示
	maxMintTimes  = 40          // 最大mint次数
	gasTop        = 30014126189 // 最大gas价格
)

// 寻找解决方案
func findSolution(currentChallenge string) (string, error) {
	startTime := time.Now()
	count := 0

	for {
		randomValue := make([]byte, 32)
		_, err := rand.Read(randomValue)
		if err != nil {
			return "", err
		}
		potentialSolution := hex.EncodeToString(randomValue)

		hashedSolution := crypto.Keccak256Hash([]byte(potentialSolution + currentChallenge)).Hex()
		count++

		fmt.Println(count)

		if hashedSolution[:6] == "0x7777" {
			endTime := time.Now()
			fmt.Printf("use time: %f s\n", endTime.Sub(startTime).Seconds())
			return potentialSolution, nil
		}
	}
}

func main() {
	// 连接到以太坊节点
	client, err := ethclient.Dial("https://rpc.flashbots.net")
	if err != nil {
		fmt.Println("Error connecting to provider:", err)
		return
	}

	// 解析私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		fmt.Println("Error parsing private key:", err)
		return
	}

	// 获取公钥和地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("Error casting public key to ECDSA")
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取当前挑战的哈希值
	currentChallenge := hexutil.Encode(crypto.Keccak256([]byte("rETH")))

	for i := 1; i <= maxMintTimes; i++ {
		// 寻找解决方案
		solution, err := findSolution(currentChallenge)
		if err != nil {
			fmt.Printf("Error finding solution: %s\n", err)
			continue
		}

		// 构建交易数据
		jsonData := map[string]interface{}{
			"p":    "rerc-20",
			"op":   "mint",
			"tick": "rETH",
			"id":   solution,
			"amt":  "10000",
		}
		dataHex := hexutil.Encode([]byte(fmt.Sprintf("data:application/json,%s", jsonData)))

		// 获取gas价格
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			fmt.Println("Error getting gas price:", err)
			continue
		}

		// gas价格过高时等待
		for gasPrice.Cmp(big.NewInt(gasTop)) >= 0 {
			time.Sleep(time.Second * 10)
			gasPrice, err = client.SuggestGasPrice(context.Background())
			if err != nil {
				fmt.Println("Error getting gas price:", err)
				continue
			}
		}

		// 构建交易并签名
		auth := bind.NewKeyedTransactor(privateKey)
		auth.Nonce = big.NewInt(int64(i)) // 使用nonce作为交易序号
		auth.Value = big.NewInt(0)        // 转账金额，这里设置为0
		auth.GasLimit = uint64(3000000)   // gas限制
		auth.GasPrice = gasPrice          // gas价格

		tx := types.NewTransaction(auth.Nonce.Uint64(), fromAddress, auth.Value, auth.GasLimit, auth.GasPrice, []byte(dataHex))
		signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
		if err != nil {
			fmt.Println("Error signing transaction:", err)
			continue
		}

		// 发送交易
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			fmt.Printf("Error sending transaction: %s\n", err)
			continue
		}

		fmt.Printf("第%d次mint成功，transaction hash为：%s\n", i, signedTx.Hash().Hex())
	}
}
