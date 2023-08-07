package main

import "fmt"

func addI( a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

// Type constraints
func addT[T int | float64](a, b T) T {
	return a + b
}

type myNum interface {
	~int | ~float64
}

func addT2[T myNum](a, b T) T {
	return a + b
}

type myAlias int

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
	fmt.Printf("GENERIC FUNC [float64]: %v\n\n", ansT2)

	ansT3 := addT2(a, b)
	ansT4 := addT2(c, d)
	fmt.Printf("(type set interface) GENERIC FUNC [int]: %v\n", ansT3)
	fmt.Printf("(type set interface) GENERIC FUNC [float64]: %v\n", ansT4)

	var x myAlias = 15
	ansT5 := addT2(x, 2)
	fmt.Printf("(Undelying types) GENERIC FUNC [int]: %v\n", ansT5)
}