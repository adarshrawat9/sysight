package cmd

import (
	"fmt"
	"sysight/internal/collector"

	"github.com/spf13/cobra"
)

var logical bool

var cpuCmd = &cobra.Command{
	Use : "cpu",
	Short : "A tool that would get you any information regarding the cpu",
	Run : func(cmd *cobra.Command , args []string){
		
			core, err := collector.GetCpuCores(logical)
		if err != nil{
			fmt.Println("Error: ", err)
			return
		}
		if logical{
			fmt.Println("your pc has", core, "logical cores")
		}else{
			fmt.Println("your pc has", core, "physical cores")
		}

			usage, err := collector.CpuPercentUsage(true)
			if err != nil{
				fmt.Println("some unknown error occured")
				return
			}
			fmt.Println("cores: " , core)
			fmt.Println("cpu usage")
			for i,use := range usage{
				fmt.Printf("core %d  %.2f%%\n" , i , use)

			}

	},



}

var cpuUsageCmd = cobra.Command{
	Use: "usage",
	Short: "use for showing the cpu usage",
	Run: func(cmd *cobra.Command, args []string) {
		usage, err := collector.CpuPercentUsage(true)
			if err != nil{
				fmt.Println("some unknown error occured")
				return
			}
			fmt.Println("cpu usage")
			for i,use := range usage{
				fmt.Printf("core %d  %.2f%%\n" , i , use)

			}
	},
}

var cpuCoresCmd = cobra.Command{
	Use: "core",
	Short: "use for showing physical and logical cores",
	Run: func (cmd *cobra.Command , args[] string){
		core, err := collector.GetCpuCores(logical)
		if err != nil{
			fmt.Println("Error: ", err)
			return
		}
		if logical{
			fmt.Println("your pc has", core, "logical cores")
		}else{
			fmt.Println("your pc has", core, "physical cores")
		}

	},
}

func init(){
	
	rootCmd.AddCommand(cpuCmd)
	cpuCmd.AddCommand(&cpuCoresCmd)
	cpuCmd.AddCommand(&cpuUsageCmd)
	cpuCoresCmd.Flags().BoolVarP(&logical, "core" , "l" , false , "use for showing number of cores")
}