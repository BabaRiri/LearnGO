package main

import (
	"fmt"
	"sort"
)

type Planet struct {
    name     string
    numMoons int
}

func main() {
    planets := []Planet{
        {"Mercury", 0},
        {"Venus", 0},
        {"Earth", 1},
        {"Mars", 2},
        {"Jupiter", 79},
        {"Saturn", 82},
        {"Uranus", 27},
        {"Neptune", 14},
    }

    sort.Slice(planets, func(i, j int) bool {
        return planets[i].numMoons < planets[j].numMoons
    	},
	)

    fmt.Println(planets)
}
