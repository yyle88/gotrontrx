package gotronhash_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gotrontrx/gotrongrpc"
	"github.com/yyle88/gotrontrx/gotronhash"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
)

var caseClientShasta *gotrongrpc.Client

func TestMain(m *testing.M) {
	caseClientShasta = gotrongrpc.NewClient(rese.P1(gotrongrpc.NewGrpcClient(gotrongrpc.ShastaNetGrpc)))
	m.Run()
}

func TestGetTxHash(t *testing.T) {
	rawTransaction := rese.P1(caseClientShasta.GetGrpc().Transfer(
		"TBYHGsFkshasvB3R6Zys4627h98owvUNFn", // 发送者，也就是您，的钱包地址
		"TEucCiybmbLhG8it1RM31js91fMLCjEARY", // 接收者的钱包地址
		5000000,                              // 要发送的数量，注意因为这里是整数，需要自行转换单位
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
	rawTransaction := rese.P1(caseClientShasta.GetGrpc().Transfer(
		"TBYHGsFkshasvB3R6Zys4627h98owvUNFn", // 发送者，也就是您，的钱包地址
		"TEucCiybmbLhG8it1RM31js91fMLCjEARY", // 接收者的钱包地址
		5000000,                              // 要发送的数量，注意因为这里是整数，需要自行转换单位
	))
	fmt.Println(neatjsons.S(rawTransaction))
	require.Len(t, rawTransaction.Txid, 32)

	newTxHash, err := gotronhash.RawTxHash(rawTransaction.Transaction.RawData)
	require.NoError(t, err)
	t.Log(len(newTxHash))

	require.Equal(t, rawTransaction.Txid, newTxHash)
}
