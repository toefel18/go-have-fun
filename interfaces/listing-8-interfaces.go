package main

import (
"encoding/json"
"fmt"
)
// FROM LISTING 4

type message struct {
    Text   string
    Author string
}

//ToJson returns the message as JSON
func (msg message) ToJson() string {
    jsonMsg, _ := json.Marshal(msg)
    return string(jsonMsg)
}

type Jsonizer interface {
    ToJson() string
}

func printJson(something Jsonizer) {
    fmt.Println(something.ToJson())
}

func main() {
    printJson(message{Text: "Hi Gopher", Author: "Duke"})
}


