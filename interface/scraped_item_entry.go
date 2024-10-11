package interfaces

type ScrapedItemEntry struct {
	ScrapedAt 	string		`json:"scraped_at"`
	Price 			float64		`json:"price"`
}