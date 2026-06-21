SYSight 🖥️

A fast, beautiful system monitor for your terminal — built with Go.

Features


CPU — model name, core count (physical/logical), usage per core
Memory — RAM usage, swap/page file info
Disk — all drives, space used/free, file system type
Host — OS info, hostname, uptime, boot time
Live Dashboard — real-time updating TUI with progress bars


Installation

bashgit clone https://github.com/adarshrawat9/sysight.git
cd sysight
go build -o sysight.exe .

Usage

Live Dashboard

bashsysight -d

CPU Information

bashsysight cpu              # full CPU overview
sysight cpu cores        # physical core count
sysight cpu cores -l     # logical core count
sysight cpu usage        # per core usage %
sysight cpu info         # model, speed, cores

Memory Information

bashsysight memory           # RAM and swap usage

Disk Information

bashsysight disk             # all drives with usage

Host Information

bashsysight host             # OS, hostname, uptime, boot time

Dashboard Preview

╔══════════════════════════════════════════╗
║  SYSIGHT — System Monitor                ║
╠══════════════════════════════════════════╣
║  CPU  : 12th Gen Intel Core i5-12450H    ║
║  Usage: [████████░░░░░░░░░░░░] 40.0%     ║
╠══════════════════════════════════════════╣
║  RAM  : 8.4 GB / 16.0 GB                 ║
║  Usage: [██████████░░░░░░░░░░] 52.5%     ║
╠══════════════════════════════════════════╣
║  Host : DESKTOP-ABC123                   ║
║  OS   : Windows 11                       ║
╚══════════════════════════════════════════╝

Tech Stack

LibraryPurposecobraCLI commands and flagsgopsutilSystem metrics collectionbubbleteaLive TUI dashboardlipglossTerminal styling

Project Structure

sysight/
├── main.go
├── cmd/
│   ├── root.go         ← entry point, dashboard flag
│   ├── cpu.go          ← cpu command
│   ├── memory.go       ← memory command
│   ├── disk.go         ← disk command
│   └── host.go         ← host command
└── internal/
    ├── collector/
    │   ├── cpu.go      ← CPU metrics
    │   ├── mem.go      ← memory metrics
    │   ├── disk.go     ← disk metrics
    │   └── host.go     ← host metrics
    └── ui/
        ├── dashboard.go   ← bubbletea app
        ├── styles.go      ← lipgloss styles
        └── components.go  ← progress bars, rows

Requirements


Go 1.21+
Windows 10/11


License

MIT