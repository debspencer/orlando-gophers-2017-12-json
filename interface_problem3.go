package main

import "encoding/json"
import "fmt"

func main() {
    js := []byte(`
{
    "colors": {
        "red": "#FF0000",
        "green": "#00FF00",
        "blue": "#0000FF"
    }
}`)

    m := make(map[string]interface{})
    json.Unmarshal(js, &m)
    for k, v := range m {
        switch k {
        case "colors":
            colors := v.(map[string]interface{})
            for k2, v2 := range colors {
                fmt.Printf("color %s = %s\n", k2, v2.(string))
            }
        }
    }
}
