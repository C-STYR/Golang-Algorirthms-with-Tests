package nth_fibonacci

func nthFibonacciSolution(n int) int {
	var recursiveCheck func(n int) int
	cache := make(map[int]int)

	recursiveCheck = func(n int) int {
		if val, ok := cache[n]; ok {
			return val
		}
		if n < 2 {
			return n
		} else {
			result := (recursiveCheck(n-1) + recursiveCheck(n-2))
			cache[n] = result
			return result
		}
	}
	return recursiveCheck(n)
}
