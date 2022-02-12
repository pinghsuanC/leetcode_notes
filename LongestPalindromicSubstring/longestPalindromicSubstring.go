package solutions

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	result := ""
	maxLen := 0

	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			subStr := s[i:j]
			if isPa(subStr) {
				curLen := len(subStr)
				if maxLen < curLen {
					maxLen = curLen
					result = subStr
				}
			}
		}
	}
	return result
}

func isPa(s string) bool {
	var n int
	if len(s)%2 == 0 {
		n = len(s) / 2
	} else {
		n = len(s)/2 + 1
	}
	for i := 0; i < n; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}