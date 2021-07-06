package main

import (
	Download "sawonersm/geonames-processor/business/service/download"

	Processor "sawonersm/geonames-processor/business/processor"
	Features "sawonersm/geonames-processor/business/service/features"
	Inspector "sawonersm/geonames-processor/business/service/inspector"
	Di "sawonersm/geonames-processor/kernel/di"

	Mode "sawonersm/geonames-processor/kernel/arguments/enum/mode"
)

func main() {
	di := Di.New()

	switch di.Arguments.Mode {
	case Mode.COUNTRY:
		Download.DownloadCountry(di)
		processables := Inspector.Inspect(di)
		Processor.Process(di, processables)
		break
	case Mode.FEATURES:
		file := Download.DownloadFeatures(di)
		Features.ProcessFile(di, file)
		break
	}

}
