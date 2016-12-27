package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Text   string
	Author string
}

//ToJson returns the message as JSON
func (msg *message) ToJson() string {
	jsonMsg, _ := json.Marshal(msg)
	return string(jsonMsg)
}

func main() {
	defer fmt.Println("Closing! Bye!")
	fromChris := message{Text: "Hi Gopher", Author: "Christophe"}
	fromRob := message{"Hi Gopher", "Rob"}
	fmt.Println(fromRob.ToJson())
	printJson(fromChris)
}

type Jsonizer interface {
	ToJson() string
}

func printJson(something Jsonizer) {
	fmt.Println(something.ToJson())
}
