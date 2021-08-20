package script

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

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
