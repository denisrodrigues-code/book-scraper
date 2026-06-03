package scraper

import "strings"

func ratingToNumber(class string) int {

	ratings := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
	}

	for key, value := range ratings {
		if strings.Contains(class, key) {
			return value
		}
	}

	return 0
}
