package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

type shape interface {
	Area() float64
}

func (s Square) Area() float64 {
	return s.length * s.width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func Info(s shape) float64 {
	return s.Area()
}

func main() {
	s := Square{
		length: 10,
		width:  10,
	}
	c := Circle{
		radius: 5,
	}

	fmt.Printf("Area of square: %v\n",s.Area())
	fmt.Printf("Area of circle: %v\n", c.Area())

	fmt.Println("Using the interface:")
	
	fmt.Printf("Area of square: %v\n",Info(s))
	fmt.Printf("Area of circle: %v\n", Info(c))
}
