package os_api_model

import "time"

type PortInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	NetworkID    string `json:"network_id"`
	ProjectID    string `json:"project_id"`
	TenantID     string `json:"tenant_id"`
	MacAddress   string `json:"mac_address"`
	AdminStateUp bool   `json:"admin_state_up"`
	Status       string `json:"status"`
	DeviceID     string `json:"device_id"`
	DeviceOwner  string `json:"device_owner"`
	FixedIps     []struct {
		SubnetID  string `json:"subnet_id"`
		IPAddress string `json:"ip_address"`
	} `json:"fixed_ips"`
	AllowedAddressPairs []interface{} `json:"allowed_address_pairs"`
	ExtraDhcpOpts       []interface{} `json:"extra_dhcp_opts"`
	SecurityGroups      []interface{} `json:"security_groups"`
	Description         string        `json:"description"`
	BindingVnicType     string        `json:"binding:vnic_type"`
	BindingProfile      struct {
	} `json:"binding:profile"`
	BindingHostID     string `json:"binding:host_id"`
	BindingVifType    string `json:"binding:vif_type"`
	BindingVifDetails struct {
		Connectivity  string `json:"connectivity"`
		PortFilter    bool   `json:"port_filter"`
		OvsHybridPlug bool   `json:"ovs_hybrid_plug"`
		DatapathType  string `json:"datapath_type"`
		BridgeName    string `json:"bridge_name"`
		BoundDrivers  struct {
			Num0 string `json:"0"`
		} `json:"bound_drivers"`
	} `json:"binding:vif_details"`
	QosPolicyID         interface{}   `json:"qos_policy_id"`
	QosNetworkPolicyID  interface{}   `json:"qos_network_policy_id"`
	PortSecurityEnabled bool          `json:"port_security_enabled"`
	ResourceRequest     interface{}   `json:"resource_request"`
	IPAllocation        string        `json:"ip_allocation"`
	Tags                []interface{} `json:"tags"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
	RevisionNumber      int           `json:"revision_number"`
}

type Ports struct {
	Ports []*PortInfo `json:"ports"`
}
