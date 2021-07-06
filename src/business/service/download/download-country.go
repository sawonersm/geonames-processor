package download

import (
	Di "sawonersm/geonames-processor/kernel/di"
)

func DownloadCountry(di *Di.Di) string {
	url := di.Configuration.Geonames.Country.Url + "/" + *di.Arguments.Country + ".zip"
	return DownloadFile(
		url,
		di.Configuration.Temporary.Path,
	)
}
