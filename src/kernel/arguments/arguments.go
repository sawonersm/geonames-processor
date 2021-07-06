package parameters

import (
	"os"
	Mode "sawonersm/geonames-processor/kernel/arguments/enum/mode"
)

type Arguments struct {
	Mode Mode.Mode
}

func New() *Arguments {
	arguments := os.Args[1:]
	return &Arguments{
		Mode.FromString(
			arguments[0],
		),
	}
}
