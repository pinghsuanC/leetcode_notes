package solutions

func isValid(s string) bool {
	stringlen := len(s)
	helperMap := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	if stringlen == 0 {
		return true
	}
	if stringlen%2 != 0 {
		return false
	}

	stack := make([]rune, 0)
	for _, char := range s {
		if match, ok := helperMap[char]; !ok {
			// push left column onto stack
			stack = append(stack, char)
		} else {
			// if there is a right column without a left one, return false
			if len(stack) < 1 {
				return false
			}
			preVal := stack[len(stack)-1]
			if preVal == match {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

	}

	return len(stack) == 0
}