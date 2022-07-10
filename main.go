package main

import (
	"fmt"
	"monitor/global"
	"monitor/initialize"
	"monitor/services"
)

type Stat struct {
	Pro      string             `csv:"pro"`
	IP       string             `csv:"ip"`
	Name     string             `csv:"name"`
	UpTime   float64            `csv:"up_time"`
	TotalMem float64            `csv:"total_mem"`
	UsedMem  float64            `csv:"used_mem"`
	TotalCpu float64            `csv:"total_cpu"`
	UsedCpu  float64            `csv:"used_cpu"`
	Disk     map[string]float64 `csv:"disk"`
}

func main() {

	initialize.InitConfig() // 初始化配置文件
	initialize.InitApis()   // 初始化 prometheus API

	var stats []Stat
	for _, v := range global.PromConfig.Services {

		hostName, _ := services.QueryHostName(v.Name)
		upTime, _ := services.QueryUpTime(v.Name)
		cpuTotal, _ := services.QueryCpuNum(v.Name)
		cpuUse, _ := services.QueryCPUUseAvg(v.Name)
		memoryTotal, _ := services.QueryMemoryTotal(v.Name)
		memoryUse, _ := services.QueryMemoryAvg(v.Name)
		diskStat, _ := services.QueryDiskUse(v.Name)

		for k, hostNames := range hostName {
			fmt.Println(k)
			stat := Stat{
				Pro:      v.Name,
				IP:       k,
				Name:     hostNames,
				UpTime:   upTime[k],
				TotalCpu: cpuTotal[k],
				UsedCpu:  cpuUse[k],
				TotalMem: memoryTotal[k],
				UsedMem:  memoryUse[k],
				Disk:     diskStat[k],
			}
			stats = append(stats, stat)
		}

	}
	for _, v := range stats {
		if v.Pro == "azure" {
			fmt.Println("azure")
		}
		if v.Pro == "test" {
			fmt.Println("test")
		}
	}

}
