package download

import (
	Di "sawonersm/geonames-processor/kernel/di"
)

func DownloadFeatures(di *Di.Di) string {
	return DownloadFile(
		di.Configuration.Geonames.Features.Url,
		di.Configuration.Temporary.Path,
	)
}
