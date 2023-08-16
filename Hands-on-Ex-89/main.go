package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(){
		c <- 42
	}()

	x := make(chan int, 1)
	
	x <- 16

	fmt.Println(<-c)
	fmt.Println(<-x)
}
