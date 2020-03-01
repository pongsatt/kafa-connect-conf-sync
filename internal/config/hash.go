package config

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"

	"github.com/pongsatt/kafa-connect-conf-sync/internal/model"
)

// ComputeHash computes hash for config content
func ComputeHash(config *model.Config) (string, error) {

	if config.Config == nil {
		return "", nil
	}

	keys := make([]string, 0)

	for k := range config.Config {
		if strings.ToLower(k) != "name" {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)

	var all bytes.Buffer

	for _, k := range keys {
		v := config.Config[k]

		all.WriteString(k)
		all.WriteString(v)
	}

	hash := sha256.Sum256(all.Bytes())

	return hex.EncodeToString(hash[:]), nil
}
