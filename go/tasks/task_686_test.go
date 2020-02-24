package tasks

import "testing"

func TestRepeatedStringMatch(t *testing.T) {
	var testCases = []struct {
		inString      string
		desiredString string
		wanted        int
	}{
		{
			inString:      "abcd",
			desiredString: "cdabcdab",
			wanted:        3,
		},
		{
			inString:      "a",
			desiredString: "aa",
			wanted:        2,
		},
		{
			inString:      "a",
			desiredString: "a",
			wanted:        1,
		},
		{
			inString:      "",
			desiredString: "a",
			wanted:        1,
		},
		{
			inString:      "a",
			desiredString: "",
			wanted:        1,
		},
		{
			inString:      "aa",
			desiredString: "a",
			wanted:        1,
		},
		{
			inString:      "abc",
			desiredString: "cabcabca",
			wanted:        4,
		},
		{
			inString:      "aaaaaaaaaaaaaaaaaaaaaab",
			desiredString: "ba",
			wanted:        2,
		},
		{
			inString:      "abcde",
			desiredString: "cea",
			wanted:        -1,
		},
	}

	for _, currentCase := range testCases {
		v := repeatedStringMatch(currentCase.inString, currentCase.desiredString)
		if v != currentCase.wanted {
			t.Errorf("Wanted: %d, got %d\n", currentCase.wanted, v)
		}
	}
}
