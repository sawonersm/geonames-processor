package enum

type Mode string

const (
	FEATURES Mode = "features"
	COUNTRY  Mode = "country"
)

func FromString(mode string) Mode {
	switch mode {
	case "features":
		return FEATURES
	case "country":
		return COUNTRY
	}

	panic("Mode unavailable")
}
