package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("014-Interfaces-01/note.txt")
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

	b := bytes.NewBufferString(string(info)) //Created a string from the buffer (New Buffer)
	fmt.Println(b.String())
	b.WriteString(" A good mantra to live by!") //Added more text to the buffer
	fmt.Println(b.String())
	b.Reset() //Reset the buffer
	b.WriteString("The buffer was reset")
	fmt.Println(b.String())

	b.Write([]byte("\nThis is another way to write to a buffer")) //Alternate way to write to a buffer
	fmt.Println(b.String())


}