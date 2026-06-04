package scraper

import "testing"

func TestRatingToNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"One", 1},
		{"Two", 2},
		{"Three", 3},
		{"Four", 4},
		{"Five", 5},
	}

	for _, test := range tests {
		result := ratingToNumber(test.input)

		if result != test.expected {
			t.Errorf(
				"expected %d, got %d",
				test.expected,
				result,
			)
		}
	}
}
