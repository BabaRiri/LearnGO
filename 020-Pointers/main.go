package main

import "fmt"

func main() {
	x := 10
	fmt.Println("VALUE:", x)
	fmt.Println("ADRESS:", &x)

	fmt.Printf("%v\t%T\n", &x, &x)
	s := "Rethabile"
	fmt.Printf("%v\t%T\n", &s, &s)

}