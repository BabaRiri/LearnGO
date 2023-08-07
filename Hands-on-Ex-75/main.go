package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("VALUE[a]: %v\n", a)
	fmt.Printf("TYPE[a]: %T\n", a)
	fmt.Printf("DATA[a]: %v\n\n", *a)

	fmt.Printf("VALUE[b]: %v\n", b)
	fmt.Printf("TYPE[b]: %T\n", b)
	fmt.Printf("DATA[b]: %v\n\n", *b)

	fmt.Printf("VALUE[c]: %v\n", c)
	fmt.Printf("TYPE[c]: %T\n", c)
	fmt.Printf("DATA[c]: %v\n\n", *c)

	fmt.Printf("VALUE[d]: %v\n", d)
	fmt.Printf("TYPE[d]: %T\n", d)
	fmt.Printf("DATA[d]: %v\n\n", *d)

}