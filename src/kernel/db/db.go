package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	connection *gorm.DB
}

func NewMysqlConnection() *Connection {

	db, err := gorm.Open(
		mysql.Open(
			"root:root@tcp(processor-db:3306)/geonames?charset=utf8mb4&parseTime=True&loc=Local",
		),
		&gorm.Config{},
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
