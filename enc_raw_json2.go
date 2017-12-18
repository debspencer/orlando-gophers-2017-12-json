package main

import "bytes"
import "fmt"
import "encoding/json"

type Wrapper struct {
	Name       string
	JsonData   json.RawMessage
}

func main() {
	w := &Wrapper{
		Name: "Wikipedia",
		JsonData: json.RawMessage(`{"PageData":"<html>goes here</html>","Action":"update"}`),
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "    ")
	enc.Encode(&w)
	fmt.Println(buf.String())
}
