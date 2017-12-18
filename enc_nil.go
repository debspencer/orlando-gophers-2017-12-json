package main

import (
	"encoding/json"
	"fmt"
)

type Pizza struct {
	Name       string    `json:"name"`
	Size       string    `json:"size"`
	NumSlices  *int      `json:"slices"`
	Toppings   []string  `json:"toppings"`
}

func main() {
	pizza := Pizza{
		Name:       "Pineapple Lovers",
		Size:       "Large",
	}
	js, _ := json.MarshalIndent(&pizza, "", "    ")
	fmt.Println(string(js))
}
