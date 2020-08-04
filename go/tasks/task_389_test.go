package tasks

import "testing"

func TestFindDifference(t *testing.T) {
	var testCases = []struct {
		s        string
		t        string
		expected byte
	}{
		{
			s:        "abcd",
			t:        "abcde",
			expected: 'e',
		},
		{
			s:        "abcd",
			t:        "bcatd",
			expected: 't',
		},
		{
			s:        "abcdabbde",
			t:        "bddbbezcaa",
			expected: 'z',
		},
		{
			s:        "a",
			t:        "aa",
			expected: 'a',
		},
	}

	for _, testCase := range testCases {
		result := findTheDifference(testCase.s, testCase.t)
		if result != testCase.expected {
			t.Errorf("[%s, %s]: expected %v, got %v", testCase.s, testCase.t, testCase.expected, result)
		}
	}
}
