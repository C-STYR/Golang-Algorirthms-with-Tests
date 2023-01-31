package is_palindrome

import "testing"

type palindromeTest struct {
	arg1     string
	expected bool
}

var palindromeTests = []palindromeTest{
	palindromeTest{"llama mall", true},
	palindromeTest{"acrobatic coitus", false},
	palindromeTest{"Anne, I vote more cars race Rome-to-Vienna", true},
	palindromeTest{"If I had a HI-FI", true},
	palindromeTest{"Please go the fuck away", true},
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range palindromeTests {
		if output := isPalindrome(test.arg1); output != test.expected {
			t.Errorf("Result %q not equal t expected %q", output, test.expected)
		}
	}
}
