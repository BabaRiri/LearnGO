package main

import "fmt"

type dog struct {
	name string
}

func (d dog) walk() {
	fmt.Printf("My name is %v and i am walking!\n", d.name)
}

func (d *dog) run() {
	fmt.Printf("My name is %v and i am running!\n", d.name)
}

func main() {
	pet1 := dog{
		name: "Baki",
	}

	pet1.walk()
	pet1.run()

	pet2 := &dog{
		name: "Yugiro",
	}

	pet2.walk()
	pet2.run()
}