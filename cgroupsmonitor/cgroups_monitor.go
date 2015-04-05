package cgroupsmonitor

import (
	"os/exec"
)

func GetMetricsString() string {
	cmd := exec.Command("top", "-l", "1")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return string(output)
}
