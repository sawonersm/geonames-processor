package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
	Path "path"
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
	configuration.Temporary = &Temporary{}
	tmpFolder, _ := ioutil.TempDir("", "geonames")
	configuration.Temporary.Input = Path.Join(tmpFolder, "input")
	os.Mkdir(configuration.Temporary.Input, 0755)

	return configuration
}
