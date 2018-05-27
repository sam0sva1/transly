package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/satori/go.uuid"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// type Config struct {
// 	Host string
// 	Port string
// }

// func CreateConfig(fileName ...string) (*Config, error) {
// 	var config *Config
// 	if len(fileName) > 0 {
// 		file, err := os.Open(fileName[0])
// 		if err != nil {
// 			return nil, err
// 		}

// 		config = &Config{}
// 		json.NewDecoder(file).Decode(config)
// 	} else {
// 		config = &Config{
// 			"127.0.0.1",
// 			"8080",
// 		}
// 	}

// 	return config, nil
// }

type OldEx struct {
	ID     string   `json:"id"`
	Number int      `json:"number"`
	Name   string   `json:"name"`
	Ru     []string `json:"ru"`
	En     []string `json:"en"`
}

type NewEx struct {
	ID   string  `json:"id"`
	Rus  string  `json:"rus"`
	Eng  string  `json:"eng"`
	Rank float32 `json:"rank"`
}

func Convert() []OldEx {
	file, err := os.Open("./data.json")
	check(err)

	var oldExList []OldEx
	json.NewDecoder(file).Decode(&oldExList)

	return oldExList
}

func Collect(list []OldEx) []NewEx {
	var newExList []NewEx

	for _, oldOne := range list {
		if len(oldOne.Ru) == len(oldOne.En) {
			innerLen := len(oldOne.Ru)
			for i := 0; i < innerLen; i++ {
				id := uuid.Must(uuid.NewV4())
				ex := NewEx{
					ID:   fmt.Sprintf("%s", id),
					Rus:  oldOne.Ru[i],
					Eng:  oldOne.En[i],
					Rank: 0,
				}
				newExList = append(newExList, ex)
			}
		}
	}

	return newExList
}

func main() {
	oldList := Convert()
	newList := Collect(oldList)

	fmt.Println(len(newList))
}
