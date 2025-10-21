// Package gotronsign provides transaction signing functionality for TRON blockchain
// Implements ECDSA signature generation for transactions
//
// gotronsign 提供 TRON 区块链的交易签名功能
// 实现交易的 ECDSA 签名生成
package gotronsign

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/pkg/errors"
	"github.com/yyle88/gotrontrx/gotronhash"
	"google.golang.org/protobuf/encoding/protojson"
)

// SignMessage signs a JSON-formatted transaction message with the given private key
// Unmarshals the JSON into TransactionRaw before signing
//
// SignMessage 使用给定的私钥签名 JSON 格式的交易消息
// 签名前将 JSON 解组为 TransactionRaw
func SignMessage(privateKeyHex string, message json.RawMessage) ([]byte, error) {
	txRaw := core.TransactionRaw{}
	if err := protojson.Unmarshal(message, &txRaw); err != nil {
		return nil, errors.WithMessage(err, "wrong to unmarshal JSON message into TransactionRaw")
	}
	return Sign(privateKeyHex, &txRaw)
}

// Sign generates an ECDSA signature for a transaction using the provided private key
// Returns the signature bytes that can be attached to the transaction
//
// Sign 使用提供的私钥为交易生成 ECDSA 签名
// 返回可附加到交易的签名字节
func Sign(privateKeyHex string, txRaw *core.TransactionRaw) ([]byte, error) {
	// Calculate transaction hash
	// 计算交易哈希
	txHash, err := gotronhash.RawTxHash(txRaw)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to calculate transaction hash")
	}

	// Parse private key from hexadecimal format
	// 从十六进制格式解析私钥
	privateECDSA, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, errors.New("wrong private key: failed to parse Hex to ECDSA")
	}

	// Generate ECDSA signature
	// 生成 ECDSA 签名
	signatureBytes, err := crypto.Sign(txHash, privateECDSA)
	if err != nil {
		return nil, errors.New("wrong to generate transaction signature")
	}

	return signatureBytes, nil
}
