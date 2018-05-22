package main

import "fmt"

func main() {
	config, err := CreateConfig("config.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config)
}