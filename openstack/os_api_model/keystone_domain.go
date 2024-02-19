package os_api_model

type DomainInfo struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Enabled     bool          `json:"enabled"`
	Tags        []interface{} `json:"tags"`
	Options     interface{}   `json:"options"`

	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type Domains struct {
	Links   interface{}   `json:"links"`
	Domains []*DomainInfo `json:"domains"`
}
