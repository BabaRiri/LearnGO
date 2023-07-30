package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(100)
	for a >= 0 {
		fmt.Printf("%d\n", a)
		a--
	}
}
