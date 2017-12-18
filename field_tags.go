package main

import "encoding/json"
import "fmt"

type Pizza struct {
	Name       string    `json:"name"`
	Size       string    `json:"size"`
	NumSlices  int       `json:"slices"`
	Toppings   []string  `json:"toppings"`
}

func main() {
	pizza := Pizza{
		Name:       "Veggie Lovers",
		Size:       "Large",
		NumSlices:  8,
		Toppings:   []string{"Peppers", "Onions", "Tomatos"},
	}
	js, _ := json.MarshalIndent(&pizza, "", "    ")
	fmt.Println(string(js))
}
