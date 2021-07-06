package download

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func DownloadFile(url string, folder string) string {

	filename := filepath.Base(url)

	path := path.Join(folder, filename)
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}

	return path
}
