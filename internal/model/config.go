package model

// Config is a model for config
type Config struct {
	Name       string
	ConfigHash string
	Config     map[string]string
}

// Equals compares two configs
func (config *Config) Equals(other *Config) bool {
	if other == nil && config == nil {
		return true
	}

	if other == nil || config == nil {
		return false
	}

	if other.Name != config.Name {
		return false
	}

	if other.ConfigHash != config.ConfigHash {
		return false
	}

	return true
}
