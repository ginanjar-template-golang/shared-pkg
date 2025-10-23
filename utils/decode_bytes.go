package utils

import (
	"encoding/base64"
	"encoding/json"
)

func DecodeBytesToJSON(b []byte) (any, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(b))
	if err == nil {
		return tryUnmarshal(decoded)
	}

	return tryUnmarshal(b)
}

func tryUnmarshal(data []byte) (any, error) {
	var result any
	if err := json.Unmarshal(data, &result); err != nil {
		return string(data), nil
	}
	return result, nil
}
