package main

import (
	"fmt"
	"learngo/Hands-on-Ex-100/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	dog1 := canine{
		name: "Baki",
		age:  dog.Years(18),
	}

	fmt.Printf("%v is %v old in dog years", dog1.name, dog1.age)
}