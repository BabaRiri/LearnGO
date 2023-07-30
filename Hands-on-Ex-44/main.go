package main

import "fmt"

func main() {
	slice := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 100}

	fmt.Println("ORIGINAL SLICE")
	for i, num := range slice {
		fmt.Printf("INDEX: %d\t VALUE: %v\n", i, num)
	}

	fmt.Printf("----------------------------------------------------------------\n")
	fmt.Println("BOTTOM SLICE")
	slice1 := slice[:5]

	for i, num := range slice1 {
		fmt.Printf("INDEX: %d\t VALUE: %v\n", i, num)
	}

	fmt.Printf("----------------------------------------------------------------\n")
	fmt.Println("TOP SLICE")
	slice2 := slice[5:]

	for i, num := range slice2 {
		fmt.Printf("INDEX: %d\t VALUE: %v\n", i, num)
	}
}
