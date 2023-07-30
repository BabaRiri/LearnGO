package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var x int
	for i := 1; i <= 42; i++ {
		x = rand.Intn(5)
		fmt.Printf("i = %d\t", i)

		switch x{
			case 0:
				fmt.Printf("x = %v\n", x)
			case 1:
				fmt.Printf("x = %v\n", x)
			case 2:
				fmt.Printf("x = %v\n", x)
			case 3:
				fmt.Printf("x = %v\n", x)
			case 4:
				fmt.Printf("x = %v\n", x)
			default:
				fmt.Println("Not possible")
		}
	}
}
