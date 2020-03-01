package kafaconnect

import (
	"fmt"
	"net/http"

	"github.com/pongsatt/kafa-connect-conf-sync/internal/model"
)

// Delete delete config
func (kc *Client) Delete(config *model.Config) error {
	// Create client
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", kc.baseURL, config.Name), nil)

	// Fetch Request
	_, err = client.Do(req)
	return err
}
