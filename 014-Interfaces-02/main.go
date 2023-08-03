package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	name string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.name))
}

func main() {

	p := person {
		name: "Rethabile",
	}

	f, err := os.Create("014-Interfaces-02/note2.txt")
	if err!= nil {
        log.Fatalf("Error: %v", err)
    }
	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())

	//Clearing the contents of a file
	if err := os.Truncate("014-Interfaces-02/note2.txt", 0); err != nil {
    	log.Printf("Failed to truncate: %v", err)
	}
}