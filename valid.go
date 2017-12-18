package main

import "encoding/json"
import "fmt"

func main() {
	good := []byte(`
{
    "name": "Veggie Lovers",
    "size": "Large",
    "slices": 8,
    "toppings":["Peppers","Onions","Tomatos"]
}`)

	bad := []byte(`this is not valid json`)

	fmt.Println("Good = ", json.Valid(good))
	fmt.Println("Bad  = ", json.Valid(bad))
}
