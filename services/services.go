package services

import (
	"fmt"
	"monitor/global"
	"monitor/prom"
	"monitor/utils"
)

// hostName
func QueryHostName(serviceName string) (map[string]string, error) {
	pql := "node_uname_info"
	value, err := prom.QueryVectorValue(serviceName, pql)
	if err != nil {
		return nil, err
	}
	hostName, err := utils.ParseLabelValue(value, "nodename")
	if err != nil {
		return nil, err
	}
	return hostName, nil

}

// 在线时间
func QueryUpTime(serviceName string) (map[string]float64, error) {
	qpl := "sum(time() - node_boot_time_seconds)by(instance)"
	value, err := prom.QueryVectorValue(serviceName, qpl)
	if err != nil {
		return nil, err
	}
	upTime, err := utils.ParseVectorValue(value)
	if err != nil {
		fmt.Println("ParseVectorValue error : ", err)
	}
	return upTime, nil
}

// totalCpu
func QueryCpuNum(serviceName string) (map[string]float64, error) {
	pql := "count(node_cpu_seconds_total) by (instance)"
	value, err := prom.QueryVectorValue(serviceName, pql)
	if err != nil {
		fmt.Println("QueryCpuNum error : ", err)
	}
	cpuNumStat, err := utils.ParseVectorValue(value)
	if err != nil {
		fmt.Println("ParseVectorValue error : ", err)
	}
	return cpuNumStat, err
}

// totalMem
func QueryMemoryTotal(serviceName string) (map[string]float64, error) {
	pql := "node_memory_MemTotal_bytes"
	value, err := prom.QueryVectorValue(serviceName, pql)
	if err != nil {
		fmt.Println("QueryMemory error : ", err)
	}
	memoryStat, err := utils.ParseVectorValue(value)
	if err != nil {
		fmt.Println("ParseVectorValue error : ", err)
	}
	return memoryStat, err
}

// avgCpu
func QueryCPUUseAvg(serviceName string) (map[string]float64, error) {
	pql := fmt.Sprintf("(1 - avg(rate(node_cpu_seconds_total{mode='idle'}[%s])) by (instance)) * 100", global.PromConfig.Day)
	value, err := prom.QueryVectorValue(serviceName, pql)
	if err != nil {
		fmt.Println("QueryCPU error : ", err)
	}
	cpuStat, err := utils.ParseVectorValue(value)
	if err != nil {
		fmt.Println("ParseVectorValue error : ", err)
	}
	return cpuStat, err
}

// avgMem
func QueryMemoryAvg(serviceName string) (map[string]float64, error) {
	pql := fmt.Sprintf("100 * (1 - ((avg_over_time(node_memory_MemFree_bytes[%s]) + avg_over_time(node_memory_Cached_bytes[%s]) + avg_over_time(node_memory_Buffers_bytes[%s])) / avg_over_time(node_memory_MemTotal_bytes[24h])))", global.PromConfig.Day, global.PromConfig.Day, global.PromConfig.Day)
	value, err := prom.QueryVectorValue(serviceName, pql)
	if err != nil {
		fmt.Println("QueryMemory error : ", err)
	}
	memoryStat, err := utils.ParseVectorValue(value)
	if err != nil {
		fmt.Println("ParseVectorValue error : ", err)
	}
	return memoryStat, err
}

// disk
func QueryDiskUse(serviceName string) (map[string]map[string]float64, error) {
	qpl := "max((node_filesystem_size_bytes{fstype=~'ext.?|xfs',mountpoint=~'/data.*|/web|/'}-node_filesystem_free_bytes{fstype=~'ext.?|xfs',mountpoint=~'/data.*|/web|/'}) *100/(node_filesystem_avail_bytes {fstype=~'ext.?|xfs'}+(node_filesystem_size_bytes{fstype=~'ext.?|xfs',mountpoint=~'/data.*|/web|/'}-node_filesystem_free_bytes{fstype=~'ext.?|xfs',mountpoint=~'/data.*|/web|/'})))by(instance,mountpoint)"
	value, err := prom.QueryVectorValue(serviceName, qpl)
	if err != nil {
		fmt.Println("QueryDisk error : ", err)
	}
	diskStat, err := utils.ParseMultiVectorValue(value, "mountpoint")
	if err != nil {
		fmt.Println("ParseVectorValue error : ", err)
	}
	return diskStat, nil
}
