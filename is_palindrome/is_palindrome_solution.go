package is_palindrome

import (
	"regexp"
	"strings"
)

func isPalindromeSolution(input string) bool {
	// unify case
	lowered := strings.ToLower(input)

	// strip whitespace & punctuation
	puncs := regexp.MustCompile(`[.,%_:{}*=!-\/'? ]`)
	stripped := puncs.ReplaceAll([]byte(lowered), []byte(""))

	// check if palindrome
	for i := range stripped {
		if stripped[i] != stripped[len(stripped)-1-i] {
			return false
		}
	}

	return true
}
