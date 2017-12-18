package main

import "encoding/json"
import "fmt"

type PizzaSize string

type Pizza struct {
	Name       string
	Size       PizzaSize
}
func (sz PizzaSize) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case "Medium", "Large", "Extra-Large":
	default: return fmt.Errorf("Bad Pizza size: '%s'", string(b))
	}
	return nil
}
func main() {
	js := []byte(`{"name": "Veggie Lovers","size": "Large"}`)

	var pizza Pizza
	err := json.Unmarshal(js, &pizza)
	fmt.Printf("err = %v\n%#v\n", err, pizza)
}
