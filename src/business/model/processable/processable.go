package processable

import (
	Country "sawonersm/geonames-processor/business/model/country"
)

type Processable struct {
	File             string
	Path             string
	Filepath         string
	UncompressedPath string
	Country          *Country.Country
}

func (self *Processable) IsCountry() bool {
	return self.Country != nil
}
