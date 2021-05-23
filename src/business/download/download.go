package download

import (
	"io"
	"net/http"
	"os"
	"path"

	Di "sawonersm/geonames-processor/kernel/di"
)

func DownloadCountries(di *Di.Di) []*os.File {
	var countries []*os.File
	for _, country := range di.Configuration.Geonames.Countries {
		countries = append(
			countries,
			downloadFile(di.Configuration.Geonames.Url, country, di.Configuration.Temporary.Input),
		)
	}

	return countries
}

func downloadFile(url string, country string, folder string) *os.File {
	filePath := path.Join(folder, country+".zip")
	file, _ := os.Create(filePath)
	defer file.Close()

	url = url + "/" + country + ".zip"
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}

	return file
}
