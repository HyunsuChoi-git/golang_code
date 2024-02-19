package os_api_model

type RegionInfo struct {
	ID             string `json:"id"`
	Description    string `json:"description"`
	ParentRegionID string `json:"parent_region_id"`

	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type Regions struct {
	Links   interface{}   `json:"links"`
	Regions []*RegionInfo `json:"regions"`
}
