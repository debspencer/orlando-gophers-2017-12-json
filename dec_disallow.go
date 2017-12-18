package main

import "bytes"
import "encoding/json"
import "fmt"

type Pizza struct {
    Name       string
    Size       string
    NumSlices  int
    Toppings   []string
}

func main() {
    js := []byte(`{"name": "Veggie Lovers","size": "Large","slices": 8,
"toppings":["Peppers","Onions","Tomatos"],"ExtraStuff":"goes here"}`)

    dec := json.NewDecoder(bytes.NewReader(js))
    dec.UseNumber()
    dec.DisallowUnknownFields()
    err := dec.Decode(&pizza)
    fmt.Println("Error = ", err)
}
