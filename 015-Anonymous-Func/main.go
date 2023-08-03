package main

import "fmt"

func main() {
	
	testFunc()

	func(){
		fmt.Println("Inside Anonymous function")
	}()


	x := func(s string){
		fmt.Println(s)
	}

	x("Function as a type")
}

func testFunc() {
	fmt.Println("I am in the TEST FUNCTION!")
}