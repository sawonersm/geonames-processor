package inspector

import (
	"io/ioutil"
	"os"
	Path "path"
	"regexp"
	Di "sawonersm/geonames-processor/kernel/di"
	Country "sawonersm/geonames-processor/model/country"
	Processable "sawonersm/geonames-processor/model/processable"
)

const (
	PATTERN_COUNTRY = "^(.*)(.zip)$"
)

func Inspect(di *Di.Di) []*Processable.Processable {
	files, err := ioutil.ReadDir(di.Configuration.Temporary.Input)
	if err != nil {
		panic(err)
	}

	processables := []*Processable.Processable{}
	for _, file := range files {
		if isProcessableFile(file) {
			processables = append(
				processables,
				getProcessable(di, file),
			)
		}
	}
	return processables
}

func isProcessableFile(file os.FileInfo) bool {
	return isCountryFile(
		file.Name(),
	)
}

func isCountryFile(name string) bool {
	r, _ := regexp.Match(PATTERN_COUNTRY, []byte(name))
	return r
}

func getProcessable(di *Di.Di, file os.FileInfo) *Processable.Processable {
	if isCountryFile(file.Name()) {
		return getProcessableCountry(di, file)
	}

	panic("The given file is not processable")
}

func getProcessableCountry(di *Di.Di, file os.FileInfo) *Processable.Processable {
	regex := regexp.MustCompile(PATTERN_COUNTRY)
	matches := regex.FindAllSubmatch([]byte(file.Name()), -1)

	if len(matches) == 0 {
		panic("The given country file is not processable")
	}

	submatch := matches[0]
	if len(submatch) != 3 {
		panic("The given country file has wrong name format")
	}

	code := string(submatch[1])
	filename := string(submatch[0])
	filepath := Path.Join(di.Configuration.Temporary.Input, filename)

	return &Processable.Processable{
		File:     filename,
		Path:     di.Configuration.Temporary.Input,
		Filepath: filepath,
		Country: &Country.Country{
			Code: code,
		},
	}
}
