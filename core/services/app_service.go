package services

import (
	"beer/core/model"
	"beer/util"
	"fmt"
	"os"
	"slices"

	"github.com/shirou/gopsutil/net"

	"gopkg.in/yaml.v3"
)

func GetAppsFromConfig() model.Apps {
	configPath := util.Config("BEER_CONFIG_PATH", "../../config.beer.yml")

	config, error := os.ReadFile(configPath)
	if error != nil {
		panic(error)
	}

	apps := model.Apps{}
	err := yaml.Unmarshal([]byte(config), &apps)
	if err != nil {
		panic(err)
	}

	return apps

}

func IsAppPortsInProcessConnections(app model.App, connections []net.ConnectionStat) bool {
	ports := make([]int, 0)
	for _, connection := range connections {
		if slices.Contains(app.Ports, int(connection.Laddr.Port)) {
			ports = append(ports, int(connection.Laddr.Port))
		}
	}

	return len(ports) > 0
}

func GetApps() []model.App {
	apps := GetAppsFromConfig()
	processes := GetProcessInfo()

	fmt.Printf("Apps: %v\n", len(apps.Apps))
	fmt.Printf("Processes: %v\n", len(processes))

	result := make([]model.App, 0)

	for _, app := range apps.Apps {
		for _, process := range processes {
			if IsAppPortsInProcessConnections(app, process.Connections) {
				app.Process = process
				result = append(result, app)
			}
		}
	}

	return result
}
