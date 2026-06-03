package scraper

import "strings"

const baseURL = "https://books.toscrape.com/"

func buildImageURL(path string) string {
	path = strings.Replace(path, "../../", "", 1)
	return baseURL + path
}

func buildProductURL(path string) string {
	return baseURL + path
}
