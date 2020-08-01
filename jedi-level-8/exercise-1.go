package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "M",
		Age:   75,
	}

	u2 := user{
		First: "N",
		Age:   35,
	}

	users := []user{u1, u2}

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
