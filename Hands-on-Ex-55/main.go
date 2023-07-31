package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {

	myCar1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Rivian",
		model: "R1S",
		doors: 5,
		color: "El Cap Granite",
	}

	myCar2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Toyota",
		model: "Land Cruiser",
		doors: 5,
		color: "White",
	}

	fmt.Println(myCar1)
	fmt.Println(myCar2)
	fmt.Printf("\n")

	myMap := map[string]vehicle{
		myCar1.make: myCar1,
		myCar2.make: myCar2,
	}

	for i, v := range myMap {
		fmt.Printf("MAKE: %v\n", i)
		fmt.Printf("ELECTRIC ENGINE: %v\n\n", v.engine.electric)
	}

}
