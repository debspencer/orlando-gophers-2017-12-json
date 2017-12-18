package main

import "encoding/json"
import "fmt"

type Pizza struct {
	Name       string
	Size       string
}
type alias Pizza
func (p *Pizza) UnmarshalJSON(b []byte) error {
	var a alias
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	switch a.Size {
	case "Medium", "Large", "Extra-Large":
	default: return fmt.Errorf("Bad Pizza size: '%s'", a.Size)
	}
	*p = Pizza(a)
	return nil
}
func main() {
	js := []byte(`{"name": "Veggie Lovers","size": "Small"}`)

	var pizza Pizza
	err := json.Unmarshal(js, &pizza)
	fmt.Printf("err = %v\n%#v\n", err, pizza)
}
