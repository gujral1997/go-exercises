package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		"a": []string{"1", "2", "3"},
	}

	fmt.Println(x)
	for k, v := range x {
		fmt.Println(k, v)
	}
}
