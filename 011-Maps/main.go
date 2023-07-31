package main

import "fmt"

func main() {

	//creating a map
	myMap := map[string]float64{
		"Alphabet":  1.0,
		"Apple":     1.3,
		"Microsoft": 1.2,
		"Amazon":    0.93,
		"Facebook":  0.63,
	}

	fmt.Println("--------------------------------")
	fmt.Printf("LENGTH: %v\n", len(myMap))
	fmt.Printf("TYPE: %T\n", myMap)
	fmt.Println(myMap)

	//Accessing elements of a map
	fmt.Printf("\nThe value for Apple is %v\n", myMap["Apple"])
	fmt.Println("--------------------------------")

	//Creating a map using make
	myMap1 := make(map[string]int)
	myMap1["a"] = 1 // adding elements to map
	myMap1["b"] = 2
	myMap1["c"] = 3

	fmt.Printf("\nLENGTH: %v\n", len(myMap1))
	fmt.Printf("TYPE: %T\n", myMap1)
	fmt.Println(myMap1)

	fmt.Printf("\nThe value for 'a' is %v\n", myMap1["a"])
	fmt.Println("--------------------------------")

	//Interate a map with the for range loop
	for k, v := range myMap {
		fmt.Printf("KEY: %v\t\tVALUE: %v\n", k, v)
	}

	fmt.Println("--------------------------------")
	//Printing just the key
	for k := range myMap {
		fmt.Printf("KEY: %v\n", k)
	}

	fmt.Println("--------------------------------")
	//Printing just the value
	for _, v := range myMap {
		fmt.Printf("VALUE: %v\n", v)
	}
}
