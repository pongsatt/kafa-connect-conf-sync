package kafaconnect

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Read reads configs data from kafa-connect
func (kc *Client) Read() ([][]byte, error) {
	connectors, err := kc.getConnectors()

	if err != nil {
		return nil, err
	}

	results := make([][]byte, 0)

	for _, connector := range connectors {
		data, err := kc.getConfig(connector)

		if err != nil {
			return nil, err
		}

		results = append(results, data)
	}

	return results, nil
}

func (kc *Client) getConnectors() ([]string, error) {
	resp, err := http.Get(fmt.Sprintf("%s", kc.baseURL))

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var connectors []string
	err = json.Unmarshal(data, &connectors)

	return connectors, err
}

func (kc *Client) getConfig(connector string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", kc.baseURL, connector))

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}
