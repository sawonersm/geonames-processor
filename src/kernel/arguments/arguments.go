package parameters

import (
	"os"
	Mode "sawonersm/geonames-processor/kernel/arguments/enum/mode"
)

type Arguments struct {
	Mode    Mode.Mode
	Country *string
}

func New() *Arguments {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		panic("You must specify the mode: features|country")
	}

	mode := Mode.FromString(arguments[0])

	var country *string = nil
	if mode == Mode.COUNTRY {
		if len(arguments) < 2 {
			panic("You must specify the country. Countries availables: ES")
		}
		country = &arguments[1]
	}

	return &Arguments{
		Mode:    mode,
		Country: country,
	}
}
