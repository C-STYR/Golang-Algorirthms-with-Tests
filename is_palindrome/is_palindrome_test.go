package is_palindrome

import "testing"

type palindromeTest struct {
	arg1     string
	expected bool
}

var palindromeTests = []palindromeTest{
	{"llama mall", true},
	{"acrobatic coitus", false},
	{"Anne, I vote more   cars race Rome-to-Vienna", true},
	{" If I had a HI-FI", true},
	{"Please go the *fuck* away", false},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range palindromeTests {
		if output := isPalindrome(test.arg1); output != test.expected {
			t.Errorf("Result %s not equal to expected %t", test.arg1, test.expected)
		}
	}
}
