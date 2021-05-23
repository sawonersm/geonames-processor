package processable

import (
	Country "sawonersm/geonames-processor/model/country"
)

type Processable struct {
	File     string
	Path     string
	Filepath string
	Country  *Country.Country
}

func (self *Processable) IsCountry() bool {
	return self.Country != nil
}
