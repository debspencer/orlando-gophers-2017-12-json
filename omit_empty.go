package main

import "encoding/json"
import "fmt"

type Pizza struct {
	Name       string    `json:"name,omitempty"`
	Size       string    `json:"size,omitempty"`
	NumSlices  int       `json:"slices,omitempty"`
	Toppings   []string  `json:"toppings,omitempty"`
}

func main() {
	pizza := Pizza{
		Name:       "Pineapple Lovers",
		Size:       "Large",
	}
	js, _ := json.MarshalIndent(&pizza, "", "    ")
	fmt.Println(string(js))
}
