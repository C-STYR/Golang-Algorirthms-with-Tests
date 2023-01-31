package linear_search

/*
linearSearch should return a bool regarding the presence of a target int in a slice of ints
*/

func linearSearch(collection []int, target int) bool {
	for _, item := range collection {
		if item == target {
			return true
		}
	}
	return false
}