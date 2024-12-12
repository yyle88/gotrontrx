package gotrongrpc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gotrontrx/gotrongrpc"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
)

var caseClient *gotrongrpc.Client

func TestMain(m *testing.M) {
	caseClient = gotrongrpc.NewClient(rese.P1(gotrongrpc.NewGrpcClient(gotrongrpc.ShastaNetGrpc)))
	m.Run()
}

func TestGrpcClient_Balance(t *testing.T) {
	account, err := caseClient.GetGrpc().GetAccountDetailed("TBYHGsFkshasvB3R6Zys4627h98owvUNFn")
	require.NoError(t, err)
	t.Log(neatjsons.S(account))
}

func TestGrpcClient_TRC20ContractBalance(t *testing.T) {
	value, err := caseClient.GetGrpc().TRC20ContractBalance(
		"THidAkJ7mmWPuEuM9BsoNYsXLpWU6SBJ7h",
		"TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs",
	)
	require.NoError(t, err)
	t.Log(neatjsons.S(value))
}

func TestGrpcClient_Transfer_Transaction(t *testing.T) {
	rawTransaction, err := caseClient.GetGrpc().Transfer(
		"TBYHGsFkshasvB3R6Zys4627h98owvUNFn",
		"TEucCiybmbLhG8it1RM31js91fMLCjEARY",
		5000000,
	)
	require.NoError(t, err)
	t.Log(neatjsons.S(rawTransaction))
}
