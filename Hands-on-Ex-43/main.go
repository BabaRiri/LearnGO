package main

import "fmt"

func main() {
	slice := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 100}

	for i, num := range slice {
		fmt.Printf("INDEX: %d\t VALUE: %v\n", i, num)
	}
}
