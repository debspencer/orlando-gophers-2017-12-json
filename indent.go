package main

import "bytes"
import "encoding/json"
import "fmt"

func main() {
	js := []byte(`{"name": "Veggie Lovers","size": "Large","slices": 8,
"toppings":["Peppers","Onions","Tomatos"]}`)

	var b bytes.Buffer
	json.Indent(&b, js, ">", "    ")
	fmt.Println(string(js))
	fmt.Println(b.String())
}
