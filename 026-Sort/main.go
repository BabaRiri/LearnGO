package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{4, 7, 3, 42, 99, 16, 31, 10, 9, 12}
	strSlice := []string{"Rethabile", "Tafadzwa", "Blessing", "Seward", "Richard", "Nthabiseng", "Hapiness"}

	fmt.Println(intSlice)
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	fmt.Println(strSlice)
	sort.Strings(strSlice)
	fmt.Println(strSlice)

	
}