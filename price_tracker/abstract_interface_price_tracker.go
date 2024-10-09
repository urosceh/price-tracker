package price_tracker

import (
	"fmt"
	"regexp"
)

type IPriceTracker interface {
	GetName() string
	GetPrice() float64
	GetUrl() string
}

func GetPriceTracker(name string, url string) (IPriceTracker, error) {
	hostNameRegex := regexp.MustCompile(`(?:https?:\/\/)?(www\.)?([a-zA-Z0-9-]+)\.(com|rs)`)

	hostName := hostNameRegex.FindStringSubmatch(url)[2]

	switch hostName {
	case "sportvision":
		return NewSportVisionPriceTracker(name, url), nil
	case "planetasport":
		return NewPlanetaSportPriceTracker(name, url), nil
	default:
		return nil, fmt.Errorf("unsupported website: %s", url)
	}
}