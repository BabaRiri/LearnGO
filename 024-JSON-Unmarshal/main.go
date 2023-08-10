package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Rethabile","Last":"Mupereri","Age":3},{"First":"Nthabiseng","Last":"Mupereri","Age":25}]`
	byteSlice := []byte(s)
	
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(string(byteSlice))
	fmt.Printf("%T\n",string(byteSlice))

	var family []person
	err := json.Unmarshal(byteSlice, &family)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("All of the data", family)
	for i, v := range family {
		fmt.Println("\nMEMBER NUMBER:", i)
		fmt.Println(v.First, v.Last, v.Age)
		
	}
}
