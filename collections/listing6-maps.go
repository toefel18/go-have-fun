package main

import "fmt"

func main() {
	age := make(map[string]int)
	age["Christophe"] = 28
	age["Rob"] = 34

	for name, age := range age {
		fmt.Printf("%v has age %v\n", name, age)
	}
}
