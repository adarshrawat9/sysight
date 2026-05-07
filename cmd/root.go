package cmd

import (
	"fmt"
	"os"
	"sysight/internal/collector"

	"github.com/spf13/cobra"
)


var rootCmd = cobra.Command{
	Use: "sysight",
	Short: "A system monitor tool",
}

func Execute(){
	if err := rootCmd.Execute();err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
