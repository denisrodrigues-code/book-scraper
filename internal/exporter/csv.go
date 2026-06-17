package exporter

import (
	"encoding/csv"
	"os"
	"strconv"

	"book-scraper/internal/models"
)

func SaveCSV(books []models.Book) error {

	file, err := os.Create("data/books.csv")

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	err = writer.Write([]string{
		"title",
		"price",
		"rating",
		"availability",
		"image_url",
		"product_url",
	})

	if err != nil {
		return err
	}

	for _, book := range books {

		err = writer.Write([]string{
			book.Title,
			strconv.FormatFloat(
				book.Price,
				'f',
				2,
				64,
			),
			strconv.Itoa(book.Rating),
			book.Availability,
			book.ImageURL,
			book.ProductURL,
		})

		if err != nil {
			return err
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
