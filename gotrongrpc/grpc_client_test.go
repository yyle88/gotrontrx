// Package gotrongrpc_test: Comprehensive test suite for TRON gRPC client operations
// Tests account balance queries, TRC20 token balances, and transaction creation
// Validates integration with TRON testnet (Shasta) blockchain operations
//
// gotrongrpc_test: TRON gRPC 客户端操作的综合测试套件
// 测试账户余额查询、TRC20 代币余额和交易创建
// 验证与 TRON 测试网（Shasta）区块链操作的集成
package gotrongrpc_test

import (
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/gotrontrx/gotrongrpc"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
)

var grpcClient *client.GrpcClient

func TestMain(m *testing.M) {
	grpcClient = rese.P1(gotrongrpc.NewGrpcClient(gotrongrpc.ShastaNetGrpc))
	defer grpcClient.Stop()
	m.Run()
}

func TestGrpcClient_Balance(t *testing.T) {
	account, err := grpcClient.GetAccountDetailed("TBYHGsFkshasvB3R6Zys4627h98owvUNFn")
	require.NoError(t, err)
	t.Log(neatjsons.S(account))
}

func TestGrpcClient_TRC20ContractBalance(t *testing.T) {
	value, err := grpcClient.TRC20ContractBalance(
		"THidAkJ7mmWPuEuM9BsoNYsXLpWU6SBJ7h",
		"TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs",
	)
	require.NoError(t, err)
	t.Log(neatjsons.S(value))
}

func TestGrpcClient_Transfer_Transaction(t *testing.T) {
	rawTransaction, err := grpcClient.Transfer(
		"TBYHGsFkshasvB3R6Zys4627h98owvUNFn",
		"TEucCiybmbLhG8it1RM31js91fMLCjEARY",
		5000000,
	)
	require.NoError(t, err)
	t.Log(neatjsons.S(rawTransaction))
}
