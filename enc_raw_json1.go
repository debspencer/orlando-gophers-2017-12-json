package main

import "bytes"
import "fmt"
import "encoding/json"

type Wrapper struct {
	Name       string
	JsonData   string
}

func main() {
	w := &Wrapper{
		Name: "Wikipedia",
		JsonData: `{"PageData":"<html>goes here</html>","Action":"update"}`,
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "    ")
	enc.Encode(&w)
	fmt.Println(buf.String())
}
