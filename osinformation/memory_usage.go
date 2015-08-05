package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.CPUInfo()
	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
	for _, v := range c {
		fmt.Printf("Percentage: %f\n", v.Mhz)
	}
	// convert to JSON. String() is also implemented
}
