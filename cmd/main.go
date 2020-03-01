package main

import (
	"fmt"
	"os"
	"path"

	"github.com/pongsatt/kafa-connect-conf-sync/internal/config"
	"github.com/pongsatt/kafa-connect-conf-sync/internal/file"
	"github.com/pongsatt/kafa-connect-conf-sync/internal/kafaconnect"
	"github.com/pongsatt/kafa-connect-conf-sync/internal/sync"
)

func main() {
	// list configs from files
	rootPath, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	fileReader := file.NewFileReader(path.Join(rootPath, "/files"))
	fileConfigs, err := config.Read(fileReader)

	if err != nil {
		panic(err)
	}

	fmt.Printf("found files %d\n", len(fileConfigs))

	// list configs from server
	baseURL := "http://localhost:8083/connectors"
	kcClient := kafaconnect.NewClient(baseURL)
	remoteConfigs, err := config.Read(kcClient)

	if err != nil {
		panic(err)
	}

	fmt.Printf("found remote configs %d\n", len(remoteConfigs))

	// compare configs
	diffs, err := sync.Configs(fileConfigs, remoteConfigs)

	if err != nil {
		panic(err)
	}

	fmt.Printf("configs to create %d\n", len(diffs.Creates))
	// create configs to server
	if len(diffs.Creates) > 0 {
		for _, config := range diffs.Creates {
			_, err := kcClient.Save(&config)

			if err != nil {
				panic(err)
			}

			fmt.Printf("%s created\n", config.Name)
		}
	}

	fmt.Printf("configs to update %d\n", len(diffs.Updates))
	// update configs to server
	if len(diffs.Updates) > 0 {
		for _, config := range diffs.Updates {
			_, err := kcClient.Save(&config)

			if err != nil {
				panic(err)
			}

			fmt.Printf("%s updated\n", config.Name)
		}
	}

	fmt.Printf("configs to delete %d\n", len(diffs.Deletes))
	// delete configs to server
	if len(diffs.Deletes) > 0 {
		for _, config := range diffs.Deletes {
			err := kcClient.Delete(&config)

			if err != nil {
				panic(err)
			}

			fmt.Printf("%s deleted\n", config.Name)
		}
	}

	fmt.Println("All done.")
}
