package dbdrive

import (
	"fmt"
	"transly/config"
	"transly/tools"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var err error

type Connection struct {
	db     *gorm.DB
	config *config.Config
}

func Connect(config *config.Config, withMigrations bool) *Connection {
	dbConf := config.DB
	params := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=%s", dbConf.Host, dbConf.Port, dbConf.Name, dbConf.Sslmode)
	db, err := gorm.Open("postgres", params)

	db.LogMode(true)

	tools.Chk(err)

	conn := &Connection{
		db:     db,
		config: config,
	}

	if withMigrations {
		db.AutoMigrate(&User{}, &Exercise{})
	}

	return conn
}
