package os_api_model

import "time"

type Networks struct {
	Networks []*NetworkInfo `json:"networks"`
}

type Network struct {
	Network NetworkInfo `json:"network"`
}

type NetworkInfo struct {
	AdminStateUp            bool          `json:"admin_state_up"`
	AvailabilityZoneHints   []string      `json:"availability_zone_hints"`
	AvailabilityZones       []string      `json:"availability_zones"`
	CreatedAt               string        `json:"created_at"`
	DNSDomain               string        `json:"dns_domain"`
	ID                      string        `json:"id"`
	IPv4AddressScope        string        `json:"ipv4_address_scope"`
	IPv6AddressScope        string        `json:"ipv6_address_scope"`
	L2Adjacency             bool          `json:"l2_adjacency"`
	MTU                     int           `json:"mtu"`
	Name                    string        `json:"name"`
	PortSecurityEnabled     bool          `json:"port_security_enabled"`
	ProjectID               string        `json:"project_id"`
	QoSPolicyID             string        `json:"qos_policy_id"`
	RevisionNumber          int           `json:"revision_number"`
	RouterExternal          bool          `json:"router:external"`
	Shared                  bool          `json:"shared"`
	Status                  string        `json:"status"`
	Subnets                 []string      `json:"subnets"`
	TenantID                string        `json:"tenant_id"`
	UpdatedAt               time.Time     `json:"updated_at"`
	VLANTransparent         bool          `json:"vlan_transparent"`
	Description             string        `json:"description"`
	ProviderNetworkType     string        `json:"provider:network_type"`
	ProviderPhysicalNetwork string        `json:"provider:physical_network"`
	ProviderSegmentationID  int           `json:"provider:segmentation_id"`
	Tags                    []interface{} `json:"tags"`
	IsDefault               bool          `json:"is_default"`
}
