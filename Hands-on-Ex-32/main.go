package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(100)

	for {

		fmt.Printf("%d\n", a)
		if a == 0 {
			break
		}
		a--
	}
}
