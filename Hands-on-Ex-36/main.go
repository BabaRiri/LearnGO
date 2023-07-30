package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
	}

	for i, v := range m {
		fmt.Printf("KEY: %v\t VALUE: %v\n", i, v)

	}
}