package price_tracker

import (
	"fmt"
	"net/http"
	"strconv"

	"price-tracker/helper"

	"github.com/PuerkitoBio/goquery"
)

type AbstractPriceTracker struct {
	name string
	price float64
	url string
}

func NewAbstractPriceTracker(itemName string, url string, priceFields []string) *AbstractPriceTracker {
	price := generatePrice(url, priceFields)

	return &AbstractPriceTracker{
		name: itemName,
		price: price,
		url: url,
	}
}

func (pt *AbstractPriceTracker) GetPrice() float64 {
	return pt.price
}

func (pt *AbstractPriceTracker) GetUrl() string {
	return pt.url
}

func (pt *AbstractPriceTracker) String() string {
	return fmt.Sprintf("Name: %s\nPrice: %.2f\n", pt.name, pt.price)
}

func generatePrice(url string, priceFields []string) float64 {
	result, err := http.Get(url)

	if err != nil || result.StatusCode >= 300 {
		panic(err)
	} else if result.StatusCode >= 300 {
		panic("Status code: " + strconv.Itoa(result.StatusCode))
	}
	defer result.Body.Close()

	doc, err := goquery.NewDocumentFromReader(result.Body)
	if err != nil {
		fmt.Println("Error loading HTTP response body. ", err)
	}

	var priceNode *goquery.Selection

	for index, priceField := range priceFields {
		if (index == 0) {
			priceNode = doc.Find(priceField)
		} else {
			priceNode = priceNode.Find(priceField)
		}
	}

	return helper.ParsePriceString(priceNode.Text())
}