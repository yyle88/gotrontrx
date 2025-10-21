[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gotrontrx/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gotrontrx/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gotrontrx)](https://pkg.go.dev/github.com/yyle88/gotrontrx)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gotrontrx/main.svg)](https://coveralls.io/github/yyle88/gotrontrx?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://github.com/yyle88/gotrontrx)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gotrontrx.svg)](https://github.com/yyle88/gotrontrx/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gotrontrx)](https://goreportcard.com/report/github.com/yyle88/gotrontrx)

<p align="center">
  <img
    alt="wojack-cartoon logo"
    src="assets/wojack-cartoon.jpeg"
    style="max-height: 500px; width: auto; max-width: 100%;"
  />
</p>
<h3 align="center">golang-tron</h3>
<p align="center">使用 golang 创建/签名 <code>tron transaction</code></p>

---

# gotrontrx

`gotrontrx` 是一个 Go 工具包，用于探索 TRON 区块链技术，而不参与加密货币活动。

`gotrontrx` 包主要通过 gRPC 客户端与波场网络交互，开发者可以轻松完成交易创建、签名和广播的完整流程。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

---

## 功能概述

- **gRPC 客户端支持**：通过 gRPC 与波场节点建立连接，支持主网和测试网的节点访问。
- **账户与转账操作**：实现了转账功能，包括发送方、接收方和金额的指定。
- **交易签名**：支持基于私钥的交易签名操作，确保交易的安全性。
- **交易广播**：支持将签名后的交易广播到区块链网络中。
- **交易哈希计算**：提供方便的工具计算交易的哈希值。
- **响应处理**：内置了对波场 gRPC API 响应的结构化处理。

## 依赖模块

- `github.com/fbsobreira/gotron-sdk`：基础的波场 gRPC API 支持。
- `github.com/yyle88/gotrontrx`：封装的波场操作工具包，便于开发者高效使用。
- `neatjson` 信息的格式化输出
- `must` 基本的条件断言
- `rese` 能减少错误处理样板代码。

## 安装

```bash
go get github.com/yyle88/gotrontrx
```

## 快速开始

以下是一些核心功能的作用：

- **`gotrongrpc.NewGrpcClient`**: 初始化 gRPC 客户端，用于连接波场节点。
- **`grpcClient.Transfer`**: 构建转账交易。
- **`gotronsign.Sign`**: 使用私钥对交易进行签名。
- **`grpcClient.Broadcast`**: 广播签名后的交易到网络中。

通过这些功能，开发者可以快速实现基于波场区块链的应用开发。

## 使用方法

### 基本 TRX 转账交易

此示例展示创建、签名和广播 TRX 转账交易的完整工作流程。

```go
package main

import (
	"encoding/base64"
	"flag"
	"fmt"

	"github.com/fbsobreira/gotron-sdk/pkg/client"
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

func main() {
	var grpcAddress string
	var fromAddress string
	var toAddress string
	var amount int64
	var privateKeyHex string

	flag.StringVar(&grpcAddress, "grpc", gotrongrpc.ShastaNetGrpc, "波场节点的grpc地址")
	flag.StringVar(&fromAddress, "from", "", "发送方地址")
	flag.StringVar(&toAddress, "to", "", "接收方地址")
	flag.Int64Var(&amount, "amount", 0, "金额(基础单位)")
	flag.StringVar(&privateKeyHex, "pk", "", "私钥十六进制")
	flag.Parse()

	// 显示配置参数
	fmt.Println("grpc_address:", grpcAddress)
	fmt.Println("from_address:", fromAddress)
	fmt.Println("to_address:", toAddress)
	fmt.Println("amount:", amount)
	fmt.Println("private_key_hex_length:", len(privateKeyHex))

	// 验证必需参数是否已提供
	must.Nice(grpcAddress)
	must.Nice(fromAddress)
	must.Nice(toAddress)
	must.Nice(amount)
	must.Nice(privateKeyHex)

	// 连接到 TRON 网络
	grpcClient := rese.P1(gotrongrpc.NewGrpcClient(grpcAddress))

	// 创建 TRX 交易
	rawTransaction := rese.P1(grpcClient.Transfer(
		fromAddress, // 发送者钱包地址
		toAddress,   // 接收者钱包地址
		amount,      // 基础单位金额
	))
	fmt.Println(neatjsons.S(rawTransaction))

	// 使用私钥签名交易
	signature := signTransaction(privateKeyHex, rawTransaction.Transaction.RawData)

	// 广播已签名交易到网络
	sendTransaction(grpcClient, rawTransaction.Transaction.RawData, signature)
}

func signTransaction(privateKeyHex string, txRaw *core.TransactionRaw) []byte {
	// 以 JSON 格式显示交易数据
	fmt.Println("tx_data:", neatjsons.SxB(rese.V1(protojson.Marshal(txRaw))))

	// 显示交易哈希
	fmt.Println("tx_hash:", rese.C1(gotronhash.GetTxHash(txRaw)))

	// 生成 ECDSA 签名
	signature := rese.V1(gotronsign.Sign(privateKeyHex, txRaw))
	fmt.Println(len(signature))
	fmt.Println(base64.StdEncoding.EncodeToString(signature))
	return signature
}

func sendTransaction(grpcClient *client.GrpcClient, txRaw *core.TransactionRaw, signature []byte) {
	// 构建已签名交易
	signedTransaction := &core.Transaction{
		RawData:   txRaw,
		Signature: [][]byte{signature},
	}

	// 广播交易到网络
	resp := rese.P1(grpcClient.Broadcast(signedTransaction))
	fmt.Println(neatjsons.S(resp))

	// 验证广播响应
	must.True(resp.GetResult())
	must.Equals(resp.Code, api.Return_SUCCESS)
	must.None(string(resp.Message))

	fmt.Println("success")
}
```

⬆️ **源码：** [源码](internal/demos/demo1x/main.go)

## 注意事项

1. **注意安全**：不要在生产环境中直接在终端中输入私钥，以免造成泄露。
2. **测试环境**：建议在测试网中进行调试，避免误操作带来的经济损失。
3. **数据验证**：使用时请确保输入的地址和金额均合法且符合波场区块链的要求。

## 波场教程

通过 `gotrontrx` 简单介绍波场 `tron` 的入门知识，以下是个简单的入门教程。

### 第一步-创建钱包：

您可以使用这段代码，离线创建钱包：
https://gist.github.com/motopig/c680f53897429fd15f5b3ca9aa6f6ed2
把代码拷贝下来，在自己电脑里运行。

当然其它的也能创建钱包，需要保证是 **离线** 的，而不是通过在线网页创建钱包。

区块链的钱包创建是离线的，你能使用任意你觉得趁手的离线工具生成你的钱包（任何通过网页在线创建私钥的行为都是耍流氓）

### 第二步-去链上查钱包信息-主要是余额信息：

创建完即可在测试链查看资产情况：
https://shasta.tronscan.org/#/address/TBYHGsFkshasvB3R6Zys4627h98owvUNFn

只要你创建完钱包，在任意波场链（即 mainNet 主网，shasta 测试网，nile 测试网）查询都是能查到信息（但资产都是空的）

注意保存好你的私钥

你能在任意波场链里给这个地址转账，结果都会让它有资金，也都可以向外转账。

但是请注意通常一个私钥只用在一个网络，特别是测试网络的私钥不要与主网相同，这样能避免私钥泄露（只能在某种程度上避免）。

### 第三步-领取测试币TRX

在波场官方中文客服群里即可领取测试币5000个TRX：
https://t.me/TronOfficialTechSupport

当然在官方英文Telegram群里也能领取5000个TRX：
@TronOfficialDevelopersGroupEn

具体操作进群以后输入消息（在以上两个群里均可）:
```
!help
```
你自然就会啦，即可给自己的测试钱包领5000个TRX。

### 第四步-使用本SDK进行转账等操作吧

当然为避免私钥泄露，只建议使用测试链钱包验证以下的Demo。

⬆️ **源码：** [基本 TRX 转账演示](internal/demos/demo1x/main.go)

---

## 免责声明：

数字货币都是骗局

都是以空气币掠夺平民财富

没有公平正义可言

数字货币对中老年人是极不友好的，因为他们没有机会接触这类披着高科技外衣的割韭菜工具

数字货币对青少年也是极不友好的，因为当他们接触的时候，前面的人已经占据了大量的资源

因此妄图以数字货币，比如稍微主流的 BTC ETH TRX 代替世界货币的操作，都是不可能实现的

都不过是早先持有数字货币的八零后们的无耻幻想

扪心自问，持有几千甚至数万个比特币的人会觉得公平吗，其实不会的

因此未来还会有新事物来代替它们，而我现在也不过只是了解其中的技术，仅此而已。

该项目仅以技术学习和探索为目的而存在。

该项目作者坚定持有"坚决抵制数字货币"的立场。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![starring](https://starchart.cc/yyle88/gotrontrx.svg?variant=adaptive)](https://starchart.cc/yyle88/gotrontrx)
