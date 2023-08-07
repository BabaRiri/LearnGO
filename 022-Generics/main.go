package main

import "fmt"

func addI( a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {
	a := 2
	b := 4
	c := 1.2
	d := 3.5

	ansI := addI(a, b)
	ansF := addF(c, d)
	ansT1 := addT(a, b)
	ansT2 := addT(c, d)

	fmt.Printf("INT FUNC: %v\n", ansI)
	fmt.Printf("FLOAT64 FUNC: %v\n\n", ansF)
	fmt.Printf("GENERIC FUNC [int]: %v\n", ansT1)
	fmt.Printf("GENERIC FUNC [float64]: %v\n", ansT2)
	
}