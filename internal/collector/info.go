package collector

import(
	"github.com/shirou/gopsutil/v3/cpu"
)

type CpuInfo struct{
	Cpu  int32
	Model string
	ModelName string
	Mhz float64
	CacheSize int32
	Cores int32
}

func ShowInfo ()(CpuInfo, error){
	info, err := cpu.Info()
	if err != nil{
		return CpuInfo{},err
	}
	return CpuInfo{
		Cpu: info[0].CPU,
		Model: info[0].Model,
		ModelName: info[0].ModelName,
		Mhz: info[0].Mhz,
		Cores: info[0].Cores,
		CacheSize: info[0].CacheSize,
	}, nil


}