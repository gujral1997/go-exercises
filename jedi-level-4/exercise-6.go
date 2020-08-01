package main

import (
	"fmt"
)

func main() {
	x := make([]string, 50, 50)
	x = []string{"1", "2", "22"}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
