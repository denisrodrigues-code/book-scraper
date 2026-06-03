package scraper

import (
	"fmt"
	"time"

	"book-scraper/internal/models"
)

func ScrapeBooks() ([]models.Book, error) {

	var allBooks []models.Book

	for page := 1; page <= 50; page++ {

		var url string

		if page == 1 {
			url = baseURL
		} else {
			url = fmt.Sprintf(
				"%scatalogue/page-%d.html",
				baseURL,
				page,
			)
		}

		books, err := ScrapePage(url)

		if err != nil {
			return nil, err
		}

		allBooks = append(allBooks, books...)

		time.Sleep(1 * time.Second)
	}

	return allBooks, nil
}
