package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2}
	y := []int{1, 2}
	z := [][]int{x, y}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	for _, xs := range z {
		for _, v := range xs {
			fmt.Println(v)
		}
	}
}
