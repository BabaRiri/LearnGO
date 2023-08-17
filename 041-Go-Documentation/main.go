package main

import (
    "fmt"
)

// Factorial function takes an integer n as input and returns the factorial of n
func factorial(n int) int {
    // If n is 0 or 1, return 1
    if n == 0 || n == 1 {
        return 1
    }
    // Otherwise, return n multiplied by the factorial of n-1
    return n * factorial(n-1)
}

func main() {
    // Define a variable to hold the input number
    var num int

    // Prompt the user to enter a number
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    // Calculate the factorial of the input number
    result := factorial(num)

    // Print the result
    fmt.Printf("The factorial of %d is %d\n", num, result)
}
