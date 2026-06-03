package main

import (
	"log"

	"book-scraper/internal/scraper"
)

func main() {

	err := scraper.ScrapeTitles()

	if err != nil {
		log.Fatal(err)
	}
}
