package main

import "fmt"

func main() {
	x := 23

	if x > 0 && x <= 10 {
		fmt.Println("x is less than 10")
	} else if x >=10 && x <= 20 {
		fmt.Println("x is greater than 10 but less than 20")
	} else {
		fmt.Println("x is greater than 20")
	}

	if x == 0 || x == 1 {
		fmt.Println("x is 0 or 1")
	}

	if x != 0 && x != 1 {
        fmt.Println("x is neither 0 nor 1")
    }
}