package dbdrive

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Login    string `gorm:"type:text;unique;not null"`
	Password string `gorm:"type:text;not null"`
	Name     string `gorm:"type:text"`
}
