package collector

import (

	"github.com/shirou/gopsutil/v3/disk"
)

type usageStats struct{
	Path              string
	Fstype            string  
	Total             uint64 
	Free              uint64
	Used              uint64
}

func Getpartitions()([]disk.PartitionStat, error){
	partition, err := disk.Partitions(false)
	if err != nil{
		return []disk.PartitionStat{}, err
	}
	return partition, nil
}
var gb float64 = float64(1024*1024*1024)

func Usage ()([]usageStats, error){

	partition, err := Getpartitions()
	if err != nil{
		return nil,err
	}

	var result []usageStats

	for _,p := range partition{
		usage, err := disk.Usage(p.Mountpoint)
		if err !=  nil{
			continue
		}

		result = append(result, usageStats{
			Path: p.Mountpoint ,
			Fstype: usage.Fstype,
			Total: usage.Total/uint64(gb),
			Used: usage.Used/uint64(gb),
			Free: usage.Free/uint64(gb),

		})
		
	}
	return result, nil
}