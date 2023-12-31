package main

import "fmt"

type person struct {
	name string
}

func main() {

	f1() // Function with no arguments and no return values

	defer f2(69) //Function with 1 argument and 0 return values

	fmt.Printf("I send: 2000 as an argument and got: %v returned back\n", f3(2000)) //Function with 1 argument and 1 return value
	
	x, y := f4(7, 16) //Function with 2 arguments and 2 return values
	fmt.Printf("I send: 7 and 16 as arguments and got: %v and %v returned back\n", x, y)

	mySlice := []int {10, 20 , 30, 40, 50}
	total := variadicFunc(mySlice...) //Function call for the variadic function + Unfurling a slice
	fmt.Printf("The sum of the arguments is: %v\n", total)

	p := person {
		name: "Seward",
	}
	
	p.show()
}

func f1() {
	fmt.Println("I come from a function with no parameters and no returns")
}

func f2(num int) {
	fmt.Printf("\nDefered(f2): I recieved: %d as an parameter\n", num)
}

func f3(num int) int {
	return num * 100
}

func f4(num int, num1 int) (int, int) {
	return num*2, num1*10
}

//Creating a variadic function that returns a value
func variadicFunc(x ...int) (int) {
	sum := 0
    for _, i := range x {
        sum += i
    }
    return sum
}

//Adding methods to types
func (p person) show() {
	fmt.Printf("\nThe name is %v\n", p.name)
   }
