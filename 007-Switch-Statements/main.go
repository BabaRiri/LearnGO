package main

import "fmt"

func main() {
    x := 10

	switch {
	case x == 10:
		fmt.Println("10")
	
	case x == 20:
		fmt.Println("20")

	case x == 30:
		fmt.Println("30")
	
    default:
		fmt.Println("default")
	}

	x = 30

	switch x {
	case 10:
		fmt.Println("10")

	case 20:
		fmt.Println("20")
	
    case 30:
		fmt.Println("30")
	
    default:
		fmt.Println("default")
		
	}

	switch x {
	case 10:
		fmt.Println("10")

	case 20:
		fmt.Println("20")
	
    case 30:
		fmt.Println("30")
		fallthrough
	
    default:
		fmt.Println("Fallthrought made it get here")
		
	}
}