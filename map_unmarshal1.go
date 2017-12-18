package main

import "encoding/json"
import "fmt"

func main() {
    js := []byte(`{"name": "Veggie Lovers","size": "Large","slices": 8,
"toppings":["Peppers","Onions","Tomatos"]}`)

    m := make(map[string]interface{})
    err := json.Unmarshal(js, &m)
    fmt.Printf("err = %v\n", err)
    for k, v := range m {
        fmt.Printf("%s = %v\n", k, v)
    }
}
