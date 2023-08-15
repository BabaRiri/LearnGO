package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("THE END!")
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("EVEN CHANNEL VALUE: ", v)
		case v := <-o:
			fmt.Println("ODD CHANNEL VALUE: ", v)
		case v := <-q:
			fmt.Println("QUIT CHANNEL VALUE: ", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}