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

	var family []person
	err := json.Unmarshal(byteSlice, &family)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All of the data: ", family)
	for _, v := range family {
		fmt.Printf("\n\nFirst: %v\nLast: %v\nAge: %v", v.First, v.Last, v.Age)
		
	}

}