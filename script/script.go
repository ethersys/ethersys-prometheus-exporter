package script

import (
	"strings"
	"github.com/shirou/gopsutil/v3/process"
)

func ProcessFilter(processname string) ([]*process.Process, error) {
	var ret []*process.Process
	processes, _ := process.Processes()
	for _, proc := range processes {
		cmd, _ := proc.Cmdline()
		if strings.Contains( cmd, processname) {
			ret = append(ret, proc)
		}
	}
	return ret, nil
}

func UsedMemory(soft string) float64 {
	var ret float64 = 0
	var val float64
	var mem *process.MemoryInfoStat
	processes, _ := process.Processes()
	for _, proc := range processes {
		mem, _ = proc.MemoryInfo()
		val = float64(mem.RSS)
		ret += val / 1024
	}
	return ret
}

func UsedCPU(soft string) float64 {
	var ret float64 = 0
	var val float64
	processes, _ := ProcessFilter(soft)
	for _, proc := range processes {
		val, _ = proc.CPUPercent()
		ret += val
	}
	return ret
}