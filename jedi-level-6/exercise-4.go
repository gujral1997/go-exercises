package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, p.age)
}

func main() {
	p1 := person{"Ansh", "Gujral", 22}
	p1.speak()
}
