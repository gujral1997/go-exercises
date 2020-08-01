package main

import (
	"fmt"
)

type hotdog int

var x hotdog

func main() {
	fmt.Printf("%T\n", x)
	x = 2
	fmt.Println(x)
}
