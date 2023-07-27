package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	x := 10

	for x > 5 {
		fmt.Println(x)
		x--
	}

	for {
		fmt.Println(x)
		x--
		if x == 0 {
			fmt.Println("About to break")
			break
		}
		fmt.Println("Don't break")
	}

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("counting even numbers: ", i)
	}
	
	provinces := map[string]string{
		"EC": "Eastern Cape",
		"FS": "Free State",
		"GP": "Gauteng",
		"KZN": "KwaZulu-Natal",
		"LP": "Limpopo",
		"MP": "Mpumalanga",
		"NC": "Northern Cape",
		"NW": "North West",
		"WC": "Western Cape",
	}
	
	for key, value := range provinces {
		fmt.Printf("%s: %s\n", key, value)
	}
	
}