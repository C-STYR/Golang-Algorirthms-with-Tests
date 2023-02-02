package nth_fibonacci

import "testing"

type nthFibTest struct {
	arg int
	expected int
}

var nthFibTests = []nthFibTest{
	{0, 0},
	{1, 1},
	{2, 1},
	{11, 89},
	{15, 610},
	{19, 4181},
}

func TestNthFibonacci(t *testing.T) {
	for _, test := range nthFibTests {
		if result := nthFibonacci(test.arg); result != test.expected {
			t.Errorf("For n %d, got result %d but expected %d", test.arg, result, test.expected)
		}
	}
}
func TestNthFibonacciSolution(t *testing.T) {
	for _, test := range nthFibTests {
		if result := nthFibonacciSolution(test.arg); result != test.expected {
			t.Errorf("For n %d, got result %d but expected %d", test.arg, result, test.expected)
		}
	}
}