// add list of all items
// create all factories
// createPriceTracker(type string) { result = getWebsite; getAbstractFactory(string, result) }
// pasteToOutputFile(pt *PriceTracker) { paste to pt.getOutputFile() }
package main

import (
	"fmt"
	"price-tracker/price_tracker"
)

func main() {
	price_tracker1 := price_tracker.GetPriceTracker("Asics jakna", "https://www.sportvision.rs/jakna/70013410-asics-icon-jacket")

	price_tracker2 := price_tracker.GetPriceTracker("Adidas Soulstride", "https://planetasport.rs/patike-terrex-soulstride-r-rdy-m-ig8029.html")

	fmt.Println(price_tracker1)
	fmt.Println(price_tracker2)	
}