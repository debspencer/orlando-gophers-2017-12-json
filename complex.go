package main

import "encoding/json"
import "fmt"

type Toppings struct {
	Name       string
	Quantity   int
}
type Cheese struct {
	ChesseType  string
	Extra       bool
}
type Pizza struct {
	Name       string
	Size       string
	Whole      []Toppings  `json:",omitempty"`
	Sides      map[string][]Toppings
	Cheese     *Cheese     `json:",omitempty"`
	Cost       float64
}

func main() {
	pizza := Pizza{
		Name:       "Veggie Lovers",
		Size:       "Large",
		Whole:      []Toppings{ { Name: "Spinach", Quantity: 1 }, { Name: "Tomatos", Quantity: 1 } },
		Sides:      map[string][]Toppings{  "Side1": []Toppings{ { Name: "Onions",  Quantity: 2 } },
                                            "Side2": []Toppings{ { Name: "Peppers", Quantity: 1 } } },
		Cheese:     &Cheese{ ChesseType: "Mozzarella", Extra: true },
		Cost:       18.59,
	}
	js, _ := json.MarshalIndent(&pizza, "", "    ")
	fmt.Println(string(js))
}
