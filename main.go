// add list of all items
// create all factories
// createPriceTracker(type string) { result = getWebsite; getAbstractFactory(string, result) }
// pasteToOutputFile(pt *PriceTracker) { paste to pt.getOutputFile() }
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	interfaces "price-tracker/interface"
	price_tracker "price-tracker/price_tracker"
)

func main() {
	filePath := "prices.json"

	itemsForScrape := getItemsForScrape()

	scrapedItemEntriesMap := getScrapedItemEntriesMap(filePath, itemsForScrape)

	writeToFile(filePath, scrapedItemEntriesMap)
}

func getItemsForScrape() []interfaces.ItemForScrape {
	file, err := os.Open("items_for_scrape.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1);
	}

	defer file.Close()

	var itemsForScrape []interfaces.ItemForScrape

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&itemsForScrape)
	if err != nil {
			fmt.Println("Error decoding JSON:", err)
			os.Exit(1)
	}

	return itemsForScrape
}

func getScrapedItemEntriesMap(filePath string, itemsForScrape []interfaces.ItemForScrape) map[string][]interfaces.ScrapedItemEntry {
	scrapedItemEntriesMap := loadScrapedItemEntriesFromFile(filePath)

	priceTrackers := price_tracker.GetPriceTrackers(itemsForScrape)

	for _, priceTracker := range priceTrackers {
		length := len(scrapedItemEntriesMap[priceTracker.GetUrl()])

		if length >= 10 {
			scrapedItemEntriesMap[priceTracker.GetUrl()] = scrapedItemEntriesMap[priceTracker.GetUrl()][1:]
		}

		if length == 0 || scrapedItemEntriesMap[priceTracker.GetUrl()][length-1].Price != priceTracker.GetPrice() {
			scrapedItemEntriesMap[priceTracker.GetUrl()] = append(scrapedItemEntriesMap[priceTracker.GetUrl()], interfaces.ScrapedItemEntry{
				ScrapedAt: time.Now().Format(time.RFC3339),
				Price:     priceTracker.GetPrice(),
			})
		} else {
			scrapedItemEntriesMap[priceTracker.GetUrl()][len(scrapedItemEntriesMap[priceTracker.GetUrl()])-1].ScrapedAt = time.Now().Format(time.RFC3339)
		}
		
		fmt.Println(priceTracker)
	}

	return scrapedItemEntriesMap
}

func loadScrapedItemEntriesFromFile(filePath string) map[string][]interfaces.ScrapedItemEntry {
	scrapedItemEntryMap := map[string][]interfaces.ScrapedItemEntry{}

	if _, err := os.Stat(filePath); err == nil {
		fileData, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			panic (err)
		}

		err = json.Unmarshal(fileData, &scrapedItemEntryMap)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			panic (err)
		}
	}

	return scrapedItemEntryMap
}

func writeToFile(filePath string, items map[string][]interfaces.ScrapedItemEntry) {
	for itemName, priceEntries := range items {
		fmt.Printf("Item: %s\n", itemName)
		for _, entry := range priceEntries {
			fmt.Printf("  Scraped at: %s, Price: %.2f\n", entry.ScrapedAt, entry.Price)
		}
	}

	newFileData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile(filePath, newFileData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}