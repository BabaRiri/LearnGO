package main

import "fmt"

func main() {
    s := myFunc1()
	fmt.Println(s)

	s1 := myFunc2()
	s3 := s1()
	fmt.Println(s3)
}

func myFunc1() string {
	return "Returned a string"
}

func myFunc2() func() string {
	return func() string {
        return "Returned a function"
    }
}