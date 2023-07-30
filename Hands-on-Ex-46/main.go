package main

import "fmt"

func main() {
	numSlice := []int{1, 2, 3, 4}
	fmt.Println(numSlice)
	numSlice = append(numSlice[:1], numSlice[3:]...)
	fmt.Println(numSlice)
}
