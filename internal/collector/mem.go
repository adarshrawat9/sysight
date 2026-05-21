package collector

import(
	"github.com/shirou/gopsutil/v3/mem"
)

type MemoryStats struct{
	TotalRam float64
	UsedRam float64
	AvialableRam float64
}

func GetMemoryStats()(MemoryStats, error){
	memStats, err := mem.VirtualMemory()
	if err != nil{
		return MemoryStats{}, err
	}
	gb := float64(1024*1024*1024)
	return  MemoryStats{
		TotalRam: float64(memStats.Total) /gb,
		UsedRam: float64(memStats.Used) /gb,
		AvialableRam: float64(memStats.Available) /gb,

	}, nil
}


