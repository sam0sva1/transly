package main

import (
	"transly/config"
	"transly/dbdrive"
	"transly/serving"
	"transly/tools"
)

var err error

func main() {
	config, err := config.Create("config.json")
	tools.Chk(err)

	connection := dbdrive.Connect(config, true)

	server := serving.Server{
		Config:          config,
		ExerciseService: connection.CreateExerciseService(),
		UserService:     connection.CreateUserService(),
	}

	server.Run()
}
