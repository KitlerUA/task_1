// Solver178C
package algsolvers

import (
	"math"
)

func Solver178C(input []int) int {
	var result int = 0
	for i := range input {
		if input[i]%2 == 0 {
			var sqrt float64 = math.Sqrt(float64(input[i]))
			if math.Abs(sqrt-math.Trunc(sqrt+0.5)) < 0.00001 {
				result++
			}
		}
	}
	return result
}
