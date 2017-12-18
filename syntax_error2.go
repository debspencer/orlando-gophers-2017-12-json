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
    js := []byte(`
{
    "name": "Veggie Lovers",
    "size": "Large",
    "slices": 8,
    "toppings":["Peppers","Onions","Tomatos"],
}`)

    var pizza Pizza
    err := json.Unmarshal(js, &pizza)
    fmt.Printf("err = %v\n", err)
}
