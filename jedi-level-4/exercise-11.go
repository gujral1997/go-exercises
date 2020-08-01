package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		"a": []string{"1", "2", "3"},
	}

	x["gg"] = []string{"1", "wp", "3"}

	delete(x, "gg")

	fmt.Println(x)
	for k, v := range x {
		fmt.Println(k, v)
	}
}
