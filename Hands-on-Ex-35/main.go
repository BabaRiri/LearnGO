package main

import "fmt"

func main() {
	xi := []int{41, 42, 43, 44, 55, 46, 47, 48, 49}

	for i, v := range xi {
		fmt.Printf("INDEX: %d\tVALUE: %d\n", i, v)
	}
}
