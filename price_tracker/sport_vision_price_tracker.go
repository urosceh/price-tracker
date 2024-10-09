package price_tracker

type SportVisionPriceTracker struct {
	*AbstractPriceTracker
}

/*
 * Constructor
 * @param itemName string -> name of the item whose price is to be tracked
 * @param result *TemuPriceTracker -> pointer to the result object
*/
func NewSportVisionPriceTracker(itemName string, url string) *SportVisionPriceTracker {
	abstractPriceTracker := NewAbstractPriceTracker(itemName, url, []string{".product-price-value.value"})

	return &SportVisionPriceTracker{
		AbstractPriceTracker: abstractPriceTracker,
	}
}