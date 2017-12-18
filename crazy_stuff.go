package main

import (
	"encoding/json"
	"fmt"
)

type Pizza struct {
	Name       string    `json:"-,"`
	DontShow   string    `json:"-"`
	Size       string
	NumSlices  int
	Toppings   []string
}

func main() {
	pizza := Pizza{
		Name:       "Pineapple Lovers",
		Size:       "Large",
	}
	js, _ := json.MarshalIndent(&pizza, "", "    ")
	fmt.Println(string(js))
}
