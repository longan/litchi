package agent

import ()

type Resources struct {
	CpuTotal    uint64 `json:cpu_total`
	CpuUsage    uint64 `json:cpu_usage`
	MemoryTotal uint64 `json:memory_total`
	MemoryUsage uint64 `json:memory_usage`
	DiskTotal   uint64 `json:disk_total`
	DiskUsage   uint64 `json:disk_usage`
}

func NewResources(cpuTotal uint64, cpuUsage uint64, memoryTotal uint64, memoryUsage uint64, diskTotal uint64, diskUsage uint64) *Resources {
	resources := Resources{}

	resources.CpuTotal = cpuTotal
	resources.CpuUsage = cpuUsage
	resources.MemoryTotal = memoryTotal
	resources.MemoryUsage = memoryUsage
	resources.DiskTotal = diskTotal
	resources.DiskUsage = diskUsage

	return &resources
}
