package main

import (
	"encoding/json"
	"fmt"
)

type Pizza struct {
	Name       string    `json:",omitempty"`
	Size       string    `json:",omitempty"`
	NumSlices  int       `json:",omitempty"`
	Toppings   []string  `json:",omitempty"`
}

func main() {
	pizza := Pizza{
		Name:       "Pineapple Lovers",
		Size:       "Large",
	}
	js, _ := json.MarshalIndent(&pizza, "", "    ")
	fmt.Println(string(js))
}
