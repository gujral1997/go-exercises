package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("Not Print")
	case true:
		fmt.Println("Print")
	}
}
