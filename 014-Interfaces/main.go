package main

import (
	"bytes"
	"fmt"
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
    } else {
		fmt.Println("Infomation successfully written!")
	}

	b := bytes.NewBufferString(string(info))
	fmt.Println(b.String())
	b.WriteString(" A good mantra to live by!")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("The buffer was reset")
	fmt.Println(b.String())

	b.Write([]byte("This is another way to write to a buffer"))
	fmt.Println(b.String())


}