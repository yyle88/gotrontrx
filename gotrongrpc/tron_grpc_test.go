package gotrongrpc

import (
	"encoding/base64"
	"math/big"
	"testing"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/gotron/gotronhash"
	"github.com/yyle88/gotron/gotronsign"
	"github.com/yyle88/gotron/internal/utils"
	"google.golang.org/protobuf/encoding/protojson"
)

var caseClient *Client
var caseClientShasta *Client

func TestMain(m *testing.M) {
	caseClient = mustNewClient(MainnetGrpc)
	caseClientShasta = mustNewClient(ShastaNetGrpc)
	m.Run()
}

func TestGrpcClient_Transfer(t *testing.T) {
	txn, err := caseClient.grpcClient.Transfer(
		"TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
		"TRJvQFUWwSmnk5rgM8m4HgE6Csj2qPEupX",
		1000,
	)
	require.NoError(t, err)
	t.Log(utils.SoftNeatString(txn))
}

func TestGrpcClient_Balance(t *testing.T) {
	txn, err := caseClientShasta.grpcClient.GetAccountDetailed("TBYHGsFkshasvB3R6Zys4627h98owvUNFn")
	require.NoError(t, err)
	t.Log(utils.SoftNeatString(txn))
}

func TestGrpcClient_TRC20Send(t *testing.T) {
	txn, err := caseClient.grpcClient.TRC20Send(
		"TWd4WrZ9wn84f5x1hZhL4DHvk738ns5jwb",
		"TEpjT8xbAe3FPCPFziqFfEjLVXaw9NbGXj",
		"TAFjULxiVgT4qWk6UZwjqwZXTSaGaqnVp4",
		big.NewInt(100),
		200,
	)
	require.NoError(t, err)
	t.Log(utils.SoftNeatString(txn))
}

func TestGrpcClient_TRC20ContractBalance(t *testing.T) {
	value, err := caseClientShasta.grpcClient.TRC20ContractBalance(
		"THidAkJ7mmWPuEuM9BsoNYsXLpWU6SBJ7h",
		"TG3XXyExBkPp9nzdajDZsozEu4BkaSJozs",
	)
	require.NoError(t, err)
	t.Log(utils.SoftNeatString(value))
}

func TestGrpcClient_TRC20Allowance(t *testing.T) {
	value, err := caseClient.TRC20Allowance(
		"TWd4WrZ9wn84f5x1hZhL4DHvk738ns5jwb",
		"TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
		"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
	)
	require.NoError(t, err)
	t.Log(value.String())
}

func TestGrpcClient_TRC20Allowance_2(t *testing.T) {
	value, err := caseClient.TRC20Allowance(
		"TDToUxX8sH4z6moQpK3ZLAN24eupu2ivA4",
		"TVh1PF9xr4zC5uAqRcCbxF1By6ucp95G4i",
		"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
	)
	require.NoError(t, err)
	t.Log(value.String())
}

func TestGrpcClient_TRC20Allowance_3(t *testing.T) {
	value, err := caseClient.TRC20Allowance(
		"TMPZ9SCZAgPJZh6zKjj2HJeWCD4xxvohGv",
		"TSUfx6xDEUGNzWQVMbqjf1zD5xwX4DQ9PL",
		"TNojYuHohVLPxb4MxjuUKnQJX36h3uB9Wg",
	)
	require.NoError(t, err)
	t.Log(value.String())
}

func TestGrpcClient_TRC20Allowance_4(t *testing.T) {
	value, err := caseClient.TRC20Allowance(
		"TWd4WrZ9wn84f5x1hZhL4DHvk738ns5jwb",
		"TEpjT8xbAe3FPCPFziqFfEjLVXaw9NbGXj",
		"TAFjULxiVgT4qWk6UZwjqwZXTSaGaqnVp4",
	)
	require.NoError(t, err)
	t.Log(value.String())
}

func TestGrpcClient_Transfer_LuoLuoToFengGe(t *testing.T) {
	txn, err := caseClientShasta.grpcClient.Transfer(
		"TBYHGsFkshasvB3R6Zys4627h98owvUNFn",
		"TEucCiybmbLhG8it1RM31js91fMLCjEARY",
		5000000,
	)
	require.NoError(t, err)
	t.Log(utils.SoftNeatString(txn))

	caseSignTronTransactionSendBroadcast(t, txn)
}

func caseSignTronTransactionSendBroadcast(t *testing.T, txn *api.TransactionExtention) {
	rawTx := txn.Transaction.RawData

	{
		txData, err := protojson.Marshal(rawTx)
		require.NoError(t, err)
		t.Log(utils.NeatStringXBytes(txData))
	}
	{
		txHash, err := gotronhash.GetTxHashHex(rawTx)
		require.NoError(t, err)
		t.Log(txHash)
	}

	privateKeyHex := "-YOUR PRIVATE KEY-"

	signature, err := gotronsign.Sign(privateKeyHex, rawTx)
	require.NoError(t, err)
	t.Log(len(signature))
	t.Log(base64.StdEncoding.EncodeToString(signature))

	paramX := &core.Transaction{
		RawData:   rawTx,
		Signature: [][]byte{signature},
	}

	res, err := caseClientShasta.grpcClient.Broadcast(paramX)
	require.NoError(t, err)
	t.Log(utils.SoftNeatString(res))
}
