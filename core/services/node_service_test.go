package services

import (
	"beer/core/entities"
	"reflect"
	"testing"
)

func TestGetNodes(t *testing.T) {
	want := entities.Nodes{
		Nodes: []entities.Node{
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
		},
	}

	got := GetNodes()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("LoadNodes() = %v, want %v", got, want)
	}

}
