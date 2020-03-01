package sync

import "github.com/pongsatt/kafa-connect-conf-sync/internal/model"

// Diffs stores configuration differences
type Diffs struct {
	Creates []model.Config
	Updates []model.Config
	Deletes []model.Config
}

// Configs sync all configurations and report the differences
func Configs(newConfigs []*model.Config, oldConfigs []*model.Config) (*Diffs, error) {
	oldConfigsMap := toMap(oldConfigs)
	toCreates := make([]model.Config, 0)
	toUpdates := make([]model.Config, 0)

	for _, newConfig := range newConfigs {
		oldConfig, match := oldConfigsMap[newConfig.Name]

		if !match {
			toCreates = append(toCreates, *newConfig)
			continue
		}

		if newConfig.ConfigHash != oldConfig.ConfigHash {
			toUpdates = append(toUpdates, *newConfig)
		}

		delete(oldConfigsMap, newConfig.Name)

	}

	toDeletes := toList(oldConfigsMap)

	diffs := &Diffs{
		Creates: toCreates,
		Updates: toUpdates,
		Deletes: toDeletes,
	}

	return diffs, nil
}

func toList(configsMap map[string]*model.Config) []model.Config {
	list := make([]model.Config, 0)

	for _, v := range configsMap {
		list = append(list, *v)
	}
	return list
}

func toMap(configs []*model.Config) map[string]*model.Config {
	configMap := make(map[string]*model.Config)

	for _, config := range configs {
		configMap[config.Name] = config
	}

	return configMap
}
