package dbdrive

import (
	"time"
	"transly/config"

	"github.com/jinzhu/gorm"
)

type Exercise struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Rus       string     `json:"rus" gorm:"type:text;not null"`
	Eng       string     `json:"eng" gorm:"type:text;not null"`
	Rank      float32    `json:"rank" gorm:"type:integer;not null;default:0"`
}

type ExerciseRepository struct {
	db *gorm.DB
}

type ExerciseParams struct {
	Limit  int
	Offset int
}

func (repo *ExerciseRepository) GetCollection(params ExerciseParams) []*Exercise {
	var exercises []*Exercise

	repo.db.Limit(params.Limit).Offset(params.Offset).Find(&exercises)
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
func (es *ExerciseService) GetCollection(params ExerciseParams) []*Exercise {
	return es.repository.GetCollection(params)
}

func (conn *Connection) CreateExerciseService() *ExerciseService {
	repo := CreateExerciseRepository(conn.db)
	return &ExerciseService{config: conn.config, repository: repo}
}
