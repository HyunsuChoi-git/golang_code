package libvirt

var C libvirtConst

var MetricNames map[string]bool

type libvirtConst struct {
	METRIC_TYPE metricType
}

type metricType struct {
	INFO_STATE                             string
	INFO_VSTATE                            string
	INFO_META                              string
	INFO_VIRTUAL_CPUS                      string
	INFO_CPU_TIME_SECONDS_TOTAL            string
	INFO_MAXIMUM_MEMORY_BYTES              string
	BLOCK_STATS_READ_BYTES_TOTAL           string
	BLOCK_STATS_READ_REQUESTS_TOTAL        string
	BLOCK_STATS_WRITE_BYTES_TOTAL          string
	BLOCK_STATS_WRITE_REQUESTS_TOTAL       string
	BLOCK_STATS_CAPACITY_BYTES             string
	BLOCK_CAPACITY_BYTES                   string
	BLOCK_STATS_ALLOCATION                 string
	MEMORY_STATS_USABLE_BYTES              string
	MEMORY_USABLE_BYTES                    string
	MEMORY_STATS_AVAILABLE_BYTES           string
	MEMORY_STATS_USED_PERCENT              string
	INTERFACE_STATS_RECEIVE_BYTES_TOTAL    string
	INTERFACE_STATS_RECEIVE_DROPS_TOTAL    string
	INTERFACE_STATS_RECEIVE_ERRORS_TOTAL   string
	INTERFACE_STATS_RECEIVE_PACKETS_TOTAL  string
	INTERFACE_STATS_TRANSMIT_BYTES_TOTAL   string
	INTERFACE_STATS_TRANSMIT_DROPS_TOTAL   string
	INTERFACE_STATS_TRANSMIT_ERRORS_TOTAL  string
	INTERFACE_STATS_TRANSMIT_PACKETS_TOTAL string

	//신정원 추가
	STATE_CODE                         string
	BLOCK_STATS_CAPACITY               string
	BLOCK_STATS_ALLOCATION2            string
	MEMORY_STAT_MEMORY_USABLE_BYTES    string
	MEMORY_STAT_MEMORY_AVAILABLE_BYTES string
}

