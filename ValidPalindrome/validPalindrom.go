package solutions

import "strings"

func isPalindrome(s string) bool {
	leftIndex := 0
	rightIndex := len(s) - 1
	s = strings.ToLower(s)
	for leftIndex < rightIndex {
		leftChar := s[leftIndex]
		rightChar := s[rightIndex]
		if (leftChar < 'a' || leftChar > 'z') && (leftChar < 'A' || leftChar > 'Z') && (leftChar < '0' || leftChar > '9') {
			leftIndex++
			continue
		}
		if (rightChar < 'a' || rightChar > 'z') && (rightChar < 'A' || rightChar > 'Z') && (rightChar < '0' || rightChar > '9') {
			rightIndex--
			continue
		}
		if s[leftIndex] != s[rightIndex] {
			return false
		}
		leftIndex++
		rightIndex--
	}

	return true
}