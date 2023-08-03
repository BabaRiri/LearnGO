package main

import "fmt"

func main() {
    x := func(){
		fmt.Printf("My name is Baba Riri\n")
	}

	x()
	y("I am a variable")

}

var y = func(s string){
	fmt.Println(s)

}
