package processor

import (
	Unzipper "sawonersm/geonames-processor/business/service/unzipper"
	Di "sawonersm/geonames-processor/kernel/di"
	Processable "sawonersm/geonames-processor/model/processable"
)

func Process(di *Di.Di, processables []*Processable.Processable) {
	for _, processable := range processables {
		processFile(processable)
	}
}

func processFile(processable *Processable.Processable) {
	if processable.IsCountry() {
		Unzipper.UnzipCountry(processable)
	}
}
