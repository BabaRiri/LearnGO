package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walk.")
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Baki"}
	youngRun(d1)

	d2 := dog{"Yugiro"}
	youngRun(d2)
	
	d2 = d2.changeName("Mr Orge")
	youngRun(d2)
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d
}
