package services

import (
	"core/entities"
	"core/util"
)

func GetNodes() []entities.Node {
	return util.LoadNodes()
}
