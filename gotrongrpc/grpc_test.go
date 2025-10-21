// Package gotrongrpc_test: Test suite for TRON gRPC client functionality
// Verifies connection establishment and basic node information queries
// Tests against TRON testnet (Nile) to validate client operations
//
// gotrongrpc_test: TRON gRPC 客户端功能的测试套件
// 验证连接建立和基本节点信息查询
// 针对 TRON 测试网（Nile）测试以验证客户端操作
package gotrongrpc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gotrontrx/gotrongrpc"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestNewGrpcClient(t *testing.T) {
	client, err := gotrongrpc.NewGrpcClient(gotrongrpc.NileNetGrpc)
	require.NoError(t, err)

	nodeInfo, err := client.GetNodeInfo()
	require.NoError(t, err)
	t.Log(neatjsons.S(nodeInfo))
}
