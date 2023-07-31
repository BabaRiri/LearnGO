package main

import "fmt"

type planet struct {
	name string
	galaxy string
	numMoons string
	distance float64
}

func main() {

	pEarth := planet{
		name: "Earth",
		galaxy: "Milkyway",
        numMoons: "4",
        distance: 0,
	}

	pMars := planet{
		name: "Mars",
        galaxy: "Milkyway",
        numMoons: "2",
        distance: 1737,
	}

	pPluto := planet{
		name: "Pluto",
        galaxy: "Milkyway",
        numMoons: "1",
        distance: 11.8,
	}

	fmt.Printf("NAME: %v\tGALAXY: %v\tMOONS: %v\tDISTANCE: %v\n", pEarth.name, pEarth.galaxy, pEarth.numMoons, pEarth.distance)
	fmt.Printf("NAME: %v\tGALAXY: %v\tMOONS: %v\tDISTANCE: %v\n", pMars.name, pMars.galaxy, pMars.numMoons, pMars.distance)
	fmt.Printf("NAME: %v\tGALAXY: %v\tMOONS: %v\tDISTANCE: %v\n", pPluto.name, pPluto.galaxy, pPluto.numMoons, pPluto.distance)
	
}