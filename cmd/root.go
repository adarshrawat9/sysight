package cmd

import (
	"fmt"
	"os"
	"sysight/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)


var dashboard bool

var rootCmd = cobra.Command{
	Use: "sysight",
	Short: "A system monitor tool",
	Run: func(cmd *cobra.Command, args []string) {
		if dashboard{
			p := tea.NewProgram(ui.InitialModel())
			if err := p.Start();err != nil{
				fmt.Println("error :" , err)
			}
		}
	},
}


func init(){
	rootCmd.Flags().BoolVarP(&dashboard, "dashboard" , "d" , false, "launch live dashbaord")
}



func Execute(){
	if err := rootCmd.Execute();err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
