package scraper

import (
	"book-scraper/internal/models"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeBooks() ([]models.Book, error) {

	var books []models.Book

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(
		"GET",
		"https://books.toscrape.com/",
		nil,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"User-Agent",
		"BookScraper/1.0",
	)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return nil, err
	}

	doc.Find("article.product_pod").Each(func(i int, s *goquery.Selection) {

		title, _ := s.Find("h3 a").Attr("title")

		priceText := s.Find(".price_color").Text()

		ratingClass, _ := s.Find(".star-rating").Attr("class")

		availability := strings.TrimSpace(
			s.Find(".availability").Text(),
		)

		imageURL, _ := s.Find("img").Attr("src")

		productURL, _ := s.Find("h3 a").Attr("href")

		book := models.Book{
			Title:        title,
			Price:        parsePrice(priceText),
			Rating:       ratingToNumber(ratingClass),
			Availability: availability,
			ImageURL:     buildImageURL(imageURL),
			ProductURL:   buildProductURL(productURL),
		}

		books = append(books, book)
	})

	return books, nil

}
