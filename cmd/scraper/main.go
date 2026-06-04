package main

import (
	"log"

	"book-scraper/internal/exporter"
	"book-scraper/internal/scraper"
)

func main() {

	books, err := scraper.ScrapeBooks()

	if err != nil {
		log.Fatal(err)
	}

	err = exporter.SaveJSON(books)

	if err != nil {
		log.Fatal(err)
	}

	err = exporter.SaveCSV(books)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf(
		"%d books exported successfully",
		len(books),
	)
}
