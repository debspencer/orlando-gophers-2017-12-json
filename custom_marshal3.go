package main

import "encoding/json"
import "fmt"

type Pizza struct {
	Name       string
	Size       string
}
type Alias Pizza

func (pizza *Pizza) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`{"myname":"%s","mysize":"%s"}`, pizza.Name, pizza.Size)

	return []byte(s), nil
}

func main() {
	pizza := Pizza{
		Name: "Pineapple Lovers",
		Size: "Large",
	}
	js, err := json.MarshalIndent(&pizza, "", "    ")
	fmt.Println(err, string(js))
}
