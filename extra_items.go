package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Pizza struct {
	Name       string    `json:"name,omitempty"`
	Size       string    `json:"size,omitempty"`
	NumSlices  int       `json:"slices,omitempty"`
	Toppings   []string  `json:"toppings,omitempty"`
	OrderTime  time.Time
	Cost       float64
}

func main() {
	pizza := Pizza{
		Name:       "Pineapple Lovers",
		Size:       "Large",
		OrderTime:  time.Now(),
		Cost:       9.99,
	}
	js, _ := json.MarshalIndent(&pizza, "", "    ")
	fmt.Println(string(js))
}
