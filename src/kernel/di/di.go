package di

import (
	"os"
	Path "path"
	Arguments "sawonersm/geonames-processor/kernel/arguments"
	Configuration "sawonersm/geonames-processor/kernel/configuration"
	Db "sawonersm/geonames-processor/kernel/db"
)

type Di struct {
	Arguments     *Arguments.Arguments
	Configuration *Configuration.Configuration
	db            *Db.Connection
}

func New() *Di {
	path, _ := os.Getwd()
	path = Path.Join(path, "../configuration.json")
	configuration := Configuration.New(path)

	return &Di{
		Arguments.New(),
		configuration,
		Db.NewMysqlConnection(),
	}
}

func (self *Di) GetDb() *Db.Connection {
	return self.db
}
