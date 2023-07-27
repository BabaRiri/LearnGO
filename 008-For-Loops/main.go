package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	x := 10

	for x > 5 {
		fmt.Println(x)
		x--
	}
}