package tasks

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	var v string

	v = longestCommonPrefix([]string{"flower", "flight", "fleet"})
	if v != "fl" {
		t.Errorf("Wanted <%s>, got <%s>\n", "fl", v)
	}

	v = longestCommonPrefix([]string{"flower", "light", "etc"})
	if v != "" {
		t.Errorf("Wanted <%s>, got <%s>\n", "", v)
	}

	v = longestCommonPrefix([]string{"flower", "flow", "flowers are cool"})
	if v != "flow" {
		t.Errorf("Wanted <%s>, got <%s>\n", "flow", v)
	}

	v = longestCommonPrefix([]string{"", ""})
	if v != "" {
		t.Errorf("Wanted <%s>, got <%s>\n", "", v)
	}

	v = longestCommonPrefix([]string{"c", "c"})
	if v != "c" {
		t.Errorf("Wanted <%s>, got <%s>\n", "c", v)
	}
}
