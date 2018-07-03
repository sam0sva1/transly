package dbdrive

import (
	"transly/config"

	"github.com/jinzhu/gorm"
)

type Exercise struct {
	gorm.Model
	Rus  string  `json:"rus" gorm:"type:text;not null"`
	Eng  string  `json:"eng" gorm:"type:text;not null"`
	Rank float32 `json:"rank" gorm:"type:decimal;not null;default:0"`
}

type ExerciseRepository struct {
	db *gorm.DB
}

func (repo *ExerciseRepository) GetCollection(args ...int) []*Exercise {
	var exercises []*Exercise
	limit := 10
	offset := -1

	switch len(args) {
	case 1:
		limit = args[0]
	case 2:
		limit = args[0]
		offset = args[1]
	}

	repo.db.Limit(limit).Offset(offset).Find(&exercises)
	return exercises
}

func CreateExerciseRepository(database *gorm.DB) *ExerciseRepository {
	return &ExerciseRepository{db: database}
}

type ExerciseService struct {
	config     *config.Config
	repository *ExerciseRepository
}

// GetCollection is the best place to make some constrains according to config
func (es *ExerciseService) GetCollection(args ...int) []*Exercise {
	// es.config
	return es.repository.GetCollection(args...)
}

func (conn *Connection) CreateExerciseService() *ExerciseService {
	repo := CreateExerciseRepository(conn.db)
	return &ExerciseService{config: conn.config, repository: repo}
}
