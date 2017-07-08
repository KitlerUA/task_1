// Solver558
package algsolvers

import (
	"fmt"
)

func Solver554(n int) {
	if n >= 5 {
		var arrayOfSquares = make([]int, n)
		for i := 1; i <= n; i++ {
			arrayOfSquares[i-1] = i * i
		}
		for i := n - 1; i >= 2; i-- {
			for j := 2; j < i; j++ {
				for k := j + 1; k < i; k++ {
					if arrayOfSquares[j]+arrayOfSquares[k]-arrayOfSquares[i] == 0 {
						fmt.Print("a = ", j+1, ", b = ", k+1, ", c = ", i+1)
						fmt.Println(" (", arrayOfSquares[j], " + ", arrayOfSquares[k], " = ", arrayOfSquares[i], ")")
					}
				}
			}
		}
	}

}
