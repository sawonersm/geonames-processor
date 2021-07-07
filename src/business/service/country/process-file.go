package country

import (
	"bufio"
	"os"
	Model "sawonersm/geonames-processor/data/model"
	FeatureRepository "sawonersm/geonames-processor/data/repository/feature"
	Di "sawonersm/geonames-processor/kernel/di"
	"strconv"
	"strings"
)

func ProcessFile(di *Di.Di, path string) {
	file, _ := os.Open(path)

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
	parts := strings.Split(line, "\t")

	featureCode := parts[7]
	feature := FeatureRepository.GetByCode(di, featureCode)
	if feature == nil {
		return
	}

	latitude, _ := strconv.ParseFloat(parts[4], 32)
	longitude, _ := strconv.ParseFloat(parts[5], 32)

	location := &Model.Location{
		GeonameId: parts[0],
		Name:      parts[1],
		Feature:   feature.Id,
		Latitude:  float32(latitude),
		Longitude: float32(longitude),
	}

	di.GetDb().GetConnection().Save(location)
}
