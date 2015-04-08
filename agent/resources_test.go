package agent

import (
	"testing"
)

func TestNewResource(t *testing.T) {
	r := NewResources(1, 1, 1, 1, 1, 1)

	if r.CpuTotal != 1 || r.CpuUsage != 1 || r.MemoryTotal != 1 || r.MemoryUsage != 1 || r.DiskTotal != 1 || r.DiskUsage != 1 {
		t.Fatal("Fail to construct resources struct")
	}
}
