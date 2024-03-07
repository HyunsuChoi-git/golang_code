package libvirt

//	type VmOutputMetrics struct {
//		ProvID        string         `json:"prov_id"`
//		CollNo        string         `json:"coll_no"`
//		Agent         config.Agent   `json:"agent"`
//		Service       config.Service `json:"service"`
//		CollectorTime int64          `json:"collector_time"`
//		ApiInfo       config.ApiInfo `json:"api_info"`
//		OutInfo       config.OutInfo `json:"out_info"`
//		Metrics       Basic          `json:"basic"`
//	}
type Basic struct {
	Vm           Vm     `json:"vm"`
	System       System `json:"system"`
	HypervisorIp string `json:"hypervisor_ip"`
}

type Vm struct {
	VmId       string  `json:"vm_id"`
	Domain     string  `json:"domain"`
	PowerState string  `json:"power_state"`
	State      float64 `json:"state"`
}

type System struct {
	Diskio     Diskio     `json:"diskio"`
	Filesystem Filesystem `json:"filesystem"`
	Network    Network    `json:"network"`
	Cpu        Cpu        `json:"cpu"`
	Memory     Memory     `json:"memory"`
}

type Filesystem struct {
	All  AllFilesystemDetail `json:"all"`
	Each []FilesystemDetail  `json:"each"`
}

type AllFilesystemDetail struct {
	Total     float64 `json:"total"`
	Available float64 `json:"available"`
}

type Diskio struct {
	All  AllDiskioDetail `json:"all"`
	Each []DiskioDetail  `json:"each"`
}

type AllDiskioDetail struct {
	Read  ReadWrite `json:"read"`
	Write ReadWrite `json:"write"`
}

type Network struct {
	All  AllNetworkDetail `json:"all"`
	Each []NetworkDetail  `json:"each"`
}

type AllNetworkDetail struct {
	In  InOut `json:"in"`
	Out InOut `json:"out"`
}

type Cpu struct {
	Total Total   `json:"total"`
	Cores float64 `json:"cores"`
}

type Memory struct {
	Actual Actual  `json:"actual"`
	Free   float64 `json:"free"`
	Total  float64 `json:"total"`
}

//

type ReadWrite struct {
	Bytes float64 `json:"bytes"`
	Count float64 `json:"count"`
}

type InOut struct {
	Packets float64 `json:"packets"`
	Bytes   float64 `json:"bytes"`
	Dropped float64 `json:"dropped"`
	Errors  float64 `json:"errors"`
}

type Total struct {
	Norm Norm `json:"norm"`
}

type Norm struct {
	Total float64 `json:"total"`
	Pct   float64 `json:"pct"`
}
type Actual struct {
	Free float64 `json:"free"`
	Used Used    `json:"used"`
}

type Used struct {
	Pct   float64 `json:"pct"`
	Bytes float64 `json:"bytes"`
}

// diskio
type DiskioDetail struct {
	Name  string    `json:"name"`
	Read  ReadWrite `json:"read"`
	Write ReadWrite `json:"write"`
}

// filesystem
type FilesystemDetail struct {
	Name      string  `json:"name"`
	Available float64 `json:"available"`
	Total     float64 `json:"total"`
}

// network
type NetworkDetail struct {
	Name string `json:"name"`
	In   InOut  `json:"in"`
	Out  InOut  `json:"out"`
}
