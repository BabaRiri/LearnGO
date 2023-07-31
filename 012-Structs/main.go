package main

import "fmt"

//Define strutc to be used
type planet struct {
	name     string
	galaxy   string
	numMoons string
	distance float64
}

type habitable struct {
	planet
	habitable bool
}

func main() {

	//Created structs of the defined structs
	pEarth := planet{
		name:     "Earth",
		galaxy:   "Milkyway",
		numMoons: "4",
		distance: 0,
	}

	pMars := planet{
		name:     "Mars",
		galaxy:   "Milkyway",
		numMoons: "2",
		distance: 1737,
	}

	pPluto := planet{
		name:     "Pluto",
		galaxy:   "Milkyway",
		numMoons: "1",
		distance: 11.8,
	}

	fmt.Println(pEarth)
	fmt.Println(pMars)
	fmt.Println(pPluto)

	fmt.Printf("NAME: %v\tGALAXY: %v\tMOONS: %v\tDISTANCE: %v\n", pEarth.name, pEarth.galaxy, pEarth.numMoons, pEarth.distance)
	fmt.Printf("NAME: %v\tGALAXY: %v\tMOONS: %v\tDISTANCE: %v\n", pMars.name, pMars.galaxy, pMars.numMoons, pMars.distance)
	fmt.Printf("NAME: %v\tGALAXY: %v\tMOONS: %v\tDISTANCE: %v\n\n", pPluto.name, pPluto.galaxy, pPluto.numMoons, pPluto.distance)

	//created an embedded struct
	earthHabitable := habitable{
		planet:    pEarth,
		habitable: true,
	}

	fmt.Println(earthHabitable)
	fmt.Printf("PLANET: %v\tGALAXY: %v\tHABITABLE: %v\n\n", earthHabitable.planet.name, earthHabitable.planet.galaxy, earthHabitable.habitable)

	//Create an anonymous struct
	marsHabitable := struct {
		planet
		habitable bool
	}{
		planet:    pMars,
        habitable: false,
	}

	fmt.Println(marsHabitable)
	fmt.Printf("PLANET: %v\tGALAXY: %v\tHABITABLE: %v\n", marsHabitable.planet.name, marsHabitable.planet.galaxy, marsHabitable.habitable)

}
