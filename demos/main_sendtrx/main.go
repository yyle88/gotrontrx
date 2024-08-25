package main

import (
	"encoding/base64"
	"flag"
	"fmt"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/api"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/pkg/errors"
	"github.com/yyle88/done"
	"github.com/yyle88/gotron/gotrongrpc"
	"github.com/yyle88/gotron/gotronhash"
	"github.com/yyle88/gotron/gotronsign"
	"github.com/yyle88/neatjson"
	"github.com/yyle88/neatjson/neatjsons"
	"google.golang.org/protobuf/encoding/protojson"
)

/*
请注意这只是个demo，理论上不要在你的命令行终端里输入私钥，以免泄露
因此只建议使用测试网的私钥运行

go run main.go -from=TBYHGsFkshasvB3R6Zys4627h98owvUNFn -to=TEucCiybmbLhG8it1RM31js91fMLCjEARY -amount=5000000 -pk=PRIVATE-KEY-HEX
*/
func main() {
	var grpcAddress string
	var fromAddress, toAddress string
	var amount int64
	var privateKeyHex string

	flag.StringVar(&grpcAddress, "grpc", gotrongrpc.ShastaNetGrpc, "波场节点的grpc地址")
	flag.StringVar(&fromAddress, "from", "", "发送方地址 example: TBYHGsFkshasvB3R6Zys4627h98owvUNFn")
	flag.StringVar(&toAddress, "to", "", "接收方地址 example: TEucCiybmbLhG8it1RM31js91fMLCjEARY")
	flag.Int64Var(&amount, "amount", 0, "金额 example: 5000000")
	flag.StringVar(&privateKeyHex, "pk", "", "私钥 example: cdad54217914446687162216a4de99c28b9915c5683c4cf58aa525606d5f9c2f")
	flag.Parse()

	fmt.Println("grpc_address:", grpcAddress)
	fmt.Println("from_address:", fromAddress)
	fmt.Println("to_address:", toAddress)
	fmt.Println("amount:", amount)
	fmt.Println("private_key_hex_size:", len(privateKeyHex))
	if grpcAddress == "" || fromAddress == "" || toAddress == "" || amount == 0 || privateKeyHex == "" {
		panic(errors.New("param is missing"))
	}

	client, err := gotrongrpc.NewClientWithAddress(grpcAddress)
	done.Done(err)

	txn, err := client.GetGrpcClient().Transfer(
		fromAddress,
		toAddress,
		amount,
	)
	done.Done(err)
	fmt.Println(neatjsons.S(txn))

	signature := signTx(privateKeyHex, txn.Transaction.RawData)

	sendTx(client, txn.Transaction.RawData, signature)
}

func signTx(privateKeyHex string, rawTx *core.TransactionRaw) []byte {
	{
		txData, err := protojson.Marshal(rawTx)
		done.Done(err)
		txNeat, err := neatjson.TAB.SxB(txData)
		done.Done(err)
		fmt.Println("tx_neat:", txNeat)
	}
	{
		txHash, err := gotronhash.GetTxHashHex(rawTx)
		done.Done(err)
		fmt.Println("tx_hash:", txHash)
	}

	signature, err := gotronsign.Sign(privateKeyHex, rawTx)
	done.Done(err)

	{
		fmt.Println(len(signature))
		fmt.Println(base64.StdEncoding.EncodeToString(signature))
	}

	return signature
}

func sendTx(client *gotrongrpc.Client, rawTx *core.TransactionRaw, signature []byte) {
	paramX := &core.Transaction{
		RawData:   rawTx,
		Signature: [][]byte{signature},
	}

	res, err := client.GetGrpcClient().Broadcast(paramX)
	done.Done(err)
	fmt.Println(neatjsons.S(res))

	if !res.GetResult() {
		panic(errors.New("result is false"))
	}
	if res.Code != api.Return_SUCCESS {
		panic(errors.Errorf("res.code=%v != %v", res.Code, api.Return_SUCCESS))
	}
	if len(res.Message) != 0 {
		panic(errors.Errorf("message=%s is not none", string(res.Message)))
	}
	fmt.Println("success")
}
