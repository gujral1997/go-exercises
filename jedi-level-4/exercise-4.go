package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 52}
	x = append(x, 53)
	fmt.Println(x)
}
