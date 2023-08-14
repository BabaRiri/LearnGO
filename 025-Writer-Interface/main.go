package main

import (
	"fmt"
	"io"
	"os"
)

//This program illustrates how the type writer interface is used with Displaying out put to Stdout

func main() {
	fmt.Println("Hello Riri")
	fmt.Fprintln(os.Stdout, "Hello Riri")
	io.WriteString(os.Stdout, "Hello Riri")
}
