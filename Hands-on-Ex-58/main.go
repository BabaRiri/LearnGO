package main

import "fmt"

func main() {
    fmt.Printf("foo() returns: %v\n", foo())

	a, b := bar()
	fmt.Printf("bar() returns: %v and %v", a, b)
}

func foo() int {
	return 100
}

func bar() (int, string) {
	return 200, "I am a string!!"
}