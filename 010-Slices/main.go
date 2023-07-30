package main

import "fmt"

func main() {

	//Creating a slice of car brands
	slice1 := []string{
		"Tesla",
		"Nissan",
		"Chevrolet",
		"BMW",
		"Audi",
	}

	//Printing some values from the slice
	fmt.Println(slice1)
	fmt.Printf("%T\n", slice1)
	fmt.Printf("LENGTH: %v\n\n", len(slice1))

	//For range loop Iterating over the slice
	for _, brand := range slice1 {
		fmt.Printf("%v\n", brand)
	}

	//Accessing elements of the slice using indexes
	fmt.Printf("-------------------------------\n")
	fmt.Printf("%v\n", slice1[0])
	fmt.Printf("%v\n", slice1[1])
	fmt.Printf("%v\n", slice1[2])
	fmt.Printf("-------------------------------\n")

	//Accessing elements of the slice using for loop and indexes
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%v\n", slice1[i])
	}

	//Adding vlaues to a slice using append and printint new slice
	slice1 = append(slice1, "Rivian", "Lucid", "BYO")
	
	fmt.Printf("-------------------------------\n")
	
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%v\n", slice1[i])
	}

	fmt.Printf("-------------------------------\n")

	//Slicing a slice
	fmt.Println(slice1)
	fmt.Printf("%v\n", slice1[:2])
	fmt.Printf("%v\n", slice1[0:2])
	fmt.Printf("%v\n", slice1[2:])
	
	fmt.Printf("-------------------------------\n")

	//Deleting an element from a slice
	fmt.Println(slice1)
	slice1 = append(slice1[:3], slice1[4:]...)
	fmt.Println(slice1)

	fmt.Printf("-------------------------------\n")

	//Declaring a slice using make
	slice2 := make([]string, 0, 5)
	fmt.Println(slice2)
	fmt.Printf("%v\n", len(slice2))
	fmt.Printf("%v\n", cap(slice2))

	slice2 = append(slice2, "10", "20", "30", "40", "50")
	fmt.Println(slice2)
	fmt.Printf("%v\n", len(slice2))
	fmt.Printf("%v\n", cap(slice2))
	
	slice2 = append(slice2, "60", "70", "80")
	fmt.Println(slice2)
	fmt.Printf("%v\n", len(slice2))
	fmt.Printf("%v\n", cap(slice2))

	fmt.Printf("-------------------------------\n")

	//Creating a multi-dimensional slice
	slice3 := [][]string{slice1, slice2}
	fmt.Println(slice3)

}
