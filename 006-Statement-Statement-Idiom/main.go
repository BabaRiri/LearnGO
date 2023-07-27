package main

import "fmt"

func main() {
    x := 10

	if z := x%2; z != 0 {
		fmt.Printf("x = (%v) is odd", x)
	} else {
		fmt.Printf("x = (%v) is even", x)
	}
}