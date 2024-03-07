package libvirt

type MetricDatas struct {
	VmName                      string
	VmUuId                      string
	VmDomain                    string
	HypervisorMetrics           HypervisorMetrics
	Info_state                  CommonMetric
	Info_virtual_cpus           CommonMetric
	Info_cpu_time_seconds_total CommonMetric
	Info_maximum_memory_bytes   CommonMetric

	Block_stats_read_bytes_total     []TargetMetric
	Block_stats_read_requests_total  []TargetMetric
	Block_stats_write_bytes_total    []TargetMetric
	Block_stats_write_requests_total []TargetMetric
	Block_stats_allocation           []TargetMetric
	Block_stats_capacity_bytes       []TargetMetric

	Memory_stats_usable_bytes    CommonMetric
	Memory_stats_used_percent    CommonMetric
	Memory_stats_available_bytes CommonMetric

	Interface_stats_receive_bytes_total    []TargetMetric
	Interface_stats_receive_drops_total    []TargetMetric
	Interface_stats_receive_packets_total  []TargetMetric
	Interface_stats_receive_errors_total   []TargetMetric
	Interface_stats_transmit_bytes_total   []TargetMetric
	Interface_stats_transmit_drops_total   []TargetMetric
	Interface_stats_transmit_packets_total []TargetMetric
	Interface_stats_transmit_errors_total  []TargetMetric
}

type CommonMetric struct {
	Value float64
}
type TargetMetric struct {
	TargetDevice string
	Value        float64
}

type HypervisorMetrics struct {
	HostIp string
	//Vcpu   HyperMetrics
	//Memory HyperMetrics
	//Disk   HyperMetrics
}

type HyperMetrics struct {
	Total float64 `json:"total"`
	Used  float64 `json:"used"`
	Usage float64 `json:"usage"`
}
