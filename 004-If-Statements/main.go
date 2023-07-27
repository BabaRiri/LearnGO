package main

import "fmt"

func main() {
	x := 23

	if x == 0 {
		fmt.Println("x is zero")
	}

	if x > 0 {
		fmt.Println("x is positive")
	} else {
		fmt.Println("x is negative")
	}

	if x < 10 {
		fmt.Println("x is less than 10")
	} else if x < 20 {
		fmt.Println("x is greate than 10 but less than 20")
	} else {
		fmt.Println("x is greater than 20")
	}
}