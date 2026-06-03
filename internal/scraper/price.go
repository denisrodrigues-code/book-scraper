package scraper

import (
	"strconv"
	"strings"
)

func parsePrice(price string) float64 {

	price = strings.ReplaceAll(price, "£", "")

	value, err := strconv.ParseFloat(price, 64)

	if err != nil {
		return 0
	}

	return value
}
