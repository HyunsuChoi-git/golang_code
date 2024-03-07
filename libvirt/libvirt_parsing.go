package libvirt

import (
	"fmt"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"strings"
)

func LibvirtParsingTest() {
	var err error

	allVmMetrics := map[string]MetricDatas{}
	response := ""
	{
		fname := "libvirt/libvirt_response.txt"

		// 파일 읽기
		content, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Println("파일을 읽는 동안 오류가 발생했습니다:", err)
			return
		}

		// 파일 내용을 문자열 변수에 할당
		response = string(content)
	}

	hypervisorMetrics := HypervisorMetrics{
		HostIp: "10.111.11.21",
	}

	parser := expfmt.TextParser{}
	originMetricsMap, err := parser.TextToMetricFamilies(strings.NewReader(response))
	if err != nil {
		fmt.Println("Error parsing input:", err)
	}

	if len(originMetricsMap) == 0 {
		fmt.Println("Libvirt Parsing error: ", response)
	}

	/* 3-4) vm filtering */
	var metaInfo *io_prometheus_client.MetricFamily
	if _, ok := originMetricsMap[C.METRIC_TYPE.INFO_META]; ok {
		metaInfo, _ = originMetricsMap[C.METRIC_TYPE.INFO_META]
	} else if _, ok := originMetricsMap[C.METRIC_TYPE.STATE_CODE]; ok {
		metaInfo, _ = originMetricsMap[C.METRIC_TYPE.STATE_CODE]
	}

	if metaInfo == nil {

		metricName := ""
		for key, _ := range originMetricsMap {
			metricName = metricName + " / " + key
		}
		metricName = strings.Trim(metricName, " / ")

		log.Error().Msgf("No 'libvirt_domain_info_meta' or 'libvirt_domain_state_code'.  metric cnt(%d), metric name(%s)", len(originMetricsMap), metricName)
	}

	for _, k := range metaInfo.Metric {
		var vmName string
		var projectUuid string
		var uuid string
		for _, l := range k.Label {
			if *l.Name == "domain" {
				vmName = *l.Value
			} else if *l.Name == "project_uuid" {
				projectUuid = *l.Value
			} else if *l.Name == "projectId" {
				projectUuid = *l.Value
			} else if *l.Name == "uuid" {
				uuid = *l.Value
			} else if *l.Name == "instanceId" {
				uuid = *l.Value
			}
		}

		// 수집대상 프로젝트 내의 vm이면 저장
		//if _, ok := osProjectList[projectUuid]; ok {
		//	vm := MetricDatas{}
		//	vm.VmDomain = osProjectList[projectUuid]
		//	vm.VmName = vmName
		//	vm.VmUuId = uuid
		//	vm.HypervisorMetrics = hypervisorMetrics
		//	allVmMetrics[vmName] = vm
		//}

		vm := MetricDatas{}
		vm.VmDomain = projectUuid
		vm.VmName = vmName
		vm.VmUuId = uuid
		vm.HypervisorMetrics = hypervisorMetrics
		allVmMetrics[vmName] = vm

	}

	/* 3-5) vm별 필요 메트릭 추출 */
	ExtractMetrics(originMetricsMap, &allVmMetrics)

	convert(allVmMetrics)
}

