package main

import (
	"fmt"

	"github.com/BabaRiri/puppy"
)

func main() {
    fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.BigBarks())

	puppy.BigBark()
}