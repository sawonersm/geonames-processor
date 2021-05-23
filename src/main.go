package main

import (
	// Download "sawonersm/geonames-processor/business/download"
	Processor "sawonersm/geonames-processor/business/processor"
	Inspector "sawonersm/geonames-processor/business/service/inspector"
	Di "sawonersm/geonames-processor/kernel/di"
)

func main() {
	di := Di.New()

	// Download.DownloadCountries(di)

	processables := Inspector.Inspect(di)

	Processor.Process(di, processables)
}
