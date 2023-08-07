package main

import "fmt"

func intDelta( n *int) {
	*n = 100
}

func sliceDelta( ii []int) {
	ii[0] = 888
}

func mapDelta( m map[string]int, s string) {
	m[s] = 4
}

func main() {
	x := 10
	fmt.Println("VALUE:", x)
	fmt.Println("ADRESS:", &x)
	fmt.Printf("&x Value: %v\t&x Type: %T\n\n", &x, &x)

	s := "Rethabile"
	fmt.Printf("&s Value: %v\t&s Type: %T\n\n", &s, &s)

	y := &s
	fmt.Printf("y Value: %v\ty Type :%T\n", y, y)
	fmt.Printf("y Value: %v\n", y)
	fmt.Printf("*y Value: %v\n\n", *y)
	
	fmt.Println("VALUE:", x)
	intDelta(&x)
	fmt.Println("VALUE:", x)

	xi := []int{1, 2, 3, 4}
	fmt.Println("\n",xi)
	sliceDelta(xi)
	fmt.Println(xi)

	mp := map[string]int{
		"Rethabile": 3,
		"Nthabiseng": 25,
		"Seward": 25,
	}
	fmt.Println("\n",mp)
	mapDelta(mp, "Rethabile")
	fmt.Println(mp)

}