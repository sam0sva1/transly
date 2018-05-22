package main

import (
	"os"
	"encoding/json"
)

type Config struct {
	Port string
}

func CreateConfig(fileName ...string) (*Config, error) {
	var config *Config
	if len(fileName) > 0 {
		file, err := os.Open(fileName[0])
		if err != nil {
			return nil, err
		}

		config = &Config{}
		json.NewDecoder(file).Decode(config)
	} else {
		config = &Config{
			"8080",
		}
	}

	return config, nil
}