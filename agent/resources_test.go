package agent

import (
	"testing"
)

func TestConstructResource(t *testing.T) {
	r := ConstructResources(1, 1, 1, 1, 1, 1)

	if r.CpuTotal != 1 || r.CpuUsage != 1 || r.MemoryTotal != 1 || r.MemoryUsage != 1 || r.DiskTotal != 1 || r.DiskUsage != 1 {
		t.Fatal("disk stat read fail")
	}
}
