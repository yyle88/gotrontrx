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

func SignMessage(privateHex string, message json.RawMessage) ([]byte, error) {
	txRaw := core.TransactionRaw{}
	if err := protojson.Unmarshal(message, &txRaw); err != nil {
		return nil, errors.WithMessage(err, "protojson unmarshal message is wrong")
	}
	return Sign(privateHex, &txRaw)
}

func Sign(privateHex string, txRaw *core.TransactionRaw) ([]byte, error) {
	sign, err := proto.Marshal(txRaw)
	if err != nil {
		return nil, errors.WithMessage(err, "proto marshal raw_tx is wrong")
	}
	h256h := sha256.New()
	h256h.Write(sign)
	hash := h256h.Sum(nil)
	priKey, err := crypto.HexToECDSA(privateHex)
	if err != nil {
		return nil, errors.New("private is wrong")
	}
	signature, err := crypto.Sign(hash, priKey)
	if err != nil {
		return nil, errors.New("sign is wrong")
	}
	return signature, nil
}
