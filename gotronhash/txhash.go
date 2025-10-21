// Package gotronhash provides transaction hash calculation utilities for TRON blockchain
// Implements SHA256 hashing for transaction data
//
// gotronhash 提供 TRON 区块链的交易哈希计算工具
// 实现交易数据的 SHA256 哈希计算
package gotronhash

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// GetTxHash calculates the SHA256 hash of a transaction and encodes it as a hex-string
// Returns the transaction hash in hexadecimal format
//
// GetTxHash 计算交易的 SHA256 哈希值并编码为十六进制字符串
// 返回十六进制格式的交易哈希
func GetTxHash(txRaw *core.TransactionRaw) (string, error) {
	rawHash, err := RawTxHash(txRaw)
	if err != nil {
		return "", errors.Wrap(err, "wrong to get raw tx hash")
	}
	return hex.EncodeToString(rawHash), nil
}

// RawTxHash calculates the raw SHA256 hash of a transaction
// Returns the hash as raw bytes
//
// RawTxHash 计算交易的原始 SHA256 哈希值
// 返回原始字节格式的哈希
func RawTxHash(txRaw *core.TransactionRaw) ([]byte, error) {
	raw, err := proto.Marshal(txRaw)
	if err != nil {
		return nil, errors.WithMessage(err, "wrong to serialize transaction raw using proto")
	}
	sha256h := sha256.New()
	sha256h.Write(raw)
	rawHash := sha256h.Sum(nil)
	return rawHash, nil
}
