package main

import "fmt"

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		c = gen(c)
	}

	receive(c)

	fmt.Println("THE END!")
}

func receive(c <-chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}

func gen(c chan int) chan int {

	go func() {
		for i := 1; i <= 100; i++ {
			c <- i
		}
	}()

	return c
}