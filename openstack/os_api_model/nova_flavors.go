package os_api_model

type FlavorInfo struct {
	OSFLVDISABLEDDisabled  bool   `json:"OS-FLV-DISABLED:disabled"`
	Disk                   int    `json:"disk"`
	OSFLVEXTDATAEphemeral  int    `json:"OS-FLV-EXT-DATA:ephemeral"`
	OsFlavorAccessIsPublic bool   `json:"os-flavor-access:is_public"`
	ID                     string `json:"id"`
	Links                  []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Name        string      `json:"name"`
	RAM         int         `json:"ram"`
	Swap        string      `json:"swap"`
	Vcpus       int         `json:"vcpus"`
	RxtxFactor  float64     `json:"rxtx_factor"`
	Description string      `json:"description"`
	ExtraSpecs  interface{} `json:"extra_specs"`
}

type Flavors struct {
	Flavors []*FlavorInfo `json:"flavors"`
}
