package config

import (
	"encoding/json"

	"github.com/pongsatt/kafa-connect-conf-sync/internal/model"
)

// Reader reads configs
type Reader interface {
	Read() ([][]byte, error)
}

// Read all configs from a given provider
func Read(reader Reader) ([]*model.Config, error) {
	rawConfigs, err := reader.Read()

	if err != nil {
		return nil, err
	}

	configs := make([]*model.Config, 0)

	for _, data := range rawConfigs {
		config, err := readConfigFromRaw(data)

		if err != nil {
			return nil, err
		}

		configs = append(configs, config)
	}
	return configs, nil
}

func readConfigFromRaw(data []byte) (*model.Config, error) {
	var config *model.Config

	err := json.Unmarshal(data, &config)

	if err != nil {
		return nil, err
	}

	config.ConfigHash, err = ComputeHash(config)

	return config, err
}
