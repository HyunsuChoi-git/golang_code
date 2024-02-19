package os_api_model

type HypervisorDetails struct {
	Hypervisors []*HypervisorInfo
}

type HypervisorInfo struct {
	ID                 int64  `json:"id"`
	HypervisorHostname string `json:"hypervisor_hostname"`
	State              string `json:"state"`
	Status             string `json:"status"`
	HypervisorType     string `json:"hypervisor_type"`
	HypervisorVersion  float64
	HostIp             string `json:"host_ip"`
	Service            struct {
		ID             int    `json:"id"`
		Host           string `json:"host"`
		DisabledReason string `json:"disabled_reason"`
	} `json:"service"`
	Vcpus              float64 `json:"vcpus"`
	MemoryMb           float64 `json:"memory_mb"`
	LocalGb            float64 `json:"local_gb"`
	VcpusUsed          float64 `json:"vcpus_used"`
	MemoryMbUsed       float64 `json:"memory_mb_used"`
	LocalGbUsed        float64 `json:"local_gb_used"`
	FreeRAMMb          float64 `json:"free_ram_mb"`
	FreeDiskGb         float64 `json:"free_disk_gb"`
	CurrentWorkload    float64 `json:"current_workload"`
	RunningVms         float64 `json:"running_vms"`
	DiskAvailableLeast float64 `json:"disk_available_least"`
	CPUInfo            string  `json:"cpu_info"`
}
