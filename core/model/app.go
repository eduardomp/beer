package model

type App struct {
	Id      string
	Name    string
	Ports   []int
	Host    string
	Type    string
	Volumes []string
	Dirs    []string
	Process Process
}

type Apps struct {
	Apps []App
}
