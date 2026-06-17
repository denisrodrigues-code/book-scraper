package scraper

import (
	"fmt"
	"sync"
	"time"

	"book-scraper/internal/models"
)

const (
	totalPages   = 50
	workerCount  = 5
	requestDelay = 300 * time.Millisecond
)

func ScrapeBooks() ([]models.Book, error) {
	jobs := make(chan string)
	results := make(chan []models.Book)
	errors := make(chan error, 1)

	var wg sync.WaitGroup

	limiter := time.NewTicker(requestDelay)
	defer limiter.Stop()

	for i := 0; i < workerCount; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for url := range jobs {
				<-limiter.C

				books, err := ScrapePage(url)
				if err != nil {
					select {
					case errors <- err:
					default:
					}
					return
				}

				results <- books
			}
		}()
	}

	go func() {
		defer close(jobs)

		for page := 1; page <= totalPages; page++ {
			jobs <- buildPageURL(page)
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var allBooks []models.Book

	for {
		select {
		case err := <-errors:
			return nil, err

		case books, ok := <-results:
			if !ok {
				return allBooks, nil
			}

			allBooks = append(allBooks, books...)
		}
	}
}

func buildPageURL(page int) string {
	if page == 1 {
		return baseURL
	}

	return fmt.Sprintf(
		"%scatalogue/page-%d.html",
		baseURL,
		page,
	)
}
