package os_api_model

type GroupInfo struct {
	ID          string `json:"id"`
	DomainID    string `json:"domain_id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type Groups struct {
	Links  interface{}  `json:"links"`
	Groups []*GroupInfo `json:"groups"`
}
