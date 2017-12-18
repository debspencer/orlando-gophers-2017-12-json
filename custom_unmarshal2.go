package main

import "encoding/json"
import "fmt"
import "strings"

type PizzaSize string

type Pizza struct {
	Name       string
	Size       PizzaSize
}
func (sz *PizzaSize) UnmarshalJSON(b []byte) error {
	s := string(b)
	switch s {
	case `"Medium"`, `"Large"`, `"Extra-Large"`: s = strings.Replace(s, `"`, ``, -1)
	default: return fmt.Errorf("Bad Pizza size: '%s'", s)
	}
	*sz = PizzaSize(s)
	return nil
}
func main() {
	js := []byte(`{"name": "Veggie Lovers","size": "Large"}`)

	var pizza Pizza
	err := json.Unmarshal(js, &pizza)
	fmt.Printf("err = %v\n%#v\n", err, pizza)
}
