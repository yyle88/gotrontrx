// Package gotronhash_test: Test suite for TRON transaction hash calculations
// Validates SHA256 hash generation against testnet transaction responses
// Ensures hash calculation consistency with TRON network standards
//
// gotronhash_test: TRON 交易哈希计算的测试套件
// 针对测试网交易响应验证 SHA256 哈希生成
// 确保哈希计算与 TRON 网络标准保持一致
package gotronhash_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/gotrontrx/gotrongrpc"
	"github.com/yyle88/gotrontrx/gotronhash"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
)

var grpcClientShasta *client.GrpcClient

func TestMain(m *testing.M) {
	grpcClientShasta = rese.P1(gotrongrpc.NewGrpcClient(gotrongrpc.ShastaNetGrpc))
	m.Run()
}

func TestGetTxHash(t *testing.T) {
	rawTransaction := rese.P1(grpcClientShasta.Transfer(
		"TBYHGsFkshasvB3R6Zys4627h98owvUNFn", // Sender wallet address // 发送者钱包地址
		"TEucCiybmbLhG8it1RM31js91fMLCjEARY", // Recipient wallet address // 接收者钱包地址
		5000000,                              // Amount in base units // 基础单位金额
	))
	fmt.Println(neatjsons.S(rawTransaction))

	rawTxHash := hex.EncodeToString(rawTransaction.Txid)
	t.Log(rawTxHash)

	newTxHash, err := gotronhash.GetTxHash(rawTransaction.Transaction.RawData)
	require.NoError(t, err)
	t.Log(newTxHash)

	require.Equal(t, rawTxHash, newTxHash)
}

func TestRawTxHash(t *testing.T) {
	rawTransaction := rese.P1(grpcClientShasta.Transfer(
		"TBYHGsFkshasvB3R6Zys4627h98owvUNFn", // Sender wallet address // 发送者钱包地址
		"TEucCiybmbLhG8it1RM31js91fMLCjEARY", // Recipient wallet address // 接收者钱包地址
		5000000,                              // Amount in base units // 基础单位金额
	))
	fmt.Println(neatjsons.S(rawTransaction))
	require.Len(t, rawTransaction.Txid, 32)

	newTxHash, err := gotronhash.RawTxHash(rawTransaction.Transaction.RawData)
	require.NoError(t, err)
	t.Log(len(newTxHash))

	require.Equal(t, rawTransaction.Txid, newTxHash)
}
