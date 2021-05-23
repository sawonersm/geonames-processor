package main

import (
	"fmt"
	Di "sawonersm/geonames-processor/kernel/di"

	Download "sawonersm/geonames-processor/business/download"
)

func main() {
	di := Di.New()

	files := Download.DownloadCountries(di)
	fmt.Println(files)
}
