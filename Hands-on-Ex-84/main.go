package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func (p *person) speak() {
	fmt.Printf("My name is %s\n", p.Name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		Name: "Rethabile",
		Age: 3,
	}

	saySomething(&p1)
	//saySomething(p1) wont work
	p1.speak()
}