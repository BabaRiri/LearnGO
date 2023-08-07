package main

import "fmt"

func main() {
	x := 10
	fmt.Println("VALUE:", x)
	fmt.Println("ADRESS:", &x)
	fmt.Printf("&x Value: %v\t&x Type: %T\n", &x, &x)

	s := "Rethabile"
	fmt.Printf("&s Value: %v\t&s Type: %T\n", &s, &s)

	y := &s
	fmt.Printf("y Value: %v\ty Type :%T\n", y, y)
	fmt.Printf("y Value: %v\n", y)
	fmt.Printf("*y Value: %v\n", *y)
}