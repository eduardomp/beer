package util

import "core/entities"

func LoadNodes() []entities.Node {

	//TODO: Will read the config.yml file to load the nodes

	return []entities.Node{
		{
			Name:    "node1",
			Ports:   []int{3000, 3001, 3002},
			Host:    "localhost",
			Type:    "container",
			Volumes: []string{"/var/lib/mysql"},
			Dirs:    []string{"/var/lib/mysql"},
		},
		{
			Name:    "node2",
			Ports:   []int{3003, 3004},
			Host:    "localhost",
			Type:    "application",
			Volumes: []string{"/var/lib/mysql"},
			Dirs:    []string{"/var/lib/mysql"},
		},
	}
}
