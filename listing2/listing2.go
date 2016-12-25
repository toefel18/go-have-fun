package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Text   string
	Author string
}

//toJson returns the message as JSON
func (msg message) toJson() string {
	jsonMsg, _ := json.Marshal(msg)
	return string(jsonMsg)
}

func main() {
	fromRob := message{"Hi Gopher", "Rob"}
	fromChris := message{Text: "Hi Gopher", Author: "Christophe"}
	fmt.Println(fromRob.toJson())
	writeJson(fromChris)
}

func writeJson(something Jsonizable) {
	fmt.Println(something.toJson())
}

type Jsonizable interface {
	toJson() string
}
