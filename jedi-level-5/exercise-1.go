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
	fmt.Println(x)
	fmt.Println(x.favFlavors)
}
