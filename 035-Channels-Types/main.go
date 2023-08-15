package main

import "fmt"

func main() {
	c := make(chan int) //normal channel [biderectional]
	cr := make(<-chan int) //receiver only channel
	cs := make(chan<- int) //sender only channel

	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)
}

//channels are gonna be converted from general to specific but not specific to general