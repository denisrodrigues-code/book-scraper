package exporter

import (
	"encoding/json"
	"os"

	"book-scraper/internal/models"
)

func SaveJSON(books []models.Book) error {

	data, err := json.MarshalIndent(
		books,
		"",
		" ",
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		"data/books.json",
		data,
		0644,
	)
}
