package os_api_model

type LoadbalancerList struct {
	Loadbalancers []*Loadbalancer `json:"loadbalancers"`
}

type LoadbalancerDetail struct {
	Loadbalancer Loadbalancer `json:"loadbalancer"`
}

type Loadbalancer struct {
	ID               string              `json:"id"`
	Name             string              `json:"name"`
	Operating_status string              `json:"operating_status"`
	Listeners        []map[string]string `json:"listeners"`

	description         string
	admin_state_up      bool
	project_id          string
	provisioning_status string
	flavor_id           string
	vip_subnet_id       string
	vip_address         string
	vip_network_id      string
	vip_port_id         string
	additional_vips     []string
	provider            string
	pools               interface{}
	created_at          string
	updated_at          string
	vip_qos_policy_id   string
	availability_zone   string
	tags                []string
	tenant_id           string
}
