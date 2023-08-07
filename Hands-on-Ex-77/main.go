package main

import "fmt"

type person struct {
	name string
}

func myFunc1(p person, s string) person {
	p.name = s
	return p
}

func myFunc2(p *person, s string) {
	p.name = s
}

func main() {
	x := person{
		name: "Hanma",
	}
	fmt.Println("x: ",x)

	y := myFunc1(x, "Baki")
	fmt.Println("x: ",x)
	fmt.Println("y: ",y)

	myFunc2(&x, "Yugiro")
	fmt.Println("x: ",x)
}