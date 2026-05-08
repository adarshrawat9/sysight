package cmd

import (
	"fmt"
	"sysight/internal/collector"

	"github.com/spf13/cobra"
)


var memoryCmd = &cobra.Command{
	Use: "mem",
	Short: "show memory stats",
	Run: func(cmd *cobra.Command , args []string){
		memStats, err := collector.GetMemoryStats()
		if err !=  nil{
			println("error getting memory stats")
			return
		}
		fmt.Printf("total ram : %.2f gb\n" , memStats.TotalRam)
		fmt.Printf("used ram : %.2f gb\n" , memStats.UsedRam)
		fmt.Printf("avialable ram : %.2f gb\n" , memStats.AvialableRam)


	},

}


func init(){
	rootCmd.AddCommand(memoryCmd)
}