package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	Logger "gorm.io/gorm/logger"
)

type Connection struct {
	connection *gorm.DB
}

func NewMysqlConnection() *Connection {
	logger := Logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		Logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  Logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	db, err := gorm.Open(
		mysql.Open(
			"root:root@tcp(processor-db:3306)/geonames?charset=utf8mb4&parseTime=True&loc=Local",
		),
		&gorm.Config{
			Logger: logger,
		},
	)

	if err != nil {
		panic(err)
	}

	return &Connection{
		db,
	}
}

func (self *Connection) GetConnection() *gorm.DB {
	return self.connection
}
