package dbdrive

import (
	"time"
	"transly/config"

	"github.com/jinzhu/gorm"
)

type User2Exer struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	UserID     string  `gorm:"type:text;not null" json:"user_id" binding:"required"`
	ExerciseID string  `gorm:"type:text;not null" json:"exercise_id" binding:"required"`
	Rank       float32 `gorm:"type:integer;not null;default:0" json:"rank"`
	Progress   int     `gorm:"type:integer;default:'active'" json:"progress"`
}

type User2ExerRepository struct {
	db *gorm.DB
}

func (u2er *User2ExerRepository) CreateRelation(u2e *User2Exer) error {
	return nil
}

func CreateUser2ExerRepository(database *gorm.DB) *User2ExerRepository {
	return &User2ExerRepository{db: database}
}

type User2ExerService struct {
	config     *config.Config
	repository *User2ExerRepository
}

func (server *User2ExerService) CreateRelation(u2e *User2Exer) error {
	return server.CreateRelation(u2e)
}

func (conn *Connection) CreateUser2ExerService() *User2ExerService {
	repo := CreateUser2ExerRepository(conn.db)
	return &User2ExerService{config: conn.config, repository: repo}
}
