package main

import "fmt"

func main() {
	x := myFunc()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func myFunc() func() int {
	x := 0

	return func() int {
			x++
			return x
	}
}