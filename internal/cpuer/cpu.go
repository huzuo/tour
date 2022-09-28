package cpuer

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func GetCpuInfo() []float64 {
	p, _ := cpu.Percent(time.Second, false)
	return p
}
