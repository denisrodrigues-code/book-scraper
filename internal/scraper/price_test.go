package scraper

import "testing"

func TestParsePrice(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"£10.99", 10.99},
		{"£50.00", 50.00},
		{"£0.99", 0.99},
	}

	for _, test := range tests {
		result := parsePrice(test.input)

		if result != test.expected {
			t.Errorf(
				"expected %.2f, got %.2f",
				test.expected,
				result,
			)
		}
	}
}
