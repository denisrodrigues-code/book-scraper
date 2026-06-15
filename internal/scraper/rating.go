package scraper

import "strings"

func ratingToNumber(class string) int {
	switch {
	case strings.Contains(class, "One"):
		return 1
	case strings.Contains(class, "Two"):
		return 2
	case strings.Contains(class, "Three"):
		return 3
	case strings.Contains(class, "Four"):
		return 4
	case strings.Contains(class, "Five"):
		return 5
	default:
		return 0
	}
}
