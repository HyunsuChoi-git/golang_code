package os_api_model

import "time"

type VolumeInfo struct {
	Attachments        []interface{} `json:"attachments"`
	AvailabilityZone   string        `json:"availability_zone"`
	Bootable           string        `json:"bootable"`
	ConsistencygroupID interface{}   `json:"consistencygroup_id"`
	CreatedAt          *time.Time    `json:"created_at"`
	Description        string        `json:"description"`
	Encrypted          bool          `json:"encrypted"`
	ID                 string        `json:"id"`
	Links              []struct {
		Href string `json:"href"`
		Rel  string `json:"rel"`
	} `json:"links"`
	Metadata struct {
	} `json:"metadata"`
	MigrationStatus           interface{} `json:"migration_status"`
	Multiattach               bool        `json:"multiattach"`
	Name                      string      `json:"name"`
	OsVolHostAttrHost         string      `json:"os-vol-host-attr:host"`
	OsVolMigStatusAttrMigstat string      `json:"os-vol-mig-status-attr:migstat"`
	OsVolMigStatusAttrNameID  string      `json:"os-vol-mig-status-attr:name_id"`
	OsVolTenantAttrTenantID   string      `json:"os-vol-tenant-attr:tenant_id"`
	ReplicationStatus         interface{} `json:"replication_status"`
	Size                      int         `json:"size"`
	SnapshotID                interface{} `json:"snapshot_id"`
	SourceVolid               interface{} `json:"source_volid"`
	Status                    string      `json:"status"`
	UpdatedAt                 *time.Time  `json:"updated_at"`
	UserID                    string      `json:"user_id"`
	VolumeType                string      `json:"volume_type"`
	VolumeTypeID              string      `json:"volume_type_id"`
	ProviderID                interface{} `json:"provider_id"`
	GroupID                   interface{} `json:"group_id"`
	ServiceUUID               interface{} `json:"service_uuid"`
	SharedTargets             bool        `json:"shared_targets"`
	ClusterName               interface{} `json:"cluster_name"`
	ConsumesQuota             bool        `json:"consumes_quota"`
}

type VolumesDetail struct {
	Volumes []*VolumeInfo `json:"volumes"`
}
