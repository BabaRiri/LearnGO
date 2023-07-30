package main

import "fmt"

func main() {
	slice1 := []string{
		"Tesla",
		"Nissan",
		"Chevrolet",
		"BMW",
		"Audi",
	}
	fmt.Println(slice1)
	fmt.Printf("%T\n", slice1)
	fmt.Printf("LENGTH: %v\n\n", len(slice1))

	for _, brand := range slice1 {
		fmt.Printf("%v\n", brand)
	}

	fmt.Printf("-------------------------------\n")
	fmt.Printf("%v\n", slice1[0])
	fmt.Printf("%v\n", slice1[1])
	fmt.Printf("%v\n", slice1[2])
	fmt.Printf("-------------------------------\n")

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%v\n", slice1[i])
	}
}