func convert(allVmMetrics map[string]MetricDatas) {
	// 2. 음답 값 convert
	//		vm 별로 메트릭 정리. 한 vm당 전체/diskio/network/filestorage 4가지 메트릭 존재
	for _, vmMetrics := range allVmMetrics {

		basic := Basic{}
		basic.HypervisorIp = vmMetrics.HypervisorMetrics.HostIp

		// Metric.Vm
		basic.Vm.VmId = vmMetrics.VmUuId
		basic.Vm.Domain = vmMetrics.VmDomain

		// Hypervisor

		basic.Vm.State = vmMetrics.Info_state.Value
		if vmMetrics.Info_state.Value == 1 {
			basic.Vm.PowerState = "on"
		} else {
			basic.Vm.PowerState = "off"
		}

		//Metric.CPU
		basic.System.Cpu.Total.Norm.Pct = 0
		basic.System.Cpu.Total.Norm.Total = vmMetrics.Info_cpu_time_seconds_total.Value
		basic.System.Cpu.Cores = vmMetrics.Info_virtual_cpus.Value

		//Metric.Memory
		basic.System.Memory.Actual.Used.Bytes =
			vmMetrics.Memory_stats_available_bytes.Value - vmMetrics.Memory_stats_usable_bytes.Value
		if basic.System.Memory.Actual.Used.Bytes == 0 {
			basic.System.Memory.Actual.Used.Pct = 0
		} else {
			basic.System.Memory.Actual.Used.Pct =
				(vmMetrics.Memory_stats_available_bytes.Value - vmMetrics.Memory_stats_usable_bytes.Value) / vmMetrics.Memory_stats_available_bytes.Value * 100
		}
		basic.System.Memory.Actual.Free = vmMetrics.Memory_stats_usable_bytes.Value
		basic.System.Memory.Free = vmMetrics.Memory_stats_usable_bytes.Value
		basic.System.Memory.Total = vmMetrics.Info_maximum_memory_bytes.Value

		//Metric.Filesystem
		basic.System.Filesystem = Filesystem{}

		capaBytesList := vmMetrics.Block_stats_capacity_bytes
		var capaBytes float64
		for _, v := range capaBytesList {
			capaBytes += v.Value
		}
		basic.System.Filesystem.All.Total = capaBytes

		allocationList := vmMetrics.Block_stats_allocation
		var allocationTotal float64
		for _, v := range allocationList {
			allocationTotal += v.Value
		}
		basic.System.Filesystem.All.Available = capaBytes - allocationTotal

		//Metric.Diskio
		readBytesList := vmMetrics.Block_stats_read_bytes_total
		var readBytes float64
		for _, v := range readBytesList {
			readBytes += v.Value
		}
		basic.System.Diskio.All.Read.Bytes = readBytes

		readRequestList := vmMetrics.Block_stats_read_requests_total
		var readRequest float64
		for _, v := range readRequestList {
			readRequest += v.Value
		}
		basic.System.Diskio.All.Read.Count = readRequest

		writeBytesList := vmMetrics.Block_stats_write_bytes_total
		var writeTotal float64
		for _, v := range writeBytesList {
			writeTotal += v.Value
		}
		basic.System.Diskio.All.Write.Bytes = writeTotal

		writeRequestList := vmMetrics.Block_stats_write_requests_total
		var writeRequest float64
		for _, v := range writeRequestList {
			writeRequest += v.Value
		}
		basic.System.Diskio.All.Write.Count = writeRequest

		//Metric.Network
		receiveBytesList := vmMetrics.Interface_stats_receive_bytes_total
		var receiveBytes float64
		for _, v := range receiveBytesList {
			receiveBytes += v.Value
		}
		basic.System.Network.All.In.Bytes = receiveBytes

		receivePacketsList := vmMetrics.Interface_stats_receive_packets_total
		var receivePackets float64
		for _, v := range receivePacketsList {
			receivePackets += v.Value
		}
		basic.System.Network.All.In.Packets = receivePackets

		receiveDropsList := vmMetrics.Interface_stats_receive_drops_total
		var receiveDrops float64
		for _, v := range receiveDropsList {
			receiveDrops += v.Value
		}
		basic.System.Network.All.In.Dropped = receiveDrops

		receiveErrorList := vmMetrics.Interface_stats_receive_errors_total
		var receiveError float64
		for _, v := range receiveErrorList {
			receiveError += v.Value
		}
		basic.System.Network.All.In.Errors = receiveError

		transmitBytesList := vmMetrics.Interface_stats_transmit_bytes_total
		var transmitBytes float64
		for _, v := range transmitBytesList {
			transmitBytes += v.Value
		}
		basic.System.Network.All.Out.Bytes = transmitBytes

		transmitPacketsList := vmMetrics.Interface_stats_transmit_packets_total
		var transmitPackets float64
		for _, v := range transmitPacketsList {
			transmitPackets += v.Value
		}
		basic.System.Network.All.Out.Packets = transmitPackets

		transmitDropsList := vmMetrics.Interface_stats_transmit_drops_total
		var transmitDrops float64
		for _, v := range transmitDropsList {
			transmitDrops += v.Value
		}
		basic.System.Network.All.Out.Dropped = transmitDrops

		transmitErrorList := vmMetrics.Interface_stats_transmit_errors_total
		var transmitError float64
		for _, v := range transmitErrorList {
			transmitError += v.Value
		}
		basic.System.Network.All.Out.Errors = transmitError

		//Metric.Diskio.each
		basic.System.Diskio.Each = extractDiskioMetric(readBytesList, readRequestList, writeBytesList, writeRequestList)

		//Metric.Network.each
		basic.System.Network.Each = extractNetworkMetric(receiveBytesList, receivePacketsList, receiveDropsList, receiveErrorList,
			transmitBytesList, transmitPacketsList, transmitDropsList, transmitErrorList)

		//Metric.Filesystem
		basic.System.Filesystem.Each = extractFilesystemMetric(allocationList, capaBytesList)

		//Metric List에 추가
		//vmOutputMetricList = append(vmOutputMetricList, metricData)

		//바로 전송

		log.Info().Msgf(" Basic >> %v", basic)
	}
}

