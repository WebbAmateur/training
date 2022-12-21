package palindrome

import "strings"

func IsPalindrome(candidate string) bool {

	candidate = strings.ToLower(candidate)
	candidate = strings.ReplaceAll(candidate, " ", "")
	for front, back := 0, len(candidate)-1; front <= back; front, back = front+1, back-1 {
		if candidate[front] != candidate[back] {
			return false
		}
	}

	// We have checked all characters in candidate

	return true

}

// len(candidate) = 10
// 1  front=0, back=9
// 2  front=1, back=8
