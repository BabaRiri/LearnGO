package main

import (
    "fmt"
)

// factorial function takes an integer n as input and returns the factorial of n
func factorial(n int) int {
    // if n is 0 or 1, return 1
    if n == 0 || n == 1 {
        return 1
    }
    // otherwise, return n multiplied by the factorial of n-1
    return n * factorial(n-1)
}

func main() {
    // define a variable to hold the input number
    var num int

    // prompt the user to enter a number
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    // calculate the factorial of the input number
    result := factorial(num)

    // print the result
    fmt.Printf("The factorial of %d is %d\n", num, result)
}
