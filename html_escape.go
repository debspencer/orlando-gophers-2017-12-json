package main

import "bytes"
import "encoding/json"
import "fmt"

func main() {
	html := []byte(`<a href="https://wikipedia.org/index.php?title=<Go>&action=edit">GO</a>`)

	var b bytes.Buffer
	json.HTMLEscape(&b, html)
	fmt.Println(string(html))
	fmt.Println(b.String())
}
