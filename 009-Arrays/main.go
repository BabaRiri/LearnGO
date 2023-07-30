package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	fmt.Println(array1)

	array2 := [...]int{1, 2, 3}
	fmt.Println(array2)

	var array3 [2]string
	array3[0] = "First Elelemnt\n"
	array3[1] = "Second Element"
	fmt.Println(array3)
}
