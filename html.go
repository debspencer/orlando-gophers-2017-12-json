package main

import "bytes"
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
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "    ")
	enc.Encode(w)
	fmt.Println(buf.String())
}
