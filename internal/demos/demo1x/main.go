// Package main: TRON blockchain transaction demonstration
// Shows complete workflow of creating, signing, and broadcasting TRX transactions
// Uses command-line flags to configure transaction parameters
//
// 演示 TRON 区块链交易的完整流程
// 展示创建、签名和广播 TRX 交易的完整工作流
// 使用命令行标志配置交易参数
package main

import (
	"encoding/base64"
	"flag"
	"fmt"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/yyle88/gotrontrx/gotrongrpc"
	"github.com/yyle88/gotrontrx/gotronhash"
	"github.com/yyle88/gotrontrx/gotronsign"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"google.golang.org/protobuf/encoding/protojson"
)

// main demonstrates the complete TRON transaction workflow
// IMPORTANT: This is a demo - avoid entering private keys in terminal to prevent leaks
// RECOMMENDED: Use testnet private keys when running this demo
//
// Usage example:
// go run main.go -from=TBYHGsFkshasvB3R6Zys4627h98owvUNFn -to=TEucCiybmbLhG8it1RM31js91fMLCjEARY -amount=5000000 -pk=PRIVATE-KEY-HEX
//
// 演示完整的 TRON 交易工作流程
// 重要提示：这是演示程序 - 避免在终端输入私钥以防泄露
// 建议：运行此演示时使用测试网私钥
//
// 使用示例：
// go run main.go -from=TBYHGsFkshasvB3R6Zys4627h98owvUNFn -to=TEucCiybmbLhG8it1RM31js91fMLCjEARY -amount=5000000 -pk=PRIVATE-KEY-HEX
func main() {
	var grpcAddress string
	var fromAddress string
	var toAddress string
	var amount int64
	var privateKeyHex string

	flag.StringVar(&grpcAddress, "grpc", gotrongrpc.ShastaNetGrpc, "波场节点的grpc地址")
	flag.StringVar(&fromAddress, "from", "", "发送方地址 example: TBYHGsFkshasvB3R6Zys4627h98owvUNFn")
	flag.StringVar(&toAddress, "to", "", "接收方地址 example: TEucCiybmbLhG8it1RM31js91fMLCjEARY")
	flag.Int64Var(&amount, "amount", 0, "金额(当金额是整数时表示没带单位需要你自己转换) example: 5000000")
	flag.StringVar(&privateKeyHex, "pk", "", "私钥 example: cdad54217914446687162216a4de99c28b9915c5683c4cf58aa525606d5f9c2f")
	flag.Parse()

	// Display configuration parameters
	// 显示配置参数
	fmt.Println("grpc_address:", grpcAddress)
	fmt.Println("from_address:", fromAddress)
	fmt.Println("to_address:", toAddress)
	fmt.Println("amount:", amount)
	fmt.Println("private_key_hex_length:", len(privateKeyHex))

	// Validate required parameters are provided
	// 验证必需参数是否已提供
	must.Nice(grpcAddress)
	must.Nice(fromAddress)
	must.Nice(toAddress)
	must.Nice(amount)
	must.Nice(privateKeyHex)

	// Connect to TRON network
	// 连接到 TRON 网络
	grpcClient := rese.P1(gotrongrpc.NewGrpcClient(grpcAddress))

	// Create TRX transaction
	// 创建 TRX 交易
	rawTransaction := rese.P1(grpcClient.Transfer(
		fromAddress, // Sender wallet address // 发送者钱包地址
		toAddress,   // Recipient wallet address // 接收者钱包地址
		amount,      // Amount in base units (needs manual conversion) // 基础单位金额（需手动转换）
	))
	fmt.Println(neatjsons.S(rawTransaction))

	// Sign transaction with private key
	// 使用私钥签名交易
	signature := signTransaction(privateKeyHex, rawTransaction.Transaction.RawData)

	// Broadcast signed transaction to network
	// 广播已签名交易到网络
	sendTransaction(grpcClient, rawTransaction.Transaction.RawData, signature)
}

// signTransaction generates signature for transaction using private key
// Displays transaction data and hash before signing
// Returns signature bytes ready to attach to transaction
//
// signTransaction 使用私钥为交易生成签名
// 签名前显示交易数据和哈希值
// 返回可附加到交易的签名字节
func signTransaction(privateKeyHex string, txRaw *core.TransactionRaw) []byte {
	// Display transaction data in JSON format
	// 以 JSON 格式显示交易数据
	fmt.Println("tx_data:", neatjsons.SxB(rese.V1(protojson.Marshal(txRaw))))

	// Display transaction hash
	// 显示交易哈希
	fmt.Println("tx_hash:", rese.C1(gotronhash.GetTxHash(txRaw)))

	// Generate ECDSA signature
	// 生成 ECDSA 签名
	signature := rese.V1(gotronsign.Sign(privateKeyHex, txRaw))
	fmt.Println(len(signature))
	fmt.Println(base64.StdEncoding.EncodeToString(signature))
	return signature
}

// sendTransaction broadcasts signed transaction to TRON network
// Validates response to ensure transaction was accepted
// Panics if transaction broadcast fails
//
// sendTransaction 将已签名交易广播到 TRON 网络
// 验证响应以确保交易被接受
// 如果交易广播失败则 panic
func sendTransaction(grpcClient *client.GrpcClient, txRaw *core.TransactionRaw, signature []byte) {
	// Construct signed transaction
	// 构建已签名交易
	signedTransaction := &core.Transaction{
		RawData:   txRaw,
		Signature: [][]byte{signature},
	}

	// Broadcast transaction to network
	// 广播交易到网络
	resp := rese.P1(grpcClient.Broadcast(signedTransaction))
	fmt.Println(neatjsons.S(resp))

	// Validate broadcast response
	// 验证广播响应
	must.True(resp.GetResult())
	must.Equals(resp.Code, api.Return_SUCCESS)
	must.None(string(resp.Message))

	fmt.Println("success")
}
