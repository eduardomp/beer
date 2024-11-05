package services

import (
	"beer/core/model"
	"reflect"
	"testing"
)

func TestGetApps(t *testing.T) {
	want := model.Apps{
		Apps: []model.App{
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

	got := GetApps()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetApps() = %v, want %v", got, want)
	}

}
