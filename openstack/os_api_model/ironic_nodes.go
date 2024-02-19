package os_api_model

type Node struct {
	Uuid            string
	Power_state     string `json:"power_state"`
	istrance_uuid   string `json:"istrance_uuid"`
	links           interface{}
	maintenance     string
	name            string
	provision_state string
}

type NodeList struct {
	Nodes []*Node
}

type Nodes struct {
	Uuid                   string `json:"uuid"`
	istrance_uuid          string `json:"istrance_uuid"`
	Power_state            string `json:"power_state"`
	target_power_state     string
	last_error             string
	provision_updated_at   string
	maintenance            string
	maintenance_reason     string
	target_provision_state string
	console_enabled        string
	instance_info          interface{}
	driver                 string
	deiver_info            interface{}
	extra                  interface{}
	propertise             interface{}
	chassis_uuid           string
	links                  interface{}
	ports                  interface{}
	updated_at             string
	created_at             string
}
