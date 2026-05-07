package collector

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
)

var cores int

func GetCpuCores(islogical bool)(int , error){

	core , err := cpu.Counts(islogical)
	if err != nil{
		fmt.Println("some error occured while getting cores")
		return cores , err
	}
	cores = core


	return core, nil

}
