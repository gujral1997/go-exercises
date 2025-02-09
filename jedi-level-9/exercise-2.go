package main

import (
	"fmt"
)

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Println("Hello")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Ansh",
	}
	saySomething(&p1)
	p1.speak()
}
