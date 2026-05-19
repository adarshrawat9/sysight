package collector

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/host"
)

type HostInfo struct{
	Hostname string
	Uptime string
	BootTime string
	Procs uint64
	OS string
	PlatformVersion string
	HostId string
}

func GetHostInfo() (HostInfo, error){
  
	info, err := host.Info()
	if err != nil{
		return HostInfo{}, err
	}

	return HostInfo{
		Hostname: info.Hostname,
		Uptime: formatTime(info.Uptime),
		BootTime: time.Unix(int64(info.BootTime), 0).Format("Jan 02 2006 15:04:05"),
		Procs: info.Procs,
		OS: info.OS,
		PlatformVersion: info.Platform,
		HostId: info.HostID,
	}, nil

	

}

func formatTime(seconds uint64) string{
	hours := seconds/3600
	minutes := (seconds % 3600) / 60
	return fmt.Sprintf("%dh %dm" , hours, minutes)
}