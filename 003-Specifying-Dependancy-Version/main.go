package main

import (
	"fmt"

	"github.com/BabaRiri/puppy"
)

func main() {
	str1 := puppy.Bark()
	str2 := puppy.Barks()

	fmt.Printf("%v %v\n", str1, str2)

	//alternative way
	fmt.Println(puppy.Barks(), puppy.Bark())

	//Accessing dog module through puppy
	fmt.Printf("%v\n%v\n", puppy.BigBark(), puppy.BigBarks())

	puppy.Version11()
	puppy.Version12()
}
