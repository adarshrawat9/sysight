package cmd

import (
	"fmt"
	"sysight/internal/collector"

	"github.com/spf13/cobra"
)


var hostCmd = &cobra.Command{
	Use: "host",
	Short: "use for host info",
	Run: func(cmd *cobra.Command, args []string){
		info, err := collector.GetHostInfo()
		if err != nil{
			fmt.Println("some error caught while getting hsot info", err)
			return
		}
		fmt.Println("Host name : ", info.Hostname)
		fmt.Println("Hostid: ", info.HostId)
		fmt.Println("Host OS : ", info.OS)
		fmt.Println("Host platform version : ", info.PlatformVersion)
		fmt.Println("Uptime : ", info.Uptime)
		fmt.Println("processes : ", info.Procs)
		fmt.Println("Boot time : ", info.BootTime)
	},
}



func init(){
	rootCmd.AddCommand(hostCmd)
}