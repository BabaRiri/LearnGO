package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	gs := 100
	wg.Add(gs)
	var incrementor int64

	for i := 0; i < gs; i++ {
		go func(){
			atomic.AddInt64(&incrementor,1)
			fmt.Println(incrementor)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("FINAL: ", incrementor)
}