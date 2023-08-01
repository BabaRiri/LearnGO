package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("014-Interfaces/note.txt")
	if err!= nil {
        log.Fatalf("Error: %v", err)
    }
	defer f.Close()

	info := []byte("Drop by drop, the bucket gets filled!")

	_, err = f.Write(info)
	if err!= nil {
        log.Fatalf("Error: %v", err)
    }
}