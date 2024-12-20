package model

type Connection struct {
	SourceApp      App
	TargetApp      App
	Protocol       string
	TargetPort     int
	Status         string
	StartTimestamp string
	EndTimestamp   string
	Duration       string
}

type Connections struct {
	Connections []Connection
}
