package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization of the program begins.")
}

func main() {
    x := rand.Intn(251)

	fmt.Printf("x = %v\n", x)

	switch {
	case (x >= 0 && x <= 100):
		fmt.Printf("x is between 0 and 100")
	case (x > 100 && x <= 200):
		fmt.Printf("x is between 101 and 200")
	default:
		fmt.Printf("x is not between 201 and 250")
	}
}