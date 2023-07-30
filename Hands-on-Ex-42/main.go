package main

import "fmt"

func main() {
	arr := [5]int {}
	for i := 0; i < 5; i++ {
		arr[i] = (i*10) +10
	}
	fmt.Printf("LENGTH: %v\n", len(arr))
	fmt.Printf("TYPE: %T\n", arr)

	for i := 0; i < 5; i++ {
		fmt.Printf("INDEX: %v\tVALUE: %v\n", i, arr[i])
	}
}
