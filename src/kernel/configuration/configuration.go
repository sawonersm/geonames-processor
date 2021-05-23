package configuration

import (
	"encoding/json"
	"io/ioutil"
	Path "path"
	Filepath "path/filepath"
)

type Geonames struct {
	Url       string   `json:"url"`
	Countries []string `json:"countries"`
}

type Temporary struct {
	Input string
}

type Configuration struct {
	Temporary *Temporary
	Geonames  Geonames `json:"geonames"`
}

func New(path string) *Configuration {
	var err error
	var file []byte

	configuration := &Configuration{}

	file, err = ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &configuration)
	if err != nil {
		panic(err)
	}

	// Temporary configurations
	tmpPath := Filepath.Dir(path)
	tmpPath = Path.Join(tmpPath, "tmp")
	configuration.Temporary = &Temporary{tmpPath}
	return configuration
}
