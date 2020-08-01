package main

import (
	"fmt"
)

type person struct {
	name string
}

func changeMe(p *person) {
	(p).name = "Arush Gujral"
	// (*p).name = "Arush Gujral"
}

func main() {
	p1 := person{"Ansh Gujral"}
	fmt.Println(p1.name)
	changeMe(&p1)
	fmt.Println(p1.name)
}
