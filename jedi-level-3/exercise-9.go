package main

import (
	"fmt"
)

func main() {
	favSport := "surfing"
	switch favSport {
	case "surfing":
		fmt.Println("Print")
	case "dancing":
		fmt.Println("Not Print")
	default:
		fmt.Println("Default")
	}
}
