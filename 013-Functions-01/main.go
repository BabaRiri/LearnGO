package main

import "fmt"

//Creating types with composite types
type person struct {
	name string
}

type programmer struct {
	person
	coder bool
}

//Creating a method for a type
func (p person) say() {
	fmt.Printf("I am %v\n", p.name)
}

func (p programmer) say() {
	fmt.Printf("I am a Programmer named %v\n", p.name)
}

//Creating an interface for types using a certain method
type human interface {
	say()
}

func intro(h human) {
	h.say()

}

func main() {
	p1 := person {
		name: "Rethabile",
	}

	p2 := programmer {
        person: person {
			name: "Baba Riri",
		},
        coder: true,
    }

	p1.say()
	p2.say()

	fmt.Printf("\n")

	intro(p1)
	intro(p2)



}