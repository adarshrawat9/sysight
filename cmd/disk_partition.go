package cmd

import (
	"fmt"
	"sysight/internal/collector"

	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use: "disk",
	Short: "used to show partitions",
	Run: func(cmd *cobra.Command, args[] string){
		partition, err := collector.Getpartitions()
		if err != nil{
			fmt.Print("an error occurced while fetching partitions", err)
			return
		}
		for _,partitions := range partition{
			fmt.Printf("Drive : %s   File type : %s  ", partitions.Device, partitions.Fstype)
			fmt.Println()

		}

		usage , err := collector.Usage()
		if err != nil{
			fmt.Println("error while fetching disk usage")
			return
		}
		for _,u := range usage{
			 fmt.Println("path : " , u.Path)
			 fmt.Printf("total : %v gb \n" , u.Total)
			 fmt.Printf("usage : %v gb \n" , u.Used)
			 fmt.Printf("free : %v gb \n" , u.Free)

		}

	},
}


func init(){
	rootCmd.AddCommand(diskCmd)
}