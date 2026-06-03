package scraper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeTitles() error {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(
		"GET",
		"http://books.toscrape.com/",
		nil,
	)

	if err != nil {
		return err
	}

	req.Header.Set(
		"User-Agent",
		"BookScraper/1.0",
	)

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return err
	}

	doc.Find("article.product_pod").Each(func(i int, s *goquery.Selection) {

		title, exists := s.Find("h3 a").Attr("title")

		if exists {
			fmt.Println(title)
		}

	})

	return nil

}
