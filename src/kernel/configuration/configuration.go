package configuration

import (
	"encoding/json"
	"io/ioutil"
)

type Geonames struct {
	Url       string   `json:"url"`
	Countries []string `json:"countries"`
	Features  Features `json:"features"`
}

type Temporary struct {
	Path string `json:"path"`
}

type Configuration struct {
	Temporary Temporary `json:"temporary"`
	Geonames  Geonames  `json:"geonames"`
}

type Features struct {
	Url string `json:"url"`
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

	return configuration
}
