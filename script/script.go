package script

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"github.com/shirou/gopsutil/v3/process"
)

func ProcessFilter(processname string) ([]*process.Process, error) {
	var ret []*process.Process
	processes, _ := process.Processes()
	for i, proc := range processes {
		name, _ := proc.Name()
		if strings.Contains( name, processname) {
			ret = append(ret, proc)
		}
	}
	return ret, nil
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
	commande := "echo $(ps aux | grep \"" + soft + "\"| grep -v grep | awk 'BEGIN { sum=0 } {sum=sum+$3; } END {printf(\"%s\",sum)}')"
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