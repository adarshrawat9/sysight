package cmd

import (
	"fmt"
	"os"
	"sysight/internal/collector"

	"github.com/spf13/cobra"
)

var cores bool

var coresCmd = &cobra.Command{
	Use : "core",
	Short : "A system that counts no of cores",
	Run : func(cmd *cobra.Command , args []string){
		core, err := collector.GetCpuCores(cores)
		if err != nil{
			fmt.Println("Error: ", err)
		}
		if cores{
			fmt.Println("your pc has ", core, "logical cores")
		}else{
			fmt.Println("your pc has ", core, "physical cores")
		}
		
	},

}

func Execute(){
	if err := coresCmd.Execute();err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func init(){
	coresCmd.Flags().BoolVarP(&cores , "cores" , "c" , false , "print static system summary and exit")
}