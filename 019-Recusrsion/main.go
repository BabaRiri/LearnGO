package main

import "fmt"

func main() {
    fmt.Printf("Recursive factorial: %v\n", factorial(4))
    fmt.Printf("Iterative factorial: %v\n", factLoop(4))
}

func factorial(a int) int {
	if a == 0 {
        return 1
    }
    return a * factorial(a-1)
}

func factLoop(a int) int {
	result := 1
	for i := 0; i <= a; i++ {
		result *= a
		a--
	}
	return result
}