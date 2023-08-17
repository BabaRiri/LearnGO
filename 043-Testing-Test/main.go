package main

import "fmt"

func main() {
	fmt.Printf("1 + 1 = %v\n", mySum(1, 1))
	fmt.Printf("2 + 2 = %v\n", mySum(2, 2))
	fmt.Printf("3 + 3 = %v\n", mySum(3, 3))
	fmt.Printf("4 + 4 = %v\n", mySum(4, 4))
	fmt.Printf("5 + 5 = %v\n", mySum(5, 5))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}