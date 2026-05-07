package collector

import (
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func CpuPercentUsage(percpu bool)([]float64, error){
	cpuUse, err := cpu.Percent(time.Second, percpu)
	if err != nil{
		return nil, err
	}
	return cpuUse , nil
}
