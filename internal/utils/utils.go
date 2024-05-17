package utils

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

func SoftNeat(v interface{}) string {
	data, err := NeatBytes(v)
	if err != nil {
		return "" //when the result is empty string, means wrong
	}
	return string(data)
}

func NeatBytes(v interface{}) ([]byte, error) {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, errors.WithMessage(err, "marshal object is wrong")
	}
	return data, nil
}

func NeatStringFromBytes(data []byte) (string, error) {
	res, err := NeatFromBytes(data)
	if err != nil {
		return string(data), err
	}
	return string(res), nil
}

// NeatFromBytes 把json的bytes变为neat的，当然假如不是Json就会报错，这块常用于测试的打印
func NeatFromBytes(data []byte) ([]byte, error) {
	var ob bytes.Buffer
	if err := json.Indent(&ob, data, "", "\t"); err != nil {
		return data, err
	}
	return ob.Bytes(), nil
}
