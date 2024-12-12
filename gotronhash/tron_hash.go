package gotronhash

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func GetTxHash(txRaw *core.TransactionRaw) (string, error) {
	raw, err := proto.Marshal(txRaw)
	if err != nil {
		return "", errors.WithMessage(err, "proto.marshal raw_tx is wrong")
	}
	sha256h := sha256.New()
	sha256h.Write(raw)
	rawHash := sha256h.Sum(nil)
	return hex.EncodeToString(rawHash), nil
}
