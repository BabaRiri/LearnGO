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
		favIceCream: []string{"Rum and Raisins", "Vanilla", "Chocolate"},
	}

	p2 := person{
		firstName:   "Nthabiseng",
		lastName:    "Mopelong",
		favIceCream: []string{"Vanilla", "Chocolate", "Cookies & Cream"},
	}

	myMap := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(myMap)
	for k, v := range myMap {
		fmt.Printf("%v\n", k)
		for _, v2 := range v.favIceCream {
			fmt.Printf("  %v\n", v2)
		}
		fmt.Printf("\n")
	}
}