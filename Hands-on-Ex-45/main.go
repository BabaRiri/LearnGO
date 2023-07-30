package main

import "fmt"

func main() {
	numSlice := []int{1, 2, 3, 4}
	fmt.Println(numSlice)
	numSlice = append(numSlice, 5)
	fmt.Println(numSlice)
	numSlice = append(numSlice, 6, 7, 8, 9, 10)
	fmt.Println(numSlice)

	numSlice1 := []int{20, 21, 22, 23, 24, 25}
	numSlice = append(numSlice, numSlice1...)
	fmt.Println(numSlice)

}
