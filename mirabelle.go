package client

type Event struct {
	Host        string            `json:"host"`
	Service     string            `json:"service"`
	Time        int64             `json:"time"`
	Name        string            `json:"name"`
	State       string            `json:"state"`
	Description string            `json:"description"`
	Tags        []string          `json:"tags"`
	Attributes  map[string]string `json:"attributes"`
}
