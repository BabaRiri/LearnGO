package main

import "fmt"

func main() {
	ans := doMath(100, 100, add)
	fmt.Printf("If we add we get: %v\n",ans)
	
	ans1 := doMath(100, 100, subtract)
	fmt.Printf("If we subtract we get: %v\n",ans1)

	fmt.Printf("add : %T\n", add)
	fmt.Printf("subtract : %T\n", subtract)
	fmt.Printf("doMath : %T\n", doMath)
}

func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}