package download

import (
	"os"

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