func init() {
	// 수집할 메트릭 추가 시 MetricNames, C 두 군데에 모두 추가

	C.METRIC_TYPE.INFO_STATE = "libvirt_domain_info_state" // libvirt_domain_state_code
	C.METRIC_TYPE.STATE_CODE = "libvirt_domain_state_code"
	C.METRIC_TYPE.INFO_VSTATE = "libvirt_domain_info_vstate"
	C.METRIC_TYPE.INFO_META = "libvirt_domain_info_meta" // x
	C.METRIC_TYPE.INFO_VIRTUAL_CPUS = "libvirt_domain_info_virtual_cpus"
	C.METRIC_TYPE.INFO_CPU_TIME_SECONDS_TOTAL = "libvirt_domain_info_cpu_time_seconds_total"
	C.METRIC_TYPE.INFO_MAXIMUM_MEMORY_BYTES = "libvirt_domain_info_maximum_memory_bytes"
	C.METRIC_TYPE.BLOCK_STATS_READ_BYTES_TOTAL = "libvirt_domain_block_stats_read_bytes_total"
	C.METRIC_TYPE.BLOCK_STATS_READ_REQUESTS_TOTAL = "libvirt_domain_block_stats_read_requests_total"
	C.METRIC_TYPE.BLOCK_STATS_WRITE_BYTES_TOTAL = "libvirt_domain_block_stats_write_bytes_total"
	C.METRIC_TYPE.BLOCK_STATS_WRITE_REQUESTS_TOTAL = "libvirt_domain_block_stats_write_requests_total"
	C.METRIC_TYPE.BLOCK_STATS_CAPACITY_BYTES = "libvirt_domain_block_stats_capacity_bytes"
	C.METRIC_TYPE.BLOCK_CAPACITY_BYTES = "libvirt_domain_block_capacity_bytes"
	C.METRIC_TYPE.BLOCK_STATS_CAPACITY = "libvirt_domain_block_stats_Capacity" //
	C.METRIC_TYPE.BLOCK_STATS_ALLOCATION = "libvirt_domain_block_stats_allocation"
	C.METRIC_TYPE.BLOCK_STATS_ALLOCATION2 = "libvirt_domain_block_stats_Allocation" //
	C.METRIC_TYPE.MEMORY_STATS_USABLE_BYTES = "libvirt_domain_memory_stats_usable_bytes"
	C.METRIC_TYPE.MEMORY_STAT_MEMORY_USABLE_BYTES = "libvirt_domain_stat_memory_usable_bytes" //
	C.METRIC_TYPE.MEMORY_USABLE_BYTES = "libvirt_domain_memory_usable_bytes"
	C.METRIC_TYPE.MEMORY_STATS_AVAILABLE_BYTES = "libvirt_domain_memory_stats_available_bytes"
	C.METRIC_TYPE.MEMORY_STAT_MEMORY_AVAILABLE_BYTES = "libvirt_domain_stat_memory_available_bytes" //
	C.METRIC_TYPE.MEMORY_STATS_USED_PERCENT = "libvirt_domain_memory_stats_used_percent"
	C.METRIC_TYPE.INTERFACE_STATS_RECEIVE_BYTES_TOTAL = "libvirt_domain_interface_stats_receive_bytes_total"
	C.METRIC_TYPE.INTERFACE_STATS_RECEIVE_DROPS_TOTAL = "libvirt_domain_interface_stats_receive_drops_total"
	C.METRIC_TYPE.INTERFACE_STATS_RECEIVE_ERRORS_TOTAL = "libvirt_domain_interface_stats_receive_errors_total"
	C.METRIC_TYPE.INTERFACE_STATS_RECEIVE_PACKETS_TOTAL = "libvirt_domain_interface_stats_receive_packets_total"
	C.METRIC_TYPE.INTERFACE_STATS_TRANSMIT_BYTES_TOTAL = "libvirt_domain_interface_stats_transmit_bytes_total"
	C.METRIC_TYPE.INTERFACE_STATS_TRANSMIT_DROPS_TOTAL = "libvirt_domain_interface_stats_transmit_drops_total"
	C.METRIC_TYPE.INTERFACE_STATS_TRANSMIT_ERRORS_TOTAL = "libvirt_domain_interface_stats_transmit_errors_total"
	C.METRIC_TYPE.INTERFACE_STATS_TRANSMIT_PACKETS_TOTAL = "libvirt_domain_interface_stats_transmit_packets_total"

	MetricNames = map[string]bool{
		C.METRIC_TYPE.INFO_STATE:                             true,
		C.METRIC_TYPE.INFO_VSTATE:                            true,
		C.METRIC_TYPE.INFO_META:                              true,
		C.METRIC_TYPE.INFO_VIRTUAL_CPUS:                      true,
		C.METRIC_TYPE.INFO_MAXIMUM_MEMORY_BYTES:              true,
		C.METRIC_TYPE.INFO_CPU_TIME_SECONDS_TOTAL:            true,
		C.METRIC_TYPE.BLOCK_STATS_READ_BYTES_TOTAL:           true,
		C.METRIC_TYPE.BLOCK_STATS_READ_REQUESTS_TOTAL:        true,
		C.METRIC_TYPE.BLOCK_STATS_WRITE_BYTES_TOTAL:          true,
		C.METRIC_TYPE.BLOCK_STATS_WRITE_REQUESTS_TOTAL:       true,
		C.METRIC_TYPE.BLOCK_STATS_CAPACITY_BYTES:             true,
		C.METRIC_TYPE.BLOCK_STATS_ALLOCATION:                 true,
		C.METRIC_TYPE.BLOCK_CAPACITY_BYTES:                   true,
		C.METRIC_TYPE.MEMORY_STATS_USABLE_BYTES:              true,
		C.METRIC_TYPE.MEMORY_USABLE_BYTES:                    true,
		C.METRIC_TYPE.MEMORY_STATS_AVAILABLE_BYTES:           true,
		C.METRIC_TYPE.MEMORY_STATS_USED_PERCENT:              true,
		C.METRIC_TYPE.INTERFACE_STATS_RECEIVE_BYTES_TOTAL:    true,
		C.METRIC_TYPE.INTERFACE_STATS_RECEIVE_DROPS_TOTAL:    true,
		C.METRIC_TYPE.INTERFACE_STATS_RECEIVE_ERRORS_TOTAL:   true,
		C.METRIC_TYPE.INTERFACE_STATS_RECEIVE_PACKETS_TOTAL:  true,
		C.METRIC_TYPE.INTERFACE_STATS_TRANSMIT_BYTES_TOTAL:   true,
		C.METRIC_TYPE.INTERFACE_STATS_TRANSMIT_DROPS_TOTAL:   true,
		C.METRIC_TYPE.INTERFACE_STATS_TRANSMIT_ERRORS_TOTAL:  true,
		C.METRIC_TYPE.INTERFACE_STATS_TRANSMIT_PACKETS_TOTAL: true,

		// 신정원 추가
		C.METRIC_TYPE.STATE_CODE:                         true,
		C.METRIC_TYPE.BLOCK_STATS_CAPACITY:               true,
		C.METRIC_TYPE.BLOCK_STATS_ALLOCATION2:            true,
		C.METRIC_TYPE.MEMORY_STAT_MEMORY_USABLE_BYTES:    true,
		C.METRIC_TYPE.MEMORY_STAT_MEMORY_AVAILABLE_BYTES: true,
	}

}
