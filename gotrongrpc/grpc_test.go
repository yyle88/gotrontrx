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
