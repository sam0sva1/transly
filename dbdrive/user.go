package dbdrive

import (
	"time"
	"transly/config"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Login     string     `gorm:"type:text;unique;not null" json:"login" binding:"required"`
	Password  string     `gorm:"type:text;not null" json:"password" binding:"required"`
	Name      string     `gorm:"type:text" json:"name"`
	Status    string     `gorm:"type:text;default:'active'" json:"status"`
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) GetUserById(id int) (*User, bool) {
	user := User{}

	isExist := repo.db.First(&user, id).RecordNotFound()
	return &user, isExist
}

func (repo *UserRepository) Check(user *User) bool {
	isExist := repo.db.Where("login = ?", user.Login).Find(&user).RecordNotFound()

	return !isExist
}

func (repo *UserRepository) Create(user *User) error {
	db := repo.db.Create(user)
	return db.Error
}

func CreateUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{db: database}
}

type UserService struct {
	config     *config.Config
	repository *UserRepository
}

func (us *UserService) Create(user *User) error {
	return us.repository.Create(user)
}

func (us *UserService) Check(user *User) bool {
	return us.repository.Check(user)
}

func (us *UserService) GetUserById(id int) (*User, bool) {
	return us.repository.GetUserById(id)
}

func (conn *Connection) CreateUserService() *UserService {
	repo := CreateUserRepository(conn.db)
	return &UserService{config: conn.config, repository: repo}
}
