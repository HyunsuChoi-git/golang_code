package os_api_model

import "time"

type SubnetInfo struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	ProjectID       string      `json:"project_id"`
	TenantID        string      `json:"tenant_id"`
	NetworkID       string      `json:"network_id"`
	SegmentID       string      `json:"segment_id"`
	IPVersion       int         `json:"ip_version"`
	SubnetpoolID    interface{} `json:"subnetpool_id"`
	EnableDhcp      bool        `json:"enable_dhcp"`
	Ipv6RaMode      interface{} `json:"ipv6_ra_mode"`
	Ipv6AddressMode interface{} `json:"ipv6_address_mode"`
	GatewayIP       string      `json:"gateway_ip"`
	Cidr            string      `json:"cidr"`
	AllocationPools []struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"allocation_pools"`
	HostRoutes     []interface{} `json:"host_routes"`
	DNSNameservers []interface{} `json:"dns_nameservers"`
	Description    string        `json:"description"`
	ServiceTypes   []interface{} `json:"service_types"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	RevisionNumber int           `json:"revision_number"`
	Tags           []interface{} `json:"tags"`
}

type Subnets struct {
	Subnets []*SubnetInfo `json:"subnets"`
}
