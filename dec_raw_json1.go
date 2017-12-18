package main

import "fmt"
import "encoding/json"

type Wrapper struct {
	Name       string
	JsonData   json.RawMessage
}

func main() {
    js := []byte(`
{
    "Name": "Wikipedia",
    "JsonData": {
        "PageData": "<html>goes here</html>",
        "Action": "update"
    }
}`)
    var w Wrapper
    err := json.Unmarshal(js, &w)
    fmt.Printf("err = %v %#v\n", err, w)
}
