package examples

import (
	"math/big"
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/gotrontrx/gotrongrpc"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
)

var grpcClient *client.GrpcClient

func TestMain(m *testing.M) {
	grpcClient = rese.P1(gotrongrpc.NewGrpcClient(gotrongrpc.MainNetGrpc))
	m.Run()
}

func TestGrpcClient_Transfer_Transaction(t *testing.T) {
	rawTransaction, err := grpcClient.Transfer(
		"TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
		"TRJvQFUWwSmnk5rgM8m4HgE6Csj2qPEupX",
		1000,
	)
	require.NoError(t, err)
	t.Log(neatjsons.S(rawTransaction))
}

func TestGrpcClient_TRC20Send(t *testing.T) {
	rawTransaction, err := grpcClient.TRC20Send(
		"TWd4WrZ9wn84f5x1hZhL4DHvk738ns5jwb",
		"TEpjT8xbAe3FPCPFziqFfEjLVXaw9NbGXj",
		"TAFjULxiVgT4qWk6UZwjqwZXTSaGaqnVp4",
		big.NewInt(100),
		200,
	)
	require.NoError(t, err)
	t.Log(neatjsons.S(rawTransaction))
}

func TestGrpcClient_TRC20Allowance(t *testing.T) {
	value, err := gotrongrpc.TRC20Allowance(
		grpcClient,
		"TWd4WrZ9wn84f5x1hZhL4DHvk738ns5jwb",
		"TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
		"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
	)
	require.NoError(t, err)
	t.Log(value.String())
}

func TestGrpcClient_TRC20Allowance_2(t *testing.T) {
	value, err := gotrongrpc.TRC20Allowance(
		grpcClient,
		"TDToUxX8sH4z6moQpK3ZLAN24eupu2ivA4",
		"TVh1PF9xr4zC5uAqRcCbxF1By6ucp95G4i",
		"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
	)
	require.NoError(t, err)
	t.Log(value.String())
}

func TestGrpcClient_TRC20Allowance_3(t *testing.T) {
	value, err := gotrongrpc.TRC20Allowance(
		grpcClient,
		"TMPZ9SCZAgPJZh6zKjj2HJeWCD4xxvohGv",
		"TSUfx6xDEUGNzWQVMbqjf1zD5xwX4DQ9PL",
		"TNojYuHohVLPxb4MxjuUKnQJX36h3uB9Wg",
	)
	require.NoError(t, err)
	t.Log(value.String())
}

func TestGrpcClient_TRC20Allowance_4(t *testing.T) {
	value, err := gotrongrpc.TRC20Allowance(
		grpcClient,
		"TWd4WrZ9wn84f5x1hZhL4DHvk738ns5jwb",
		"TEpjT8xbAe3FPCPFziqFfEjLVXaw9NbGXj",
		"TAFjULxiVgT4qWk6UZwjqwZXTSaGaqnVp4",
	)
	require.NoError(t, err)
	t.Log(value.String())
}
