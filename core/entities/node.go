package entities

type Node struct {
	Name    string
	Ports   []int
	Host    string
	Type    string
	Volumes []string
	Dirs    []string
}