func extractDiskioMetric(readBytesList []TargetMetric, readRequestList []TargetMetric,
	writeBytesList []TargetMetric, writeRequestList []TargetMetric) (diskioDetailList []DiskioDetail) {

	// ( read bytes의 인덱스 값을 기준으로 read bytes, read nytes, write count도 구해서 하나의 메트릭으로 만든다.
	for _, readBytes := range readBytesList {
		diskioDetail := DiskioDetail{}

		// 1. Metric 추출
		// 1) name 추출
		name := readBytes.TargetDevice
		diskioDetail.Name = name

		// 2) read byte 추출
		diskioDetail.Read.Bytes = readBytes.Value

		// 3) read count 추출
		for _, readRequest := range readRequestList {
			if readRequest.TargetDevice == name {
				diskioDetail.Read.Count = readRequest.Value
				break
			}
		}
		// 4) write byte 추출
		for _, writeBytes := range writeBytesList {
			if writeBytes.TargetDevice == name {
				diskioDetail.Write.Bytes = writeBytes.Value
				break
			}
		}
		// 5) write count 추출
		for _, writeRequest := range writeRequestList {
			if writeRequest.TargetDevice == name {
				diskioDetail.Write.Count = writeRequest.Value
				break
			}
		}

		diskioDetailList = append(diskioDetailList, diskioDetail)
	}

	return
}
func extractNetworkMetric(receiveBytesList []TargetMetric, receivePacketsList []TargetMetric,
	receiveDropsList []TargetMetric, receiveErrorList []TargetMetric, transmitBytesList []TargetMetric,
	transmitPacketsList []TargetMetric, transmitDropsList []TargetMetric, transmitErrorList []TargetMetric) (networkDetailList []NetworkDetail) {

	// readBytes의 'deviceName'값을 기준으로 나머지 metric 들도 구해서 하나의 메트릭으로 만든다.
	for _, receiveBytes := range receiveBytesList {

		var networkDetail NetworkDetail
		name := receiveBytes.TargetDevice
		networkDetail.Name = name
		networkDetail.In.Bytes = receiveBytes.Value

		for _, receivePackets := range receivePacketsList {
			if receivePackets.TargetDevice == name {
				networkDetail.In.Packets = receivePackets.Value
				break
			}
		}
		for _, receiveDrops := range receiveDropsList {
			if receiveDrops.TargetDevice == name {
				networkDetail.In.Dropped = receiveDrops.Value
				break
			}
		}
		for _, receiveErrors := range receiveErrorList {
			if receiveErrors.TargetDevice == name {
				networkDetail.In.Errors = receiveErrors.Value
				break
			}
		}
		for _, transmitBytes := range transmitBytesList {
			if transmitBytes.TargetDevice == name {
				networkDetail.Out.Bytes = transmitBytes.Value
				break
			}
		}
		for _, transmitPackets := range transmitPacketsList {
			if transmitPackets.TargetDevice == name {
				networkDetail.Out.Packets = transmitPackets.Value
				break
			}
		}
		for _, transmitDrops := range transmitDropsList {
			if transmitDrops.TargetDevice == name {
				networkDetail.Out.Dropped = transmitDrops.Value
				break
			}
		}
		for _, transmitErrors := range transmitErrorList {
			if transmitErrors.TargetDevice == name {
				networkDetail.Out.Errors = transmitErrors.Value
				break
			}
		}

		networkDetailList = append(networkDetailList, networkDetail)
	}

	return
}

func extractFilesystemMetric(allocationList []TargetMetric, capaBytesList []TargetMetric) (filesystemDetailList []FilesystemDetail) {

	// ( listWriteBytesTotal의 'deviceName'값을 기준으로 available, total 값을 구해 하나의 메트릭을 만든다.)
	for _, allocationBytes := range allocationList {
		var filesystemDetail FilesystemDetail

		name := allocationBytes.TargetDevice
		filesystemDetail.Name = name

		for _, capaBytes := range capaBytesList {
			if capaBytes.TargetDevice == name {
				filesystemDetail.Total = capaBytes.Value
				filesystemDetail.Available = capaBytes.Value - allocationBytes.Value // total - 사용량 = 사용가능량
				break
			}
		}

		filesystemDetailList = append(filesystemDetailList, filesystemDetail)
	}

	return
}
