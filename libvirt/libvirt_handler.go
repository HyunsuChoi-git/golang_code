package libvirt

import (
	"encoding/json"
	"github.com/prometheus/client_model/go"
	"github.com/rs/zerolog/log"
)

var (
	ServiceType = "libvirt"
)

func ExtractMetrics(originMetricsMap map[string]*io_prometheus_client.MetricFamily, allVmMetricsP *map[string]MetricDatas) {
	//fnc := "ExtractMetrics"

	for key, om := range originMetricsMap {
		// key 값으로 내가 원하는 메트릭인지 체크
		if _, ok := MetricNames[key]; !ok {
			continue
		}

		// vm별로 메트릭 추출
		// name{lavel}Gauge.value
		for _, metric := range om.Metric {
			var metricData MetricDatas
			var targetDevice string
			var vmName string
			//var vmUuid string

			for _, lavel := range metric.Label { // { 괄호 안 데이터들 }
				l := *lavel.Name
				switch l {
				case "domain":
					vmName = *lavel.Value
				case "target_device":
					targetDevice = *lavel.Value
					//case "uuid":
					//	vmUuid = *lavel.Value
				}
			}

			/* target vm check if it is tartget. */
			if _, ok := (*allVmMetricsP)[vmName]; !ok { // target vm 아니면 continue
				continue
			}

			// Metric Type 별로 값 추출해서 변수에 저장
			var metricValue float64
			if om.Type.String() == "COUNTER" {
				//counter
				metricValue = *metric.Counter.Value
			} else if om.Type.String() == "GAUGE" {
				//gauge
				metricValue = *metric.Gauge.Value
			}

			/* extract vm metric */
			metricData = (*allVmMetricsP)[vmName]

			switch key {
			case C.METRIC_TYPE.INFO_STATE:
				metricData.Info_state = CommonMetric{metricValue}
			case C.METRIC_TYPE.INFO_VSTATE:
				metricData.Info_state = CommonMetric{metricValue}
			case C.METRIC_TYPE.STATE_CODE:
				metricData.Info_state = CommonMetric{metricValue}
			case C.METRIC_TYPE.INFO_META:
				continue
			case C.METRIC_TYPE.INFO_VIRTUAL_CPUS:
				metricData.Info_virtual_cpus = CommonMetric{metricValue}
			case C.METRIC_TYPE.INFO_CPU_TIME_SECONDS_TOTAL:
				metricData.Info_cpu_time_seconds_total = CommonMetric{metricValue}
			case C.METRIC_TYPE.INFO_MAXIMUM_MEMORY_BYTES:
				metricData.Info_maximum_memory_bytes = CommonMetric{metricValue}

			case C.METRIC_TYPE.BLOCK_STATS_READ_BYTES_TOTAL:
				metricData.Block_stats_read_bytes_total =
					append(metricData.Block_stats_read_bytes_total,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.BLOCK_STATS_READ_REQUESTS_TOTAL:
				metricData.Block_stats_read_requests_total =
					append(metricData.Block_stats_read_requests_total,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.BLOCK_STATS_WRITE_BYTES_TOTAL:
				metricData.Block_stats_write_bytes_total =
					append(metricData.Block_stats_write_bytes_total,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.BLOCK_STATS_WRITE_REQUESTS_TOTAL:
				metricData.Block_stats_write_requests_total =
					append(metricData.Block_stats_write_requests_total,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.BLOCK_STATS_CAPACITY_BYTES:
				metricData.Block_stats_capacity_bytes =
					append(metricData.Block_stats_capacity_bytes,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.BLOCK_CAPACITY_BYTES:
				metricData.Block_stats_capacity_bytes =
					append(metricData.Block_stats_capacity_bytes,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.BLOCK_STATS_CAPACITY: // 신정원 추가
				metricData.Block_stats_capacity_bytes =
					append(metricData.Block_stats_capacity_bytes,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.BLOCK_STATS_ALLOCATION:
				metricData.Block_stats_allocation =
					append(metricData.Block_stats_allocation,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.BLOCK_STATS_ALLOCATION2:
				metricData.Block_stats_allocation =
					append(metricData.Block_stats_allocation,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.MEMORY_STATS_USABLE_BYTES:
				metricData.Memory_stats_usable_bytes = CommonMetric{metricValue}
			case C.METRIC_TYPE.MEMORY_STAT_MEMORY_USABLE_BYTES:
				metricData.Memory_stats_usable_bytes = CommonMetric{metricValue}
			case C.METRIC_TYPE.MEMORY_USABLE_BYTES:
				metricData.Memory_stats_usable_bytes = CommonMetric{metricValue}
			case C.METRIC_TYPE.MEMORY_STATS_AVAILABLE_BYTES:
				metricData.Memory_stats_available_bytes = CommonMetric{metricValue}
			case C.METRIC_TYPE.MEMORY_STAT_MEMORY_AVAILABLE_BYTES:
				metricData.Memory_stats_available_bytes = CommonMetric{metricValue}
			case C.METRIC_TYPE.MEMORY_STATS_USED_PERCENT:
				metricData.Memory_stats_used_percent = CommonMetric{metricValue}

			case C.METRIC_TYPE.INTERFACE_STATS_RECEIVE_BYTES_TOTAL:
				metricData.Interface_stats_receive_bytes_total =
					append(metricData.Interface_stats_receive_bytes_total,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.INTERFACE_STATS_RECEIVE_DROPS_TOTAL:
				metricData.Interface_stats_receive_drops_total =
					append(metricData.Interface_stats_receive_drops_total,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.INTERFACE_STATS_RECEIVE_ERRORS_TOTAL:
				metricData.Interface_stats_receive_errors_total =
					append(metricData.Interface_stats_receive_errors_total,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.INTERFACE_STATS_RECEIVE_PACKETS_TOTAL:
				metricData.Interface_stats_receive_packets_total =
					append(metricData.Interface_stats_receive_packets_total,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.INTERFACE_STATS_TRANSMIT_BYTES_TOTAL:
				metricData.Interface_stats_transmit_bytes_total =
					append(metricData.Interface_stats_transmit_bytes_total,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.INTERFACE_STATS_TRANSMIT_DROPS_TOTAL:
				metricData.Interface_stats_transmit_drops_total =
					append(metricData.Interface_stats_transmit_drops_total,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.INTERFACE_STATS_TRANSMIT_ERRORS_TOTAL:
				metricData.Interface_stats_transmit_errors_total =
					append(metricData.Interface_stats_transmit_errors_total,
						TargetMetric{targetDevice, metricValue})
			case C.METRIC_TYPE.INTERFACE_STATS_TRANSMIT_PACKETS_TOTAL:
				metricData.Interface_stats_transmit_packets_total =
					append(metricData.Interface_stats_transmit_packets_total,
						TargetMetric{targetDevice, metricValue})
			default:
				continue
			}

			(*allVmMetricsP)[vmName] = metricData
		}

	}

}

func unmarshal(dataByte []byte, mappingStruct interface{}) {
	fnc := "unmarshal"
	err := json.Unmarshal(dataByte, &mappingStruct)
	if err != nil {
		log.Error().Msgf("%s : Failed to retrieve domain name. - error : %s", fnc, err)
	}
}
