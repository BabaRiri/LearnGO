package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Username  string
	FirstName string
	LastName  string
	Age       int
}

func main() {
	user1 := user{
		Username:  "NanaRiri",
		FirstName: "Rethabile",
		LastName:  "Mupereri",
		Age:       3,
	}

	user2 := user{
		Username:  "MamaRiri",
		FirstName: "Nthabiseng",
		LastName:  "Mupereri",
		Age:       25,
	}

	user3 := user{
		Username:  "BabaRiri",
		FirstName: "Seward",
		LastName:  "Mupereri",
		Age:       25,
	}

	users := []user{
		user1,
		user2,
		user3,
	}

	//Marshaling a slice of type user to JSON
	byteSlice, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteSlice))
}