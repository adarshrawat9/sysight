package collector

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

type usageStats struct{
	Path              string
	Fstype            string  
	Total             uint64 
	Free              uint64
	Used              uint64
}

func getpartitions()([]disk.PartitionStat, error){
	partition, err := disk.Partitions(false)
	if err != nil{
		return []disk.PartitionStat{}, err
	}
	return partition, nil
}
var gb float64 = float64(1024*1024*1024)

func Usage ()(*usageStats, error){

	partition, err := getpartitions()
	if err != nil{
		return &usageStats{},err
	}

	for _,p := range partition{
		usage, err disk.Usage(p.Mountpoint)
		if err !=  nil{
			return &usageStats{}, err
		}
		fmt.Println(p.Mountpoint)
    fmt.Printf("Total : %.1f GB\n", float64(usage.Total) / gb)
    fmt.Printf("Used  : %.1f GB\n", float64(usage.Used) / gb)
    fmt.Printf("Free  : %.1f GB\n", float64(usage.Free) / gb)
    fmt.Printf("Usage : %.1f%%\n\n", usage.UsedPercent)

		
	}
}