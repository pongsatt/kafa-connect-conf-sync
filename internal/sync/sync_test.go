package sync_test

import (
	"testing"

	"github.com/pongsatt/kafa-connect-conf-sync/internal/model"
	"github.com/pongsatt/kafa-connect-conf-sync/internal/sync"
)

func TestEmptyConfig(t *testing.T) {
	newConfigs := []*model.Config{}

	oldConfigs := []*model.Config{}

	diffs, err := sync.Configs(newConfigs, oldConfigs)

	if err != nil {
		t.Error(err)
	}

	if len(diffs.Creates) != 0 {
		t.Errorf("Expect creates cases > 0 but %d", len(diffs.Creates))
	}

	if len(diffs.Updates) != 0 {
		t.Errorf("Expect updates cases > 0 but %d", len(diffs.Updates))
	}

	if len(diffs.Deletes) != 0 {
		t.Errorf("Expect deletes cases > 0 but %d", len(diffs.Deletes))
	}
}
func TestFirstTimeConfigs(t *testing.T) {
	newConfigs := []*model.Config{
		{
			Name:       "test1",
			ConfigHash: "testhash",
		},
		{
			Name:       "test2",
			ConfigHash: "testhash",
		},
		{
			Name:       "test3",
			ConfigHash: "testhash",
		},
	}

	oldConfigs := []*model.Config{}

	diffs, err := sync.Configs(newConfigs, oldConfigs)

	if err != nil {
		t.Error(err)
	}

	if len(diffs.Creates) != 3 {
		t.Errorf("Expect creates cases > 0 but %d", len(diffs.Creates))
	}

	if len(diffs.Updates) != 0 {
		t.Errorf("Expect updates cases > 0 but %d", len(diffs.Updates))
	}

	if len(diffs.Deletes) != 0 {
		t.Errorf("Expect deletes cases > 0 but %d", len(diffs.Deletes))
	}
}
func TestSyncConfigys(t *testing.T) {
	newConfigs := []*model.Config{
		{
			Name:       "test1",
			ConfigHash: "testhash",
		},
		{
			Name:       "test2",
			ConfigHash: "testhash",
		},
		{
			Name:       "test3",
			ConfigHash: "testhash",
		},
	}

	oldConfigs := []*model.Config{
		{
			Name:       "test1",
			ConfigHash: "testhash",
		},
		{
			Name:       "test2",
			ConfigHash: "testhash_updated",
		},
		{
			Name:       "test4",
			ConfigHash: "testhash",
		},
	}

	diffs, err := sync.Configs(newConfigs, oldConfigs)

	if err != nil {
		t.Error(err)
	}

	if len(diffs.Creates) != 1 {
		t.Errorf("Expect creates cases > 0 but %d", len(diffs.Creates))
	}

	if len(diffs.Updates) != 1 {
		t.Errorf("Expect updates cases > 0 but %d", len(diffs.Updates))
	}

	if len(diffs.Deletes) != 1 {
		t.Errorf("Expect deletes cases > 0 but %d", len(diffs.Deletes))
	}
}
