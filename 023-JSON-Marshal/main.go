// this program takes a struct and marshal it into json
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last string
	Age int
}

func main() {
	p1 := person{
		First: "Rethabile",
		Last: "Mupereri",
		Age: 3,
	}

	p2 := person{
		First: "Nthabiseng",
		Last: "Mupereri",
		Age: 25,
	}

	people := []person{
		p1,
		p2,
	}

	fmt.Println(people)

	byteSlice, err := json.Marshal(people)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byteSlice))
}