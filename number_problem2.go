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
"toppings":["Peppers","Onions","Tomatos"]}`)

    m := make(map[string]interface{})

    dec := json.NewDecoder(bytes.NewReader(js))
    dec.UseNumber()
    err := dec.Decode(&m)

    var pizza Pizza
    for k, v := range m {
        switch k {
        case "name":      pizza.Name = v.(string)
        case "size":      pizza.Size = v.(string)
        case "slices":    pizza.NumSlices = v.(int)
        }
    }
    fmt.Printf("err = %v\n%#v\n", err, pizza)
}
