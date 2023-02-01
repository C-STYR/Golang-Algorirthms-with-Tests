package is_prime

import "testing"

type isPrimeTest struct {
	arg int
	expected bool
}

var isPrimeTests = []isPrimeTest{
	{0, false},
	{1, false},
	{2, true},
	{3, true},
	{169, false},
	{256, false},
	{953, true},
	{927, false},
}

func TestIsPrime(t *testing.T) {
	for _, test := range isPrimeTests {
		if output := isPrime(test.arg); output != test.expected {
			t.Errorf("For input %d, got result %t but expected %t", test.arg, output, test.expected)
		}
	}
}
func TestIsPrimeSolution(t *testing.T) {
	for _, test := range isPrimeTests {
		if output := isPrimeSolution(test.arg); output != test.expected {
			t.Errorf("For input %d, got result %t but expected %t", test.arg, output, test.expected)
		}
	}
}