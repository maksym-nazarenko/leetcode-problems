package tasks

import "testing"

func TestIsValid(t *testing.T) {
	var results = map[string]bool{
		"":       true,
		"()":     true,
		"()[]{}": true,
		"(]":     false,
		"([)]":   false,
		"{[]}":   true,
	}

	for s, expected := range results {
		actual := isValid(s)
		if expected != actual {
			t.Errorf("<%s>: wanted %v, got %v\n", s, expected, actual)
		}
	}

}
