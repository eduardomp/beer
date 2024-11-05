package services

import (
	"beer/core/model"
	"beer/util"
	"os"

	"gopkg.in/yaml.v3"
)

func GetApps() model.Apps {
	configPath := util.Config("BEER_CONFIG_PATH", "../../config.yml")

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
