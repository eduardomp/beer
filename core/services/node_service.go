package services

import (
	"beer/core/entities"
	"beer/util"
	"os"

	"gopkg.in/yaml.v3"
)

func GetNodes() entities.Nodes {
	configPath := util.Config("BEER_CONFIG_PATH", "../../config.yml")

	config, error := os.ReadFile(configPath)
	if error != nil {
		panic(error)
	}

	nodes := entities.Nodes{}
	err := yaml.Unmarshal([]byte(config), &nodes)
	if err != nil {
		panic(err)
	}

	return nodes

}
