package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			fmt.Println(incrementer)
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end", incrementer)
}
