package main

import "bytes"
import "fmt"
import "encoding/json"

type Pizza struct {
	Name       string    `json:"name,omitempty"`
	Size       string    `json:"size,omitempty"`
	Toppings   []string  `json:"toppings,omitempty"`
}
func main() {
	pizza := Pizza{
		Name:       "Pineapple Lovers",
		Size:       "Large" }
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "    ")
	err := enc.Encode(&pizza)
	fmt.Printf("err = %v\njson = %s\n", err, buf.String())
}
