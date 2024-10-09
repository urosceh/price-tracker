// add list of all items
// create all factories
// createPriceTracker(type string) { result = getWebsite; getAbstractFactory(string, result) }
// pasteToOutputFile(pt *PriceTracker) { paste to pt.getOutputFile() }
package main

import (
	"encoding/json"
	"fmt"
	"os"

	interfaces "price-tracker/interface"
	price_tracker "price-tracker/price_tracker"
)

func main() {
	// get file linkes_for_scrape.json from root folder
	// get all items from file
	// create all factories
	file, err := os.Open("items_for_scrape.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1);
	}

	defer file.Close()

	var products []interfaces.ItemForScrape

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&products)
	if err != nil {
			fmt.Println("Error decoding JSON:", err)
			os.Exit(1)
	}

	for _, product := range products {
		priceTracker, err := price_tracker.GetPriceTracker(product.Name, product.Url)
		if err != nil {
			fmt.Println("Failed to get price tracker for", product.Name, ":", err)
			continue
		}
		fmt.Println(priceTracker)
	}
}