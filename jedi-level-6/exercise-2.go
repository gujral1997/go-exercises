package main

import (
	"fmt"
)

func foo(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func bar() {
	sum := []int{1, 2, 3, 4, 5, 6, 7, 8}
	println(foo(sum...))
}

func main() {
	fmt.Println(foo(1, 2, 3, 4))
	bar()
}
