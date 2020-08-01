package main

import (
	"fmt"
)

func foo() int {
	return 1
}

func bar() (int, string) {
	return 1, "Hey"
}

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}
