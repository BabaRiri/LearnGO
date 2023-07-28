package main

import (
	"fmt"
	"math/rand"
)

func main() {
    x := rand.Intn(251)

	fmt.Printf("x = %v\n", x)

	if (x >= 0 && x <= 100) {
		fmt.Printf("x is between 0 and 100")
	} else if (x > 100 && x <= 200) {
		fmt.Printf("x is between 101 and 200")
	} else {
		fmt.Printf("x is not between 201 and 250")
	}
}