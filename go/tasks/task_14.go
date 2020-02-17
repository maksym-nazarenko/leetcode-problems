package tasks

import "strings"

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".

func longestCommonPrefix(strs []string) string {
	var index int
	var prefix strings.Builder

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	for {
		if len(strs[0]) == 0 {
			return ""
		}

		if index >= len(strs[0]) {
			return prefix.String()
		}

		var curChar byte = strs[0][index]

		for _, str := range strs[1:] {
			if len(str) == 0 {
				return ""
			}

			if index >= len(str) || curChar != str[index] {
				return prefix.String()
			}
		}

		prefix.WriteByte(curChar)
		index++
	}

}
