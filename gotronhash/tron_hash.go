package gotronhash

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// GetTxHash calculates the SHA256 hash of a transaction and encodes it as a hex-string.
func GetTxHash(txRaw *core.TransactionRaw) (string, error) {
	rawHash, err := RawTxHash(txRaw)
	if err != nil {
		return "", errors.Wrap(err, "wrong to get raw tx hash")
	}
	return hex.EncodeToString(rawHash), nil
}

// RawTxHash calculates the raw SHA256 hash of a transaction.
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
