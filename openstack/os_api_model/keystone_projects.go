package os_api_model

type Projects struct {
	Projects []*ProjectInfo
}

type ProjectInfo struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	DomainID    string        `json:"domain_id"`
	Description string        `json:"description"`
	Enabled     bool          `json:"enabled"`
	ParentID    string        `json:"parent_id"`
	IsDomain    bool          `json:"is_domain"`
	Tags        []interface{} `json:"tags"`
	Options     interface{}   `json:"options"`
}
