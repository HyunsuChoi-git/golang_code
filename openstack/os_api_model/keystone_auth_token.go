package os_api_model

import "time"

type ResAuthBody struct {
	Token ResAuthToken
}

type ResAuthToken struct {
	methodes   []string
	user       []string
	audit_ids  []string
	Expires_at time.Time `json:"expires_at"`
	Issued_at  time.Time `json:"issued_at"`
	roles      []interface{}
	system     interface{}
	domain     interface{}
	Catalog    []ResAuthTokenCatalog `json:"catalog"`
}
type ResAuthTokenCatalog struct {
	Endpoints []ResAuthTokenCatalogEndPoints
	Id        string
	TypeName  string `json:"type"`
	Name      string
}
type ResAuthTokenCatalogEndPoints struct {
	id            string
	InterfaceName string `json:"interface"`
	region_id     string
	Url           string
	region        string
	service_id    string
	enabled       bool
}
