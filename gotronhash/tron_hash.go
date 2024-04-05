package gotronhash

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func GetTxHashHex(rawTx *core.TransactionRaw) (string, error) {
	raw, err := proto.Marshal(rawTx)
	if err != nil {
		return "", errors.WithMessage(err, "proto.marshal raw_tx is wrong")
	}
	return GetRawTxHashHex(raw), nil
}

func GetRawTxHashHex(raw []byte) string {
	return hex.EncodeToString(GetRawTxHash(raw))
}

func GetRawTxHash(raw []byte) []byte {
	h256h := sha256.New()
	h256h.Write(raw)
	rawHash := h256h.Sum(nil)
	return rawHash
}
