package main

import (
	"encoding/base64"
	"flag"
	"fmt"

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

/*
请注意这只是个demo，理论上不要在你的命令行终端里输入私钥，以免泄露
因此只建议使用测试网的私钥运行

go run main.go -from=TBYHGsFkshasvB3R6Zys4627h98owvUNFn -to=TEucCiybmbLhG8it1RM31js91fMLCjEARY -amount=5000000 -pk=PRIVATE-KEY-HEX
*/
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

	fmt.Println("grpc_address:", grpcAddress)
	fmt.Println("from_address:", fromAddress)
	fmt.Println("to_address:", toAddress)
	fmt.Println("amount:", amount)
	fmt.Println("private_key_hex_length:", len(privateKeyHex))

	must.Nice(grpcAddress)
	must.Nice(fromAddress)
	must.Nice(toAddress)
	must.Nice(amount)
	must.Nice(privateKeyHex)

	client := gotrongrpc.NewClient(rese.P1(gotrongrpc.NewGrpcClient(grpcAddress)))

	rawTransaction := rese.P1(client.GetGrpc().Transfer(
		fromAddress, // 发送者，也就是您，的钱包地址
		toAddress,   // 接收者的钱包地址
		amount,      // 要发送的数量，注意因为这里是整数，需要自行转换单位
	))
	fmt.Println(neatjsons.S(rawTransaction))

	signature := signTransaction(privateKeyHex, rawTransaction.Transaction.RawData)

	sendTransaction(client, rawTransaction.Transaction.RawData, signature)
}

func signTransaction(privateKeyHex string, txRaw *core.TransactionRaw) []byte {
	fmt.Println("tx_data:", neatjsons.SxB(rese.V1(protojson.Marshal(txRaw))))
	fmt.Println("tx_hash:", rese.C1(gotronhash.GetTxHash(txRaw)))

	signature := rese.V1(gotronsign.Sign(privateKeyHex, txRaw))
	fmt.Println(len(signature))
	fmt.Println(base64.StdEncoding.EncodeToString(signature))
	return signature
}

func sendTransaction(client *gotrongrpc.Client, txRaw *core.TransactionRaw, signature []byte) {
	signedTransaction := &core.Transaction{
		RawData:   txRaw,
		Signature: [][]byte{signature},
	}

	resp := rese.P1(client.GetGrpc().Broadcast(signedTransaction))
	fmt.Println(neatjsons.S(resp))

	must.True(resp.GetResult())
	must.Equals(resp.Code, api.Return_SUCCESS)
	must.None(string(resp.Message))

	fmt.Println("success")
}
