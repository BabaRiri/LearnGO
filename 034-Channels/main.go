package main

import "fmt"

func main() {
	c := make(chan int, 2)
	/*go func() {
		c <- 16
	}()*/
	c <- 30
	c <- 11

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("--------")
	fmt.Printf("%T\n", c)
}