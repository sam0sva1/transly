package tools

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// It's all because of gorm work
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satori/go.uuid"
)

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

func convert() []OldEx {
	file, err := os.Open("./data.json")
	Chk(err)

	var oldExList []OldEx
	json.NewDecoder(file).Decode(&oldExList)

	return oldExList
}

func collect(list []OldEx) []NewEx {
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

type User struct {
	gorm.Model
	Login    string `gorm:"type:text;unique;not null"`
	Password string `gorm:"type:text;not null"`
	Name     string `gorm:"type:text"`
}

type Exercise struct {
	gorm.Model
	Rus  string  `gorm:"type:text;not null"`
	Eng  string  `gorm:"type:text;not null"`
	Rank float32 `gorm:"type:decimal;not null;default:0"`
}

var err error

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=transly sslmode=disable")
	Chk(err)
	defer db.Close()

	oldList := convert()
	_ = collect(oldList)

	// for _, exercise := range newList {
	// 	dbEx := Exercise{
	// 		Rus:  exercise.Rus,
	// 		Eng:  exercise.Eng,
	// 		Rank: exercise.Rank,
	// 	}

	// 	db.Create(&dbEx)
	// }
}
