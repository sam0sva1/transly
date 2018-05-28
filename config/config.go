package config

import (
	"encoding/json"
	"os"
)

type DBConfig struct {
	Name    string
	Host    string
	Port    int
	Sslmode string
}

type Config struct {
	Host string
	Port string
	DB   DBConfig
}

func Create(fileName ...string) (*Config, error) {
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
			"127.0.0.1",
			"8080",
			DBConfig{
				"transly",
				"localhost",
				5432,
				"disable",
			},
		}
	}

	return config, nil
}
