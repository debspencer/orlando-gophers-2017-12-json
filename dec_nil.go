package main

import "encoding/json"
import "fmt"

type Pizza struct {
    Name       *string    `json:"name"`
    Size       *string    `json:"size"`
    NumSlices  *int       `json:"slices"`
    Toppings   []string   `json:"toppings"`
}

func main() {
	decode("All Filled out", `{"name": "Veggie Lovers","size": "Large","slices": 8,
"toppings":["Peppers","Onions","Tomatos"]}`)
	decode("Empty", "{}")
}

func decode(name, js string) {
	fmt.Println(name, "====")
	var pizza Pizza
	json.Unmarshal([]byte(js), &pizza)
	if pizza.Name != nil { fmt.Println("Got Name: ", *pizza.Name) }
	if pizza.Size != nil { fmt.Println("Got Size: ", *pizza.Size) }
	if pizza.NumSlices != nil { fmt.Println("Got Size: ", *pizza.NumSlices) }
	if len(pizza.Toppings) > 0 { fmt.Println("Got Toppings: ", pizza.Toppings) }
	fmt.Println(name, "====", "Done")
}
