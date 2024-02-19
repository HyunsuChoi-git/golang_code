package os_api_model

import "time"

type RouterInfo struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	ProjectID           string `json:"project_id"`
	TenantID            string `json:"tenant_id"`
	AdminStateUp        bool   `json:"admin_state_up"`
	Status              string `json:"status"`
	ExternalGatewayInfo struct {
		NetworkID        string `json:"network_id"`
		ExternalFixedIps []struct {
			SubnetID  string `json:"subnet_id"`
			IPAddress string `json:"ip_address"`
		} `json:"external_fixed_ips"`
		EnableSnat bool `json:"enable_snat"`
	} `json:"external_gateway_info"`
	Description           string        `json:"description"`
	AvailabilityZones     []string      `json:"availability_zones"`
	Distributed           bool          `json:"distributed"`
	Ha                    bool          `json:"ha"`
	AvailabilityZoneHints []interface{} `json:"availability_zone_hints"`
	Routes                []interface{} `json:"routes"`
	FlavorID              string        `json:"flavor_id"`
	Tags                  []interface{} `json:"tags"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
	RevisionNumber        int           `json:"revision_number"`
}

type Routers struct {
	Routers []*RouterInfo `json:"routers"`
}
