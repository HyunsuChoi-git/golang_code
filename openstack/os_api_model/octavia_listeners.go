package os_api_model

type ListenerDetail struct {
	Listener Listener `json:"listener"`
}

type Listener struct {
	Id                          string `json:"id"`
	Name                        string `json:"name"`
	Description                 string `json:"description"`
	Provisioning_status         string `json:"provisioning_status"`
	Operating_status            string `json:"operating_status"`
	admin_state_up              string
	protocol                    string
	protocol_port               string
	connection_limit            string
	default_tls_container_ref   string
	sni_container_refs          string
	project_id                  string
	default_pool_id             string
	policies                    interface{} `json:"17policies"`
	insert_headers              interface{}
	created_at                  string
	updated_at                  string
	loadbalancers               interface{}
	timeout_client_data         int
	timeout_member_connect      int
	timeout_member_data         int
	timeout_tcp_inspect         int
	tags                        interface{}
	client_ca_tls_container_ref string
	client_authentication       string
	client_crl_container_ref    string
	tls_versions                interface{}
	alpn_protocols              interface{}
	tenant_id                   string
}
