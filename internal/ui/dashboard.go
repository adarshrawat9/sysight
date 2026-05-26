package ui

import (
	"fmt"
	"sysight/internal/collector"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//model app state change
type model struct{
	//static data
	cpu collector.CpuInfo
	host collector.HostInfo

	//dynamic data
	memory collector.MemoryStats
}

type staticdata struct{
	cpu collector.CpuInfo
	host collector.HostInfo
}

type tickMsg time.Time

func tick() tea.Cmd{
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func InitialModel()model{
	return model{}
}

func (m model) Update (msg tea.Msg)(tea.Model, tea.Cmd){
	switch msg := msg.(type){


	case staticdata:
		m.cpu = msg.cpu
		m.host = msg.host
		
		return m, nil
		
	case tickMsg:
	
		memory, _ := collector.GetMemoryStats()
		
		m.memory = memory
		return m, tick()

	case tea.KeyMsg:
		switch msg.String(){
		case "q","ctrl+c":
			return m, tea.Quit
		}	
	}
	return m, nil
}


func (m model)Init() tea.Cmd{


	return tea.Batch(fetchstaticdata, tick())
}

func fetchstaticdata()tea.Msg{
	cpu, _ := collector.ShowInfo()
	host,_ := collector.GetHostInfo()


	return staticdata{cpu: cpu, host: host}

}


func (m model) View() string{

	if m.cpu.ModelName == ""{
		return "loading...\n"
	}
	title := sectionStyle.Render(
		titleStyle.Render("Sysight"),
	)

	cpuSection := sectionStyle.Render(
		fmt.Sprintf("%s\n%s", row("Cpu  : " , m.cpu.ModelName) ,
		 row("usage  : " , progressBar(m.memory.UsedRam, 20))),
	)

	ramSection := sectionStyle.Render(
		fmt.Sprintf("%s\n%s" , 
	row("Ram  : " , fmt.Sprintf("%.1f GB / %.1f Gb" , m.memory.UsedRam, m.memory.TotalRam)),
    row("Usage  : " , progressBar(m.memory.UsedPercentage , 20))),
	)

	hostSection := sectionStyle.Render(
		fmt.Sprintf("%s", 
	    row("Host  : " , m.host.Hostname),
	
	),
)


	return lipgloss.JoinVertical(lipgloss.Left,
	title,
     cpuSection,
    ramSection,
	hostSection,
	labelStyle.Render("\n Press q to quit"),
	)
}

