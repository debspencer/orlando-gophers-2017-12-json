package main

import "fmt"
import "encoding/json"

type Website struct {
	Name       string
	Url        string
}

func main() {
	w := &Website{
		Name: "Wikipedia",
		Url: "https://wikipedia.org/index.php?title=<Go>&action=edit",
	}
	js, _ := json.MarshalIndent(&w, "", "    ")
	fmt.Println(string(js))
}
