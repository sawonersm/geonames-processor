package di

import (
	"os"
	Path "path"
	Configuration "sawonersm/geonames-processor/kernel/configuration"
)

type Di struct {
	Configuration *Configuration.Configuration
}

func New() *Di {
	path, _ := os.Getwd()
	path = Path.Join(path, "../configuration.json")

	return &Di{
		Configuration.New(path),
	}
}
