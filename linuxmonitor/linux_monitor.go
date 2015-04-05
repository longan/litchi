package linuxmonitor

import (
	"github.com/c9s/goprocinfo/linux"
)

func GetCpu() (*linux.Stat, error) {
	return linux.ReadStat("/proc/stat")
}
