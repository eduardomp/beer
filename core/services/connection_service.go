package services

import (
	"beer/core/model"
	"time"
)

func GetConnectionsToChannel(connectionsChan chan model.Connections) {
	for {
		time.Sleep(time.Duration(30) * time.Second)
		connections := GetConnections()
		connectionsChan <- connections
	}
}

func GetConnections() model.Connections {
	timestamp := time.Now()
	connections := model.Connections{
		Connections: []model.Connection{
			{
				SourceApp: model.App{
					Name:  "web-1",
					Host:  "localhost",
					Ports: []int{4000},
				},
				TargetApp: model.App{
					Name:  "db-1",
					Host:  "localhost",
					Ports: []int{5432},
				},
				Protocol:       "HTTP",
				TargetPort:     5432,
				Status:         "ESTABLISHED",
				StartTimestamp: timestamp.Format(time.RFC3339),
				EndTimestamp:   timestamp.Add(time.Duration(10) * time.Second).Format(time.RFC3339),
				Duration:       "10s",
			},
		},
	}
	return connections
}
