package os_api_model

import (
	"time"
)

type ServerInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	TenantID string `json:"tenant_id"`
	UserID   string `json:"user_id"`
	Metadata struct {
	} `json:"metadata"`
	HostID string `json:"hostId"`
	//Image  struct {
	//	ID    string `json:"id"`
	//	Links []struct {
	//		Rel  string `json:"rel"`
	//		Href string `json:"href"`
	//	} `json:"links"`
	//} `json:"image"`
	Image  interface{}
	Flavor struct {
		ID    string `json:"id"`
		Links []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
	} `json:"flavor"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
	Addresses map[string][]struct {
		Version            int    `json:"version"`
		Addr               string `json:"addr"`
		OSEXTIPSType       string `json:"OS-EXT-IPS:type"`
		OSEXTIPSMACMacAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`
	} `json:"addresses"`
	AccessIPv4 string `json:"accessIPv4"`
	AccessIPv6 string `json:"accessIPv6"`
	Links      []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	OSDCFDiskConfig                  string      `json:"OS-DCF:diskConfig"`
	Progress                         int         `json:"progress"`
	OSEXTAZAvailabilityZone          string      `json:"OS-EXT-AZ:availability_zone"`
	ConfigDrive                      string      `json:"config_drive"`
	KeyName                          string      `json:"key_name"`
	OSSRVUSGLaunchedAt               string      `json:"OS-SRV-USG:launched_at"`
	OSSRVUSGTerminatedAt             interface{} `json:"OS-SRV-USG:terminated_at"`
	OSEXTSRVATTRHost                 string      `json:"OS-EXT-SRV-ATTR:host"`
	OSEXTSRVATTRInstanceName         string      `json:"OS-EXT-SRV-ATTR:instance_name"`
	OSEXTSRVATTRHypervisorHostname   string      `json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`
	OSEXTSTSTaskState                interface{} `json:"OS-EXT-STS:task_state"`
	OSEXTSTSVMState                  string      `json:"OS-EXT-STS:vm_state"`
	OSEXTSTSPowerState               int         `json:"OS-EXT-STS:power_state"`
	OsExtendedVolumesVolumesAttached []struct {
		ID string `json:"id"`
	} `json:"os-extended-volumes:volumes_attached"`
	SecurityGroups []struct {
		Name string `json:"name"`
	} `json:"security_groups"`
}

type Servers struct {
	Links   interface{}   `json:"links"`
	Servers []*ServerInfo `json:"servers"`
}
