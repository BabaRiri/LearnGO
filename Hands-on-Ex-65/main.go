package main

import "fmt"

func main() {
    ans := multiply(5, 10)
	fmt.Println(ans)
}

func multiply(a, b float64) float64 {
	return a * b
}