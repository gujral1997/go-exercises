package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	fourWheel bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func main() {

	t := truck{vehicle{2, "white"}, true}
	s := sedan{vehicle{4, "black"}, true}
	fmt.Println(s)
	fmt.Println(t)
}
