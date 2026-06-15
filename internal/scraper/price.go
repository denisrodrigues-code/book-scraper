package scraper

import (
	"strconv"
	"strings"
)

func parsePrice(price string) (float64, error) {

	price = strings.ReplaceAll(price, "£", "")

	value, err := strconv.ParseFloat(price, 64)

	if err != nil {
		return 0, err
	}

	return value, nil
}
