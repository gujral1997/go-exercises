package main

import (
	"fmt"
)

func main() {

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "James",
		friends:   map[string]int{"Ram": 1, "Sham": 2},
		favDrinks: []string{"Fanta", "Pepsi"},
	}
	fmt.Println(s)
}
