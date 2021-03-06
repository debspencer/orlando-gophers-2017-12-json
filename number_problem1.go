package main

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
"toppings":["Peppers","Onions","Tomatos"]}`)

    m := make(map[string]interface{})
    json.Unmarshal(js, &m)
    var pizza Pizza
    for k, v := range m {
        switch k {
        case "name":      pizza.Name = v.(string)
        case "size":      pizza.Size = v.(string)
        case "slices":    pizza.NumSlices = v.(int)
        }
    }
    fmt.Printf("%#v\n", pizza)
}
