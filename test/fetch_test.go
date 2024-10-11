package price_tracker

import (
	"price-tracker/price_tracker"
	"testing"
)

func TestFetchPriceFromUrl(t *testing.T) {
	t.Run("Test fetch price from planetasport url", func(t *testing.T) {
		planetaPriceTracker, _ := price_tracker.GetPriceTracker("Adidas Soulstride", "https://planetasport.rs/patike-terrex-soulstride-r-rdy-m-ig8029.html")
		if planetaPriceTracker.GetPrice() != 9749.99 {
			t.Errorf("Price is not correct. Expected: 9749.99, got: %.2f", planetaPriceTracker.GetPrice())
		}
	})

	t.Run("Test fetch price from sportvision url", func(t *testing.T) {
		sportVisionPriceTracker, _ := price_tracker.GetPriceTracker("Adidas Soulstride", "https://www.sportvision.rs/patike/69947717-adidas-terrex-soulstride")
		if sportVisionPriceTracker.GetPrice() != 11999.00 {
			t.Errorf("Price is not correct. Expected: 11999.00, got: %.2f", sportVisionPriceTracker.GetPrice())
		}
	})
}