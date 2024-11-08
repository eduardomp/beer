package model

type Process struct {
	Name    string
	Pid     int
	Status  string
	Cmdline string
}

type Processes struct {
	Processes []Process
}
