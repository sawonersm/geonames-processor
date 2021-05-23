package unzipper

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"regexp"
	Processable "sawonersm/geonames-processor/model/processable"
)

func UnzipCountry(processable *Processable.Processable) string {
	var filename string

	r, err := zip.OpenReader(processable.Filepath)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	for _, f := range r.File {
		// Prevent readmes and kind of stuff
		if !isUncompressedCountry(processable.Country.Code, f) {
			continue
		}

		fpath := filepath.Join(processable.Path, f.Name)
		filename = fpath

		targetFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		sourceFile, err := f.Open()
		if err != nil {
			panic(err)
		}

		_, err = io.Copy(targetFile, sourceFile)

		targetFile.Close()
		sourceFile.Close()
	}

	return filename
}

func isUncompressedCountry(code string, file *zip.File) bool {
	is, _ := regexp.Match(code+"(.*)", []byte(file.Name))
	return is
}
