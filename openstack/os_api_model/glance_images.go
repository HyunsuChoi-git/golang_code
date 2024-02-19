package os_api_model

type ImageList struct {
	Images []*Image `json:"images"`
}

type Image struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Size   int    `json:"size"`
}
