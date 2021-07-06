package unzipper

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Unzip(path string) []string {

	var files = []string{}
	folder := filepath.Dir(path)

	r, err := zip.OpenReader(path)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	fmt.Println()

	for _, f := range r.File {

		fpath := filepath.Join(folder, f.Name)
		filename := fpath

		files = append(files, filename)

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

	return files
}
