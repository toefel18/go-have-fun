package main

import "fmt"

func main() {
	numbers := []int{1,2,3,4,5}
	print("all    ",  numbers)
	print("from  2 ", numbers[2:])
	print("until 2 ", numbers[:2])
	print("middle  ", numbers[2:3])
}
func print(name string, nrs []int) {
	for index, nr := range nrs {
		fmt.Println(name, index, ":", nr)
	}
}