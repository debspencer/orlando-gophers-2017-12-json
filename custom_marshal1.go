package main

import "encoding/json"
import "fmt"

type Pizza struct {
	Name       string
	Size       string
}
type Alias Pizza
func (pizza *Pizza) MarshalJSON() ([]byte, error) {
	switch pizza.Size {
	case "Medium", "Large", "Extra-Large":
	default: return nil, fmt.Errorf("Bad Pizza size: %s\n", pizza.Size)
	}
	a := Alias(*pizza)
	return json.Marshal(&a)
}
func main() {
	pizza := Pizza{
		Name: "Pineapple Lovers",
		Size: "Small" }
	js, err := json.MarshalIndent(&pizza, "", "    ")
	fmt.Println(err, string(js))
}
