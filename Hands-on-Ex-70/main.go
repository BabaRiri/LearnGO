package main

func myFunc1() func() int {
	return func() int {
        return 1
    }
}

func main() {
	f := myFunc1()
    println(f())
}