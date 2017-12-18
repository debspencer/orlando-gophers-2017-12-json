package main

import "encoding/json"
import "fmt"

type Pizza struct {
	Name      string
	Size      string
	Slices    int
	Toppings  []string
}

func main() {
	pizza := Pizza{
		Name:      "Veggie Lovers",
		Size:      "Large",
		Slices:    8,
		Toppings:  []string{"Peppers", "Onions", "Tomatos"},
	}
	js, err := json.Marshal(&pizza)
	fmt.Printf("err=%v\n%s\n", err, string(js))
}
