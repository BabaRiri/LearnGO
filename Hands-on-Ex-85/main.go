package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementor := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func(){
			v := incrementor
			runtime.Gosched()
			v++
			incrementor = v
			fmt.Println(incrementor)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("FINAL: ", incrementor)
}