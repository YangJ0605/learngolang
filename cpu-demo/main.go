package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Println("get cpu info failed, err: ", err)
	}

	for _, ci := range cpuInfos {
		fmt.Println("cpu", ci)
	}

	// for {
	// 	percent, _ := cpu.Percent(time.Second, false)
	// 	fmt.Println("cpu百分比: ", percent)
	// }

	info, _ := load.Avg()
	fmt.Println("info: ", info)

	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("memInfo: %v\n", memInfo)

	hInfo, _ := host.Info()
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
}
