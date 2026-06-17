package scraper

import "net/url"

const baseURL = "https://books.toscrape.com/"

func buildURL(path string) string {
	base, _ := url.Parse(baseURL)
	ref, _ := url.Parse(path)

	return base.ResolveReference(ref).String()
}
