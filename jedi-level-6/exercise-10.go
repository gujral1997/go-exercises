package main

import (
	"fmt"
)

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	x := foo()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}
