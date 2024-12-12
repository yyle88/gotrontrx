package gotronsign

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbsobreira/gotron-sdk/pkg/proto/core"
	"github.com/pkg/errors"
	"github.com/yyle88/gotrontrx/gotronhash"
	"google.golang.org/protobuf/encoding/protojson"
)

func SignMessage(privateKeyHex string, message json.RawMessage) ([]byte, error) {
	txRaw := core.TransactionRaw{}
	if err := protojson.Unmarshal(message, &txRaw); err != nil {
		return nil, errors.WithMessage(err, "wrong to unmarshal JSON message into TransactionRaw")
	}
	return Sign(privateKeyHex, &txRaw)
}

func Sign(privateKeyHex string, txRaw *core.TransactionRaw) ([]byte, error) {
	// Get transaction hash using GetTxHash
	txHash, err := gotronhash.RawTxHash(txRaw)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to calculate transaction hash")
	}

	// Parse the private key from Hex
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, errors.New("wrong private key: failed to parse Hex to ECDSA")
	}

	// Generate the signature
	signatureBytes, err := crypto.Sign(txHash, ecdsaPrivateKey)
	if err != nil {
		return nil, errors.New("wrong to generate transaction signature")
	}

	return signatureBytes, nil
}
