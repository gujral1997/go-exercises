package main

import (
	"fmt"
)

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	x := person{"Ansh", "Gujral", []string{"Strawberry", "Cocolate"}}
	m := map[string]person{
		"A": x,
	}
	fmt.Println(x)
	fmt.Println(x.favFlavors)
	fmt.Println(m)
}
