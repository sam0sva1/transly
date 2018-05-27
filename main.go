package main

import (
	//"fmt"

)

var err error

func main() {
	config, err := CreateConfig("config.json")
	chkErr(err)

	server := Server{config:config}

	server.Run()
}