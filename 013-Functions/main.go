package main

import "fmt"

func main() {

	f1() // Function with no arguments and no return values
	f2(69) //Function with 1 argument and 0 return values
	fmt.Printf("I send: 2000 as an argument and got: %v returned back", f3(2000)) //Function with 1 argument and 1 return value
	x, y := f4(7, 16) //Function with 2 arguments and 2 return values
	fmt.Printf("I send: 7 and 16 as arguments and got: %v and %v returned back", x, y) 

}

func f1() {
	fmt.Println("I come from a function with no parameters and no returns")
}

func f2(num int) {
	fmt.Printf("I recieved: %d as an parameter\n", num)
}

func f3(num int) int {
	return num * 100
}

func f4(num int, num1 int) (int, int) {
	return num*2, num1*10
}