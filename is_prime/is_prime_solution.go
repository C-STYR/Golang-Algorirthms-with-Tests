package is_prime

import (
	"math"
)

func isPrimeSolution(input int) bool {
	// there is no reason to test further than the sqrt of the input
	limit := int(math.Sqrt(float64(input)))
	for i := 2; i <= limit; i++ {
		if input % i == 0 {
			return false}
	}
	return input > 1
}