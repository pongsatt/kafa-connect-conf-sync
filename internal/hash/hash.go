package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// ComputeHash computes hash for config content
func ComputeHash(content interface{}) (string, error) {
	data, err := json.Marshal(content)

	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:]), nil
}
