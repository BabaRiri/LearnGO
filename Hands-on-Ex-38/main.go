package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c:=0
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("i = %v\tc = %d\tx = %v\n", i, c, x)
			c++
		}
	}
}
