package script

import (
	"strings"
	"github.com/shirou/gopsutil/v3/process"
)

func ListContains(terms []string, searchTerm string) bool {
    for _, term := range terms {
        if strings.Contains( searchTerm, term) {
            return true
        }
    }
    return false
}

func ProcessesExcludeFilter(processesList []string) ([]*process.Process, error) {
	var ret []*process.Process
	processes, _ := process.Processes()
	for _, proc := range processes {
		cmd, _ := proc.Cmdline()
		if (!ListContains(processesList, cmd)) {
			ret = append(ret, proc)
		}
	}
	return ret, nil
}

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
	var err error
	var memoryMapsStats *[]process.MemoryMapsStat
	processes, _ := ProcessFilter(soft)
	for _, proc := range processes {
		memoryMapsStats, err = proc.MemoryMaps(true)
		if err == nil {
			for _, memoryMapsStat := range *memoryMapsStats {
				val = float64(memoryMapsStat.Rss)
				ret += val / 1024
			}
		}
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

func UnmonitoredMemory(processesList []string) float64 {
	var ret float64 = 0
	var val float64
	var err error
	var memoryMapsStats *[]process.MemoryMapsStat
	processes, _ := ProcessesExcludeFilter(processesList)
	for _, proc := range processes {
		memoryMapsStats, err = proc.MemoryMaps(true)
		if err == nil {
			for _, memoryMapsStat := range *memoryMapsStats {
				val = float64(memoryMapsStat.Rss)
				ret += val / 1024
			}
		}
	}
	return ret
}

func UnmonitoredCPU(processesList []string) float64 {
	var ret float64 = 0
	var val float64
	processes, _ := ProcessesExcludeFilter(processesList)
	for _, proc := range processes {
		val, _ = proc.CPUPercent()
		ret += val
	}
	return ret
}