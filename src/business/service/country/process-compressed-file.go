package country

import (
	"path/filepath"
	Unzipper "sawonersm/geonames-processor/business/service/unzipper"
	Di "sawonersm/geonames-processor/kernel/di"
)

func ProcessCompressedFile(di *Di.Di, path string) {
	files := Unzipper.Unzip(path)

	var txtPath string
	filename := *di.Arguments.Country + ".txt"
	for _, file := range files {
		if filename == filepath.Base(file) {
			txtPath = file
			break
		}
	}

	ProcessFile(di, txtPath)
}
