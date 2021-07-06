package features

import (
	"bufio"
	"os"
	Model "sawonersm/geonames-processor/data/model"
	FeatureRepository "sawonersm/geonames-processor/data/repository/feature"
	Di "sawonersm/geonames-processor/kernel/di"
	"strings"
)

func ProcessFile(di *Di.Di, path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	for _, line := range lines {
		processLine(di, line)
	}

}

func processLine(di *Di.Di, line string) {
	var feature *Model.Feature
	parts := strings.Split(line, "\t")

	code := parts[0]
	feature = FeatureRepository.GetByCode(di, code)

	if feature != nil {
		return
	}

	feature = &Model.Feature{
		Code:        parts[0], // Code
		Name:        parts[1], // Name
		Description: parts[2], // Description
	}

	di.GetDb().GetConnection().Save(feature)
}
