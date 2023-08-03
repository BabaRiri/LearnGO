package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
    p := person{
		name: "Rethabile",
		age: 3,
	}

	p.speak()
}


func (p person) speak() {
	fmt.Printf("My name is %s. I am %v years old.\n", p.name, p.age)
}