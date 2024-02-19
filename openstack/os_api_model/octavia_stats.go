package os_api_model

type LoadbalancerStats struct {
	Stats Stats `json:"stats"`
}

type ListenerStats struct {
	Stats Stats `json:"stats"`
}

type Stats struct {
	Bytes_in           float64 `json:"bytes_in"`
	Bytes_out          float64 `json:"bytes_out"`
	Active_connections float64 `json:"active_connections"`
	Total_connections  float64 `json:"total_connections"`
	Request_errors     float64 `json:"request_errors"`
}
