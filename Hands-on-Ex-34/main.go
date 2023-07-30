package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("I am in MAIN loop...\tMAIN INTERATION: %d\n", i)
		for b := 0; b < 5; b++ {
			fmt.Printf("I am in NESTED loop...\tNESTED ITERATION: %d\n", b)
		}
	}
}
