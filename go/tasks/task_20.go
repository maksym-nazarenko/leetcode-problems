package tasks

func isValid(s string) bool {
	var stack []byte
	var brackets = map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}

		if s[i] == ')' || s[i] == '}' || s[i] == ']' {
			if len(stack) == 0 {
				return false
			}

			if stack[len(stack)-1] != brackets[s[i]] {
				return false
			}

			stack = stack[:len(stack)-1]
		}

	}

	return len(stack) == 0
}
