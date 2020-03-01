package kafaconnect

// Client provides kafka connect client
type Client struct {
	baseURL string
}

// NewClient creates new client
func NewClient(baseURL string) *Client {
	if baseURL == "" {
		baseURL = "http://localhost:8083/connectors"
	}

	return &Client{
		baseURL: baseURL,
	}
}
