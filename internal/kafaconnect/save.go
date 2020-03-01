package kafaconnect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pongsatt/kafa-connect-conf-sync/internal/model"
)

// Save creates or updates config
func (kc *Client) Save(config *model.Config) (*model.Config, error) {
	inputData, err := json.Marshal(config.Config)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s/config", kc.baseURL, config.Name), bytes.NewBuffer(inputData))

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var newConfig *model.Config
	err = json.Unmarshal(data, &newConfig)

	return newConfig, err
}
