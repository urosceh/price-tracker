package helper

import (
	"regexp"
	"strconv"
	"strings"
)

func ParsePriceString(priceString string) float64 {
	value := strings.ReplaceAll(strings.ReplaceAll(priceString, ".", ""), ",", ".")

	whitespaceAndLettersRegex := regexp.MustCompile(`[a-zA-Z\s\n]`)

	price, err := strconv.ParseFloat(whitespaceAndLettersRegex.ReplaceAllString(value, ""), 64)
	if err != nil {
		panic(err)
	}

	return price
}