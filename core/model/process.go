package model

import "github.com/shirou/gopsutil/net"

type Process struct {
	Name        string
	Pid         int
	Status      string
	Cmdline     string
	Connections []net.ConnectionStat
}

type Processes struct {
	Processes []Process
}
