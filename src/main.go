package main

import (
	Download "sawonersm/geonames-processor/business/service/download"

	Country "sawonersm/geonames-processor/business/service/country"
	Features "sawonersm/geonames-processor/business/service/features"
	Di "sawonersm/geonames-processor/kernel/di"

	Mode "sawonersm/geonames-processor/kernel/arguments/enum/mode"
)

func main() {
	di := Di.New()

	switch di.Arguments.Mode {
	case Mode.COUNTRY:
		file := Download.DownloadCountry(di)
		Country.ProcessCompressedFile(di, file)
		break
	case Mode.FEATURES:
		file := "/tmp/ES.zip"
		// file := Download.DownloadFeatures(di)
		Features.ProcessFile(di, file)
		break
	}

}
