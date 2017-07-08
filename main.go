// Task_1_split_files project main.go
package main

import (
	"Task_1_split_files/algsolvers"
	"Task_1_split_files/arrreader"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var inputArray, _ = arrreader.ReadArray()
	fmt.Println("Result of B: ", algsolvers.Solver178B(inputArray))
	fmt.Println("Result of C: ", algsolvers.Solver178C(inputArray))
	var n int
	fmt.Print("Enter maximal number: ")
	fmt.Scan(&n)
	algsolvers.Solver554(n)
}
