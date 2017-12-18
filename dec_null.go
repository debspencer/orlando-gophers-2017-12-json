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
    js := []byte(`{"name": null, "size": null,"slices": null, "toppings": null}`)

    var pizza Pizza
    err := json.Unmarshal(js, &pizza)
    fmt.Printf("err = %v %#v\n", err, pizza)
}
