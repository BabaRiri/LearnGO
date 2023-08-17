package main

import (
	"fmt"
)

type customErr struct {
	msg string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("It did not work my gee!: [ERROR] - %s", ce.msg)
}

func main() {
	c1 := customErr{
		msg: "You did something wrong!",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
