package gotronsign

import (
	"crypto/sha256"
	"encoding/json"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func SignMessage(private string, message json.RawMessage) ([]byte, error) {
	rawTx := core.TransactionRaw{}
	if err := protojson.Unmarshal(message, &rawTx); err != nil {
		return nil, errors.WithMessage(err, "protojson unmarshal message is wrong")
	}
	return Sign(private, &rawTx)
}

func Sign(private string, rawTx *core.TransactionRaw) ([]byte, error) {
	sign, err := proto.Marshal(rawTx)
	if err != nil {
		return nil, errors.WithMessage(err, "proto marshal raw_tx is wrong")
	}
	h256h := sha256.New()
	h256h.Write(sign)
	hash := h256h.Sum(nil)
	priKey, err := crypto.HexToECDSA(private)
	if err != nil {
		return nil, errors.New("private is wrong")
	}
	signature, err := crypto.Sign(hash, priKey)
	if err != nil {
		return nil, errors.New("sign is wrong")
	}
	return signature, nil
}
