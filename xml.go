package main

import "fmt"
import "encoding/json"
import "encoding/xml"

type Pizza struct {
	Name       string    `json:"name,omitempty" xml:"name,omitempty"`
	Size       string    `json:"size,omitempty" xml:"size,omitempty"`
	NumSlices  int       `json:"slices,omitempty" xml:"slices,omitempty"`
	Toppings   []string  `json:"toppings,omitempty" xml:"toppings,omitempty"`
}

func main() {
	pizza := Pizza{
		Name:       "Pineapple Lovers",
		Size:       "Large",
	}
	js, _ := json.MarshalIndent(&pizza, "", "    ")
	x, _ := xml.MarshalIndent(&pizza, "", "    ")
	fmt.Println(string(js))
	fmt.Println(string(x))
}
