package os_api_model

type UserInfo struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	DomainID          string `json:"domain_id"`
	DefaultProjectID  string `json:"default_project_id"`
	Enabled           bool   `json:"enabled"`
	PasswordExpiresAt string `json:"password_expires_at"`

	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type Users struct {
	Links interface{} `json:"links"`
	Users []*UserInfo `json:"users"`
}
