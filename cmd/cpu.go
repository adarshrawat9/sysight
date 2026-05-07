package cmd

import (
	"fmt"
	"sysight/internal/collector"

	"github.com/spf13/cobra"
)

var cores bool

var cpuCmd = &cobra.Command{
	Use : "cpu",
	Short : "A tool that would get you any information regarding the cpu",
	Run : func(cmd *cobra.Command , args []string){	
		core, err := collector.GetCpuCores(cores)
		if err != nil{
			fmt.Println("Error: ", err)
			return
		}
		if cores{
			fmt.Println("your pc has ", core, "logical cores")
		}else{
			fmt.Println("your pc has ", core, "physical cores")
		}
		
	},

}

func init(){
	cpuCmd.Flags().BoolVarP(&cores , "cores" , "c" , false , "print static system summary and exit")
	rootCmd.AddCommand(cpuCmd)
}