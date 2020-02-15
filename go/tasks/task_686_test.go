package tasks

import "testing"

func TestRepeatedStringMatch(t *testing.T) {
	v := repeatedStringMatch("abcd", "cdabcdab")
	if v != 3 {
		t.Errorf("Wanted: 3, got %d\n", v)
	}

	v = repeatedStringMatch("a", "aa")
	if v != 2 {
		t.Errorf("Wanted: 2, got %d\n", v)
	}

	v = repeatedStringMatch("a", "a")
	if v != 1 {
		t.Errorf("Wanted: 1, got %d\n", v)
	}

	v = repeatedStringMatch("", "a")
	if v != 1 {
		t.Errorf("Wanted: 0, got %d\n", v)
	}

	v = repeatedStringMatch("a", "")
	if v != 1 {
		t.Errorf("Wanted: 0, got %d\n", v)
	}

	v = repeatedStringMatch("aa", "a")
	if v != 1 {
		t.Errorf("Wanted: 1, got %d\n", v)
	}

	v = repeatedStringMatch("abc", "cabcabca")
	if v != 4 {
		t.Errorf("Wanted: 4, got %d\n", v)
	}

	v = repeatedStringMatch("aaaaaaaaaaaaaaaaaaaaaab", "ba")
	if v != 2 {
		t.Errorf("Wanted: 2, got %d\n", v)
	}

	v = repeatedStringMatch("abcde", "cea")
	if v != -1 {
		t.Errorf("Wanted: -1, got %d\n", v)
	}
}
