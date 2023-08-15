package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("[BEFORE] CPU: ", runtime.NumCPU())
	fmt.Println("[BEFORE] Goroutines: ", runtime.NumGoroutine())

	wg.Add(2)
	go func (){
		fmt.Println("I am in the first function!")
		wg.Done()
	}()

	go func (){
		fmt.Println("I am in the second function!")
		wg.Done()
	}()

	fmt.Println("[DURING] CPU: ", runtime.NumCPU())
	fmt.Println("[DURING] Goroutines: ", runtime.NumGoroutine())

	
	wg.Wait()
	
	fmt.Println("I am in the main function!")
	fmt.Println("[AFTER] CPU: ", runtime.NumCPU())
	fmt.Println("[AFTER] Goroutines: ", runtime.NumGoroutine())


}