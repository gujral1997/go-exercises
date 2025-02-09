package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementer int64
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end", incrementer)
}
