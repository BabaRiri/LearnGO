package main

import "fmt"

func main() {
    mySlc := []int { 1, 2, 3, 4, 5}
	fmt.Println(sum(mySlc))
}

func sum(slc []int) (total int) {
	total = 0
	for _, v := range slc {
        total += v
    }
	return
}