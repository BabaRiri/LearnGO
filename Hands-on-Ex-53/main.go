package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {
	p1 := person{
		firstName:   "Seward",
		lastName:    "Mupereri",
		favIceCream: []string{"Rum and Chocolate", "Vanilla", "Chocolate"},
	}
	p2 := person{
		firstName:   "Nthabiseng",
		lastName:    "Mupereri",
		favIceCream: []string{"Vanilla", "Chocolate", "Cookies & Cream"},
	}

	fmt.Printf("%v\n%v\n", p1, p2)
	
	fmt.Printf("\n")
	fmt.Printf("%v's favourite ice creams are: \n", p1.firstName)
	for i, v := range p1.favIceCream {
		fmt.Printf("%v - %v\n", i+1, v)
	}
	fmt.Printf("\n")
	
	fmt.Printf("%v's favourite ice creams are: \n", p2.firstName)
	for i, v := range p2.favIceCream {
		fmt.Printf("%v - %v\n", i+1, v)
	}
}
