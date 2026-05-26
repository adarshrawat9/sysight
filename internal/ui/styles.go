package ui

import "github.com/charmbracelet/lipgloss"

var titleStyle = lipgloss.NewStyle().
Bold(true).
Foreground(lipgloss.Color("99")).
Padding(1)


var sectionStyle = lipgloss.NewStyle().
Border(lipgloss.DoubleBorder()).
BorderForeground(lipgloss.Color("63")).
Padding(0,1).
Width(44)

var labelStyle = lipgloss.NewStyle().
Foreground(lipgloss.Color("241"))



var valueStyle = lipgloss.NewStyle().
Foreground(lipgloss.Color("255")).
Bold(true)



var barFilledStyle = lipgloss.NewStyle().
Foreground(lipgloss.Color("46")) 


var barEmptyStyle = lipgloss.NewStyle().
Foreground(lipgloss.Color("238"))