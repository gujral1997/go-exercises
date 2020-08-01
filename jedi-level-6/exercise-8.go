package main

import (
	"fmt"
)

func foo() func() int {
	return func() int {
		fmt.Println("Returned As Function")
		return 1
	}
}

func main() {
	x := foo()
	fmt.Println(x())
}
