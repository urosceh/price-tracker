package price_tracker

type PlanetaSportPriceTracker struct {
	*AbstractPriceTracker
}

/*
 * Constructor
 * @param itemName string -> name of the item whose price is to be tracked
 * @param result *TemuPriceTracker -> pointer to the result object
*/
func NewPlanetaSportPriceTracker(itemName string, url string) *PlanetaSportPriceTracker {
	abstractPriceTracker := NewAbstractPriceTracker(itemName, url, []string{".product-info-price",".zsdev-special-price",".price.tooltip-toggle"})

	return &PlanetaSportPriceTracker{
		AbstractPriceTracker: abstractPriceTracker,
	}
}