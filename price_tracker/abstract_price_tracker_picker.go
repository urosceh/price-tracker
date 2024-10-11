package price_tracker

import (
	"fmt"
	interfaces "price-tracker/interface"
	"regexp"
	"sync"
	"time"
)

type PriceTrackerResponse struct {
	PriceTracker interfaces.IPriceTracker
	Error        error
}

func GetPriceTrackers(itemsForScrape []interfaces.ItemForScrape) []interfaces.IPriceTracker {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(itemsForScrape))

	channel := make(chan PriceTrackerResponse, len(itemsForScrape))

	for _, item := range itemsForScrape {
		go getPriceTracker(item, &waitGroup, channel)
	}

	waitGroup.Wait()
	close(channel)

	var priceTrackers []interfaces.IPriceTracker

	for response := range channel {
		if response.Error != nil {
			fmt.Println("Failed to get price tracker for", response.PriceTracker.GetUrl(), ":", response.Error)
			continue
		} else {
			priceTrackers = append(priceTrackers, response.PriceTracker)
		}
	}

	return priceTrackers

}

func getPriceTracker(item interfaces.ItemForScrape, waitGroup *sync.WaitGroup, channel chan PriceTrackerResponse) {
	defer waitGroup.Done()

	startTime := time.Now()
	fmt.Println("Getting price tracker for", item.Name, "at", time.Now())


	hostNameRegex := regexp.MustCompile(`(?:https?:\/\/)?(www\.)?([a-zA-Z0-9-]+)\.(com|rs)`)

	hostName := hostNameRegex.FindStringSubmatch(item.Url)[2]

	switch hostName {
	case "sportvision":
		channel <- PriceTrackerResponse{NewSportVisionPriceTracker(item.Name, item.Url), nil}
	case "planetasport":
		channel <- PriceTrackerResponse{NewPlanetaSportPriceTracker(item.Name, item.Url), nil}
	default:
		channel <- PriceTrackerResponse{nil, fmt.Errorf("unsupported host name: %s", hostName)}
	}

	fmt.Println("Got price tracker for", item.Name, "in", time.Since(startTime))
}