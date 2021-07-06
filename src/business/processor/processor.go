package processor

import (
	"bufio"
	"fmt"
	"os"
	Processable "sawonersm/geonames-processor/business/model/processable"
	Unzipper "sawonersm/geonames-processor/business/service/unzipper"
	Di "sawonersm/geonames-processor/kernel/di"
	"strings"
)

func Process(di *Di.Di, processables []*Processable.Processable) {
	for _, processable := range processables {
		processFile(di, processable)
	}
}

func processFile(di *Di.Di, processable *Processable.Processable) {
	if processable.IsCountry() {
		processable.UncompressedPath = Unzipper.UnzipCountry(processable)
	}

	file, _ := os.Open(processable.UncompressedPath)

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

	fmt.Println(parts[0]) // geonameid
	fmt.Println(parts[1]) // name
	fmt.Println(parts[4]) // latitude
	fmt.Println(parts[5]) // longitude
	fmt.Println("===============")

}
