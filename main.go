package main

import (
	"transly/config"
	"transly/dbdrive"
	"transly/tools"
)

var err error

func main() {
	config, err := config.Create("config.json")
	tools.Chk(err)

	connection := dbdrive.Connect(config)

	server := Server{
		config:          config,
		exerciseService: connection.CreateExerciseService(),
	}

	server.Run()
}
