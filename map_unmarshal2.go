package main

import "bytes"
import "encoding/json"
import "fmt"

func main() {
    js := []byte(`{"name": "Veggie Lovers","size": "Large","slices": 8,
"toppings":["Peppers","Onions","Tomatos"]}`)

    m := make(map[string]interface{})
    dec := json.NewDecoder(bytes.NewReader(js))
    err := dec.Decode(&m)
    fmt.Printf("err = %v\n", err)
    for k, v := range m {
        fmt.Printf("%s = %v\n", k, v)
    }
}
