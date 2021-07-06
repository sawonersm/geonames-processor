package country

import (
	"fmt"
	Unzipper "sawonersm/geonames-processor/business/service/unzipper"
	Di "sawonersm/geonames-processor/kernel/di"
)

func ProcessCompressedFile(di *Di.Di, path string) {
	files := Unzipper.Unzip(path)

	fmt.Println(files)
}
