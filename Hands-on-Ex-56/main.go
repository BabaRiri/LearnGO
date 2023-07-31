package main

import "fmt"

func main() {
    anonStruct := struct {
		first string
		friends map[string]int
		favDrink []string
	}{
		first: "Marcus",
        friends: map[string]int{
            "Aristotle": 1005,
            "Plato": 2635,
        },
        favDrink: []string{
            "Tea",
            "Coffee",
        },
	}

	fmt.Printf("NAME: %v\n", anonStruct.first)

	fmt.Println("FRIENDS:")
	for i, v := range anonStruct.friends {
		fmt.Printf("  %v (%v years old!)\n", i, v)
	}

	fmt.Println("FAVOURITE DRINK:")
	for _, v := range anonStruct.favDrink {
		fmt.Printf("  %v\n", v)
	}
}