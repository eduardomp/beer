package services

import (
	"beer/core/model"
	"beer/util"
	"reflect"
	"testing"
)

func TestGetAppsFromConfig(t *testing.T) {
	want := model.Apps{
		Apps: []model.App{
			{
				Name:    "db-1",
				Ports:   []int{5432},
				Host:    "localhost",
				Type:    "container",
				Volumes: []string{"/var/lib/mysql"},
				Dirs:    []string{"/var/lib/mysql"},
			},
			{
				Name:    "web-1",
				Ports:   []int{4000},
				Host:    "localhost",
				Type:    "container",
				Volumes: []string{"/var/lib/mysql"},
				Dirs:    []string{"/var/lib/mysql"},
			},
		},
	}
	configPath := util.Config("BEER_CONFIG_PATH", "../../config.beer.yml")

	got := GetAppsFromConfig(configPath)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetApps() = %v, want %v", got, want)
	}

}
