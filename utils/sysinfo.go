package utils

import (
	"gin-blog-server/models"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/disk"

	"github.com/shirou/gopsutil/mem"

	"github.com/shirou/gopsutil/cpu"
)

const (
	B = iota * 10
	KiB
	MiB
	GiB
)

func OSInfo() *models.OS {
	os := models.OS{
		GOOS:         runtime.GOOS,
		NumCPU:       runtime.NumCPU(),
		Compiler:     runtime.Compiler,
		GoVersion:    runtime.Version(),
		NumGoroutine: runtime.NumGoroutine(),
	}
	return &os
}

func CPUInfo() (*models.Cpu, error) {
	usedCpus, err := cpu.Percent(200*time.Millisecond, true)
	if err != nil {
		return nil, err
	}
	cores, err := cpu.Counts(true)
	if err != nil {
		return nil, err
	}
	c := models.Cpu{
		UsedCpus: usedCpus,
		Cores:    cores,
	}
	return &c, nil
}

func MemInfo() (*models.Ram, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	ram := models.Ram{
		UsedMiB:  vm.Used >> MiB,
		TotalMiB: vm.Total >> MiB,
	}
	return &ram, nil
}

func DiskInfo() (*models.Disk, error) {
	usage, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}
	d := models.Disk{
		UsedMiB:  usage.Used >> MiB,
		UsedGiB:  usage.Used >> GiB,
		TotalMiB: usage.Total >> GiB,
		TotalGiB: usage.Total >> GiB,
	}
	return &d, nil
}
