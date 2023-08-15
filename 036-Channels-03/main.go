package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c)

	receive(c)

	fmt.Println("The end!")
}

// send channel
func send(c chan<- int) {
	c <- 42
}

// receive channel
func receive(c <-chan int) {
	fmt.Println("VALUE RECIEVED:", <-c)
}
