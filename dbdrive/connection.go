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

func Connect(config *config.Config) *Connection {
	dbConf := config.DB
	params := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=%s", dbConf.Host, dbConf.Port, dbConf.Name, dbConf.Sslmode)
	db, err := gorm.Open("postgres", params)

	tools.Chk(err)

	conn := &Connection{
		db:     db,
		config: config,
	}
	return conn
}
