package main

import "fmt"

func main() {
	x := []int { 1, 2, 3, 4}

	sumFoo := foo(x)
	sumBar := bar(x...)

    fmt.Printf("Sum(foo): %v\n", sumFoo)
	fmt.Printf("Sum(bar): %v\n", sumBar)
}

func foo(mySlc []int) int {
	sum := 0
	for _, v := range mySlc {
		sum += v
	}
	return sum
}

func bar(mySlc ...int) int {
	sum := 0
	for _, v := range mySlc {
		sum += v
	}
	return sum
}