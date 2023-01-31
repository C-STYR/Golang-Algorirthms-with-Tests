package linear_search

import "testing"

func TestLinearSearch(t *testing.T){
	testSlice := []int{1, 2, 3, 4, 5}

	result1 := linearSearch(testSlice, 2)
	result2 := linearSearch(testSlice, 6)

	if result1 != true {
		t.Error("got false when slice contains target")
	}

	if result2 == true {
		t.Error("got true when slice does not contain target")
	}
}