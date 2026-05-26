package ui

import (
	"fmt"
	"strings"
)


func progressBar(percent float64 , width int) string{
	if percent < 0{
		percent = 0 
	}
	if percent > 100{
		percent = 100
	} 
	filled := int(percent / 100*float64(width))
	empty := width - filled

	bar := barFilledStyle.Render(strings.Repeat("█", filled)) +
	barEmptyStyle.Render(strings.Repeat("░", empty))


	return fmt.Sprintf("[%s] %.1f%%", bar, percent)
}


func row(lable string, value string) string{
	return fmt.Sprintf("%s %s" , labelStyle.Render(lable), labelStyle.Render(value))
}