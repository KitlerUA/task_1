// Solver178B
package algsolvers

func Solver178B(arrayOfNumbers []int) int {
	var result int = 0
	for i := range arrayOfNumbers {
		if arrayOfNumbers[i]%3 == 0 && arrayOfNumbers[i]%5 != 0 {
			result++
		}
	}
	return result
}
