package main

import "bytes"
import "encoding/json"
import "fmt"
import "log"
import "io"

func main() {
    js := []byte(`{"name": "Veggie Lovers","size": "Large","slices": 8,
"toppings":["Peppers","Onions","Tomatos"]}`)
    dec := json.NewDecoder(bytes.NewReader(js))
    dec.UseNumber()
    for {
        t, err := dec.Token()
        if err == io.EOF { break }
        if err != nil { log.Fatal(err) }
        fmt.Printf("%T: %v\n", t, t)
        if dec.More() {
            reader := dec.Buffered()
            remaining := make([]byte, 1024)
            n, err := reader.Read(remaining)
            fmt.Printf(" -- %d bytes remaing, '%s' (%v)\n", n, string(remaining), err)
        }
    }
}
