package price_tracker

import "regexp"

type IPriceTracker interface {
	GetName() string
	GetPrice() float64
	GetUrl() string
}

func GetPriceTracker(name string, url string) IPriceTracker {
	hostNameRegex := regexp.MustCompile(`(?:https?:\/\/)?(www\.)?([a-zA-Z0-9-]+)\.(com|rs)`)

	hostName := hostNameRegex.FindStringSubmatch(url)[2]

	switch hostName {
	case "sportvision":
		return NewSportVisionPriceTracker(name, url)
	case "planetasport":
		return NewPlanetaSportPriceTracker(name, url)
	default:
		return nil
	}
}