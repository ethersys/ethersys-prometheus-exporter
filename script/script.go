package script

import (
	"strings"
	"log"
	"os/exec"
	"strconv"
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

func UsedMemoryPercent(soft string) float64 {
	var ret float64 = 0
	var val float32
	processes, _ := ProcessFilter(soft)
	for _, proc := range processes {
		val, _ = proc.MemoryPercent()
		ret += float64(val)
	}
	return ret
}

func UsedMemory(soft string) float64 {
	commande := "echo $(ps aux | grep \"" + soft + "\"| grep -v grep | awk 'BEGIN { sum=0 } {sum=sum+$6; } END {printf(\"%s\",sum / 1024)}')"
	cmd, err := exec.Command("bash", "-c", commande).Output()
	if err != nil {
		log.Fatal(err)
	}
	res, err := strconv.ParseFloat(strings.Replace(string(cmd), "\n", "", 1), 64)
	if err != nil {
		log.Fatal(err)
	}
	return res
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