package main

import "fmt"

func main() {
	testFunc()

	func(s string){
		fmt.Println(s)
	}("ANONYMOUS function here")

	func(){}()
}

func testFunc() {
	fmt.Println("I am in the TEST FUNCTION!")
}