package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("A")
		wg.Done()
	}()
	go func() {
		fmt.Println("B")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Ran")
}
